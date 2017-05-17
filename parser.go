// The TableSet parser reads text and parses it into a set of one or more tables in a TableSet.

package gotable

/*
# 04/03/2016  Allow int format in float regular expressions.
*/

import (
	"fmt"
//	"os"
	"bufio"
//	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	//	"unsafe"
	"math"
)

/*
Copyright (c) 2017 Malcolm Gorman

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

func init() {
	log.SetFlags(log.Llongfile) // For var where
}

var where = log.Print

var globalLineNum int

const _ALL_SUBSTRINGS = -1

// Constants for strconv parse functions.
const (
	_DECIMAL = 10
	_BITS_8  = 8  // Bit width.
	_BITS_16 = 16 // Bit width.
	_BITS_32 = 32 // Bit width.
	_BITS_64 = 64 // Bit width.
)

// Compiled regular expressions.
// From http://stackoverflow.com/questions/249791/regex-for-quoted-string-with-escaping-quotes:  /"(?:[^"\\]|\\.)*"/
// var stringRegexp *regexp.Regexp = regexp.MustCompile("/"(?:[^"\\]|\\.)*"/")
var stringRegexp *regexp.Regexp = regexp.MustCompile(`^"(?:[^"\\]*(?:\\.)?)*"`)
var boolRegexp *regexp.Regexp = regexp.MustCompile(`^\b(true|false)\b`)
var uintRegexp *regexp.Regexp = regexp.MustCompile(`^[+]?\b\d+\b`)
var intRegexp *regexp.Regexp = regexp.MustCompile(`^[-+]?\b\d+\b`)

// Allow negative float numbers! 15/03/2017 Amazed that this was missed during initial testing!
// var floatRegexp		*regexp.Regexp = regexp.MustCompile(`(^[-+]?(\b[0-9]+\.([0-9]+\b)?|\b\.[0-9]+\b))|([Nn][Aa][Nn])|(\b[-+]?\d+\b)`)
// From Regular Expressions Cookbook.
var floatRegexp *regexp.Regexp = regexp.MustCompile(`^([-+]?([0-9]+(\.[0-9]*)?|\.[0-9]+)([eE][-+]?[0-9]+)?)|([Nn][Aa][Nn])`)
var namePattern = `^[a-zA-Z_][a-zA-Z0-9_]*$`
var tableNamePattern = `^\[[a-zA-Z_][a-zA-Z0-9_]*\]$`
var tableNameRegExp *regexp.Regexp = regexp.MustCompile(tableNamePattern)
var colNameRegExp *regexp.Regexp = regexp.MustCompile(namePattern)
var whiteRegExp *regexp.Regexp = regexp.MustCompile(`\s+`)
var equalsRegExp *regexp.Regexp = regexp.MustCompile(`=`)

// Oct regular expression (for integral types)
// Hex regular expression (for integral types)

/*
type DataType int
const (
	bool		DataType = iota
	byte        DataType = uint8 // alias for uint8
	complex128  DataType = iota	 // The set of all complex numbers with float64 real and imaginary parts
	complex64   DataType = iota	 // The set of all complex numbers with float32 real and imaginary parts
	float32     DataType = iota	 // The set of all IEEE-754 32-bit floating-point numbers
	float64     DataType = iota	 // The set of all IEEE-754 64-bit floating-point numbers
	int16       DataType = iota	 // The set of all signed 16-bit integers (-32768 to 32767)
	int32       DataType = iota	 // The set of all signed 32-bit integers (-2147483648 to 2147483647)
	int64       DataType = iota	 // The set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	int8        DataType = iota	 // The set of all signed  8-bit integers (-128 to 127)
	rune        DataType = int32 // alias for int32
	string		DataType = iota
	uint16      DataType = iota	 // The set of all unsigned 16-bit integers (0 to 65535)
	uint32      DataType = iota	 // The set of all unsigned 32-bit integers (0 to 4294967295)
	uint64      DataType = iota	 // The set of all unsigned 64-bit integers (0 to 18446744073709551615)
	uint8       DataType = iota	 // The set of all unsigned  8-bit integers (0 to 255)

)
*/

var globalColTypesMap = map[string]int{
	"bool":    0,
	"float32": 0,
	"float64": 0,
	"int":     0,
	"int16":   0,
	"int32":   0,
	"int64":   0,
	"int8":    0,
	"string":  0,
	"uint":    0,
	"uint16":  0,
	"uint32":  0,
	"uint64":  0,
	"uint8":   0,
}

var globalNumericColTypesMap = map[string]int{
	"float32": 0,
	"float64": 0,
	"int":     0,
	"int16":   0,
	"int32":   0,
	"int64":   0,
	"int8":    0,
	"uint":    0,
	"uint16":  0,
	"uint32":  0,
	"uint64":  0,
	"uint8":   0,
}

func (p *parser) parseString(s string) (*TableSet, error) {
	type _TableSection string
	const (
		_TABLE_NAME _TableSection = "_TABLE_NAME"
		_COL_NAMES  _TableSection = "_COL_NAMES"
		_COL_TYPES  _TableSection = "_COL_TYPES"
		_COL_ROWS   _TableSection = "_COL_ROWS"
	)
	var expecting _TableSection = _TABLE_NAME

	type _TableShape string
	const (
		_TABLE_SHAPE     _TableShape = "_TABLE_SHAPE"
		_STRUCT_SHAPE    _TableShape = "_STRUCT_SHAPE"
		_UNDEFINED_SHAPE _TableShape = "_UNDEFINED_SHAPE"
	)
	var tableShape _TableShape = _UNDEFINED_SHAPE
	// Note: tableShape variable is used for parsing. Not sure it's needed.
	// Note: It's not worth the trouble of printing a table as a struct.
	// Let's give it a try ... 29/03/2017

	var structHasRowData bool

	var parserColNames []string
	var parserColTypes []string
	var rowMapOfStruct tableRow // Needs to persist over multiple lines.

	unnamed := ""
	tables, err := NewTableSet(unnamed)
	if err != nil {
		return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
	}

	var table *Table

	var line string
	var readError error
	var stringReader *strings.Reader = strings.NewReader(s)
	var inputReader *bufio.Reader = bufio.NewReader(stringReader)
	globalLineNum = 1
	for ; ; globalLineNum++ {
		line, readError = inputReader.ReadString('\n')
		line = strings.TrimSpace(line)

		// Skip commented lines.
		if len(line) > 0 && line[0] == '#' {
			continue
		}

		if len(line) == 0 {
			if expecting == _COL_TYPES {
				return nil, fmt.Errorf("%s expecting col names to be followed by a line of col types", p.gotFilePos())
				// A blank line is okay after table name, col types or row values, but col names must always have col types.
			}
			expecting = _TABLE_NAME
			tableShape = _UNDEFINED_SHAPE
			// The only other place there is a need to test len(line) is in case _TABLE_NAME.
		}

		switch expecting {

		case _TABLE_NAME:

			if len(line) > 0 {
				// Skip blank lines before table.
				var tableName string
				tableName, err = p.getTableName(line)
				if err != nil {
					//						return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
					return nil, err
				}
				table, err = NewTable(tableName)
				if err != nil {
					return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
				}
				tableShape = _UNDEFINED_SHAPE
				// Add this table to tables. Do it immediately to allow empty tables. 02.08.2016
				err = tables.AppendTable(table)
				if err != nil {
					return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
				}
				expecting = _COL_NAMES
			}

		case _COL_NAMES: // Also proxy for a line of a table struct in the form: name type = value

			// EITHER (1) read a line of a table struct OR (2) read col names of a tabular table.
			var lineSplit []string = whiteRegExp.Split(line, _ALL_SUBSTRINGS)
			const structNameIndex = 0
			const structTypeIndex = 1
			const structEqualsIndex = 2
			const tokenCountForNameType = 2					// (a) name type
			var isNameTypeStruct bool						// (a) name type
			const minTokenCountForNameTypeEqualsValue = 4	// (b) name type = value
			var isNameTypeEqualsValueStruct bool			// (b) name type = value

			// This is a recognition step.
			// Determine whether this is a candidate struct of either:
			// (a) name type
			// (b) name type = value

			// Note: strings can mean len(lineSplit) is > 4 but still valid. So can't just test for exactly 4.
			var lenLineSplit int = len(lineSplit)
			var looksLikeStructShape bool
			if lenLineSplit != tokenCountForNameType && lenLineSplit < minTokenCountForNameTypeEqualsValue {
				looksLikeStructShape = false
			} else if lenLineSplit == 2 {
				secondTokenIsType, _ := IsValidColType(lineSplit[structTypeIndex])
				if secondTokenIsType {
					isNameTypeStruct = true
					looksLikeStructShape = true
				}
			} else {	// lenLineSplit must be >= 4
				secondTokenIsType, _ := IsValidColType(lineSplit[structTypeIndex])
				if secondTokenIsType && lineSplit[structEqualsIndex] == "=" {
					isNameTypeEqualsValueStruct = true
					looksLikeStructShape = true
				}
			}

			if looksLikeStructShape {

				// (1) Get the table struct (name, type and optional equals value) of this line.

				tableShape = _STRUCT_SHAPE
				table.structShape = true
				var colName string = lineSplit[structNameIndex]
				var colType string = lineSplit[structTypeIndex]
				var isValid bool
				if isValid, err = IsValidColName(colName); !isValid {
					// return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
					return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
				}
				var colNameSlice []string = []string{colName}
				var colTypeSlice []string = []string{colType}

				err = table.AppendCol(colName, colType)
				if err != nil {
					return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
				}

				// where(p.gotFilePos())
				// Set this only once (for each table). Base on the first "col", which is <name> <type> = <value>|no-value
				if table.ColCount() == 1 {	// The first struct item.
					structHasRowData = isNameTypeEqualsValueStruct
					// where(fmt.Sprintf("%s: setting structHasRowData = %t", p.gotFilePos(), structHasRowData))
				}
				// where(fmt.Sprintf("%s: structHasRowData = %t", p.gotFilePos(), structHasRowData))

				if structHasRowData && isNameTypeStruct {
					return nil, fmt.Errorf("%s expecting: %s %s = <value> but found: %s %s",
						p.gotFilePos(), colName, colType, lineSplit[0], lineSplit[1])
				}

/* Unreachable because len(lineSplit) here can only be 2 or 4.
				if !structHasRowData && len(lineSplit) > tokenCountForNameType {
					// An approximate way to construct text for the error message. Spacing may be different.
					var remaining string
					for i := 3; i < len(lineSplit); i++ {
						remaining += lineSplit[i] + " "
					}
					return nil, fmt.Errorf("%s expecting 0 values after = but found: %s", p.gotFilePos(), remaining)
				}
*/

				if structHasRowData {
					// where(fmt.Sprintf("%s: if structHasRowData = %t", p.gotFilePos(), structHasRowData))
					// Find the equals sign byte location within the string so we can locate the value data after equals.
					// We know it's there (from the line split above), so don't check for nil returned.
					var rangeFound []int = equalsRegExp.FindStringIndex(line)
					var valueData string = line[rangeFound[1]:]        // Just after = equals sign.
					valueData = strings.TrimLeft(valueData, " \t\r\n") // Remove leading space.
//where()
					rowMapOfStruct, err = p.getRowData(valueData, colNameSlice, colTypeSlice)
					// where(fmt.Sprintf("len(lineSplit) = %d", len(lineSplit)))
					if err != nil {
						//							return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
						// where(err)
						return nil, err
					}
//where()
					// Still expecting _COL_NAMES which is where we find struct: name type = value

					// Handle the first iteration (parse a line) through a struct, where the table has no rows.
					// Exactly one row is needed for a struct table.
					if table.RowCount() == 0 {
						err = table.appendRowOfNil()
						if err != nil {
							return nil, err
						}
					}

					var val interface{} = rowMapOfStruct[colName]
					err = table.SetVal(colName, 0, val)
					if err != nil {
						return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
					}
				}
			} else {
				if tableShape == _STRUCT_SHAPE {
					return nil, fmt.Errorf("%s expecting more structs ( <name> <type> = <value> ) but found: %s", p.gotFilePos(), line)
				}

				tableShape = _TABLE_SHAPE

				// (2) Get the col names.

				parserColNames, err = p.getColNames(lineSplit)
				if err != nil {
					return nil, err
				}
				err = table.appendColNames(parserColNames)
				if err != nil {
					return nil, err
				}
				expecting = _COL_TYPES
			}

		case _COL_TYPES:

			parserColTypes, err = p.getColTypes(line)
			if err != nil {
				return nil, err
			}
			lenColNames := len(parserColNames)
			lenColTypes := len(parserColTypes)
			if lenColTypes != lenColNames {
				return nil, fmt.Errorf("%s expecting: %d col type%s but found: %d", p.gotFilePos(), lenColNames, plural(lenColNames), lenColTypes)
			}
			err = table.appendColTypes(parserColTypes)
			if err != nil {
				return nil, err
			}
			expecting = _COL_ROWS

		case _COL_ROWS:

			// Found data.
			var rowMap tableRow
			rowMap, err = p.getRowData(line, parserColNames, parserColTypes)
			if err != nil {
				//					return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
				return nil, err
			}
			lenColTypes := len(parserColTypes)
			lenRowMap := len(rowMap)
			if lenColTypes != lenRowMap {
				return nil, fmt.Errorf("%s expecting: %d value%s but found: %d", p.gotFilePos(), lenColTypes, plural(lenColTypes), lenRowMap)
			}
			err = table.appendRowMap(rowMap)
			if err != nil {
				return tables, err
			}
		}

		if readError == io.EOF {
			return tables, nil // It's not an error to reach EOF. It just means end of document.
		}
	}
}

func (p *parser) parseFile(fileName string) (*TableSet, error) {
	var err error
	var fileBytes []byte

	p.SetFileName(fileName) // For file and line diagnostics.

	fileBytes, err = ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	tables, err := p.parseString(string(fileBytes))
	if err != nil {
		return nil, err
	}

	tables.SetFileName(fileName)

	return tables, err
}

func (p *parser) gotFilePos() string {
	return fmt.Sprintf("%s[%d]", p.fileName, globalLineNum)
}

func file_line() string {
	_, fileName, fileLine, ok := runtime.Caller(2)
	var s string
	if ok {
		s = fmt.Sprintf("%s:%d", fileName, fileLine)
	} else {
		s = ""
	}
	return s
}

func (p *parser) getTableName(line string) (string, error) {

	fields := strings.Fields(line)
	if len(fields) != 1 {
		return "", fmt.Errorf("%s expecting a table name (without spaces) but found: %s", p.gotFilePos(), line)
	}
	tableName := fields[0]

	result := tableNameRegExp.MatchString(tableName)
	if !result {
		return "", fmt.Errorf("%s expecting a valid alpha-numeric table name, eg [_Foo2Bar3] but found: %s", p.gotFilePos(), tableName)
	}

	// Strip off surrounding []
	tableName = tableName[1 : len(tableName)-1]
	return tableName, nil
}

func (p *parser) getColNames(colNames []string) ([]string, error) {
	for i := 0; i < len(colNames); i++ {
		isValid, err := IsValidColName(colNames[i])
		if !isValid {
			if i == 1 {
				_, contains := globalColTypesMap[colNames[1]]
				if contains {
					return nil, fmt.Errorf("%s %s did you perhaps mean either: %s %s or %s %s = <val>",
						p.gotFilePos(), err, colNames[0], colNames[1], colNames[0], colNames[1])
				} else {
					return nil, fmt.Errorf("%s %s", p.gotFilePos(), err) // Default error.
				}
			} else {
				return nil, fmt.Errorf("%s %s", p.gotFilePos(), err) // Default error.
			}
		}
	}

	return colNames, nil
}

func (p *parser) getColTypes(line string) ([]string, error) {

	var colTypes []string = whiteRegExp.Split(line, _ALL_SUBSTRINGS)
	if len(colTypes) == 0 {
		return nil, fmt.Errorf("%s expecting col types", p.gotFilePos())
	}

	for i := 0; i < len(colTypes); i++ {
		valid, err := IsValidColType(colTypes[i])
		if !valid {
			//			log.Printf("%s", msg)	// To print the source file name and line number.
			msg := fmt.Sprintf("%s %s", p.gotFilePos(), err)
			return nil, errors.New(msg)
		}
	}

	return colTypes, nil
}

/*
Returns true for those Go types that Table supports.

Go types NOT (yet) supported: complex64 complex128 byte rune
*/
func IsValidColType(colType string) (bool, error) {
	_, contains := globalColTypesMap[colType]
	if !contains {
		msg := fmt.Sprintf("invalid col type: %s (Valid types:", colType)
		// Note: Because maps are not ordered, this (desirably) shuffles the order of valid col types with each call.
		for typeName, _ := range globalColTypesMap {
			msg += fmt.Sprintf(" %s", typeName)
		}
		msg += ")"
		err := errors.New(msg)
		return false, err
	}
	return true, nil
}

/*
Returns true for those Go types that are numeric.

Go types NOT (yet) supported: complex64 complex128 byte rune
*/
func IsNumericColType(colType string) (bool, error) {
	_, contains := globalNumericColTypesMap[colType]
	if !contains {
		msg := fmt.Sprintf("Non-numeric col type: %s (Numeric types:", colType)
		// Note: Because maps are not ordered, this (desirably) shuffles the order of valid col types with each call.
		for typeName, _ := range globalNumericColTypesMap {
			msg += fmt.Sprintf(" %s", typeName)
		}
		msg += ")"
		err := errors.New(msg)
		return false, err
	}
	return true, nil
}

func IsValidColName(colName string) (bool, error) {

	result := colNameRegExp.MatchString(colName)
	if !result {
		return false, fmt.Errorf("invalid col name: %q (Valid example: \"_Foo2Bar3\")", colName)
	}

	_, contains := globalColTypesMap[colName]
	if contains {
		return false, fmt.Errorf("invalid col name: %s (cannot use Go types as col names)", colName)
	}

	return true, nil
}

func IsValidTableName(tableName string) (bool, error) {
	// Same regular expression as table name without square brackets.
	result := colNameRegExp.MatchString(tableName)
	if !result {
		return false, fmt.Errorf("invalid table name: %q (Valid example: \"_Foo1Bar2\")", tableName)
	}
	return true, nil
}

func (p *parser) getRowData(line string, colNames, colTypes []string) (tableRow, error) {

	var err error
	rowMap := make(tableRow)

	remaining := line // Remainder of line left to parse.
	var rangeFound []int
	var textFound string
	var colCount = 0
	var lenColTypes = len(colTypes)
	var i int

	var stringVal string
	var boolVal bool
	var uint8Val uint8
	var uint16Val uint16
	var uint32Val uint32
	var uint64Val uint64
	var uintVal uint
	var intVal int
	var int8Val int8
	var int16Val int16
	var int32Val int32
	var int64Val int64
	var float32Val float32
	var float64Val float64

	for i = 0; i < lenColTypes; i++ {
		if len(remaining) == 0 { // End of line
			return nil, fmt.Errorf("%s expecting %d value%s but found only %d", p.gotFilePos(), lenColTypes, plural(lenColTypes), colCount)
		}
		switch colTypes[i] {
		case "string":
			rangeFound = stringRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s expecting a valid value of type %s but found: %s", p.gotFilePos(), colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			stringVal = textFound[1 : len(textFound)-1] // Strip off leading and trailing "" quotes.
			rowMap[colNames[i]] = stringVal

		case "bool":
			rangeFound = boolRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s expecting a valid value of type %s but found: %s", p.gotFilePos(), colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			boolVal, err = strconv.ParseBool(textFound)
			if err != nil { // This error check probably redundant.
				return nil, fmt.Errorf("%s %s for type %s", p.gotFilePos(), err, colTypes[i])
			}
			rowMap[colNames[i]] = boolVal
		case "uint8":
			rangeFound = uintRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s expecting a valid value of type %s but found: %s", p.gotFilePos(), colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			uint64Val, err = strconv.ParseUint(textFound, _DECIMAL, _BITS_8)
			if err != nil {
				rangeMsg := rangeForIntegerType(0, math.MaxUint8)
				return nil, fmt.Errorf("%s %s for type %s %s", p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			uint8Val = uint8(uint64Val)
			rowMap[colNames[i]] = uint8Val
		case "uint16":
			rangeFound = uintRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s expecting a valid value of type %s but found: %s", p.gotFilePos(), colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			uint64Val, err = strconv.ParseUint(textFound, _DECIMAL, _BITS_16)
			if err != nil {
				rangeMsg := rangeForIntegerType(0, math.MaxUint16)
				return nil, fmt.Errorf("%s %s for type %s %s", p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			uint16Val = uint16(uint64Val)
			rowMap[colNames[i]] = uint16Val
		case "uint32":
			rangeFound = uintRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s expecting a valid value of type %s but found: %s", p.gotFilePos(), colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			uint64Val, err = strconv.ParseUint(textFound, _DECIMAL, _BITS_32)
			if err != nil {
				rangeMsg := rangeForIntegerType(0, math.MaxUint32)
				return nil, fmt.Errorf("%s %s for type %s %s", p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			uint32Val = uint32(uint64Val)
			rowMap[colNames[i]] = uint32Val
		case "uint64":
			rangeFound = uintRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s expecting a valid value of type %s but found: %s", p.gotFilePos(), colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			uint64Val, err = strconv.ParseUint(textFound, _DECIMAL, _BITS_64)
			if err != nil {
				rangeMsg := rangeForIntegerType(0, math.MaxUint64)
				return nil, fmt.Errorf("%s %s for type %s %s", p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			rowMap[colNames[i]] = uint64Val
		case "uint":
			rangeFound = uintRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s expecting a valid value of type %s but found: %s", p.gotFilePos(), colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			intBits := strconv.IntSize // uint and int are the same size.
			uint64Val, err = strconv.ParseUint(textFound, _DECIMAL, intBits)
			if err != nil {
				var minVal int64
				var maxVal uint64
				switch intBits {
				case 32:
					minVal = 0
					maxVal = math.MaxUint32
				case 64:
					minVal = 0
					maxVal = math.MaxUint64
				default:
					msg := fmt.Sprintf("CHECK int or uint ON THIS SYSTEM: Unknown int size: %d bits", intBits)
					log.Printf("%s", msg)
					return nil, fmt.Errorf("%s", msg)
				}
				rangeMsg := rangeForIntegerType(minVal, maxVal)
				return nil, fmt.Errorf("%s %s for type %s %s", p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			uintVal = uint(uint64Val) // May be unnecessary.
			rowMap[colNames[i]] = uintVal
		case "int":
			rangeFound = intRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s expecting a valid value of type %s but found: %s", p.gotFilePos(), colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			intBits := strconv.IntSize
			int64Val, err = strconv.ParseInt(textFound, _DECIMAL, intBits)
			if err != nil {
				var minVal int64
				var maxVal uint64
				switch intBits {
				case 32:
					minVal = math.MinInt32
					maxVal = math.MaxInt32
				case 64:
					minVal = math.MinInt64
					maxVal = math.MaxInt64
				default:
					msg := fmt.Sprintf("CHECK int ON THIS SYSTEM: Unknown int size: %d bits", intBits)
					log.Printf("%s", msg)
					return nil, fmt.Errorf("%s", msg)
				}
				rangeMsg := rangeForIntegerType(minVal, maxVal)
				return nil, fmt.Errorf("%s %s for type %s %s", p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			intVal = int(int64Val) // May be unnecessary.
			rowMap[colNames[i]] = intVal
		case "int8":
			rangeFound = intRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s expecting a valid value of type %s but found: %s", p.gotFilePos(), colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			int64Val, err = strconv.ParseInt(textFound, _DECIMAL, _BITS_8)
			if err != nil {
				// Example: data.got[55] strconv.ParseInt: parsing "-129": value out of range for type int8
				rangeMsg := rangeForIntegerType(math.MinInt8, math.MaxInt8)
				return nil, fmt.Errorf("%s %s for type %s %s", p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			int8Val = int8(int64Val)
			rowMap[colNames[i]] = int8Val
		case "int16":
			rangeFound = intRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s expecting a valid value of type %s but found: %s", p.gotFilePos(), colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			int64Val, err = strconv.ParseInt(textFound, _DECIMAL, _BITS_16)
			if err != nil {
				rangeMsg := rangeForIntegerType(math.MinInt16, math.MaxInt16)
				return nil, fmt.Errorf("%s %s for type %s %s", p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			int16Val = int16(int64Val)
			rowMap[colNames[i]] = int16Val
		case "int32":
			rangeFound = intRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s expecting a valid value of type %s but found: %s", p.gotFilePos(), colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			int64Val, err = strconv.ParseInt(textFound, _DECIMAL, _BITS_32)
			if err != nil {
				rangeMsg := rangeForIntegerType(math.MinInt32, math.MaxInt32)
				return nil, fmt.Errorf("%s %s for type %s%s ", p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			int32Val = int32(int64Val)
			rowMap[colNames[i]] = int32Val
		case "int64":
			rangeFound = intRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s expecting a valid value of type %s but found: %s", p.gotFilePos(), colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			int64Val, err = strconv.ParseInt(textFound, _DECIMAL, _BITS_64)
			if err != nil {
				rangeMsg := rangeForIntegerType(math.MinInt64, math.MaxInt64)
				return nil, fmt.Errorf("%s %s for type %s %s", p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			rowMap[colNames[i]] = int64Val
		case "float32":
			rangeFound = floatRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s expecting a valid value of type %s but found: %s", p.gotFilePos(), colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			float64Val, err = strconv.ParseFloat(textFound, _BITS_32)
			if err != nil {
				return nil, fmt.Errorf("%s %s for type %s", p.gotFilePos(), err, colTypes[i])
			}
			if math.IsNaN(float64Val) && textFound != "NaN" {
				//					return nil, fmt.Errorf("%s expecting NaN as Not-a-Number for type %s but found: %s ", p.gotFilePos(), colTypes[i], textFound)
				return nil, fmt.Errorf("%s col %s: expecting NaN as Not-a-Number for type %s but found: %s ",
					p.gotFilePos(), colNames[i], colTypes[i], textFound)
			}
			float32Val = float32(float64Val)
			rowMap[colNames[i]] = float32Val
		case "float64":
			rangeFound = floatRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s expecting a valid value of type %s but found: %s", p.gotFilePos(), colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			//				where(fmt.Sprintf("textFound: %s", textFound))
			float64Val, err = strconv.ParseFloat(textFound, _BITS_64)
			if err != nil {
				return nil, fmt.Errorf("%s %s for type %s", p.gotFilePos(), err, colTypes[i])
			}
			//				where(fmt.Sprintf("float64Val: %f", float64Val))
			if math.IsNaN(float64Val) && textFound != "NaN" {
				//					return nil, fmt.Errorf("%s expecting NaN as Not a Number for type %s but found: %s ", p.gotFilePos(), colTypes[i], textFound)
				return nil, fmt.Errorf("%s col %s: expecting NaN as Not-a-Number for type %s but found: %s ",
					p.gotFilePos(), colNames[i], colTypes[i], textFound)
			}
			rowMap[colNames[i]] = float64Val
		default:
			log.Printf("Unreachable code in getRowCol()") // Need to define another type?
			return nil, fmt.Errorf("%s Unreachable code in getRowCol()", p.gotFilePos())
		}
//		remaining = remaining[rangeFound[1]:len(remaining)]
		remaining = remaining[rangeFound[1]:]
		remaining = strings.TrimLeft(remaining, " \t\r\n") // Remove leading spaces.
		colCount++
	}

	if len(remaining) > 0 { // Still one or more columns to parse.
		// This handles both table shape and struct shape columns.
		return nil, fmt.Errorf("%s expecting %d value%s but found more: %s", p.gotFilePos(), lenColTypes, plural(lenColTypes), remaining)

	}

	return rowMap, nil
}

func rangeForIntegerType(min int64, max uint64) string {
	return fmt.Sprintf("(%d to %d)", min, max)
}

// parser definition: fields and methods.
type parser struct {
	fileName string // Needed for printing file and line diagnostics.
}

// Needed for printing file and line diagnostics.
func (p *parser) SetFileName(fileName string) {
	p.fileName = fileName
}

/*
func (p parser) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("table.parser: ")
	buffer.WriteString(fmt.Sprintf("fileName='%s' ", p.fileName))
	return buffer.String()
}
*/

/*
func (p *parser) Parse() (*TableSet, error) {
	return p.parseFile(p.fileName)
}
*/

func plural(items int) string {
	if items == 1 || items == -1 {
		// Singular
		return ""
	} else {
		// Plural
		return "s"
	}
}
