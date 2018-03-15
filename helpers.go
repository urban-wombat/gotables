package gotables

/*
	helpers.go
*/

import (
	"bytes"
	"fmt"
//	"os"
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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Set<type>ByColIndex() functions

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

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(string)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]string)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v (rowIndex = %d col = %v)", new_model_value, old_model_value,
				rowIndex, col)
		}
	}

	value = old_model_value

	return
}

//	Get bool table cell from colName at rowIndex
func (table *Table) GetBool(colName string, rowIndex int) (value bool, err error) {

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(bool)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]bool)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v (rowIndex = %d col = %v)", new_model_value, old_model_value,
				rowIndex, col)
		}
	}

	value = old_model_value

	return
}

//	Get int table cell from colName at rowIndex
func (table *Table) GetInt(colName string, rowIndex int) (value int, err error) {

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(int)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]int)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v (rowIndex = %d col = %v)", new_model_value, old_model_value,
				rowIndex, col)
		}
	}

	value = old_model_value

	return
}

//	Get int8 table cell from colName at rowIndex
func (table *Table) GetInt8(colName string, rowIndex int) (value int8, err error) {

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(int8)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]int8)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v (rowIndex = %d col = %v)", new_model_value, old_model_value,
				rowIndex, col)
		}
	}

	value = old_model_value

	return
}

//	Get int16 table cell from colName at rowIndex
func (table *Table) GetInt16(colName string, rowIndex int) (value int16, err error) {

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(int16)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]int16)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v (rowIndex = %d col = %v)", new_model_value, old_model_value,
				rowIndex, col)
		}
	}

	value = old_model_value

	return
}

//	Get int32 table cell from colName at rowIndex
func (table *Table) GetInt32(colName string, rowIndex int) (value int32, err error) {

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(int32)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]int32)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v (rowIndex = %d col = %v)", new_model_value, old_model_value,
				rowIndex, col)
		}
	}

	value = old_model_value

	return
}

//	Get int64 table cell from colName at rowIndex
func (table *Table) GetInt64(colName string, rowIndex int) (value int64, err error) {

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(int64)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]int64)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v (rowIndex = %d col = %v)", new_model_value, old_model_value,
				rowIndex, col)
		}
	}

	value = old_model_value

	return
}

//	Get uint table cell from colName at rowIndex
func (table *Table) GetUint(colName string, rowIndex int) (value uint, err error) {

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(uint)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]uint)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v (rowIndex = %d col = %v)", new_model_value, old_model_value,
				rowIndex, col)
		}
	}

	value = old_model_value

	return
}

//	Get byte table cell from colName at rowIndex
func (table *Table) GetByte(colName string, rowIndex int) (value byte, err error) {

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(byte)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]byte)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v (rowIndex = %d col = %v)", new_model_value, old_model_value,
				rowIndex, col)
		}
	}

	value = old_model_value

	return
}

//	Get uint8 table cell from colName at rowIndex
func (table *Table) GetUint8(colName string, rowIndex int) (value uint8, err error) {

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(uint8)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]uint8)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v (rowIndex = %d col = %v)", new_model_value, old_model_value,
				rowIndex, col)
		}
	}

	value = old_model_value

	return
}

//	Get uint16 table cell from colName at rowIndex
func (table *Table) GetUint16(colName string, rowIndex int) (value uint16, err error) {

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(uint16)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]uint16)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v (rowIndex = %d col = %v)", new_model_value, old_model_value,
				rowIndex, col)
		}
	}

	value = old_model_value

	return
}

//	Get uint32 table cell from colName at rowIndex
func (table *Table) GetUint32(colName string, rowIndex int) (value uint32, err error) {

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(uint32)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]uint32)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v (rowIndex = %d col = %v)", new_model_value, old_model_value,
				rowIndex, col)
		}
	}

	value = old_model_value

	return
}

//	Get uint64 table cell from colName at rowIndex
func (table *Table) GetUint64(colName string, rowIndex int) (value uint64, err error) {

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(uint64)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]uint64)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v (rowIndex = %d col = %v)", new_model_value, old_model_value,
				rowIndex, col)
		}
	}

	value = old_model_value

	return
}

//	Get float32 table cell from colName at rowIndex
func (table *Table) GetFloat32(colName string, rowIndex int) (value float32, err error) {

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(float32)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]float32)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v (rowIndex = %d col = %v)", new_model_value, old_model_value,
				rowIndex, col)
		}
	}

	value = old_model_value

	return
}

//	Get float64 table cell from colName at rowIndex
func (table *Table) GetFloat64(colName string, rowIndex int) (value float64, err error) {

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.(float64)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([]float64)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

		if new_model_value != old_model_value {
			return value, fmt.Errorf("new_model_value %v != old_model_value %v (rowIndex = %d col = %v)", new_model_value, old_model_value,
				rowIndex, col)
		}
	}

	value = old_model_value

	return
}

//	Get []byte table cell from colName at rowIndex
func (table *Table) GetByteSlice(colName string, rowIndex int) (value []byte, err error) {

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.([]byte)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([][]byte)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

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

	// See: Get<type>() functions

	if table == nil { return value, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil { return value, err }

	old_model_value, valid := interfaceType.([]uint8)
	if !valid {
		_, err = table.IsColType(colName, "string") // Get an error message.
		return value, err
	}

	if new_model {
		colIndex, err := table.ColIndex(colName)
		if err != nil { return value, err}

		col := table.cols[colIndex].([][]uint8)
where(fmt.Sprintf("len(col) = %d col = %v", len(col), col))
where(fmt.Sprintf("len(rowsIndex) = %d rowsIndex = %v", len(table.rowsIndex), table.rowsIndex))
		new_model_value := col[rowIndex]

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

	// See: Get<type>ByColIndex() functions

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

