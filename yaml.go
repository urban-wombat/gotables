package gotables

import (
	"fmt"
	"encoding/json"
	_ "os"
	"strings"
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

	var tableSetMap map[string]interface{}
	err = yaml.Unmarshal([]byte(yamlTableSetString), &tableSetMap)
	if err != nil {
		return
	}

	// (1) Retrieve and process TableSet name.
	var tableSetName string
	var exists bool
	tableSetName, exists = tableSetMap["tableSetName"].(string)
	if !exists {
		return nil, fmt.Errorf("%s %s: in YAML doc: 'tableSetName' is missing", UtilFuncSource(), UtilFuncName())
	}

	tableSet, err = NewTableSet(tableSetName)
	if err != nil {
		return
	}

	// (2) Retrieve and process tables.
	var tablesMap []interface{}
	var anyVal interface{}
	anyVal, exists = tableSetMap["tables"]
	tablesMap = anyVal.([]interface{})

	if !exists {
		return nil, fmt.Errorf("%s %s: in YAML doc: 'tables' is missing", UtilFuncSource(), UtilFuncName())
	}

	var tableMap map[string]interface{}

	// (3) Loop through the array of tables.
	var tableIndex int
	for tableIndex, anyVal = range tablesMap {
where(fmt.Sprintf("tableIndex:%d anyVal:%v", tableIndex, anyVal))
println(" ttt ")
where(fmt.Sprintf(" ttt %v", anyVal))
where(fmt.Sprintf(" ttt %#v", anyVal))

		tableMap = anyVal.(map[string]interface{})
where(fmt.Sprintf(" ttt %v", tableMap))
//where(fmt.Sprintf(" ttt %#v", tableMap))

		var table *Table
		table, err = newTableFromYAML_recursive(tableMap)
where()
		if err != nil {
where()
			return
		}
where(table)
println()
where(fmt.Sprintf("RowCount()=%d ttt %s", table.RowCount(), table.String()))
println()
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

println()
where(fmt.Sprintf("yyy %s", tableSet.String()))
println()

	return
}

func newTableFromYAML_recursive(tableMap map[string]interface{}) (table *Table, err error) {
where("func newTableFromYAML_recursive()")
where(fmt.Sprintf("tableMap: %v", tableMap))
var marshalledBytes []byte
marshalledBytes, err = json.MarshalIndent(tableMap, "", "  ")
where(fmt.Sprintf("ttt %s", string(marshalledBytes)))

// DATA PRESENT
//where(fmt.Sprintf("\n%v", tableMap))

where()
	var exists bool

	/*
		We don't know the order map values will be returned if we iterate over the map:
		(1) tableName
		(2) isStructShape (if there)
		(3) metadata
		(4) data (if any)
		So we retrieve each of the 3 (possibly 2) top-level map values individually.
	*/

	// (1) Retrieve and process table name.
	var tableName string
	tableName, exists = tableMap["tableName"].(string)
where(fmt.Sprintf("tableName exists = %t", exists))
//where("fff tableName:\n" + fmt.Sprintf("%v", tableName) + "\n\n")
	if !exists {
//where(tableMap)
		err = fmt.Errorf("%s %s: in YAML doc: table is missing 'tableName'", UtilFuncSource(), UtilFuncName())
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

where()
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

where()
	// (2) Retrieve and process metadata.
	var metadata []interface{}
	metadata, exists = tableMap["metadata"].([]interface{})
//where(fmt.Sprintf("metadata exists = %t", exists))
//where("fff metadata:\n" + fmt.Sprintf("%v", metadata) + "\n\n")
	if !exists {
		err = fmt.Errorf("%s %s: in YAML doc: table 'metadata' is missing", UtilFuncSource(), UtilFuncName())
		return
	}
	// Loop through the array of metadata.
	for _, colNameAndType := range metadata {
where(fmt.Sprintf("colNameAndType type: %T", colNameAndType))
		var colName string
		var colType string
		var typeVal interface{}
		for colName, typeVal = range colNameAndType.(map[string]interface{}) {
			// There's only one map element here: colName and colType.
		}
		colType, ok := typeVal.(string)
		if !ok {
			err = fmt.Errorf("expecting colType string value from YAML but got type %T and value: %v", typeVal, typeVal)
			return
		}

where()
		colType = trimQuote(colType)	// YAML likes to quote some strings.
		err = table.AppendCol(colName, colType)
		if err != nil {
			table = nil
			return
		}
	}
where("\n" + table.String() + "\n")

where()
	// (3) Retrieve and process data (if any).
	var data [][]interface{}

where("fff tableMap:\n" + fmt.Sprintf("%v",  tableMap) + "\n\n")
where("fff tableMap:\n" + fmt.Sprintf("%#v", tableMap) + "\n\n")
where("fff tableMap:\n" + fmt.Sprintf("%T",  tableMap) + "\n\n")
UtilPrintCaller()

marshalledBytes, err = json.MarshalIndent(tableMap, "", "  ")
where(fmt.Sprintf("ttt %s", string(marshalledBytes)))

where(exists)
	data, exists = tableMap["data"].([][]interface{})
	whatever, _ := tableMap["data"]

where(fmt.Sprintf("data type: %T", data))

marshalledBytes, err = json.MarshalIndent(data, "", "  ")

	var dataMapSlice []interface{}
	var dataMapSliceSlice [][]interface{}

	switch whatever.(type) {
	case [][]interface{}:
where(fmt.Sprintf("case 1 %#v", whatever))
		dataMapSliceSlice = whatever.([][]interface{})
		data = dataMapSliceSlice
	case []interface{}:
		// Rewrite this nominal slice of interface{} as [][]interface{}
		// Underneath it is "really" [][]interface{}
		// Why this is so is a real puzzler, but this solution works.
		dataMapSlice = whatever.([]interface{})
		data = make([][]interface{}, len(dataMapSlice))
		for i, el1 := range dataMapSlice {
			for _, el2 := range el1.([]interface{}) {
				data[i] = append(data[i], el2)
			}
		}
	default:
	}


	whatever, exists = existsInMap(tableMap, "data")

UtilPrintCaller()

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
		row = data[rowIndex]
		for colIndex := 0; colIndex < len(row); colIndex++ {
			switch row[colIndex].(type) {
			case uint:
				err = table.SetUintByColIndex(colIndex, rowIndex, row[colIndex].(uint))
			case int:
				var intVal = row[colIndex].(int)
				switch table.colTypes[colIndex] {
				case "int":
					err = table.SetIntByColIndex(colIndex, rowIndex, row[colIndex].(int))
				case "uint":
					var intVal int = row[colIndex].(int)
					var uintVal uint = uint(intVal)
					err = table.SetUintByColIndex(colIndex, rowIndex, uintVal)
				case "int8":
					err = table.SetInt8ByColIndex(colIndex, rowIndex, int8(intVal))
				case "int16":
					err = table.SetInt16ByColIndex(colIndex, rowIndex, int16(intVal))
				case "int32":
					err = table.SetInt32ByColIndex(colIndex, rowIndex, int32(intVal))
				case "int64":
					err = table.SetInt64ByColIndex(colIndex, rowIndex, int64(intVal))
				case "uint8":
					err = table.SetUint8ByColIndex(colIndex, rowIndex, uint8(intVal))
				case "byte":
					err = table.SetByteByColIndex(colIndex, rowIndex, uint8(intVal))
				case "uint16":
					err = table.SetUint16ByColIndex(colIndex, rowIndex, uint16(intVal))
				case "uint32":
					err = table.SetUint32ByColIndex(colIndex, rowIndex, uint32(intVal))
				case "uint64":
					err = table.SetUint64ByColIndex(colIndex, rowIndex, uint64(intVal))
				case "float32":
					err = table.SetFloat32ByColIndex(colIndex, rowIndex, float32(intVal))
				case "time.Time":
					err = table.SetTimeByColIndex(colIndex, rowIndex, row[colIndex].(time.Time))
				case "rune":
					err = table.SetRuneByColIndex(colIndex, rowIndex, rune(intVal))
				default:
					msg := invalidColTypeMsg(table.Name(), table.colTypes[colIndex])
					err = fmt.Errorf("%s %s: %s", UtilFuncSource(), UtilFuncName(), msg)
					table = nil
					return
				}
				// Error handler for all cases.
				if err != nil {
					table = nil
					return
				}
			case float32:
				var float32Val float32 = row[colIndex].(float32)
				err = table.SetFloat32ByColIndex(colIndex, rowIndex, float32Val)
			case float64:
				switch table.colTypes[colIndex] {
				case "float32":
					var float64Val float64 = row[colIndex].(float64)
					err = table.SetFloat32ByColIndex(colIndex, rowIndex, float32(float64Val))
				case "float64":
					err = table.SetFloat64ByColIndex(colIndex, rowIndex, row[colIndex].(float64))
				default:
					msg := invalidColTypeMsg(table.Name(), table.colTypes[colIndex])
					err = fmt.Errorf("%s %s: %s", UtilFuncSource(), UtilFuncName(), msg)
					table = nil
					return
				}
			case byte:
				var byteVal byte = row[colIndex].(byte)
				err = table.SetByteByColIndex(colIndex, rowIndex, byteVal)
			case rune:
				var runeVal rune = row[colIndex].(rune)
				err = table.SetRuneByColIndex(colIndex, rowIndex, runeVal)
			case string:
				var stringVal string = row[colIndex].(string)
				err = table.SetStringByColIndex(colIndex, rowIndex, stringVal)
			case bool:
				err = table.SetBoolByColIndex(colIndex, rowIndex, row[colIndex].(bool))
			case time.Time:
				err = table.SetTimeByColIndex(colIndex, rowIndex, row[colIndex].(time.Time))
			case interface{}:
				switch table.colTypes[colIndex] {
				case "byte":
					err = table.SetByteByColIndex(colIndex, rowIndex, row[colIndex].(byte))

				case "uint":
					var uint64Val uint64 = row[colIndex].(uint64)
					err = table.SetUintByColIndex(colIndex, rowIndex, uint(uint64Val))

				case "uint16":
					err = table.SetUint16ByColIndex(colIndex, rowIndex, row[colIndex].(uint16))

				case "uint32":
					err = table.SetUint32ByColIndex(colIndex, rowIndex, row[colIndex].(uint32))

				case "uint64":
					err = table.SetUint64ByColIndex(colIndex, rowIndex, row[colIndex].(uint64))

				case "int8":
					err = table.SetInt8ByColIndex(colIndex, rowIndex, row[colIndex].(int8))

				case "int16":
					err = table.SetInt16ByColIndex(colIndex, rowIndex, row[colIndex].(int16))

				case "int64":
					err = table.SetInt64ByColIndex(colIndex, rowIndex, row[colIndex].(int64))

				case "*Table":
					var tableOut *Table
					tableOut, err = newTableFromYAML_recursive(row[colIndex].(map[string]interface{}))
					if err != nil {
						return
					}
					err = table.SetTableByColIndex(colIndex, rowIndex, tableOut)

				case "[]byte", "[]uint8":
					var sliceVal []interface{}
					var byteSliceVal []byte

					switch row[colIndex].(type) {
						case []byte:
							byteSliceVal = row[colIndex].([]byte)
						case []interface{}:
							sliceVal = row[colIndex].([]interface{})
							byteSliceVal = make([]byte, len(sliceVal))
							for i := 0; i < len(sliceVal); i++ {
								var intVal int = sliceVal[i].(int)
								byteSliceVal[i] = byte(intVal)
							}
						case interface{}:
						default:
					}
					switch table.colTypes[colIndex] {
						case "[]byte":
							err = table.SetByteSliceByColIndex(colIndex, rowIndex, byteSliceVal)
						case "[]uint8":
							err = table.SetUint8SliceByColIndex(colIndex, rowIndex, byteSliceVal)
						default:
							msg := invalidColTypeMsg(table.Name(), table.colTypes[colIndex])
							err = fmt.Errorf("#1 %s %s: %s", UtilFuncSource(), UtilFuncName(), msg)
							table = nil
							return
					}

				default:
					msg := invalidColTypeMsg(table.Name(), table.colTypes[colIndex])
					err = fmt.Errorf("#2 %s %s: %s", UtilFuncSource(), UtilFuncName(), msg)
					table = nil
					return
				}
				// #1 Error handler for all cases.
				if err != nil {
					table = nil
					return
				}
			case []interface{}:
			case []byte:
			case *Table:
			default:
				msg := invalidColTypeMsg(table.Name(), table.colTypes[colIndex])
				err = fmt.Errorf("#3 %s %s: %s", UtilFuncSource(), UtilFuncName(), msg)
				table = nil
				return
			}
			// #3 Error handler for all cases.
			if err != nil {
				table = nil
where("#2 Error handler for all cases.")
				return
			}
		}
where(table.String())
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
where(table.String())

	return
}

func (tableSet *TableSet) GetTableSetAsYAML() (yamlString string, err error) {
where("func GetTableSetAsYAML()")

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
where("func GetTableSetAsMap()")
where("this func: ", UtilFuncName(), " caller: ", UtilFuncCaller(), " caller caller: ", UtilFuncCallerCaller())

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
		yamlTable, err = table.getTableAsMap()
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
println(" ttt ")
where(fmt.Sprintf("ttt %v", yamlTable))
where(fmt.Sprintf("ttt %#v", yamlTable))
var tableOut *Table
var tableOut2 *Table
tableOut, err = newTableFromYAML_recursive(yamlTable)
if err != nil {
	return
}
println()
where(fmt.Sprintf("RowCount()=%d ttt %s", tableOut.RowCount(), tableOut.String())
println()
//where("uuu OUTPUT:\n" + tableOut.String() + "\n")
//where("uuu OUTPUT2:\n" + tableOut2.String() + "\n")
*/

		yamlTables[tableIndex] = yamlTable
	}

	yamlDoc["tables"] = yamlTables
	yamlMap = yamlDoc

	return
}

func (table *Table) getTableAsMap() (yamlTable map[string]interface{}, err error) {
where("BEGIN func getTableAsMap(): ", table.Name())
where("\n" + table.String() + "\n")
where(UtilFuncCaller())

	var yamlObject map[string]interface{}	// Cell name and value pair.
	var yamlTableData [][]interface{}
	var yamlTableRow []interface{}

	var visitTable = func(table *Table) (err error) {
where("visitTable()")

		// Used (and re-used) only in visitTable() function.
		var yamlTableMetadata = make([]interface{}, table.ColCount())
where()

		yamlTable = make(map[string]interface{}, 0)
		yamlTableData = make([][]interface{}, table.RowCount())
where()

		yamlTable["tableName"] = table.Name()
		yamlTable["data"] = yamlTableData

		if table.isStructShape {
			yamlTable["isStructShape"] = true
		}

		// Build metadata map.
		for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
//where(fmt.Sprintf("table: [%s] colName: %s colIndex: %d colType: %q", table.Name(), table.colNames[colIndex], colIndex, table.colTypes[colIndex]))
			yamlObject = make(map[string]interface{}, 0)
			if table.colTypes[colIndex] == "*Table" {
				// Quote "*Table" to avoid YAML interpreting it as an alias.
				yamlObject[table.colNames[colIndex]] = fmt.Sprintf("%q", table.colTypes[colIndex])
//where(yamlObject[table.colNames[colIndex]])
			} else {
				yamlObject[table.colNames[colIndex]] = table.colTypes[colIndex]
			}
			yamlTableMetadata[colIndex] = yamlObject
		}

where(yamlTableMetadata)
println()
		yamlTable["metadata"] = yamlTableMetadata
where(yamlTable)
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
where("visitRow()")

		// Make a fresh yamlTableRow.
		yamlTableRow = make([]interface{}, row.Table.ColCount())
where(fmt.Sprintf("should be blank: yamlTableRow = %v", yamlTableRow))

		// Assign it to yamlTableData.
		yamlTableData[row.RowIndex] = yamlTableRow

// Check to see if previous row has been assigned to table data.
if row.RowIndex > 0 {
where(fmt.Sprintf("previous row should be populated: yamlTableData[%d] = %v", row.RowIndex-1, yamlTableData[row.RowIndex-1]))
}

		return
	}

	var visitCell = func(walkDeep bool, cell Cell) (err error) {
where("visitCell()")
where(table.Name())

		var anyVal interface{}
//		yamlObject = make(map[string]interface{}, 1)

where("switch cell.ColType")
println()
		switch cell.ColType {
		case "string":
			anyVal, err = cell.Table.GetStringByColIndex(cell.ColIndex, cell.RowIndex)
var sss string = anyVal.(string)
if strings.HasPrefix(sss, "sss") {
where(fmt.Sprintf("ggg %q", sss))
}
where(anyVal)
println()
		case "bool":
			anyVal, err = cell.Table.GetBoolByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "rune":
			anyVal, err = cell.Table.GetRuneByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "byte":
			anyVal, err = cell.Table.GetByteByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "int":
			anyVal, err = cell.Table.GetIntByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "int8":
			anyVal, err = cell.Table.GetInt8ByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "int16":
			anyVal, err = cell.Table.GetInt16ByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "int32":
			anyVal, err = cell.Table.GetInt32ByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "int64":
			anyVal, err = cell.Table.GetInt64ByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "uint":
			anyVal, err = cell.Table.GetUintByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "uint8":
			anyVal, err = cell.Table.GetUint8ByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "uint16":
			anyVal, err = cell.Table.GetUint16ByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "uint32":
			anyVal, err = cell.Table.GetUint32ByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "uint64":
			anyVal, err = cell.Table.GetUint64ByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "float32":
			anyVal, err = cell.Table.GetFloat32ByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "float64":
			anyVal, err = cell.Table.GetFloat64ByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "[]uint8":
			anyVal, err = cell.Table.GetUint8SliceByColIndex(cell.ColIndex, cell.RowIndex)
where(anyVal)
println()
		case "[]byte":
			var byteSlice []byte
			byteSlice, err = cell.Table.GetByteSliceByColIndex(cell.ColIndex, cell.RowIndex)
			anyVal = byteSlice
			for i := 0; i < len(byteSlice); i++ {
				where(fmt.Sprintf("byteSlice[%d] type = %T", i, byteSlice[i]))
			}
println()
		case "time.Time":
			anyVal, err = cell.Table.GetTimeByColIndex(cell.ColIndex, cell.RowIndex)
where(fmt.Sprintf("[%-8s] col:%2d row:%2d anyVal:%v", cell.Table.Name(), cell.ColIndex, cell.RowIndex, anyVal))
println()
		case "*Table":
			var nestedTable *Table
			nestedTable, err = cell.Table.GetTableByColIndex(cell.ColIndex, cell.RowIndex)
			if err != nil {
				return err
			}

			anyVal, err = nestedTable.getTableAsMap()
			if err != nil {
				return err
			}

		default:
			msg := invalidColTypeMsg(table.Name(), cell.ColType)
			err = fmt.Errorf("visitCell() YAML: %s", msg)
			return err
		}
		// All errors in this switch are handled here.
		if err != nil {
			return err
		}

		yamlTableRow[cell.ColIndex] = anyVal

		return
	}

	const walkDeep = false
	err = table.Walk(walkDeep, visitTable, visitRow, visitCell)
	if err != nil {
		return
	}

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

	if yamlString == "" && yamlMap == nil {
		err = fmt.Errorf("%s: yamlString and yamlMap are both empty", UtilFuncName())
	}

	var tableSet *TableSet

	if yamlString != "" {
		tableSet, err = NewTableSetFromYAML(yamlString)
		if err != nil {
			err = fmt.Errorf("isValidYAML(): %v", err)
			return
		}

		var rowCount = 0
		for tableIndex := 0; tableIndex < tableSet.TableCount(); tableIndex++ {
			var table *Table
			table, err = tableSet.GetTableByTableIndex(tableIndex)
			if err != nil {
				err = fmt.Errorf("isValidYAML(): %v", err)
				return
			}

			rowCount += table.RowCount()
		}
		if rowCount == 0 {
			err = fmt.Errorf("%s: from yamlString: tableSet [[%s]] with %d tables has 0 rows",
				UtilFuncName(), tableSet.Name(), tableSet.TableCount())
			err = fmt.Errorf("isValidYAML(): %v", err)
			return
		}

		isValid = true
	}

	if yamlMap != nil {
		var yamlBytes []byte
		yamlBytes, err = yaml.Marshal(yamlMap)
		if err != nil {
			err = fmt.Errorf("isValidYAML(): %v", err)
			return
		}
		yamlString = string(yamlBytes)

		tableSet, err = NewTableSetFromYAML(yamlString)
		if err != nil {
			// See if it is a table.
			errFromTableSet := err
			_, err = newTableFromYAML_recursive(yamlMap)
			if err != nil {
				err = fmt.Errorf("isValidYAML(): %v (also %v)", err, errFromTableSet)
				return
			}
		}

		var rowCount = 0
		for tableIndex := 0; tableIndex < tableSet.TableCount(); tableIndex++ {
			var table *Table
			table, err = tableSet.GetTableByTableIndex(tableIndex)
			if err != nil {
				err = fmt.Errorf("isValidYAML(): %v", err)
				return
			}

			rowCount += table.RowCount()
		}
		if rowCount == 0 {
			err = fmt.Errorf("%s: from yamlString: tableSet [[%s]] with %d tables has 0 rows",
				UtilFuncName(), tableSet.Name(), tableSet.TableCount())
			err = fmt.Errorf("isValidYAML(): %v", err)
			return
		}

		isValid = true
	}

	return
}

func dataExists(yamlTable map[string]interface{}, sourceLineNumber int) string {
	_, exists := yamlTable["data"]
	return fmt.Sprintf(`ggg "data" exists = %t called by %s at %d`, exists, UtilFuncCaller(), sourceLineNumber)
}

func metadataExists(yamlTable map[string]interface{}, sourceLineNumber int) string {
	_, exists := yamlTable["metadata"].([]interface{})
	return fmt.Sprintf(`ggg "metadata" exists = %t called by %s at %d`, exists, UtilFuncCaller(), sourceLineNumber)
}

func tableNameExists(yamlTable map[string]interface{}, sourceLineNumber int) string {
	_, exists := yamlTable["tableName"].(interface{})
	return fmt.Sprintf(`ggg "tableName" exists = %t called by %s at %d`, exists, UtilFuncCaller(), sourceLineNumber)
}

func existsInMap(tableMap map[string]interface{}, searchKey string) (element interface{}, exists bool) {
	var key string
	for key, element = range tableMap {
		if key == searchKey {
			return element, true
		}
	}
	return nil, false
}
