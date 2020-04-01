// The TableSet parser reads text and parses it into a set of one or more tables in a TableSet.

package gotables

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
	"os"
	//	"path"
	"math"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
	"unicode"
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

// true means decimal int and also bin oct hex literals
// false means decimal int literals only
const go_1_13_number_literals = true

func init() {
	// Note: map items need to go both ways.
	typeAliasMap = map[string]string{
		"uint8":   "byte",
		"byte":    "uint8",
		"[]uint8": "[]byte",
		"[]byte":  "[]uint8",
		"int32":   "rune",
		"rune":    "int32",
		//		"[]int32" : "[]rune",	// Proposed?
		"*Table":          "*gotables.Table",
		"*gotables.Table": "*Table",
	}

	if go_1_13_number_literals {
		// int with bin oct hex literals
		// Same as uint but adds minus (-) sign for decimals.
		intRegexpString = `^((0[bB])[0-1]+|(0[oO])[0-7]+|(0[xX])[0-9A-Fa-f]+|[+-]?\d+)`

		// uint with bin oct hex literals
		// Without ^ so we can use uintRegexpString in uintSliceRegexpString
		uintRegexpString = `((0[bB])[0-1]+|(0[oO])[0-7]+|(0[xX])[0-9A-Fa-f]+|[+]?\d+)`
	} else {
		// int with dec only literals
		intRegexpString = `^[-+]?\b\d+\b`

		// uint with dec only literals
		// Without ^ so we can use uintRegexpString in uintSliceRegexpString
		uintRegexpString = `[+]?\b\d+\b`
	}

	intRegexp = regexp.MustCompile(intRegexpString)
	uintRegexp = regexp.MustCompile(fmt.Sprintf(`^%s`, uintRegexpString)) // Prepend ^

	// Handles: [] [num] [num num]
	uintSliceRegexpString = fmt.Sprintf(`^\[(%s)*(\s%s)*\]`, uintRegexpString, uintRegexpString)
	uintSliceRegexp = regexp.MustCompile(uintSliceRegexpString)
}

var globalLineNum int

const globalErrorLimit int = 10

var globalErrorCount int

const _ALL_SUBSTRINGS = -1

// Constants for strconv parse functions.
const (
	_DEC     = 10
	_BIN     = 2
	_OCT     = 8
	_HEX     = 16
	_BITS_8  = 8  // Bit width.
	_BITS_16 = 16 // Bit width.
	_BITS_32 = 32 // Bit width.
	_BITS_64 = 64 // Bit width.
)

// Compiled regular expressions.
// From http://stackoverflow.com/questions/249791/regex-for-quoted-string-with-escaping-quotes:  /"(?:[^"\\]|\\.)*"/
// 25/07/2018: See also: [^"\\\r\n]*(?:\\.[^"\\\r\n]*)* from https://www.regular-expressions.info/examplesprogrammer.html
// var stringRegexp *regexp.Regexp = regexp.MustCompile("/"(?:[^"\\]|\\.)*"/")
var stringRegexp *regexp.Regexp = regexp.MustCompile(`^"(?:[^"\\]*(?:\\.)?)*"`)

var boolRegexp *regexp.Regexp = regexp.MustCompile(`^\b(true|false)\b`)

// Note: (\\') successfully parses '\'' It needs to go before ([^']*)
//             This may not be terribly efficient, awaiting a more specific regular expresssion.
const runeRegexpString string = `'((\\n)|(\\')|([^']\\[xuU].*)|([^']*))'`

var runeRegexp *regexp.Regexp = regexp.MustCompile(runeRegexpString)

// Covers all unsigned integrals, including byte.
// See init()
var intRegexpString string
var uintRegexpString string
var uintRegexp *regexp.Regexp
var uintSliceRegexpString string
var uintSliceRegexp *regexp.Regexp
var intRegexp *regexp.Regexp

// See init()

// Allow negative float numbers! 15/03/2017 Amazed that this was missed during initial testing!
// var floatRegexp		*regexp.Regexp = regexp.MustCompile(`(^[-+]?(\b[0-9]+\.([0-9]+\b)?|\b\.[0-9]+\b))|([Nn][Aa][Nn])|(\b[-+]?\d+\b)`)
// From Regular Expressions Cookbook.
var floatRegexp *regexp.Regexp = regexp.MustCompile(`^([-+]?([0-9]+(\.[0-9]*)?|\.[0-9]+)([eE][-+]?[0-9]+)?)|([Nn][Aa][Nn])`)

const tableSetNamePattern string = `^\[\[[a-zA-Z_][a-zA-Z0-9_]*\]\]$`
const namePattern string = `^[a-zA-Z_][a-zA-Z0-9_]*$`
const tableNamePattern string = `^\[[a-zA-Z_][a-zA-Z0-9_]*\]$`
const tableNameNilPattern string = `^(\[\])`

// From: https://golang.org/pkg/time
// RFC3339     = "2006-01-02T15:04:05Z07:00"
// const rfc3339TimePattern string = `\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(Z|([+-]\d{2}:\d{2}))`
// RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
const rfc3339TimePattern string = `\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|([+-]\d{2}:\d{2}))`

var tableSetNameRegexp *regexp.Regexp = regexp.MustCompile(tableSetNamePattern)
var tableNameRegexp *regexp.Regexp = regexp.MustCompile(tableNamePattern)
var tableNameNilRegexp *regexp.Regexp = regexp.MustCompile(tableNameNilPattern)
var colNameRegexp *regexp.Regexp = regexp.MustCompile(namePattern)
var whiteRegexp *regexp.Regexp = regexp.MustCompile(`\s+`)
var equalsRegexp *regexp.Regexp = regexp.MustCompile(`=`)
var rfc3339TimeRegexp *regexp.Regexp = regexp.MustCompile(rfc3339TimePattern)

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

var typeAliasMap map[string]string

// GO_TYPES
// Keep in sync with helpersmain/helpersmain.go globalColTypesMap
var globalColTypesMap = map[string]int{
	"[]byte":  0,
	"[]uint8": 0,
	"bool":    0,
	"byte":    0,
	"float32": 0,
	"float64": 0,
	"int":     0,
	"int16":   0,
	"int32":   0,
	"rune":    0,
	"int64":   0,
	"int8":    0,
	"string":  0,
	"uint":    0,
	"uint16":  0,
	"uint32":  0,
	"uint64":  0,
	"uint8":   0,
	//	"complex":         0, // Not yet implemented.
	"*Table":          0,
	"*gotables.Table": 0,
	"time.Time":       0,
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
	"byte":    0,
	"complex": 0, // Not yet implemented.
}

var globalSliceColTypesMap = map[string]int{
	"[]byte":  0,
	"[]uint8": 0,
}

var globalTableColTypesMap = map[string]int{
	"*Table":          0,
	"*gotables.Table": 0,
}

const structNameIndex = 0
const structTypeIndex = 1
const structEqualsIndex = 2

type _TableSection string

const (
	_TABLE_NAME _TableSection = "_TABLE_NAME"
	_COL_NAMES  _TableSection = "_COL_NAMES"
	_COL_TYPES  _TableSection = "_COL_TYPES"
	_COL_ROWS   _TableSection = "_COL_ROWS"
)

type _TableShape string

const (
	_TABLE_SHAPE     _TableShape = "_TABLE_SHAPE"
	_STRUCT_SHAPE    _TableShape = "_STRUCT_SHAPE"
	_UNDEFINED_SHAPE _TableShape = "_UNDEFINED_SHAPE"
)

func (p *parser) parseString(s string) (*TableSet, error) {
	// Note: tableShape variable is used for parsing. Not sure it's needed.
	// Note: It's not worth the trouble of printing a table as a struct.
	// Let's give it a try ... 29/03/2017

	var expecting _TableSection = _TABLE_NAME
	var tableShape _TableShape = _UNDEFINED_SHAPE

	var structHasRowData bool

	var parserColNames []string
	var parserColTypes []string
	var rowSliceOfStructTable tableRow // Needs to persist over multiple lines.

	var tableSetNameHasBeenSet bool = false
	unnamedTableSet := ""
	tables, err := NewTableSet(unnamedTableSet)
	if err != nil {
		return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
	}

	var table *Table // Reused with each new table.

	// Has to be initialised with each call to parseString() because
	// the first thing we always expect is a table name.
	expecting = _TABLE_NAME

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
				/*
					A blank line is okay after any of
						(1) table name,
						(2) col types or
						(3) row values,
					but a row of col names must always have (be followed by) a row of col types.
				*/
				return nil, fmt.Errorf("%s expecting row of col names to be followed by a row of col types", p.gotFilePos())
			}
			expecting = _TABLE_NAME
			// The only other place there is a need to test len(line) is in case _TABLE_NAME.
		}

		var lineSplit []string = whiteRegexp.Split(line, _ALL_SUBSTRINGS)

		switch expecting {

		case _TABLE_NAME:

			if len(line) > 0 {
				// Skip blank lines before table.

				if tableSetNameHasBeenSet == false { // Test is for efficiency (not bench tested).
					// Try to get tableSetName here.
					tableSetName, err := p.getTableSetName(line)
					if err == nil { // No error means: got a TableSet name
						err = tables.SetName(tableSetName)
						if err != nil {
							return nil, err
						}

						tableSetNameHasBeenSet = true

						continue
					}
				}

				var tableName string
				tableName, err = p.getTableName(line)
				if err != nil {
					return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
				}

				table, err = NewTable(tableName)
				if err != nil {
					return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
				}

				// Add this table to tables. Do it immediately to allow empty tables. 02.08.2016
				err = tables.AppendTable(table)
				if err != nil {
					return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
				}

				tableShape = _UNDEFINED_SHAPE
				expecting = _COL_NAMES
			}
			// Otherwise loop to the next line, and we are still expecting table name.

		case _COL_NAMES: // Also proxy for first row of a table struct in the form: <name> <type> = <value>

			// EITHER (1) read a line of a table struct OR (2) read col names of a tabular table.

			if tableShape == _UNDEFINED_SHAPE {
				tableShape, structHasRowData, err = getTableShape(p, line, lineSplit, tableShape)
				if err != nil {
					return nil, err
				}
			}

			if tableShape == _STRUCT_SHAPE {

				// (1) Get the table struct (name, type and optional equals value) of this line.

				table.structShape = true
where(lineSplit)
				var colName string = lineSplit[structNameIndex]
				var colType string = lineSplit[structTypeIndex]

				var isValid bool
				if isValid, err = IsValidColName(colName); !isValid {
					return nil, fmt.Errorf("#2 %s looks like struct but %s", p.gotFilePos(), err)
				}

				var colNameSlice []string = []string{colName}
				if isValid, err = IsValidColType(colType); !isValid {
					return nil, fmt.Errorf("#3 %s looks like struct but %s", p.gotFilePos(), err)
				}
				var colTypeSlice []string = []string{colType}

				err = table.AppendCol(colName, colType)
				if err != nil {
					return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
				}

				/*
					// Set this only once (for each table). Base on the first "col", which is <name> <type> = <value>
					if table.ColCount() == 1 { // The first struct item.
						structHasRowData = isNameAndTypeEqualsValueStruct
					}
				*/

				if structHasRowData {
					// Find the equals sign byte location within the string so we can locate the value data after equals.
					// We know it's there (from the line split above), so don't check for nil returned.
					var rangeFound []int = equalsRegexp.FindStringIndex(line)
					if rangeFound == nil { // This avoids a runtime error.
						return nil, fmt.Errorf("%s expecting <name> <type> = <value> but found: %s", p.gotFilePos(), line)
					}
					var valueData string = line[rangeFound[1]:]        // Just after = equals sign.
					valueData = strings.TrimLeft(valueData, " \t\r\n") // Remove leading space.

					// Handle the first iteration (parse a line) through a struct, where the table has no rows.
					// Exactly one row is needed for a struct table which has data. Zero rows if no data.
					if table.RowCount() == 0 {
						err = table.AppendRow()
						if err != nil {
							return nil, err
						}

						// By default, if a parsed table is struct-shape, it will be set internally to struct-shape.
						// Of course, if any further (than its single) rows are appended, it will display as table-shape,
						// and will revert to struct-shape if it is reduced back to a single row.
						err = table.SetStructShape(true)
						if err != nil {
							return nil, err
						}
					}

					if debugging {
						// where(fmt.Sprintf("table.RowCount() = %d\n", table.RowCount()))
						// where(fmt.Sprintf("len(table.rows) = %d\n", len(table.rows)))
					}

					rowSliceOfStructTable, err = p.getRowSlice(valueData, colNameSlice, colTypeSlice)
					if err != nil {
						return nil, err
					}

					// Using table.SetValByColIndex() is less efficient but the volume of structs is small.
					var val interface{} = rowSliceOfStructTable[0]
					var colIndex int = len(table.rows[0]) - 1
					const rowIndexAlwaysZero int = 0
					err = table.SetValByColIndex(colIndex, rowIndexAlwaysZero, val)
					if err != nil {
						return nil, fmt.Errorf("%s %s", p.gotFilePos(), err)
					}

					// Still expecting _COL_NAMES which is where we find struct: <name> <type> = <value>
					// rowMapOfStruct is a variable of type tableRow which is a map: map[string]interface{}
					// Look up the value by reference to the colName.
				}
			} else {

				// (2) Get the col names.

				parserColNames, err = p.getColNames(lineSplit)
				if err != nil {
					return nil, err
				}

				expecting = _COL_TYPES
			}

		case _COL_TYPES:

			parserColTypes, err = p.getColTypes(line)
			if err != nil {
				return nil, fmt.Errorf("table [%s] %s", table.Name(), err)
			}
			lenColNames := len(parserColNames)
			lenColTypes := len(parserColTypes)
			if lenColTypes != lenColNames {
				return nil,
					fmt.Errorf("%s expecting: %d col type%s but found: %d",
						p.gotFilePos(), lenColNames, plural(lenColNames), lenColTypes)
			}

			// Append cols here now that both parserColNames and parserColTypes are available.
			// Trust that gotables syntax error handling will ensure both are available here.
			err = table.appendCols(parserColNames, parserColTypes)
			if err != nil {
				return nil, err
			}

			expecting = _COL_ROWS

		case _COL_ROWS:

			// Found data.

			lenColTypes := len(parserColTypes)

			var rowSlice tableRow
			rowSlice, err = p.getRowSlice(line, parserColNames, parserColTypes)
			if err != nil {
				return nil, err
			}

			err = table.appendRowSlice(rowSlice)
			if err != nil {
				return tables, err
			}

			lenRowSlice := len(rowSlice)
			if lenColTypes != lenRowSlice {
				return nil, fmt.Errorf("%s expecting: %d value%s but found: %d",
					p.gotFilePos(), lenColTypes, plural(lenColTypes), lenRowSlice)
			}

			if lenColTypes != lenRowSlice {
				return nil, fmt.Errorf("%s expecting: %d value%s but found: %d",
					p.gotFilePos(), lenColTypes, plural(lenColTypes), lenRowSlice)
			}

		default:
			return nil, fmt.Errorf("%s expecting table name, col names or type names but found: %s",
				p.gotFilePos(), expecting)
		}

		if readError == io.EOF {
			return tables, nil // It's not an error to reach EOF. It just means end of document.
		}
	}
}

func getTableShape(p *parser, line string, lineSplit []string, tableShapeIn _TableShape) (tableShapeOut _TableShape, hasStructRowData bool, err error) {

	/*
		This is a recognition step.
		Determine whether this is a candidate struct of either:-
		(a) <name> <type>
		(b) <name> <type> =		// INVALID
		(c) <name> <type> = <value>
		(d) <name> <type> = <value>	// Note strings with whitespace can mean > 4
		OR
		A row of col names or a row of col types:
		n1 n2 n3
		t1 t2 t3
	*/

	if tableShapeIn != _UNDEFINED_SHAPE {
		err = fmt.Errorf("tableShape has already been set")
		return
	}

	// Note: strings can mean len(lineSplit) > 4 but still valid. So can't just test for exactly 4.
	var lenLineSplit int = len(lineSplit)

	switch lenLineSplit {
	case 1:
		// <name>
		// <type> possibly
		tableShapeOut = _TABLE_SHAPE
		hasStructRowData = false
		return
	case 2:
		// (a) <name> <type>
		firstTokenIsName, _ := IsValidColName(lineSplit[structNameIndex])
		secondTokenIsType, _ := IsValidColType(lineSplit[structTypeIndex])
		if firstTokenIsName && secondTokenIsType {
			tableShapeOut = _STRUCT_SHAPE
			return
		} else {
			tableShapeOut = _TABLE_SHAPE
			return
		}
	case 3:
		// (b) <name> <type> =      // INVALID
		if lineSplit[structEqualsIndex] == "=" {
			tableShapeOut = _STRUCT_SHAPE
			err = fmt.Errorf("#1 %s looks like struct but missing value after equals: %s %s %s",
				p.gotFilePos(), lineSplit[0], lineSplit[1], lineSplit[2])
			return
		} else {
			tableShapeOut = _TABLE_SHAPE
			return
		}
	case 4:
		// (c) <name> <type> = <value>
		firstTokenIsName, _ := IsValidColName(lineSplit[structNameIndex])
		secondTokenIsType, _ := IsValidColType(lineSplit[structTypeIndex])
		thirdTokenIsEquals := lineSplit[structEqualsIndex] == "="
		if firstTokenIsName && secondTokenIsType && thirdTokenIsEquals {
			tableShapeOut = _STRUCT_SHAPE
			hasStructRowData = true
			return
		} else {
			tableShapeOut = _TABLE_SHAPE
			return
		}
	default: // > 4
		// (d) <name> <type> = <value>	// Note strings with whitespace can mean > 4
		firstTokenIsName, _ := IsValidColName(lineSplit[structNameIndex])
		secondTokenIsType, _ := IsValidColType(lineSplit[structTypeIndex])
		thirdTokenIsEquals := lineSplit[structEqualsIndex] == "="
		if firstTokenIsName && secondTokenIsType && thirdTokenIsEquals {
			tableShapeOut = _STRUCT_SHAPE
			hasStructRowData = true
			return
		} else {
			tableShapeOut = _TABLE_SHAPE
			return
		}
	}
}

func (p *parser) parseFile(inputFileName string) (*TableSet, error) {
	var fileName string = inputFileName
	var err error
	var fileBytes []byte

	p.SetFileName(fileName) // For file and line diagnostics.

	if filepath.IsAbs(fileName) {
		fileName, err = UtilFilepathAbs(fileName)
		if err != nil {
			return nil, err
		}
		fileName = filepath.ToSlash(fileName)
	}
	// Check that this is not a directory. Would like to check more.
	var info os.FileInfo
	info, err = os.Stat(fileName)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		err = fmt.Errorf("FILE ERROR: %q is a directory", fileName)
		return nil, err
	}

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
	return fmt.Sprintf("%s:%d:", p.fileName, globalLineNum)
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

func (p *parser) getTableSetName(line string) (string, error) {

	fields := strings.Fields(line)
	if len(fields) != 1 { // Note: len(fields) cannot be 0, because len(line) > 0 has been tested before call.
		return "", fmt.Errorf("%s expecting a table set name in double square brackets but found: %s", p.gotFilePos(), fields[0])
	}

	tableSetName := fields[0]
	result := tableSetNameRegexp.MatchString(tableSetName)
	if !result {
		return "", fmt.Errorf("%s expecting a valid alpha-numeric table set name in double square brackets, eg [[_Foo2Bar3]] but found: %s",
			p.gotFilePos(), tableSetName)
	}

	// Must be at least 5 chars to have matched tableSetNameRegexp.

	// Strip off surrounding [[]]
	tableSetName = tableSetName[2 : len(tableSetName)-2]

	return tableSetName, nil
}

func (p *parser) getTableName(line string) (string, error) {

	fields := strings.Fields(line)
	if len(fields) != 1 { // Note: len(fields) cannot be 0, because len(line) > 0 has been tested before call.
		return "", fmt.Errorf("%s expecting a table name in square brackets but found: %s", p.gotFilePos(), fields[0])
	}

	tableName := fields[0]
	result := tableNameRegexp.MatchString(tableName)
	if !result {
		return "", fmt.Errorf("%s expecting a valid alpha-numeric table name in square brackets, eg [_Foo2Bar3] but found: %s",
			p.gotFilePos(), tableName)
	}

	// Must be at least 3 chars to have matched tableNameRegexp.

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
					return nil, fmt.Errorf("%s %s did you perhaps mean either: %s %s OR %s %s = <val>",
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

	var colTypes []string = whiteRegexp.Split(line, _ALL_SUBSTRINGS)
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

	gotables.Table supports:-
	* Go types - except complex64 complex128
	* []byte and []uint8
*/
func IsValidColType(colType string) (bool, error) {

	_, contains := globalColTypesMap[colType]
	if !contains {
		err := fmt.Errorf("invalid col type: %s (allowed: most Go built-in types, except complex)", colType)
		return false, err
	}

	return true, nil
}

/*
Returns true for those Go types that are numeric.

Go types NOT supported: complex64 complex128
*/
func IsNumericColType(colType string) (bool, error) {
	_, contains := globalNumericColTypesMap[colType]

	if !contains {
		msg := fmt.Sprintf("non-numeric col type: %s (numeric types:", colType)
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

func IsSliceColType(colType string) (bool, error) {
	_, contains := globalSliceColTypesMap[colType]

	if !contains {
		msg := fmt.Sprintf("non-slice col type: %s (slice types:", colType)
		// Note: Because maps are not ordered, this (desirably) shuffles the order of valid col types with each call.
		for typeName, _ := range globalSliceColTypesMap {
			msg += fmt.Sprintf(" %s", typeName)
		}
		msg += ")"
		err := errors.New(msg)
		return false, err
	}

	return true, nil
}

func IsTableColType(colType string) (bool, error) {
	_, contains := globalTableColTypesMap[colType]

	if !contains {
		msg := fmt.Sprintf("non-table col type: %s (table types:", colType)
		// Note: Because maps are not ordered, this (desirably) shuffles the order of valid col types with each call.
		for typeName, _ := range globalTableColTypesMap {
			msg += fmt.Sprintf(" %s", typeName)
		}
		msg += ")"
		err := errors.New(msg)
		return false, err
	}

	return true, nil
}

/*
	Following Rob Pike and avoiding a regular expression where a simple loop will do.
	https://commandcenter.blogspot.com/2011/08/regular-expressions-in-lexing-and.html
	This function led to a substantial benchmark improvement.
*/
func isValidName(name string) (bool, error) {
	if len(name) == 0 {
		return false, fmt.Errorf("%s: len(name) == 0", UtilFuncName())
	}
	if !unicode.IsLetter(rune(name[0])) && name[0] != '_' {
		return false, fmt.Errorf("%s: invalid first char: %c (expecting letter or '_')", UtilFuncName(), name[0])
	}
	for i := 1; i < len(name); i++ {
		if !unicode.IsLetter(rune(name[i])) && name[i] != '_' && !unicode.IsNumber(rune(name[i])) {
			return false, fmt.Errorf("%s: invalid char: %c (expecting letter or '_' or number)", UtilFuncName(), name[0])
		}
	}

	return true, nil
}

// Note: The same validity rules apply to both table names and col names.
func IsValidColName(colName string) (bool, error) {

	// Following Rob Pike and avoiding a regular expression where a simple loop will do.
	isValid, _ := isValidName(colName)
	if !isValid {
		return false, fmt.Errorf("#1 invalid col name: %q (valid example: \"_Foo2Bar3\")", colName)
	}

	_, contains := globalColTypesMap[colName]
	if contains {
		return false, fmt.Errorf("#2 invalid col name: %s (cannot use Go type as col name)", colName)
	}

	return true, nil
}

// Note: The same validity rules apply to both table names and col names.
// This tests table name WITHOUT surrounding square brackets. Text part only.
func IsValidTableName(tableName string) (bool, error) {
	if len(tableName) < 1 {
		return false, errors.New("invalid table name has zero length")
	}

	// Following Rob Pike and avoiding a regular expression where a simple loop will do.
	isValid, _ := isValidName(tableName)
	if !isValid {
		return false, fmt.Errorf("invalid table name: %q (valid example: \"_Foo1Bar2\")", tableName)
	}

	_, contains := globalColTypesMap[tableName]
	if contains {
		return false, fmt.Errorf("invalid table name: %s (cannot use a Go type or '*Table' as a table name)", tableName)
	}

	return true, nil
}

func (p *parser) getRowSlice(line string, colNames []string, colTypes []string) (tableRow, error) {
	var err error
	rowSlice := make(tableRow, len(colNames))

	remaining := line // Remainder of line left to parse.
	var rangeFound []int
	var textFound string
	var colCount = 0
	var lenColTypes = len(colTypes)
	var i int

	var boolVal bool
	var uint8Val uint8
	var uint8SliceVal []uint8
	var byteSliceVal []byte
	var uint16Val uint16
	var uint32Val uint32
	var uint64Val uint64
	var uintVal uint
	var intVal int
	var int8Val int8
	var int16Val int16
	var int32Val int32
	var runeVal rune
	var int64Val int64
	var float32Val float32
	var float64Val float64
	var tableVal *Table
	var timeVal time.Time

	for i = 0; i < lenColTypes; i++ {
		if len(remaining) == 0 { // End of line
			return nil, fmt.Errorf("%s expecting %d value%s but found only %d", p.gotFilePos(), lenColTypes, plural(lenColTypes), colCount)
		}
		switch colTypes[i] {
		case "string":
			rangeFound = stringRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s expecting a valid value of double-quoted %s but found: %s (Need backticks? Use []byte)",
					p.gotFilePos(), colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			unquoted, err := strconv.Unquote(textFound) // Note: strconv.Unquote() strips off surrounding double-quotes.
			if err != nil {
				return nil, fmt.Errorf("%s error: %v of string: %s",
					p.gotFilePos(), err, textFound)
			}
			rowSlice[i] = unquoted
		case "bool":
			rangeFound = boolRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			boolVal, err = strconv.ParseBool(textFound)
			if err != nil { // This error check probably redundant.
				return nil, fmt.Errorf("%s %s for type %s", p.gotFilePos(), err, colTypes[i])
			}
			rowSlice[i] = boolVal
		case "uint8", "byte":
			rangeFound = uintRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			//			uint64Val, err = strconv.ParseUint(textFound, _DEC, _BITS_8)
			if go_1_13_number_literals {
				uint64Val, err = parseUint(textFound, _BITS_8)
			} else {
				uint64Val, err = strconv.ParseUint(textFound, _DEC, _BITS_8)
			}
			if err != nil {
				rangeMsg := rangeForIntegerType(0, math.MaxUint8)
				return nil, fmt.Errorf("#1 %s: %s %s for type %s %s", UtilFuncName(), p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			uint8Val = uint8(uint64Val)
			rowSlice[i] = uint8Val
		case "[]uint8":
			// Go stores byte as uint8, so there's no need to process byte differently. ???
			rangeFound = uintSliceRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			var sliceString string = textFound[1 : len(textFound)-1] // Strip off leading and trailing [] slice delimiters.
			var sliceStringSplit []string = splitSliceString(sliceString)
			uint8SliceVal = make([]uint8, len(sliceStringSplit))
			for el := 0; el < len(sliceStringSplit); el++ {
				if go_1_13_number_literals {
					uint64Val, err = parseUint(sliceStringSplit[el], _BITS_8)
				} else {
					uint64Val, err = strconv.ParseUint(sliceStringSplit[el], _DEC, _BITS_8)
				}
				if err != nil {
					rangeMsg := rangeForIntegerType(0, math.MaxUint8)
					return nil, fmt.Errorf("#2 %s: %s %s for type %s %s", UtilFuncName(), p.gotFilePos(), err, colTypes[i], rangeMsg)
				}
				uint8SliceVal[el] = uint8(uint64Val)
			}
			rowSlice[i] = uint8SliceVal
		case "[]byte":
			// Go stores byte as uint8, so there's no need to process byte differently. ???
			rangeFound = uintSliceRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			var sliceString string = textFound[1 : len(textFound)-1] // Strip off leading and trailing [] slice delimiters.
			var sliceStringSplit []string = splitSliceString(sliceString)
			byteSliceVal = make([]uint8, len(sliceStringSplit))
			for el := 0; el < len(sliceStringSplit); el++ {
				//				uint64Val, err = strconv.ParseUint(sliceStringSplit[el], _DEC, _BITS_8)
				if go_1_13_number_literals {
					uint64Val, err = parseUint(sliceStringSplit[el], _BITS_8)
				} else {
					uint64Val, err = strconv.ParseUint(sliceStringSplit[el], _DEC, _BITS_8)
				}
				if err != nil {
					rangeMsg := rangeForIntegerType(0, math.MaxUint8)
					return nil, fmt.Errorf("#3 %s: %s %s for type %s %s", UtilFuncName(), p.gotFilePos(), err, colTypes[i], rangeMsg)
				}
				byteSliceVal[el] = byte(uint64Val)
			}
			rowSlice[i] = byteSliceVal
		case "uint16":
			rangeFound = uintRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			//			uint64Val, err = strconv.ParseUint(textFound, _DEC, _BITS_16)
			if go_1_13_number_literals {
				uint64Val, err = parseUint(textFound, _BITS_16)
			} else {
				uint64Val, err = strconv.ParseUint(textFound, _DEC, _BITS_16)
			}
			if err != nil {
				rangeMsg := rangeForIntegerType(0, math.MaxUint16)
				return nil, fmt.Errorf("#3 %s: %s %s for type %s %s", UtilFuncName(), p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			uint16Val = uint16(uint64Val)
			rowSlice[i] = uint16Val
		case "uint32":
			rangeFound = uintRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			//			uint64Val, err = strconv.ParseUint(textFound, _DEC, _BITS_32)
			if go_1_13_number_literals {
				uint64Val, err = parseUint(textFound, _BITS_32)
			} else {
				uint64Val, err = strconv.ParseUint(textFound, _DEC, _BITS_32)
			}
			if err != nil {
				rangeMsg := rangeForIntegerType(0, math.MaxUint32)
				return nil, fmt.Errorf("#4 %s: %s %s for type %s %s", UtilFuncName(), p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			uint32Val = uint32(uint64Val)
			rowSlice[i] = uint32Val
		case "uint64":
			rangeFound = uintRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			//			uint64Val, err = strconv.ParseUint(textFound, _DEC, _BITS_64)
			if go_1_13_number_literals {
				uint64Val, err = parseUint(textFound, _BITS_64)
			} else {
				uint64Val, err = strconv.ParseUint(textFound, _DEC, _BITS_64)
			}
			if err != nil {
				rangeMsg := rangeForIntegerType(0, math.MaxUint64)
				return nil, fmt.Errorf("#5 %s: %s %s for type %s %s", UtilFuncName(), p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			rowSlice[i] = uint64Val
		case "uint":
			rangeFound = uintRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			// uint and int are the same size.
			//			uint64Val, err = strconv.ParseUint(textFound, _DEC, strconv.IntSize)
			if go_1_13_number_literals {
				uint64Val, err = parseUint(textFound, strconv.IntSize)
			} else {
				uint64Val, err = strconv.ParseUint(textFound, _DEC, strconv.IntSize)
			}
			if err != nil {
				var minVal int64
				var maxVal uint64
				switch strconv.IntSize {
				case 32:
					minVal = 0
					maxVal = math.MaxUint32
				case 64:
					minVal = 0
					maxVal = math.MaxUint64
				default:
					msg := fmt.Sprintf("#6 %s: unknown int size: %d bits", UtilFuncName(), strconv.IntSize)
					log.Printf("%s", msg)
					return nil, fmt.Errorf("%s", msg)
				}
				rangeMsg := rangeForIntegerType(minVal, maxVal)
				return nil, fmt.Errorf("#7 %s: %s %s for type %s %s", UtilFuncName(), p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			uintVal = uint(uint64Val) // May be unnecessary.
			rowSlice[i] = uintVal
		case "int":
			rangeFound = intRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			if go_1_13_number_literals {
				int64Val, err = parseInt(textFound, strconv.IntSize)
			} else {
				int64Val, err = strconv.ParseInt(textFound, _DEC, strconv.IntSize)
			}
			if err != nil {
				var minVal int64
				var maxVal uint64
				switch strconv.IntSize {
				case 32:
					minVal = math.MinInt32
					maxVal = math.MaxInt32
				case 64:
					minVal = math.MinInt64
					maxVal = math.MaxInt64
				default:
					msg := fmt.Sprintf("unknown int size: %d bits", strconv.IntSize)
					log.Printf("%s", msg)
					return nil, fmt.Errorf("%s", msg)
				}
				rangeMsg := rangeForIntegerType(minVal, maxVal)
				return nil, fmt.Errorf("%s %s for type %s %s", p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			intVal = int(int64Val) // May be unnecessary.
			rowSlice[i] = intVal
		case "int8":
			rangeFound = intRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			//			int64Val, err = strconv.ParseInt(textFound, _DEC, _BITS_8)
			//			int64Val, err = parseInt(textFound, _BITS_8)
			if go_1_13_number_literals {
				int64Val, err = parseInt(textFound, _BITS_8)
			} else {
				int64Val, err = strconv.ParseInt(textFound, _DEC, _BITS_8)
			}
			if err != nil {
				// Example: data.got[55] strconv.ParseInt: parsing "-129": value out of range for type int8
				rangeMsg := rangeForIntegerType(math.MinInt8, math.MaxInt8)
				return nil, fmt.Errorf("%s %s for type %s %s", p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			int8Val = int8(int64Val)
			rowSlice[i] = int8Val
		case "int16":
			rangeFound = intRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			//			int64Val, err = strconv.ParseInt(textFound, _DEC, _BITS_16)
			if go_1_13_number_literals {
				int64Val, err = parseInt(textFound, _BITS_16)
			} else {
				int64Val, err = strconv.ParseInt(textFound, _DEC, _BITS_16)
			}
			if err != nil {
				rangeMsg := rangeForIntegerType(math.MinInt16, math.MaxInt16)
				return nil, fmt.Errorf("%s %s for type %s %s", p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			int16Val = int16(int64Val)
			rowSlice[i] = int16Val
		case "int32":
			rangeFound = intRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			//			int64Val, err = strconv.ParseInt(textFound, _DEC, _BITS_32)
			if go_1_13_number_literals {
				int64Val, err = parseInt(textFound, _BITS_32)
			} else {
				int64Val, err = strconv.ParseInt(textFound, _DEC, _BITS_32)
			}
			if err != nil {
				rangeMsg := rangeForIntegerType(math.MinInt32, math.MaxInt32)
				return nil, fmt.Errorf("%s %s for type %s%s ", p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			int32Val = int32(int64Val)
			rowSlice[i] = int32Val
		case "rune":
			rangeFound = runeRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			if rangeFound[1]-rangeFound[0] < 3 { // Expecting 2 delimeters surrounding at least 1 char.
				return nil, fmt.Errorf("%s invalid rune with zero length: ''", p.gotFilePos())
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			var runeText string = textFound[1 : len(textFound)-1] // Strip off leading and trailing '' quotes.
			runeVal, err = parseRune(runeText)
			if err != nil {
				return nil, fmt.Errorf("%s %v", p.gotFilePos(), err)
			}
			rowSlice[i] = runeVal
		case "int64":
			rangeFound = intRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			// int64Val, err = strconv.ParseInt(textFound, _DEC, _BITS_64)
			if go_1_13_number_literals {
				int64Val, err = parseInt(textFound, _BITS_64)
			} else {
				int64Val, err = strconv.ParseInt(textFound, _DEC, _BITS_64)
			}
			if err != nil {
				rangeMsg := rangeForIntegerType(math.MinInt64, math.MaxInt64)
				return nil, fmt.Errorf("%s %s for type %s %s", p.gotFilePos(), err, colTypes[i], rangeMsg)
			}
			rowSlice[i] = int64Val
		case "float32":
			rangeFound = floatRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			float64Val, err = strconv.ParseFloat(textFound, _BITS_32)
			if err != nil {
				return nil, fmt.Errorf("%s %s for type %s", p.gotFilePos(), err, colTypes[i])
			}
			if math.IsNaN(float64Val) && textFound != "NaN" {
				return nil, fmt.Errorf("%s col %s: expecting NaN as Not-a-Number for type %s but found: %s ",
					p.gotFilePos(), colNames[i], colTypes[i], textFound)
			}
			float32Val = float32(float64Val)
			rowSlice[i] = float32Val
		case "float64":
			rangeFound = floatRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s", p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			float64Val, err = strconv.ParseFloat(textFound, _BITS_64)
			if err != nil {
				return nil, fmt.Errorf("%s %s for type %s", p.gotFilePos(), err, colTypes[i])
			}
			if math.IsNaN(float64Val) && textFound != "NaN" {
				return nil, fmt.Errorf("%s col %s: expecting NaN as Not-a-Number for type %s but found: %s",
					p.gotFilePos(), colNames[i], colTypes[i], textFound)
			}
			rowSlice[i] = float64Val
		case "*Table", "*gotables.Table":
			rangeFound = tableNameRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				// See if it's a nil table.
				rangeFound = tableNameNilRegexp.FindStringIndex(remaining)
				if rangeFound != nil {
					tableVal = NewNilTable()
				} else {
					return nil, fmt.Errorf("%s expecting a valid place-holder value of type %s, in square brackets, but found: %s",
						p.gotFilePos(), colTypes[i], remaining)
				}
			} else {
				textFound = remaining[rangeFound[0]:rangeFound[1]]
				// tableNameRegexp matches to [name] therefore Trim() will work safely.
				var tableName string = strings.Trim(textFound, "[]")
				tableVal, err = NewTable(tableName)
				if err != nil {
					return nil, fmt.Errorf("%s expecting a valid table name in the form [name] but found: %s",
						p.gotFilePos(), remaining)
				}
			}
			rowSlice[i] = tableVal
		case "time.Time":
			rangeFound = rfc3339TimeRegexp.FindStringIndex(remaining)
			if rangeFound == nil {
				return nil, fmt.Errorf("%s col %s expecting a valid value of type %s but found: %s",
					p.gotFilePos(), colNames[i], colTypes[i], remaining)
			}
			textFound = remaining[rangeFound[0]:rangeFound[1]]
			timeVal, err = time.Parse(time.RFC3339, textFound)
			if err != nil {
				// Example error message: expecting 1 value but found more:
				return nil, fmt.Errorf("%s %s for type %s", p.gotFilePos(), err, colTypes[i])
			}
			rowSlice[i] = timeVal
		default:
			log.Printf("Managed to reach unreachable code in getRowCol()") // Need to define another type?
			return nil, fmt.Errorf("line %s Unreachable code in getRowCol(): Need to define another type?", p.gotFilePos())
		}
		remaining = remaining[rangeFound[1]:]
		//		remaining = strings.TrimLeft(remaining, " \t\r\n") // Remove leading whitespace. Is \t\r\n overkill?
		remaining = strings.TrimLeft(remaining, " \t") // Remove leading whitespace.
		colCount++
	}

	if len(remaining) > 0 { // Still one or more columns to parse.
		// This handles both table shape and struct shape columns.
		where()
		return nil, fmt.Errorf("%s expecting %d value%s but found more: %s", p.gotFilePos(), lenColTypes, plural(lenColTypes), remaining)

	}

	return rowSlice, nil
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

func plural(items int) string {
	if items == 1 || items == -1 {
		// Singular
		return ""
	} else {
		// Plural
		return "s"
	}
}

/*
	whiteRegexp.Split returns a slice with 1 empty string element if the
	input sliceString is empty. But we want a slice with 0 elements.
*/
func splitSliceString(sliceString string) (sliceStringSplit []string) {
	if len(sliceString) == 0 {
		sliceStringSplit = []string{} // 0 elements, not 1 element.
	} else {
		sliceStringSplit = whiteRegexp.Split(sliceString, _ALL_SUBSTRINGS)
	}
	return
}

func hasDelims(s string, delim string) bool {
	var lens int = len(s)

	if lens < 1 {
		return false
	}

	/*
	   	if lens >= 2 {
	   		if s[0] == delim[0] && s[lens-1] == delim[0] {
	   where()
	   			return true
	   		}
	   	}
	*/

	if strings.HasSuffix(string(s[0]), delim) && strings.HasSuffix(string(s[len(s)-1]), delim) {
		return true
	}

	return false
}

func trimDelims(s string, delim string) string {
	if len(s) < 2 {
		return s
	}
	s = strings.TrimPrefix(s, delim)
	s = strings.TrimSuffix(s, delim)
	return s
}

// Convert rune or int32 to uint64 to call this function.
func isValidUnicode(code uint64) (bool, error) {
	// See: https://en.wikipedia.org/wiki/Unicode
	const maxUnicode = 0x10FFFF
	const minReservedUnicode = 0xD800
	const maxReservedUnicode = 0xDFFF
	fmt.Printf("maxUnicode    =    %d\n", maxUnicode)
	fmt.Printf("math.MaxInt32 = %d\n", math.MaxInt32)

	if code > maxUnicode {
		return false, fmt.Errorf("invalid unicode number > %X (%d) for rune: %X (%d)", code, code, code, code)
	}

	return true, nil
}

var globalPrevRuneVal rune

func parseRune(runeText string) (rune, error) {
	var err error
	const delim = '\'' // Single quote for single-quote unquote rules.
	var runeVal rune
	var tail string // remainder (if any) of string after first rune is unquoted.

	runeVal, _, tail, err = strconv.UnquoteChar(runeText, delim)
	if err != nil {

		/*
			// Work around an apparent bug in strconv.UnquoteChar() where it returns an error trying to parse "\\"
			if runeText == "\\" {
				return '\\', nil
			}
			// Turns out not to be a bug. String should be "'\\\\'"
		*/

		return 0, fmt.Errorf("invalid rune literal '%s' with %s: %s", runeText, err, runeText)
	}

	if len(tail) > 0 {
		return 0, fmt.Errorf("invalid rune literal '%s' with trailing char%s: %q", runeText, plural(len(tail)), tail)
	}

	return runeVal, nil
}

// Parse int strings with non-decimal literal prefixes: 0b, 0o, 0x
func parseInt(s string, bitSize int) (int64Val int64, err error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("gotables.parseInt: parsing %q: invalid syntax", s)
	}

	if s[0] == '0' && len(s) >= 2 {
		// May be BIN, OCT or HEX.
		switch s[0:2] {
		case "0b", "0B":
			int64Val, err = strconv.ParseInt(s[2:], _BIN, bitSize)
		case "0o", "0O":
			int64Val, err = strconv.ParseInt(s[2:], _OCT, bitSize)
		case "0x", "0X":
			int64Val, err = strconv.ParseInt(s[2:], _HEX, bitSize)
		default: // octal with leading 0 only: not leading 0o or 0O
			int64Val, err = strconv.ParseInt(s, _OCT, bitSize)
		}
	} else {
		int64Val, err = strconv.ParseInt(s, _DEC, bitSize)
	}

	return
}

// Parse uint strings with non-decimal literal prefixes: 0b, 0o, 0x
func parseUint(s string, bitSize int) (uint64Val uint64, err error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("gotables.parseUint: parsing %q: invalid syntax", s)
	}

	if s[0] == '0' && len(s) >= 2 {
		// May be BIN, OCT or HEX.
		switch s[0:2] {
		case "0b", "0B":
			uint64Val, err = strconv.ParseUint(s[2:], _BIN, bitSize)
		case "0o", "0O":
			uint64Val, err = strconv.ParseUint(s[2:], _OCT, bitSize)
		case "0x", "0X":
			uint64Val, err = strconv.ParseUint(s[2:], _HEX, bitSize)
		default: // octal with leading 0 only: not leading 0o or 0O
			uint64Val, err = strconv.ParseUint(s, _OCT, bitSize)
		}
	} else {
		uint64Val, err = strconv.ParseUint(s, _DEC, bitSize)
	}

	return
}

func writeHeapDump() (err error) {
	fmt.Printf("writing %s.heapdump", os.Args[0])
	f, err := os.Create(os.Args[0] + ".heapdump")
	if err != nil {
		panic(err)
	}

	// Note: didn't defer f.Close() for the reasons in the following post:
	//       https://www.joeshaw.org/dont-defer-close-on-writable-files
	err = f.Close()
	if err != nil {
		return err
	}

	debug.WriteHeapDump(f.Fd())

	return nil
}
