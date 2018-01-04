package gotables

/*
	helpers.go
	DO NOT MODIFY
	Generated Thursday 4 Jan 2018 from template: ../gotables/helpers.template

	go run ../gotablesmain/helpersmain.go
*/

import (
	"fmt"
)

func (table *Table) model_AppendRowMap(newRow tableRow) error {
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

		switch colType {
			case "string":
				if len(table.cols[colIndex].([]string)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] string len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]string)), table.RowCount()))
				}
			case "bool":
				if len(table.cols[colIndex].([]bool)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] bool len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]bool)), table.RowCount()))
				}
			case "int":
				if len(table.cols[colIndex].([]int)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] int len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int)), table.RowCount()))
				}
			case "int8":
				if len(table.cols[colIndex].([]int8)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] int8 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int8)), table.RowCount()))
				}
			case "int16":
				if len(table.cols[colIndex].([]int16)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] int16 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int16)), table.RowCount()))
				}
			case "int32":
				if len(table.cols[colIndex].([]int32)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] int32 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int32)), table.RowCount()))
				}
			case "int64":
				if len(table.cols[colIndex].([]int64)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] int64 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int64)), table.RowCount()))
				}
			case "uint":
				if len(table.cols[colIndex].([]uint)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] uint len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint)), table.RowCount()))
				}
			case "byte":
				if len(table.cols[colIndex].([]byte)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] byte len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]byte)), table.RowCount()))
				}
			case "uint8":
				if len(table.cols[colIndex].([]uint8)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] uint8 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint8)), table.RowCount()))
				}
			case "uint16":
				if len(table.cols[colIndex].([]uint16)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] uint16 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint16)), table.RowCount()))
				}
			case "uint32":
				if len(table.cols[colIndex].([]uint32)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] uint32 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint32)), table.RowCount()))
				}
			case "uint64":
				if len(table.cols[colIndex].([]uint64)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] uint64 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint64)), table.RowCount()))
				}
			case "float32":
				if len(table.cols[colIndex].([]float32)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] float32 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]float32)), table.RowCount()))
				}
			case "float64":
				if len(table.cols[colIndex].([]float64)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] float64 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]float64)), table.RowCount()))
				}
			case "[]byte":
				if len(table.cols[colIndex].([][]byte)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] []byte len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([][]byte)), table.RowCount()))
				}
			case "[]uint8":
				if len(table.cols[colIndex].([][]uint8)) != table.RowCount() {
					panic(fmt.Sprintf("%s() table [%s] []uint8 len(table.cols[%d]) %d != table.RowCount() %d",
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
where(fmt.Sprintf("[]string len(col.([]string)) = %d", len(col.([]string))))
		case "bool":
			col = make([]bool, 0)
where(fmt.Sprintf("[]bool len(col.([]bool)) = %d", len(col.([]bool))))
		case "int":
			col = make([]int, 0)
where(fmt.Sprintf("[]int len(col.([]int)) = %d", len(col.([]int))))
		case "int8":
			col = make([]int8, 0)
where(fmt.Sprintf("[]int8 len(col.([]int8)) = %d", len(col.([]int8))))
		case "int16":
			col = make([]int16, 0)
where(fmt.Sprintf("[]int16 len(col.([]int16)) = %d", len(col.([]int16))))
		case "int32":
			col = make([]int32, 0)
where(fmt.Sprintf("[]int32 len(col.([]int32)) = %d", len(col.([]int32))))
		case "int64":
			col = make([]int64, 0)
where(fmt.Sprintf("[]int64 len(col.([]int64)) = %d", len(col.([]int64))))
		case "uint":
			col = make([]uint, 0)
where(fmt.Sprintf("[]uint len(col.([]uint)) = %d", len(col.([]uint))))
		case "byte":
			col = make([]byte, 0)
where(fmt.Sprintf("[]byte len(col.([]byte)) = %d", len(col.([]byte))))
		case "uint8":
			col = make([]uint8, 0)
where(fmt.Sprintf("[]uint8 len(col.([]uint8)) = %d", len(col.([]uint8))))
		case "uint16":
			col = make([]uint16, 0)
where(fmt.Sprintf("[]uint16 len(col.([]uint16)) = %d", len(col.([]uint16))))
		case "uint32":
			col = make([]uint32, 0)
where(fmt.Sprintf("[]uint32 len(col.([]uint32)) = %d", len(col.([]uint32))))
		case "uint64":
			col = make([]uint64, 0)
where(fmt.Sprintf("[]uint64 len(col.([]uint64)) = %d", len(col.([]uint64))))
		case "float32":
			col = make([]float32, 0)
where(fmt.Sprintf("[]float32 len(col.([]float32)) = %d", len(col.([]float32))))
		case "float64":
			col = make([]float64, 0)
where(fmt.Sprintf("[]float64 len(col.([]float64)) = %d", len(col.([]float64))))
		case "[]byte":
			col = make([][]byte, 0)
where(fmt.Sprintf("[][]byte len(col.([][]byte)) = %d", len(col.([][]byte))))
		case "[]uint8":
			col = make([][]uint8, 0)
where(fmt.Sprintf("[][]uint8 len(col.([][]uint8)) = %d", len(col.([][]uint8))))

		default:
			err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
			return nil, err
	}

	return col, nil
}
