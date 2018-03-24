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

	var expected string

	expected = "abc"

	rowIndex := 0

	err = table.SetString(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetString(colName, rowIndex)

	if value != expected {

		t.Errorf("expecting .GetString() value %v, not %v", expected, value)
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

	var expected bool

	expected = true

	rowIndex := 0

	err = table.SetBool(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetBool(colName, rowIndex)

	if value != expected {

		t.Errorf("expecting .GetBool() value %v, not %v", expected, value)
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

	var expected int

	expected = 9223372036854775807

	rowIndex := 0

	err = table.SetInt(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetInt(colName, rowIndex)

	if value != expected {

		t.Errorf("expecting .GetInt() value %v, not %v", expected, value)
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

	var expected int8

	expected = 127

	rowIndex := 0

	err = table.SetInt8(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetInt8(colName, rowIndex)

	if value != expected {

		t.Errorf("expecting .GetInt8() value %v, not %v", expected, value)
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

	var expected int16

	expected = 32767

	rowIndex := 0

	err = table.SetInt16(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetInt16(colName, rowIndex)

	if value != expected {

		t.Errorf("expecting .GetInt16() value %v, not %v", expected, value)
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

	var expected int32

	expected = 2147483647

	rowIndex := 0

	err = table.SetInt32(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetInt32(colName, rowIndex)

	if value != expected {

		t.Errorf("expecting .GetInt32() value %v, not %v", expected, value)
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

	var expected int64

	expected = 9223372036854775807

	rowIndex := 0

	err = table.SetInt64(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetInt64(colName, rowIndex)

	if value != expected {

		t.Errorf("expecting .GetInt64() value %v, not %v", expected, value)
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

	var expected uint

	expected = 18446744073709551615

	rowIndex := 0

	err = table.SetUint(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetUint(colName, rowIndex)

	if value != expected {

		t.Errorf("expecting .GetUint() value %v, not %v", expected, value)
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

	var expected byte

	expected = 255

	rowIndex := 0

	err = table.SetByte(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetByte(colName, rowIndex)

	if value != expected {

		t.Errorf("expecting .GetByte() value %v, not %v", expected, value)
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

	var expected uint8

	expected = 255

	rowIndex := 0

	err = table.SetUint8(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetUint8(colName, rowIndex)

	if value != expected {

		t.Errorf("expecting .GetUint8() value %v, not %v", expected, value)
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

	var expected uint16

	expected = 65535

	rowIndex := 0

	err = table.SetUint16(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetUint16(colName, rowIndex)

	if value != expected {

		t.Errorf("expecting .GetUint16() value %v, not %v", expected, value)
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

	var expected uint32

	expected = 4294967295

	rowIndex := 0

	err = table.SetUint32(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetUint32(colName, rowIndex)

	if value != expected {

		t.Errorf("expecting .GetUint32() value %v, not %v", expected, value)
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

	var expected uint64

	expected = 18446744073709551615

	rowIndex := 0

	err = table.SetUint64(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetUint64(colName, rowIndex)

	if value != expected {

		t.Errorf("expecting .GetUint64() value %v, not %v", expected, value)
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

	var expected float32

	expected = 3.4028234663852886e+38

	rowIndex := 0

	err = table.SetFloat32(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetFloat32(colName, rowIndex)

	if value != expected {

		t.Errorf("expecting .GetFloat32() value %v, not %v", expected, value)
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

	var expected float64

	expected = 1.7976931348623157e+308

	rowIndex := 0

	err = table.SetFloat64(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetFloat64(colName, rowIndex)

	if value != expected {

		t.Errorf("expecting .GetFloat64() value %v, not %v", expected, value)
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

	var expected []byte

	expected = []byte{ 255 }

	rowIndex := 0

	err = table.SetByteSlice(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetByteSlice(colName, rowIndex)

	if !bytes.Equal(value, expected) {

		t.Errorf("expecting .GetByteSlice() value %v, not %v", expected, value)
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

	var expected []uint8

	expected = []uint8{ 255 }

	rowIndex := 0

	err = table.SetUint8Slice(colName, rowIndex, expected)
    if err != nil { t.Error(err) }

	value, err := table.GetUint8Slice(colName, rowIndex)

	if !bytes.Equal(value, expected) {

		t.Errorf("expecting .GetUint8Slice() value %v, not %v", expected, value)
	}
}

