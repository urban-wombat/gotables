package gotables

/*
	helpers.go
*/

import (
//	"bytes"
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

	value = old_model_value

	return
}

