package gotables

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"regexp"
)

type circRefMap map[*Table]struct{}

var empty struct{}

var replaceSpaces *regexp.Regexp = regexp.MustCompile(` `)

func (table *Table) GetTableAsJSON() (jsonString string, err error) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	if table == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
	}

	tableSet, err := NewTableSet("")
	if err != nil {
		return "", err
	}

	tableSet.Append(table)
	if err != nil {
		return "", err
	}

	jsonString, err = tableSet.GetTableSetAsJSON()
	if err != nil {
		return "", err
	}

	return
}

func (table *Table) getTableAsJSON_private() (json string, err error) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	if table == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
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

	json = buf.String()

	return
}

func getTableAsJSON_recursive(table *Table, buf *bytes.Buffer, refMap circRefMap, topTable *Table) (err error) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
	}

	// Add this table to the circular reference map.
	refMap[table] = empty

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

	//	buf.WriteString(fmt.Sprintf(`"%s%s":[`, dataTableNamePrefix, table.Name()))	// Begin array of rows.
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

			case bool, int, uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64:
				var valStr string
				valStr, err = table.GetValAsStringByColIndex(colIndex, rowIndex)
				if err != nil {
					return err
				}
				buf.WriteString(valStr)

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
		return fmt.Errorf("%s ERROR %s: unexpected EOF", UtilFuncSource(), UtilFuncName())
	}

	if checkErr != nil {
		return fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
	}

	return nil
}

func newTableFromJSON_recursive(m map[string]interface{}) (table *Table, err error) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	var exists bool

	/*
		We don't know the order map values will be returned if we iterate of the map:
		(1) tableName
		(2) metadata
		(3) data (if any)
		So we retrieve each of the 3 (possibly 2) top-level map values individually.
	*/

	// (1) Retrieve and process table name.
	var tableName string
	tableName, exists = m["tableName"].(string)
	if !exists {
		return nil, fmt.Errorf("JSON is missing table name")
	}
	table, err = NewTable(tableName)
	if err != nil {
		return nil, err
	}

	// If this optional isStructShape element is present, use it.
	var isStructShape bool
	isStructShape, exists = m["isStructShape"].(bool)
	if exists {
		err = table.SetStructShape(isStructShape)
		if err != nil {
			return nil, err
		}
	}

	// (2) Retrieve and process metadata.
	var metadata []interface{}
	metadata, exists = m["metadata"].([]interface{})
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
		colType, ok := val.(string)
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
	data, exists = m["data"].([]interface{})
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
					err = table.SetStringByColIndex(colIndex, rowIndex, cell.(string))
				case float64: // All JSON number values are stored as float64
					switch colType { // We need to convert them back to gotables numeric types
					case "int":
						err = table.SetIntByColIndex(colIndex, rowIndex, int(cell.(float64)))
					case "uint":
						err = table.SetUintByColIndex(colIndex, rowIndex, uint(cell.(float64)))
					case "byte":
						err = table.SetByteByColIndex(colIndex, rowIndex, byte(cell.(float64)))
					case "int8":
						err = table.SetInt8ByColIndex(colIndex, rowIndex, int8(cell.(float64)))
					case "int16":
						err = table.SetInt16ByColIndex(colIndex, rowIndex, int16(cell.(float64)))
					case "int32":
						err = table.SetInt32ByColIndex(colIndex, rowIndex, int32(cell.(float64)))
					case "int64":
						err = table.SetInt64ByColIndex(colIndex, rowIndex, int64(cell.(float64)))
					case "uint8":
						err = table.SetUint8ByColIndex(colIndex, rowIndex, uint8(cell.(float64)))
					case "uint16":
						err = table.SetUint16ByColIndex(colIndex, rowIndex, uint16(cell.(float64)))
					case "uint32":
						err = table.SetUint32ByColIndex(colIndex, rowIndex, uint32(cell.(float64)))
					case "uint64":
						err = table.SetUint64ByColIndex(colIndex, rowIndex, uint64(cell.(float64)))
					case "float32":
						err = table.SetFloat32ByColIndex(colIndex, rowIndex, float32(cell.(float64)))
					case "float64":
						err = table.SetFloat64ByColIndex(colIndex, rowIndex, float64(cell.(float64)))
					}
					// Single error handler for all the calls in this switch statement.
					if err != nil {
						err := fmt.Errorf("could not convert JSON float64 to gotables %s", colType)
						return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
					}
				case bool:
					err = table.SetBoolByColIndex(colIndex, rowIndex, cell.(bool))
				case []interface{}: // This cell is a slice
					var interfaceSlice []interface{} = cell.([]interface{})
					var byteSlice []byte = []byte{}
					for _, sliceVal := range interfaceSlice {
						byteSlice = append(byteSlice, byte(sliceVal.(float64)))
					}
					err = table.SetByteSliceByColIndex(colIndex, rowIndex, byteSlice)
					if err != nil {
						return nil, err
					}
				case map[string]interface{}: // This cell is a table.
					switch colType {
					case "*Table", "*gotables.Table":
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
					case "*Table", "*gotables.Table":
						var tableNested *Table = NewNilTable()
						err = table.SetTableByColIndex(colIndex, rowIndex, tableNested)
						if err != nil {
							return nil, err
						}
					default:
						return nil, fmt.Errorf("newTableFromJSON_recursive(): unexpected nil value at [%s].(%d,%d)",
							tableName, colIndex, rowIndex)
					}
				default:
					return nil, fmt.Errorf("%s ERROR %s: unexpected value of type: %v",
						UtilFuncSource(), UtilFuncName(), reflect.TypeOf(val))
				}
				// Single error handler for all the calls in this switch statement.
				if err != nil {
					return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
				}
			}
		}
	}

	return
}

func (tableSet *TableSet) GetTableSetAsJSONIndent() (jsonTableSetIndented string, err error) {

	if tableSet == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
	}

	jsonTableSet, err := tableSet.GetTableSetAsJSON()
	if err != nil {
		return "", err
	}

	jsonTableSetIndented, err = indentJSON(jsonTableSet)
	if err != nil {
		return "", err
	}

	return
}

/*
	Marshal gotables TableSet to JSON
*/
func (tableSet *TableSet) GetTableSetAsJSON() (jsonTableSet string, err error) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	if tableSet == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
	}

	var buf bytes.Buffer

	buf.WriteByte(123) // Opening brace outermost
	buf.WriteString(fmt.Sprintf(`"tableSetName":%q,`, tableSet.Name()))
	buf.WriteString(`"tables":[`) // Opening array of tables

	var tableCount int = tableSet.TableCount()
	for tableIndex := 0; tableIndex < tableCount; tableIndex++ {
		table, err := tableSet.TableByTableIndex(tableIndex)
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

	jsonTableSet = buf.String()

	return
}

/*
	Unmarshal JSON documents to a *gotables.TableSet
*/
func NewTableSetFromJSON(jsonTableSet string) (tableSet *TableSet, err error) {

	if jsonTableSet == "" {
		return nil, fmt.Errorf("%s: jsonTableSet is empty", UtilFuncName())
	}

	var m map[string]interface{}
	err = json.Unmarshal([]byte(jsonTableSet), &m)
	if err != nil {
		return nil, err
	}

	// (1) Retrieve and process TableSet name.
	var tableSetName string
	var exists bool
	tableSetName, exists = m["tableSetName"].(string)
	if !exists {
		return nil, fmt.Errorf("JSON is missing tableSet name")
	}

	tableSet, err = NewTableSet(tableSetName)
	if err != nil {
		return nil, err
	}

	// (2) Retrieve and process tables.
	var tablesMap []interface{}
	tablesMap, exists = m["tables"].([]interface{})
	if !exists {
		return nil, fmt.Errorf("JSON is missing tables")
	}

	var tableMap map[string]interface{}
	var tableMapInterface interface{}

	// Loop through the array of tables.
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

	table, err = tableSet.TableByTableIndex(0)
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

	table, err = tableSet.Table(tableName)
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
