package gotables

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	_ "math"
	_ "os"
	"reflect"
	"regexp"
	_ "strconv"
	"strings"
	"time"
)

type circRefMap map[*Table]struct{}

var EmptyStruct struct{}

var replaceSpaces *regexp.Regexp = regexp.MustCompile(` `)

var ubjsonTypesMap map[string]byte

func init() {
	ubjsonTypesMap = map[string]byte{
		"null":    'Z',
		"no-op":   'N', // no operation, to be ignored by the receiving end
		"true":    'T',
		"false":   'F',
		"int8":    'i',
		"uint8":   'U',
		"int16":   'I',
		"int32":   'l',
		"int64":   'L',
		"float32": 'd',
		"float64": 'D',
		"ASCII":   'C', // ASCII character
		"string":  'S', // UTF-8 string
	}
}

func (table *Table) GetTableAsJSON() (jsonString string, err error) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	if table == nil {
		return "", fmt.Errorf("%s table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
	}

	tableSet, err := NewTableSet("")
	if err != nil {
		return "", err
	}

	err = tableSet.Append(table)
	if err != nil {
		return "", err
	}

	jsonString, err = tableSet.GetTableSetAsJSON()
	if err != nil {
		return "", err
	}

	return
}

func (table *Table) getTableAsJSON_private() (jsonString string, err error) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	if table == nil {
		return "", fmt.Errorf("%s table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
	}

	var refMap circRefMap = map[*Table]struct{}{}
	var buf bytes.Buffer

	buf.WriteByte(123) // Opening brace outermost

	//where("***CALLING** getTableAsJSON_recursive()")
	err = getTableAsJSON_recursive(table, &buf, refMap, table)
	if err != nil {
		return "", err
	}

	buf.WriteByte(125) // Closing brace outermost

	jsonString = buf.String()

	return
}

func getTableAsJSON_recursive(table *Table, buf *bytes.Buffer, refMap circRefMap, topTable *Table) (err error) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	if table == nil {
		return fmt.Errorf("%s table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
	}

	// Add this table to the circular reference map.
	refMap[table] = EmptyStruct

	buf.WriteString(fmt.Sprintf(`"tableName":%q,`, table.Name()))

	isStructShape, err := table.IsStructShape()
	if err != nil {
		return err
	}
	if isStructShape {
		buf.WriteString(`"isStructShape":true,`)
	}

	buf.WriteString(`"metadata":[`)
	for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
		buf.WriteByte(123) // Opening brace around heading element (name: type)
		buf.WriteByte('"')
		buf.WriteString(table.colNames[colIndex])
		buf.WriteString(`":"`)
		buf.WriteString(table.colTypes[colIndex])
		buf.WriteByte('"')
		buf.WriteByte(125) // Closing brace around heading element (name: type)
		if colIndex < len(table.colNames)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']')
	buf.WriteByte(',') // Between metadata and data.

	// Get data

	buf.WriteString(`"data":[`)
	for rowIndex := 0; rowIndex < len(table.rows); rowIndex++ {
		buf.WriteByte('[') // Begin array of column cells.
		for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
			buf.WriteByte(123) // Opening brace
			buf.WriteString(fmt.Sprintf("%q:", table.colNames[colIndex]))
			var val interface{}
			val, err = table.GetValByColIndex(colIndex, rowIndex)
			if err != nil {
				return err
			}

			switch val.(type) {

			case string:
				buf.WriteString(fmt.Sprintf("%q", val.(string)))

			case bool, int, uint, int8, int16, int64, uint8, uint16, uint32, uint64, float32, float64:
				var valStr string
				valStr, err = table.GetValAsStringByColIndex(colIndex, rowIndex)
				if err != nil {
					return err
				}
				buf.WriteString(valStr)

			// rune vs int32
			case int32:
				var valStr string
				valStr, err = table.GetValAsStringByColIndex(colIndex, rowIndex)
				if err != nil {
					return err
				}
				switch table.colTypes[colIndex] {
				case "int32":
					buf.WriteString(valStr)
				case "rune":
					buf.WriteString(fmt.Sprintf("%q", valStr))
				default:
					msg := invalidColTypeMsg(table.Name(), fmt.Sprintf("%s: %s", UtilFuncName(), table.colTypes[colIndex]))
					buf.WriteString(msg)
				}

			case []byte:
				var valStr string
				valStr, err := table.GetValAsStringByColIndex(colIndex, rowIndex)
				if err != nil {
					return err
				}
				// Insert comma delimiters between slice elements.
				//				valStr = strings.ReplaceAll(valStr, " ", ",")	// New in Go 1.11?
				valStr = replaceSpaces.ReplaceAllString(valStr, ",")
				buf.WriteString(valStr)

			case *Table:
				var nestedTable *Table
				nestedTable, err = table.GetTableByColIndex(colIndex, rowIndex)
				if err != nil {
					return err
				}

				_, exists := refMap[nestedTable]
				if exists {
					err = fmt.Errorf("%s: circular reference in table [%s]: a reference to table [%s] already exists",
						UtilFuncName(), topTable.Name(), nestedTable.Name())
					return
				}

				isNilTable, err := nestedTable.IsNilTable()
				if err != nil {
					return err
				}
				if isNilTable {
					buf.WriteString("null")
				} else {
					buf.WriteByte(123) // Begin nested table.
					err = getTableAsJSON_recursive(nestedTable, buf, refMap, topTable)
					if err != nil {
						return err
					}
					buf.WriteByte(125) // End nested table.
				}

			case time.Time:
				var timeVal time.Time = val.(time.Time)
				if timeVal.Nanosecond() > 0 {
					buf.WriteByte('"')
					buf.WriteString(timeVal.Format(time.RFC3339Nano))
					buf.WriteByte('"')
				} else {
					buf.WriteByte('"')
					buf.WriteString(timeVal.Format(time.RFC3339))
					buf.WriteByte('"')
				}

			default:
				buf.WriteString(`"TYPE UNKNOWN"`)
			}

			buf.WriteByte(125) // Closing brace
			if colIndex < len(table.colNames)-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte(']') // End array of column cells.
		if rowIndex < len(table.rows)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']') // End array of rows.

	return
}

func (table *Table) GetTableAsJSONIndent() (jsonStringIndented string, err error) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	jsonString, err := table.GetTableAsJSON()
	if err != nil {
		return "", err
	}

	jsonStringIndented, err = indentJSON(jsonString)
	if err != nil {
		return "", err
	}

	return
}

func checkJsonDecodeError(checkErr error) (err error) {
	if checkErr == io.EOF {
		return fmt.Errorf("%s %s: unexpected EOF", UtilFuncSource(), UtilFuncName())
	}

	if checkErr != nil {
		return fmt.Errorf("%s %s: %v", UtilFuncSource(), UtilFuncName(), err)
	}

	return nil
}

func newTableFromJSON_recursive(jsonMap map[string]interface{}) (table *Table, err error) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	var exists bool

	/*
		We may not know the order map values will be returned if we iterate of the map:
		(1) tableName
		(2) isStructShape (if there)
		(3) metadata
		(4) data (if any)
		So we retrieve each of the 3 (possibly 2) top-level map values individually.
	*/

	// (1) Retrieve and process table name.
	var tableName string
	tableName, exists = jsonMap["tableName"].(string)
	if !exists {
		return nil, fmt.Errorf("JSON is missing table name")
	}
	table, err = NewTable(tableName)
	if err != nil {
		return nil, err
	}

	// If this optional isStructShape element is present, use it.
	var isStructShape bool
	isStructShape, exists = jsonMap["isStructShape"].(bool)
	if exists {
		err = table.SetStructShape(isStructShape)
		if err != nil {
			return nil, err
		}
	}

	// (2) Retrieve and process metadata.
	var metadata []interface{}
	metadata, exists = jsonMap["metadata"].([]interface{})
	if !exists {
		return nil, fmt.Errorf("JSON is missing table metadata")
	}
	// Loop through the array of metadata.
	for _, colNameAndType := range metadata {
		var colName string
		var colType string
		var val interface{}
		for colName, val = range colNameAndType.(map[string]interface{}) {
			// There's only one map element here: colName and colType.
		}
		var ok bool
		colType, ok = val.(string)
		if !ok {
			return nil, fmt.Errorf("expecting col type value from JSON string value but got type %T: %v", val, val)
		}

		err = table.AppendCol(colName, colType)
		if err != nil {
			return nil, err
		}
	}

	// (3) Retrieve and process data (if any).
	var data []interface{}
	data, exists = jsonMap["data"].([]interface{})
	if !exists {
		// Zero rows in this table. That's okay.
		return table, nil
	}

	// Loop through the array of rows.
	for rowIndex, val := range data {
		// where(fmt.Sprintf("row [%d] %v", rowIndex, val))
		err = table.AppendRow()
		if err != nil {
			return nil, err
		}

		var row []interface{} = val.([]interface{})
		for colIndex, val := range row {
			// where(fmt.Sprintf("\t\tcol [%d] %v", colIndex, val))
			var cell interface{}
			for _, cell = range val.(map[string]interface{}) {
				// There's only one map element here: colName and colType.
				// where(fmt.Sprintf("\t\t\tcol=%d row=%d celltype=%T cell=%v", colIndex, rowIndex, cell, cell))

				var colType string = table.colTypes[colIndex]
				switch cell.(type) {
				case string:
					switch colType { // We need to convert time string format to time.Time
					case "rune":
						var stringVal = cell.(string)
						var runeSlice []rune = []rune(stringVal)
						var runeVal rune = runeSlice[0]
						err = table.SetRuneByColIndex(colIndex, rowIndex, runeVal)
					case "time.Time":
						var timeVal time.Time
						timeVal, err = time.Parse(time.RFC3339, cell.(string))
						if err != nil { // We need this extra error check here
							err := fmt.Errorf("could not convert JSON time string to gotables %s", colType)
							return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), UtilFuncName(), err)
						}
						err = table.SetTimeByColIndex(colIndex, rowIndex, timeVal)
						if err != nil {
							err := fmt.Errorf("could not convert JSON string to gotables %s", colType)
							return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), UtilFuncName(), err)
						}
					case "string":
						err = table.SetStringByColIndex(colIndex, rowIndex, cell.(string))
					default:
						return nil, fmt.Errorf("%s %s: unexpected value of type: %s",
							UtilFuncSource(), UtilFuncName(), colType)
					}
					// Single error handler for all the calls in this switch statement.
					if err != nil {
						err := fmt.Errorf("could not convert JSON string to gotables %s", colType)
						return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), UtilFuncName(), err)
					}

				case float64:
					// Should never happen. All numbers are (now) type json.Number
					return nil, fmt.Errorf("%s %s: unexpected value of type: float64",
						UtilFuncSource(), UtilFuncName())

				// Deal with conversions to larger ints: int64 uint64
				case json.Number: // We set to json.Number with: decoder.UseNumber()
					var float64Val float64
					float64Val, err = cell.(json.Number).Float64()
					if err != nil {
						err := fmt.Errorf("could not convert json.Number to float64")
						return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), UtilFuncName(), err)
					}

					switch colType { // We need to convert them back to gotables numeric types
					case "int64":
						err = table.SetInt64ByColIndex(colIndex, rowIndex, int64(float64Val))
					case "uint64":
						err = table.SetUint64ByColIndex(colIndex, rowIndex, uint64(float64Val))
					case "float32":
						err = table.SetFloat32ByColIndex(colIndex, rowIndex, float32(float64Val))
					case "float64":
						err = table.SetFloat64ByColIndex(colIndex, rowIndex, float64Val)
					case "int":
						err = table.SetIntByColIndex(colIndex, rowIndex, int(float64Val))
					case "uint":
						err = table.SetUintByColIndex(colIndex, rowIndex, uint(float64Val))
					case "byte":
						err = table.SetByteByColIndex(colIndex, rowIndex, byte(float64Val))
					case "int8":
						err = table.SetInt8ByColIndex(colIndex, rowIndex, int8(float64Val))
					case "int16":
						err = table.SetInt16ByColIndex(colIndex, rowIndex, int16(float64Val))
					case "int32":
						err = table.SetInt32ByColIndex(colIndex, rowIndex, int32(float64Val))
					case "uint8":
						err = table.SetUint8ByColIndex(colIndex, rowIndex, uint8(float64Val))
					case "uint16":
						err = table.SetUint16ByColIndex(colIndex, rowIndex, uint16(float64Val))
					case "uint32":
						err = table.SetUint32ByColIndex(colIndex, rowIndex, uint32(float64Val))
					default:
						return nil, fmt.Errorf("%s %s: unexpected value of type: %s",
							UtilFuncSource(), UtilFuncName(), colType)
					}
					// Single error handler for all the calls in this switch statement.
					if err != nil {
						err := fmt.Errorf("could not convert JSON float64 to gotables %s", colType)
						return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), UtilFuncName(), err)
					}
				case bool:
					err = table.SetBoolByColIndex(colIndex, rowIndex, cell.(bool))

				case []interface{}: // This cell is a slice (probably either byte or uint8)
					var interfaceSlice []interface{} = cell.([]interface{})
					var byteSlice []byte = []byte{} // Ready to append to.
					var float64Val float64
					for _, sliceVal := range interfaceSlice {
						float64Val, err = sliceVal.(json.Number).Float64()
						if err != nil {
							err := fmt.Errorf("could not convert json.Number to float64")
							return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), UtilFuncName(), err)
						}
						byteSlice = append(byteSlice, byte(float64Val))
					}
					err = table.SetByteSliceByColIndex(colIndex, rowIndex, byteSlice)
					if err != nil {
						return nil, err
					}

				case map[string]interface{}: // This cell is a table.
					switch colType {
					case "*Table":
						tableNested, err := newTableFromJSON_recursive(cell.(map[string]interface{}))
						if err != nil {
							return nil, err
						}
						err = table.SetTableByColIndex(colIndex, rowIndex, tableNested)
						if err != nil {
							return nil, err
						}
					default:
						return nil, fmt.Errorf("newTableFromJSON_recursive(): unexpected cell value at [%s].(%d,%d)",
							tableName, colIndex, rowIndex)
					}
				case nil: // This cell is a nil table.
					switch colType {
					case "*Table":
						var tableNested *Table = NewNilTable()
						err = table.SetTableByColIndex(colIndex, rowIndex, tableNested)
						if err != nil {
							return nil, err
						}
					default:
						return nil, fmt.Errorf("%s %s: unexpected nil value at [%s].(%d,%d)",
							UtilFuncSource(), UtilFuncName(), tableName, colIndex, rowIndex)
					}
				default:
					return nil, fmt.Errorf("%s %s: unexpected value of type: %v",
						UtilFuncSource(), UtilFuncName(), reflect.TypeOf(cell))
				}
				// Single error handler for all the calls in this switch statement.
				if err != nil {
					return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), UtilFuncName(), err)
				}
			}
		}
	}

	return
}

func (tableSet *TableSet) GetTableSetAsJSONIndent() (jsonTableSetStringIndented string, err error) {

	if tableSet == nil {
		return "", fmt.Errorf("%s table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
	}

	jsonTableSetString, err := tableSet.GetTableSetAsJSON()
	if err != nil {
		return "", err
	}

	jsonTableSetStringIndented, err = indentJSON(jsonTableSetString)
	if err != nil {
		return "", err
	}

	return
}

/*
	Marshal gotables TableSet to JSON
*/
func (tableSet *TableSet) GetTableSetAsJSON() (jsonTableSetString string, err error) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	if tableSet == nil {
		return "", fmt.Errorf("%s table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
	}

	var buf bytes.Buffer

	buf.WriteByte(123) // Opening brace outermost
	buf.WriteString(fmt.Sprintf(`"tableSetName":%q,`, tableSet.Name()))
	buf.WriteString(`"tables":[`) // Opening array of tables

	var tableCount int = tableSet.TableCount()
	for tableIndex := 0; tableIndex < tableCount; tableIndex++ {
		table, err := tableSet.GetTableByTableIndex(tableIndex)
		if err != nil {
			return "", err
		}

		var jsonTable string
		jsonTable, err = table.getTableAsJSON_private()
		if err != nil {
			return "", err
		}

		buf.WriteString(jsonTable)

		if tableIndex < tableCount-1 {
			buf.WriteByte(',') // Delimiter between tables
		}
	}

	buf.WriteByte(']') // Closing array of tables
	buf.WriteByte(125) // Closing brace outermost

	jsonTableSetString = buf.String()

	return
}

/*
	Unmarshal JSON documents to a *gotables.TableSet

	See also:
		GetTableSetAsJSON()
*/
func NewTableSetFromJSON(jsonTableSetString string) (tableSet *TableSet, err error) {

	if jsonTableSetString == "" {
		return nil, fmt.Errorf("%s: jsonTableSetString is empty", UtilFuncName())
	}

	var jsonMap map[string]interface{}
	/*
		err = json.Unmarshal([]byte(jsonTableSetString), &jsonMap)
		if err != nil {
			return nil, err
		}
	*/

	decoder := json.NewDecoder(strings.NewReader(jsonTableSetString))
	decoder.UseNumber() // This prevents json from messing with int64 and uint64 values.
	err = decoder.Decode(&jsonMap)
	if err != nil {
		return nil, err
	}

	// (1) Retrieve and process TableSet name.
	var tableSetName string
	var exists bool
	tableSetName, exists = jsonMap["tableSetName"].(string)
	if !exists {
		return nil, fmt.Errorf("%s: JSON is missing tableSet name", UtilFuncName())
	}

	tableSet, err = NewTableSet(tableSetName)
	if err != nil {
		return nil, err
	}

	// (2) Retrieve and process tables.
	var tablesMap []interface{}
	tablesMap, exists = jsonMap["tables"].([]interface{})
	if !exists {
		return nil, fmt.Errorf("%s: JSON is missing tables", UtilFuncName())
	}

	var tableMap map[string]interface{}
	var tableMapInterface interface{}

	// (3) Loop through the array of tables.
	for _, tableMapInterface = range tablesMap {

		tableMap = tableMapInterface.(map[string]interface{})

		var table *Table
		table, err = newTableFromJSON_recursive(tableMap)
		if err != nil {
			return nil, err
		}

		err = tableSet.Append(table)
		if err != nil {
			return nil, err
		}
	}

	return
}

func NewTableFromJSON(jsonString string) (table *Table, err error) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	// This is similar to NewTableFromString which first gets a TableSet.

	if jsonString == "" {
		return nil, fmt.Errorf("%s: jsonString is empty", UtilFuncName())
	}

	tableSet, err := NewTableSetFromJSON(jsonString)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", UtilFuncName(), err)
	}

	tableCount := tableSet.TableCount()
	if tableCount != 1 {
		return nil, fmt.Errorf("%s: expecting a JSON string containing 1 table but found %d table%s",
			UtilFuncName(), tableCount, plural(tableCount))
	}

	table, err = tableSet.GetTableByTableIndex(0)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", UtilFuncName(), err)
	}

	return table, nil
}

func NewTableFromJSONByTableName(jsonString string, tableName string) (table *Table, err error) {
	tableSet, err := NewTableSetFromJSON(jsonString)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", UtilFuncName(), err)
	}

	table, err = tableSet.GetTable(tableName)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", UtilFuncName(), err)
	}

	return
}

func indentJSON(jsonString string) (jsonStringIndented string, err error) {

	var buf bytes.Buffer

	err = json.Indent(&buf, []byte(jsonString), "", "\t")
	if err != nil {
		return "", err
	}

	jsonStringIndented = buf.String()

	return
}
