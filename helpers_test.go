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

//	-----------------------------------------------------------
//	TestSet<type>() functions for each of 17 types.
//	-----------------------------------------------------------

//	Test Set table cell in colName at rowIndex to newValue string
func TestSetString(t *testing.T) {

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

		value, _ := table.GetString(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting .GetString() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue bool
func TestSetBool(t *testing.T) {

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

		value, _ := table.GetBool(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting .GetBool() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue int
func TestSetInt(t *testing.T) {

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

		value, _ := table.GetInt(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting .GetInt() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue int8
func TestSetInt8(t *testing.T) {

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

		value, _ := table.GetInt8(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting .GetInt8() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue int16
func TestSetInt16(t *testing.T) {

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

		value, _ := table.GetInt16(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting .GetInt16() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue int32
func TestSetInt32(t *testing.T) {

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

		value, _ := table.GetInt32(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting .GetInt32() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue int64
func TestSetInt64(t *testing.T) {

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

		value, _ := table.GetInt64(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting .GetInt64() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue uint
func TestSetUint(t *testing.T) {

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

		value, _ := table.GetUint(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting .GetUint() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue byte
func TestSetByte(t *testing.T) {

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

		value, _ := table.GetByte(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting .GetByte() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue uint8
func TestSetUint8(t *testing.T) {

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

		value, _ := table.GetUint8(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting .GetUint8() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue uint16
func TestSetUint16(t *testing.T) {

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

		value, _ := table.GetUint16(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting .GetUint16() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue uint32
func TestSetUint32(t *testing.T) {

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

		value, _ := table.GetUint32(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting .GetUint32() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue uint64
func TestSetUint64(t *testing.T) {

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

		value, _ := table.GetUint64(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting .GetUint64() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue float32
func TestSetFloat32(t *testing.T) {

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

		value, _ := table.GetFloat32(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting .GetFloat32() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue float64
func TestSetFloat64(t *testing.T) {

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

		value, _ := table.GetFloat64(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting .GetFloat64() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue []byte
func TestSetByteSlice(t *testing.T) {

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

		value, _ := table.GetByteSlice(colName, rowIndex)

		if !bytes.Equal(value, test.expected) {

			t.Errorf("expecting .GetByteSlice() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set table cell in colName at rowIndex to newValue []uint8
func TestSetUint8Slice(t *testing.T) {

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

		value, _ := table.GetUint8Slice(colName, rowIndex)

		if !bytes.Equal(value, test.expected) {

			t.Errorf("expecting .GetUint8Slice() value %v, not %v", test.expected, value)
		}
	}
}

