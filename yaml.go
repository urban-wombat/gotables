package gotables

import (
	_ "bytes"
	"fmt"
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
//println()
//where("\n" + printMap(m))

where()
	// (1) Retrieve and process TableSet name.
	var tableSetName string
	var exists bool
	tableSetName, exists = m["tableSetName"].(string)
	if !exists {
		return nil, fmt.Errorf("%s: YAML is missing tableSet name", UtilFuncName())
	}

where()
	tableSet, err = NewTableSet(tableSetName)
	if err != nil {
		return
	}

where()
	// (2) Retrieve and process tables.
	var tablesMap []interface{}
	tablesMap, exists = m["tables"].([]interface{})
	if !exists {
		return nil, fmt.Errorf("%s: YAML is missing tables", UtilFuncName())
	}

where()
	var tableMap map[string]interface{}
	var tableMapInterface interface{}

where()
	// (3) Loop through the array of tables.
	for _, tableMapInterface = range tablesMap {

where()
		tableMap = tableMapInterface.(map[string]interface{})
// where(printMap(tableMap))

where()
		var table *Table
		table, err = newTableFromYAML_recursive(tableMap)
		if err != nil {
			return
		}

where()
		err = tableSet.Append(table)
		if err != nil {
where(err)
			return
		}
	}

where()
	return
}

func newTableFromYAML_recursive(m map[string]interface{}) (table *Table, err error) {

//	const TopFuncName string = "NewTableFromYAML()"

where()
	var exists bool

where()
	/*
		We don't know the order map values will be returned if we iterate of the map:
		(1) tableName
		(2) isStructShape (if there)
		(3) metadata
		(4) data (if any)
		So we retrieve each of the 3 (possibly 2) top-level map values individually.
	*/

where()
	// (1) Retrieve and process table name.
	var tableName string
	tableName, exists = m["tableName"].(string)
where(tableName)
	if !exists {
where()
		return nil, fmt.Errorf("YAML is missing table name")
	}
where()
	if len(tableName) > 0 {
		table, err = NewTable(tableName)
		if err != nil {
where()
			return nil, err
		}
	} else {
		table = NewNilTable()
	}

where()
	// If this optional isStructShape element is present, use it.
	var isStructShape bool
	isStructShape, exists = m["isStructShape"].(bool)
	if exists {
		err = table.SetStructShape(isStructShape)
		if err != nil {
			return nil, err
		}
	}

where()
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

where()
where(fmt.Sprintf("[%s].AppendCol(%q, %q)", table.Name(), colName, colType))
		colType = trimQuote(colType)	// YAML likes to quote some strings.
where(fmt.Sprintf("[%s].AppendCol(%q, %q)", table.Name(), colName, colType))
		err = table.AppendCol(colName, colType)
where(err)
		if err != nil {
			return nil, err
		}
	}

where()
	// (3) Retrieve and process data (if any).
	var data []interface{}
	data, exists = m["data"].([]interface{})
	if !exists {
		// Zero rows in this table. That's okay.
		return table, nil
	}
where("\n" + table.String())

	// Loop through the array of rows.
	for rowIndex, rowVal := range data {
where(fmt.Sprintf("rowIndex=%d: %v", rowIndex, rowVal))
		err = table.AppendRow()
		if err != nil {
			return nil, err
		}

where()
		var row []interface{} = rowVal.([]interface{})
where(fmt.Sprintf("rowVal type: %T", rowVal))
where(fmt.Sprintf("rowVal     : %#v", rowVal))
where(fmt.Sprintf("row    type: %T", row))
//where(printMap(row))
//		for colIndex, val := range row {
		for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
			// where(fmt.Sprintf("\t\tcol [%d] %v", colIndex, val))

where()
			switch table.colTypes[colIndex] {
			case "int":
				err = table.SetIntByColIndex(colIndex, rowIndex, row[colIndex].(int))
			case "int8":
				err = table.SetInt8ByColIndex(colIndex, rowIndex, row[colIndex].(int8))
			case "int16":
				err = table.SetInt16ByColIndex(colIndex, rowIndex, row[colIndex].(int16))
			case "int32":
				err = table.SetInt32ByColIndex(colIndex, rowIndex, row[colIndex].(int32))
			case "int64":
				err = table.SetInt64ByColIndex(colIndex, rowIndex, row[colIndex].(int64))
			case "uint":
				err = table.SetUintByColIndex(colIndex, rowIndex, row[colIndex].(uint))
			case "uint8":
				// We know this is safe because the encoding was from a uint8 value.
				var intVal int = row[colIndex].(int)
				var uint8Val uint8 = uint8(intVal)
				err = table.SetUint8ByColIndex(colIndex, rowIndex, uint8Val)
			case "byte":
				// We know this is safe because the encoding was from a byte value.
				var intVal int = row[colIndex].(int)
				var byteVal byte = byte(intVal)
				err = table.SetByteByColIndex(colIndex, rowIndex, byteVal)
			case "uint16":
				err = table.SetUint16ByColIndex(colIndex, rowIndex, row[colIndex].(uint16))
			case "uint32":
				err = table.SetUint32ByColIndex(colIndex, rowIndex, row[colIndex].(uint32))
			case "uint64":
				err = table.SetUint64ByColIndex(colIndex, rowIndex, row[colIndex].(uint64))
			case "float32":
				// We know this is safe because the encoding was from a float32 value.
				var float64Val float64
				var float32Val float32
				var intVal int
				switch row[colIndex].(type) {
				case float64:
					float64Val = row[colIndex].(float64)
					float32Val = float32(float64Val)
				case int:	// I don't know why sometimes YAML stores a float32 as an int. But it does.
					intVal = row[colIndex].(int)
					float32Val = float32(intVal)
				default:
					err = fmt.Errorf("%s: invalid type from YAML", UtilFuncName())
					return nil, err
				}
					err = table.SetFloat32ByColIndex(colIndex, rowIndex, float32Val)
			case "float64":
				err = table.SetFloat64ByColIndex(colIndex, rowIndex, row[colIndex].(float64))
			case "string":
				err = table.SetStringByColIndex(colIndex, rowIndex, row[colIndex].(string))
			case "bool":
				err = table.SetBoolByColIndex(colIndex, rowIndex, row[colIndex].(bool))
			case "time.Time":
				err = table.SetTimeByColIndex(colIndex, rowIndex, row[colIndex].(time.Time))
			case "rune":
				var intVal int = row[colIndex].(int)
				var runeVal rune = rune(intVal)
				err = table.SetRuneByColIndex(colIndex, rowIndex, runeVal)
			case "*Table":
where(fmt.Sprintf("table [%s]", table.Name()))
where(fmt.Sprintf("row[%d] %v type %T", colIndex, row[colIndex], row[colIndex]))
where(row[colIndex])
				var tableNested *Table
				if row[colIndex] == nil {
					tableNested = NewNilTable()
				} else {
					var mapVal map[string]interface{} = row[colIndex].(map[string]interface{})
					tableNested, err = newTableFromYAML_recursive(mapVal)
					if err != nil {
where()
						return nil, err
					}
				}
where()
				err = table.SetTableByColIndex(colIndex, rowIndex, tableNested)
				if err != nil {
where()
					return nil, err
				}
where()
			case "[]byte":
where(table.Name())
where(rowIndex)
where(colIndex)
where(table.colNames[colIndex])
where(fmt.Sprintf("row[%d] = %v  type = %T", colIndex, row[colIndex], row[colIndex]))
				var sliceVal []interface{} = row[colIndex].([]interface{})
				var byteSliceVal []byte = make([]byte, len(sliceVal))
				for i := 0; i < len(sliceVal); i++ {
					var intVal = sliceVal[i].(int)
					var byteVal = byte(intVal)
					byteSliceVal[i] = byteVal
				}
				err = table.SetByteSliceByColIndex(colIndex, rowIndex, byteSliceVal)
			case "[]uint8":
				var sliceVal []interface{} = row[colIndex].([]interface{})
				var uint8SliceVal []uint8 = make([]uint8, len(sliceVal))
				for i := 0; i < len(sliceVal); i++ {
					var intVal = sliceVal[i].(int)
					var uint8Val = uint8(intVal)
					uint8SliceVal[i] = uint8Val
				}
				err = table.SetUint8SliceByColIndex(colIndex, rowIndex, uint8SliceVal)
			default:
where()
				msg := invalidColTypeMsg(table.colTypes[colIndex])
				err = fmt.Errorf("%s: %s", UtilFuncName(), msg)
				return nil, err
			}
			// Error handler for all cases.
			if err != nil {
where()
				return nil, err
			}
		}
where("\n" + table.String())

//			for colName, cell = range val.(map[string]interface{}) {
//where(fmt.Sprintf("colName  = %q", colName))
//where(fmt.Sprintf("colIndex = %d", colIndex))
//where(fmt.Sprintf("rowIndex = %d", rowIndex))
//where(fmt.Sprintf("cell     = %v", cell))
//				// There's only one map element here: colName and colType.
//				// where(fmt.Sprintf("\t\t\tcol=%d row=%d celltype=%T cell=%v", colIndex, rowIndex, cell, cell))
//
//				var colType string = table.colTypes[colIndex]
//where(fmt.Sprintf("colType  = %q", colType))
//where(fmt.Sprintf("cell type: %T", cell))
//println()
////				switch cell.(type) {
//				switch colType {
//				case "string":
//					switch colType { // We need to convert time string format to time.Time
//					case "time.Time":
//						var timeVal time.Time
//						timeVal, err = time.Parse(time.RFC3339, cell.(string))
//						if err != nil { // We need this extra error check here
//							err := fmt.Errorf("could not convert YAML time string to gotables %s", colType)
//							return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), TopFuncName, err)
//						}
//						err = table.SetTimeByColIndex(colIndex, rowIndex, timeVal)
//						if err != nil {
//							err := fmt.Errorf("could not convert YAML string to gotables %s", colType)
//							return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), TopFuncName, err)
//						}
//					default: // Is a string
//						err = table.SetStringByColIndex(colIndex, rowIndex, cell.(string))
//					}
//					// Single error handler for all the calls in this switch statement.
//					if err != nil {
//						err := fmt.Errorf("could not convert YAML string to gotables %s", colType)
//						return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), TopFuncName, err)
//					}
//				case "int":
//					err = table.SetIntByColIndex(colIndex, rowIndex, cell.(int))
//				case "float64":
//					switch colType { // We need to convert them back to gotables numeric types
//					case "int":
//						err = table.SetIntByColIndex(colIndex, rowIndex, int(cell.(float64)))
//					case "uint":
//						err = table.SetUintByColIndex(colIndex, rowIndex, uint(cell.(float64)))
//					case "byte":
//						err = table.SetByteByColIndex(colIndex, rowIndex, byte(cell.(float64)))
//					case "int8":
//						err = table.SetInt8ByColIndex(colIndex, rowIndex, int8(cell.(float64)))
//					case "int16":
//						err = table.SetInt16ByColIndex(colIndex, rowIndex, int16(cell.(float64)))
//					case "int32":
//						err = table.SetInt32ByColIndex(colIndex, rowIndex, int32(cell.(float64)))
//					case "int64":
//						err = table.SetInt64ByColIndex(colIndex, rowIndex, int64(cell.(float64)))
//					case "uint8":
//						err = table.SetUint8ByColIndex(colIndex, rowIndex, uint8(cell.(float64)))
//					case "uint16":
//						err = table.SetUint16ByColIndex(colIndex, rowIndex, uint16(cell.(float64)))
//					case "uint32":
//						err = table.SetUint32ByColIndex(colIndex, rowIndex, uint32(cell.(float64)))
//					case "uint64":
//						err = table.SetUint64ByColIndex(colIndex, rowIndex, uint64(cell.(float64)))
//					case "float32":
//						err = table.SetFloat32ByColIndex(colIndex, rowIndex, float32(cell.(float64)))
//					case "float64":
//						err = table.SetFloat64ByColIndex(colIndex, rowIndex, float64(cell.(float64)))
//					}
//					// Single error handler for all the calls in this switch statement.
//					if err != nil {
//						err := fmt.Errorf("could not convert YAML float64 to gotables %s", colType)
//						return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), TopFuncName, err)
//					}
//				case "bool":
//where(cell.(bool))
//					err = table.SetBoolByColIndex(colIndex, rowIndex, cell.(bool))
////				case []interface{}: // This cell is a slice
//				case "[]byte", "[]uint": // This cell is a slice
//					var interfaceSlice []interface{} = cell.([]interface{})
//					var byteSlice []byte = []byte{}
//					for _, sliceVal := range interfaceSlice {
//						byteSlice = append(byteSlice, byte(sliceVal.(float64)))
//					}
//					err = table.SetByteSliceByColIndex(colIndex, rowIndex, byteSlice)
//					if err != nil {
//						return nil, err
//					}
////				case map[string]interface{}: // This cell is a table.
////				case *Table: // This cell is a table.
//				case "*Table":
//where("*Table")
//					switch colType {
//					case "*Table":
//where()
//						tableNested, err := newTableFromYAML_recursive(cell.(map[string]interface{}))
//						if err != nil {
//where()
//							return nil, err
//						}
//where()
//						err = table.SetTableByColIndex(colIndex, rowIndex, tableNested)
//						if err != nil {
//where()
//							return nil, err
//						}
//					default:
//where()
//						return nil, fmt.Errorf("newTableFromYAML_recursive(): unexpected cell value at [%s].(%d,%d)",
//							tableName, colIndex, rowIndex)
//					}
//				case "nil": // This cell is a nil table.
//where()
//					switch colType {
//					case "*Table":
//						var tableNested *Table = NewNilTable()
//						err = table.SetTableByColIndex(colIndex, rowIndex, tableNested)
//						if err != nil {
//							return nil, err
//						}
//					default:
//						return nil, fmt.Errorf("newTableFromYAML_recursive(): unexpected nil value at [%s].(%d,%d)",
//							tableName, colIndex, rowIndex)
//					}
//					/*
//						case time.Time:
//							err = table.SetTimeByColIndex(colIndex, rowIndex, cell.(time.Time))
//					*/
//				default:
//where(fmt.Sprintf("val type: %T", val))
//					return nil, fmt.Errorf("%s %s: unexpected value of type: %v",
//						UtilFuncSource(), TopFuncName, reflect.TypeOf(val))
//				}
//				// Single error handler for all the calls in this switch statement.
//				if err != nil {
//					return nil, fmt.Errorf("%s %s: %v", UtilFuncSource(), TopFuncName, err)
//				}
//			}
//		}
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
//	var yamlTableData [][]map[string]interface{}
	var yamlTableData [][]interface{}
//	var yamlTableRow []map[string]interface{}
	var yamlTableRow []interface{}
	var yamlTable map[string]interface{}

	var visitTableSet = func(tableSet *TableSet) (err error) {

		yamlDoc["tableSetName"] = tableSet.Name()

		return
	}

	var visitTable = func(table *Table) (err error) {

		// Used only in visitTable() function.
		var yamlTableMetadata = make([]interface{}, table.ColCount())

		yamlTable = make(map[string]interface{}, 0)
//		yamlTableData = make([][]map[string]interface{}, table.RowCount())
		yamlTableData = make([][]interface{}, table.RowCount())

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

//		yamlTableRow = make([]map[string]interface{}, row.Table.ColCount())
		yamlTableRow = make([]interface{}, row.Table.ColCount())
		yamlTableData[row.RowIndex] = yamlTableRow

		return
	}

	var visitCell = func(cell Cell) (err error) {

		var anyVal interface{}
		yamlObject = make(map[string]interface{}, 1)

		switch cell.ColType {
		case "string":
			anyVal, err = cell.Table.GetStringByColIndex(cell.ColIndex, cell.RowIndex)
		case "bool":
			anyVal, err = cell.Table.GetBoolByColIndex(cell.ColIndex, cell.RowIndex)
		case "rune":
			anyVal, err = cell.Table.GetRuneByColIndex(cell.ColIndex, cell.RowIndex)
		case "byte":
			anyVal, err = cell.Table.GetByteByColIndex(cell.ColIndex, cell.RowIndex)
		case "int":
			anyVal, err = cell.Table.GetIntByColIndex(cell.ColIndex, cell.RowIndex)
		case "int8":
			anyVal, err = cell.Table.GetInt8ByColIndex(cell.ColIndex, cell.RowIndex)
		case "int16":
			anyVal, err = cell.Table.GetInt16ByColIndex(cell.ColIndex, cell.RowIndex)
		case "int32":
			anyVal, err = cell.Table.GetInt32ByColIndex(cell.ColIndex, cell.RowIndex)
		case "int64":
			anyVal, err = cell.Table.GetInt64ByColIndex(cell.ColIndex, cell.RowIndex)
		case "uint":
			anyVal, err = cell.Table.GetUintByColIndex(cell.ColIndex, cell.RowIndex)
		case "uint8":
			anyVal, err = cell.Table.GetUint8ByColIndex(cell.ColIndex, cell.RowIndex)
where(fmt.Sprintf("uint8 anyVal %v type %T", anyVal, anyVal))
		case "uint16":
			anyVal, err = cell.Table.GetUint16ByColIndex(cell.ColIndex, cell.RowIndex)
		case "uint32":
			anyVal, err = cell.Table.GetUint32ByColIndex(cell.ColIndex, cell.RowIndex)
		case "uint64":
			anyVal, err = cell.Table.GetUint64ByColIndex(cell.ColIndex, cell.RowIndex)
		case "float32":
			anyVal, err = cell.Table.GetFloat32ByColIndex(cell.ColIndex, cell.RowIndex)
		case "float64":
			anyVal, err = cell.Table.GetFloat64ByColIndex(cell.ColIndex, cell.RowIndex)
		case "[]uint8":
			anyVal, err = cell.Table.GetUint8SliceByColIndex(cell.ColIndex, cell.RowIndex)
		case "[]byte":
			anyVal, err = cell.Table.GetByteSliceByColIndex(cell.ColIndex, cell.RowIndex)
where(fmt.Sprintf("[]byte anyVal %v type %T", anyVal, anyVal))
		case "time.Time":
			anyVal, err = cell.Table.GetTimeByColIndex(cell.ColIndex, cell.RowIndex)
		case "*Table":
/*
			var nestedTable *Table
			nestedTable, err = cell.Table.GetTableByColIndex(cell.ColIndex, cell.RowIndex)
			if err != nil {
				return err
			}
*/

// TO DO:
where()
		default:
where()
			msg := invalidColTypeMsg(cell.ColType)
where(msg)
			err = fmt.Errorf("visitCell() YAML: %s", msg)
			return err
		}
		// All errors in this switch are handled here.
		if err != nil {
			return err
		}

		/*
		yamlObject[cell.ColName] = anyVal
		yamlTableRow[cell.ColIndex] = yamlObject
		*/
		yamlTableRow[cell.ColIndex] = anyVal
/*
where("yamlTableRow")
where(printSlice(yamlTableRow))
where("yamlObject")
printYaml(nil, nil, yamlObject)
printYaml(nil, yamlTableRow, nil)
println()
*/

		return
	}

	err = tableSet.Walk(visitTableSet, visitTable, visitRow, visitCell)
	if err != nil {
		return "", nil
	}

	var yamlBytes []byte
	yamlBytes, err = yaml.Marshal(yamlDoc)
	if err != nil {
		return "", nil
	}
	yamlString = string(yamlBytes)

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

func trimQuote(s string) string {
	// From: https://stackoverflow.com/questions/44222554/how-to-remove-quotes-from-around-a-string-in-golang

	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}
