package gotables

/*
	helpers.go
*/

import (
	"bytes"
	//	"fmt"
	//	"os"
	//	"runtime/debug"
	"testing"
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

// Note: time.Type has been eliminated from helpers.test.go due to great difficulty automating it.

//	-----------------------------------------------------------------------
//	next group: TestSet<type>() TestGet<type>() functions for each of 20 types.
//	-----------------------------------------------------------------------

//	Test Set and Get table cell in colName at rowIndex to newValue []byte
func TestSetAndGetByteSlice(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "ByteSliceValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "[]byte"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected []byte
	}{
		{[]byte{0}},
		{[]byte{255}},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetByteSlice(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetByteSlice(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(value, test.expected) {
			t.Fatalf("expecting GetByteSlice() bytes %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue []uint8
func TestSetAndGetUint8Slice(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint8SliceValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "[]uint8"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected []uint8
	}{
		{[]uint8{0}},
		{[]uint8{255}},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint8Slice(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetUint8Slice(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(value, test.expected) {
			t.Fatalf("expecting GetUint8Slice() bytes %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue bool
func TestSetAndGetBool(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "BoolValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "bool"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected bool
	}{
		{false},
		{true},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetBool(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetBool(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetBool() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue byte
func TestSetAndGetByte(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "ByteValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "byte"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected byte
	}{
		{0},
		{255},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetByte(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetByte(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetByte() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue float32
func TestSetAndGetFloat32(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Float32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "float32"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected float32
	}{
		{1.401298464324817e-45},
		{3.4028234663852886e+38},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetFloat32(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetFloat32(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetFloat32() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue float64
func TestSetAndGetFloat64(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Float64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "float64"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected float64
	}{
		{5e-324},
		{1.7976931348623157e+308},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetFloat64(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetFloat64(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetFloat64() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int
func TestSetAndGetInt(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "IntValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "int"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected int
	}{
		{-9223372036854775808},
		{9223372036854775807},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetInt(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetInt(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetInt() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int16
func TestSetAndGetInt16(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int16Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "int16"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected int16
	}{
		{-32768},
		{32767},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetInt16(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetInt16(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetInt16() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int32
func TestSetAndGetInt32(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "int32"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected int32
	}{
		{-2147483648},
		{2147483647},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetInt32(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetInt32(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetInt32() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int64
func TestSetAndGetInt64(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "int64"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected int64
	}{
		{-9223372036854775808},
		{9223372036854775807},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetInt64(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetInt64(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetInt64() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int8
func TestSetAndGetInt8(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int8Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "int8"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected int8
	}{
		{-128},
		{127},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetInt8(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetInt8(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetInt8() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue rune
func TestSetAndGetRune(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "RuneValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "rune"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected rune
	}{
		{'A'},
		{'Z'},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetRune(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetRune(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetRune() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue string
func TestSetAndGetString(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "StringValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "string"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected string
	}{
		{"ABC"},
		{"abc"},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetString(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetString(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetString() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint
func TestSetAndGetUint(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "UintValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "uint"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected uint
	}{
		{0},
		{18446744073709551615},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetUint(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetUint() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint16
func TestSetAndGetUint16(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint16Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "uint16"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected uint16
	}{
		{0},
		{65535},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint16(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetUint16(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetUint16() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint32
func TestSetAndGetUint32(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "uint32"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected uint32
	}{
		{0},
		{4294967295},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint32(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetUint32(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetUint32() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint64
func TestSetAndGetUint64(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "uint64"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected uint64
	}{
		{0},
		{18446744073709551615},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint64(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetUint64(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetUint64() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint8
func TestSetAndGetUint8(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint8Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "uint8"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected uint8
	}{
		{0},
		{255},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint8(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetUint8(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetUint8() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue *Table
func TestSetAndGetTable(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "TableValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "*Table"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected *Table
	}{
		{NewNilTable()},
		{NewNilTable()},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetTable(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetTable(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetTable() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue time.Time
func TestSetAndGetTime(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "TimeValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	var colType string = "time.Time"
	err = table.AppendCol(colName, colType)
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected time.Time
	}{
		{MinTime},
		{MaxTime},
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetTime(colName, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetTime(colName, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetTime() value %v, not %v", test.expected, value)
		}
	}
}

//	--------------------------------------------------------------------------------
//	next group: TestSet<type>ByColIndex() TestGet<type>ByColIndex() functions for each of 20 types.
//	--------------------------------------------------------------------------------

//	Test Set and Get table cell in colIndex at rowIndex to newValue []byte
func TestHelperSetAndGetByteSliceByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "ByteSliceValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "[]byte")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected []byte
	}{
		{[]byte{0}},
		{[]byte{255}},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetByteSliceByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetByteSliceByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(value, test.expected) {
			t.Fatalf("expecting GetByteSliceByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue []uint8
func TestHelperSetAndGetUint8SliceByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint8SliceValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "[]uint8")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected []uint8
	}{
		{[]uint8{0}},
		{[]uint8{255}},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUint8SliceByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetUint8SliceByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(value, test.expected) {
			t.Fatalf("expecting GetUint8SliceByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue bool
func TestHelperSetAndGetBoolByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "BoolValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "bool")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected bool
	}{
		{false},
		{true},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetBoolByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetBoolByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetBoolByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue byte
func TestHelperSetAndGetByteByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "ByteValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "byte")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected byte
	}{
		{0},
		{255},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetByteByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetByteByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetByteByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue float32
func TestHelperSetAndGetFloat32ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Float32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "float32")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected float32
	}{
		{1.401298464324817e-45},
		{3.4028234663852886e+38},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetFloat32ByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetFloat32ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetFloat32ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue float64
func TestHelperSetAndGetFloat64ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Float64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "float64")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected float64
	}{
		{5e-324},
		{1.7976931348623157e+308},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetFloat64ByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetFloat64ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetFloat64ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int
func TestHelperSetAndGetIntByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "IntValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "int")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected int
	}{
		{-9223372036854775808},
		{9223372036854775807},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetIntByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetIntByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetIntByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int16
func TestHelperSetAndGetInt16ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int16Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "int16")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected int16
	}{
		{-32768},
		{32767},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetInt16ByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetInt16ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetInt16ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int32
func TestHelperSetAndGetInt32ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "int32")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected int32
	}{
		{-2147483648},
		{2147483647},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetInt32ByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetInt32ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetInt32ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int64
func TestHelperSetAndGetInt64ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "int64")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected int64
	}{
		{-9223372036854775808},
		{9223372036854775807},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetInt64ByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetInt64ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetInt64ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int8
func TestHelperSetAndGetInt8ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int8Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "int8")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected int8
	}{
		{-128},
		{127},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetInt8ByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetInt8ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetInt8ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue rune
func TestHelperSetAndGetRuneByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "RuneValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "rune")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected rune
	}{
		{'A'},
		{'Z'},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetRuneByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetRuneByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetRuneByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue string
func TestHelperSetAndGetStringByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "StringValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "string")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected string
	}{
		{"ABC"},
		{"abc"},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetStringByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetStringByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetStringByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint
func TestHelperSetAndGetUintByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "UintValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "uint")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected uint
	}{
		{0},
		{18446744073709551615},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUintByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetUintByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetUintByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint16
func TestHelperSetAndGetUint16ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint16Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "uint16")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected uint16
	}{
		{0},
		{65535},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUint16ByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetUint16ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetUint16ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint32
func TestHelperSetAndGetUint32ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "uint32")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected uint32
	}{
		{0},
		{4294967295},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUint32ByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetUint32ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetUint32ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint64
func TestHelperSetAndGetUint64ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "uint64")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected uint64
	}{
		{0},
		{18446744073709551615},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUint64ByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetUint64ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetUint64ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint8
func TestHelperSetAndGetUint8ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint8Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "uint8")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected uint8
	}{
		{0},
		{255},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUint8ByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetUint8ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetUint8ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue *Table
func TestHelperSetAndGetTableByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "TableValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "*Table")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected *Table
	}{
		{NewNilTable()},
		{NewNilTable()},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetTableByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetTableByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetTableByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue time.Time
func TestHelperSetAndGetTimeByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "TimeValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendCol(colName, "time.Time")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		expected time.Time
	}{
		{MinTime},
		{MaxTime},
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetTimeByColIndex(colIndex, rowIndex, test.expected)
		if err != nil {
			t.Fatal(err)
		}

		value, err := table.GetTimeByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if value != test.expected {
			t.Fatalf("expecting GetTimeByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	-----------------------------------------------------------------------
//	bench test
//	next group: BenchMarkHelpersSet<type>() functions for each of 20 types.
//	-----------------------------------------------------------------------

//	Test Set and Get table cell in colName at rowIndex to newValue []byte
func BenchmarkHelperSetByteSlice(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "ByteSliceValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "[]byte")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected []byte
	}{
		{[]byte{0}},
		{[]byte{255}},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetByteSlice(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue []uint8
func BenchmarkHelperSetUint8Slice(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Uint8SliceValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "[]uint8")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected []uint8
	}{
		{[]uint8{0}},
		{[]uint8{255}},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint8Slice(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue bool
func BenchmarkHelperSetBool(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "BoolValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "bool")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected bool
	}{
		{false},
		{true},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetBool(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue byte
func BenchmarkHelperSetByte(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "ByteValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "byte")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected byte
	}{
		{0},
		{255},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetByte(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue float32
func BenchmarkHelperSetFloat32(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Float32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "float32")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected float32
	}{
		{1.401298464324817e-45},
		{3.4028234663852886e+38},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetFloat32(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue float64
func BenchmarkHelperSetFloat64(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Float64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "float64")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected float64
	}{
		{5e-324},
		{1.7976931348623157e+308},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetFloat64(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int
func BenchmarkHelperSetInt(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "IntValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int
	}{
		{-9223372036854775808},
		{9223372036854775807},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetInt(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int16
func BenchmarkHelperSetInt16(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Int16Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int16")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int16
	}{
		{-32768},
		{32767},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetInt16(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int32
func BenchmarkHelperSetInt32(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Int32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int32")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int32
	}{
		{-2147483648},
		{2147483647},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetInt32(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int64
func BenchmarkHelperSetInt64(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Int64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int64")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int64
	}{
		{-9223372036854775808},
		{9223372036854775807},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetInt64(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int8
func BenchmarkHelperSetInt8(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Int8Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int8")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int8
	}{
		{-128},
		{127},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetInt8(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue rune
func BenchmarkHelperSetRune(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "RuneValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "rune")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected rune
	}{
		{'A'},
		{'Z'},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetRune(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue string
func BenchmarkHelperSetString(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "StringValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "string")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected string
	}{
		{"ABC"},
		{"abc"},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetString(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint
func BenchmarkHelperSetUint(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "UintValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint
	}{
		{0},
		{18446744073709551615},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint16
func BenchmarkHelperSetUint16(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Uint16Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint16")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint16
	}{
		{0},
		{65535},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint16(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint32
func BenchmarkHelperSetUint32(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Uint32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint32")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint32
	}{
		{0},
		{4294967295},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint32(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint64
func BenchmarkHelperSetUint64(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Uint64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint64")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint64
	}{
		{0},
		{18446744073709551615},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint64(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint8
func BenchmarkHelperSetUint8(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Uint8Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint8")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint8
	}{
		{0},
		{255},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint8(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue *Table
func BenchmarkHelperSetTable(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "TableValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "*Table")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected *Table
	}{
		{NewNilTable()},
		{NewNilTable()},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetTable(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue time.Time
func BenchmarkHelperSetTime(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "TimeValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "time.Time")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected time.Time
	}{
		{MinTime},
		{MaxTime},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetTime(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	--------------------------------------------------------------------------------
//	bench test
//	next group: BenchmarkHelperSetAndGet<type>ByColIndex() functions for each of 20 types.
//	--------------------------------------------------------------------------------

//	Test Set and Get table cell in colIndex at rowIndex to newValue []byte
func BenchmarkHelperSetAndGetByteSliceByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "ByteSliceValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "[]byte")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected []byte
	}{
		{[]byte{0}},
		{[]byte{255}},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetByteSliceByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetByteSliceByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if !bytes.Equal(value, test.expected) {
				b.Errorf("expecting GetByteSliceByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue []uint8
func BenchmarkHelperSetAndGetUint8SliceByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Uint8SliceValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "[]uint8")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected []uint8
	}{
		{[]uint8{0}},
		{[]uint8{255}},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUint8SliceByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetUint8SliceByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if !bytes.Equal(value, test.expected) {
				b.Errorf("expecting GetUint8SliceByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue bool
func BenchmarkHelperSetAndGetBoolByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "BoolValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "bool")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected bool
	}{
		{false},
		{true},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetBoolByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetBoolByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetBoolByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue byte
func BenchmarkHelperSetAndGetByteByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "ByteValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "byte")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected byte
	}{
		{0},
		{255},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetByteByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetByteByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetByteByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue float32
func BenchmarkHelperSetAndGetFloat32ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Float32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "float32")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected float32
	}{
		{1.401298464324817e-45},
		{3.4028234663852886e+38},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetFloat32ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetFloat32ByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetFloat32ByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue float64
func BenchmarkHelperSetAndGetFloat64ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Float64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "float64")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected float64
	}{
		{5e-324},
		{1.7976931348623157e+308},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetFloat64ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetFloat64ByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetFloat64ByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int
func BenchmarkHelperSetAndGetIntByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "IntValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int
	}{
		{-9223372036854775808},
		{9223372036854775807},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetIntByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetIntByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetIntByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int16
func BenchmarkHelperSetAndGetInt16ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Int16Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int16")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int16
	}{
		{-32768},
		{32767},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetInt16ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetInt16ByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetInt16ByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int32
func BenchmarkHelperSetAndGetInt32ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Int32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int32")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int32
	}{
		{-2147483648},
		{2147483647},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetInt32ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetInt32ByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetInt32ByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int64
func BenchmarkHelperSetAndGetInt64ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Int64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int64")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int64
	}{
		{-9223372036854775808},
		{9223372036854775807},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetInt64ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetInt64ByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetInt64ByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int8
func BenchmarkHelperSetAndGetInt8ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Int8Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int8")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int8
	}{
		{-128},
		{127},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetInt8ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetInt8ByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetInt8ByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue rune
func BenchmarkHelperSetAndGetRuneByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "RuneValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "rune")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected rune
	}{
		{'A'},
		{'Z'},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetRuneByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetRuneByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetRuneByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue string
func BenchmarkHelperSetAndGetStringByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "StringValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "string")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected string
	}{
		{"ABC"},
		{"abc"},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetStringByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetStringByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetStringByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint
func BenchmarkHelperSetAndGetUintByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "UintValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint
	}{
		{0},
		{18446744073709551615},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUintByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetUintByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetUintByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint16
func BenchmarkHelperSetAndGetUint16ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Uint16Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint16")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint16
	}{
		{0},
		{65535},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUint16ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetUint16ByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetUint16ByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint32
func BenchmarkHelperSetAndGetUint32ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Uint32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint32")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint32
	}{
		{0},
		{4294967295},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUint32ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetUint32ByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetUint32ByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint64
func BenchmarkHelperSetAndGetUint64ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Uint64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint64")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint64
	}{
		{0},
		{18446744073709551615},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUint64ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetUint64ByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetUint64ByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint8
func BenchmarkHelperSetAndGetUint8ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Uint8Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint8")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint8
	}{
		{0},
		{255},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUint8ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetUint8ByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetUint8ByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue *Table
func BenchmarkHelperSetAndGetTableByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "TableValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "*Table")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected *Table
	}{
		{NewNilTable()},
		{NewNilTable()},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetTableByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetTableByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetTableByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue time.Time
func BenchmarkHelperSetAndGetTimeByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "TimeValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "time.Time")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected time.Time
	}{
		{MinTime},
		{MaxTime},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetTimeByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetTimeByColIndex(colIndex, rowIndex)
			if err != nil {
				b.Error(err)
			}
			if value != test.expected {
				b.Errorf("expecting GetTimeByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	--------------------------------------------------------------------------------
//	bench test
//	next group: BenchmarkHelperSetAndGet<type>ByColIndex() functions for each of 20 types.
//	--------------------------------------------------------------------------------

//	Test Set and Get table cell in colIndex at rowIndex to newValue []byte
func BenchmarkHelperSetByteSliceByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "ByteSliceValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "[]byte")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected []byte
	}{
		{[]byte{0}},
		{[]byte{255}},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetByteSliceByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue []uint8
func BenchmarkHelperSetUint8SliceByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Uint8SliceValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "[]uint8")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected []uint8
	}{
		{[]uint8{0}},
		{[]uint8{255}},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUint8SliceByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue bool
func BenchmarkHelperSetBoolByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "BoolValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "bool")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected bool
	}{
		{false},
		{true},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetBoolByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue byte
func BenchmarkHelperSetByteByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "ByteValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "byte")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected byte
	}{
		{0},
		{255},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetByteByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue float32
func BenchmarkHelperSetFloat32ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Float32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "float32")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected float32
	}{
		{1.401298464324817e-45},
		{3.4028234663852886e+38},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetFloat32ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue float64
func BenchmarkHelperSetFloat64ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Float64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "float64")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected float64
	}{
		{5e-324},
		{1.7976931348623157e+308},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetFloat64ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int
func BenchmarkHelperSetIntByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "IntValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int
	}{
		{-9223372036854775808},
		{9223372036854775807},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetIntByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int16
func BenchmarkHelperSetInt16ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Int16Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int16")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int16
	}{
		{-32768},
		{32767},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetInt16ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int32
func BenchmarkHelperSetInt32ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Int32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int32")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int32
	}{
		{-2147483648},
		{2147483647},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetInt32ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int64
func BenchmarkHelperSetInt64ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Int64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int64")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int64
	}{
		{-9223372036854775808},
		{9223372036854775807},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetInt64ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int8
func BenchmarkHelperSetInt8ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Int8Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int8")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int8
	}{
		{-128},
		{127},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetInt8ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue rune
func BenchmarkHelperSetRuneByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "RuneValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "rune")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected rune
	}{
		{'A'},
		{'Z'},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetRuneByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue string
func BenchmarkHelperSetStringByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "StringValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "string")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected string
	}{
		{"ABC"},
		{"abc"},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetStringByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint
func BenchmarkHelperSetUintByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "UintValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint
	}{
		{0},
		{18446744073709551615},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUintByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint16
func BenchmarkHelperSetUint16ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Uint16Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint16")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint16
	}{
		{0},
		{65535},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUint16ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint32
func BenchmarkHelperSetUint32ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Uint32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint32")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint32
	}{
		{0},
		{4294967295},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUint32ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint64
func BenchmarkHelperSetUint64ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Uint64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint64")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint64
	}{
		{0},
		{18446744073709551615},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUint64ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint8
func BenchmarkHelperSetUint8ByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "Uint8Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint8")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint8
	}{
		{0},
		{255},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUint8ByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue *Table
func BenchmarkHelperSetTableByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "TableValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "*Table")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected *Table
	}{
		{NewNilTable()},
		{NewNilTable()},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetTableByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue time.Time
func BenchmarkHelperSetTimeByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "TimeValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "time.Time")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected time.Time
	}{
		{MinTime},
		{MaxTime},
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetTimeByColIndex(colIndex, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue []byte
func BenchmarkHelperSetAndGetByteSlice(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "ByteSliceValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "[]byte")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected []byte
	}{
		{[]byte{0}},
		{[]byte{255}},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetByteSlice(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetByteSlice(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if !bytes.Equal(value, test.expected) {
				b.Errorf("expecting GetByteSlice() bytes %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue []uint8
func BenchmarkHelperSetAndGetUint8Slice(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Uint8SliceValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "[]uint8")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected []uint8
	}{
		{[]uint8{0}},
		{[]uint8{255}},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint8Slice(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetUint8Slice(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if !bytes.Equal(value, test.expected) {
				b.Errorf("expecting GetUint8Slice() bytes %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue bool
func BenchmarkHelperSetAndGetBool(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "BoolValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "bool")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected bool
	}{
		{false},
		{true},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetBool(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetBool(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetBool() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue byte
func BenchmarkHelperSetAndGetByte(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "ByteValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "byte")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected byte
	}{
		{0},
		{255},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetByte(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetByte(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetByte() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue float32
func BenchmarkHelperSetAndGetFloat32(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Float32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "float32")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected float32
	}{
		{1.401298464324817e-45},
		{3.4028234663852886e+38},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetFloat32(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetFloat32(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetFloat32() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue float64
func BenchmarkHelperSetAndGetFloat64(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Float64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "float64")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected float64
	}{
		{5e-324},
		{1.7976931348623157e+308},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetFloat64(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetFloat64(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetFloat64() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int
func BenchmarkHelperSetAndGetInt(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "IntValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int
	}{
		{-9223372036854775808},
		{9223372036854775807},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetInt(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetInt(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetInt() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int16
func BenchmarkHelperSetAndGetInt16(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Int16Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int16")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int16
	}{
		{-32768},
		{32767},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetInt16(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetInt16(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetInt16() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int32
func BenchmarkHelperSetAndGetInt32(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Int32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int32")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int32
	}{
		{-2147483648},
		{2147483647},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetInt32(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetInt32(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetInt32() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int64
func BenchmarkHelperSetAndGetInt64(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Int64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int64")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int64
	}{
		{-9223372036854775808},
		{9223372036854775807},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetInt64(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetInt64(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetInt64() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int8
func BenchmarkHelperSetAndGetInt8(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Int8Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "int8")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected int8
	}{
		{-128},
		{127},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetInt8(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetInt8(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetInt8() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue rune
func BenchmarkHelperSetAndGetRune(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "RuneValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "rune")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected rune
	}{
		{'A'},
		{'Z'},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetRune(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetRune(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetRune() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue string
func BenchmarkHelperSetAndGetString(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "StringValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "string")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected string
	}{
		{"ABC"},
		{"abc"},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetString(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetString(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetString() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint
func BenchmarkHelperSetAndGetUint(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "UintValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint
	}{
		{0},
		{18446744073709551615},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetUint(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetUint() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint16
func BenchmarkHelperSetAndGetUint16(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Uint16Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint16")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint16
	}{
		{0},
		{65535},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint16(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetUint16(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetUint16() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint32
func BenchmarkHelperSetAndGetUint32(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Uint32Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint32")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint32
	}{
		{0},
		{4294967295},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint32(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetUint32(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetUint32() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint64
func BenchmarkHelperSetAndGetUint64(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Uint64Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint64")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint64
	}{
		{0},
		{18446744073709551615},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint64(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetUint64(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetUint64() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint8
func BenchmarkHelperSetAndGetUint8(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "Uint8Value"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "uint8")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected uint8
	}{
		{0},
		{255},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint8(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetUint8(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetUint8() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue *Table
func BenchmarkHelperSetAndGetTable(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "TableValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "*Table")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected *Table
	}{
		{NewNilTable()},
		{NewNilTable()},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetTable(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetTable(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetTable() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue time.Time
func BenchmarkHelperSetAndGetTime(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "TimeValue"

	table, err := NewTable("SetAndGet")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendCol(colName, "time.Time")
	if err != nil {
		b.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		b.Error(err)
	}

	var tests = []struct {
		expected time.Time
	}{
		{MinTime},
		{MaxTime},
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetTime(colName, rowIndex, test.expected)
			if err != nil {
				b.Error(err)
			}

			value, err := table.GetTime(colName, rowIndex)
			if err != nil {
				b.Error(err)
			}

			if value != test.expected {
				b.Errorf("expecting GetTime() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	Test GetByteSliceMustGet()
//  Test that the method panics on error.
func TestGetByteSliceMustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "[]byte")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("[]byte")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetByteSlice(colName, 0, expecting.([]byte))
	if err != nil {
		t.Fatal(err)
	}

	var got []byte
	got = table.GetByteSliceMustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches, err = Uint8SliceEquals(expecting.([]byte), got)
	if err != nil {
		t.Fatal(err)
	}

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetByteSliceMustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetByteSliceMustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetByteSliceMustGet(colName, minusIndex)
}

//	Test GetUint8SliceMustGet()
//  Test that the method panics on error.
func TestGetUint8SliceMustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "[]uint8")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("[]uint8")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetUint8Slice(colName, 0, expecting.([]uint8))
	if err != nil {
		t.Fatal(err)
	}

	var got []uint8
	got = table.GetUint8SliceMustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches, err = Uint8SliceEquals(expecting.([]uint8), got)
	if err != nil {
		t.Fatal(err)
	}

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetUint8SliceMustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetUint8SliceMustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetUint8SliceMustGet(colName, minusIndex)
}

//	Test GetBoolMustGet()
//  Test that the method panics on error.
func TestGetBoolMustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "bool")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("bool")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetBool(colName, 0, expecting.(bool))
	if err != nil {
		t.Fatal(err)
	}

	var got bool
	got = table.GetBoolMustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetBoolMustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetBoolMustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetBoolMustGet(colName, minusIndex)
}

//	Test GetByteMustGet()
//  Test that the method panics on error.
func TestGetByteMustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "byte")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("byte")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetByte(colName, 0, expecting.(byte))
	if err != nil {
		t.Fatal(err)
	}

	var got byte
	got = table.GetByteMustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetByteMustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetByteMustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetByteMustGet(colName, minusIndex)
}

//	Test GetFloat32MustGet()
//  Test that the method panics on error.
func TestGetFloat32MustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "float32")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("float32")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetFloat32(colName, 0, expecting.(float32))
	if err != nil {
		t.Fatal(err)
	}

	var got float32
	got = table.GetFloat32MustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetFloat32MustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetFloat32MustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetFloat32MustGet(colName, minusIndex)
}

//	Test GetFloat64MustGet()
//  Test that the method panics on error.
func TestGetFloat64MustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "float64")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("float64")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetFloat64(colName, 0, expecting.(float64))
	if err != nil {
		t.Fatal(err)
	}

	var got float64
	got = table.GetFloat64MustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetFloat64MustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetFloat64MustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetFloat64MustGet(colName, minusIndex)
}

//	Test GetIntMustGet()
//  Test that the method panics on error.
func TestGetIntMustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "int")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("int")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetInt(colName, 0, expecting.(int))
	if err != nil {
		t.Fatal(err)
	}

	var got int
	got = table.GetIntMustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetIntMustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetIntMustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetIntMustGet(colName, minusIndex)
}

//	Test GetInt16MustGet()
//  Test that the method panics on error.
func TestGetInt16MustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "int16")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("int16")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetInt16(colName, 0, expecting.(int16))
	if err != nil {
		t.Fatal(err)
	}

	var got int16
	got = table.GetInt16MustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetInt16MustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetInt16MustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetInt16MustGet(colName, minusIndex)
}

//	Test GetInt32MustGet()
//  Test that the method panics on error.
func TestGetInt32MustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "int32")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("int32")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetInt32(colName, 0, expecting.(int32))
	if err != nil {
		t.Fatal(err)
	}

	var got int32
	got = table.GetInt32MustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetInt32MustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetInt32MustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetInt32MustGet(colName, minusIndex)
}

//	Test GetInt64MustGet()
//  Test that the method panics on error.
func TestGetInt64MustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "int64")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("int64")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetInt64(colName, 0, expecting.(int64))
	if err != nil {
		t.Fatal(err)
	}

	var got int64
	got = table.GetInt64MustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetInt64MustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetInt64MustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetInt64MustGet(colName, minusIndex)
}

//	Test GetInt8MustGet()
//  Test that the method panics on error.
func TestGetInt8MustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "int8")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("int8")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetInt8(colName, 0, expecting.(int8))
	if err != nil {
		t.Fatal(err)
	}

	var got int8
	got = table.GetInt8MustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetInt8MustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetInt8MustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetInt8MustGet(colName, minusIndex)
}

//	Test GetRuneMustGet()
//  Test that the method panics on error.
func TestGetRuneMustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "rune")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("rune")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetRune(colName, 0, expecting.(rune))
	if err != nil {
		t.Fatal(err)
	}

	var got rune
	got = table.GetRuneMustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetRuneMustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetRuneMustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetRuneMustGet(colName, minusIndex)
}

//	Test GetStringMustGet()
//  Test that the method panics on error.
func TestGetStringMustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "string")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("string")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetString(colName, 0, expecting.(string))
	if err != nil {
		t.Fatal(err)
	}

	var got string
	got = table.GetStringMustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetStringMustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetStringMustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetStringMustGet(colName, minusIndex)
}

//	Test GetUintMustGet()
//  Test that the method panics on error.
func TestGetUintMustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "uint")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("uint")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetUint(colName, 0, expecting.(uint))
	if err != nil {
		t.Fatal(err)
	}

	var got uint
	got = table.GetUintMustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetUintMustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetUintMustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetUintMustGet(colName, minusIndex)
}

//	Test GetUint16MustGet()
//  Test that the method panics on error.
func TestGetUint16MustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "uint16")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("uint16")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetUint16(colName, 0, expecting.(uint16))
	if err != nil {
		t.Fatal(err)
	}

	var got uint16
	got = table.GetUint16MustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetUint16MustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetUint16MustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetUint16MustGet(colName, minusIndex)
}

//	Test GetUint32MustGet()
//  Test that the method panics on error.
func TestGetUint32MustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "uint32")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("uint32")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetUint32(colName, 0, expecting.(uint32))
	if err != nil {
		t.Fatal(err)
	}

	var got uint32
	got = table.GetUint32MustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetUint32MustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetUint32MustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetUint32MustGet(colName, minusIndex)
}

//	Test GetUint64MustGet()
//  Test that the method panics on error.
func TestGetUint64MustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "uint64")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("uint64")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetUint64(colName, 0, expecting.(uint64))
	if err != nil {
		t.Fatal(err)
	}

	var got uint64
	got = table.GetUint64MustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetUint64MustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetUint64MustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetUint64MustGet(colName, minusIndex)
}

//	Test GetUint8MustGet()
//  Test that the method panics on error.
func TestGetUint8MustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "uint8")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("uint8")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetUint8(colName, 0, expecting.(uint8))
	if err != nil {
		t.Fatal(err)
	}

	var got uint8
	got = table.GetUint8MustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetUint8MustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetUint8MustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetUint8MustGet(colName, minusIndex)
}

//	Test GetTableMustGet()
//  Test that the method panics on error.
func TestGetTableMustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "*Table")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("*Table")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetTable(colName, 0, expecting.(*Table))
	if err != nil {
		t.Fatal(err)
	}

	var got *Table
	got = table.GetTableMustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetTableMustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetTableMustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetTableMustGet(colName, minusIndex)
}

//	Test GetTimeMustGet()
//  Test that the method panics on error.
func TestGetTimeMustGet(t *testing.T) {

	// See: TestGet<type>MustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	err = table.AppendCol(colName, "time.Time")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("time.Time")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetTime(colName, 0, expecting.(time.Time))
	if err != nil {
		t.Fatal(err)
	}

	var got time.Time
	got = table.GetTimeMustGet(colName, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetTimeMustGet(%q, 0) expecting %v, but got %v", colName, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetTimeMustGet(%q, %d) expecting panic(), but didn't panic()", colName, minusIndex)
		}
	}()

	table.GetTimeMustGet(colName, minusIndex)
}

//	Test GetByteSliceByColIndexMustGet()
//  Test that the method panics on error.
func TestGetByteSliceByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "[]byte")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("[]byte")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetByteSliceByColIndex(colIndex, 0, expecting.([]byte))
	if err != nil {
		t.Fatal(err)
	}

	var got []byte
	got = table.GetByteSliceByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches, err = Uint8SliceEquals(expecting.([]byte), got)
	if err != nil {
		t.Fatal(err)
	}

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetByteSliceByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetByteSliceByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetByteSliceByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetUint8SliceByColIndexMustGet()
//  Test that the method panics on error.
func TestGetUint8SliceByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "[]uint8")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("[]uint8")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetUint8SliceByColIndex(colIndex, 0, expecting.([]uint8))
	if err != nil {
		t.Fatal(err)
	}

	var got []uint8
	got = table.GetUint8SliceByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches, err = Uint8SliceEquals(expecting.([]uint8), got)
	if err != nil {
		t.Fatal(err)
	}

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetUint8SliceByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetUint8SliceByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetUint8SliceByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetBoolByColIndexMustGet()
//  Test that the method panics on error.
func TestGetBoolByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "bool")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("bool")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetBoolByColIndex(colIndex, 0, expecting.(bool))
	if err != nil {
		t.Fatal(err)
	}

	var got bool
	got = table.GetBoolByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetBoolByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetBoolByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetBoolByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetByteByColIndexMustGet()
//  Test that the method panics on error.
func TestGetByteByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "byte")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("byte")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetByteByColIndex(colIndex, 0, expecting.(byte))
	if err != nil {
		t.Fatal(err)
	}

	var got byte
	got = table.GetByteByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetByteByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetByteByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetByteByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetFloat32ByColIndexMustGet()
//  Test that the method panics on error.
func TestGetFloat32ByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "float32")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("float32")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetFloat32ByColIndex(colIndex, 0, expecting.(float32))
	if err != nil {
		t.Fatal(err)
	}

	var got float32
	got = table.GetFloat32ByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetFloat32ByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetFloat32ByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetFloat32ByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetFloat64ByColIndexMustGet()
//  Test that the method panics on error.
func TestGetFloat64ByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "float64")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("float64")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetFloat64ByColIndex(colIndex, 0, expecting.(float64))
	if err != nil {
		t.Fatal(err)
	}

	var got float64
	got = table.GetFloat64ByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetFloat64ByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetFloat64ByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetFloat64ByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetIntByColIndexMustGet()
//  Test that the method panics on error.
func TestGetIntByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "int")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("int")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetIntByColIndex(colIndex, 0, expecting.(int))
	if err != nil {
		t.Fatal(err)
	}

	var got int
	got = table.GetIntByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetIntByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetIntByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetIntByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetInt16ByColIndexMustGet()
//  Test that the method panics on error.
func TestGetInt16ByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "int16")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("int16")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetInt16ByColIndex(colIndex, 0, expecting.(int16))
	if err != nil {
		t.Fatal(err)
	}

	var got int16
	got = table.GetInt16ByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetInt16ByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetInt16ByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetInt16ByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetInt32ByColIndexMustGet()
//  Test that the method panics on error.
func TestGetInt32ByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "int32")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("int32")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetInt32ByColIndex(colIndex, 0, expecting.(int32))
	if err != nil {
		t.Fatal(err)
	}

	var got int32
	got = table.GetInt32ByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetInt32ByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetInt32ByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetInt32ByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetInt64ByColIndexMustGet()
//  Test that the method panics on error.
func TestGetInt64ByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "int64")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("int64")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetInt64ByColIndex(colIndex, 0, expecting.(int64))
	if err != nil {
		t.Fatal(err)
	}

	var got int64
	got = table.GetInt64ByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetInt64ByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetInt64ByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetInt64ByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetInt8ByColIndexMustGet()
//  Test that the method panics on error.
func TestGetInt8ByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "int8")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("int8")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetInt8ByColIndex(colIndex, 0, expecting.(int8))
	if err != nil {
		t.Fatal(err)
	}

	var got int8
	got = table.GetInt8ByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetInt8ByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetInt8ByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetInt8ByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetRuneByColIndexMustGet()
//  Test that the method panics on error.
func TestGetRuneByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "rune")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("rune")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetRuneByColIndex(colIndex, 0, expecting.(rune))
	if err != nil {
		t.Fatal(err)
	}

	var got rune
	got = table.GetRuneByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetRuneByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetRuneByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetRuneByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetStringByColIndexMustGet()
//  Test that the method panics on error.
func TestGetStringByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "string")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("string")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStringByColIndex(colIndex, 0, expecting.(string))
	if err != nil {
		t.Fatal(err)
	}

	var got string
	got = table.GetStringByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetStringByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetStringByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetStringByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetUintByColIndexMustGet()
//  Test that the method panics on error.
func TestGetUintByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "uint")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("uint")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetUintByColIndex(colIndex, 0, expecting.(uint))
	if err != nil {
		t.Fatal(err)
	}

	var got uint
	got = table.GetUintByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetUintByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetUintByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetUintByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetUint16ByColIndexMustGet()
//  Test that the method panics on error.
func TestGetUint16ByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "uint16")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("uint16")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetUint16ByColIndex(colIndex, 0, expecting.(uint16))
	if err != nil {
		t.Fatal(err)
	}

	var got uint16
	got = table.GetUint16ByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetUint16ByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetUint16ByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetUint16ByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetUint32ByColIndexMustGet()
//  Test that the method panics on error.
func TestGetUint32ByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "uint32")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("uint32")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetUint32ByColIndex(colIndex, 0, expecting.(uint32))
	if err != nil {
		t.Fatal(err)
	}

	var got uint32
	got = table.GetUint32ByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetUint32ByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetUint32ByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetUint32ByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetUint64ByColIndexMustGet()
//  Test that the method panics on error.
func TestGetUint64ByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "uint64")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("uint64")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetUint64ByColIndex(colIndex, 0, expecting.(uint64))
	if err != nil {
		t.Fatal(err)
	}

	var got uint64
	got = table.GetUint64ByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetUint64ByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetUint64ByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetUint64ByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetUint8ByColIndexMustGet()
//  Test that the method panics on error.
func TestGetUint8ByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "uint8")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("uint8")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetUint8ByColIndex(colIndex, 0, expecting.(uint8))
	if err != nil {
		t.Fatal(err)
	}

	var got uint8
	got = table.GetUint8ByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetUint8ByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetUint8ByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetUint8ByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetTableByColIndexMustGet()
//  Test that the method panics on error.
func TestGetTableByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "*Table")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("*Table")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetTableByColIndex(colIndex, 0, expecting.(*Table))
	if err != nil {
		t.Fatal(err)
	}

	var got *Table
	got = table.GetTableByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetTableByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetTableByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetTableByColIndexMustGet(colIndex, minusIndex)
}

//	Test GetTimeByColIndexMustGet()
//  Test that the method panics on error.
func TestGetTimeByColIndexMustGet(t *testing.T) {

	// See: TestGet<type>ByColIndexMustGet() functions

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName string = "MyCol"
	var colIndex int = 0
	err = table.AppendCol(colName, "time.Time")
	if err != nil {
		t.Fatal(err)
	}

	err = table.AppendRows(1)
	if err != nil {
		t.Fatal(err)
	}

	// Test a simple get.

	var expecting interface{}
	expecting, err = nonZeroValue("time.Time")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetTimeByColIndex(colIndex, 0, expecting.(time.Time))
	if err != nil {
		t.Fatal(err)
	}

	var got time.Time
	got = table.GetTimeByColIndexMustGet(colIndex, 0)
	if err != nil {
		t.Fatal(err)
	}

	var matches bool
	matches = (expecting == got)

	//	where(table)
	//	where(fmt.Sprintf("got == expecting = %t", matches))

	if !matches {
		t.Fatalf("func TestGetTimeByColIndexMustGet(%d, 0) expecting %v, but got %v", colIndex, expecting, got)
	}

	// Test that the method panics with an invalid argument.

	const minusIndex = -1 // Will trigger error, therefore panic.

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("func TestGetTimeByColIndexMustGet(0, %d) expecting panic(), but didn't panic()", minusIndex)
		}
	}()

	table.GetTimeByColIndexMustGet(colIndex, minusIndex)
}

//  Test that the gotables column type constants are correct.
func TestColumnConstants(t *testing.T) {

	var err error
	var table *Table

	table, err = NewTable("MyTable")
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetStructShape(true)
	if err != nil {
		t.Fatal(err)
	}

	var colName rune = 'a'

	err = table.AppendCol(string(colName), ByteSlice)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Uint8Slice)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Bool)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Byte)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Float32)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Float64)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Int)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Int16)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Int32)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Int64)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Int8)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Rune)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), String)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Uint)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Uint16)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Uint32)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Uint64)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), Uint8)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), gotables_Table)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	err = table.AppendCol(string(colName), time_Time)
	if err != nil {
		t.Fatal(err)
	}
	colName++

	// println(table.String())
}
