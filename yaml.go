package gotables

import (
	_ "bytes"
	"fmt"
	"reflect"
	"time"

	yaml "gopkg.in/yaml.v3"
)

//func (tableSet *TableSet) GetTableSetAsYAML() (yamlString string, err error) {
//
//	if tableSet == nil {
//		return "", fmt.Errorf("%s tableSet.%s: table set is <nil>", UtilFuncSource(), UtilFuncName())
//	}
//
//	const twoSpaces string = "t "
//	var eightSpaces string
//	eightSpaces = "T< 4 6 >"	// XXXX
//	eightSpaces = "        "
//	var tableIndent string
//
//	var buf bytes.Buffer
//
//	buf.WriteString("---\n")	// Start of YAML document
//
//	var visitTableSet = func(tableSet *TableSet) (err error) {
//		buf.WriteString(`tableSetName: "`)
//		buf.WriteString(tableSet.Name())
//		buf.WriteByte('"')
//		buf.WriteByte('\n')
//		buf.WriteString("tables:\n")
//		return
//	}
//
//	var visitTable = func(table *Table) (err error) {
//
//		tableIndent = strings.Repeat(eightSpaces, table.depth*1)
//		var metadataIndent string
//		metadataIndent = "m "	// XXXX
//		metadataIndent = "  "
//		metadataIndent = "  "
//
//		buf.WriteString(tableIndent + "- tableName: ")
//		buf.WriteString(table.Name())
//		buf.WriteByte('\n')
//
//		if table.isStructShape {
//			buf.WriteString(tableIndent + metadataIndent + "isStructShape: true\n")
//		}
//
//		buf.WriteString(tableIndent + metadataIndent + "metadata:\n")
//
//		for i := 0; i < table.ColCount(); i++ {
//			buf.WriteString(tableIndent)
//			if table.colTypes[i] == "*Table" {
//				// Quote "*Table" to avoid YAML interpreting it as an alias.
//				buf.WriteString(fmt.Sprintf("%s- %s: %q\n", metadataIndent, table.colNames[i], table.colTypes[i]))
//			} else {
//				buf.WriteString(fmt.Sprintf("%s- %s: %s\n", metadataIndent, table.colNames[i], table.colTypes[i]))
//			}
//		}
//
//		buf.WriteString(tableIndent + metadataIndent + "data:\n")
//
//		return
//	}
//
//	var visitRow = func(row Row) (err error) {
//		return
//	}
//	_ = visitRow
//
//	var visitCell = func(cell Cell) (err error) {
//
//		var dataIndent string
//		if cell.ColIndex == 0 {
//			dataIndent = tableIndent + "D2-4- "	// rowIndent. Move it to visitRow() XXXX
//			dataIndent = tableIndent + "  - - "	// rowIndent. Move it to visitRow()
//		} else {
//			dataIndent = tableIndent + "d2 4- " // XXXX
//			dataIndent = tableIndent + "    - "
//		}
//		buf.WriteString(dataIndent)
//
//		var valString string
//		switch cell.ColType {
//		case "string":
//			valString, err = cell.Table.GetStringByColIndex(cell.ColIndex, cell.RowIndex)
//			if err != nil {
//				return err
//			}
//			buf.WriteString(fmt.Sprintf("%s: %q\n", cell.ColName, valString))
//		case "bool", "int", "uint", "int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64", "float32", "float64":
//			valString, err = cell.Table.GetValAsStringByColIndex(cell.ColIndex, cell.RowIndex)
//			if err != nil {
//				return err
//			}
//			buf.WriteString(fmt.Sprintf("%s: %s\n", cell.ColName, valString))
//		case "*Table":
//			buf.WriteString(fmt.Sprintf("%s:\n", cell.ColName))
//		default:
//			err = fmt.Errorf("%s: ERROR IN visitCell(): unknown type: %s\n", UtilFuncSource(), cell.ColType)
//		}
//
//		return
//	}
//
//	err = tableSet.Walk(visitTableSet, visitTable, nil, visitCell)
//	if err != nil {
//		return "", nil
//	}
//
//	yamlString = buf.String()
//
//	return
//}

/*
	Unmarshal YAML document to a *gotables.TableSet
*/
func NewTableSetFromYAML(yamlTableSetString string) (tableSet *TableSet, err error) {

	if yamlTableSetString == "" {
		return nil, fmt.Errorf("%s: yamlTableSetString is empty", UtilFuncName())
	}

	var m map[string]interface{}
	err = yaml.Unmarshal([]byte(yamlTableSetString), &m)
	if err != nil {
		return
	}
println()
where("\n" + printMap(m))

	// (1) Retrieve and process TableSet name.
	var tableSetName string
	var exists bool
	tableSetName, exists = m["tableSetName"].(string)
	if !exists {
		return nil, fmt.Errorf("%s: YAML is missing tableSet name", UtilFuncName())
	}

	tableSet, err = NewTableSet(tableSetName)
	if err != nil {
		return
	}

	// (2) Retrieve and process tables.
	var tablesMap []interface{}
	tablesMap, exists = m["tables"].([]interface{})
	if !exists {
		return nil, fmt.Errorf("%s: YAML is missing tables", UtilFuncName())
	}

	var tableMap map[string]interface{}
	var tableMapInterface interface{}

	// (3) Loop through the array of tables.
	for _, tableMapInterface = range tablesMap {

		tableMap = tableMapInterface.(map[string]interface{})
// where(printMap(tableMap))

		var table *Table
		table, err = newTableFromYAML_recursive(tableMap)
		if err != nil {
			return
		}

		err = tableSet.Append(table)
		if err != nil {
			return
		}
	}

	return
}

func newTableFromYAML_recursive(m map[string]interface{}) (table *Table, err error) {

	const TopFuncName string = "NewTableFromYAML()"

	var exists bool

	/*
		We don't know the order map values will be returned if we iterate of the map:
		(1) tableName
		(2) isStructShape (if there)
		(3) metadata
		(4) data (if any)
		So we retrieve each of the 3 (possibly 2) top-level map values individually.
	*/

	// (1) Retrieve and process table name.
	var tableName string
	tableName, exists = m["tableName"].(string)
	if !exists {
		return nil, fmt.Errorf("YAML is missing table name")
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
		return nil, fmt.Errorf("YAML is missing table metadata")
	}
	// Loop through the array of metadata.
	for _, colNameAndType := range metadata {
		var colName string
		var colType string
		var typeVal interface{}
		for colName, typeVal = range colNameAndType.(map[string]interface{}) {
			// There's only one map element here: colName and colType.
		}
		colType, ok := typeVal.(string)
		if !ok {
			return nil, fmt.Errorf("expecting col type value from YAML string value but got type %T: %v", typeVal, typeVal)
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
//where(printMap(data))

	// Loop through the array of rows.
	for rowIndex, rowVal := range data {
//where(printMap(rowVal))
		// where(fmt.Sprintf("rowIndex=%d: %v", rowIndex, rowVal))
		err = table.AppendRow()
		if err != nil {
			return nil, err
		}

		var row []interface{} = rowVal.([]interface{})
where(fmt.Sprintf("rowVal type: %T", rowVal))
where(fmt.Sprintf("row    type: %T", row))
//where(printMap(row))
		for colIndex, val := range row {
			// where(fmt.Sprintf("\t\tcol [%d] %v", colIndex, val))
			var cell interface{}
			var colName string
			for colName, cell = range val.(map[string]interface{}) {
where(fmt.Sprintf("colName  = %q", colName))
where(fmt.Sprintf("colIndex = %d", colIndex))
where(fmt.Sprintf("rowIndex = %d", rowIndex))
				// There's only one map element here: colName and colType.
				// where(fmt.Sprintf("\t\t\tcol=%d row=%d celltype=%T cell=%v", colIndex, rowIndex, cell, cell))

				var colType string = table.colTypes[colIndex]
where(fmt.Sprintf("colType  = %q", colType))
where(fmt.Sprintf("cell type: %T", cell))
println()
//				switch cell.(type) {
				switch colType {
				case "string":
					switch colType { // We need to convert time string format to time.Time
					case "time.Time":
						var timeVal time.Time
						timeVal, err = time.Parse(time.RFC3339, cell.(string))
						if err != nil { // We need this extra error check here
							err := fmt.Errorf("could not convert YAML time string to gotables %s", colType)
							return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), TopFuncName, err)
						}
						err = table.SetTimeByColIndex(colIndex, rowIndex, timeVal)
						if err != nil {
							err := fmt.Errorf("could not convert YAML string to gotables %s", colType)
							return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), TopFuncName, err)
						}
					default: // Is a string
						err = table.SetStringByColIndex(colIndex, rowIndex, cell.(string))
					}
					// Single error handler for all the calls in this switch statement.
					if err != nil {
						err := fmt.Errorf("could not convert YAML string to gotables %s", colType)
						return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), TopFuncName, err)
					}
				case "int":
					err = table.SetIntByColIndex(colIndex, rowIndex, cell.(int))
				case "float64":
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
						err := fmt.Errorf("could not convert YAML float64 to gotables %s", colType)
						return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), TopFuncName, err)
					}
				case "bool":
where(cell.(bool))
					err = table.SetBoolByColIndex(colIndex, rowIndex, cell.(bool))
//				case []interface{}: // This cell is a slice
				case "[]byte", "[]uint": // This cell is a slice
					var interfaceSlice []interface{} = cell.([]interface{})
					var byteSlice []byte = []byte{}
					for _, sliceVal := range interfaceSlice {
						byteSlice = append(byteSlice, byte(sliceVal.(float64)))
					}
					err = table.SetByteSliceByColIndex(colIndex, rowIndex, byteSlice)
					if err != nil {
						return nil, err
					}
//				case map[string]interface{}: // This cell is a table.
//				case *Table: // This cell is a table.
				case "*Table":
where("*Table")
					switch colType {
					case "*Table":
where()
						tableNested, err := newTableFromYAML_recursive(cell.(map[string]interface{}))
						if err != nil {
where()
							return nil, err
						}
where()
						err = table.SetTableByColIndex(colIndex, rowIndex, tableNested)
						if err != nil {
where()
							return nil, err
						}
					default:
where()
						return nil, fmt.Errorf("newTableFromYAML_recursive(): unexpected cell value at [%s].(%d,%d)",
							tableName, colIndex, rowIndex)
					}
				case "nil": // This cell is a nil table.
where()
					switch colType {
					case "*Table":
						var tableNested *Table = NewNilTable()
						err = table.SetTableByColIndex(colIndex, rowIndex, tableNested)
						if err != nil {
							return nil, err
						}
					default:
						return nil, fmt.Errorf("newTableFromYAML_recursive(): unexpected nil value at [%s].(%d,%d)",
							tableName, colIndex, rowIndex)
					}
					/*
						case time.Time:
							err = table.SetTimeByColIndex(colIndex, rowIndex, cell.(time.Time))
					*/
				default:
where(fmt.Sprintf("val type: %T", val))
					return nil, fmt.Errorf("%s %s: unexpected value of type: %v",
						UtilFuncSource(), TopFuncName, reflect.TypeOf(val))
				}
				// Single error handler for all the calls in this switch statement.
				if err != nil {
					return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), TopFuncName, err)
				}
			}
		}
	}

	return
}

func (tableSet *TableSet) GetTableSetAsYAML() (yamlString string, err error) {

	if tableSet == nil {
		return "", fmt.Errorf("%s tableSet.%s: table set is <nil>", UtilFuncSource(), UtilFuncName())
	}

	var yamlDoc map[string]interface{} = make(map[string]interface{}, 0)
	var yamlTables []map[string]interface{} = make([]map[string]interface{}, 0)
	var yamlObject map[string]interface{}	// Cell name and value pair.
	var yamlTableData [][]map[string]interface{}
	var yamlTableRow []map[string]interface{}
	var yamlTable map[string]interface{}

	var visitTableSet = func(tableSet *TableSet) (err error) {

		yamlDoc["tableSetName"] = tableSet.Name()

		return
	}

	var visitTable = func(table *Table) (err error) {

		// Used only in visitTable() function.
		var yamlTableMetadata = make([]interface{}, table.ColCount())

		yamlTable = make(map[string]interface{}, 0)
		yamlTableData = make([][]map[string]interface{}, table.RowCount())

		yamlTable["tableName"] = table.Name()
		yamlTable["data"] = yamlTableData

		if table.isStructShape {
			yamlTable["isStructShape"] = true
		}

		for colIndex := 0; colIndex < table.ColCount(); colIndex++ {

			yamlObject = make(map[string]interface{}, 0)

			if table.colTypes[colIndex] == "*Table" {
				// Quote "*Table" to avoid YAML interpreting it as an alias.
where()
				yamlObject[table.colNames[colIndex]] = fmt.Sprintf("%q", table.colTypes[colIndex])
			} else {
				yamlObject[table.colNames[colIndex]] = table.colTypes[colIndex]
			}
			yamlTableMetadata[colIndex] = yamlObject
		}

		yamlTable["metadata"] = yamlTableMetadata
		yamlTables = append(yamlTables, yamlTable)
		yamlDoc["tables"] = yamlTables

		return
	}

	var visitRow = func(row Row) (err error) {

		yamlTableRow = make([]map[string]interface{}, row.Table.ColCount())
		yamlTableData[row.RowIndex] = yamlTableRow

		return
	}

	var visitCell = func(cell Cell) (err error) {

		var valString string
		var anyVal interface{}
		yamlObject = make(map[string]interface{}, 1)

		switch cell.ColType {
		case "string":
			var s1 string
//			var s2 string
			s1, err = cell.Table.GetStringByColIndex(cell.ColIndex, cell.RowIndex)
//			s2 = fmt.Sprintf("%q", s1)
			anyVal = s1
//			anyVal = s2
			yamlObject[cell.ColName] = anyVal
		case "uint", "int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64", "float32", "float64":
			valString, err = cell.Table.GetValAsStringByColIndex(cell.ColIndex, cell.RowIndex)
			if err != nil {
				return err
			}
			yamlObject[cell.ColName] = valString
		case "bool":
			anyVal, err = cell.Table.GetBoolByColIndex(cell.ColIndex, cell.RowIndex)
			if err != nil {
				return err
			}
			yamlObject[cell.ColName] = anyVal
		case "byte":
			anyVal, err = cell.Table.GetByteByColIndex(cell.ColIndex, cell.RowIndex)
			if err != nil {
				return err
			}
			yamlObject[cell.ColName] = anyVal
		case "int":
			anyVal, err = cell.Table.GetIntByColIndex(cell.ColIndex, cell.RowIndex)
			if err != nil {
				return err
			}
			yamlObject[cell.ColName] = anyVal
/*
		case "float64":
			float64Val, err = cell.Table.GetFloat64ByColIndex(cell.ColIndex, cell.RowIndex)
			if err != nil {
				return err
			}
			yamlTableRow[cell.ColName] = float64Val
*/
		case "*Table":
where()
		default:
where()
			err = fmt.Errorf("%s: ERROR IN visitCell(): unknown type: %s\n", UtilFuncSource(), cell.ColType)
		}
where(err)
		// All errors in this switch are handled here.
		if err != nil {
			return err
		}

		yamlTableRow[cell.ColIndex] = yamlObject
where("yamlTableRow")
where(printSlice(yamlTableRow))
printYaml(nil, nil, yamlObject)
printYaml(nil, yamlTableRow, nil)
println()

		return
	}

	err = tableSet.Walk(visitTableSet, visitTable, visitRow, visitCell)
where(err)
	if err != nil {
		return "", nil
	}

where()
	var yamlBytes []byte
	yamlBytes, err = yaml.Marshal(yamlDoc)
	if err != nil {
		return "", nil
	}
	yamlString = string(yamlBytes)

where()
	return
}

func printYaml(yamlSlices []map[string][]map[string]interface{}, yamlArray []map[string]interface{}, yamlObject map[string]interface{}) {
	var out []byte
	var err error
	if yamlArray != nil {
		out, err = yaml.Marshal(yamlArray)
		if err != nil {
			println("PARSE ERROR")
		}
	} else if yamlSlices != nil {
		out, err = yaml.Marshal(yamlSlices)
		if err != nil {
			println("PARSE ERROR")
		}
	} else if yamlObject != nil {
		out, err = yaml.Marshal(yamlObject)
		if err != nil {
			println("PARSE ERROR")
		}
	} else {
		println("ARG ERROR!")
	}
	println("---\n" + string(out))
}

func printSlice(s []map[string]interface{}) string {
	return fmt.Sprintf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceOfSlice(s []map[string][]map[string]interface{}) string {
	return fmt.Sprintf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printMap(m map[string]interface{}) string {
	return fmt.Sprintf("len=%d %v\n", len(m), m)
}
