package gotables

/*
	helpers.go
*/

import (
//	"bytes"
	"errors"
	"fmt"
//	"reflect"
//	"os"
//	"runtime/debug"
)

/*
Copyright (c) 2017-2018 Malcolm Gorman

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

//	------------------------------------------------------------------
//	next group: Set<type>() functions for each of 18 types.
//	------------------------------------------------------------------

//	Set table cell in colName at rowIndex to newVal []byte
func (table *Table) SetByteSlice(colName string, rowIndex int, newVal []byte) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "[]byte"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal []uint8
func (table *Table) SetUint8Slice(colName string, rowIndex int, newVal []uint8) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "[]uint8"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal bool
func (table *Table) SetBool(colName string, rowIndex int, newVal bool) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "bool"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal byte
func (table *Table) SetByte(colName string, rowIndex int, newVal byte) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "byte"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal float32
func (table *Table) SetFloat32(colName string, rowIndex int, newVal float32) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "float32"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal float64
func (table *Table) SetFloat64(colName string, rowIndex int, newVal float64) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "float64"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal int
func (table *Table) SetInt(colName string, rowIndex int, newVal int) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "int"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal int16
func (table *Table) SetInt16(colName string, rowIndex int, newVal int16) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "int16"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal int32
func (table *Table) SetInt32(colName string, rowIndex int, newVal int32) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "int32"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal int64
func (table *Table) SetInt64(colName string, rowIndex int, newVal int64) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "int64"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal int8
func (table *Table) SetInt8(colName string, rowIndex int, newVal int8) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "int8"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal rune
func (table *Table) SetRune(colName string, rowIndex int, newVal rune) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "rune"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal string
func (table *Table) SetString(colName string, rowIndex int, newVal string) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "string"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal uint
func (table *Table) SetUint(colName string, rowIndex int, newVal uint) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "uint"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal uint16
func (table *Table) SetUint16(colName string, rowIndex int, newVal uint16) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "uint16"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal uint32
func (table *Table) SetUint32(colName string, rowIndex int, newVal uint32) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "uint32"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal uint64
func (table *Table) SetUint64(colName string, rowIndex int, newVal uint64) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "uint64"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal uint8
func (table *Table) SetUint8(colName string, rowIndex int, newVal uint8) error {

	// See: Set<type>() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "uint8"

	colType, err := table.ColType(colName)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average %30 speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	----------------------------------------------------------------------------
//	next group: Set<type>ByColIndex() functions for each of 18 types.
//	----------------------------------------------------------------------------

//	Set table cell in colIndex at rowIndex to newVal []byte
func (table *Table) SetByteSliceByColIndex(colIndex int, rowIndex int, newVal []byte) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "[]byte"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal []uint8
func (table *Table) SetUint8SliceByColIndex(colIndex int, rowIndex int, newVal []uint8) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "[]uint8"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal bool
func (table *Table) SetBoolByColIndex(colIndex int, rowIndex int, newVal bool) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "bool"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal byte
func (table *Table) SetByteByColIndex(colIndex int, rowIndex int, newVal byte) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "byte"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal float32
func (table *Table) SetFloat32ByColIndex(colIndex int, rowIndex int, newVal float32) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "float32"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal float64
func (table *Table) SetFloat64ByColIndex(colIndex int, rowIndex int, newVal float64) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "float64"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal int
func (table *Table) SetIntByColIndex(colIndex int, rowIndex int, newVal int) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "int"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal int16
func (table *Table) SetInt16ByColIndex(colIndex int, rowIndex int, newVal int16) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "int16"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal int32
func (table *Table) SetInt32ByColIndex(colIndex int, rowIndex int, newVal int32) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "int32"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal int64
func (table *Table) SetInt64ByColIndex(colIndex int, rowIndex int, newVal int64) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "int64"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal int8
func (table *Table) SetInt8ByColIndex(colIndex int, rowIndex int, newVal int8) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "int8"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal rune
func (table *Table) SetRuneByColIndex(colIndex int, rowIndex int, newVal rune) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "rune"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal string
func (table *Table) SetStringByColIndex(colIndex int, rowIndex int, newVal string) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "string"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal uint
func (table *Table) SetUintByColIndex(colIndex int, rowIndex int, newVal uint) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "uint"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal uint16
func (table *Table) SetUint16ByColIndex(colIndex int, rowIndex int, newVal uint16) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "uint16"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal uint32
func (table *Table) SetUint32ByColIndex(colIndex int, rowIndex int, newVal uint32) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "uint32"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal uint64
func (table *Table) SetUint64ByColIndex(colIndex int, rowIndex int, newVal uint64) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "uint64"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal uint8
func (table *Table) SetUint8ByColIndex(colIndex int, rowIndex int, newVal uint8) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil { return fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "uint8"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				funcName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return err }

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	------------------------------------------------------------------
//	next group: Get<type>() functions for each of 18 types.
//	------------------------------------------------------------------

//	Get []byte table cell from colName at rowIndex
func (table *Table) GetByteSlice(colName string, rowIndex int) (val []byte, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "[]byte"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].([]byte)

	return
}

//	Get []uint8 table cell from colName at rowIndex
func (table *Table) GetUint8Slice(colName string, rowIndex int) (val []uint8, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "[]uint8"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].([]uint8)

	return
}

//	Get bool table cell from colName at rowIndex
func (table *Table) GetBool(colName string, rowIndex int) (val bool, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "bool"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(bool)

	return
}

//	Get byte table cell from colName at rowIndex
func (table *Table) GetByte(colName string, rowIndex int) (val byte, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "byte"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(byte)

	return
}

//	Get float32 table cell from colName at rowIndex
func (table *Table) GetFloat32(colName string, rowIndex int) (val float32, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "float32"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(float32)

	return
}

//	Get float64 table cell from colName at rowIndex
func (table *Table) GetFloat64(colName string, rowIndex int) (val float64, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "float64"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(float64)

	return
}

//	Get int table cell from colName at rowIndex
func (table *Table) GetInt(colName string, rowIndex int) (val int, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "int"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(int)

	return
}

//	Get int16 table cell from colName at rowIndex
func (table *Table) GetInt16(colName string, rowIndex int) (val int16, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "int16"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(int16)

	return
}

//	Get int32 table cell from colName at rowIndex
func (table *Table) GetInt32(colName string, rowIndex int) (val int32, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "int32"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(int32)

	return
}

//	Get int64 table cell from colName at rowIndex
func (table *Table) GetInt64(colName string, rowIndex int) (val int64, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "int64"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(int64)

	return
}

//	Get int8 table cell from colName at rowIndex
func (table *Table) GetInt8(colName string, rowIndex int) (val int8, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "int8"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(int8)

	return
}

//	Get rune table cell from colName at rowIndex
func (table *Table) GetRune(colName string, rowIndex int) (val rune, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "rune"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(rune)

	return
}

//	Get string table cell from colName at rowIndex
func (table *Table) GetString(colName string, rowIndex int) (val string, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "string"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(string)

	return
}

//	Get uint table cell from colName at rowIndex
func (table *Table) GetUint(colName string, rowIndex int) (val uint, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "uint"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(uint)

	return
}

//	Get uint16 table cell from colName at rowIndex
func (table *Table) GetUint16(colName string, rowIndex int) (val uint16, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "uint16"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(uint16)

	return
}

//	Get uint32 table cell from colName at rowIndex
func (table *Table) GetUint32(colName string, rowIndex int) (val uint32, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "uint32"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(uint32)

	return
}

//	Get uint64 table cell from colName at rowIndex
func (table *Table) GetUint64(colName string, rowIndex int) (val uint64, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "uint64"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(uint64)

	return
}

//	Get uint8 table cell from colName at rowIndex
func (table *Table) GetUint8(colName string, rowIndex int) (val uint8, err error) {

	// See: Get<type>() functions

	if table == nil { return val, fmt.Errorf("table.%s(): table is <nil>", funcName()) }

	const valType string = "uint8"

	colType, err := table.ColType(colName)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				funcName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil { return val, err }

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %15 speedup.
	val = table.rows[rowIndex][colIndex].(uint8)

	return
}

//	----------------------------------------------------------------------------
//	next group: Get<type>ByColIndex() functions for each of 18 types.
//	----------------------------------------------------------------------------

//	Set table cell in colIndex at rowIndex to newValue []byte
func (table *Table) GetByteSliceByColIndex(colIndex int, rowIndex int) (val []byte, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "[]byte"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].([]byte)

	return
}

//	Set table cell in colIndex at rowIndex to newValue []uint8
func (table *Table) GetUint8SliceByColIndex(colIndex int, rowIndex int) (val []uint8, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "[]uint8"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].([]uint8)

	return
}

//	Set table cell in colIndex at rowIndex to newValue bool
func (table *Table) GetBoolByColIndex(colIndex int, rowIndex int) (val bool, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "bool"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(bool)

	return
}

//	Set table cell in colIndex at rowIndex to newValue byte
func (table *Table) GetByteByColIndex(colIndex int, rowIndex int) (val byte, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "byte"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(byte)

	return
}

//	Set table cell in colIndex at rowIndex to newValue float32
func (table *Table) GetFloat32ByColIndex(colIndex int, rowIndex int) (val float32, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "float32"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(float32)

	return
}

//	Set table cell in colIndex at rowIndex to newValue float64
func (table *Table) GetFloat64ByColIndex(colIndex int, rowIndex int) (val float64, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "float64"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(float64)

	return
}

//	Set table cell in colIndex at rowIndex to newValue int
func (table *Table) GetIntByColIndex(colIndex int, rowIndex int) (val int, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "int"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(int)

	return
}

//	Set table cell in colIndex at rowIndex to newValue int16
func (table *Table) GetInt16ByColIndex(colIndex int, rowIndex int) (val int16, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "int16"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(int16)

	return
}

//	Set table cell in colIndex at rowIndex to newValue int32
func (table *Table) GetInt32ByColIndex(colIndex int, rowIndex int) (val int32, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "int32"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(int32)

	return
}

//	Set table cell in colIndex at rowIndex to newValue int64
func (table *Table) GetInt64ByColIndex(colIndex int, rowIndex int) (val int64, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "int64"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(int64)

	return
}

//	Set table cell in colIndex at rowIndex to newValue int8
func (table *Table) GetInt8ByColIndex(colIndex int, rowIndex int) (val int8, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "int8"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(int8)

	return
}

//	Set table cell in colIndex at rowIndex to newValue rune
func (table *Table) GetRuneByColIndex(colIndex int, rowIndex int) (val rune, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "rune"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(rune)

	return
}

//	Set table cell in colIndex at rowIndex to newValue string
func (table *Table) GetStringByColIndex(colIndex int, rowIndex int) (val string, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "string"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(string)

	return
}

//	Set table cell in colIndex at rowIndex to newValue uint
func (table *Table) GetUintByColIndex(colIndex int, rowIndex int) (val uint, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "uint"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(uint)

	return
}

//	Set table cell in colIndex at rowIndex to newValue uint16
func (table *Table) GetUint16ByColIndex(colIndex int, rowIndex int) (val uint16, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "uint16"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(uint16)

	return
}

//	Set table cell in colIndex at rowIndex to newValue uint32
func (table *Table) GetUint32ByColIndex(colIndex int, rowIndex int) (val uint32, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "uint32"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(uint32)

	return
}

//	Set table cell in colIndex at rowIndex to newValue uint64
func (table *Table) GetUint64ByColIndex(colIndex int, rowIndex int) (val uint64, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "uint64"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(uint64)

	return
}

//	Set table cell in colIndex at rowIndex to newValue uint8
func (table *Table) GetUint8ByColIndex(colIndex int, rowIndex int) (val uint8, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", funcName())
		return
	}

	const valType string = "uint8"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil { return val, err }

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				funcName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow { return val, err }

	// Get the val
	// Note: This essentially inlines GetVal(): an average %25 speedup.
	val = table.rows[rowIndex][colIndex].(uint8)

	return
}

func (table *Table) SetCellToZeroValueByColIndex(colIndex int, rowIndex int) error {
	// TODO: Test for colIndex or rowIndex out of range? Or is this done by underlying functions?

	if table == nil { return fmt.Errorf("table.%s: table is <nil>", funcName()) }

	var err error
	var colType string

	colType, err = table.ColTypeByColIndex(colIndex)
	if err != nil {
		return err
	}

	switch colType {
		case "[]byte":
			err = table.SetByteSliceByColIndex(colIndex, rowIndex, []byte{})
		case "[]uint8":
			err = table.SetUint8SliceByColIndex(colIndex, rowIndex, []uint8{})
		case "bool":
			err = table.SetBoolByColIndex(colIndex, rowIndex, false)
		case "byte":
			err = table.SetByteByColIndex(colIndex, rowIndex, 0)
		case "float32":
			err = table.SetFloat32ByColIndex(colIndex, rowIndex, 0.0)
		case "float64":
			err = table.SetFloat64ByColIndex(colIndex, rowIndex, 0.0)
		case "int":
			err = table.SetIntByColIndex(colIndex, rowIndex, 0)
		case "int16":
			err = table.SetInt16ByColIndex(colIndex, rowIndex, 0)
		case "int32":
			err = table.SetInt32ByColIndex(colIndex, rowIndex, 0)
		case "int64":
			err = table.SetInt64ByColIndex(colIndex, rowIndex, 0)
		case "int8":
			err = table.SetInt8ByColIndex(colIndex, rowIndex, 0)
		case "rune":
			err = table.SetRuneByColIndex(colIndex, rowIndex, 0)
		case "string":
			err = table.SetStringByColIndex(colIndex, rowIndex, "")
		case "uint":
			err = table.SetUintByColIndex(colIndex, rowIndex, 0)
		case "uint16":
			err = table.SetUint16ByColIndex(colIndex, rowIndex, 0)
		case "uint32":
			err = table.SetUint32ByColIndex(colIndex, rowIndex, 0)
		case "uint64":
			err = table.SetUint64ByColIndex(colIndex, rowIndex, 0)
		case "uint8":
			err = table.SetUint8ByColIndex(colIndex, rowIndex, 0)
		default:
			msg := fmt.Sprintf("invalid type: %s (Valid types:", colType)
			// Note: Because maps are not ordered, this (desirably) shuffles the order of valid col types with each call.
			for typeName, _ := range globalColTypesMap {
				msg += fmt.Sprintf(" %s", typeName)
			}
			msg += ")"
			err = errors.New(msg)
			return err
	}
	if err != nil {
		return err
	}

	return nil
}
