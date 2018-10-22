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

//	-----------------------------------------------------------------------
//	next group: TestSet<type>() functions for each of 18 types.
//	-----------------------------------------------------------------------

//	Test Set and Get table cell in colName at rowIndex to newValue []byte
func TestSetAndGetByteSlice(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "ByteSliceValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "[]byte")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected []byte
	}{
		{ []byte{ 0 } },
		{ []byte{ 255 } },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetByteSlice(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetByteSlice(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if !bytes.Equal(value, test.expected) {
			t.Errorf("expecting GetByteSlice() bytes %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue []uint8
func TestSetAndGetUint8Slice(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint8SliceValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "[]uint8")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected []uint8
	}{
		{ []uint8{ 0 } },
		{ []uint8{ 255 } },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint8Slice(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetUint8Slice(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if !bytes.Equal(value, test.expected) {
			t.Errorf("expecting GetUint8Slice() bytes %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue bool
func TestSetAndGetBool(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "BoolValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "bool")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected bool
	}{
		{ false },
		{ true },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetBool(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetBool(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetBool() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue byte
func TestSetAndGetByte(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "ByteValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "byte")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected byte
	}{
		{ 0 },
		{ 255 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetByte(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetByte(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetByte() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue float32
func TestSetAndGetFloat32(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Float32Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "float32")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected float32
	}{
		{ 1.401298464324817e-45 },
		{ 3.4028234663852886e+38 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetFloat32(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetFloat32(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetFloat32() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue float64
func TestSetAndGetFloat64(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Float64Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "float64")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected float64
	}{
		{ 5e-324 },
		{ 1.7976931348623157e+308 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetFloat64(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetFloat64(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetFloat64() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int
func TestSetAndGetInt(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "IntValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int
	}{
		{ -9223372036854775808 },
		{ 9223372036854775807 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetInt(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetInt(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetInt() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int16
func TestSetAndGetInt16(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int16Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int16")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int16
	}{
		{ -32768 },
		{ 32767 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetInt16(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetInt16(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetInt16() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int32
func TestSetAndGetInt32(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int32Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int32")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int32
	}{
		{ -2147483648 },
		{ 2147483647 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetInt32(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetInt32(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetInt32() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int64
func TestSetAndGetInt64(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int64Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int64")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int64
	}{
		{ -9223372036854775808 },
		{ 9223372036854775807 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetInt64(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetInt64(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetInt64() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int8
func TestSetAndGetInt8(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int8Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int8")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int8
	}{
		{ -128 },
		{ 127 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetInt8(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetInt8(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetInt8() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue rune
func TestSetAndGetRune(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "RuneValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "rune")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected rune
	}{
		{ 'A' },
		{ 'Z' },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetRune(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetRune(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetRune() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue string
func TestSetAndGetString(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "StringValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "string")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected string
	}{
		{ "ABC" },
		{ "abc" },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetString(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetString(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetString() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint
func TestSetAndGetUint(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "UintValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint
	}{
		{ 0 },
		{ 18446744073709551615 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetUint(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetUint() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint16
func TestSetAndGetUint16(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint16Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint16")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint16
	}{
		{ 0 },
		{ 65535 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint16(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetUint16(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetUint16() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint32
func TestSetAndGetUint32(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint32Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint32")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint32
	}{
		{ 0 },
		{ 4294967295 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint32(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetUint32(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetUint32() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint64
func TestSetAndGetUint64(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint64Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint64")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint64
	}{
		{ 0 },
		{ 18446744073709551615 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint64(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetUint64(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetUint64() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint8
func TestSetAndGetUint8(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint8Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint8")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint8
	}{
		{ 0 },
		{ 255 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint8(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetUint8(colName, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetUint8() value %v, not %v", test.expected, value)
		}
	}
}

//	--------------------------------------------------------------------------------
//	next group: TestSet<type>ByColIndex() functions for each of 18 types.
//	--------------------------------------------------------------------------------

//	Test Set and Get table cell in colIndex at rowIndex to newValue []byte
func TestHelperSetAndGetByteSliceByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "ByteSliceValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "[]byte")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected []byte
	}{
		{ []byte{ 0 } },
		{ []byte{ 255 } },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetByteSliceByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetByteSliceByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if !bytes.Equal(value, test.expected) {
			t.Errorf("expecting GetByteSliceByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue []uint8
func TestHelperSetAndGetUint8SliceByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint8SliceValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "[]uint8")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected []uint8
	}{
		{ []uint8{ 0 } },
		{ []uint8{ 255 } },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUint8SliceByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetUint8SliceByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if !bytes.Equal(value, test.expected) {
			t.Errorf("expecting GetUint8SliceByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue bool
func TestHelperSetAndGetBoolByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "BoolValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "bool")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected bool
	}{
		{ false },
		{ true },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetBoolByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetBoolByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetBoolByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue byte
func TestHelperSetAndGetByteByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "ByteValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "byte")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected byte
	}{
		{ 0 },
		{ 255 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetByteByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetByteByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetByteByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue float32
func TestHelperSetAndGetFloat32ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Float32Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "float32")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected float32
	}{
		{ 1.401298464324817e-45 },
		{ 3.4028234663852886e+38 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetFloat32ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetFloat32ByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetFloat32ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue float64
func TestHelperSetAndGetFloat64ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Float64Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "float64")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected float64
	}{
		{ 5e-324 },
		{ 1.7976931348623157e+308 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetFloat64ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetFloat64ByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetFloat64ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int
func TestHelperSetAndGetIntByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "IntValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int
	}{
		{ -9223372036854775808 },
		{ 9223372036854775807 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetIntByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetIntByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetIntByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int16
func TestHelperSetAndGetInt16ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int16Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int16")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int16
	}{
		{ -32768 },
		{ 32767 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetInt16ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetInt16ByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetInt16ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int32
func TestHelperSetAndGetInt32ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int32Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int32")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int32
	}{
		{ -2147483648 },
		{ 2147483647 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetInt32ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetInt32ByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetInt32ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int64
func TestHelperSetAndGetInt64ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int64Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int64")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int64
	}{
		{ -9223372036854775808 },
		{ 9223372036854775807 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetInt64ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetInt64ByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetInt64ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int8
func TestHelperSetAndGetInt8ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int8Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int8")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int8
	}{
		{ -128 },
		{ 127 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetInt8ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetInt8ByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetInt8ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue rune
func TestHelperSetAndGetRuneByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "RuneValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "rune")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected rune
	}{
		{ 'A' },
		{ 'Z' },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetRuneByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetRuneByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetRuneByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue string
func TestHelperSetAndGetStringByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "StringValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "string")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected string
	}{
		{ "ABC" },
		{ "abc" },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetStringByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetStringByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetStringByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint
func TestHelperSetAndGetUintByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "UintValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint
	}{
		{ 0 },
		{ 18446744073709551615 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUintByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetUintByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetUintByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint16
func TestHelperSetAndGetUint16ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint16Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint16")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint16
	}{
		{ 0 },
		{ 65535 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUint16ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetUint16ByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetUint16ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint32
func TestHelperSetAndGetUint32ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint32Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint32")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint32
	}{
		{ 0 },
		{ 4294967295 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUint32ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetUint32ByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetUint32ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint64
func TestHelperSetAndGetUint64ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint64Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint64")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint64
	}{
		{ 0 },
		{ 18446744073709551615 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUint64ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetUint64ByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetUint64ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint8
func TestHelperSetAndGetUint8ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint8Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint8")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint8
	}{
		{ 0 },
		{ 255 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUint8ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, err := table.GetUint8ByColIndex(colIndex, rowIndex)
	    if err != nil { t.Error(err) }

		if value != test.expected {
			t.Errorf("expecting GetUint8ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	-----------------------------------------------------------------------
//	bench test
//	next group: TestSet<type>() functions for each of 18 types.
//	-----------------------------------------------------------------------

//	Test Set and Get table cell in colName at rowIndex to newValue []byte
func BenchmarkHelperSetAndGetByteSlice(b *testing.B) {

	// See: TestSet<type>() functions

	// Set up for benchmark.

	const colName string = "ByteSliceValue"

    table, err := NewTable("SetAndGet")
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "[]byte")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected []byte
	}{
		{ []byte{ 0 } },
		{ []byte{ 255 } },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetByteSlice(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetByteSlice(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "[]uint8")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected []uint8
	}{
		{ []uint8{ 0 } },
		{ []uint8{ 255 } },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint8Slice(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetUint8Slice(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "bool")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected bool
	}{
		{ false },
		{ true },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetBool(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetBool(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "byte")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected byte
	}{
		{ 0 },
		{ 255 },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetByte(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetByte(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "float32")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected float32
	}{
		{ 1.401298464324817e-45 },
		{ 3.4028234663852886e+38 },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetFloat32(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetFloat32(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "float64")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected float64
	}{
		{ 5e-324 },
		{ 1.7976931348623157e+308 },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetFloat64(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetFloat64(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "int")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected int
	}{
		{ -9223372036854775808 },
		{ 9223372036854775807 },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetInt(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetInt(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "int16")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected int16
	}{
		{ -32768 },
		{ 32767 },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetInt16(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetInt16(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "int32")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected int32
	}{
		{ -2147483648 },
		{ 2147483647 },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetInt32(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetInt32(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "int64")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected int64
	}{
		{ -9223372036854775808 },
		{ 9223372036854775807 },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetInt64(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetInt64(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "int8")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected int8
	}{
		{ -128 },
		{ 127 },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetInt8(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetInt8(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "rune")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected rune
	}{
		{ 'A' },
		{ 'Z' },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetRune(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetRune(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "string")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected string
	}{
		{ "ABC" },
		{ "abc" },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetString(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetString(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "uint")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected uint
	}{
		{ 0 },
		{ 18446744073709551615 },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetUint(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "uint16")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected uint16
	}{
		{ 0 },
		{ 65535 },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint16(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetUint16(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "uint32")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected uint32
	}{
		{ 0 },
		{ 4294967295 },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint32(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetUint32(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "uint64")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected uint64
	}{
		{ 0 },
		{ 18446744073709551615 },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint64(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetUint64(colName, rowIndex)
		    if err != nil { b.Error(err) }

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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "uint8")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected uint8
	}{
		{ 0 },
		{ 255 },
	}

	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {

			err = table.SetUint8(colName, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetUint8(colName, rowIndex)
		    if err != nil { b.Error(err) }

			if value != test.expected {
				b.Errorf("expecting GetUint8() value %v, not %v", test.expected, value)
			}
		}
	}
}

//	--------------------------------------------------------------------------------
//	bench test
//	next group: TestSet<type>ByColIndex() functions for each of 18 types.
//	--------------------------------------------------------------------------------

//	Test Set and Get table cell in colIndex at rowIndex to newValue []byte
func BenchmarkHelperSetAndGetByteSliceByColIndex(b *testing.B) {

	// See: TestSet<type>() functions

	const colName string = "ByteSliceValue"

    table, err := NewTable("SetAndGet")
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "[]byte")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected []byte
	}{
		{ []byte{ 0 } },
		{ []byte{ 255 } },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetByteSliceByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetByteSliceByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "[]uint8")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected []uint8
	}{
		{ []uint8{ 0 } },
		{ []uint8{ 255 } },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUint8SliceByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetUint8SliceByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "bool")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected bool
	}{
		{ false },
		{ true },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetBoolByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetBoolByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "byte")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected byte
	}{
		{ 0 },
		{ 255 },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetByteByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetByteByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "float32")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected float32
	}{
		{ 1.401298464324817e-45 },
		{ 3.4028234663852886e+38 },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetFloat32ByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetFloat32ByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "float64")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected float64
	}{
		{ 5e-324 },
		{ 1.7976931348623157e+308 },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetFloat64ByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetFloat64ByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "int")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected int
	}{
		{ -9223372036854775808 },
		{ 9223372036854775807 },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetIntByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetIntByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "int16")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected int16
	}{
		{ -32768 },
		{ 32767 },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetInt16ByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetInt16ByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "int32")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected int32
	}{
		{ -2147483648 },
		{ 2147483647 },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetInt32ByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetInt32ByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "int64")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected int64
	}{
		{ -9223372036854775808 },
		{ 9223372036854775807 },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetInt64ByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetInt64ByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "int8")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected int8
	}{
		{ -128 },
		{ 127 },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetInt8ByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetInt8ByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "rune")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected rune
	}{
		{ 'A' },
		{ 'Z' },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetRuneByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetRuneByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "string")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected string
	}{
		{ "ABC" },
		{ "abc" },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetStringByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetStringByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "uint")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected uint
	}{
		{ 0 },
		{ 18446744073709551615 },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUintByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetUintByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "uint16")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected uint16
	}{
		{ 0 },
		{ 65535 },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUint16ByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetUint16ByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "uint32")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected uint32
	}{
		{ 0 },
		{ 4294967295 },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUint32ByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetUint32ByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "uint64")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected uint64
	}{
		{ 0 },
		{ 18446744073709551615 },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUint64ByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetUint64ByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
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
    if err != nil { b.Error(err) }

	err = table.AppendCol(colName, "uint8")
    if err != nil { b.Error(err) }

	err = table.AppendRow()
    if err != nil { b.Error(err) }

	var tests = []struct {
		expected uint8
	}{
		{ 0 },
		{ 255 },
	}

	const colIndex = 0
	const rowIndex = 0

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			err = table.SetUint8ByColIndex(colIndex, rowIndex, test.expected)
		    if err != nil { b.Error(err) }

			value, err := table.GetUint8ByColIndex(colIndex, rowIndex)
		    if err != nil { b.Error(err) }
			if value != test.expected {
				b.Errorf("expecting GetUint8ByColIndex() value %v, not %v", test.expected, value)
			}
		}
	}
}

