package gotables

/*
	helpers.go
*/

import (
//	"runtime/debug"
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
col := table.cols[colIndex].([]string)
x := len(col)
y := len(table.cols[colIndex].([]string))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]string)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] string len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]string)), table.RowCount()))
				}
			case "bool":
col := table.cols[colIndex].([]bool)
x := len(col)
y := len(table.cols[colIndex].([]bool))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]bool)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] bool len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]bool)), table.RowCount()))
				}
			case "int":
col := table.cols[colIndex].([]int)
x := len(col)
y := len(table.cols[colIndex].([]int))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]int)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] int len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int)), table.RowCount()))
				}
			case "int8":
col := table.cols[colIndex].([]int8)
x := len(col)
y := len(table.cols[colIndex].([]int8))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]int8)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] int8 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int8)), table.RowCount()))
				}
			case "int16":
col := table.cols[colIndex].([]int16)
x := len(col)
y := len(table.cols[colIndex].([]int16))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]int16)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] int16 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int16)), table.RowCount()))
				}
			case "int32":
col := table.cols[colIndex].([]int32)
x := len(col)
y := len(table.cols[colIndex].([]int32))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]int32)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] int32 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int32)), table.RowCount()))
				}
			case "int64":
col := table.cols[colIndex].([]int64)
x := len(col)
y := len(table.cols[colIndex].([]int64))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]int64)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] int64 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int64)), table.RowCount()))
				}
			case "uint":
col := table.cols[colIndex].([]uint)
x := len(col)
y := len(table.cols[colIndex].([]uint))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]uint)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] uint len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint)), table.RowCount()))
				}
			case "byte":
col := table.cols[colIndex].([]byte)
x := len(col)
y := len(table.cols[colIndex].([]byte))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]byte)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] byte len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]byte)), table.RowCount()))
				}
			case "uint8":
col := table.cols[colIndex].([]uint8)
x := len(col)
y := len(table.cols[colIndex].([]uint8))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]uint8)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] uint8 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint8)), table.RowCount()))
				}
			case "uint16":
col := table.cols[colIndex].([]uint16)
x := len(col)
y := len(table.cols[colIndex].([]uint16))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]uint16)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] uint16 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint16)), table.RowCount()))
				}
			case "uint32":
col := table.cols[colIndex].([]uint32)
x := len(col)
y := len(table.cols[colIndex].([]uint32))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]uint32)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] uint32 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint32)), table.RowCount()))
				}
			case "uint64":
col := table.cols[colIndex].([]uint64)
x := len(col)
y := len(table.cols[colIndex].([]uint64))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]uint64)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] uint64 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint64)), table.RowCount()))
				}
			case "float32":
col := table.cols[colIndex].([]float32)
x := len(col)
y := len(table.cols[colIndex].([]float32))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]float32)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] float32 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]float32)), table.RowCount()))
				}
			case "float64":
col := table.cols[colIndex].([]float64)
x := len(col)
y := len(table.cols[colIndex].([]float64))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]float64)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] float64 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]float64)), table.RowCount()))
				}
			case "[]byte":
col := table.cols[colIndex].([][]byte)
x := len(col)
y := len(table.cols[colIndex].([][]byte))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([][]byte)) != rowCount {
					panic(fmt.Sprintf("*** %s() table [%s] []byte len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([][]byte)), table.RowCount()))
				}
			case "[]uint8":
col := table.cols[colIndex].([][]uint8)
x := len(col)
y := len(table.cols[colIndex].([][]uint8))
if x != y {
where(fmt.Sprintf("WARNING: x %d != y %d", x, y))
}
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

func (table *Table) model_AppendCol(colType string) error {
	// new memory model

	var col interface{}

	// Make new column the same rowCount size as (any) existing columns.
	var rowCount int = table.model_RowCount()
where(fmt.Sprintf("table.model_RowCount() = %d", table.model_RowCount()))

	switch colType {
		case "string":
			col = make([]string, rowCount)
		case "bool":
			col = make([]bool, rowCount)
		case "int":
			col = make([]int, rowCount)
		case "int8":
			col = make([]int8, rowCount)
		case "int16":
			col = make([]int16, rowCount)
		case "int32":
			col = make([]int32, rowCount)
		case "int64":
			col = make([]int64, rowCount)
		case "uint":
			col = make([]uint, rowCount)
		case "byte":
			col = make([]byte, rowCount)
		case "uint8":
			col = make([]uint8, rowCount)
		case "uint16":
			col = make([]uint16, rowCount)
		case "uint32":
			col = make([]uint32, rowCount)
		case "uint64":
			col = make([]uint64, rowCount)
		case "float32":
			col = make([]float32, rowCount)
		case "float64":
			col = make([]float64, rowCount)
		case "[]byte":
			col = make([][]byte, rowCount)
		case "[]uint8":
			col = make([][]uint8, rowCount)

		default:
			err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
			return err
	}

	table.cols = append(table.cols, col)

	return nil
}

/*
	Append a new row to this table.
*/
func (table *Table) model_AppendRow() error {
where(fmt.Sprintf("*** [%s].model_AppendRow()", table.Name()))
where(fmt.Sprintf("table.model_RowCount() = %d", table.model_RowCount()))
// debug.PrintStack()
	// new memory model
	// Note: Simpler and probably more efficient to append a row at a time.
	// See: "Growing slices" at https://blog.golang.org/go-slices-usage-and-internals

	// Note technique for appending a zero value to a slice without knowing the type.

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }
where(fmt.Sprintf("table.model_ColCount() = %d", table.model_ColCount()))
where(fmt.Sprintf("table.model_RowCount() = %d", table.model_RowCount()))

	if len(table.cols) == 0 {
		return fmt.Errorf("[%s].%s(): cannot append row to table with zero cols",
			table.Name(), funcName())
	}

where(fmt.Sprintf("table.model_RowCount() = %d", table.model_RowCount()))
	for colIndex, colName := range table.colNames {

		colType, err := table.ColType(colName)
		if err != nil { return err }

		switch colType {
			case "string":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]string))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new(string)
				table.cols[colIndex] = append(table.cols[colIndex].([]string), *val)
				rowCount = len(table.cols[colIndex].([]string))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "bool":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]bool))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new(bool)
				table.cols[colIndex] = append(table.cols[colIndex].([]bool), *val)
				rowCount = len(table.cols[colIndex].([]bool))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "int":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]int))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new(int)
				table.cols[colIndex] = append(table.cols[colIndex].([]int), *val)
				rowCount = len(table.cols[colIndex].([]int))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "int8":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]int8))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new(int8)
				table.cols[colIndex] = append(table.cols[colIndex].([]int8), *val)
				rowCount = len(table.cols[colIndex].([]int8))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "int16":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]int16))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new(int16)
				table.cols[colIndex] = append(table.cols[colIndex].([]int16), *val)
				rowCount = len(table.cols[colIndex].([]int16))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "int32":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]int32))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new(int32)
				table.cols[colIndex] = append(table.cols[colIndex].([]int32), *val)
				rowCount = len(table.cols[colIndex].([]int32))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "int64":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]int64))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new(int64)
				table.cols[colIndex] = append(table.cols[colIndex].([]int64), *val)
				rowCount = len(table.cols[colIndex].([]int64))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "uint":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]uint))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new(uint)
				table.cols[colIndex] = append(table.cols[colIndex].([]uint), *val)
				rowCount = len(table.cols[colIndex].([]uint))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "byte":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]byte))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new(byte)
				table.cols[colIndex] = append(table.cols[colIndex].([]byte), *val)
				rowCount = len(table.cols[colIndex].([]byte))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "uint8":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]uint8))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new(uint8)
				table.cols[colIndex] = append(table.cols[colIndex].([]uint8), *val)
				rowCount = len(table.cols[colIndex].([]uint8))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "uint16":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]uint16))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new(uint16)
				table.cols[colIndex] = append(table.cols[colIndex].([]uint16), *val)
				rowCount = len(table.cols[colIndex].([]uint16))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "uint32":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]uint32))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new(uint32)
				table.cols[colIndex] = append(table.cols[colIndex].([]uint32), *val)
				rowCount = len(table.cols[colIndex].([]uint32))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "uint64":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]uint64))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new(uint64)
				table.cols[colIndex] = append(table.cols[colIndex].([]uint64), *val)
				rowCount = len(table.cols[colIndex].([]uint64))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "float32":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]float32))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new(float32)
				table.cols[colIndex] = append(table.cols[colIndex].([]float32), *val)
				rowCount = len(table.cols[colIndex].([]float32))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "float64":
				var rowCount int
				rowCount = len(table.cols[colIndex].([]float64))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new(float64)
				table.cols[colIndex] = append(table.cols[colIndex].([]float64), *val)
				rowCount = len(table.cols[colIndex].([]float64))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "[]byte":
				var rowCount int
				rowCount = len(table.cols[colIndex].([][]byte))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new([]byte)
				table.cols[colIndex] = append(table.cols[colIndex].([][]byte), *val)
				rowCount = len(table.cols[colIndex].([][]byte))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			case "[]uint8":
				var rowCount int
				rowCount = len(table.cols[colIndex].([][]uint8))
where(fmt.Sprintf("[%s].%s", table.Name(), colName))
where(fmt.Sprintf("BEFORE %s() rowCount = %d", funcName(), rowCount))
				val := new([]uint8)
				table.cols[colIndex] = append(table.cols[colIndex].([][]uint8), *val)
				rowCount = len(table.cols[colIndex].([][]uint8))
where(fmt.Sprintf("AFTER  %s() rowCount = %d", funcName(), rowCount))
			default:
				err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
				return err
		}
	}
where(fmt.Sprintf("table.model_RowCount() = %d", table.model_RowCount()))

	return nil
}

/*
	Return the number of rows in this table.
*/
func (table *Table) model_RowCount() (rowCount int) {
	// new memory model

	if table == nil {
		_,_ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: %s(): table is <nil>\n", funcSource(), funcName()))
where("model_RowCount()")
		return -1
	}

	if table.cols == nil {
		_,_ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: %s(): [%s].cols = nil\n", funcSource(), table.Name(), funcName()))
where("model_RowCount()")
		return -1
	}

	if len(table.cols) == 0 {
		// Avoid index out of range indexing into table.cols in switch statement.
		// This implies rows cannot be added before at least one column is present.
where("model_RowCount() return 0")
		return 0
	}
// where("model_RowCount()")
// where(fmt.Sprintf("len(table.cols) = %d", len(table.cols)))

	rowCount = -1
	var prevRowCount int = -1

// where(fmt.Sprintf("len(table.colNames) = %d", len(table.colNames)))
	for colIndex := 0; colIndex < len(table.cols); colIndex++ {
// where(fmt.Sprintf("colIndex = %d", colIndex))

		colType := table.colTypes[colIndex]

		switch colType {
			case "string":
				rowCount = len(table.cols[colIndex].([]string))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "bool":
				rowCount = len(table.cols[colIndex].([]bool))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "int":
				rowCount = len(table.cols[colIndex].([]int))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "int8":
				rowCount = len(table.cols[colIndex].([]int8))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "int16":
				rowCount = len(table.cols[colIndex].([]int16))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "int32":
				rowCount = len(table.cols[colIndex].([]int32))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "int64":
				rowCount = len(table.cols[colIndex].([]int64))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "uint":
				rowCount = len(table.cols[colIndex].([]uint))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "byte":
				rowCount = len(table.cols[colIndex].([]byte))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "uint8":
				rowCount = len(table.cols[colIndex].([]uint8))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "uint16":
				rowCount = len(table.cols[colIndex].([]uint16))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "uint32":
				rowCount = len(table.cols[colIndex].([]uint32))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "uint64":
				rowCount = len(table.cols[colIndex].([]uint64))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "float32":
				rowCount = len(table.cols[colIndex].([]float32))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "float64":
				rowCount = len(table.cols[colIndex].([]float64))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "[]byte":
				rowCount = len(table.cols[colIndex].([][]byte))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			case "[]uint8":
				rowCount = len(table.cols[colIndex].([][]uint8))
				if prevRowCount > -1 && rowCount != prevRowCount {
					panic(fmt.Errorf("col %s (prevRowCount) %d != col %s rowCount %d",
						table.colNames[colIndex-1], prevRowCount, table.colNames[colIndex], rowCount))
				}
				prevRowCount = rowCount
			default:
				_,_ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR IN %s(): unknown type: %s\n", funcSource(), funcName(), colType))
				return -1
		}
	}

	return rowCount
}

/*
	Delete a row from this table.
*/
func (table *Table) model_DeleteRow(rowIndex int) error {
	// new memory model

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	if len(table.cols) == 0 {
		return fmt.Errorf("[%s].%s(): cannot delete rows from table with zero cols",
			table.Name(), funcName())
	}

	if rowIndex < 0 || rowIndex > table.model_RowCount()-1 {
		return fmt.Errorf("%s(): in table [%s] with %d rows, row index %d does not exist",
			funcName(), table.tableName, table.model_RowCount(), rowIndex)
	}
where(fmt.Sprintf("BEFORE deleting row %d table.model_RowCount() = %d", rowIndex, table.model_RowCount()))

	for colIndex, colName := range table.colNames {

		colType, err := table.ColType(colName)
		if err != nil { return err }
// where(colIndex)
// where(colName)
// where(colType)
// where(fmt.Sprintf("[%s] %d %s %s []col type = %T", table.Name(), colIndex, colName, colType, table.cols[colIndex]))

		switch colType {
			case "string":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([]string)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "bool":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([]bool)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "int":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([]int)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "int8":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([]int8)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "int16":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([]int16)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "int32":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([]int32)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "int64":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([]int64)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "uint":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([]uint)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "byte":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([]byte)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "uint8":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([]uint8)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "uint16":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([]uint16)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "uint32":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([]uint32)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "uint64":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([]uint64)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "float32":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([]float32)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "float64":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([]float64)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "[]byte":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([][]byte)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			case "[]uint8":
where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
				col := table.cols[colIndex].([][]uint8)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
				// From Ivo Balbaert p182 for deleting a single element from a slice.
				col = append(col[:rowIndex], col[rowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
// where()
			default:
				err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
				return err
		}
	}
where(fmt.Sprintf("AFTER  deleting row %d table.model_RowCount() = %d", rowIndex, table.model_RowCount()))

	return nil
}

/*
	Delete a col from this table.
*/
func (table *Table) model_DeleteColByColIndex(colIndex int) error {
	// new memory model

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	if len(table.cols) == 0 {
		return fmt.Errorf("[%s].%s(): cannot delete cols from table with zero cols",
			table.Name(), funcName())
	}

	if colIndex < 0 || colIndex > table.model_ColCount()-1 {
		err := fmt.Errorf("in table [%s] with %d cols, col index %d does not exist",
			table.tableName, table.model_ColCount(), colIndex)
		return err
	}

	// new memory model
	// From Ivo Balbaert p182 for deleting a single element from a slice.
	table.cols = append(table.cols[:colIndex], table.cols[colIndex+1:]...)

	return nil
}

/*
	Delete rows from this table.

	Delete rows from firstRowIndex to lastRowIndex inclusive.
	This means lastRowIndex will be deleted.
*/
func (table *Table) model_DeleteRows(firstRowIndex int, lastRowIndex int) error {
	// new memory model

	if table == nil { return fmt.Errorf("(model) table.%s(): table is <nil>", funcName()) }

	if len(table.cols) == 0 {
		return fmt.Errorf("(model) [%s].%s(): cannot delete rows from table with zero cols",
			table.Name(), funcName())
	}

	if firstRowIndex < 0 || firstRowIndex > table.model_RowCount()-1 {
		return fmt.Errorf("(model) in table [%s] with %d rows, firstRowIndex %d does not exist",
			table.tableName, table.model_RowCount(), firstRowIndex)
	}

    if lastRowIndex < 0 || lastRowIndex > table.model_RowCount()-1 {
        return fmt.Errorf("(model) in table [%s] with %d rows, lastRowIndex %d does not exist",
            table.tableName, table.model_RowCount(), lastRowIndex)
    }

    if firstRowIndex > lastRowIndex {
        return fmt.Errorf("(model) invalid row index range: firstRowIndex %d > lastRowIndex %d", firstRowIndex, lastRowIndex)
    }
where(fmt.Sprintf("BEFORE model_DeleteRows() model_RowCount() = %d", table.model_RowCount()))

	for colIndex, colName := range table.colNames {

		colType, err := table.ColType(colName)
		if err != nil { return err }
// where(colIndex)
// where(colName)
// where(colType)
// where(fmt.Sprintf("[%s] %d %s %s []col type = %T", table.Name(), colIndex, colName, colType, table.cols[colIndex]))

		switch colType {
			case "string":
				col := table.cols[colIndex].([]string)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "bool":
				col := table.cols[colIndex].([]bool)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "int":
				col := table.cols[colIndex].([]int)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "int8":
				col := table.cols[colIndex].([]int8)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "int16":
				col := table.cols[colIndex].([]int16)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "int32":
				col := table.cols[colIndex].([]int32)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "int64":
				col := table.cols[colIndex].([]int64)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "uint":
				col := table.cols[colIndex].([]uint)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "byte":
				col := table.cols[colIndex].([]byte)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "uint8":
				col := table.cols[colIndex].([]uint8)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "uint16":
				col := table.cols[colIndex].([]uint16)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "uint32":
				col := table.cols[colIndex].([]uint32)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "uint64":
				col := table.cols[colIndex].([]uint64)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "float32":
				col := table.cols[colIndex].([]float32)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "float64":
				col := table.cols[colIndex].([]float64)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "[]byte":
				col := table.cols[colIndex].([][]byte)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "[]uint8":
				col := table.cols[colIndex].([][]uint8)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			default:
				err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
				return err
		}
	}
where(fmt.Sprintf("AFTER  model_DeleteRows() model_RowCount() = %d", table.model_RowCount()))

	return nil
}
