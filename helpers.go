package gotables

/*
	helpers.go
*/

import (
	"bytes"
	"fmt"
	"os"
//	"runtime/debug"
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

func (table *Table) new_model_appendRowMap(rowMap tableRow) error {
where()
where(fmt.Sprintf("[%s].new_model_AppendRowMap(%v)", table.Name(), rowMap))
	// new memory model
	// Note: Simpler and probably more efficient to append a row at a time.
	// See: "Growing slices" at https://blog.golang.org/go-slices-usage-and-internals
	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

//	var err error

where(fmt.Sprintf("BEFORE %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("BEFORE %s(): [%s].new_model_RowCount() = %d", funcName(), table.Name(), table.new_model_RowCount()))

	for colIndex, colName := range table.colNames {

//		colType, err := table.ColType(colName)
//		if err != nil { return err }
		var colType string = table.colTypes[colIndex]
// where(fmt.Sprintf("%s(): BEFORE append colIndex = %d [%s].new_model_RowCount() = %d",
// funcName(), colIndex, table.Name(), table.new_model_RowCount()))

		switch colType {
			case "string":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col []string = make([]string, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col []string = table.cols[colIndex].([]string)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.(string))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "bool":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col []bool = make([]bool, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col []bool = table.cols[colIndex].([]bool)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.(bool))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "int":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col []int = make([]int, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col []int = table.cols[colIndex].([]int)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.(int))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "int8":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col []int8 = make([]int8, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col []int8 = table.cols[colIndex].([]int8)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.(int8))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "int16":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col []int16 = make([]int16, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col []int16 = table.cols[colIndex].([]int16)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.(int16))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "int32":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col []int32 = make([]int32, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col []int32 = table.cols[colIndex].([]int32)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.(int32))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "int64":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col []int64 = make([]int64, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col []int64 = table.cols[colIndex].([]int64)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.(int64))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "uint":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col []uint = make([]uint, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col []uint = table.cols[colIndex].([]uint)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.(uint))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "byte":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col []byte = make([]byte, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col []byte = table.cols[colIndex].([]byte)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.(byte))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "uint8":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col []uint8 = make([]uint8, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col []uint8 = table.cols[colIndex].([]uint8)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.(uint8))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "uint16":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col []uint16 = make([]uint16, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col []uint16 = table.cols[colIndex].([]uint16)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.(uint16))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "uint32":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col []uint32 = make([]uint32, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col []uint32 = table.cols[colIndex].([]uint32)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.(uint32))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "uint64":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col []uint64 = make([]uint64, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col []uint64 = table.cols[colIndex].([]uint64)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.(uint64))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "float32":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col []float32 = make([]float32, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col []float32 = table.cols[colIndex].([]float32)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.(float32))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "float64":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col []float64 = make([]float64, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col []float64 = table.cols[colIndex].([]float64)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.(float64))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "[]byte":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col [][]byte = make([][]byte, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col [][]byte = table.cols[colIndex].([][]byte)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.([]byte))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			case "[]uint8":
				colCount := table.new_model_ColCount()
where(fmt.Sprintf("MIDDLE 1 %s(): table.cols = %v", funcName(), table.cols))
				if colCount < colIndex+1 {
where(fmt.Sprintf("MIDDLE 2 %s(): table.cols = %v", funcName(), table.cols))
					// Column doesn't exist. Create and append it.
// where(fmt.Sprintf("%s(): colCount %d < colIndex %d + 1 = %d", funcName(), colCount, colIndex, colIndex + 1))
where(fmt.Sprintf("append col: %s", colName))
					// This is the first cell of this new column. Make just one row.
					// Don't call table.new_model_AppendCol() which will attempt to make zero rows.
//					if colCount >= 1 && table.new_model_RowCount() > 1 {
//						// Something has gone seriously wrong. Not expecting existing values for this column.
//						err = fmt.Errorf("%s(): colCount %d >= 1 && table.new_model_RowCount() %d > 1 something has gone seriously wrong",
//							funcName(), colCount, table.new_model_RowCount())
//						return err
//					}
					var col [][]uint8 = make([][]uint8, 0)	// ???
where(fmt.Sprintf("MIDDLE 3 %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("MIDDLE 3a %s(): col = %v len(col) = %d", funcName(), col, len(col)))
where(fmt.Sprintf("AAA table.cols = append(table.cols, %v)", col))
					table.cols = append(table.cols, col)
where(fmt.Sprintf("MIDDLE 4 %s(): table.cols = %v", funcName(), table.cols))
//					err = table.new_model_AppendCol(colName, colType)
//					if err != nil { return err }
				}
				val, _ := rowMap[colName]
				var col [][]uint8 = table.cols[colIndex].([][]uint8)
where(fmt.Sprintf("MIDDLE 5 %s(): table.cols = %v", funcName(), table.cols))
				col = append(col, val.([]uint8))
where(fmt.Sprintf("MIDDLE 6 %s(): table.cols = %v", funcName(), table.cols))
				table.cols[colIndex] = col
where(fmt.Sprintf("MIDDLE 7 %s(): table.cols = %v", funcName(), table.cols))
// where(fmt.Sprintf("val =  %v", val))
			default:
				err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
				return err
		}
where(fmt.Sprintf("MIDDLE 8 %s(): table.cols = %v", funcName(), table.cols))
	}
// where(fmt.Sprintf("%s(): [%s].new_model_ColCount() = %d", funcName(), table.Name(), table.new_model_ColCount()))
// where(fmt.Sprintf("%s(): AFTER  append colIndex = %d [%s].new_model_RowCount() = %d",
// funcName(), colIndex, table.Name(), table.new_model_RowCount()))
where(fmt.Sprintf("AFTER  %s(): table.cols = %v", funcName(), table.cols))
where(fmt.Sprintf("AFTER  %s(): [%s].new_model_RowCount() = %d", funcName(), table.Name(), table.new_model_RowCount()))

	return nil
}

func (table *Table) new_model_rowsEqualRows() error {
	// new memory model

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	old_model_rowCount := table.RowCount()

	// Loop through all the cols defined in the table.
	for colIndex, colName := range table.colNames {

		colType, err := table.ColType(colName)
		if err != nil { return err }

		switch colType {
			case "string":
col := table.cols[colIndex].([]string)
x := len(col)
y := len(table.cols[colIndex].([]string))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]string)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] string len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]string)), table.RowCount())
				}
			case "bool":
col := table.cols[colIndex].([]bool)
x := len(col)
y := len(table.cols[colIndex].([]bool))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]bool)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] bool len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]bool)), table.RowCount())
				}
			case "int":
col := table.cols[colIndex].([]int)
x := len(col)
y := len(table.cols[colIndex].([]int))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]int)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] int len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int)), table.RowCount())
				}
			case "int8":
col := table.cols[colIndex].([]int8)
x := len(col)
y := len(table.cols[colIndex].([]int8))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]int8)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] int8 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int8)), table.RowCount())
				}
			case "int16":
col := table.cols[colIndex].([]int16)
x := len(col)
y := len(table.cols[colIndex].([]int16))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]int16)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] int16 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int16)), table.RowCount())
				}
			case "int32":
col := table.cols[colIndex].([]int32)
x := len(col)
y := len(table.cols[colIndex].([]int32))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]int32)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] int32 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int32)), table.RowCount())
				}
			case "int64":
col := table.cols[colIndex].([]int64)
x := len(col)
y := len(table.cols[colIndex].([]int64))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]int64)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] int64 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]int64)), table.RowCount())
				}
			case "uint":
col := table.cols[colIndex].([]uint)
x := len(col)
y := len(table.cols[colIndex].([]uint))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]uint)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] uint len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint)), table.RowCount())
				}
			case "byte":
col := table.cols[colIndex].([]byte)
x := len(col)
y := len(table.cols[colIndex].([]byte))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]byte)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] byte len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]byte)), table.RowCount())
				}
			case "uint8":
col := table.cols[colIndex].([]uint8)
x := len(col)
y := len(table.cols[colIndex].([]uint8))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]uint8)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] uint8 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint8)), table.RowCount())
				}
			case "uint16":
col := table.cols[colIndex].([]uint16)
x := len(col)
y := len(table.cols[colIndex].([]uint16))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]uint16)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] uint16 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint16)), table.RowCount())
				}
			case "uint32":
col := table.cols[colIndex].([]uint32)
x := len(col)
y := len(table.cols[colIndex].([]uint32))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]uint32)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] uint32 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint32)), table.RowCount())
				}
			case "uint64":
col := table.cols[colIndex].([]uint64)
x := len(col)
y := len(table.cols[colIndex].([]uint64))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]uint64)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] uint64 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]uint64)), table.RowCount())
				}
			case "float32":
col := table.cols[colIndex].([]float32)
x := len(col)
y := len(table.cols[colIndex].([]float32))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]float32)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] float32 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]float32)), table.RowCount())
				}
			case "float64":
col := table.cols[colIndex].([]float64)
x := len(col)
y := len(table.cols[colIndex].([]float64))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([]float64)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] float64 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([]float64)), table.RowCount())
				}
			case "[]byte":
col := table.cols[colIndex].([][]byte)
x := len(col)
y := len(table.cols[colIndex].([][]byte))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([][]byte)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] []byte len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([][]byte)), table.RowCount())
				}
			case "[]uint8":
col := table.cols[colIndex].([][]uint8)
x := len(col)
y := len(table.cols[colIndex].([][]uint8))
if x != y {
where(fmt.Sprintf("NNN WARNING: x %d != y %d", x, y))
}
				if len(table.cols[colIndex].([][]uint8)) != old_model_rowCount {
					return fmt.Errorf("NNN %s() table [%s] []uint8 len(table.cols[%d]) %d != table.RowCount() %d",
						funcName(), table.Name(), colIndex, len(table.cols[colIndex].([][]uint8)), table.RowCount())
				}
			default:
				err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
				return err
		}
	}

	return nil
}

func (table *Table) new_model_AppendCol(colName string, colType string) error {
	// new memory model
where(fmt.Sprintf("FFF %s(colName=%s, colType=%s)", funcName(), colName, colType))
// if debugging { debug.PrintStack() }

	var err error
	var col interface{}

	// Make new column the same rowCount size as (any) existing columns.
	var rowCount int = table.new_model_RowCount()
where(fmt.Sprintf("FFF BEFORE append() in %s(): [%s].new_model_RowCount(%s, %s) = %d",
	funcName(), table.Name(), colName, colType, rowCount))
where(fmt.Sprintf("FFF ZZZ BEFORE table.cols = %v", table.cols))
where(fmt.Sprintf("FFF ZZZ BEFORE len(table.cols) = %d", len(table.cols)))

	switch colType {
		case "string":
where(fmt.Sprintf("%s(): col = make([]string, rowCount=%d)", funcName(), rowCount))
			col = make([]string, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([]string))))
		case "bool":
where(fmt.Sprintf("%s(): col = make([]bool, rowCount=%d)", funcName(), rowCount))
			col = make([]bool, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([]bool))))
		case "int":
where(fmt.Sprintf("%s(): col = make([]int, rowCount=%d)", funcName(), rowCount))
			col = make([]int, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([]int))))
		case "int8":
where(fmt.Sprintf("%s(): col = make([]int8, rowCount=%d)", funcName(), rowCount))
			col = make([]int8, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([]int8))))
		case "int16":
where(fmt.Sprintf("%s(): col = make([]int16, rowCount=%d)", funcName(), rowCount))
			col = make([]int16, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([]int16))))
		case "int32":
where(fmt.Sprintf("%s(): col = make([]int32, rowCount=%d)", funcName(), rowCount))
			col = make([]int32, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([]int32))))
		case "int64":
where(fmt.Sprintf("%s(): col = make([]int64, rowCount=%d)", funcName(), rowCount))
			col = make([]int64, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([]int64))))
		case "uint":
where(fmt.Sprintf("%s(): col = make([]uint, rowCount=%d)", funcName(), rowCount))
			col = make([]uint, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([]uint))))
		case "byte":
where(fmt.Sprintf("%s(): col = make([]byte, rowCount=%d)", funcName(), rowCount))
			col = make([]byte, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([]byte))))
		case "uint8":
where(fmt.Sprintf("%s(): col = make([]uint8, rowCount=%d)", funcName(), rowCount))
			col = make([]uint8, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([]uint8))))
		case "uint16":
where(fmt.Sprintf("%s(): col = make([]uint16, rowCount=%d)", funcName(), rowCount))
			col = make([]uint16, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([]uint16))))
		case "uint32":
where(fmt.Sprintf("%s(): col = make([]uint32, rowCount=%d)", funcName(), rowCount))
			col = make([]uint32, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([]uint32))))
		case "uint64":
where(fmt.Sprintf("%s(): col = make([]uint64, rowCount=%d)", funcName(), rowCount))
			col = make([]uint64, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([]uint64))))
		case "float32":
where(fmt.Sprintf("%s(): col = make([]float32, rowCount=%d)", funcName(), rowCount))
			col = make([]float32, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([]float32))))
		case "float64":
where(fmt.Sprintf("%s(): col = make([]float64, rowCount=%d)", funcName(), rowCount))
			col = make([]float64, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([]float64))))
		case "[]byte":
where(fmt.Sprintf("%s(): col = make([][]byte, rowCount=%d)", funcName(), rowCount))
			col = make([][]byte, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([][]byte))))
		case "[]uint8":
where(fmt.Sprintf("%s(): col = make([][]uint8, rowCount=%d)", funcName(), rowCount))
			col = make([][]uint8, rowCount)
where(fmt.Sprintf("BBB table.cols = append(table.cols, %v)", col))
			table.cols = append(table.cols, col)
where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len((col).([][]uint8))))

		default:
			err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
			return err
	}
	_, err = table.IsValidTable()
	if err != nil { return err }
where(fmt.Sprintf("AFTER  append() in %s(): [%s].new_model_RowCount(%s, %s) = %d",
funcName(), table.Name(), colName, colType, table.new_model_RowCount()))
where(fmt.Sprintf("FFF ZZZ AFTER  table.cols = %v", table.cols))
where(fmt.Sprintf("FFF ZZZ AFTER  len(table.cols) = %d", len(table.cols)))

/*
where("WHAT?")
where(fmt.Sprintf("CCC table.cols = append(table.cols, %v)", col))
	table.cols = append(table.cols, col)
*/

	return nil
}

//Append a new row to this table.
func (table *Table) new_model_AppendRow() error {
where(fmt.Sprintf("AAA [%s].new_model_AppendRow()", table.Name()))
where(fmt.Sprintf("table.new_model_RowCount() = %d", table.new_model_RowCount()))
// debug.PrintStack()
	// new memory model
	// Note: Simpler and probably more efficient to append a row at a time.
	// See: "Growing slices" at https://blog.golang.org/go-slices-usage-and-internals

	// Note technique for appending a zero value to a slice without knowing the type.

	/*
		new_model_AppendRow replaces need for appendRowOfNil() and SetRowCellsToZeroValue()
		because Go initialises elements to zero value.
	*/

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }
where(fmt.Sprintf("table.new_model_ColCount() = %d", table.new_model_ColCount()))
where(fmt.Sprintf("table.new_model_RowCount() = %d", table.new_model_RowCount()))

	_, err := table.IsValidTable()
	if err != nil { return err }

	if len(table.cols) == 0 {
//debug.PrintStack()
		return fmt.Errorf("[%s].%s(): cannot append row to table with zero cols",
			table.Name(), funcName())
	}

	var rowCount int

where(fmt.Sprintf("table.new_model_RowCount() = %d", table.new_model_RowCount()))
	for colIndex, colName := range table.colNames {

		colType, err := table.ColType(colName)
		if err != nil { return err }

where(fmt.Sprintf("[%s].%s", table.Name(), colName))
		switch colType {
			case "string":
				rowCount = len(table.cols[colIndex].([]string))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new(string)	// Pointer to zero value string
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([]string), val.(string)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([]string), *val)
//				rowCount = len(table.cols[colIndex].([]string))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "bool":
				rowCount = len(table.cols[colIndex].([]bool))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new(bool)	// Pointer to zero value bool
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([]bool), val.(bool)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([]bool), *val)
//				rowCount = len(table.cols[colIndex].([]bool))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "int":
				rowCount = len(table.cols[colIndex].([]int))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new(int)	// Pointer to zero value int
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([]int), val.(int)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([]int), *val)
//				rowCount = len(table.cols[colIndex].([]int))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "int8":
				rowCount = len(table.cols[colIndex].([]int8))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new(int8)	// Pointer to zero value int8
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([]int8), val.(int8)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([]int8), *val)
//				rowCount = len(table.cols[colIndex].([]int8))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "int16":
				rowCount = len(table.cols[colIndex].([]int16))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new(int16)	// Pointer to zero value int16
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([]int16), val.(int16)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([]int16), *val)
//				rowCount = len(table.cols[colIndex].([]int16))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "int32":
				rowCount = len(table.cols[colIndex].([]int32))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new(int32)	// Pointer to zero value int32
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([]int32), val.(int32)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([]int32), *val)
//				rowCount = len(table.cols[colIndex].([]int32))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "int64":
				rowCount = len(table.cols[colIndex].([]int64))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new(int64)	// Pointer to zero value int64
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([]int64), val.(int64)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([]int64), *val)
//				rowCount = len(table.cols[colIndex].([]int64))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "uint":
				rowCount = len(table.cols[colIndex].([]uint))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new(uint)	// Pointer to zero value uint
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([]uint), val.(uint)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([]uint), *val)
//				rowCount = len(table.cols[colIndex].([]uint))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "byte":
				rowCount = len(table.cols[colIndex].([]byte))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new(byte)	// Pointer to zero value byte
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([]byte), val.(byte)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([]byte), *val)
//				rowCount = len(table.cols[colIndex].([]byte))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "uint8":
				rowCount = len(table.cols[colIndex].([]uint8))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new(uint8)	// Pointer to zero value uint8
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([]uint8), val.(uint8)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([]uint8), *val)
//				rowCount = len(table.cols[colIndex].([]uint8))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "uint16":
				rowCount = len(table.cols[colIndex].([]uint16))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new(uint16)	// Pointer to zero value uint16
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([]uint16), val.(uint16)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([]uint16), *val)
//				rowCount = len(table.cols[colIndex].([]uint16))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "uint32":
				rowCount = len(table.cols[colIndex].([]uint32))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new(uint32)	// Pointer to zero value uint32
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([]uint32), val.(uint32)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([]uint32), *val)
//				rowCount = len(table.cols[colIndex].([]uint32))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "uint64":
				rowCount = len(table.cols[colIndex].([]uint64))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new(uint64)	// Pointer to zero value uint64
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([]uint64), val.(uint64)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([]uint64), *val)
//				rowCount = len(table.cols[colIndex].([]uint64))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "float32":
				rowCount = len(table.cols[colIndex].([]float32))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new(float32)	// Pointer to zero value float32
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([]float32), val.(float32)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([]float32), *val)
//				rowCount = len(table.cols[colIndex].([]float32))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "float64":
				rowCount = len(table.cols[colIndex].([]float64))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new(float64)	// Pointer to zero value float64
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([]float64), val.(float64)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([]float64), *val)
//				rowCount = len(table.cols[colIndex].([]float64))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "[]byte":
				rowCount = len(table.cols[colIndex].([][]byte))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new([]byte)	// Pointer to zero value []byte
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([][]byte), val.([]byte)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([][]byte), *val)
//				rowCount = len(table.cols[colIndex].([][]byte))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			case "[]uint8":
				rowCount = len(table.cols[colIndex].([][]uint8))
where(fmt.Sprintf("BEFORE append new row %s() rowCount = %d", funcName(), rowCount))
				val := new([]uint8)	// Pointer to zero value []uint8
where(fmt.Sprintf("%s(): table.cols[%d] = append(table.cols[%d].([][]uint8), val.([]uint8)) %v", funcName(), colIndex, colIndex, *val))
				table.cols[colIndex] = append(table.cols[colIndex].([][]uint8), *val)
//				rowCount = len(table.cols[colIndex].([][]uint8))
where(fmt.Sprintf("AFTER  append new value %v in %s() rowCount = %d", val, funcName(), rowCount))
			default:
				err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
				return err
		}
	}
where(fmt.Sprintf("table.new_model_RowCount() = %d", table.new_model_RowCount()))

	_, err = table.IsValidTable()
	if err != nil { return err }
	table.rowsIndex = append(table.rowsIndex, rowCount-1)
	_, err = table.IsValidTable()
	if err != nil { return err }

	return nil
}

/*
	Return the number of rows in this table.
*/
func (table *Table) new_model_checkRowCount() (rowCount int) {
where(fmt.Sprintf("inside [%s].%s()", table.Name(), funcName()))
// debug.PrintStack()

	if new_model {
		if table == nil {
			_,_ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: %s(): table is <nil>\n", funcSource(), funcName()))
			return -1
		}

		if table.cols == nil {
			_,_ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: %s(): [%s].cols = nil\n", funcSource(), table.Name(), funcName()))
			return -1
		}
	}

	// Temporary check? This should never happen in production.
	if new_model {
	    if len(table.cols) > len(table.colNames) {
	        err := fmt.Errorf("%s ERROR: %s(): len([%s].cols) %d > len([%s].colNames) %d\n",
	            funcSource(),
	            funcName(),
	            table.Name(),
	            len(table.cols),
	            table.Name(),
	            len(table.colNames),
	        )
	        _,_ = os.Stderr.WriteString(err.Error())
	        panic(err)
	    }
	}

	if new_model {
		// Skip these checks. Old model will have already created table.colNames and table.colTypes

		if len(table.colNames) != len(table.cols) {
			err := fmt.Errorf("%s ERROR: %s(): len([%s].colNames) %d != len([%s].cols) %d\n",
				funcSource(),
				funcName(),
				table.Name(),
				len(table.colNames),
				table.Name(),
				len(table.cols),
			)
			_,_ = os.Stderr.WriteString(err.Error())
			panic(err)
		}

		if len(table.colTypes) != len(table.cols) {
			err := fmt.Errorf("%s ERROR: %s(): len([%s].colTypes) %d != len([%s].cols) %d\n",
				funcSource(),
				funcName(),
				table.Name(),
				len(table.colTypes),
				table.Name(),
				len(table.cols),
			)
			_,_ = os.Stderr.WriteString(err.Error())
			panic(err)
		}
	}

	if len(table.colTypes) != len(table.colNames) {
		err := fmt.Errorf("%s ERROR: %s(): len([%s].colTypes) %d != len([%s].colNames) %d\n",
			funcSource(),
			funcName(),
			table.Name(),
			len(table.colTypes),
			table.Name(),
			len(table.colNames),
		)
		_,_ = os.Stderr.WriteString(err.Error())
		panic(err)
	}

	if len(table.cols) == 0 {
		// Avoid index out of range indexing into table.cols in switch statement.
		// This implies rows cannot be appended before at least one column has been appended.
//debug.PrintStack()
		rowCount = 0
where(fmt.Sprintf("MIDDLE OF FUNC [%s].%s() = %d", table.Name(), funcName(), rowCount))
		return
	}
// where(fmt.Sprintf("len(table.cols) = %d", len(table.cols)))

	var modelRowCount int = -1
	var prevModelRowCount int = -1

// debug.PrintStack()
where(fmt.Sprintf("len(table.colNames) = %d", len(table.colNames)))
where(fmt.Sprintf("len(table.cols) = %d", len(table.cols)))

	for colIndex := 0; colIndex < len(table.cols); colIndex++ {

where(fmt.Sprintf("YYY [%s] colNames = %v", table.Name(), table.colNames))
where(fmt.Sprintf("YYY [%s] colTypes = %v", table.Name(), table.colTypes))
where(fmt.Sprintf("YYY [%s] table.cols = %v", table.Name(), table.cols))
where(fmt.Sprintf("YYY [%s] len(table.cols) = %d", table.Name(), len(table.cols)))
where(fmt.Sprintf("YYY [%s] colIndex = %d", table.Name(), colIndex))

		colType := table.colTypes[colIndex]

		switch colType {
			case "string":
				modelRowCount = len(table.cols[colIndex].([]string))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "bool":
				modelRowCount = len(table.cols[colIndex].([]bool))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "int":
				modelRowCount = len(table.cols[colIndex].([]int))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "int8":
				modelRowCount = len(table.cols[colIndex].([]int8))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "int16":
				modelRowCount = len(table.cols[colIndex].([]int16))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "int32":
				modelRowCount = len(table.cols[colIndex].([]int32))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "int64":
				modelRowCount = len(table.cols[colIndex].([]int64))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "uint":
				modelRowCount = len(table.cols[colIndex].([]uint))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "byte":
				modelRowCount = len(table.cols[colIndex].([]byte))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "uint8":
				modelRowCount = len(table.cols[colIndex].([]uint8))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "uint16":
				modelRowCount = len(table.cols[colIndex].([]uint16))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "uint32":
				modelRowCount = len(table.cols[colIndex].([]uint32))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "uint64":
				modelRowCount = len(table.cols[colIndex].([]uint64))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "float32":
				modelRowCount = len(table.cols[colIndex].([]float32))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "float64":
				modelRowCount = len(table.cols[colIndex].([]float64))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "[]byte":
				modelRowCount = len(table.cols[colIndex].([][]byte))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "[]uint8":
				modelRowCount = len(table.cols[colIndex].([][]uint8))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			default:
				_,_ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR IN %s(): unknown type: %s\n", funcSource(), funcName(), colType))
				return -1
		}
	}

where(fmt.Sprintf("END OF FUNC [%s].%s() = %d", table.Name(), funcName(), modelRowCount))
	rowCount = modelRowCount

	return
}

/*
	Return the number of rows in this table.
*/
func (table *Table) new_model_RowCount() (rowCount int) {
where(fmt.Sprintf("inside [%s].%s()", table.Name(), funcName()))
// debug.PrintStack()

	if new_model {
		if table == nil {
			_,_ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: %s(): table is <nil>\n", funcSource(), funcName()))
			return -1
		}

		if table.cols == nil {
			_,_ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: %s(): [%s].cols = nil\n", funcSource(), table.Name(), funcName()))
			return -1
		}
	}

	if len(table.cols) == 0 {
		// Avoid index out of range indexing into table.cols in switch statement.
		// This implies rows cannot be appended before at least one column has been appended.
//debug.PrintStack()
		rowCount = 0
where(fmt.Sprintf("MIDDLE OF FUNC [%s].%s() = %d", table.Name(), funcName(), rowCount))
		return
	}
// where(fmt.Sprintf("len(table.cols) = %d", len(table.cols)))

	var modelRowCount int = -1
	var prevModelRowCount int = -1

// debug.PrintStack()
where(fmt.Sprintf("len(table.colNames) = %d", len(table.colNames)))
where(fmt.Sprintf("len(table.cols) = %d", len(table.cols)))

	for colIndex := 0; colIndex < len(table.cols); colIndex++ {

where(fmt.Sprintf("YYY [%s] colNames = %v", table.Name(), table.colNames))
where(fmt.Sprintf("YYY [%s] colTypes = %v", table.Name(), table.colTypes))
where(fmt.Sprintf("YYY [%s] table.cols = %v", table.Name(), table.cols))
where(fmt.Sprintf("YYY [%s] len(table.cols) = %d", table.Name(), len(table.cols)))
where(fmt.Sprintf("YYY [%s] colIndex = %d", table.Name(), colIndex))

		colType := table.colTypes[colIndex]

		switch colType {
			case "string":
				modelRowCount = len(table.cols[colIndex].([]string))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "bool":
				modelRowCount = len(table.cols[colIndex].([]bool))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "int":
				modelRowCount = len(table.cols[colIndex].([]int))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "int8":
				modelRowCount = len(table.cols[colIndex].([]int8))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "int16":
				modelRowCount = len(table.cols[colIndex].([]int16))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "int32":
				modelRowCount = len(table.cols[colIndex].([]int32))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "int64":
				modelRowCount = len(table.cols[colIndex].([]int64))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "uint":
				modelRowCount = len(table.cols[colIndex].([]uint))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "byte":
				modelRowCount = len(table.cols[colIndex].([]byte))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "uint8":
				modelRowCount = len(table.cols[colIndex].([]uint8))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "uint16":
				modelRowCount = len(table.cols[colIndex].([]uint16))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "uint32":
				modelRowCount = len(table.cols[colIndex].([]uint32))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "uint64":
				modelRowCount = len(table.cols[colIndex].([]uint64))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "float32":
				modelRowCount = len(table.cols[colIndex].([]float32))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "float64":
				modelRowCount = len(table.cols[colIndex].([]float64))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "[]byte":
				modelRowCount = len(table.cols[colIndex].([][]byte))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			case "[]uint8":
				modelRowCount = len(table.cols[colIndex].([][]uint8))
where(fmt.Sprintf("%s(): colIndex = %d modelRowCount = %d", funcName(), colIndex, modelRowCount))
				if prevModelRowCount > -1 && modelRowCount != prevModelRowCount {
					panic(fmt.Errorf("%s(): col %s (prevModelRowCount) %d != col %s modelRowCount %d ([%s].RowCount() = %d)",
						funcName(),
						table.colNames[colIndex-1],
						prevModelRowCount,
						table.colNames[colIndex],
						modelRowCount,
						table.Name(),
						table.RowCount()))
				}
				prevModelRowCount = modelRowCount
			default:
				_,_ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR IN %s(): unknown type: %s\n", funcSource(), funcName(), colType))
				return -1
		}
	}

where(fmt.Sprintf("END OF FUNC [%s].%s() = %d", table.Name(), funcName(), modelRowCount))
	rowCount = modelRowCount

	return
}

/*
	Delete a row from this table.
*/
func (table *Table) new_model_DeleteRow(rowIndex int) error {
	// new memory model
where(funcName())

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	_, err := table.IsValidTable()
	if err != nil { return err }

where()
	if len(table.cols) == 0 {
		return fmt.Errorf("[%s].%s(): cannot delete rows from table with zero cols",
			table.Name(), funcName())
	}

where()
	if rowIndex < 0 || rowIndex > table.new_model_RowCount()-1 {
		return fmt.Errorf("%s(): in table [%s] with %d rows, row index %d does not exist",
			funcName(), table.tableName, table.new_model_RowCount(), rowIndex)
	}

	_, err = table.IsValidTable()
	if err != nil { return err }

where()
	return table.new_model_DeleteRows(rowIndex, rowIndex)

//where(fmt.Sprintf("BEFORE deleting row %d table.new_model_RowCount() = %d", rowIndex, table.new_model_RowCount()))
//
//	for colIndex, colName := range table.colNames {
//
//		colType, err := table.ColType(colName)
//		if err != nil { return err }
//// where(colIndex)
//// where(colName)
//// where(colType)
//// where(fmt.Sprintf("[%s] %d %s %s []col type = %T", table.Name(), colIndex, colName, colType, table.cols[colIndex]))
//
//		switch colType {
//			//			case "string":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([]string)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "bool":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([]bool)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "int":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([]int)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "int8":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([]int8)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "int16":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([]int16)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "int32":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([]int32)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "int64":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([]int64)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "uint":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([]uint)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "byte":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([]byte)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "uint8":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([]uint8)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "uint16":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([]uint16)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "uint32":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([]uint32)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "uint64":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([]uint64)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "float32":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([]float32)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "float64":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([]float64)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "[]byte":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([][]byte)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//			//			case "[]uint8":
//where(fmt.Sprintf("%s(): deleting table.cols[%d][%d]", funcName(), colIndex, rowIndex))
//				col := table.cols[colIndex].([][]uint8)
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//				// From Ivo Balbaert p182 for deleting a single element from a slice.
//				col = append(col[:rowIndex], col[rowIndex+1:]...)
//				table.cols[colIndex] = col	// append may have returned a new col slice variable.
//where(fmt.Sprintf("%s(): len(col) = %d", funcName(), len(col)))
//// where()
//		//			default:
//				err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
//				return err
//		}
//	}
//where(fmt.Sprintf("AFTER  deleting row %d table.new_model_RowCount() = %d", rowIndex, table.new_model_RowCount()))
//
//	return nil
}

/*
	Delete a col from this table.
*/
func (table *Table) new_model_DeleteColByColIndex(colIndex int) error {
	// new memory model

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	if len(table.cols) == 0 {
		return fmt.Errorf("[%s].%s(): cannot delete cols from table with zero cols",
			table.Name(), funcName())
	}

	if colIndex < 0 || colIndex > table.new_model_ColCount()-1 {
		err := fmt.Errorf("in table [%s] with %d cols, col index %d does not exist",
			table.tableName, table.new_model_ColCount(), colIndex)
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
func (table *Table) new_model_DeleteRows(firstRowIndex int, lastRowIndex int) error {
	// new memory model
where(funcName())

	if table == nil { return fmt.Errorf("(new model) table.%s(): table is <nil>", funcName()) }

	if len(table.cols) == 0 {
		return fmt.Errorf("(new model) [%s].%s(): cannot delete rows from table with zero cols",
			table.Name(), funcName())
	}

	if firstRowIndex < 0 || firstRowIndex > table.new_model_RowCount()-1 {
		return fmt.Errorf("(new model) in table [%s] with %d rows, firstRowIndex %d does not exist",
			table.tableName, table.new_model_RowCount(), firstRowIndex)
	}

    if lastRowIndex < 0 || lastRowIndex > table.new_model_RowCount()-1 {
        return fmt.Errorf("(new model) in table [%s] with %d rows, lastRowIndex %d does not exist",
            table.tableName, table.new_model_RowCount(), lastRowIndex)
    }

    if firstRowIndex > lastRowIndex {
        return fmt.Errorf("(new model) invalid row index range: firstRowIndex %d > lastRowIndex %d", firstRowIndex, lastRowIndex)
    }
where(fmt.Sprintf("BEFORE new_model_DeleteRows() new_model_RowCount() = %d", table.new_model_RowCount()))

	for colIndex, colName := range table.colNames {

		colType, err := table.ColType(colName)
		if err != nil { return err }
where(colIndex)
where(colName)
where(colType)
where(fmt.Sprintf("[%s] %d %s %s []col type = %T", table.Name(), colIndex, colName, colType, table.cols[colIndex]))

		switch colType {
			case "string":
				col := table.cols[colIndex].([]string)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "bool":
				col := table.cols[colIndex].([]bool)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "int":
				col := table.cols[colIndex].([]int)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "int8":
				col := table.cols[colIndex].([]int8)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "int16":
				col := table.cols[colIndex].([]int16)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "int32":
				col := table.cols[colIndex].([]int32)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "int64":
				col := table.cols[colIndex].([]int64)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "uint":
				col := table.cols[colIndex].([]uint)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "byte":
				col := table.cols[colIndex].([]byte)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "uint8":
				col := table.cols[colIndex].([]uint8)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "uint16":
				col := table.cols[colIndex].([]uint16)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "uint32":
				col := table.cols[colIndex].([]uint32)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "uint64":
				col := table.cols[colIndex].([]uint64)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "float32":
				col := table.cols[colIndex].([]float32)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "float64":
				col := table.cols[colIndex].([]float64)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "[]byte":
				col := table.cols[colIndex].([][]byte)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			case "[]uint8":
				col := table.cols[colIndex].([][]uint8)
				// From Ivo Balbaert p182 for deleting a range of elements from a slice.
where(fmt.Sprintf("NNN BEFORE delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				col = append(col[:firstRowIndex], col[lastRowIndex+1:]...)
where(fmt.Sprintf("NNN AFTER  delete [%s].%s len(col) = %d", table.Name(), colName, len(col)))
				table.cols[colIndex] = col	// append may have returned a new col slice variable.
			default:
				err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
				return err
		}
	}
where(fmt.Sprintf("AFTER  new_model_DeleteRows() new_model_RowCount() = %d", table.new_model_RowCount()))

where("NNN BEFORE deleting from table.rowIndex")
where(table.rowsIndex)
	// From Ivo Balbaert p182 for deleting a range of elements from a slice.
	table.rowsIndex = append(table.rowsIndex[:firstRowIndex], table.rowsIndex[lastRowIndex+1:]...)
where("NNN AFTER  deleting from table.rowIndex")
where(table.rowsIndex)
	_, err := table.IsValidTable()
	if err != nil { return err }

	return nil
}

/*
	Set an interface{} val where you don't know or cannot easily select the specific type.
*/
func (table *Table) new_model_SetValByColIndex(colIndex int, rowIndex int, val interface{}) error {
where(fmt.Sprintf("%s(colIndex=%d, rowIndex=%d, val=%v)", funcName(), colIndex, rowIndex, val))
where(fmt.Sprintf("new_model_ColCount() = %d", table.new_model_ColCount()))
	// Note: With helper functions it is easy to set a val by its specific type.
	//       Essentially, it is now faster to set val by specific type than as an interface() value.
	//       For consistency, we retain this method.

	if table == nil { return fmt.Errorf("(new model) table.%s(): table is <nil>", funcName()) }

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	valType := fmt.Sprintf("%T", val)
	if valType != colType {
		if !isAlias(colType, valType) {
			colName, err := table.ColNameByColIndex(colIndex)
			if err != nil { return err }
			return fmt.Errorf("%s(): table [%s] col index %d col name %s expecting type %s not type %s",
				funcName(), table.Name(), colIndex, colName, colType, valType)
		}
	}

	switch colType {
		case "string":
			col := table.cols[colIndex].([]string)
			col[rowIndex] = val.(string)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "bool":
			col := table.cols[colIndex].([]bool)
			col[rowIndex] = val.(bool)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "int":
			col := table.cols[colIndex].([]int)
			col[rowIndex] = val.(int)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "int8":
			col := table.cols[colIndex].([]int8)
			col[rowIndex] = val.(int8)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "int16":
			col := table.cols[colIndex].([]int16)
			col[rowIndex] = val.(int16)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "int32":
			col := table.cols[colIndex].([]int32)
			col[rowIndex] = val.(int32)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "int64":
			col := table.cols[colIndex].([]int64)
			col[rowIndex] = val.(int64)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "uint":
			col := table.cols[colIndex].([]uint)
			col[rowIndex] = val.(uint)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "byte":
			col := table.cols[colIndex].([]byte)
			col[rowIndex] = val.(byte)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "uint8":
			col := table.cols[colIndex].([]uint8)
			col[rowIndex] = val.(uint8)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "uint16":
			col := table.cols[colIndex].([]uint16)
			col[rowIndex] = val.(uint16)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "uint32":
			col := table.cols[colIndex].([]uint32)
			col[rowIndex] = val.(uint32)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "uint64":
			col := table.cols[colIndex].([]uint64)
			col[rowIndex] = val.(uint64)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "float32":
			col := table.cols[colIndex].([]float32)
			col[rowIndex] = val.(float32)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "float64":
			col := table.cols[colIndex].([]float64)
			col[rowIndex] = val.(float64)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "[]byte":
			col := table.cols[colIndex].([][]byte)
			col[rowIndex] = val.([]byte)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.
		case "[]uint8":
			col := table.cols[colIndex].([][]uint8)
			col[rowIndex] = val.([]uint8)
//			table.cols[colIndex] = col	// append may have returned a new col slice variable.

		default:
			err := fmt.Errorf("ERROR IN %s(): unknown type: %s\n", funcName(), colType)
			return err
	}

	return nil
}

//	-------------------------------------------------------
//	Set<type>() functions for each of 17 types.
//	-------------------------------------------------------

//	Set table cell in colName at rowIndex to newValue string
func (table *Table) SetString(colName string, rowIndex int, newValue string) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([]string)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue bool
func (table *Table) SetBool(colName string, rowIndex int, newValue bool) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([]bool)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue int
func (table *Table) SetInt(colName string, rowIndex int, newValue int) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([]int)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue int8
func (table *Table) SetInt8(colName string, rowIndex int, newValue int8) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([]int8)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue int16
func (table *Table) SetInt16(colName string, rowIndex int, newValue int16) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([]int16)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue int32
func (table *Table) SetInt32(colName string, rowIndex int, newValue int32) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([]int32)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue int64
func (table *Table) SetInt64(colName string, rowIndex int, newValue int64) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([]int64)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue uint
func (table *Table) SetUint(colName string, rowIndex int, newValue uint) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([]uint)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue byte
func (table *Table) SetByte(colName string, rowIndex int, newValue byte) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([]byte)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue uint8
func (table *Table) SetUint8(colName string, rowIndex int, newValue uint8) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([]uint8)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue uint16
func (table *Table) SetUint16(colName string, rowIndex int, newValue uint16) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([]uint16)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue uint32
func (table *Table) SetUint32(colName string, rowIndex int, newValue uint32) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([]uint32)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue uint64
func (table *Table) SetUint64(colName string, rowIndex int, newValue uint64) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([]uint64)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue float32
func (table *Table) SetFloat32(colName string, rowIndex int, newValue float32) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([]float32)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue float64
func (table *Table) SetFloat64(colName string, rowIndex int, newValue float64) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([]float64)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue []byte
func (table *Table) SetByteSlice(colName string, rowIndex int, newValue []byte) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([][]byte)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colName at rowIndex to newValue []uint8
func (table *Table) SetUint8Slice(colName string, rowIndex int, newValue []uint8) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetVal(colName, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return err }

		col := table.cols[colIndex].([][]uint8)
		col[rowIndex] = newValue
	}

	return nil
}

//	-----------------------------------------------------------------
//	Set<type>ByColIndex() functions for each of 17 types.
//	-----------------------------------------------------------------

//	Set table cell in colIndex at rowIndex to newValue string
func (table *Table) SetStringByColIndex(colIndex int, rowIndex int, newValue string) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([]string)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue bool
func (table *Table) SetBoolByColIndex(colIndex int, rowIndex int, newValue bool) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([]bool)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue int
func (table *Table) SetIntByColIndex(colIndex int, rowIndex int, newValue int) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([]int)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue int8
func (table *Table) SetInt8ByColIndex(colIndex int, rowIndex int, newValue int8) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([]int8)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue int16
func (table *Table) SetInt16ByColIndex(colIndex int, rowIndex int, newValue int16) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([]int16)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue int32
func (table *Table) SetInt32ByColIndex(colIndex int, rowIndex int, newValue int32) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([]int32)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue int64
func (table *Table) SetInt64ByColIndex(colIndex int, rowIndex int, newValue int64) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([]int64)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue uint
func (table *Table) SetUintByColIndex(colIndex int, rowIndex int, newValue uint) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([]uint)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue byte
func (table *Table) SetByteByColIndex(colIndex int, rowIndex int, newValue byte) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([]byte)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue uint8
func (table *Table) SetUint8ByColIndex(colIndex int, rowIndex int, newValue uint8) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([]uint8)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue uint16
func (table *Table) SetUint16ByColIndex(colIndex int, rowIndex int, newValue uint16) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([]uint16)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue uint32
func (table *Table) SetUint32ByColIndex(colIndex int, rowIndex int, newValue uint32) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([]uint32)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue uint64
func (table *Table) SetUint64ByColIndex(colIndex int, rowIndex int, newValue uint64) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([]uint64)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue float32
func (table *Table) SetFloat32ByColIndex(colIndex int, rowIndex int, newValue float32) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([]float32)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue float64
func (table *Table) SetFloat64ByColIndex(colIndex int, rowIndex int, newValue float64) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([]float64)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue []byte
func (table *Table) SetByteSliceByColIndex(colIndex int, rowIndex int, newValue []byte) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([][]byte)
		col[rowIndex] = newValue
	}

	return nil
}

//	Set table cell in colIndex at rowIndex to newValue []uint8
func (table *Table) SetUint8SliceByColIndex(colIndex int, rowIndex int, newValue []uint8) error {
	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return err }

	// old_model
	err = table.SetValByColIndex(colIndex, rowIndex, newValue)
	if err != nil { return err }

	// new_model
	if new_model {
		col := table.cols[colIndex].([][]uint8)
		col[rowIndex] = newValue
	}

	return nil
}

//	-------------------------------------------------------
//	Get<type>() functions for each of 17 types.
//	-------------------------------------------------------

//	Get string table cell from colName at rowIndex
func (table *Table) GetString(colName string, rowIndex int) (value string, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(string)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]string)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get bool table cell from colName at rowIndex
func (table *Table) GetBool(colName string, rowIndex int) (value bool, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(bool)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]bool)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get int table cell from colName at rowIndex
func (table *Table) GetInt(colName string, rowIndex int) (value int, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(int)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]int)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get int8 table cell from colName at rowIndex
func (table *Table) GetInt8(colName string, rowIndex int) (value int8, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(int8)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]int8)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get int16 table cell from colName at rowIndex
func (table *Table) GetInt16(colName string, rowIndex int) (value int16, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(int16)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]int16)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get int32 table cell from colName at rowIndex
func (table *Table) GetInt32(colName string, rowIndex int) (value int32, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(int32)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]int32)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get int64 table cell from colName at rowIndex
func (table *Table) GetInt64(colName string, rowIndex int) (value int64, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(int64)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]int64)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get uint table cell from colName at rowIndex
func (table *Table) GetUint(colName string, rowIndex int) (value uint, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(uint)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]uint)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get byte table cell from colName at rowIndex
func (table *Table) GetByte(colName string, rowIndex int) (value byte, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(byte)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]byte)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get uint8 table cell from colName at rowIndex
func (table *Table) GetUint8(colName string, rowIndex int) (value uint8, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(uint8)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]uint8)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get uint16 table cell from colName at rowIndex
func (table *Table) GetUint16(colName string, rowIndex int) (value uint16, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(uint16)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]uint16)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get uint32 table cell from colName at rowIndex
func (table *Table) GetUint32(colName string, rowIndex int) (value uint32, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(uint32)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]uint32)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get uint64 table cell from colName at rowIndex
func (table *Table) GetUint64(colName string, rowIndex int) (value uint64, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(uint64)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]uint64)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get float32 table cell from colName at rowIndex
func (table *Table) GetFloat32(colName string, rowIndex int) (value float32, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(float32)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]float32)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get float64 table cell from colName at rowIndex
func (table *Table) GetFloat64(colName string, rowIndex int) (value float64, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(float64)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]float64)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get []byte table cell from colName at rowIndex
func (table *Table) GetByteSlice(colName string, rowIndex int) (value []byte, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.([]byte)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([][]byte)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		// We need to use bytes.Equal() to compare []byte and []uint8 slices.
		if !bytes.Equal(new_model_value, old_model_value) {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	Get []uint8 table cell from colName at rowIndex
func (table *Table) GetUint8Slice(colName string, rowIndex int) (value []uint8, err error) {

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.([]uint8)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	// new_model
	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([][]uint8)
		new_model_value := col[rowIndex]
		new_model_value = new_model_value	// Avoid compiler error.

		// We need to use bytes.Equal() to compare []byte and []uint8 slices.
		if !bytes.Equal(new_model_value, old_model_value) {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
		}
	}

	value = old_model_value

	return
}

//	-----------------------------------------------------------------
//	Get<type>ByColIndex() functions for each of 17 types.
//	-----------------------------------------------------------------

//	Set table cell in colIndex at rowIndex to newValue string
func (table *Table) GetStringByColIndex(colIndex int, rowIndex int) (value string, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.(string)

	// new_model
	if new_model {
		col := table.cols[colIndex].([]string)
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue bool
func (table *Table) GetBoolByColIndex(colIndex int, rowIndex int) (value bool, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.(bool)

	// new_model
	if new_model {
		col := table.cols[colIndex].([]bool)
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue int
func (table *Table) GetIntByColIndex(colIndex int, rowIndex int) (value int, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.(int)

	// new_model
	if new_model {
		col := table.cols[colIndex].([]int)
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue int8
func (table *Table) GetInt8ByColIndex(colIndex int, rowIndex int) (value int8, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.(int8)

	// new_model
	if new_model {
		col := table.cols[colIndex].([]int8)
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue int16
func (table *Table) GetInt16ByColIndex(colIndex int, rowIndex int) (value int16, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.(int16)

	// new_model
	if new_model {
		col := table.cols[colIndex].([]int16)
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue int32
func (table *Table) GetInt32ByColIndex(colIndex int, rowIndex int) (value int32, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.(int32)

	// new_model
	if new_model {
		col := table.cols[colIndex].([]int32)
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue int64
func (table *Table) GetInt64ByColIndex(colIndex int, rowIndex int) (value int64, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.(int64)

	// new_model
	if new_model {
		col := table.cols[colIndex].([]int64)
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue uint
func (table *Table) GetUintByColIndex(colIndex int, rowIndex int) (value uint, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.(uint)

	// new_model
	if new_model {
		col := table.cols[colIndex].([]uint)
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue byte
func (table *Table) GetByteByColIndex(colIndex int, rowIndex int) (value byte, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.(byte)

	// new_model
	if new_model {
		col := table.cols[colIndex].([]byte)
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue uint8
func (table *Table) GetUint8ByColIndex(colIndex int, rowIndex int) (value uint8, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.(uint8)

	// new_model
	if new_model {
		col := table.cols[colIndex].([]uint8)
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue uint16
func (table *Table) GetUint16ByColIndex(colIndex int, rowIndex int) (value uint16, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.(uint16)

	// new_model
	if new_model {
		col := table.cols[colIndex].([]uint16)
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue uint32
func (table *Table) GetUint32ByColIndex(colIndex int, rowIndex int) (value uint32, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.(uint32)

	// new_model
	if new_model {
		col := table.cols[colIndex].([]uint32)
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue uint64
func (table *Table) GetUint64ByColIndex(colIndex int, rowIndex int) (value uint64, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.(uint64)

	// new_model
	if new_model {
		col := table.cols[colIndex].([]uint64)
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue float32
func (table *Table) GetFloat32ByColIndex(colIndex int, rowIndex int) (value float32, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.(float32)

	// new_model
	if new_model {
		col := table.cols[colIndex].([]float32)
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue float64
func (table *Table) GetFloat64ByColIndex(colIndex int, rowIndex int) (value float64, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.(float64)

	// new_model
	if new_model {
		col := table.cols[colIndex].([]float64)
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue []byte
func (table *Table) GetByteSliceByColIndex(colIndex int, rowIndex int) (value []byte, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.([]byte)

	// new_model
	if new_model {
		col := table.cols[colIndex].([][]byte)
		new_model_value := col[rowIndex]

		// We need to use bytes.Equal() to compare []byte and []uint8 slices.
		if !bytes.Equal(new_model_value, old_model_value) {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

//	Set table cell in colIndex at rowIndex to newValue []uint8
func (table *Table) GetUint8SliceByColIndex(colIndex int, rowIndex int) (value []uint8, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell { return }

	// old_model
	interfaceValue, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil { return }
	old_model_value := interfaceValue.([]uint8)

	// new_model
	if new_model {
		col := table.cols[colIndex].([][]uint8)
		new_model_value := col[rowIndex]

		// We need to use bytes.Equal() to compare []byte and []uint8 slices.
		if !bytes.Equal(new_model_value, old_model_value) {
			err = fmt.Errorf("new_model_value %v != old_model_value %v", new_model_value, old_model_value)
			return
		}
	}

	value = old_model_value

	return
}

