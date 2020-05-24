package gotables

import (
	"fmt"
	_ "os"
	"time"

	yaml "gopkg.in/yaml.v3"
)

/*
	Unmarshal YAML document to a *gotables.TableSet
*/
func NewTableSetFromYAML(yamlTableSetString string) (tableSet *TableSet, err error) {

/*
	Note: Although the yamlTableSetString may (or may not) be the same in each case:
	(1) GetTableSetAsYAML() produces a YAML document marshalled from:
	map[string]interface{}
		{"data":[][]interface{}{[]interface{}{11.1, 0x2, 97, 3}, []interface{}{22.2, 0x4, 98, 4}, []interface{}{33.3, 0x6, 99, 5}}
	(2) but is parsed as:
	map[string]interface{}
		{"data":[]  interface{}{[]interface{}{11.1, 2, 97, 3},   []interface{}{22.2, 4, 98, 4},   []interface{}{33.3, 6, 99, 5}}
	The critical difference is that:
		(i)  "data" in (1) is a 2-dimensional array (rows by cols), whereas
		(ii) "data" in (2) is an array of rows of maps.
	For this reason, NewTableSetFromYAML parses the unmarshalled yaml differently from the way GetTableSetAsYAML() generates it.
*/

	if yamlTableSetString == "" {
		return nil, fmt.Errorf("%s: yamlTableSetString is empty", UtilFuncName())
	}

//where()
	var m map[string]interface{}
	err = yaml.Unmarshal([]byte(yamlTableSetString), &m)
	if err != nil {
		return
}
// DATA PRESENT
//println()
//where("\n" + printMap(m))

	// (1) Retrieve and process TableSet name.
	var tableSetName string
	var exists bool
	tableSetName, exists = m["tableSetName"].(string)
	if !exists {
		return nil, fmt.Errorf("%s: in YAML doc 'tableSetName' is missing", UtilFuncName())
	}

	tableSet, err = NewTableSet(tableSetName)
	if err != nil {
		return
	}

// DATA PRESENT
where("\n" + tableSet.String())
	// (2) Retrieve and process tables.
	var tablesMap []interface{}
	tablesMap, exists = m["tables"].([]interface{})
	if !exists {
		return nil, fmt.Errorf("%s: in YAML doc 'tables' is missing", UtilFuncName())
	}

	var tableMap map[string]interface{}
	var tableMapInterface interface{}

	// (3) Loop through the array of tables.
	for _, tableMapInterface = range tablesMap {

		tableMap = tableMapInterface.(map[string]interface{})

		var table *Table
		table, err = newTableFromYAML_recursive(tableMap)
		if err != nil {
			return
		}
if table.Name() == "Tminus1" {
// DATA PRESENT
//where(fmt.Sprintf("\n%v", tableMap))
// DATA MISSING
where("\n" + table.String())
//os.Exit(32)
}
		err = tableSet.Append(table)
		if err != nil {
//where(err)
			return
		}
	}
//where("fff LOOP END:\n" + tableSet.String() + "\n")

//where()
	return
}

func newTableFromYAML_recursive(tableMap map[string]interface{}) (table *Table, err error) {
//where("LOOP newTableFromYAML_recursive")

// DATA PRESENT
//where(fmt.Sprintf("\n%v", tableMap))

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
	tableName, exists = tableMap["tableName"].(string)
//where(fmt.Sprintf("tableName exists = %t", exists))
//where("fff tableName:\n" + fmt.Sprintf("%v", tableName) + "\n\n")
	if !exists {
//where(tableMap)
		err = fmt.Errorf("in YAML doc: table is missing 'tableName'")
		return
	}
	if len(tableName) > 0 {
		table, err = NewTable(tableName)
		if err != nil {
			return
		}
	} else {
		table = NewNilTable()
		return
	}

	// If this optional isStructShape element is present, use it.
	var isStructShape bool
	isStructShape, exists = tableMap["isStructShape"].(bool)
//where(fmt.Sprintf("isStructShape exists = %t", exists))
//where("fff isStructShape:\n" + fmt.Sprintf("%v", isStructShape) + "\n\n")
	if exists {
		err = table.SetStructShape(isStructShape)
		if err != nil {
			return
		}
	}

	// (2) Retrieve and process metadata.
	var metadata []interface{}
	metadata, exists = tableMap["metadata"].([]interface{})
//where(fmt.Sprintf("metadata exists = %t", exists))
//where("fff metadata:\n" + fmt.Sprintf("%v", metadata) + "\n\n")
	if !exists {
		err = fmt.Errorf("in YAML doc table 'metadata' is missing")
		return
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
			err = fmt.Errorf("expecting col type value from YAML string value but got type %T: %v", typeVal, typeVal)
			return
		}

		colType = trimQuote(colType)	// YAML likes to quote some strings.
		err = table.AppendCol(colName, colType)
		if err != nil {
			table = nil
			return
		}
	}
where("\n" + table.String() + "\n")

	// (3) Retrieve and process data (if any).
	var data []interface{}

where("fff tableMap:\n" + fmt.Sprintf("%v",  tableMap) + "\n\n")
where("fff tableMap:\n" + fmt.Sprintf("%#v", tableMap) + "\n\n")
where("fff tableMap:\n" + fmt.Sprintf("%T",  tableMap) + "\n\n")
	data, exists = tableMap["data"].([]interface{})

//where("newTableFromYAML_recursive()")
where(fmt.Sprintf("data exists = %t", exists))
//where("fff data:\n" + fmt.Sprintf("%#v", data) + "\n\n")

//	var exists2 bool
//	data2, exists2 = tableMap["data"].([]interface{})
//where(fmt.Sprintf("data2 exists2 = %t", exists2))
UtilPrintCaller()
UtilPrintCallerCaller()
//where("fff data2:\n" + fmt.Sprintf("%#v", data2) + "\n\n")


	if !exists {
		// Zero rows in this table. That's okay.
		return
	}

	// Loop through the array of rows.
	var row []interface{}
	for rowIndex := 0; rowIndex < len(data); rowIndex++ {
		err = table.AppendRow()
		if err != nil {
			table = nil
			return
		}
		row = data[rowIndex].([]interface{})
where(row)
		for colIndex := 0; colIndex < len(row); colIndex++ {
if table.colNames[colIndex] == "bta" || table.colNames[colIndex] == "t" {
where(table.colNames[colIndex])
}
//where(fmt.Sprintf("rowIndex=%d: %v", rowIndex, row))

//		var row []interface{} = rowVal.([]interface{})
//		var row []interface{} = rowVal
//where(fmt.Sprintf("row    type: %T", row))
//where(fmt.Sprintf("row        : %#v", row))
//where(printMap(row))
			// where(fmt.Sprintf("\t\tcol [%d] %v", colIndex, val))

//where()
			switch row[colIndex].(type) {
			case int:
where(fmt.Sprintf("%9s int is type %9s %9T %v", table.colNames[colIndex], table.colTypes[colIndex],  row[colIndex], row[colIndex]))
				var intVal = row[colIndex].(int)
				switch table.colTypes[colIndex] {
				case "int":
where(intVal)
					err = table.SetIntByColIndex(colIndex, rowIndex, row[colIndex].(int))
				case "int8":
where(intVal)
					err = table.SetInt8ByColIndex(colIndex, rowIndex, int8(intVal))
				case "int16":
where(intVal)
					err = table.SetInt16ByColIndex(colIndex, rowIndex, int16(intVal))
				case "int32":
where(intVal)
					err = table.SetInt32ByColIndex(colIndex, rowIndex, int32(intVal))
				case "int64":
where(intVal)
					err = table.SetInt64ByColIndex(colIndex, rowIndex, int64(intVal))
				case "uint":
where(intVal)
					err = table.SetUintByColIndex(colIndex, rowIndex, row[colIndex].(uint))
				case "uint8":
where(intVal)
					err = table.SetUint8ByColIndex(colIndex, rowIndex, uint8(intVal))
				case "byte":
where(intVal)
					err = table.SetByteByColIndex(colIndex, rowIndex, uint8(intVal))
				case "uint16":
where(intVal)
					err = table.SetUint16ByColIndex(colIndex, rowIndex, uint16(intVal))
				case "uint32":
where(intVal)
					err = table.SetUint32ByColIndex(colIndex, rowIndex, uint32(intVal))
				case "uint64":
where(intVal)
					err = table.SetUint64ByColIndex(colIndex, rowIndex, uint64(intVal))
				case "float32":
where(intVal)
					err = table.SetFloat32ByColIndex(colIndex, rowIndex, float32(intVal))
//				case "float64":
//where(intVal)
//					err = table.SetFloat64ByColIndex(colIndex, rowIndex, float64(intVal))
//				case "string":
//where(intVal)
//					err = table.SetStringByColIndex(colIndex, rowIndex, row[colIndex].(string))
//				case "bool":
//where(intVal)
//					err = table.SetBoolByColIndex(colIndex, rowIndex, row[colIndex].(bool))
				case "time.Time":
where(intVal)
					err = table.SetTimeByColIndex(colIndex, rowIndex, row[colIndex].(time.Time))
				case "rune":
where(intVal)
					err = table.SetRuneByColIndex(colIndex, rowIndex, rune(intVal))
				default:
where(intVal)
					msg := invalidColTypeMsg(table.colTypes[colIndex])
					err = fmt.Errorf("%s: %s", UtilFuncName(), msg)
					table = nil
					return
				}
				// Error handler for all cases.
				if err != nil {
					table = nil
					return
				}
			case float64:
where(fmt.Sprintf("%9s int is type %9s %v", table.colNames[colIndex], table.colTypes[colIndex], row[colIndex]))
where(fmt.Sprintf("%T", row[colIndex]))
				switch table.colTypes[colIndex] {
				case "float32":
					var float64Val float64 = row[colIndex].(float64)
					err = table.SetFloat32ByColIndex(colIndex, rowIndex, float32(float64Val))
				case "float64":
					err = table.SetFloat64ByColIndex(colIndex, rowIndex, row[colIndex].(float64))
				default:
					msg := invalidColTypeMsg(table.colTypes[colIndex])
					err = fmt.Errorf("%s: %s", UtilFuncName(), msg)
					table = nil
					return
				}
			case string:
where(fmt.Sprintf("%9s int is type %9s %v", table.colNames[colIndex], table.colTypes[colIndex], row[colIndex]))
				var stringVal string = row[colIndex].(string)
				err = table.SetStringByColIndex(colIndex, rowIndex, stringVal)
			case bool:
where(fmt.Sprintf("%9s int is type %9s %v", table.colNames[colIndex], table.colTypes[colIndex], row[colIndex]))
				err = table.SetBoolByColIndex(colIndex, rowIndex, row[colIndex].(bool))
			case time.Time:
where(fmt.Sprintf("%9s time.Time is type %9s %v", table.colNames[colIndex], table.colTypes[colIndex], row[colIndex]))
				err = table.SetTimeByColIndex(colIndex, rowIndex, row[colIndex].(time.Time))
			case interface{}:
where()
				switch table.colTypes[colIndex] {
				case "[]byte":
where()
					var sliceVal []interface{} = row[colIndex].([]interface{})
					var byteSliceVal []byte = make([]byte, len(sliceVal))
					for i := 0; i < len(sliceVal); i++ {
						var faceVal interface{} = sliceVal[i]
						var intVal int = faceVal.(int)
						byteSliceVal[i] = byte(intVal)
					}
					err = table.SetByteSliceByColIndex(colIndex, rowIndex, byteSliceVal)
				case "[]uint8":
where()
					var sliceVal []interface{} = row[colIndex].([]interface{})
					var uint8SliceVal []uint8 = make([]uint8, len(sliceVal))
					for i := 0; i < len(sliceVal); i++ {
						var faceVal interface{} = sliceVal[i]
						var intVal int = faceVal.(int)
						uint8SliceVal[i] = uint8(intVal)
					}
					err = table.SetUint8SliceByColIndex(colIndex, rowIndex, uint8SliceVal)
				default:
where()
					msg := invalidColTypeMsg(table.colTypes[colIndex])
					err = fmt.Errorf("%s: %s", UtilFuncName(), msg)
					table = nil
					return
				}
				// Error handler for all cases.
				if err != nil {
					table = nil
					return
				}
			case []interface{}:
where(fmt.Sprintf("%9s fac is type %9s %9T %v", table.colNames[colIndex], table.colTypes[colIndex], row[colIndex], row[colIndex]))
			case []byte:
where(fmt.Sprintf("%9s []b is type %9s %9T %v", table.colNames[colIndex], table.colTypes[colIndex], row[colIndex], row[colIndex]))
			default:
where(fmt.Sprintf("default is type %9s %9T %v", table.colTypes[colIndex], row[colIndex], row[colIndex]))
where()
				msg := invalidColTypeMsg(table.colTypes[colIndex])
				err = fmt.Errorf("%s: %s", UtilFuncName(), msg)
				table = nil
				return
			}
		}
	}
//where("LOOP RECURSIVE \n" + table.String() + "\n")
/*
if table.Name() == "T0" {
var i int
i, err = table.GetInt("k", 1)
if err != nil {
	table = nil
	return
}
//where(i)
}
*/

	return
}

func (tableSet *TableSet) GetTableSetAsYAML() (yamlString string, err error) {

	if tableSet == nil {
		return "", fmt.Errorf("%s tableSet.%s: table set is <nil>", UtilFuncSource(), UtilFuncName())
	}

	var yamlMap map[string]interface{}
	yamlMap, err = tableSet.GetTableSetAsMap()
	if err != nil {
		return "", nil
	}

	var yamlBytes []byte
	yamlBytes, err = yaml.Marshal(yamlMap)
	if err != nil {
		return "", nil
	}

	yamlString = string(yamlBytes)

	return
}

func (tableSet *TableSet) GetTableSetAsMap() (yamlMap map[string]interface{}, err error) {

	if tableSet == nil {
		return nil, fmt.Errorf("%s tableSet.%s: table set is <nil>", UtilFuncSource(), UtilFuncName())
	}

	var yamlDoc map[string]interface{} = make(map[string]interface{}, 0)
	yamlDoc["tableSetName"] = tableSet.Name()

	var yamlTables []map[string]interface{} = make([]map[string]interface{}, tableSet.TableCount())
	for tableIndex := 0; tableIndex < tableSet.TableCount(); tableIndex++ {

		var table *Table
		table, err = tableSet.GetTableByTableIndex(tableIndex)
		if err != nil {
			return
		}
//where("fff INPUT:\n" + table.String() + "\n")

		var yamlTable map[string]interface{}
		yamlTable, err = table.getTableAsYAML()
		if err != nil {
			return
		}
//where("fff YAML:\n" + fmt.Sprintf("%v", yamlTable) + "\n\n")
//where("fff YAML:\n" + fmt.Sprintf("%#v", yamlTable) + "\n\n")
/*
var valid bool
valid, err = isValidYAML("", yamlTable)
//where(fmt.Sprintf("valid: %t err: %v", valid, err))
*/
/*
var tableOut *Table
var tableOut2 *Table
tableOut, err = newTableFromYAML_recursive(yamlTable)
if err != nil {
	return
}
//where("uuu OUTPUT:\n" + tableOut.String() + "\n")
//where("uuu OUTPUT2:\n" + tableOut2.String() + "\n")
*/

		yamlTables[tableIndex] = yamlTable
	}

	yamlDoc["tables"] = yamlTables
	yamlMap = yamlDoc

	return
}

func (table *Table) getTableAsYAML() (yamlTable map[string]interface{}, err error) {
//where("getTableAsYAML()")
//where("\n" + table.String() + "\n")

	var yamlObject map[string]interface{}	// Cell name and value pair.
	var yamlTableData [][]interface{}
	var yamlTableRow []interface{}

	var visitTable = func(table *Table) (err error) {

		// Used only in visitTable() function.
		var yamlTableMetadata = make([]interface{}, table.ColCount())

		yamlTable = make(map[string]interface{}, 0)
		yamlTableData = make([][]interface{}, table.RowCount())

		yamlTable["tableName"] = table.Name()
		yamlTable["data"] = yamlTableData

		if table.isStructShape {
			yamlTable["isStructShape"] = true
		}

		for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
//where(fmt.Sprintf("table: [%s] colName: %s colIndex: %d colType: %q", table.Name(), table.colNames[colIndex], colIndex, table.colTypes[colIndex]))
			yamlObject = make(map[string]interface{}, 0)
			if table.colTypes[colIndex] == "*Table" {
				// Quote "*Table" to avoid YAML interpreting it as an alias.
//where()
				yamlObject[table.colNames[colIndex]] = fmt.Sprintf("%q", table.colTypes[colIndex])
//where(yamlObject[table.colNames[colIndex]])
			} else {
				yamlObject[table.colNames[colIndex]] = table.colTypes[colIndex]
			}
			yamlTableMetadata[colIndex] = yamlObject
		}

//where(yamlTableMetadata)
		yamlTable["metadata"] = yamlTableMetadata
//where(yamlTable)
/*
		yamlTables = append(yamlTables, yamlTable)
		yamlDoc["tables"] = yamlTables
*/

// DOING:
/*
		if table.parentTable != nil {	// Not a top-level table.
			// Add this to the parent table's cell?
			var nestedTable *Table
			nestedTable, err = cell.Table.GetTableByColIndex(cell.ColIndex, cell.RowIndex)
			if err != nil {
				return err
			}
			yamlTableRow[cell.ColIndex] = anyVal
		}
*/

		return
	}

	var visitRow = func(row Row) (err error) {

		yamlTableRow = make([]interface{}, row.Table.ColCount())
		yamlTableData[row.RowIndex] = yamlTableRow

		return
	}

	var visitCell = func(cell Cell) (err error) {

		var anyVal interface{}
//		yamlObject = make(map[string]interface{}, 1)

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
		case "time.Time":
			anyVal, err = cell.Table.GetTimeByColIndex(cell.ColIndex, cell.RowIndex)
		case "*Table":
// DOING:
			var nestedTable *Table
			nestedTable, err = cell.Table.GetTableByColIndex(cell.ColIndex, cell.RowIndex)
//where(cell.Table.Name())
//where(nestedTable.Name())
println()
//where("\n" + nestedTable.String())
println()
//where(err)
			if err != nil {
				return err
			}

			var nestedTableYAML map[string]interface{}
			nestedTableYAML, err = nestedTable.getTableAsYAML()
			if err != nil {
				return err
			}
/*
//where(fmt.Sprintf("%v", nestedTableYAML))
var t *Table
t, err = newTableFromYAML_recursive(nestedTableYAML)
if err != nil {
	return err
}
if cell.Table.Name() == "T3" {
//where("WHAT? \n" + cell.Table.String() + "\n")
//where("WHAT? \n" + t.String() + "\n")
//os.Exit(66)
}
*/
			anyVal = nestedTableYAML

//where()
		default:
//where()
			msg := invalidColTypeMsg(cell.ColType)
//where(msg)
			err = fmt.Errorf("visitCell() YAML: %s", msg)
			return err
		}
		// All errors in this switch are handled here.
		if err != nil {
			return err
		}

		yamlTableRow[cell.ColIndex] = anyVal
//where(yamlTableRow)
/*
//where(printSlice(yamlTableRow))
//where("yamlObject")
printYaml(nil, nil, yamlObject)
printYaml(nil, yamlTableRow, nil)
println()
*/

		return
	}

	err = table.Walk(visitTable, visitRow, visitCell)
	if err != nil {
		return
	}
/*
var valid bool
valid, err = isValidYAML("", yamlTable)
//where(fmt.Sprintf("valid: %t err: %v", valid, err))
*/

/*
	var yamlBytes []byte
	yamlBytes, err = yaml.Marshal(yamlDoc)
	if err != nil {
		return "", nil
	}
	yamlString = string(yamlBytes)
*/
//where(fmt.Sprintf("%#v", yamlTable))
/*
var t *Table
t, err = newTableFromYAML_recursive(yamlTable)
if err != nil {
	return
}
//where("HOW?\n" + t.String() + "\n\n")
*/

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

func isValidYAML(yamlString string, yamlMap map[string]interface{}) (isValid bool, err error) {
where("yyy INSIDE isValidYAML()")

	if yamlString == "" && yamlMap == nil {
		err = fmt.Errorf("%s: yamlString and yamlMap are both empty", UtilFuncName())
	}

	var tableSet *TableSet

	if yamlString != "" {
where(fmt.Sprintf("yyy isValidYAML() checking yamlString:\n%s", yamlString))
		tableSet, err = NewTableSetFromYAML(yamlString)
		if err != nil {
			err = fmt.Errorf("isValidYAML(): %v", err)
where(fmt.Sprintf("isValidYAML(): %v\n%s", err, yamlString))
			return
		}

where(fmt.Sprintf("yyy isValidYAML() checking tableCounts"))
		var rowCount = 0
		for tableIndex := 0; tableIndex < tableSet.TableCount(); tableIndex++ {
			var table *Table
			table, err = tableSet.GetTableByTableIndex(tableIndex)
			if err != nil {
				err = fmt.Errorf("isValidYAML(): %v", err)
where(fmt.Sprintf("isValidYAML(): %v tableIndex: %d:\n%s", err, tableIndex, yamlString))
				return
			}

			rowCount += table.RowCount()
		}
		if rowCount == 0 {
			err = fmt.Errorf("%s: from yamlString: tableSet [[%s]] with %d tables has 0 rows",
				UtilFuncName(), tableSet.Name(), tableSet.TableCount())
where(fmt.Sprintf("isValidYAML(): %v\n%s", err, tableSet.String()))
			err = fmt.Errorf("isValidYAML(): %v", err)
			return
		}

		isValid = true
	}

	if yamlMap != nil {
where(fmt.Sprintf("yyy isValidYAML() checking yamlMap: yaml.Marshal(yamlMap):\n%v", yamlMap))
		var yamlBytes []byte
		yamlBytes, err = yaml.Marshal(yamlMap)
		if err != nil {
			err = fmt.Errorf("isValidYAML(): %v", err)
where(fmt.Sprintf("isValidYAML(): %v", err))
			return
		}
		yamlString = string(yamlBytes)

where(fmt.Sprintf("yyy isValidYAML() checking yamlMap: NewTableSetFromYAML(yamlString)"))
		tableSet, err = NewTableSetFromYAML(yamlString)
		if err != nil {
where(fmt.Sprintf("yyy isValidYAML() false checked yamlMap: NewTableSetFromYAML(yamlString): %v", err))
			// See if it is a table.
			errFromTableSet := err
			var table *Table
where(fmt.Sprintf("yyy isValidYAML() checking yamlMap: newTableFromYAML_recursive(yamlMap)"))
			table, err = newTableFromYAML_recursive(yamlMap)
			if err != nil {
where(fmt.Sprintf("isValidYAML(): false %v\n%s", err, yamlString))
				err = fmt.Errorf("isValidYAML(): %v (also %v)", err, errFromTableSet)
				return
			}
where(fmt.Sprintf("yyy isValidYAML() true checked yamlMap: newTableFromYAML_recursive(yamlMap) =\n%s", table.String()))
		}
where(fmt.Sprintf("yyy isValidYAML() true checked yamlMap: NewTableSetFromYAML(yamlString) =\n%s", tableSet.String()))

where(fmt.Sprintf("yyy isValidYAML() checking tableCounts"))
		var rowCount = 0
		for tableIndex := 0; tableIndex < tableSet.TableCount(); tableIndex++ {
			var table *Table
			table, err = tableSet.GetTableByTableIndex(tableIndex)
			if err != nil {
				err = fmt.Errorf("isValidYAML(): %v", err)
where(fmt.Sprintf("isValidYAML(): %v tableIndex: %d:\n%s", err, tableIndex, yamlString))
				return
			}

			rowCount += table.RowCount()
		}
		if rowCount == 0 {
			err = fmt.Errorf("%s: from yamlString: tableSet [[%s]] with %d tables has 0 rows",
				UtilFuncName(), tableSet.Name(), tableSet.TableCount())
where(fmt.Sprintf("isValidYAML(): %v\n%s", err, tableSet.String()))
			err = fmt.Errorf("isValidYAML(): %v", err)
			return
		}

		isValid = true
	}

	return
}

//					case "*Table":
//	//where("case table")
//	//where(fmt.Sprintf("table [%s]", table.Name()))
//	//where(fmt.Sprintf("row[%d] %#v type %T", colIndex, row[colIndex], row[colIndex]))
//	//where(row[colIndex])
//						var tableNested *Table
//						if row[colIndex] == nil {
//	//where()
//							tableNested = NewNilTable()
//						} else {
//	//where()
//	//where("LOOP CALLING RECURSIVELY")
//	/*		UNDELETE!!!
//							var mapVal map[string]interface{} = row[colIndex].(map[string]interface{})
//	// UNDELETE!!!						tableNested, err = newTableFromYAML_recursive(mapVal)
//							if err != nil {
//	//where()
//								return nil, err
//							}
//	*/
//	//where(fmt.Sprintf("tableNested = %s\n", tableNested.String()))
//						}
//	//where()
//						err = table.SetTableByColIndex(colIndex, rowIndex, tableNested)
//						if err != nil {
//	//where()
//							table = nil
//							return
//						}
//	//where()
