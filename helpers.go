package gotables

/*
	helpers.go
*/

import (
	"fmt"
	"time"
)

/*
Copyright (c) 2017-2020 Malcolm Gorman

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

// Types are defined in helpersmain.go

// The 19 gotables column type constants.
const (
	ByteSlice     = "[]byte"
	Uint8Slice    = "[]uint8"
	Bool          = "bool"
	Byte          = "byte"
	Float32       = "float32"
	Float64       = "float64"
	Int           = "int"
	Int16         = "int16"
	Int32         = "int32"
	Int64         = "int64"
	Int8          = "int8"
	Rune          = "rune"
	String        = "string"
	Uint          = "uint"
	Uint16        = "uint16"
	Uint32        = "uint32"
	Uint64        = "uint64"
	Uint8         = "uint8"
	GotablesTable = "*Table"
	TimeTime      = "time.Time"
)

//	------------------------------------------------------------------
//	next group: Set<type>() functions for each of 20 types
//	20 types: *Table []byte []uint8 bool byte float32 float64 int int16 int32 int64 int8 rune string time.Time uint uint16 uint32 uint64 uint8
//  NOTE: Types are defined in helpersmain.go AND parser.go
//	------------------------------------------------------------------

//	Set table cell in colName at rowIndex to newVal []byte
func (table *Table) SetByteSlice(colName string, rowIndex int, newVal []byte) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "[]byte"

	// Do not allow <nil> for type []byte
	if newVal == nil {
		return fmt.Errorf("%s: table [%s] col %s expecting val of type []byte, not: <nil>",
			UtilFuncName(), table.Name(), colName)
	}

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal []uint8
func (table *Table) SetUint8Slice(colName string, rowIndex int, newVal []uint8) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "[]uint8"

	// Do not allow <nil> for type []uint8
	if newVal == nil {
		return fmt.Errorf("%s: table [%s] col %s expecting val of type []uint8, not: <nil>",
			UtilFuncName(), table.Name(), colName)
	}

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal bool
func (table *Table) SetBool(colName string, rowIndex int, newVal bool) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "bool"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal byte
func (table *Table) SetByte(colName string, rowIndex int, newVal byte) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "byte"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal float32
func (table *Table) SetFloat32(colName string, rowIndex int, newVal float32) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "float32"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal float64
func (table *Table) SetFloat64(colName string, rowIndex int, newVal float64) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "float64"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal int
func (table *Table) SetInt(colName string, rowIndex int, newVal int) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "int"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal int16
func (table *Table) SetInt16(colName string, rowIndex int, newVal int16) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "int16"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal int32
func (table *Table) SetInt32(colName string, rowIndex int, newVal int32) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "int32"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal int64
func (table *Table) SetInt64(colName string, rowIndex int, newVal int64) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "int64"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal int8
func (table *Table) SetInt8(colName string, rowIndex int, newVal int8) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "int8"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal rune
func (table *Table) SetRune(colName string, rowIndex int, newVal rune) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "rune"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal string
func (table *Table) SetString(colName string, rowIndex int, newVal string) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "string"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal uint
func (table *Table) SetUint(colName string, rowIndex int, newVal uint) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "uint"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal uint16
func (table *Table) SetUint16(colName string, rowIndex int, newVal uint16) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "uint16"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal uint32
func (table *Table) SetUint32(colName string, rowIndex int, newVal uint32) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "uint32"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal uint64
func (table *Table) SetUint64(colName string, rowIndex int, newVal uint64) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "uint64"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal uint8
func (table *Table) SetUint8(colName string, rowIndex int, newVal uint8) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "uint8"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal *Table
func (table *Table) SetTable(colName string, rowIndex int, newVal *Table) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "*Table"

	// Do not allow <nil> for type *Table
	if newVal == nil {
		return fmt.Errorf("%s(%s, %d, val): table [%s] col %s expecting val of type *Table, not: <nil> [use NewNilTable() instead of <nil>]",
			UtilFuncNameNoParens(), colName, rowIndex, table.Name(), colName)
	}

	// This may help with reporting errors.
	newVal.parentTable = table

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colName at rowIndex to newVal time.Time
func (table *Table) SetTime(colName string, rowIndex int, newVal time.Time) error {

	// See: Set<type>() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "time.Time"

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colName, colType, valType, newVal)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	----------------------------------------------------------------------------
//	next group: Set<type>ByColIndex() functions for each of 20 types
//	20 types: *Table []byte []uint8 bool byte float32 float64 int int16 int32 int64 int8 rune string time.Time uint uint16 uint32 uint64 uint8
//  NOTE: Types are defined in helpersmain.go AND parser.go
//	----------------------------------------------------------------------------

//	Set table cell in colIndex at rowIndex to newVal []byte
func (table *Table) SetByteSliceByColIndex(colIndex int, rowIndex int, newVal []byte) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "[]byte"

	// Do not allow <nil> for type []byte
	if newVal == nil {
		return fmt.Errorf("%s: table [%s] col %d expecting val of type []byte, not: <nil>",
			UtilFuncName(), table.Name(), colIndex)
	}

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal []uint8
func (table *Table) SetUint8SliceByColIndex(colIndex int, rowIndex int, newVal []uint8) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "[]uint8"

	// Do not allow <nil> for type []uint8
	if newVal == nil {
		return fmt.Errorf("%s: table [%s] col %d expecting val of type []uint8, not: <nil>",
			UtilFuncName(), table.Name(), colIndex)
	}

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal bool
func (table *Table) SetBoolByColIndex(colIndex int, rowIndex int, newVal bool) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "bool"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal byte
func (table *Table) SetByteByColIndex(colIndex int, rowIndex int, newVal byte) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "byte"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal float32
func (table *Table) SetFloat32ByColIndex(colIndex int, rowIndex int, newVal float32) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "float32"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal float64
func (table *Table) SetFloat64ByColIndex(colIndex int, rowIndex int, newVal float64) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "float64"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal int
func (table *Table) SetIntByColIndex(colIndex int, rowIndex int, newVal int) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "int"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal int16
func (table *Table) SetInt16ByColIndex(colIndex int, rowIndex int, newVal int16) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "int16"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal int32
func (table *Table) SetInt32ByColIndex(colIndex int, rowIndex int, newVal int32) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "int32"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal int64
func (table *Table) SetInt64ByColIndex(colIndex int, rowIndex int, newVal int64) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "int64"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal int8
func (table *Table) SetInt8ByColIndex(colIndex int, rowIndex int, newVal int8) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "int8"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal rune
func (table *Table) SetRuneByColIndex(colIndex int, rowIndex int, newVal rune) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "rune"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal string
func (table *Table) SetStringByColIndex(colIndex int, rowIndex int, newVal string) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "string"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal uint
func (table *Table) SetUintByColIndex(colIndex int, rowIndex int, newVal uint) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "uint"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal uint16
func (table *Table) SetUint16ByColIndex(colIndex int, rowIndex int, newVal uint16) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "uint16"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal uint32
func (table *Table) SetUint32ByColIndex(colIndex int, rowIndex int, newVal uint32) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "uint32"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal uint64
func (table *Table) SetUint64ByColIndex(colIndex int, rowIndex int, newVal uint64) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "uint64"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal uint8
func (table *Table) SetUint8ByColIndex(colIndex int, rowIndex int, newVal uint8) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "uint8"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal *Table
func (table *Table) SetTableByColIndex(colIndex int, rowIndex int, newVal *Table) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "*Table"

	// Do not allow <nil> for type *Table
	if newVal == nil {
		return fmt.Errorf("%s(%d, %d, val): table [%s] col %d expecting val of type *Table, not: <nil> [use NewNilTable() instead of <nil>]",
			UtilFuncNameNoParens(), colIndex, rowIndex, table.Name(), colIndex)
	}

	// This may help with reporting errors.
	newVal.parentTable = table

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	Set table cell in colIndex at rowIndex to newVal time.Time
func (table *Table) SetTimeByColIndex(colIndex int, rowIndex int, newVal time.Time) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "time.Time"

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				UtilFuncName(), table.Name(), colIndex, colType, valType, newVal)
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

//	------------------------------------------------------------------
//	next group: Get<type>() functions for each of 20 types
//	20 types: *Table []byte []uint8 bool byte float32 float64 int int16 int32 int64 int8 rune string time.Time uint uint16 uint32 uint64 uint8
//  NOTE: Types are defined in helpersmain.go AND parser.go
//	------------------------------------------------------------------

//	Get []byte table cell from colName at rowIndex
func (table *Table) GetByteSlice(colName string, rowIndex int) (val []byte, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "[]byte"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].([]byte)

	return
}

//	Get []uint8 table cell from colName at rowIndex
func (table *Table) GetUint8Slice(colName string, rowIndex int) (val []uint8, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "[]uint8"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].([]uint8)

	return
}

//	Get bool table cell from colName at rowIndex
func (table *Table) GetBool(colName string, rowIndex int) (val bool, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "bool"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(bool)

	return
}

//	Get byte table cell from colName at rowIndex
func (table *Table) GetByte(colName string, rowIndex int) (val byte, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "byte"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(byte)

	return
}

//	Get float32 table cell from colName at rowIndex
func (table *Table) GetFloat32(colName string, rowIndex int) (val float32, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "float32"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(float32)

	return
}

//	Get float64 table cell from colName at rowIndex
func (table *Table) GetFloat64(colName string, rowIndex int) (val float64, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "float64"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(float64)

	return
}

//	Get int table cell from colName at rowIndex
func (table *Table) GetInt(colName string, rowIndex int) (val int, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "int"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(int)

	return
}

//	Get int16 table cell from colName at rowIndex
func (table *Table) GetInt16(colName string, rowIndex int) (val int16, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "int16"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(int16)

	return
}

//	Get int32 table cell from colName at rowIndex
func (table *Table) GetInt32(colName string, rowIndex int) (val int32, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "int32"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(int32)

	return
}

//	Get int64 table cell from colName at rowIndex
func (table *Table) GetInt64(colName string, rowIndex int) (val int64, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "int64"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(int64)

	return
}

//	Get int8 table cell from colName at rowIndex
func (table *Table) GetInt8(colName string, rowIndex int) (val int8, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "int8"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(int8)

	return
}

//	Get rune table cell from colName at rowIndex
func (table *Table) GetRune(colName string, rowIndex int) (val rune, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "rune"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(rune)

	return
}

//	Get string table cell from colName at rowIndex
func (table *Table) GetString(colName string, rowIndex int) (val string, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "string"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(string)

	return
}

//	Get uint table cell from colName at rowIndex
func (table *Table) GetUint(colName string, rowIndex int) (val uint, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "uint"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(uint)

	return
}

//	Get uint16 table cell from colName at rowIndex
func (table *Table) GetUint16(colName string, rowIndex int) (val uint16, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "uint16"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(uint16)

	return
}

//	Get uint32 table cell from colName at rowIndex
func (table *Table) GetUint32(colName string, rowIndex int) (val uint32, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "uint32"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(uint32)

	return
}

//	Get uint64 table cell from colName at rowIndex
func (table *Table) GetUint64(colName string, rowIndex int) (val uint64, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "uint64"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(uint64)

	return
}

//	Get uint8 table cell from colName at rowIndex
func (table *Table) GetUint8(colName string, rowIndex int) (val uint8, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "uint8"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(uint8)

	return
}

//	Get *Table table cell from colName at rowIndex
func (table *Table) GetTable(colName string, rowIndex int) (val *Table, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "*Table"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(*Table)

	return
}

//	Get time.Time table cell from colName at rowIndex
func (table *Table) GetTime(colName string, rowIndex int) (val time.Time, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	const valType string = "time.Time"

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				UtilFuncName(), table.Name(), colName, colType)
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex].(time.Time)

	return
}

/*	Get []byte table cell from colName at rowIndex

	Like its non-MustGet alternative GetByteSlice(), but panics on error, and does not return an error.
*/
func (table *Table) GetByteSliceMustGet(colName string, rowIndex int) (val []byte) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetByteSlice(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get []uint8 table cell from colName at rowIndex

	Like its non-MustGet alternative GetUint8Slice(), but panics on error, and does not return an error.
*/
func (table *Table) GetUint8SliceMustGet(colName string, rowIndex int) (val []uint8) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetUint8Slice(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get bool table cell from colName at rowIndex

	Like its non-MustGet alternative GetBool(), but panics on error, and does not return an error.
*/
func (table *Table) GetBoolMustGet(colName string, rowIndex int) (val bool) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetBool(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get byte table cell from colName at rowIndex

	Like its non-MustGet alternative GetByte(), but panics on error, and does not return an error.
*/
func (table *Table) GetByteMustGet(colName string, rowIndex int) (val byte) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetByte(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get float32 table cell from colName at rowIndex

	Like its non-MustGet alternative GetFloat32(), but panics on error, and does not return an error.
*/
func (table *Table) GetFloat32MustGet(colName string, rowIndex int) (val float32) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetFloat32(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get float64 table cell from colName at rowIndex

	Like its non-MustGet alternative GetFloat64(), but panics on error, and does not return an error.
*/
func (table *Table) GetFloat64MustGet(colName string, rowIndex int) (val float64) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetFloat64(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get int table cell from colName at rowIndex

	Like its non-MustGet alternative GetInt(), but panics on error, and does not return an error.
*/
func (table *Table) GetIntMustGet(colName string, rowIndex int) (val int) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetInt(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get int16 table cell from colName at rowIndex

	Like its non-MustGet alternative GetInt16(), but panics on error, and does not return an error.
*/
func (table *Table) GetInt16MustGet(colName string, rowIndex int) (val int16) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetInt16(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get int32 table cell from colName at rowIndex

	Like its non-MustGet alternative GetInt32(), but panics on error, and does not return an error.
*/
func (table *Table) GetInt32MustGet(colName string, rowIndex int) (val int32) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetInt32(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get int64 table cell from colName at rowIndex

	Like its non-MustGet alternative GetInt64(), but panics on error, and does not return an error.
*/
func (table *Table) GetInt64MustGet(colName string, rowIndex int) (val int64) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetInt64(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get int8 table cell from colName at rowIndex

	Like its non-MustGet alternative GetInt8(), but panics on error, and does not return an error.
*/
func (table *Table) GetInt8MustGet(colName string, rowIndex int) (val int8) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetInt8(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get rune table cell from colName at rowIndex

	Like its non-MustGet alternative GetRune(), but panics on error, and does not return an error.
*/
func (table *Table) GetRuneMustGet(colName string, rowIndex int) (val rune) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetRune(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get string table cell from colName at rowIndex

	Like its non-MustGet alternative GetString(), but panics on error, and does not return an error.
*/
func (table *Table) GetStringMustGet(colName string, rowIndex int) (val string) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetString(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get uint table cell from colName at rowIndex

	Like its non-MustGet alternative GetUint(), but panics on error, and does not return an error.
*/
func (table *Table) GetUintMustGet(colName string, rowIndex int) (val uint) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetUint(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get uint16 table cell from colName at rowIndex

	Like its non-MustGet alternative GetUint16(), but panics on error, and does not return an error.
*/
func (table *Table) GetUint16MustGet(colName string, rowIndex int) (val uint16) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetUint16(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get uint32 table cell from colName at rowIndex

	Like its non-MustGet alternative GetUint32(), but panics on error, and does not return an error.
*/
func (table *Table) GetUint32MustGet(colName string, rowIndex int) (val uint32) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetUint32(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get uint64 table cell from colName at rowIndex

	Like its non-MustGet alternative GetUint64(), but panics on error, and does not return an error.
*/
func (table *Table) GetUint64MustGet(colName string, rowIndex int) (val uint64) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetUint64(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get uint8 table cell from colName at rowIndex

	Like its non-MustGet alternative GetUint8(), but panics on error, and does not return an error.
*/
func (table *Table) GetUint8MustGet(colName string, rowIndex int) (val uint8) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetUint8(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get *Table table cell from colName at rowIndex

	Like its non-MustGet alternative GetTable(), but panics on error, and does not return an error.
*/
func (table *Table) GetTableMustGet(colName string, rowIndex int) (val *Table) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetTable(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Get time.Time table cell from colName at rowIndex

	Like its non-MustGet alternative GetTime(), but panics on error, and does not return an error.
*/
func (table *Table) GetTimeMustGet(colName string, rowIndex int) (val time.Time) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetTime(colName, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*	Set []byte MustSet table cell from colName

	Like its non-MustSet alternative SetByteSlice(), but panics on error, and does not return an error.
*/
func (table *Table) SetByteSliceMustSet(colName string, rowIndex int, val []byte) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetByteSlice(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set []uint8 MustSet table cell from colName

	Like its non-MustSet alternative SetUint8Slice(), but panics on error, and does not return an error.
*/
func (table *Table) SetUint8SliceMustSet(colName string, rowIndex int, val []uint8) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetUint8Slice(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set bool MustSet table cell from colName

	Like its non-MustSet alternative SetBool(), but panics on error, and does not return an error.
*/
func (table *Table) SetBoolMustSet(colName string, rowIndex int, val bool) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetBool(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set byte MustSet table cell from colName

	Like its non-MustSet alternative SetByte(), but panics on error, and does not return an error.
*/
func (table *Table) SetByteMustSet(colName string, rowIndex int, val byte) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetByte(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set float32 MustSet table cell from colName

	Like its non-MustSet alternative SetFloat32(), but panics on error, and does not return an error.
*/
func (table *Table) SetFloat32MustSet(colName string, rowIndex int, val float32) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetFloat32(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set float64 MustSet table cell from colName

	Like its non-MustSet alternative SetFloat64(), but panics on error, and does not return an error.
*/
func (table *Table) SetFloat64MustSet(colName string, rowIndex int, val float64) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetFloat64(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set int MustSet table cell from colName

	Like its non-MustSet alternative SetInt(), but panics on error, and does not return an error.
*/
func (table *Table) SetIntMustSet(colName string, rowIndex int, val int) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetInt(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set int16 MustSet table cell from colName

	Like its non-MustSet alternative SetInt16(), but panics on error, and does not return an error.
*/
func (table *Table) SetInt16MustSet(colName string, rowIndex int, val int16) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetInt16(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set int32 MustSet table cell from colName

	Like its non-MustSet alternative SetInt32(), but panics on error, and does not return an error.
*/
func (table *Table) SetInt32MustSet(colName string, rowIndex int, val int32) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetInt32(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set int64 MustSet table cell from colName

	Like its non-MustSet alternative SetInt64(), but panics on error, and does not return an error.
*/
func (table *Table) SetInt64MustSet(colName string, rowIndex int, val int64) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetInt64(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set int8 MustSet table cell from colName

	Like its non-MustSet alternative SetInt8(), but panics on error, and does not return an error.
*/
func (table *Table) SetInt8MustSet(colName string, rowIndex int, val int8) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetInt8(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set rune MustSet table cell from colName

	Like its non-MustSet alternative SetRune(), but panics on error, and does not return an error.
*/
func (table *Table) SetRuneMustSet(colName string, rowIndex int, val rune) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetRune(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set string MustSet table cell from colName

	Like its non-MustSet alternative SetString(), but panics on error, and does not return an error.
*/
func (table *Table) SetStringMustSet(colName string, rowIndex int, val string) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetString(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set uint MustSet table cell from colName

	Like its non-MustSet alternative SetUint(), but panics on error, and does not return an error.
*/
func (table *Table) SetUintMustSet(colName string, rowIndex int, val uint) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetUint(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set uint16 MustSet table cell from colName

	Like its non-MustSet alternative SetUint16(), but panics on error, and does not return an error.
*/
func (table *Table) SetUint16MustSet(colName string, rowIndex int, val uint16) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetUint16(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set uint32 MustSet table cell from colName

	Like its non-MustSet alternative SetUint32(), but panics on error, and does not return an error.
*/
func (table *Table) SetUint32MustSet(colName string, rowIndex int, val uint32) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetUint32(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set uint64 MustSet table cell from colName

	Like its non-MustSet alternative SetUint64(), but panics on error, and does not return an error.
*/
func (table *Table) SetUint64MustSet(colName string, rowIndex int, val uint64) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetUint64(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set uint8 MustSet table cell from colName

	Like its non-MustSet alternative SetUint8(), but panics on error, and does not return an error.
*/
func (table *Table) SetUint8MustSet(colName string, rowIndex int, val uint8) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetUint8(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set *Table MustSet table cell from colName

	Like its non-MustSet alternative SetTable(), but panics on error, and does not return an error.
*/
func (table *Table) SetTableMustSet(colName string, rowIndex int, val *Table) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetTable(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set time.Time MustSet table cell from colName

	Like its non-MustSet alternative SetTime(), but panics on error, and does not return an error.
*/
func (table *Table) SetTimeMustSet(colName string, rowIndex int, val time.Time) {

	if table == nil {
		panic(fmt.Errorf("table.%s(%s, %d, val): table is <nil>", UtilFuncNameNoParens(), colName, rowIndex))
	}

	err := table.SetTime(colName, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set []byte MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetByteSliceByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetByteSliceByColIndexMustSet(colIndex int, rowIndex int, val []byte) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetByteSliceByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set []uint8 MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetUint8SliceByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetUint8SliceByColIndexMustSet(colIndex int, rowIndex int, val []uint8) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetUint8SliceByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set bool MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetBoolByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetBoolByColIndexMustSet(colIndex int, rowIndex int, val bool) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetBoolByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set byte MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetByteByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetByteByColIndexMustSet(colIndex int, rowIndex int, val byte) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetByteByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set float32 MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetFloat32ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetFloat32ByColIndexMustSet(colIndex int, rowIndex int, val float32) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetFloat32ByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set float64 MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetFloat64ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetFloat64ByColIndexMustSet(colIndex int, rowIndex int, val float64) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetFloat64ByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set int MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetIntByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetIntByColIndexMustSet(colIndex int, rowIndex int, val int) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetIntByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set int16 MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetInt16ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetInt16ByColIndexMustSet(colIndex int, rowIndex int, val int16) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetInt16ByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set int32 MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetInt32ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetInt32ByColIndexMustSet(colIndex int, rowIndex int, val int32) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetInt32ByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set int64 MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetInt64ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetInt64ByColIndexMustSet(colIndex int, rowIndex int, val int64) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetInt64ByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set int8 MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetInt8ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetInt8ByColIndexMustSet(colIndex int, rowIndex int, val int8) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetInt8ByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set rune MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetRuneByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetRuneByColIndexMustSet(colIndex int, rowIndex int, val rune) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetRuneByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set string MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetStringByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetStringByColIndexMustSet(colIndex int, rowIndex int, val string) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetStringByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set uint MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetUintByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetUintByColIndexMustSet(colIndex int, rowIndex int, val uint) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetUintByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set uint16 MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetUint16ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetUint16ByColIndexMustSet(colIndex int, rowIndex int, val uint16) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetUint16ByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set uint32 MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetUint32ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetUint32ByColIndexMustSet(colIndex int, rowIndex int, val uint32) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetUint32ByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set uint64 MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetUint64ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetUint64ByColIndexMustSet(colIndex int, rowIndex int, val uint64) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetUint64ByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set uint8 MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetUint8ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetUint8ByColIndexMustSet(colIndex int, rowIndex int, val uint8) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetUint8ByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set *Table MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetTableByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetTableByColIndexMustSet(colIndex int, rowIndex int, val *Table) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetTableByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

/*	Set time.Time MustSet table cell from colName at rowIndex

	Like its non-MustSet alternative SetTimeByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) SetTimeByColIndexMustSet(colIndex int, rowIndex int, val time.Time) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	err := table.SetTimeByColIndex(colIndex, rowIndex, val)
	if err != nil {
		panic(err)
	}
}

//	----------------------------------------------------------------------------
//	next group: Get<type>ByColIndex() functions for each of 20 types
//	20 types: *Table []byte []uint8 bool byte float32 float64 int int16 int32 int64 int8 rune string time.Time uint uint16 uint32 uint64 uint8
//  NOTE: Types are defined in helpersmain.go AND parser.go
//	----------------------------------------------------------------------------

//  Get []byte table cell from colIndex at rowIndex
func (table *Table) GetByteSliceByColIndex(colIndex int, rowIndex int) (val []byte, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "[]byte"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].([]byte)

	return
}

//  Get []uint8 table cell from colIndex at rowIndex
func (table *Table) GetUint8SliceByColIndex(colIndex int, rowIndex int) (val []uint8, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "[]uint8"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].([]uint8)

	return
}

//  Get bool table cell from colIndex at rowIndex
func (table *Table) GetBoolByColIndex(colIndex int, rowIndex int) (val bool, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "bool"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(bool)

	return
}

//  Get byte table cell from colIndex at rowIndex
func (table *Table) GetByteByColIndex(colIndex int, rowIndex int) (val byte, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "byte"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(byte)

	return
}

//  Get float32 table cell from colIndex at rowIndex
func (table *Table) GetFloat32ByColIndex(colIndex int, rowIndex int) (val float32, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "float32"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(float32)

	return
}

//  Get float64 table cell from colIndex at rowIndex
func (table *Table) GetFloat64ByColIndex(colIndex int, rowIndex int) (val float64, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "float64"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(float64)

	return
}

//  Get int table cell from colIndex at rowIndex
func (table *Table) GetIntByColIndex(colIndex int, rowIndex int) (val int, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "int"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(int)

	return
}

//  Get int16 table cell from colIndex at rowIndex
func (table *Table) GetInt16ByColIndex(colIndex int, rowIndex int) (val int16, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "int16"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(int16)

	return
}

//  Get int32 table cell from colIndex at rowIndex
func (table *Table) GetInt32ByColIndex(colIndex int, rowIndex int) (val int32, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "int32"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(int32)

	return
}

//  Get int64 table cell from colIndex at rowIndex
func (table *Table) GetInt64ByColIndex(colIndex int, rowIndex int) (val int64, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "int64"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(int64)

	return
}

//  Get int8 table cell from colIndex at rowIndex
func (table *Table) GetInt8ByColIndex(colIndex int, rowIndex int) (val int8, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "int8"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(int8)

	return
}

//  Get rune table cell from colIndex at rowIndex
func (table *Table) GetRuneByColIndex(colIndex int, rowIndex int) (val rune, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "rune"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(rune)

	return
}

//  Get string table cell from colIndex at rowIndex
func (table *Table) GetStringByColIndex(colIndex int, rowIndex int) (val string, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "string"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(string)

	return
}

//  Get uint table cell from colIndex at rowIndex
func (table *Table) GetUintByColIndex(colIndex int, rowIndex int) (val uint, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "uint"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(uint)

	return
}

//  Get uint16 table cell from colIndex at rowIndex
func (table *Table) GetUint16ByColIndex(colIndex int, rowIndex int) (val uint16, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "uint16"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(uint16)

	return
}

//  Get uint32 table cell from colIndex at rowIndex
func (table *Table) GetUint32ByColIndex(colIndex int, rowIndex int) (val uint32, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "uint32"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(uint32)

	return
}

//  Get uint64 table cell from colIndex at rowIndex
func (table *Table) GetUint64ByColIndex(colIndex int, rowIndex int) (val uint64, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "uint64"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(uint64)

	return
}

//  Get uint8 table cell from colIndex at rowIndex
func (table *Table) GetUint8ByColIndex(colIndex int, rowIndex int) (val uint8, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "uint8"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(uint8)

	return
}

//  Get *Table table cell from colIndex at rowIndex
func (table *Table) GetTableByColIndex(colIndex int, rowIndex int) (val *Table, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "*Table"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(*Table)

	return
}

//  Get time.Time table cell from colIndex at rowIndex
func (table *Table) GetTimeByColIndex(colIndex int, rowIndex int) (val time.Time, err error) {

	// See: Get<type>ByColIndex() functions

	if table == nil {
		err = fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
		return
	}

	const valType string = "time.Time"

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				UtilFuncName(), table.Name(), colIndex, colType)
		}
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex].(time.Time)

	return
}

/*  Get []byte table cell from colIndex at rowIndex

Like its non-MustGet alternative GetByteSliceByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetByteSliceByColIndexMustGet(colIndex int, rowIndex int) (val []byte) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetByteSliceByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get []uint8 table cell from colIndex at rowIndex

Like its non-MustGet alternative GetUint8SliceByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetUint8SliceByColIndexMustGet(colIndex int, rowIndex int) (val []uint8) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetUint8SliceByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get bool table cell from colIndex at rowIndex

Like its non-MustGet alternative GetBoolByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetBoolByColIndexMustGet(colIndex int, rowIndex int) (val bool) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetBoolByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get byte table cell from colIndex at rowIndex

Like its non-MustGet alternative GetByteByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetByteByColIndexMustGet(colIndex int, rowIndex int) (val byte) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetByteByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get float32 table cell from colIndex at rowIndex

Like its non-MustGet alternative GetFloat32ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetFloat32ByColIndexMustGet(colIndex int, rowIndex int) (val float32) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetFloat32ByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get float64 table cell from colIndex at rowIndex

Like its non-MustGet alternative GetFloat64ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetFloat64ByColIndexMustGet(colIndex int, rowIndex int) (val float64) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetFloat64ByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get int table cell from colIndex at rowIndex

Like its non-MustGet alternative GetIntByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetIntByColIndexMustGet(colIndex int, rowIndex int) (val int) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetIntByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get int16 table cell from colIndex at rowIndex

Like its non-MustGet alternative GetInt16ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetInt16ByColIndexMustGet(colIndex int, rowIndex int) (val int16) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetInt16ByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get int32 table cell from colIndex at rowIndex

Like its non-MustGet alternative GetInt32ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetInt32ByColIndexMustGet(colIndex int, rowIndex int) (val int32) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetInt32ByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get int64 table cell from colIndex at rowIndex

Like its non-MustGet alternative GetInt64ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetInt64ByColIndexMustGet(colIndex int, rowIndex int) (val int64) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetInt64ByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get int8 table cell from colIndex at rowIndex

Like its non-MustGet alternative GetInt8ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetInt8ByColIndexMustGet(colIndex int, rowIndex int) (val int8) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetInt8ByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get rune table cell from colIndex at rowIndex

Like its non-MustGet alternative GetRuneByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetRuneByColIndexMustGet(colIndex int, rowIndex int) (val rune) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetRuneByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get string table cell from colIndex at rowIndex

Like its non-MustGet alternative GetStringByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetStringByColIndexMustGet(colIndex int, rowIndex int) (val string) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetStringByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get uint table cell from colIndex at rowIndex

Like its non-MustGet alternative GetUintByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetUintByColIndexMustGet(colIndex int, rowIndex int) (val uint) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetUintByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get uint16 table cell from colIndex at rowIndex

Like its non-MustGet alternative GetUint16ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetUint16ByColIndexMustGet(colIndex int, rowIndex int) (val uint16) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetUint16ByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get uint32 table cell from colIndex at rowIndex

Like its non-MustGet alternative GetUint32ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetUint32ByColIndexMustGet(colIndex int, rowIndex int) (val uint32) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetUint32ByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get uint64 table cell from colIndex at rowIndex

Like its non-MustGet alternative GetUint64ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetUint64ByColIndexMustGet(colIndex int, rowIndex int) (val uint64) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetUint64ByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get uint8 table cell from colIndex at rowIndex

Like its non-MustGet alternative GetUint8ByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetUint8ByColIndexMustGet(colIndex int, rowIndex int) (val uint8) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetUint8ByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get *Table table cell from colIndex at rowIndex

Like its non-MustGet alternative GetTableByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetTableByColIndexMustGet(colIndex int, rowIndex int) (val *Table) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetTableByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*  Get time.Time table cell from colIndex at rowIndex

Like its non-MustGet alternative GetTimeByColIndex(), but panics on error, and does not return an error.
*/
func (table *Table) GetTimeByColIndexMustGet(colIndex int, rowIndex int) (val time.Time) {

	if table == nil {
		panic(fmt.Errorf("table.%s: table is <nil>", UtilFuncName()))
	}

	val, err := table.GetTimeByColIndex(colIndex, rowIndex)
	if err != nil {
		panic(err)
	}

	return val
}

/*
//func (table *Table) setCellToZeroValueByColIndexCheck(colIndex int, rowIndex int) error {
//// This is the MUCH SLOWER previous version. Is there any safety advantage in using it? Perhaps not.
//	// TODO: Test for colIndex or rowIndex out of range? Or is this done by underlying functions?
//
//	if table == nil { return fmt.Errorf("table.%s: table is <nil>", UtilFuncName()) }
//
//	var err error
//	var colType string
//
//	colType, err = table.ColTypeByColIndex(colIndex)
//	if err != nil {
//		return err
//	}
//
//	switch colType {
//
//		case "[]byte":
//			err = table.SetByteSliceByColIndex(colIndex, rowIndex, []byte{})
//
//		case "[]uint8":
//			err = table.SetUint8SliceByColIndex(colIndex, rowIndex, []uint8{})
//
//		case "bool":
//			err = table.SetBoolByColIndex(colIndex, rowIndex, false)
//
//		case "byte":
//			err = table.SetByteByColIndex(colIndex, rowIndex, 0)
//
//		case "float32":
//			err = table.SetFloat32ByColIndex(colIndex, rowIndex, 0.0)
//
//		case "float64":
//			err = table.SetFloat64ByColIndex(colIndex, rowIndex, 0.0)
//
//		case "int":
//			err = table.SetIntByColIndex(colIndex, rowIndex, 0)
//
//		case "int16":
//			err = table.SetInt16ByColIndex(colIndex, rowIndex, 0)
//
//		case "int32":
//			err = table.SetInt32ByColIndex(colIndex, rowIndex, 0)
//
//		case "int64":
//			err = table.SetInt64ByColIndex(colIndex, rowIndex, 0)
//
//		case "int8":
//			err = table.SetInt8ByColIndex(colIndex, rowIndex, 0)
//
//		case "rune":
//			err = table.SetRuneByColIndex(colIndex, rowIndex, 0)
//
//		case "string":
//			err = table.SetStringByColIndex(colIndex, rowIndex, "")
//
//		case "uint":
//			err = table.SetUintByColIndex(colIndex, rowIndex, 0)
//
//		case "uint16":
//			err = table.SetUint16ByColIndex(colIndex, rowIndex, 0)
//
//		case "uint32":
//			err = table.SetUint32ByColIndex(colIndex, rowIndex, 0)
//
//		case "uint64":
//			err = table.SetUint64ByColIndex(colIndex, rowIndex, 0)
//
//		case "uint8":
//			err = table.SetUint8ByColIndex(colIndex, rowIndex, 0)
//
//		case "*Table":
//			err = table.SetTableByColIndex(colIndex, rowIndex, NewNilTable())
//
//		case "time.Time":
//			err = table.SetTimeByColIndex(colIndex, rowIndex, MinTime)
//
//		default:
//			msg := fmt.Sprintf("invalid type: %s (Valid types:", colType)
//			// Note: Because maps are not ordered, this (desirably) shuffles the order of valid col types with each call.
//			for typeName, _ := range globalColTypesMap {
//				msg += fmt.Sprintf(" %s", typeName)
//			}
//			msg += ")"
//			err = errors.New(msg)
//			return err
//	}
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
*/

type zeroVals struct {
	byteSliceVal  []byte
	uint8SliceVal []uint8
	boolVal       bool
	byteVal       byte
	float32Val    float32
	float64Val    float64
	intVal        int
	int16Val      int16
	int32Val      int32
	int64Val      int64
	int8Val       int8
	runeVal       rune
	stringVal     string
	uintVal       uint
	uint16Val     uint16
	uint32Val     uint32
	uint64Val     uint64
	uint8Val      uint8
	tableVal      *Table
	timeVal       time.Time
}

var zeroVal zeroVals

func init() {
	// This avoids relatively expensive assignments to a local variable in SetCellToZeroValueByColIndex()

	zeroVal.byteSliceVal = []byte{}

	zeroVal.uint8SliceVal = []uint8{}

	zeroVal.boolVal = false

	zeroVal.byteVal = 0

	zeroVal.float32Val = 0.0

	zeroVal.float64Val = 0.0

	zeroVal.intVal = 0

	zeroVal.int16Val = 0

	zeroVal.int32Val = 0

	zeroVal.int64Val = 0

	zeroVal.int8Val = 0

	zeroVal.runeVal = 0

	zeroVal.stringVal = ""

	zeroVal.uintVal = 0

	zeroVal.uint16Val = 0

	zeroVal.uint32Val = 0

	zeroVal.uint64Val = 0

	zeroVal.uint8Val = 0

	zeroVal.tableVal = NewNilTable() // Beware: to avoid circular reference, this can be used just once.

	zeroVal.timeVal = MinTime
}

func (table *Table) SetCellToZeroValueByColIndex(colIndex int, rowIndex int) error {

	// Note: zeroVal is defined in the type zeroVals struct.

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	var colType = table.colTypes[colIndex]

	switch colType {
	case "[]byte":
		// This is a x10 tuning strategy to avoid type conversion []byte([]byte{})
		table.rows[rowIndex][colIndex] = zeroVal.byteSliceVal
	case "[]uint8":
		// This is a x10 tuning strategy to avoid type conversion []uint8([]uint8{})
		table.rows[rowIndex][colIndex] = zeroVal.uint8SliceVal
	case "bool":
		// This is a x10 tuning strategy to avoid type conversion bool(false)
		table.rows[rowIndex][colIndex] = zeroVal.boolVal
	case "byte":
		// This is a x10 tuning strategy to avoid type conversion byte(0)
		table.rows[rowIndex][colIndex] = zeroVal.byteVal
	case "float32":
		// This is a x10 tuning strategy to avoid type conversion float32(0.0)
		table.rows[rowIndex][colIndex] = zeroVal.float32Val
	case "float64":
		// This is a x10 tuning strategy to avoid type conversion float64(0.0)
		table.rows[rowIndex][colIndex] = zeroVal.float64Val
	case "int":
		// This is a x10 tuning strategy to avoid type conversion int(0)
		table.rows[rowIndex][colIndex] = zeroVal.intVal
	case "int16":
		// This is a x10 tuning strategy to avoid type conversion int16(0)
		table.rows[rowIndex][colIndex] = zeroVal.int16Val
	case "int32":
		// This is a x10 tuning strategy to avoid type conversion int32(0)
		table.rows[rowIndex][colIndex] = zeroVal.int32Val
	case "int64":
		// This is a x10 tuning strategy to avoid type conversion int64(0)
		table.rows[rowIndex][colIndex] = zeroVal.int64Val
	case "int8":
		// This is a x10 tuning strategy to avoid type conversion int8(0)
		table.rows[rowIndex][colIndex] = zeroVal.int8Val
	case "rune":
		// This is a x10 tuning strategy to avoid type conversion rune(0)
		table.rows[rowIndex][colIndex] = zeroVal.runeVal
	case "string":
		// This is a x10 tuning strategy to avoid type conversion string("")
		table.rows[rowIndex][colIndex] = zeroVal.stringVal
	case "uint":
		// This is a x10 tuning strategy to avoid type conversion uint(0)
		table.rows[rowIndex][colIndex] = zeroVal.uintVal
	case "uint16":
		// This is a x10 tuning strategy to avoid type conversion uint16(0)
		table.rows[rowIndex][colIndex] = zeroVal.uint16Val
	case "uint32":
		// This is a x10 tuning strategy to avoid type conversion uint32(0)
		table.rows[rowIndex][colIndex] = zeroVal.uint32Val
	case "uint64":
		// This is a x10 tuning strategy to avoid type conversion uint64(0)
		table.rows[rowIndex][colIndex] = zeroVal.uint64Val
	case "uint8":
		// This is a x10 tuning strategy to avoid type conversion uint8(0)
		table.rows[rowIndex][colIndex] = zeroVal.uint8Val
	case "*Table":
		// This is a x10 tuning strategy to avoid type conversion *Table(NewNilTable())
		table.rows[rowIndex][colIndex] = NewNilTable() // Avoid circular reference.
	case "time.Time":
		// This is a x10 tuning strategy to avoid type conversion time.Time(MinTime)
		table.rows[rowIndex][colIndex] = zeroVal.timeVal
	default:
		return fmt.Errorf("invalid type: %s", colType)
	}

	return nil
}

func (table *Table) SetRowCellsToZeroValue(rowIndex int) error {

	// Note: zeroVal is defined in the type zeroVals struct.

	if table == nil {
		return fmt.Errorf("table.%s: table is <nil>", UtilFuncName())
	}

	if table.isNilTable {
		return fmt.Errorf("table.%s: table is an unnamed NilTable. Call table.SetName() to un-Nil it", UtilFuncName())
	}

	for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
		var colType string = table.colTypes[colIndex]
		switch colType {
		case "[]byte":
			// This is a x10 tuning strategy to avoid type conversion []byte([]byte{})
			table.rows[rowIndex][colIndex] = zeroVal.byteSliceVal
		case "[]uint8":
			// This is a x10 tuning strategy to avoid type conversion []uint8([]uint8{})
			table.rows[rowIndex][colIndex] = zeroVal.uint8SliceVal
		case "bool":
			// This is a x10 tuning strategy to avoid type conversion bool(false)
			table.rows[rowIndex][colIndex] = zeroVal.boolVal
		case "byte":
			// This is a x10 tuning strategy to avoid type conversion byte(0)
			table.rows[rowIndex][colIndex] = zeroVal.byteVal
		case "float32":
			// This is a x10 tuning strategy to avoid type conversion float32(0.0)
			table.rows[rowIndex][colIndex] = zeroVal.float32Val
		case "float64":
			// This is a x10 tuning strategy to avoid type conversion float64(0.0)
			table.rows[rowIndex][colIndex] = zeroVal.float64Val
		case "int":
			// This is a x10 tuning strategy to avoid type conversion int(0)
			table.rows[rowIndex][colIndex] = zeroVal.intVal
		case "int16":
			// This is a x10 tuning strategy to avoid type conversion int16(0)
			table.rows[rowIndex][colIndex] = zeroVal.int16Val
		case "int32":
			// This is a x10 tuning strategy to avoid type conversion int32(0)
			table.rows[rowIndex][colIndex] = zeroVal.int32Val
		case "int64":
			// This is a x10 tuning strategy to avoid type conversion int64(0)
			table.rows[rowIndex][colIndex] = zeroVal.int64Val
		case "int8":
			// This is a x10 tuning strategy to avoid type conversion int8(0)
			table.rows[rowIndex][colIndex] = zeroVal.int8Val
		case "rune":
			// This is a x10 tuning strategy to avoid type conversion rune(0)
			table.rows[rowIndex][colIndex] = zeroVal.runeVal
		case "string":
			// This is a x10 tuning strategy to avoid type conversion string("")
			table.rows[rowIndex][colIndex] = zeroVal.stringVal
		case "uint":
			// This is a x10 tuning strategy to avoid type conversion uint(0)
			table.rows[rowIndex][colIndex] = zeroVal.uintVal
		case "uint16":
			// This is a x10 tuning strategy to avoid type conversion uint16(0)
			table.rows[rowIndex][colIndex] = zeroVal.uint16Val
		case "uint32":
			// This is a x10 tuning strategy to avoid type conversion uint32(0)
			table.rows[rowIndex][colIndex] = zeroVal.uint32Val
		case "uint64":
			// This is a x10 tuning strategy to avoid type conversion uint64(0)
			table.rows[rowIndex][colIndex] = zeroVal.uint64Val
		case "uint8":
			// This is a x10 tuning strategy to avoid type conversion uint8(0)
			table.rows[rowIndex][colIndex] = zeroVal.uint8Val
		case "*Table":
			// This is a x10 tuning strategy to avoid type conversion *Table(NewNilTable())
			table.rows[rowIndex][colIndex] = NewNilTable() // Avoid circular reference.
		case "time.Time":
			// This is a x10 tuning strategy to avoid type conversion time.Time(MinTime)
			table.rows[rowIndex][colIndex] = zeroVal.timeVal
		default:
			return fmt.Errorf("invalid type: %s", colType)
		}
	}

	return nil
}
