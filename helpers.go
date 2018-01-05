package gotables

/*
	helpers.go
*/

import (
	"fmt"
	"os"
)

func (table *Table) model_AppendRowMap(newRow tableRow) error {
where(fmt.Sprintf("[%s].model_AppendRowMap()", table.Name()))
	// new memory model
	// Note: Simpler and probably more efficient to append a row at a time.
	// See: "Growing slices" at https://blog.golang.org/go-slices-usage-and-internals
	if table == nil { return fmt.Errorf("table.%s() table is <nil>", funcName()) }

	for colIndex, colName := range table.colNames {

		colType, err := table.ColType(colName)
		if err != nil { return err }

		switch colType {
			case "string":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([]string), val.(string))
			case "bool":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([]bool), val.(bool))
			case "int":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([]int), val.(int))
			case "int8":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([]int8), val.(int8))
			case "int16":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([]int16), val.(int16))
			case "int32":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([]int32), val.(int32))
			case "int64":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([]int64), val.(int64))
			case "uint":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([]uint), val.(uint))
			case "byte":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([]byte), val.(byte))
			case "uint8":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([]uint8), val.(uint8))
			case "uint16":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([]uint16), val.(uint16))
			case "uint32":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([]uint32), val.(uint32))
			case "uint64":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([]uint64), val.(uint64))
			case "float32":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([]float32), val.(float32))
			case "float64":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([]float64), val.(float64))
			case "[]byte":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([][]byte), val.([]byte))
			case "[]uint8":
				val, _ := newRow[colName]
				table.cols[colIndex] = append(table.cols[colIndex].([][]uint8), val.([]uint8))
			default:
				err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
				return err
		}
	}

	return nil
}

func (table *Table) model_rowsEqualRows() error {
	// new memory model

	// Loop through all the cols defined in the table.
	for colIndex, colName := range table.colNames {

		colType, err := table.ColType(colName)
		if err != nil { return err }

		rowCount := table.RowCount()

		switch colType {
			case "string":
				if len(table.cols[colIndex].([]string)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] string len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]string)), table.RowCount()))
				}
			case "bool":
				if len(table.cols[colIndex].([]bool)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] bool len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]bool)), table.RowCount()))
				}
			case "int":
				if len(table.cols[colIndex].([]int)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] int len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int)), table.RowCount()))
				}
			case "int8":
				if len(table.cols[colIndex].([]int8)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] int8 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int8)), table.RowCount()))
				}
			case "int16":
				if len(table.cols[colIndex].([]int16)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] int16 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int16)), table.RowCount()))
				}
			case "int32":
				if len(table.cols[colIndex].([]int32)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] int32 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int32)), table.RowCount()))
				}
			case "int64":
				if len(table.cols[colIndex].([]int64)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] int64 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int64)), table.RowCount()))
				}
			case "uint":
				if len(table.cols[colIndex].([]uint)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] uint len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint)), table.RowCount()))
				}
			case "byte":
				if len(table.cols[colIndex].([]byte)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] byte len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]byte)), table.RowCount()))
				}
			case "uint8":
				if len(table.cols[colIndex].([]uint8)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] uint8 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint8)), table.RowCount()))
				}
			case "uint16":
				if len(table.cols[colIndex].([]uint16)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] uint16 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint16)), table.RowCount()))
				}
			case "uint32":
				if len(table.cols[colIndex].([]uint32)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] uint32 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint32)), table.RowCount()))
				}
			case "uint64":
				if len(table.cols[colIndex].([]uint64)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] uint64 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint64)), table.RowCount()))
				}
			case "float32":
				if len(table.cols[colIndex].([]float32)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] float32 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]float32)), table.RowCount()))
				}
			case "float64":
				if len(table.cols[colIndex].([]float64)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] float64 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]float64)), table.RowCount()))
				}
			case "[]byte":
				if len(table.cols[colIndex].([][]byte)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] []byte len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([][]byte)), table.RowCount()))
				}
			case "[]uint8":
				if len(table.cols[colIndex].([][]uint8)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] []uint8 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([][]uint8)), table.RowCount()))
				}
			default:
				err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
				return err
		}
	}

	return nil
}

func model_newCol(colType string) (interface{}, error) {
	// new memory model
	var col interface{}

	switch colType {
		case "string":
			col = make([]string, 0)
		case "bool":
			col = make([]bool, 0)
		case "int":
			col = make([]int, 0)
		case "int8":
			col = make([]int8, 0)
		case "int16":
			col = make([]int16, 0)
		case "int32":
			col = make([]int32, 0)
		case "int64":
			col = make([]int64, 0)
		case "uint":
			col = make([]uint, 0)
		case "byte":
			col = make([]byte, 0)
		case "uint8":
			col = make([]uint8, 0)
		case "uint16":
			col = make([]uint16, 0)
		case "uint32":
			col = make([]uint32, 0)
		case "uint64":
			col = make([]uint64, 0)
		case "float32":
			col = make([]float32, 0)
		case "float64":
			col = make([]float64, 0)
		case "[]byte":
			col = make([][]byte, 0)
		case "[]uint8":
			col = make([][]uint8, 0)

		default:
			err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
			return nil, err
	}

	return col, nil
}

/*
	Append a new row to this table.
*/
func (table *Table) model_AppendRow() error {
where(fmt.Sprintf("[%s].model_AppendRow()", table.Name()))
	// new memory model
	// Note: Simpler and probably more efficient to append a row at a time.
	// See: "Growing slices" at https://blog.golang.org/go-slices-usage-and-internals

	// Note technique for appending a zero value to a slice.

	if table == nil { return fmt.Errorf("table.%s() table is <nil>", funcName()) }
where(fmt.Sprintf("table.model_ColCount() = %d", table.model_ColCount()))

	for colIndex, colName := range table.colNames {

		colType, err := table.ColType(colName)
		if err != nil { return err }

		switch colType {
			case "string":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]string))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new(string)
				table.cols[colIndex] = append(table.cols[colIndex].([]string), *val)
				rowCount = len(table.cols[colIndex].([]string))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "bool":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]bool))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new(bool)
				table.cols[colIndex] = append(table.cols[colIndex].([]bool), *val)
				rowCount = len(table.cols[colIndex].([]bool))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "int":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]int))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new(int)
				table.cols[colIndex] = append(table.cols[colIndex].([]int), *val)
				rowCount = len(table.cols[colIndex].([]int))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "int8":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]int8))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new(int8)
				table.cols[colIndex] = append(table.cols[colIndex].([]int8), *val)
				rowCount = len(table.cols[colIndex].([]int8))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "int16":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]int16))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new(int16)
				table.cols[colIndex] = append(table.cols[colIndex].([]int16), *val)
				rowCount = len(table.cols[colIndex].([]int16))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "int32":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]int32))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new(int32)
				table.cols[colIndex] = append(table.cols[colIndex].([]int32), *val)
				rowCount = len(table.cols[colIndex].([]int32))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "int64":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]int64))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new(int64)
				table.cols[colIndex] = append(table.cols[colIndex].([]int64), *val)
				rowCount = len(table.cols[colIndex].([]int64))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "uint":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]uint))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new(uint)
				table.cols[colIndex] = append(table.cols[colIndex].([]uint), *val)
				rowCount = len(table.cols[colIndex].([]uint))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "byte":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]byte))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new(byte)
				table.cols[colIndex] = append(table.cols[colIndex].([]byte), *val)
				rowCount = len(table.cols[colIndex].([]byte))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "uint8":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]uint8))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new(uint8)
				table.cols[colIndex] = append(table.cols[colIndex].([]uint8), *val)
				rowCount = len(table.cols[colIndex].([]uint8))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "uint16":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]uint16))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new(uint16)
				table.cols[colIndex] = append(table.cols[colIndex].([]uint16), *val)
				rowCount = len(table.cols[colIndex].([]uint16))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "uint32":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]uint32))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new(uint32)
				table.cols[colIndex] = append(table.cols[colIndex].([]uint32), *val)
				rowCount = len(table.cols[colIndex].([]uint32))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "uint64":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]uint64))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new(uint64)
				table.cols[colIndex] = append(table.cols[colIndex].([]uint64), *val)
				rowCount = len(table.cols[colIndex].([]uint64))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "float32":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]float32))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new(float32)
				table.cols[colIndex] = append(table.cols[colIndex].([]float32), *val)
				rowCount = len(table.cols[colIndex].([]float32))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "float64":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]float64))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new(float64)
				table.cols[colIndex] = append(table.cols[colIndex].([]float64), *val)
				rowCount = len(table.cols[colIndex].([]float64))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "[]byte":
				var rowCount int
				rowCount = len(table.cols[colIndex].([][]byte))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new([]byte)
				table.cols[colIndex] = append(table.cols[colIndex].([][]byte), *val)
				rowCount = len(table.cols[colIndex].([][]byte))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			case "[]uint8":
				var rowCount int
				rowCount = len(table.cols[colIndex].([][]uint8))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE rowCount = %d", rowCount))
				val := new([]uint8)
				table.cols[colIndex] = append(table.cols[colIndex].([][]uint8), *val)
				rowCount = len(table.cols[colIndex].([][]uint8))
where(fmt.Sprintf("AFTER  rowCount = %d", rowCount))
			default:
				err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
				return err
		}
	}

	return nil
}

/*
	Return the number of rows in this table.
*/
func (table *Table) model_RowCount() int {
	// new memory model

	if table == nil {
		os.Stderr.WriteString(fmt.Sprintf("ERROR: %s(): table is <nil>\n", funcName()))
		return -1
	}

	var rowCount int = -1
	var prevRowCount int = -1

	for colIndex, colName := range table.colNames {

		colType, err := table.ColType(colName)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("ERROR: %s(): %s\n", funcName(), err))
			return -1
		}

		switch colType {
			case "string":
//				val := new(string)
//				table.cols[colIndex] = append(table.cols[colIndex].([]string), *val)
				rowCount = len(table.cols[colIndex].([]string))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "bool":
//				val := new(bool)
//				table.cols[colIndex] = append(table.cols[colIndex].([]bool), *val)
				rowCount = len(table.cols[colIndex].([]bool))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "int":
//				val := new(int)
//				table.cols[colIndex] = append(table.cols[colIndex].([]int), *val)
				rowCount = len(table.cols[colIndex].([]int))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "int8":
//				val := new(int8)
//				table.cols[colIndex] = append(table.cols[colIndex].([]int8), *val)
				rowCount = len(table.cols[colIndex].([]int8))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "int16":
//				val := new(int16)
//				table.cols[colIndex] = append(table.cols[colIndex].([]int16), *val)
				rowCount = len(table.cols[colIndex].([]int16))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "int32":
//				val := new(int32)
//				table.cols[colIndex] = append(table.cols[colIndex].([]int32), *val)
				rowCount = len(table.cols[colIndex].([]int32))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "int64":
//				val := new(int64)
//				table.cols[colIndex] = append(table.cols[colIndex].([]int64), *val)
				rowCount = len(table.cols[colIndex].([]int64))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "uint":
//				val := new(uint)
//				table.cols[colIndex] = append(table.cols[colIndex].([]uint), *val)
				rowCount = len(table.cols[colIndex].([]uint))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "byte":
//				val := new(byte)
//				table.cols[colIndex] = append(table.cols[colIndex].([]byte), *val)
				rowCount = len(table.cols[colIndex].([]byte))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "uint8":
//				val := new(uint8)
//				table.cols[colIndex] = append(table.cols[colIndex].([]uint8), *val)
				rowCount = len(table.cols[colIndex].([]uint8))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "uint16":
//				val := new(uint16)
//				table.cols[colIndex] = append(table.cols[colIndex].([]uint16), *val)
				rowCount = len(table.cols[colIndex].([]uint16))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "uint32":
//				val := new(uint32)
//				table.cols[colIndex] = append(table.cols[colIndex].([]uint32), *val)
				rowCount = len(table.cols[colIndex].([]uint32))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "uint64":
//				val := new(uint64)
//				table.cols[colIndex] = append(table.cols[colIndex].([]uint64), *val)
				rowCount = len(table.cols[colIndex].([]uint64))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "float32":
//				val := new(float32)
//				table.cols[colIndex] = append(table.cols[colIndex].([]float32), *val)
				rowCount = len(table.cols[colIndex].([]float32))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "float64":
//				val := new(float64)
//				table.cols[colIndex] = append(table.cols[colIndex].([]float64), *val)
				rowCount = len(table.cols[colIndex].([]float64))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "[]byte":
//				val := new([]byte)
//				table.cols[colIndex] = append(table.cols[colIndex].([][]byte), *val)
				rowCount = len(table.cols[colIndex].([][]byte))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			case "[]uint8":
//				val := new([]uint8)
//				table.cols[colIndex] = append(table.cols[colIndex].([][]uint8), *val)
				rowCount = len(table.cols[colIndex].([][]uint8))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
			default:
				err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
				os.Stderr.WriteString(fmt.Sprintf("ERROR: %s(): %s\n", funcName(), err))
				return -1
		}
	}

	return rowCount
}
