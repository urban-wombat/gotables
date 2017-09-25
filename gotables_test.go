package gotables

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
//	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
//	"time"
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

func TestRenameTable(t *testing.T) {
	var err error
	var table *Table
	var setupName string = "Fred"

	var tests3 = []struct {
		input    string
		succeeds bool
		output   string
	}{
		{"Barney", true, "Barney"},
		{"", false, "Fred"},
		{"$&*", false, "Fred"},
	}

	for _, test := range tests3 {

		table, err = NewTable(setupName)
		if err != nil {
			t.Error(err)
		}

		err = table.RenameTable(test.input)
		if (err == nil) != test.succeeds {
			t.Errorf("Error renaming to %q: %s", test.output, err)
		}

		var tableName string = table.Name()
		if tableName != test.output {
			t.Errorf("Expected %q, not %q", test.output, tableName)
		}
	}
}

func TestTableSetRenameTable(t *testing.T) {
	/*
		tableSet, err := NewTableSetFromString(`[Wilma]`)
		if err != nil {
			panic(err)
		}
	*/
	//	fmt.Printf("tableSet.TableCount() = %d\n", tableSet.TableCount())

	tests := []struct {
		renameFrom string
		renameTo   string
		succeeds   bool
	}{
		{"Wilma", "Betty", true},
		{"Betty", "Wilma", false},
		{"Wilma", "Wilma", false},
		{"", "Wilma", false},
		{"Wilma", "", false},
	}

	for i, test := range tests {
		tableSet, err := NewTableSetFromString(`[Wilma]`)
		if err != nil {
			t.Error(err)
		}
		err = tableSet.RenameTable(test.renameFrom, test.renameTo)
		if (err == nil) != test.succeeds {
			t.Errorf("test[%d]: Error renaming from %q to %q: %v", i, test.renameFrom, test.renameTo, err)
		}
	}
}

func TestReadString1(t *testing.T) {
	tableSet, err := NewTableSetFromString(
		`[EmptyTable1]

		[EmptyTable2]

		[TableWithColNamesAndTypes]
		A	B	C
		int	int	int

		[TableWithRow]
		D	E	F
		int	int	int
		1	2	3

		[TableWithRows]
		G	H	I
		int	int	int
		4	5	6
		7	8	9
	`)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		tableName string
		colCount  int
		rowCount  int
		succeeds  bool
	}{
		{"EmptyTable1", 0, 0, true},
		{"EmptyTable2", 0, 0, true},
		{"TableWithColNamesAndTypes", 3, 0, true},
		{"TableWithRow", 3, 1, true},
		{"TableWithRows", 3, 2, true},
	}

	for i, test := range tests {
		table, err := tableSet.Table(test.tableName)
		if err != nil {
			t.Errorf("[%d] %v", i, err)
		}
		if (err == nil) != test.succeeds {
			t.Errorf("test[%d]: err == %v but expecting succeeds == %t", i, err, test.succeeds)
		}

		colCount := table.ColCount()
		if colCount != test.colCount {
			t.Errorf("test[%d]: expecting [%s] colCount %d, not %d\n", i, test.tableName, test.colCount, colCount)
		}

		rowCount := table.RowCount()
		if rowCount != test.rowCount {
			t.Errorf("test[%d]: expecting [%s] rowCount %d, not %d\n", i, test.tableName, test.rowCount, rowCount)
		}
	}
}

func TestReadString2(t *testing.T) {
	_, err := NewTableSetFromString(
		`[EmptyTable1]

		# Should be a syntax error. Table should have both names AND types.
		[TableWithColNamesOnly]
		A	B	C

		[EmptyTable2]
	`)
	if err == nil {
		t.Errorf("Should return a syntax error. Table should have both names AND types.")
	}
}

func TestReadString3(t *testing.T) {
	_, err := NewTableSetFromString(
		`[TableWithRow]
		D	E	F
		int	int	int
		1	2	3

		A	B	C
	`)
	if err == nil {
		t.Errorf("Should return a syntax error. Col names should not follow blank lines.")
	}
}

func TestReadString4(t *testing.T) {
	_, err := NewTableSetFromString(
		`[TableWithRow]
		D	E	F
		int	int	int
		1	2	3

		4	5	6
		`)
	if err == nil {
		t.Errorf("Should return a syntax error. Col values should not follow blank lines.")
	}
}

func TestReadString5(t *testing.T) {
	tableSet, err := NewTableSetFromString(
		`[TableEmpty]
		
		`)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		tableName string
		colCount  int
		rowCount  int
		succeeds  bool
	}{
		{"TableEmpty", 0, 0, true},
	}

	for i, test := range tests {
		table, err := tableSet.Table(test.tableName)
		if err != nil {
			t.Errorf("[%d] %v", i, err)
		}
		if (err == nil) != test.succeeds {
			t.Errorf("test[%d]: err == %v but expecting succeeds == %t", i, err, test.succeeds)
		}

		colCount := table.ColCount()
		if colCount != test.colCount {
			t.Errorf("test[%d]: expecting [%s] colCount %d, not %d\n", i, test.tableName, test.colCount, colCount)
		}

		rowCount := table.RowCount()
		if rowCount != test.rowCount {
			t.Errorf("test[%d]: expecting [%s] rowCount %d, not %d\n", i, test.tableName, test.rowCount, rowCount)
		}
	}
}

func TestReadString6(t *testing.T) {
	tableSet, err := NewTableSetFromString(
		`[TableStruct]
		i int = 42
		j int = 44

		[Empty]

		`)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		tableName string
		colCount  int
		rowCount  int
		succeeds  bool
	}{
		{"TableStruct", 2, 1, true},
	}

	for i, test := range tests {
		table, err := tableSet.Table(test.tableName)
		if err != nil {
			t.Errorf("[%d] %v", i, err)
		}
		if (err == nil) != test.succeeds {
			t.Errorf("test[%d]: err == %v but expecting succeeds == %t", i, err, test.succeeds)
		}

		colCount := table.ColCount()
		if colCount != test.colCount {
			t.Errorf("test[%d]: expecting [%s] colCount %d, not %d\n", i, test.tableName, test.colCount, colCount)
		}

		rowCount := table.RowCount()
		if rowCount != test.rowCount {
			t.Errorf("test[%d]: expecting [%s] rowCount %d, not %d\n", i, test.tableName, test.rowCount, rowCount)
		}
	}
}

func TestReadString7(t *testing.T) {
	_, err := NewTableSetFromString(
		`[TableStruct]
		i int = 42
		j int = 44
		# Expecting more structs or a blank line.
		X Y Z

		[Empty]

	`)
	if err == nil {
		t.Error(err)
	}
}

func TestReadString8(t *testing.T) {
	_, err := NewTableSetFromString(
		`[TableShaped]
		X Y Z
		# Expecting col types, not structs.
		i int = 42
		j int = 44

		[Empty]

		`)
	if err == nil {
		t.Error(err)
	}
}

// Testing struct using = with zero rows.
func TestReadString9(t *testing.T) {
	table, err := NewTableFromString(
		`[TableStruct]
		i int
		j int
	`)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		tableName string
		colCount  int
		rowCount  int
		succeeds  bool
	}{
		{"TableStruct", 2, 0, true},
	}

	for i, test := range tests {

		colCount := table.ColCount()
		if colCount != test.colCount {
			t.Errorf("test[%d]: expecting [%s] colCount %d, not %d\n", i, test.tableName, test.colCount, colCount)
		}

		rowCount := table.RowCount()
		if rowCount != test.rowCount {
			t.Errorf("test[%d]: expecting [%s] rowCount %d, not %d\n", i, test.tableName, test.rowCount, rowCount)
		}
	}
}

// 02.05.2017
// Testing struct without = having zero rows.
// This is a struct format change to have = only if there is a value following it.
func TestReadString10(t *testing.T) {
	table, err := NewTableFromString(
		`[BlankTableStruct]
		i int
		j int
	`)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		tableName string
		colCount  int
		rowCount  int
		succeeds  bool
	}{
		{"BlankTableStruct", 2, 0, true},
	}

	for i, test := range tests {

		colCount := table.ColCount()
		if colCount != test.colCount {
			t.Errorf("test[%d]: expecting [%s] colCount %d, not %d\n", i, test.tableName, test.colCount, colCount)
		}

		rowCount := table.RowCount()
		if rowCount != test.rowCount {
			t.Errorf("test[%d]: expecting [%s] rowCount %d, not %d\n", i, test.tableName, test.rowCount, rowCount)
		}
	}
// fmt.Println(table)
}

// 02.05.2017
// Testing struct with name type = value
// This is a struct format change to have = only if there is a value following it.
func TestReadString11(t *testing.T) {
	table, err := NewTableFromString(
		`[ValuesTableStruct]
		i int = 1
		j int = 2
		s string = "ABC"
	`)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		tableName string
		colCount  int
		rowCount  int
		succeeds  bool
	}{
		{"ValuesTableStruct", 3, 1, true},
	}

	for i, test := range tests {

		colCount := table.ColCount()
		if colCount != test.colCount {
			t.Errorf("test[%d]: expecting [%s] colCount %d, not %d\n", i, test.tableName, test.colCount, colCount)
		}

		rowCount := table.RowCount()
		if rowCount != test.rowCount {
			t.Errorf("test[%d]: expecting [%s] rowCount %d, not %d\n", i, test.tableName, test.rowCount, rowCount)
		}
	}
// fmt.Println(table)
}

// 02.05.2017
// Testing struct with name type = value
// This is a struct format change to have = only if there is a value following it.
func TestReadString12(t *testing.T) {
	_, err := NewTableFromString(
		`[InvalidTableStruct]
		i int =
		j int =
		s string =
	`)
	if err == nil {
		t.Error(err)
	}
}

// 02/09.2017
// Testing table with slice of uint: []uint
func TestReadString13(t *testing.T) {
	var err error

	_, err = NewTableFromString(
		`[TableWithUintSlice]
		uintNums
		[]uint8
		[0 1 255 3 4]
	`)
	if err != nil {
		t.Error(err)
	}

	_, err = NewTableFromString(
		`[TableWithUintSlice]
		uintNums
		[]uint8
		[0 -1 255 3 4]
	`)
	if err == nil {
		t.Error(err)
	}

	_, err = NewTableFromString(
		`[TableWithByteSlice]
		uintNums
		[]byte
		[0 1 256 3 4]
	`)
	if err == nil {
		t.Error(err)
	}

	_, err = NewTableFromString(
		`[TableWithByteSlice]
		uintNums []byte = [0 1 256 3 4]
		i int = 42
		b []byte = [1 1 255]
		u []uint8 = [2 2 255 2]
		f float32 = 32
		x []byte = []
		s string = "In Between ..."
		y []uint8 = []
		j int64 = 99
	`)
	if err == nil {
		t.Error(err)
	}
}

// 03/09.2017
// Testing table with slice of uint: []uint
func TestReadString14(t *testing.T) {
	var err error
	s :=
	`[TableX]
	i   x                   f           bb                  s       b
	int []uint8             float64     []byte              string  byte
	1   [10 11 12 13]       1           [90 81 72 63]       "one"   11
	2   [20 21 22 23]       2           [90 81 72 63]       "two"   22
	3   [30 31 32]          3           [90 81 72]          "three" 33
	4   [40 41 42 43 44]    4           [90 81 72 63 255]   "four"  44
	`
	table, err := NewTableFromString(s)
	if err != nil {
		t.Error(err)
	}

	s =
	`[StructY]
	i int = 42
	bb []byte = [1 1 255]
	u []uint8 = [2 2 255 2]
	f float32 = 32
	x []byte = []
	y []uint8 = []
	b byte = 55
	`
	table, err = NewTableFromString(s)
	if err != nil {
		t.Error(err)
	}

	err = table.AppendCol("bite", "[]byte")
	if err != nil {
		t.Error(err)
	}
}

func ExampleNewTableSet() {
	tableSetName := "MyTableSet"
	tableSet, err := NewTableSet(tableSetName)
	if err != nil {
		log.Println(err)
	}
	tableCount := tableSet.TableCount()
	name := tableSet.Name()
	fmt.Println(tableCount)
	fmt.Println(name)
	// Output:
	// 0
	// MyTableSet
}

func ExampleRound() {
	numberToRound := 12.326
	places := 2 // The rounded fractional part will have 2 decimal places.

	rounded := Round(numberToRound, places)
	fmt.Println(rounded)
	// Output:
	// 12.33
}

func TestRound(t *testing.T) {
	tests := []struct {
		val     float64
		places  int
		rounded float64
	}{
		{12.326, 2, 12.33},
		{12.325, 2, 12.33},
		{12.324, 2, 12.32},
		{12.32, 2, 12.32},
		{12.3, 2, 12.3},
	}

	for i, test := range tests {
		rounded := Round(test.val, test.places)
		if rounded != test.rounded {
			t.Errorf("test[%d]: expecting %f, not %f", i, test.rounded, rounded)
		}
	}
}

func TestSetAndGetFunctions(t *testing.T) {
	var bVal bool
	var byteVal byte   // alias for uint8
	var byteSlice []byte
	var ui8Slice  []uint8
	//	complex128 		// The set of all complex numbers with float64 real and imaginary parts
	//	complex64		// The set of all complex numbers with float32 real and imaginary parts
	var f32Val float32 // The set of all IEEE-754 32-bit floating-point numbers
	var f64Val float64 // The set of all IEEE-754 64-bit floating-point numbers
	var iVal int       // Machine-dependent
	var i16Val int16   // The set of all signed 16-bit integers (-32768 to 32767)
	var i32Val int32   // The set of all signed 32-bit integers (-2147483648 to 2147483647)
	var i64Val int64   // The set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	var i8Val int8     // The set of all signed  8-bit integers (-128 to 127)
	//	rune 			// alias for int32
	var uiVal uint     // Machine-dependent
	var ui16Val uint16 // The set of all unsigned 16-bit integers (0 to 65535)
	var ui32Val uint32 // The set of all unsigned 32-bit integers (0 to 4294967295)
	var ui64Val uint64 // The set of all unsigned 64-bit integers (0 to 18446744073709551615)
	var ui8Val uint8   // The set of all unsigned  8-bit integers (0 to 255)
	var sVal string

	var err error
	var table *Table
	const rowIndex = 0
	var colIndex = 0

	table, err = NewTable("SetAndGet")
	if err != nil {
		t.Error(err)
	}

	err = table.AppendRow()
	if err != nil {
		t.Error(err)
	}

	// Note: Tests are collected inside code blocks for human readability.

	{ // bool tests

		err = table.AppendCol("bVal", "bool")
		if err != nil {
			t.Error(err)
		}

		expected := true
		err = table.SetBool("bVal", rowIndex, expected)
		if err != nil {
			t.Error(err)
		}
		bVal, err = table.GetBool("bVal", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if bVal != expected {
			t.Errorf("expecting GetBool() value %t, not %t\n", expected, bVal)
		}

		expected = false
		err = table.SetBoolByColIndex(colIndex, rowIndex, expected)
		if err != nil {
			t.Error(err)
		}
		bVal, err = table.GetBoolByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if bVal != expected {
			t.Errorf("expecting GetBoolByColIndex() value %t, not %t\n", expected, bVal)
		}
	}

	{ // float32 tests

		err = table.AppendCol("f32Val", "float32")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetFloat32("f32Val", rowIndex, 55.1)
		if err != nil {
			t.Error(err)
		}
		f32Val, err = table.GetFloat32("f32Val", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if f32Val != 55.1 {
			t.Errorf("expecting GetFloat32() value %f, not %f\n", 55.1, f32Val)
		}

		err = table.SetFloat32ByColIndex(colIndex, rowIndex, 66.1)
		if err != nil {
			t.Error(err)
		}
		f32Val, err = table.GetFloat32ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if f32Val != 66.1 {
			t.Errorf("expecting GetFloat32ByColIndex() value %f, not %f\n", 66.1, f32Val)
		}
	}

	{ // float64 tests

		err = table.AppendCol("f64Val", "float64")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetFloat64("f64Val", rowIndex, 88.1)
		if err != nil {
			t.Error(err)
		}
		f64Val, err = table.GetFloat64("f64Val", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if f64Val != 88.1 {
			t.Errorf("expecting GetFloat64() value %f, not %f\n", 88.1, f64Val)
		}

		err = table.SetFloat64ByColIndex(colIndex, rowIndex, 77.1)
		if err != nil {
			t.Error(err)
		}
		f64Val, err = table.GetFloat64ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if f64Val != 77.1 {
			t.Errorf("expecting GetFloat64ByColIndex() value %f, not %f\n", 77.1, f64Val)
		}
	}

	{ // int tests

		err = table.AppendCol("iVal", "int")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetInt("iVal", rowIndex, 55)
		if err != nil {
			t.Error(err)
		}
		iVal, err = table.GetInt("iVal", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if iVal != 55 {
			t.Errorf("expecting GetInt() value %d, not %d\n", 55, iVal)
		}

		err = table.SetIntByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Error(err)
		}
		iVal, err = table.GetIntByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if iVal != 66 {
			t.Errorf("expecting GetIntByColIndex() value %d, not %d\n", 66, iVal)
		}
	}

	{ // int16 tests

		err = table.AppendCol("i16Val", "int16")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetInt16("i16Val", rowIndex, 55)
		if err != nil {
			t.Error(err)
		}
		i16Val, err = table.GetInt16("i16Val", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if i16Val != 55 {
			t.Errorf("expecting GetInt16() value %d, not %d\n", 55, i16Val)
		}

		err = table.SetInt16ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Error(err)
		}
		i16Val, err = table.GetInt16ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if i16Val != 66 {
			t.Errorf("expecting GetInt16ByColIndex() value %d, not %d\n", 66, i16Val)
		}
	}

	{ // int32 tests

		err = table.AppendCol("i32Val", "int32")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetInt32("i32Val", rowIndex, 55)
		if err != nil {
			t.Error(err)
		}
		i32Val, err = table.GetInt32("i32Val", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if i32Val != 55 {
			t.Errorf("expecting GetInt32() value %d, not %d\n", 55, i32Val)
		}

		err = table.SetInt32ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Error(err)
		}
		i32Val, err = table.GetInt32ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if i32Val != 66 {
			t.Errorf("expecting GetInt32ByColIndex() value %d, not %d\n", 66, i32Val)
		}
	}

	{ // int64 tests

		err = table.AppendCol("i64Val", "int64")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetInt64("i64Val", rowIndex, 55)
		if err != nil {
			t.Error(err)
		}
		i64Val, err = table.GetInt64("i64Val", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if i64Val != 55 {
			t.Errorf("expecting GetInt64() value %d, not %d\n", 55, i64Val)
		}

		err = table.SetInt64ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Error(err)
		}
		i64Val, err = table.GetInt64ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if i64Val != 66 {
			t.Errorf("expecting GetInt64ByColIndex() value %d, not %d\n", 66, i64Val)
		}
	}

	{ // int8 tests

		err = table.AppendCol("i8Val", "int8")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetInt8("i8Val", rowIndex, 55)
		if err != nil {
			t.Error(err)
		}
		i8Val, err = table.GetInt8("i8Val", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if i8Val != 55 {
			t.Errorf("expecting GetInt8() value %d, not %d\n", 55, i8Val)
		}

		err = table.SetInt8ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Error(err)
		}
		i8Val, err = table.GetInt8ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if i8Val != 66 {
			t.Errorf("expecting GetInt8ByColIndex() value %d, not %d\n", 66, i8Val)
		}
	}

	{ // uint tests

		err = table.AppendCol("uiVal", "uint")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetUint("uiVal", rowIndex, 55)
		if err != nil {
			t.Error(err)
		}
		uiVal, err = table.GetUint("uiVal", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if uiVal != 55 {
			t.Errorf("expecting GetUint() value %d, not %d\n", 55, uiVal)
		}

		err = table.SetUintByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Error(err)
		}
		uiVal, err = table.GetUintByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if uiVal != 66 {
			t.Errorf("expecting GetUintByColIndex() value %d, not %d\n", 66, uiVal)
		}
	}

	{ // uint16 tests

		err = table.AppendCol("ui16Val", "uint16")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetUint16("ui16Val", rowIndex, 55)
		if err != nil {
			t.Error(err)
		}
		ui16Val, err = table.GetUint16("ui16Val", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if ui16Val != 55 {
			t.Errorf("expecting GetUint16() value %d, not %d\n", 55, ui16Val)
		}

		err = table.SetUint16ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Error(err)
		}
		ui16Val, err = table.GetUint16ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if ui16Val != 66 {
			t.Errorf("expecting GetUint16ByColIndex() value %d, not %d\n", 66, ui16Val)
		}
	}

	{ // uint32 tests

		err = table.AppendCol("ui32Val", "uint32")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetUint32("ui32Val", rowIndex, 55)
		if err != nil {
			t.Error(err)
		}
		ui32Val, err = table.GetUint32("ui32Val", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if ui32Val != 55 {
			t.Errorf("expecting GetUint32() value %d, not %d\n", 55, ui32Val)
		}

		err = table.SetUint32ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Error(err)
		}
		ui32Val, err = table.GetUint32ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if ui32Val != 66 {
			t.Errorf("expecting GetUint32ByColIndex() value %d, not %d\n", 66, ui32Val)
		}
	}

	{ // uint64 tests

		err = table.AppendCol("ui64Val", "uint64")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetUint64("ui64Val", rowIndex, 55)
		if err != nil {
			t.Error(err)
		}
		ui64Val, err = table.GetUint64("ui64Val", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if ui64Val != 55 {
			t.Errorf("expecting GetUint64() value %d, not %d\n", 55, ui64Val)
		}

		err = table.SetUint64ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Error(err)
		}
		ui64Val, err = table.GetUint64ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if ui64Val != 66 {
			t.Errorf("expecting GetUint64ByColIndex() value %d, not %d\n", 66, ui64Val)
		}
	}

	{ // uint8 tests

		err = table.AppendCol("ui8Val", "uint8")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetUint8("ui8Val", rowIndex, 55)
		if err != nil {
			t.Error(err)
		}
		ui8Val, err = table.GetUint8("ui8Val", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if ui8Val != 55 {
			t.Errorf("expecting GetUint8() value %d, not %d\n", 55, ui8Val)
		}

		err = table.SetUint8ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Error(err)
		}
		ui8Val, err = table.GetUint8ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if ui8Val != 66 {
			t.Errorf("expecting GetUint8ByColIndex() value %d, not %d\n", 66, ui8Val)
		}
	}

	{ // uint8[] slice tests

		err = table.AppendCol("ui8Slice", "[]uint8")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetByteSlice("ui8Slice", rowIndex, []uint8{0,1,2})
		if err != nil {
			t.Error(err)
		}
		ui8Slice, err = table.GetByteSlice("ui8Slice", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if bytes.Compare(ui8Slice, []uint8{0,1,2}) != 0 {	// Slices not equal.
			t.Errorf("expecting GetByteSlice() value %d, not %d\n", []uint8{0,1,2}, ui8Slice)
		}

		err = table.SetByteSliceByColIndex(colIndex, rowIndex, []uint8{2,4,6})
		if err != nil {
			t.Error(err)
		}
		ui8Slice, err = table.GetByteSliceByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if bytes.Compare(ui8Slice, []uint8{2,4,6}) != 0 {	// Slices not equal.
			t.Errorf("expecting GetByteSliceByColIndex() value %d, not %d\n", []uint8{2,4,6}, ui8Slice)
		}
	}

	{ // byte tests

		err = table.AppendCol("byteVal", "byte")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetByte("byteVal", rowIndex, 56)
		if err != nil {
			t.Error(err)
		}
		byteVal, err = table.GetByte("byteVal", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if byteVal != 56 {
			t.Errorf("expecting GetByte() value %d, not %d\n", 56, byteVal)
		}

		err = table.SetByteByColIndex(colIndex, rowIndex, 67)
		if err != nil {
			t.Error(err)
		}
		byteVal, err = table.GetByteByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if byteVal != 67 {
			t.Errorf("expecting GetByteByColIndex() value %d, not %d\n", 67, byteVal)
		}
	}

	{ // byte[] slice tests

		err = table.AppendCol("byteSlice", "[]byte")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetByteSlice("byteSlice", rowIndex, []byte{4,5,6})
		if err != nil {
			t.Error(err)
		}
		byteSlice, err = table.GetByteSlice("byteSlice", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if bytes.Compare(byteSlice, []byte{4,5,6}) != 0 {	// Slices not equal.
			t.Errorf("expecting GetByteSlice() value %d, not %d\n", []byte{4,5,6}, byteSlice)
		}

		err = table.SetByteSliceByColIndex(colIndex, rowIndex, []byte{7,8,9})
		if err != nil {
			t.Error(err)
		}
		byteSlice, err = table.GetByteSliceByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if bytes.Compare(byteSlice, []byte{7,8,9}) != 0 {	// Slices not equal.
			t.Errorf("expecting GetByteSliceByColIndex() value %d, not %d\n", []byte{7,8,9}, byteSlice)
		}
	}

	{ // string tests

		err = table.AppendCol("sVal", "string")
		if err != nil {
			t.Error(err)
		}
		colIndex += 1

		err = table.SetString("sVal", rowIndex, "55")
		if err != nil {
			t.Error(err)
		}
		sVal, err = table.GetString("sVal", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if sVal != "55" {
			t.Errorf("expecting GetString() value %q, not %q\n", "55", sVal)
		}

		err = table.SetStringByColIndex(colIndex, rowIndex, "66")
		if err != nil {
			t.Error(err)
		}
		sVal, err = table.GetStringByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if sVal != "66" {
			t.Errorf("expecting GetStringByColIndex() value %q, not %q\n", "66", sVal)
		}
	}

//	fmt.Println(table)

	var shape bool
	var expected bool = true
	err = table.SetStructShape(expected)
	if err != nil {
		t.Error(err)
	}
	shape, err = table.IsStructShape()
	if err != nil {
		t.Error(err)
	}
	if shape != expected {
		t.Errorf("expecting [%s].IsStructShape() value %t, not %t\n", table.Name(), expected, shape)
	}

//	fmt.Println(table)

	expected = false
	err = table.SetStructShape(expected)
	if err != nil {
		t.Error(err)
	}
	shape, err = table.IsStructShape()
	if err != nil {
		t.Error(err)
	}
	if shape != expected {
		t.Errorf("expecting [%s].IsStructShape() value %t, not %t\n", table.Name(), expected, shape)
	}
}

func TestSetIntegerMinAndMax(t *testing.T) {
	var err error

	// For testing machine-dependent types
	var intBits int = strconv.IntSize // uint and int are the same bit size.
	var intMinVal int64
	var intMaxVal uint64
	var uintMaxVal uint64
	switch intBits {
	case 32:
		intMinVal = math.MinInt32
		intMaxVal = math.MaxInt32
		uintMaxVal = math.MaxUint32
	case 64:
		intMinVal = math.MinInt64
		intMaxVal = math.MaxInt64
		uintMaxVal = math.MaxUint64
	default:
		msg := fmt.Sprintf("CHECK int or uint ON THIS SYSTEM: Unknown int size: %d bits", intBits)
		t.Errorf(msg)
	}

	var tests = []struct {
		input string
		valid bool
	}{

		// uint8
		{`
			[uint8]
			i uint8 = 0`,
			true,
		},
		{`
			[uint8]
			i uint8 = 255`,
			true,
		},
		{`
			[uint8]
			i uint8 = -1`,
			false,
		},
		{`
			[uint8]
			i uint8 = 256`,
			false,
		},

		// uint16
		{`
			[uint16]
			i uint16 = 0`,
			true,
		},
		{`
			[uint16]
			i uint16 = 65535`,
			true,
		},
		{`
			[uint16]
			i uint16 = -1`,
			false,
		},
		{`
			[uint16]
			i uint16 = 65536`,
			false,
		},

		// uint32
		{`
			[uint32]
			i uint32 = 0`,
			true,
		},
		{`
			[uint32]
			i uint32 = 4294967295`,
			true,
		},
		{`
			[uint32]
			i uint32 = -1`,
			false,
		},
		{`
			[uint32]
			i uint32 = 4294967296`,
			false,
		},

		// uint64
		{`
			[uint64]
			i uint64 = 0`,
			true,
		},
		{`
			[uint64]
			i uint64 = 18446744073709551615`,
			true,
		},
		{`
			[uint64]
			i uint64 = -1`,
			false,
		},
		{`
			[uint64]
			i uint64 = 18446744073709551616`,
			false,
		},

		// uint
		{`
			[uint]
			i uint = 0`,
			true,
		},
		{fmt.Sprintf(`
			[uint]
			i uint = %d`, uintMaxVal),
			true,
		},
		{`
			[uint]
			i uint = -1`,
			false,
		},
		// Note: Cannot easily test machine-dependent types outside range values (except for uint 0 and -1), so skipping them.

		// int8
		{`
			[int8]
			i int8 = -128`,
			true,
		},
		{`
			[int8]
			i int8 = 127`,
			true,
		},
		{`
			[int8]
			i int8 = -129`,
			false,
		},
		{`
			[int8]
			i int8 = 128`,
			false,
		},

		// int16
		{`
			[int16]
			i int16 = -32768`,
			true,
		},
		{`
			[int16]
			i int16 = 32767`,
			true,
		},
		{`
			[int16]
			i int16 = -32769`,
			false,
		},
		{`
			[int16]
			i int16 = 32768`,
			false,
		},

		// int32
		{`
			[int32]
			i int32 = -2147483648`,
			true,
		},
		{`
			[int32]
			i int32 = 2147483647`,
			true,
		},
		{`
			[int32]
			i int32 = -2147483649`,
			false,
		},
		{`
			[int32]
			i int32 = 2147483648`,
			false,
		},

		// int64
		{`
			[int64]
			i int64 = -9223372036854775808`,
			true,
		},
		{`
			[int64]
			i int64 = 9223372036854775807`,
			true,
		},
		{`
			[int64]
			i int64 = -9223372036854775809`,
			false,
		},
		{`
			[int64]
			i int64 = 9223372036854775808`,
			false,
		},

		// int
		{fmt.Sprintf(`
			[int]
			i int = %d`, intMinVal),
			true,
		},
		{fmt.Sprintf(`
			[int]
			i int = %d`, intMaxVal),
			true,
		},
		// Note: Cannot easily test machine-dependent types outside range values here, so skipping them.
		//       See TestSetIntegerMinAndMaxMachineDependent()
	}

	for i, test := range tests {
		_, err = NewTableSetFromString(test.input)
		if err == nil != test.valid {
			switch test.valid {
			case true:
				t.Errorf("test[%d]: %v", i, err)
			case false:
				t.Errorf("test[%d]: NewTableSetFromString(): expecting this input to fail with a range error: %s", i, test.input)
			}
		}
	}
}

func TestSetIntegerMinAndMaxMachineDependent(t *testing.T) {
	var err error

	type testCase struct {
		input string
		valid bool
	}
	var tests []testCase

	// All of these tests are of out-of-range values (1 too small or 1 too large) which should fail when parsed.
	// NOTE: Only half of these tests will be executed. They are machine dependent: 32-bit OR 64-bit machines.

	// For testing machine-dependent types
	var intBits int = strconv.IntSize // uint and int are the same bit size.
	switch intBits {
	case 32: // NOTE: This will be executed on 32-bit machines ONLY.
		tests = append(tests, testCase{`
					[uint]
					i uint = 4294967296`,
			false,
		},
		)
		tests = append(tests, testCase{`
					[int]
					i int = -2147483649`,
			false,
		},
		)
		tests = append(tests, testCase{`
					[int]
					i int = 2147483648`,
			false,
		},
		)
	case 64: // NOTE: This will be executed on 32-bit machines ONLY.
		tests = append(tests, testCase{`
					[uint]
					i uint = 18446744073709551616`,
			false,
		},
		)
		tests = append(tests, testCase{`
					[int]
					i int = -9223372036854775809`,
			false,
		},
		)
		tests = append(tests, testCase{`
					[int]
					i int = 9223372036854775808`,
			false,
		},
		)
	default:
		msg := fmt.Sprintf("CHECK int or uint ON THIS SYSTEM: Unknown int size: %d bits", intBits)
		t.Errorf(msg)
	}

	for i, test := range tests {
		_, err = NewTableSetFromString(test.input)
		if err == nil != test.valid {
			switch test.valid {
			case true:
				t.Errorf("test[%d]: %v", i, err)
			case false:
				t.Errorf("test[%d]: NewTableSetFromString(): expecting this input to fail with a range error: %s", i, test.input)
			}
		}
	}
}

var tableSetString string = `
	[sable_fur]
    i   s       f           b
    int string  float64     bool
    1   "abc"   2.3         true
    2   "xyz"   4.5         false
    3   "ssss"  4.9         false

    [my_struct_table]
    i int    = 9223372036854775807
    i2 int64 = 9223372036854775807
    s string = "forty-two"
    f int8 = 42
    u uint8  = 255
    i81 int8 = 127
    i82 int8 = -128
    i161 int16 = 32767
    i162 int16 = -32768
    i321 int8 = 127
    i322 int8 = -128
    ui uint16 = 65535
    `

func BenchmarkNewTableSetFromString(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		_, err = NewTableSetFromString(tableSetString)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkTableSetToString(b *testing.B) {
	// Set up for benchmark.
	tableSet, err := NewTableSetFromString(tableSetString)
	if err != nil {
		b.Error(err)
	}

	var _ string
	for i := 0; i < b.N; i++ {
		_ = tableSet.String()
	}
}

var planets_padded = `[planets_padded]
name         mass distance moons index mnemonic
string    float64  float64   int   int string
"Mercury"   0.055      0.4     0     0 "my"
"Venus"     0.815      0.7     0     1 "very"
"Earth"     1.0        1.0     1     2 "elegant"
"Mars"      0.107      1.5     2     3 "mother"
"Jupiter" 318.0        5.2    67     4 "just"
"Saturn"   95.0       29.4    62     5 "sat"
"Uranus"   15.0       84.0    27     6 "upon"
"Neptune"  17.0      164.0    13     7 "nine ... porcupines"
`

var planets_unpadded = `[planets_unpadded]
name mass distance moons index mnemonic
string float64 float64 int int string
"Mercury" 0.055 0.4 0 0 "my"
"Venus" 0.815 0.7 0 1 "very"
"Earth" 1 1 1 2 "elegant"
"Mars" 0.107 1.5 2 3 "mother"
"Jupiter" 318 5.2 67 4 "just"
"Saturn" 95 29.4 62 5 "sat"
"Uranus" 15 84 27 6 "upon"
"Neptune" 17 164 13 7 "nine ... porcupines"
`

func BenchmarkNewTableSetFromString_padded(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		_, err = NewTableSetFromString(planets_padded)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkNewTableSetFromString_unpadded(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		_, err = NewTableSetFromString(planets_unpadded)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkTableSetToString_padded(b *testing.B) {
	// Set up for benchmark.
	// Note: It's irrelevant whether the input string is padded.
	tableSet, err := NewTableSetFromString(planets_padded)
	if err != nil {
		b.Error(err)
	}

	var _ string
	for i := 0; i < b.N; i++ {
		_ = tableSet.StringPadded()
	}
}

func BenchmarkTableSetToString_unpadded(b *testing.B) {
	// Set up for benchmark.
	// Note: It's irrelevant whether the input string is padded.
	tableSet, err := NewTableSetFromString(planets_unpadded)
	if err != nil {
		b.Error(err)
	}

	var _ string
	for i := 0; i < b.N; i++ {
		_ = tableSet.StringUnpadded()
	}
}

func BenchmarkGobEncode(b *testing.B) {
	// Set up for benchmark.
	tableSet, err := NewTableSetFromString(tableSetString)
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		_, err := tableSet.GobEncode()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGobDecode(b *testing.B) {
	// Set up for benchmark.
	tableSet, err := NewTableSetFromString(tableSetString)
	if err != nil {
		b.Error(err)
	}

	// Set up for benchmark.
	gobEncodedTableSet, err := tableSet.GobEncode()
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		_, err := GobDecodeTableSet(gobEncodedTableSet)
		if err != nil {
			b.Error(err)
		}
	}
}

func TestIsNumericColType(t *testing.T) {
	tableString := `
    [table]
	F_bool    bool
	F_string  string
	F__byte   []byte
	F__uint8  []uint8
	T_float32 float32
	T_float64 float64
	T_int     int
	T_int16   int16
	T_int32   int32
	T_int64   int64
	T_int8    int8
	T_uint    uint
	T_uint16  uint16
	T_uint32  uint32
	T_uint64  uint64
	T_uint8   uint8
	T_byte    byte
    `

	tableSet, err := NewTableSetFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	table, err := tableSet.Table("table")
	if err != nil {
		t.Error(err)
	}

	for colIndex := 0; colIndex < table.ColCount(); colIndex++ {

		colName, err := table.ColName(colIndex)
		if err != nil {
			t.Error(err)
		}

		colType, err := table.ColTypeByColIndex(colIndex)
		if err != nil {
			t.Error(err)
		}

		isNumeric, _ := IsNumericColType(colType)

		hasPrefixT := strings.HasPrefix(colName, "T_")
		hasPrefixF := strings.HasPrefix(colName, "F_")
		if !hasPrefixT && !hasPrefixF {
			t.Errorf("expecting col name %s to have prefix \"T_\" or \"F_\" but found: %q", colName, colName)
		}

		if isNumeric != hasPrefixT {
			err := fmt.Errorf("col %s type %s unexpected IsNumeric: %t", colName, colType, isNumeric)
			t.Error(err)
		}
	}
}

func TestAppendRow(t *testing.T) {
	tableString := `
    [table]
	F_bool bool
	F_string string
	T_float32 float32
	T_float64 float64
	T_int int
	T_int16 int16
	T_int32 int32
	T_int64 int64
	T_int8 int8
	T_uint uint
	T_uint16 uint16
	T_uint32 uint32
	T_uint64 uint64
	T_uint8 uint8
    `

	tableSet, err := NewTableSetFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	table, err := tableSet.Table("table")
	if err != nil {
		t.Error(err)
	}

	rowCount := table.RowCount()
	if rowCount != 0 {
		t.Errorf("expecting row count of 0, not: %d", rowCount)
	}

	err = table.AppendRow()
	if err != nil {
		t.Error(err)
	}

	rowCount = table.RowCount()
	if rowCount != 1 {
		t.Errorf("expecting row count of 1, not: %d", rowCount)
	}

	err = table.DeleteRow(0)
	if err != nil {
		t.Error(err)
	}

	rowCount = table.RowCount()
	if rowCount != 0 {
		t.Errorf("expecting row count of 0, not: %d", rowCount)
	}
}

func TestColCount(t *testing.T) {
	tableString := `
    [table]
	F_bool bool
	F_string string
	T_float32 float32
	T_float64 float64
	T_int int
	T_int16 int16
	T_int32 int32
	T_int64 int64
	T_int8 int8
	T_uint uint
	T_uint16 uint16
	T_uint32 uint32
	T_uint64 uint64
	T_uint8 uint8
    `

	tableSet, err := NewTableSetFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	table, err := tableSet.Table("table")
	if err != nil {
		t.Error(err)
	}

	const initialColCount = 14

	colCount := table.ColCount()
	if colCount != initialColCount {
		t.Errorf("expecting col count of %d, not: %d", initialColCount, colCount)
	}

	err = table.AppendCol("ExtraCol", "bool")
	if err != nil {
		t.Error(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount+1 {
		t.Errorf("expecting col count of %d, not: %d", initialColCount+1, colCount)
	}

	lastCol := colCount - 1
	err = table.DeleteColByColIndex(lastCol)
	if err != nil {
		t.Error(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount {
		t.Errorf("expecting col count of %d, not: %d", initialColCount, colCount)
	}

	err = table.AppendCol("AnotherCol", "string")
	if err != nil {
		t.Error(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount+1 {
		t.Errorf("expecting col count of %d, not: %d", initialColCount+1, colCount)
	}

	err = table.DeleteCol("AnotherCol")
	if err != nil {
		t.Error(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount {
		t.Errorf("expecting col count of %d, not: %d", initialColCount, colCount)
	}
}

func TestDeleteRow(t *testing.T) {
	tableString := `
	[table]
	item
	int
	0
	1
	2
	3
	4
	5
	6
	7
	8
	9
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	initialRowCount := table.RowCount()

	const deleteRow = 4

	err = table.DeleteRow(deleteRow)
	if err != nil {
		t.Error(err)
	}

	rowCount := table.RowCount()
	if rowCount != initialRowCount-1 {
		t.Errorf("expecting 1 row less than %d after DeleteRow(%d) but found %d", initialRowCount, deleteRow, rowCount)
	}

	// fmt.Println(table)

	for i := 0; i < table.RowCount(); i++ {
		item, err := table.GetInt("item", i)
		if err != nil {
			t.Error(err)
		}
		if item == deleteRow {
			t.Errorf("expecting to NOT find item %d after DeleteRow(%d) but found %d", deleteRow, deleteRow, deleteRow)
		}
	}

}

func TestDeleteRows(t *testing.T) {
	tableString := `
	[items]
	item
	int
	0
	1
	2
	3
	4
	5
	6
	7
	8
	9
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	initialRowCount := table.RowCount()

	// Test invalid row index range: first greater than last
	err = table.DeleteRows(5, 4)
	if err == nil {
		t.Error(err)
	}

	// Test invalid row index range: first below zero
	err = table.DeleteRows(-1, 4)
	if err == nil {
		t.Error(err)
	}

	// Test invalid row index range: last above initialRowCount-1
	err = table.DeleteRows(0, initialRowCount)
	if err == nil {
		t.Error(err)
	}

	var first int
	var last int
	var items int
	var rowCount int

	// Test delete 1 item: 4
	first = 4
	last = 4
	items = last - first + 1
	err = table.DeleteRows(first, last)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(table)
	rowCount = table.RowCount()
	if rowCount != initialRowCount-items {
		t.Errorf("expecting %d row less than %d after DeleteRows(%d, %d) but found %d", items, initialRowCount, first, last, rowCount)
	}
	for i := 0; i < table.RowCount(); i++ {
		item, err := table.GetInt("item", i)
		if err != nil {
			t.Error(err)
		}
		if item == first {
			t.Errorf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d", first, first, last, first)
		}
	}

	// Test delete 2 items: 4 to 5
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	first = 4
	last = 5
	items = last - first + 1
	err = table.DeleteRows(first, last)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(table)
	rowCount = table.RowCount()
	if rowCount != initialRowCount-items {
		t.Errorf("expecting %d row less than %d after DeleteRows(%d, %d) but found %d", items, initialRowCount, first, last, rowCount)
	}
	for i := 0; i < table.RowCount(); i++ {
		item, err := table.GetInt("item", i)
		if err != nil {
			t.Error(err)
		}
		if item == first {
			t.Errorf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d", first, first, last, first)
		}
		if item == last {
			t.Errorf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d", last, first, last, last)
		}
	}

	// Test delete 6 items: 4 to 9
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	first = 6
	last = 9
	items = last - first + 1
	err = table.DeleteRows(first, last)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(table)
	rowCount = table.RowCount()
	if rowCount != initialRowCount-items {
		t.Errorf("expecting %d row less than %d after DeleteRows(%d, %d) but found %d", items, initialRowCount, first, last, rowCount)
	}
	for i := 0; i < table.RowCount(); i++ {
		item, err := table.GetInt("item", i)
		if err != nil {
			t.Error(err)
		}
		if item == first {
			t.Errorf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d", first, first, last, first)
		}
		if item == last {
			t.Errorf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d", last, first, last, last)
		}
	}

	// Test delete 3 items: 0 to 9
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	first = 0
	last = 2
	items = last - first + 1
	err = table.DeleteRows(first, last)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(table)
	rowCount = table.RowCount()
	if rowCount != initialRowCount-items {
		t.Errorf("expecting %d row less than %d after DeleteRows(%d, %d) but found %d", items, initialRowCount, first, last, rowCount)
	}
	for i := 0; i < table.RowCount(); i++ {
		item, err := table.GetInt("item", i)
		if err != nil {
			t.Error(err)
		}
		if item == first {
			t.Errorf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d", first, first, last, first)
		}
		if item == last {
			t.Errorf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d", last, first, last, last)
		}
	}
}

func ExampleNewTableFromString_struct() {
	// A table literal. Sometimes easier than constructing a table programmatically.
	tableString := `[MyTable]
		MyBool bool = true
		MyString string = "The answer to life, the universe and everything is forty-two."
		MyInt int = 42`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	// Print the table in its original struct shape.
	fmt.Println(table)

	// Now change its shape to tabular.
	err = table.SetStructShape(false)
	if err != nil {
		log.Println(err)
	}

	// The table is now printed as a single row of data.
	fmt.Println(table)

	// Note: The struct/tabular shape is for readability and has no impact on its internal structure.

	// Output:
	// [MyTable]
	// MyBool bool = true
	// MyString string = "The answer to life, the universe and everything is forty-two."
	// MyInt int = 42
	//
	// [MyTable]
	// MyBool MyString                                                        MyInt
	// bool   string                                                            int
	// true   "The answer to life, the universe and everything is forty-two."    42
}

func ExampleNewTableFromFile() {
	tableString := `
	[MyTable]
	MyBool bool = true
	MyString string = "The answer to life, the universe and everything"
	MyInt int = 42
	`

	table1, err := NewTableFromStringByTableName(tableString, "MyTable")
	if err != nil {
		log.Println(err)
	}

	// For testing, we need to write this out to a file so we can read it back.
	fileName := "ExampleNewTableFromFile.txt"
	err = table1.WriteFile(fileName, 0644)
	if err != nil {
		log.Println(err)
	}

	table2, err := NewTableFromFile(fileName)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table2)

	err = table2.SetStructShape(false)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table2)

	// Output:
	// [MyTable]
	// MyBool bool = true
	// MyString string = "The answer to life, the universe and everything"
	// MyInt int = 42
	//
	// [MyTable]
	// MyBool MyString                                          MyInt
	// bool   string                                              int
	// true   "The answer to life, the universe and everything"    42
}

func ExampleNewTableFromFileByTableName() {

	/*
		NewTableFromFileByTableName() is for when you want just one table from
		(possibly) multiple tables in a file, and you don't want to bother with
		NewTableSetFromFile().

		See also NewTableFromStringByTableName().
	*/

	tableSetString := `
	[MyTable]
	MyBool bool = true
	MyString string = "The answer to life, the universe and everything"
	MyInt int = 42

	[Fred]
	i
	int
	`

	// For testing, we need to write this out to a file so we can read it back.
	tableSet, err := NewTableSetFromString(tableSetString)
	if err != nil {
		log.Println(err)
	}
	fileName := "ExampleNewTableFromFileByTableName.txt"
	err = tableSet.WriteFile(fileName, 0644)
	if err != nil {
		log.Println(err)
	}

	table, err := NewTableFromFileByTableName(fileName, "MyTable")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	// Output:
	// [MyTable]
	// MyBool bool = true
	// MyString string = "The answer to life, the universe and everything"
	// MyInt int = 42
}

func ExampleTable_DeleteRows() {
	tableString := `
	[items]
	item
	int
	0
	1
	2
	3
	4
	5
	6
	7
	8
	9
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	err = table.DeleteRows(4, 6)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	// Output:
	// [items]
	// item
	//  int
	//    0
	//    1
	//    2
	//    3
	//    4
	//    5
	//    6
	//    7
	//    8
	//    9
	//
	// [items]
	// item
	//  int
	//    0
	//    1
	//    2
	//    3
	//    7
	//    8
	//    9
}

func ExampleTable_JoinColVals() {

	tableString := `
	[commands]
	command
	string
	"echo myfile"
	"wc -l"
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	joined, err := table.JoinColVals("command", " | ")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(joined)

	// Output:
	// [commands]
	// command
	// string
	// "echo myfile"
	// "wc -l"
	//
	// echo myfile | wc -l
}

func ExampleTable_JoinColValsByColIndex() {

	tableString := `
	[commands]
	command
	string
	"echo myfile"
	"wc -l"
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	colIndex := 0
	joined, err := table.JoinColValsByColIndex(colIndex, " | ")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(joined)

	// Output:
	// [commands]
	// command
	// string
	// "echo myfile"
	// "wc -l"
	//
	// echo myfile | wc -l
}

func TestGetValAsString(t *testing.T) {
	tableString := `
	[table]
	s string = "Fred"
	t bool = true
	i int = 23
	f float64 = 55.5
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)
	}

	var expecting string
	var found string

	expecting = "Fred"
	found, err = table.GetValAsString("s", 0)
	if err != nil {
		t.Error(err)
	}
	if found != expecting {
		t.Errorf("expecting %s but found: %s", expecting, found)
	}

	expecting = "true"
	found, err = table.GetValAsString("t", 0)
	if err != nil {
		t.Error(err)
	}
	if found != expecting {
		t.Errorf("expecting %s but found: %s", expecting, found)
	}

	expecting = "23"
	found, err = table.GetValAsString("i", 0)
	if err != nil {
		t.Error(err)
	}
	if found != expecting {
		t.Errorf("expecting %s but found: %s", expecting, found)
	}

	expecting = "55.5"
	found, err = table.GetValAsString("f", 0)
	if err != nil {
		t.Error(err)
	}
	if found != expecting {
		t.Errorf("expecting %s but found: %s", expecting, found)
	}
}

func TestTableSet_FileName(t *testing.T) {
	tableString := `
		[table]
		s string = "Fred"
		t bool = true
		i int = 23
		f float64 = 55.5
	`

	// For testing, we need to write this out to a file so we can read it back.
	actualFileName := funcName() + ".txt"
	err := ioutil.WriteFile(actualFileName, []byte(tableString), 0644)
	if err != nil {
		t.Error(err)
	}

	tables, err := NewTableSetFromFile(actualFileName)
	if err != nil {
		t.Error(err)
	}

	fileName := tables.FileName()
	if fileName != actualFileName {
		t.Errorf("Expecting FileName() = %q but found %q", actualFileName, fileName)
	}
}

func TestTableSet_SetName(t *testing.T) {
	expected := ""
	tableSet, err := NewTableSet(expected)
	if err != nil {
		t.Error(err)
	}

	tableSetName := tableSet.Name()
	if tableSetName != expected {
		t.Errorf("Expecting tableSetName = %q but found %q", expected, tableSetName)
	}

	
	expected = "Musk"
	tableSet.SetName(expected)
	tableSetName = tableSet.Name()
	if tableSetName != expected {
		t.Errorf("Expecting tableSetName = %q but found %q", expected, tableSetName)
	}
}

func TestTable_SetName(t *testing.T) {
	expected := "InitialName"
	table, err := NewTable(expected)
	if err != nil {
		t.Error(err)
	}

	tableName := table.Name()
	if tableName != expected {
		t.Errorf("Expecting tableName = %q but found %q", expected, tableName)
	}

	
	expected = "Elon"
	if err = table.SetName(expected); err != nil {
		t.Error(err)
	}
	tableName = table.Name()
	if tableName != expected {
		t.Errorf("Expecting tableName = %q but found %q", expected, tableName)
	}
}

func TestMissingValueForType(t *testing.T) {
	var tests = []struct {
		typeName string
		expected bool
	}{
		{"string", false},
		{"bool", false},
		{"int", false},
		{"int32", false},
		{"int64", false},
		{"uint", false},
		{"float32", true},
		{"float64", true},
	}

	for _, test := range tests {

		_, hasMissing := missingValueForType(test.typeName)
		if hasMissing != test.expected {
			t.Errorf("Expecting missingValueForType(%q) = %t but found %t", test.typeName, test.expected, hasMissing)
		}
	}
}

func TestPreNumberOf(t *testing.T) {
	var tests = []struct {
		sNumber string
		expected int
	}{
		{"0.32", 1},
		{"0.64", 1},
		{"-0", 2},
		{"0", 1},
		{"0", 1},
		{"0", 1},
		{"1", 1},
		{"1", 1},
		{"-11", 3},
		{"-11", 3},
		{"1111", 4},
		{"1111", 4},
		{"0.1", 1},
		{"0.1", 1},
		{"0.11", 1},
		{"0.11", 1},
		{"-0.1112", 2},
		{"0.1112", 1},
		{"0.111236", 1},
		{"0.11123", 1},
		{"NaN", 3},
		{"NaN", 3},
		{"32", 2},
		{"64", 2},
	}

	for _, test := range tests {

		preNumber := preNumberOf(test.sNumber)
		if preNumber != test.expected {
			t.Errorf("Expecting preNumberOf(%q) = %d but found %d", test.sNumber, test.expected, preNumber)
		}
	}
}

func TestPointsOf(t *testing.T) {
	var tests = []struct {
		sNumber string
		expected int
	}{
		{"0.32", 1},
		{"0.64", 1},
		{"-0", 0},
		{"0", 0},
		{"0", 0},
		{"0", 0},
		{"1", 0},
		{"1", 0},
		{"-11", 0},
		{"-11", 0},
		{"1111", 0},
		{"1111", 0},
		{"0.1", 1},
		{"0.1", 1},
		{"0.11", 1},
		{"0.11", 1},
		{"-0.1112", 1},
		{"0.1112", 1},
		{"0.111236", 1},
		{"0.11123", 1},
		{"NaN", 0},
		{"NaN", 0},
		{"32", 0},
		{"64", 0},
	}

	for _, test := range tests {

		points := pointsOf(test.sNumber)
		if points != test.expected {
			t.Errorf("Expecting pointsOf(%q) = %d but found %d", test.sNumber, test.expected, points)
		}
	}
}

func TestPrecisionOf(t *testing.T) {
	var tests = []struct {
		sNumber string
		expected int
	}{
		{"0.32", 2},
		{"0.64", 2},
		{"-0", 0},
		{"0", 0},
		{"0", 0},
		{"0", 0},
		{"1", 0},
		{"1", 0},
		{"-11", 0},
		{"-11", 0},
		{"1111", 0},
		{"1111", 0},
		{"0.1", 1},
		{"0.1", 1},
		{"0.11", 2},
		{"0.11", 2},
		{"-0.1112", 4},
		{"0.1112", 4},
		{"0.111236", 6},
		{"0.11123", 5},
		{"NaN", 0},
		{"NaN", 0},
		{"32", 0},
		{"64", 0},
	}

	for _, test := range tests {

		precision := precisionOf(test.sNumber)
		if precision != test.expected {
			t.Errorf("Expecting precisionOf(%q) = %d but found %d", test.sNumber, test.expected, precision)
		}
	}
}

func TestPadTrailingZeros(t *testing.T) {
	var tests = []struct {
		trailing string
		expected string
	}{
		{  "0.0",    "0.0"},	// Leave as is.
		{  "0.00",   "0.0 "},	// Pad with space.
		{  "0.0000", "0.0   "},	// Pad with spaces.
		{   ".0",     ".0"},	// Pad.
		{   ".10",    ".1 "},	// Pad.
		{   ".100",   ".1  "},	// Pad.
		{  "0",      "0"},		// Integer. Don't pad trailing zeros in int-like floats.
		{ "10",      "10"},		// Integer. Don't pad trailing zeros in int-like floats.
		{"100",      "100"},	// Integer. Don't pad trailing zeros in int-like floats.
	}

	for _, test := range tests {

		trimmed := padTrailingZeros(test.trailing)
		if trimmed != test.expected {
			t.Errorf("Expecting TrimTrailingZeros(%q) = %q but found %q", test.trailing, test.expected, trimmed)
		}
	}
}

func TestTrimTrailingZeros(t *testing.T) {
	var tests = []struct {
		trailing string
		expected string
	}{
		{  "0.0",    "0.0"},	// Leave as is.
		{  "0.00",   "0.0"},	// Trim zeros.
		{  "0.0000", "0.0"},	// Trim zeros.
		{   ".0",     ".0"},	// Trim.
		{   ".10",    ".1"},	// Trim.
		{   ".100",   ".1"},	// Trim.
		{  "0",      "0"},		// Integer. Don't pad trailing zeros in int-like floats.
		{ "10",     "10"},		// Integer. Don't pad trailing zeros in int-like floats.
		{"100",    "100"},		// Integer. Don't pad trailing zeros in int-like floats.
	}

	for _, test := range tests {

		trimmed := trimTrailingZeros(test.trailing)
		if trimmed != test.expected {
			t.Errorf("Expecting TrimTrailingZeros(%q) = %q but found %q", test.trailing, test.expected, trimmed)
		}
	}
}

func TestIsColTypeByColIndex(t *testing.T) {

	tableString :=
	`[ColTypes]
	i int
	b bool
	s string
	f64 float64
	f32 float32
	i32 int32
	u64 uint64
	u   uint
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)	// We're not testing this function.
	}

	var tests = []struct {
		colIndex int
		colType string
		expected bool
	}{
		{0, "int", true},
		{1, "bool", true},
		{2, "string", true},
		{3, "float64", true},
		{4, "float32", true},
		{5, "int32", true},
		{6, "uint64", true},
		{7, "uint", true},
		{0, "int15", false},
		{1, "float64", false},
		{2, "bool", false},
		{3, "float", false},
		{4, "String", false},
		{5, "int16", false},
		{6, "uint8", false},
		{7, "int", false},
	}

	for _, test := range tests {
		isColType, _ := table.IsColTypeByColIndex(test.colIndex, test.colType)
		// Ignore err. Returns err if col type is false.
		if isColType != test.expected {
			t.Errorf("Expecting table.IsColTypeByColIndex(%d, %q) = %t but found %t", test.colIndex, test.colType, test.expected, isColType)
		}
	}
}

func TestIsColType(t *testing.T) {

	tableString :=
	`[ColTypes]
	i int
	b bool
	s string
	f64 float64
	f32 float32
	i32 int32
	u64 uint64
	u   uint
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)	// We're not testing this function.
	}

	var tests = []struct {
		colName string
		colType string
		expected bool
	}{
		{"i", "int", true},
		{"b", "bool", true},
		{"s", "string", true},
		{"f64", "float64", true},
		{"f32", "float32", true},
		{"i32", "int32", true},
		{"u64", "uint64", true},
		{"u", "uint", true},
		{"i", "int15", false},
		{"b", "float64", false},
		{"s", "bool", false},
		{"f64", "float", false},
		{"f32", "String", false},
		{"i32", "int16", false},
		{"u64", "uint8", false},
		{"u", "int", false},
	}

	for _, test := range tests {
		isColType, _ := table.IsColType(test.colName, test.colType)
		// Ignore err. Returns err if col type is false.
		if isColType != test.expected {
			t.Errorf("Expecting table.IsColType(%s, %q) = %t but found %t", test.colName, test.colType, test.expected, isColType)
		}
	}
}

func ExampleTable_Sort() {
	tableString :=
	`[planets]
	name         mass distance
	string    float64  float64
	"Mercury"   0.055      0.4
	"Venus"     0.815      0.7
	"Earth"     1.000      1.0
	"Mars"      0.107      1.5
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(1) Unsorted table:")
	fmt.Println(table)

	// First let's sort the table by name.
	err = table.SetSortKeys("name")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(2) Sorted table by name:")
	fmt.Println(table)

	// Now let's sort the table by name but this time in reverse.
	err = table.SetSortKeys("name")
	if err != nil {
		log.Println(err)
	}
	err = table.SetSortKeysReverse("name")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(3) Sorted table by name in reverse:")
	fmt.Println(table)

	// Output:
	// (1) Unsorted table:
	// [planets]
	// name         mass distance
	// string    float64  float64
	// "Mercury"   0.055      0.4
	// "Venus"     0.815      0.7
	// "Earth"     1.0        1.0
	// "Mars"      0.107      1.5
	//
	// (2) Sorted table by name:
	// [planets]
	// name         mass distance
	// string    float64  float64
	// "Earth"     1.0        1.0
	// "Mars"      0.107      1.5
	// "Mercury"   0.055      0.4
	// "Venus"     0.815      0.7
	//
	// (3) Sorted table by name in reverse:
	// [planets]
	// name         mass distance
	// string    float64  float64
	// "Venus"     0.815      0.7
	// "Mercury"   0.055      0.4
	// "Mars"      0.107      1.5
	// "Earth"     1.0        1.0
}

func ExampleTable_SetSortKeys() {
	tableString :=
	`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(1) Unsorted table:")
	fmt.Println(table)

	// Sort the table by user.
	if err = table.SetSortKeys("user"); err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(2) Sorted by user:")
	fmt.Println(table)

	// Sort by user and lines.
	err = table.SetSortKeys("user", "lines")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(3) Sorted by user and lines:")
	fmt.Println(table)

	// Sort the table by user but reverse lines.
	err = table.SetSortKeys("user", "lines")
	if err != nil {
		log.Println(err)
	}
	err = table.SetSortKeysReverse("lines")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(4) Sort by user but reverse lines:")
	fmt.Println(table)

	// Sort the table by language and lines.
	err = table.SetSortKeys("language", "lines")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(5) Sort by language and lines:")
	fmt.Println(table)

	// Sort the table by language and lines and user.
	err = table.SetSortKeys("language", "lines", "user")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(6) Sort by language and lines and user:")
	fmt.Println(table)

	keysTable, err := table.GetSortKeysAsTable()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(7) SortKeys as a table:")
	fmt.Println(keysTable)

	// Output:
	// (1) Unsorted table:
	// [changes]
	// user     language    lines
	// string   string        int
	// "gri"    "Go"          100
	// "ken"    "C"           150
	// "glenda" "Go"          200
	// "rsc"    "Go"          200
	// "r"      "Go"          100
	// "ken"    "Go"          200
	// "dmr"    "C"           100
	// "r"      "C"           150
	// "gri"    "Smalltalk"    80
	// 
	// (2) Sorted by user:
	// [changes]
	// user     language    lines
	// string   string        int
	// "dmr"    "C"           100
	// "glenda" "Go"          200
	// "gri"    "Go"          100
	// "gri"    "Smalltalk"    80
	// "ken"    "C"           150
	// "ken"    "Go"          200
	// "r"      "Go"          100
	// "r"      "C"           150
	// "rsc"    "Go"          200
	// 
	// (3) Sorted by user and lines:
	// [changes]
	// user     language    lines
	// string   string        int
	// "dmr"    "C"           100
	// "glenda" "Go"          200
	// "gri"    "Smalltalk"    80
	// "gri"    "Go"          100
	// "ken"    "C"           150
	// "ken"    "Go"          200
	// "r"      "Go"          100
	// "r"      "C"           150
	// "rsc"    "Go"          200
	// 
	// (4) Sort by user but reverse lines:
	// [changes]
	// user     language    lines
	// string   string        int
	// "dmr"    "C"           100
	// "glenda" "Go"          200
	// "gri"    "Go"          100
	// "gri"    "Smalltalk"    80
	// "ken"    "Go"          200
	// "ken"    "C"           150
	// "r"      "C"           150
	// "r"      "Go"          100
	// "rsc"    "Go"          200
	// 
	// (5) Sort by language and lines:
	// [changes]
	// user     language    lines
	// string   string        int
	// "dmr"    "C"           100
	// "ken"    "C"           150
	// "r"      "C"           150
	// "r"      "Go"          100
	// "gri"    "Go"          100
	// "ken"    "Go"          200
	// "glenda" "Go"          200
	// "rsc"    "Go"          200
	// "gri"    "Smalltalk"    80
	// 
	// (6) Sort by language and lines and user:
	// [changes]
	// user     language    lines
	// string   string        int
	// "dmr"    "C"           100
	// "ken"    "C"           150
	// "r"      "C"           150
	// "gri"    "Go"          100
	// "r"      "Go"          100
	// "glenda" "Go"          200
	// "ken"    "Go"          200
	// "rsc"    "Go"          200
	// "gri"    "Smalltalk"    80
	// 
	// (7) SortKeys as a table:
	// [SortKeys]
	// index colName    colType  reverse
	//   int string     string   bool
	//     0 "language" "string" false
	//     1 "lines"    "int"    false
	//     2 "user"     "string" false
}

func ExampleTable_GobEncode_table() {
	s := `[sable_fur]
    i   s      f       t     b    bb            ui8
    int string float64 bool  byte []byte        []uint8
    1   "abc"  2.3     true  11   [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     false 22   [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     false 33   [33 34 35 36] [37 38 39]
    `
	tableToBeEncoded, err := NewTableFromString(s)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(1) Table ready to encode into binary.")
	fmt.Println(tableToBeEncoded)

	// Encode into binary.
    var binary []byte
    binary, err = tableToBeEncoded.GobEncode()
    if err != nil {
		log.Println(err)
    }

	// Now decode it back from binary to type *gotables.Table
	// Note: NewTableFromGob(binary) is equivalent to GobDecodeTable(binary)
    tableDecoded, err := NewTableFromGob(binary)
    if err != nil {
		log.Println(err)
	}
	fmt.Println("(2) Table decoded from binary.")
	fmt.Println(tableDecoded)

	// Output:
	// (1) Table ready to encode into binary.
	// [sable_fur]
	//   i s            f t        b bb            ui8
	// int string float64 bool  byte []byte        []uint8
	//   1 "abc"      2.3 true    11 [11 12 13 14] [15 16 17]
	//   2 "xyz"      4.5 false   22 [22 23 24 25] [26 27 28]
	//   3 "ssss"     4.9 false   33 [33 34 35 36] [37 38 39]
	// 
	// (2) Table decoded from binary.
	// [sable_fur]
	//   i s            f t        b bb            ui8
	// int string float64 bool  byte []byte        []uint8
	//   1 "abc"      2.3 true    11 [11 12 13 14] [15 16 17]
	//   2 "xyz"      4.5 false   22 [22 23 24 25] [26 27 28]
	//   3 "ssss"     4.9 false   33 [33 34 35 36] [37 38 39]
}

func ExampleTableSet_GobEncode_tableset() {
	s := `[sable_fur]
    i   s       f           b
    int string  float64     bool
    1   "abc"   2.3         true
    2   "xyz"   4.5         false
    3   "ssss"  4.9         false

	[Struct_With_Data]
	Fred int = 42
	Wilma int = 39
	Pebbles int = 2

	[Empty_Struct]
	Fred int

	[Empty_Table]
	Fred
	int
	`
	tableSetToEncode, err := NewTableSetFromString(s)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(1) TableSet ready to encode into binary.")
	fmt.Println(tableSetToEncode)

	// Encode into binary.
    var binary []bytes.Buffer
    binary, err = tableSetToEncode.GobEncode()
    if err != nil {
		log.Println(err)
    }

	// Now decode it back from binary to type *gotables.TableSet
    tableSetDecoded, err := GobDecodeTableSet(binary)
    if err != nil {
		log.Println(err)
	}
	fmt.Println("(2) TableSet decoded from binary.")
	fmt.Println(tableSetDecoded)

	// Output:
	// (1) TableSet ready to encode into binary.
	// [sable_fur]
	//   i s            f b
	// int string float64 bool
	//   1 "abc"      2.3 true
	//   2 "xyz"      4.5 false
	//   3 "ssss"     4.9 false
	// 
	// [Struct_With_Data]
	// Fred int = 42
	// Wilma int = 39
	// Pebbles int = 2
	//
	// [Empty_Struct]
	// Fred int
	//
	// [Empty_Table]
	// Fred
	//  int
	// 
	// (2) TableSet decoded from binary.
	// [sable_fur]
	//   i s            f b
	// int string float64 bool
	//   1 "abc"      2.3 true
	//   2 "xyz"      4.5 false
	//   3 "ssss"     4.9 false
	// 
	// [Struct_With_Data]
	// Fred int = 42
	// Wilma int = 39
	// Pebbles int = 2
	//
	// [Empty_Struct]
	// Fred int
	//
	// [Empty_Table]
	// Fred
	//  int
}

func ExampleTableSet_String() {
	// Deliberately unpadded (by hand) for contrast.
	s := `[sable_fur]
    i s f b
    int string float64 bool
    1 "abc" 2.3 true
    2 "xyz" 4.5 false
    3 "ssss" 4.9 false
	`
	tableSet, err := NewTableSetFromString(s)
	if err != nil {
		log.Println(err)
	}

	// Imagine this function (in both TableSet and Table) is called StringPaddedAndAligned()
	// It isn't, because it has to be called String(), but that's the functionality.

	// Notice that the default String() output for both TableSet and Table objects
	// is padded into easy to read columns, with numeric columns aligned right.
	// The design is: readability trumps compactness.
	// There are alternatives where size matters, such as compression, and StringUnpadded()

	fmt.Println("(1) TableSet (and Table) default String() output:")
	fmt.Println(tableSet.String())

	fmt.Println("(2) TableSet (and Table) same as default String() output:")
	fmt.Println(tableSet)

	// Output:
	// (1) TableSet (and Table) default String() output:
	// [sable_fur]
	//   i s            f b
	// int string float64 bool
	//   1 "abc"      2.3 true
	//   2 "xyz"      4.5 false
	//   3 "ssss"     4.9 false
	// 
	// (2) TableSet (and Table) same as default String() output:
	// [sable_fur]
	//   i s            f b
	// int string float64 bool
	//   1 "abc"      2.3 true
	//   2 "xyz"      4.5 false
	//   3 "ssss"     4.9 false
}

func ExampleTable_StringUnpadded() {
	// Deliberately padded (by hand) for contrast.
	s := `[sable_fur]
	  i s            f  b		ff
	int string float32  bool	float64
	  1 "abc"      2.34 true	7.899
	  2 "xyz"      4.5  false	6
	  3 "s  s"     4.9  false	5.5
	`
	table, err := NewTableFromString(s)
	if err != nil {
		log.Println(err)
	}

	// Note: the default String() output for both TableSet and Table objects
	// is padded into easy to read columns, with numeric columns aligned right.
	// The design is: readability trumps compactness.
	// There are alternatives where size matters, such as compression, and StringUnpadded()

	// This is an example of StringUnpadded() which uses minimal spacing between values.

	fmt.Println("TableSet (and Table) StringUnpadded() output:")
	fmt.Println(table.StringUnpadded())

	// Output:
	// TableSet (and Table) StringUnpadded() output:
	// [sable_fur]
	// i s f b ff
	// int string float32 bool float64
	// 1 "abc" 2.34 true 7.899
	// 2 "xyz" 4.5 false 6
	// 3 "s  s" 4.9 false 5.5
}

func ExampleTableSet_StringUnpadded() {
	// Deliberately padded (by hand) for contrast.
	s :=
	`[wombat_fur]
	  i s            f b
	int string float64 bool
	  1 "abc"      2.3 true
	  2 "xyz"      4.5 false
	  3 "s  s"     4.9 false

	[various]
	i	f		u		s
	int	float32	uint	string
	3	44.55	2		"Here I am!"
	4	22.99	255		"And now I'm not ..."
	`
	tableSet, err := NewTableSetFromString(s)
	if err != nil {
		log.Println(err)
	}

	// Note: the default String() output for both TableSet and Table objects
	// is padded into easy to read columns, with numeric columns aligned right.
	// The design is: readability trumps compactness.
	// There are alternatives where size matters, such as compression, and StringUnpadded()

	// This is an example of StringUnpadded() which uses minimal spacing between values.

	fmt.Println("TableSet (and Table) StringUnpadded() output:")
	fmt.Println(tableSet.StringUnpadded())

	// Output:
	// TableSet (and Table) StringUnpadded() output:
	// [wombat_fur]
	// i s f b
	// int string float64 bool
	// 1 "abc" 2.3 true
	// 2 "xyz" 4.5 false
	// 3 "s  s" 4.9 false
	// 
	// [various]
	// i f u s
	// int float32 uint string
	// 3 44.55 2 "Here I am!"
	// 4 22.99 255 "And now I'm not ..."
}

func TestNewTableFromCols(t *testing.T) {

	var table *Table
	var err error

	var tests = []struct {
		colNames []string
		colTypes []string
		expected bool
	}{
		{[]string{"Age", "Mothballs", "delims", "tags"}, []string{"int", "bool", "string", "string"}, true},
		{[]string{"Age", "Mothballs", "delims"}, []string{"int", "bool", "string", "string"}, false},	// Missing name
		{[]string{"Age", "Mothballs", "delims", "tags"}, []string{"int", "bool", "string",}, false},		// Missing type
		{[]string{}, []string{"int", "bool", "string", "string"}, false},	// Empty name slice
		{[]string{"Age", "Mothballs", "delims", "tags"}, []string{}, false},	// Empty type slice
		{[]string{}, []string{}, true},	// Empty table is allowed
	}

	for _, test := range tests {

		table, err = newTableFromCols("Moviegoers", test.colNames, test.colTypes)
		if (err == nil) != test.expected {
			if err != nil {
				t.Error(err)
			} else {
				t.Errorf("Expecting fail: NewTableFromCols(\"Moviegoers\", %v, %v)", test.colNames, test.colTypes)
			}
		}

		_, err = table.IsValidTable()
		if (err == nil) != test.expected {
			t.Error(err)
		}

		err = table.AppendRows(1)
		if (err == nil) != test.expected {
			t.Error(err)
		}

		if table != nil {
			rowCount := table.RowCount()
			if (rowCount == 1) != test.expected {
				t.Error(err)
			}
		}

		_, err = table.IsValidRow(0)
		if (err == nil) != test.expected {
			t.Error(err)
		}
	}
}

func ExampleTable_SetRowFloatCellsToNaN() {
	s := `[three_rows]
	  i s            f b	 f2
	int string float32 bool  float64
	  0 "abc"      2.3 true  42.0
	  1 "xyz"      4.5 false 43.0
	  2 "s  s"     4.9 false 44.0
	`
	table, err := NewTableFromString(s)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Before ...")
	fmt.Println(table)

	rowIndex := 1	// The middle row.
	err = table.SetRowFloatCellsToNaN(rowIndex)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("After ...")
	fmt.Println(table)

	// Output:
	// Before ...
	// [three_rows]
	//   i s            f b          f2
	// int string float32 bool  float64
	//   0 "abc"      2.3 true       42
	//   1 "xyz"      4.5 false      43
	//   2 "s  s"     4.9 false      44
	// 
	// After ...
	// [three_rows]
	//   i s            f b          f2
	// int string float32 bool  float64
	//   0 "abc"      2.3 true       42
	//   1 "xyz"      NaN false     NaN
	//   2 "s  s"     4.9 false      44
}

func ExampleTable_SetCellToZeroValue() {
	s := `[three_rows]
	i    s        f1 	b	 f2
	int  string float32 bool  float64
	  0  "abc"      2.3 true  42.0
	  1  "xyz"      4.5 false 43.0
	  2  "s  s"     4.9 false 44.0
	`
	table, err := NewTableFromString(s)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Initial table:")
	fmt.Println(table)

	err = table.SetCellToZeroValue("s", 1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("table.SetCellToZeroValue(\"s\", 1)")
	fmt.Println(table)

	err = table.SetCellToZeroValue("f1", 0)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("table.SetCellToZeroValue(\"f1\", 0)")
	fmt.Println(table)

	err = table.SetCellToZeroValue("b", 0)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("table.SetCellToZeroValue(\"b\", 0)")
	fmt.Println(table)

	err = table.SetCellToZeroValue("i", 2)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("table.SetCellToZeroValue(\"i\", 2)")
	fmt.Println(table)

	// Output:
	// Initial table:
	// [three_rows]
	//   i s           f1 b          f2
	// int string float32 bool  float64
	//   0 "abc"      2.3 true       42
	//   1 "xyz"      4.5 false      43
	//   2 "s  s"     4.9 false      44
	// 
	// table.SetCellToZeroValue("s", 1)
	// [three_rows]
	//   i s           f1 b          f2
	// int string float32 bool  float64
	//   0 "abc"      2.3 true       42
	//   1 ""         4.5 false      43
	//   2 "s  s"     4.9 false      44
	// 
	// table.SetCellToZeroValue("f1", 0)
	// [three_rows]
	//   i s           f1 b          f2
	// int string float32 bool  float64
	//   0 "abc"      0.0 true       42
	//   1 ""         4.5 false      43
	//   2 "s  s"     4.9 false      44
	// 
	// table.SetCellToZeroValue("b", 0)
	// [three_rows]
	//   i s           f1 b          f2
	// int string float32 bool  float64
	//   0 "abc"      0.0 false      42
	//   1 ""         4.5 false      43
	//   2 "s  s"     4.9 false      44
	// 
	// table.SetCellToZeroValue("i", 2)
	// [three_rows]
	//   i s           f1 b          f2
	// int string float32 bool  float64
	//   0 "abc"      0.0 false      42
	//   1 ""         4.5 false      43
	//   0 "s  s"     4.9 false      44
}

func TestTable_RenameCol(t *testing.T) {

	tableString :=
	`[Renaming]
	i int
	j int
	k int
	`

	var tests = []struct {
		from string
		to string
		expected bool
	}{
		{"i", "m", true},
		{"i", "i", false},
		{"i", "j", false},
		{"f", "m", false},
	}

	for _, test := range tests {

		// Reinstate table for each test. For cognitive simplicity.
		table, err := NewTableFromString(tableString)
		if err != nil {
			t.Error(err)
		}

		err = table.RenameCol(test.from, test.to)
		if (err == nil) != test.expected {
			t.Errorf("Expecting table.RenameCol(%q, %q) %s but found err = %v", test.from, test.to, ternString(test.expected, "SUCCESS", "FAILURE"), err)
		}
	}
}

func ternString(itIs bool, ifTrue string, ifFalse string) string {
	if (itIs) {
		return ifTrue
	} else {
		return ifFalse
	}
}

func TestPlural(t *testing.T) {

	var tests = []struct {
		in int
		expected string
	}{
		{ 1, ""},
		{-1, ""},
		{-2, "s"},
		{ 2, "s"},
	}

	for _, test := range tests {

		var result string = plural(test.in)
		if (result != test.expected) {
			t.Errorf("Expecting plural(%d) = %q but found: %q", test.in, test.expected, result)
		}
	}
}

func TestSort(t *testing.T) {

	table, err := NewTable("HasZeroSortKeys")
	if err != nil {
		t.Error(err)
	}

	err = table.Sort()
	if err == nil {
		t.Errorf("Expecting table.Sort() err because of 0 sort keys")
	}
}

func TestSearch(t *testing.T) {

	tableString :=
	`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)
	}

	_, err = table.Search()
	if err == nil {
		t.Errorf("Expecting table.Search() err because of 0 sort keys")
	}

	// Clear sort keys (if any) by calling with empty argument list.
	err = table.SetSortKeys()	// Note: sort keys count 0
	if err != nil {
		t.Error(err)
	}

	err = table.SetSortKeys("user")	// Note: sort keys count 1
	if err != nil {
		t.Error(err)
	}

	err = table.Sort()
	if err != nil {
		t.Error(err)
	}

	_, err = table.Search()	// Note: 0 search values passed to Search()
	if err == nil {
		t.Errorf("Expecting searchValues count 0 != sort keys count 1")
	}

	_, err = table.Search("glenda")
	if err != nil {
		t.Error(err)
	}

}

func TestIsValidColValue (t *testing.T) {
	tableString :=
	`[Types]
	i int
	b bool
	f64 float64
	f32 float32
	s string
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		col string
		val interface{}
		expecting bool
	}{
		{"i", 8, true},
		{"b", true, true},
		{"f64", 23.4, true},
		{"s", "sss", true},
		{"i", false, false},
		{"b", 67.8, false},
		{"f64", 23, false},
		{"s", 8, false},
		{"f32", 23.4, false},			// Floating point constant is float64
		{"f32", float32(23.4), true},	// It's now a float32
	}

	for _, test := range tests {

		result, err := table.IsValidColValue(test.col, test.val)
		if (result != test.expecting) {
			t.Errorf("Expecting IsValidColValue(%q, %v) = %t but found: %t, err: %v", test.col, test.val, test.expecting, result, err)
		}
	}
}

func ExampleTable_Search_keys1() {
	// mass:     Earth = 1 (relative to Earth)
	// distance: Earth = 1 (relative to Earth - AU)
	// http://www.windows2universe.org/our_solar_system/planets_table.html
	// http://www.space.com/17001-how-big-is-the-sun-size-of-the-sun.html
	tableString :=
	`[planets]
	name         mass distance moons index mnemonic
	string    float64   float64   int   int string
	"Sun"      333333        0     0    -1 ""
	"Mercury"   0.055      0.4     0     0 "my"
	"Venus"     0.815      0.7     0     1 "very"
	"Earth"     1.000      1.0     1     2 "elegant"
	"Mars"      0.107      1.5     2     3 "mother"
	"Jupiter" 318.000      5.2    67     4 "just"
	"Saturn"   95.000     29.4    62     5 "sat"
	"Uranus"   15.000     84.0    27     6 "upon"
	"Neptune"  17.000    164.0    13     7 "nine ... porcupines"
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(1) Unsorted table:")
	fmt.Println(table)

	// First let's sort the table by name.
	err = table.SetSortKeys("name")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(2) Sorted table by name:")
	fmt.Println(table)

	searchValue := "Mars" // 2
	fmt.Printf("(3) Search for name: %s\n", searchValue)
	rowIndex, err := table.Search(searchValue)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Found %s at rowIndex = %d\n", searchValue, rowIndex)
	fmt.Println()

	searchValue = "Pluto" // -1
	fmt.Printf("(4) Search for name: %s\n", searchValue)
	rowIndex, _ = table.Search(searchValue)
	fmt.Printf("Found %s at rowIndex = %d (missing)\n", searchValue, rowIndex)

	// Output:
	// (1) Unsorted table:
	// [planets]
	// name            mass distance moons index mnemonic
	// string       float64  float64   int   int string
	// "Sun"     333333.0        0.0     0    -1 ""
	// "Mercury"      0.055      0.4     0     0 "my"
	// "Venus"        0.815      0.7     0     1 "very"
	// "Earth"        1.0        1.0     1     2 "elegant"
	// "Mars"         0.107      1.5     2     3 "mother"
	// "Jupiter"    318.0        5.2    67     4 "just"
	// "Saturn"      95.0       29.4    62     5 "sat"
	// "Uranus"      15.0       84.0    27     6 "upon"
	// "Neptune"     17.0      164.0    13     7 "nine ... porcupines"
	// 
	// (2) Sorted table by name:
	// [planets]
	// name            mass distance moons index mnemonic
	// string       float64  float64   int   int string
	// "Earth"        1.0        1.0     1     2 "elegant"
	// "Jupiter"    318.0        5.2    67     4 "just"
	// "Mars"         0.107      1.5     2     3 "mother"
	// "Mercury"      0.055      0.4     0     0 "my"
	// "Neptune"     17.0      164.0    13     7 "nine ... porcupines"
	// "Saturn"      95.0       29.4    62     5 "sat"
	// "Sun"     333333.0        0.0     0    -1 ""
	// "Uranus"      15.0       84.0    27     6 "upon"
	// "Venus"        0.815      0.7     0     1 "very"
	// 
	// (3) Search for name: Mars
	// Found Mars at rowIndex = 2
	// 
	// (4) Search for name: Pluto
	// Found Pluto at rowIndex = -1 (missing)
}

func ExampleTable_Search_keys1Reverse() {
	// mass:     Earth = 1 (relative to Earth)
	// distance: Earth = 1 (relative to Earth - AU)
	// http://www.windows2universe.org/our_solar_system/planets_table.html
	tableString :=
	`[planets]
	name         mass distance moons index mnemonic
	string    float64  float64   int   int string
	"Mercury"   0.055      0.4     0     0 "my"
	"Venus"     0.815      0.7     0     1 "very"
	"Earth"     1.000      1.0     1     2 "elegant"
	"Mars"      0.107      1.5     2     3 "mother"
	"Jupiter" 318.000      5.2    67     4 "just"
	"Saturn"   95.000     29.4    62     5 "sat"
	"Uranus"   15.000     84.0    27     6 "upon"
	"Neptune"  17.000    164.0    13     7 "nine ... porcupines"
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(1) Unsorted table:")
	fmt.Println(table)

	// First let's sort the table by name in reverse.
	err = table.SetSortKeys("name")
	if err != nil {
		log.Println(err)
	}
	err = table.SetSortKeysReverse("name")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(2) Sorted table by name in reverse order:")
	fmt.Println(table)

	searchValue := "Mars" // 5
	fmt.Printf("(3) Search for name: %s\n", searchValue)
	rowIndex, err := table.Search(searchValue)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Found %s at rowIndex = %d\n", searchValue, rowIndex)

	searchValue = "Pluto" // -1
	fmt.Printf("(4) Search for name: %s\n", searchValue)
	rowIndex, _ = table.Search(searchValue)
	fmt.Printf("Found %s at rowIndex = %d (missing)\n", searchValue, rowIndex)

	// Output:
	// (1) Unsorted table:
	// [planets]
	// name         mass distance moons index mnemonic
	// string    float64  float64   int   int string
	// "Mercury"   0.055      0.4     0     0 "my"
	// "Venus"     0.815      0.7     0     1 "very"
	// "Earth"     1.0        1.0     1     2 "elegant"
	// "Mars"      0.107      1.5     2     3 "mother"
	// "Jupiter" 318.0        5.2    67     4 "just"
	// "Saturn"   95.0       29.4    62     5 "sat"
	// "Uranus"   15.0       84.0    27     6 "upon"
	// "Neptune"  17.0      164.0    13     7 "nine ... porcupines"
	// 
	// (2) Sorted table by name in reverse order:
	// [planets]
	// name         mass distance moons index mnemonic
	// string    float64  float64   int   int string
	// "Venus"     0.815      0.7     0     1 "very"
	// "Uranus"   15.0       84.0    27     6 "upon"
	// "Saturn"   95.0       29.4    62     5 "sat"
	// "Neptune"  17.0      164.0    13     7 "nine ... porcupines"
	// "Mercury"   0.055      0.4     0     0 "my"
	// "Mars"      0.107      1.5     2     3 "mother"
	// "Jupiter" 318.0        5.2    67     4 "just"
	// "Earth"     1.0        1.0     1     2 "elegant"
	// 
	// (3) Search for name: Mars
	// Found Mars at rowIndex = 5
	// (4) Search for name: Pluto
	// Found Pluto at rowIndex = -1 (missing)
}

func TestTable_Search_1key(t *testing.T) {
	// mass:     Earth = 1 (relative to Earth)
	// distance: Earth = 1 (relative to Earth - AU)
	// http://www.windows2universe.org/our_solar_system/planets_table.html
	tableString :=
	`[planets]
	name         mass distance moons index mnemonic
	string    float64  float64   int   int string
	"Mercury"   0.055      0.4     0     0 "my"
	"Venus"     0.815      0.7     0     1 "very"
	"Earth"     1.000      1.0     1     2 "elegant"
	"Mars"      0.107      1.5     2     3 "mother"
	"Jupiter" 318.000      5.2    67     4 "just"
	"Saturn"   95.000     29.4    62     5 "sat"
	"Uranus"   15.000     84.0    27     6 "upon"
	"Neptune"  17.000    164.0    13     7 "nine ... porcupines"
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)
	}

	// First let's sort the table by name.
	err = table.SetSortKeys("name")
	if err != nil {
		t.Error(err)
	}
	err = table.Sort()
	if err != nil {
		t.Error(err)
	}

	var searchValue string
	var expecting int
	var rowIndex int

	// Search for entries that exist in the table.
	for i := 0; i < table.RowCount(); i++ {
		searchValue, err = table.GetString("name", i)
		if err != nil {
			t.Error(err)
		}
		expecting := i
		rowIndex, err := table.Search(searchValue)
		if err != nil {
			t.Error(err)
		}
		if rowIndex != expecting {
			t.Errorf("Expecting Search(%q) = %d but found: %d", searchValue, expecting, rowIndex)
		}
	}

	// Search for entries that don't exist.
	dontExist := []string{
		"Sun",
		"Moon",
		"Pluto",
	}
	for _, item := range dontExist {
		searchValue = item
		expecting = -1
		rowIndex, err = table.Search(searchValue)
		if err == nil {
			t.Errorf("Expecting an error with Search(%v)", searchValue)
		}
		if rowIndex != expecting {
			t.Errorf("Expecting Search(%q) = %d but found: %d", searchValue, expecting, rowIndex)
		}
	}
}

func TestTable_Search_1key_reverse(t *testing.T) {
	// mass:     Earth = 1 (relative to Earth)
	// distance: Earth = 1 (relative to Earth - AU)
	// http://www.windows2universe.org/our_solar_system/planets_table.html
	tableString :=
	`[planets]
	name         mass distance moons index mnemonic
	string    float64  float64   int   int string
	"Mercury"   0.055      0.4     0     0 "my"
	"Venus"     0.815      0.7     0     1 "very"
	"Earth"     1.000      1.0     1     2 "elegant"
	"Mars"      0.107      1.5     2     3 "mother"
	"Jupiter" 318.000      5.2    67     4 "just"
	"Saturn"   95.000     29.4    62     5 "sat"
	"Uranus"   15.000     84.0    27     6 "upon"
	"Neptune"  17.000    164.0    13     7 "nine ... porcupines"
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)
	}

	// First let's sort the table by name - in reverse order.
	err = table.SetSortKeys("name")
	if err != nil {
		t.Error(err)
	}
	err = table.SetSortKeysReverse("name")
	if err != nil {
		t.Error(err)
	}
	err = table.Sort()
	if err != nil {
		t.Error(err)
	}

	var searchValue string
	var expecting int
	var rowIndex int

	// Search for entries that exist in the table.
	for i := 0; i < table.RowCount(); i++ {
		searchValue, err = table.GetString("name", i)
		if err != nil {
			t.Error(err)
		}
		expecting := i

		rowIndex, err := table.Search(searchValue)
		if err != nil {
			t.Error(err)
		}

		if rowIndex != expecting {
			t.Errorf("Expecting Search(%q) = %d but found: %d", searchValue, expecting, rowIndex)
		}
	}
//	log.Printf("%q expecting %d found %d", searchValue, expecting, rowIndex)

	// Search for entries that don't exist.
	dontExist := []string{
		"Sun",
		"Moon",
		"Pluto",
	}
	for _, item := range dontExist {
		searchValue = item
		expecting = -1

		rowIndex, err = table.Search(searchValue)
		if err == nil {
			t.Errorf("Expecting an error with Search(%v)", searchValue)
		}
		if rowIndex != expecting {
			t.Errorf("Expecting Search(%q) = %d but found: %d", searchValue, expecting, rowIndex)
		}
	}
}

func TestTable_Search_2keys(t *testing.T) {
	tableString :=
	`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)
	}

// where("SORTING")
	// First let's sort the table by user and lines.
	err = table.SetSortKeys("user", "lines")
	if err != nil {
		t.Error(err)
	}
	err = table.Sort()
	if err != nil {
		t.Error(err)
	}
// fmt.Printf("here:\n%s", table)

// where("SEARCHING")
	var searchValues []interface{} = make([]interface{}, 2)
	var expecting int
	var found int

	// Search for entries that exist in the table.
	for i := 0; i < table.RowCount(); i++ {
		searchValues[0], err = table.GetString("user", i)
		if err != nil {
			t.Error(err)
		}
		searchValues[1], err = table.GetInt("lines", i)
		if err != nil {
			t.Error(err)
		}
// where(fmt.Sprintf("test len(searchValues) = %d", len(searchValues)))
// where(fmt.Sprintf("test searchValues = %v", searchValues))
		expecting := i
// where()
		found, err := table.Search(searchValues...)
		if err != nil {
			t.Error(err)
		}
		if found != expecting {
			t.Errorf("Expecting Search(%v) = %d but found: %d", searchValues, expecting, found)
		}
	}
//	log.Printf("%q expecting %d found %d", searchValues, expecting, found)

	// Search for entries that don't exist.
	dontExist := [][]interface{}{
		{"steve",   42},
		{"bill",  42},
		{"larry", 42},
	}
	for _, item := range dontExist {
		searchValues = item
		expecting = -1
// where()
		found, err = table.Search(searchValues...)
		if found != expecting {
			t.Errorf("Expecting Search(%q) = %d but found: %d", searchValues, expecting, found)
		}
	}
}

func TestTable_Search_2keys_reverse2nd(t *testing.T) {
	tableString :=
	`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)
	}

	// First let's sort the table by user and lines.
	err = table.SetSortKeys("user", "lines")
	if err != nil {
		t.Error(err)
	}
	err = table.SetSortKeysReverse("lines")
	if err != nil {
		t.Error(err)
	}
	err = table.Sort()
	if err != nil {
		t.Error(err)
	}
// fmt.Printf("here:\n%s", table)

// where("SEARCHING")
	var searchValues []interface{} = make([]interface{}, 2)
	var expecting int
	var found int

	// Search for entries that exist in the table.
	for i := 0; i < table.RowCount(); i++ {
		searchValues[0], err = table.GetString("user", i)
		if err != nil {
			t.Error(err)
		}
		searchValues[1], err = table.GetInt("lines", i)
		if err != nil {
			t.Error(err)
		}
// where(fmt.Sprintf("test len(searchValues) = %d", len(searchValues)))
// where(fmt.Sprintf("test searchValues = %v", searchValues))
		expecting := i
// where()
		found, err := table.Search(searchValues...)
		if err != nil {
			t.Error(err)
		}
		if found != expecting {
			t.Errorf("Expecting Search(%v) = %d but found: %d", searchValues, expecting, found)
		}
	}
//	log.Printf("%q expecting %d found %d", searchValues, expecting, found)

	// Search for entries that don't exist.
	dontExist := [][]interface{}{
		{"steve",   42},
		{"bill",  42},
		{"larry", 42},
	}
	for _, item := range dontExist {
		searchValues = item
		expecting = -1
// where()
		found, err = table.Search(searchValues...)
		if found != expecting {
			t.Errorf("Expecting Search(%q) = %d but found: %d", searchValues, expecting, found)
		}
	}
}

func TestTable_Search_2keys_reverseBoth(t *testing.T) {
	tableString :=
	`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)
	}

	err = table.SetSortKeys("user", "lines")
	if err != nil {
		t.Error(err)
	}
	err = table.SetSortKeysReverse("user", "lines")
	if err != nil {
		t.Error(err)
	}
	err = table.Sort()
	if err != nil {
		t.Error(err)
	}

	var searchValues []interface{} = make([]interface{}, 2)
	var expecting int
	var found int

	// Search for entries that exist in the table.
	for i := 0; i < table.RowCount(); i++ {
		searchValues[0], err = table.GetString("user", i)
		if err != nil {
			t.Error(err)
		}
		searchValues[1], err = table.GetInt("lines", i)
		if err != nil {
			t.Error(err)
		}
		expecting = i
		found, err := table.Search(searchValues...)
		if err != nil {
			t.Error(err)
		}
		if found != expecting {
			t.Errorf("Expecting Search(%v) = %d but found: %d", searchValues, expecting, found)
		}
	}

	// Search for entries that don't exist.
	dontExist := [][]interface{}{
		{"steve",   42},
		{"bill",  42},
		{"larry", 42},
	}
	for _, item := range dontExist {
		searchValues = item
		expecting = -1
		found, err = table.Search(searchValues...)
		if err == nil {
			t.Errorf("Expecting an error with Search(%v)", searchValues)
		}
		if found != expecting {
			t.Errorf("Expecting Search(%q) = %d but found: %d", searchValues, expecting, found)
		}
	}
}

func ExampleNewTableFromString_planets() {
	// mass:     Earth = 1 (relative to Earth)
	// distance: Earth = 1 (relative to Earth - AU)
	// http://www.windows2universe.org/our_solar_system/planets_table.html
	tableString :=
	`[planets]
	name         mass distance moons index mnemonic
	string    float64  float64   int   int string
	"Mercury"   0.055      0.4     0     0 "my"
	"Venus"     0.815      0.7     0     1 "very"
	"Earth"     1.000      1.0     1     2 "elegant"
	"Mars"      0.107      1.5     2     3 "mother"
	"Jupiter" 318.000      5.2    67     4 "just"
	"Saturn"   95.000     29.4    62     5 "sat"
	"Uranus"   15.000     84.0    27     6 "upon"
	"Neptune"  17.000    164.0    13     7 "nine ... porcupines"
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	// Simply echo it back out.
	fmt.Println(table)

/*
	Notice that by default the columns of data are padded with spaces and numeric types
	are right-aligned.
	This reflects the opinion that human readability is important.
	*Table.String() and *TableSet.String() call their underlying StringPadded() methods.
	Where human readability is not important (with messaging or as a wire format) use:
	*Table.StringUnpadded()
	*TableSet.StringUnpadded()
	StringUnpadded() is 3 to 4 times faster.
	Reading a padded table string is only slightly slower (about 2.7% slower).
*/

	// For unpadded output:
    fmt.Println(table.StringUnpadded())

	// Output:
	// [planets]
	// name         mass distance moons index mnemonic
	// string    float64  float64   int   int string
	// "Mercury"   0.055      0.4     0     0 "my"
	// "Venus"     0.815      0.7     0     1 "very"
	// "Earth"     1.0        1.0     1     2 "elegant"
	// "Mars"      0.107      1.5     2     3 "mother"
	// "Jupiter" 318.0        5.2    67     4 "just"
	// "Saturn"   95.0       29.4    62     5 "sat"
	// "Uranus"   15.0       84.0    27     6 "upon"
	// "Neptune"  17.0      164.0    13     7 "nine ... porcupines"
	//
	// [planets]
	// name mass distance moons index mnemonic
	// string float64 float64 int int string
	// "Mercury" 0.055 0.4 0 0 "my"
	// "Venus" 0.815 0.7 0 1 "very"
	// "Earth" 1 1 1 2 "elegant"
	// "Mars" 0.107 1.5 2 3 "mother"
	// "Jupiter" 318 5.2 67 4 "just"
	// "Saturn" 95 29.4 62 5 "sat"
	// "Uranus" 15 84 27 6 "upon"
	// "Neptune" 17 164 13 7 "nine ... porcupines"
}

// This is not a comprehensive test.
func TestTable_Equals(t *testing.T) {
    var err error
    var table1, table2 *Table

	_, err = table1.Equals(table2)
    if err == nil {
		t.Errorf("Expecting an error calling Equals() on nil table")
    }
//	fmt.Println(err)

    t1string :=
    `[T1]
    i   s       f       ui
    int string  float64 uint
    1   "abc"   5.50    50
    2   "def"   6.66    60
    `
    table1, err = NewTableFromString(t1string)
    if err != nil {
        t.Error(err)
    }

	_, err = table1.Equals(table2)
    if err == nil {
		t.Errorf("Expecting an error calling Equals() with nil table")
	}
//	fmt.Println(err)

    t2string :=
    `[T2]
    ui      i   s       f
    uint    int string  float64
    50      1   "abc"   5.5
    60      2   "def"   6.6600
    `
    table2, err = NewTableFromString(t2string)
    if err != nil {
        t.Error(err)
    }

	equals, err := table1.Equals(table2)
    if !equals {
		t.Errorf("Expecting table1.Equals(table2) = true but found %t", equals)
	}
    if err != nil {
        t.Error(err)
    }
}

func ExampleTable_GetSortKeysAsTable() {
	tableString :=
	`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	// Sort the table by user but reverse lines.
	err = table.SetSortKeys("user", "lines")
	if err != nil {
		log.Println(err)
	}

	err = table.SetSortKeysReverse("lines")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("(1) GetSortKeysAsTable():")
	sortKeysTable, err := table.GetSortKeysAsTable()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(sortKeysTable)

	err = table.Sort()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("(2) Sort by user but reverse lines:")
	fmt.Println(table)

	// Output:
	// (1) GetSortKeysAsTable():
	// [SortKeys]
	// index colName colType  reverse
	//   int string  string   bool
	//     0 "user"  "string" false
	//     1 "lines" "int"    true
	// 
	// (2) Sort by user but reverse lines:
	// [changes]
	// user     language    lines
	// string   string        int
	// "dmr"    "C"           100
	// "glenda" "Go"          200
	// "gri"    "Go"          100
	// "gri"    "Smalltalk"    80
	// "ken"    "Go"          200
	// "ken"    "C"           150
	// "r"      "C"           150
	// "r"      "Go"          100
	// "rsc"    "Go"          200
}

func TestTable_SortKeyCount(t *testing.T) {
	tableString :=
	`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)
	}

	// First let's sort the table by user and lines.
	err = table.SetSortKeys("user", "lines")
	if err != nil {
		t.Error(err)
	}

	expecting := 2
	count := table.SortKeyCount()
	if count != expecting {
		t.Errorf("Expecting table.SortKeyCount() = %d but found %d", expecting, count)
	}
}

func TestTable_SetSortKeysFromTable(t *testing.T) {
	fromTableString :=
	`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	fromTable, err := NewTableFromString(fromTableString)
	if err != nil {
		t.Error(err)
	}

	// First let's sort the table by user and lines.
	err = fromTable.SetSortKeys("user", "lines")
	if err != nil {
		t.Error(err)
	}

	err = fromTable.SetSortKeysReverse("lines")
	if err != nil {
		t.Error(err)
	}

	toTableString :=
	`[ToTable]
	user	string
	lines	int
	`
	toTable, err := NewTableFromString(toTableString)
	if err != nil {
		t.Error(err)
	}

	err = toTable.SetSortKeysFromTable(fromTable)
	if err != nil {
		t.Error(err)
	}

	keysTable1, err := fromTable.GetSortKeysAsTable()
	if err != nil {
		t.Error(err)
	}

	keysTable2, err := toTable.GetSortKeysAsTable()
	if err != nil {
		t.Error(err)
	}

	expecting := true

	equals, err := keysTable1.Equals(keysTable2)
	if err != nil {
		t.Error(err)
	}

	if equals != expecting {
		t.Errorf("Expecting table1.Equals(table2) = %t but found %t", expecting, equals)
	}
}

func ExampleTable_OrderColsBySortKeys() {
	tableString :=
	`[MyTable]
	ColA   ColB Key2      ColC Key1 ColD ColE
	string  int string float64  int  int bool
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	err = table.SetSortKeys("Key1", "Key2")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("(1) Here is the table in its original column order:")
	fmt.Println(table)

	fmt.Println("(2) Here are the keys:")
	sortKeysTable, err := table.GetSortKeysAsTable()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(sortKeysTable)

	fmt.Println("(3) Order the sort key columns to the left:")
	err = table.OrderColsBySortKeys()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(table)

	// Output:
	// (1) Here is the table in its original column order:
	// [MyTable]
	// ColA   ColB Key2      ColC Key1 ColD ColE
	// string  int string float64  int  int bool
	// 
	// (2) Here are the keys:
	// [SortKeys]
	// index colName colType  reverse
	//   int string  string   bool
	//     0 "Key1"  "int"    false
	//     1 "Key2"  "string" false
	// 
	// (3) Order the sort key columns to the left:
	// [MyTable]
	// Key1 Key2   ColA   ColB    ColC ColD ColE
	//  int string string  int float64  int bool
}

/*
	This tests a copy gotables.Search() of sort.Search()
	to confirm that SearchLast() is a mirror image in
	behaviour: Search() is GE and SearchLast is LE.
*/
func Test_Search(t *testing.T) {

/*
	sliceToString := func(slice []int) string {
		var s string
		for i := 0; i < len(slice); i++ {
			s += fmt.Sprintf("%3d", slice[i])
		}
		return s
	}
*/

	const tests = 40		// Make this 20 for realism.
	const elements = 10
	const intRange = 10
	slice := make([]int, elements)
	indices := make([]int, elements)

	for i := 0; i < elements; i++ {
		indices[i] = i
	}

//	rand.Seed(time.Now().UnixNano())

	for i := 0; i < tests; i++ {
		for j := 0; j < elements; j++ {
			slice[j] = rand.Intn(intRange)
		}
		sort.Ints(slice)
		// fmt.Println()
		// fmt.Printf("%s()\n", funcName())
		// fmt.Printf("test[%2d] %s\n", i, sliceToString(slice))
		// fmt.Printf("test[%2d] %s\n", i, sliceToString(indices))
		var index int
		for searchFor := -1; searchFor <= intRange; searchFor++ {
			index = sort.Search(elements, func(element int) bool {
				return slice[element] >= searchFor
			})

			// fmt.Printf("index for %d is %2d\n", searchFor, index)

			if index >= elements {
				// fmt.Printf("%d is missing but would be at (nonexistent) index %d (insert before %d)\n", searchFor, index, index)
			} else {
				if slice[index] != searchFor {
					// Have we found at the very least A right element, or if it is missing, an element less than it.
					if slice[index] < searchFor {
						t.Error(fmt.Sprintf("test[%d] Expecting Search() slice[%d] = %d or more than %d, but found %d",
							i, index, searchFor, searchFor, slice[index]))
					} else {
						// fmt.Printf("%d is missing but would be at index %d (insert before %d)\n", searchFor, index, index)
					}
				}
			}

			if index > 0 && slice[index-1] == searchFor {
				// Have we found THE right element.
				t.Error(fmt.Sprintf("test[%d] Expecting Search() slice[%d] = %d to be lowest index, but found slice[%d-1] = %d lower",
					i, index, searchFor, index, slice[index-1]))
			}
		}
	}
}

// LE: Less than or equal.
func TestSearchLast(t *testing.T) {

/*
	// Inner function to convert a slice to a string.
	sliceToString := func(slice []int) string {
		var s string
		for i := 0; i < len(slice); i++ {
			s += fmt.Sprintf("%3d", slice[i])
		}
		return s
	}
*/

	const tests = 40		// Make this 20 for realism.
	const elements = 10
	const intRange = 10
	slice := make([]int, elements)
	indices := make([]int, elements)

	for i := 0; i < elements; i++ {
		indices[i] = i
	}

//	rand.Seed(time.Now().UnixNano())

	for i := 0; i < tests; i++ {
		for j := 0; j < elements; j++ {
			slice[j] = rand.Intn(intRange)
		}
		sort.Ints(slice)
		// fmt.Println()
		// fmt.Printf("%s()\n", funcName())
		// fmt.Printf("test[%2d] %s\n", i, sliceToString(slice))
		// fmt.Printf("test[%2d] %s\n", i, sliceToString(indices))
		var index int
		for searchFor := -1; searchFor <= intRange; searchFor++ {
			index = SearchLast(elements, func(element int) bool {
				return slice[element] <= searchFor
			})

			// fmt.Printf("index for %d is %2d\n", searchFor, index)

			if index < 0 {
			// fmt.Printf("%d is missing but would be at (nonexistent) index %d (insert after %d)\n", searchFor, index, index)
			} else {
				if slice[index] != searchFor {
					// Have we found at the very least A right element, or if it is missing, an element less than it.
					if slice[index] > searchFor {
						t.Error(fmt.Sprintf("test[%d] Expecting SearchLast() slice[%d] = %d or less than %d, but found %d",
							i, index, searchFor, searchFor, slice[index]))
					} else {
						// fmt.Printf("%d is missing but would be at index %d (insert after %d)\n", searchFor, index, index)
					}
				}
			}

			if index < elements-1 && slice[index+1] == searchFor {
				// Have we found THE right element.
				t.Error(fmt.Sprintf("test[%d] Expecting SearchLast() slice[%d] = %d to be greatest index, but found slice[%d+1] = %d greater",
					i, index, searchFor, index, slice[index+1]))
			}
		}
	}
}

func ExampleSearchLast() {

	var data []int = []int { 4, 8, 10, 10, 10, 20, 23, 29 }
	fmt.Printf("data: %v\n", data)
	fmt.Println("index: 0 1  2  3  4  5  6  7")
	fmt.Println()

	fmt.Printf("(1) Find an element that is present:\n")
	x := 23
	fmt.Printf("Searching for x: %d\n", x)
	i := SearchLast(len(data), func(i int) bool { return data[i] <= x } )
	fmt.Printf("x %d is, or would be, at index i: %d\n", x, i)

	// Check whether x is actually where SearchLast() said it is, or would be inserted.
	if i >= 0 && data[i] == x {
		fmt.Printf("x %d is present at data[%d]\n", x, i)
	} else {
		fmt.Printf("x is not present in data, but i %d is the index where it would be inserted AFTER.\n", i)
		fmt.Printf("Note that i can be -1 which does not exist in data.\n")
	}
	fmt.Println()

	fmt.Printf("(2) This time find an x that is present multiple times:\n")
	x = 10
	fmt.Printf("Searching for x: %d\n", x)
	i = SearchLast(len(data), func(i int) bool { return data[i] <= x } )
	fmt.Printf("x %d is, or would be, at index i: %d\n", x, i)

	// Check whether x is actually where SearchLast() said it is, or would be inserted.
	if i >= 0 && data[i] == x {
		fmt.Printf("x %d is present at data[%d]\n", x, i)
	} else {
		fmt.Printf("x is not present in data, but i %d is the index where it would be inserted AFTER.\n", i)
		fmt.Printf("Note that i can be -1 which does not exist in data.\n")
	}
	fmt.Println()

	fmt.Printf("(3) This time find an x that is missing between items in data:\n")
	x = 15
	fmt.Printf("Searching for x: %d\n", x)
	i = SearchLast(len(data), func(i int) bool { return data[i] <= x } )
	fmt.Printf("x %d is, or would be, at index i: %d\n", x, i)

	// Check whether x is actually where SearchLast() said it is, or would be inserted.
	if i >= 0 && data[i] == x {
		fmt.Printf("x %d is present at data[%d]\n", x, i)
	} else {
		fmt.Printf("x is not present in data, but i %d is the index where it would be inserted AFTER.\n", i)
		fmt.Printf("Note that i can be -1 which does not exist in data.\n")
	}
	fmt.Println()

	fmt.Printf("(4) This time find an x that is missing below all items in data:\n")
	x = 3
	fmt.Printf("Searching for x: %d\n", x)
	i = SearchLast(len(data), func(i int) bool { return data[i] <= x } )
	fmt.Printf("x %d is, or would be, at index i: %d\n", x, i)

	// Check whether x is actually where SearchLast() said it is, or would be inserted.
	if i >= 0 && data[i] == x {
		fmt.Printf("x %d is present at data[%d]\n", x, i)
	} else {
		fmt.Printf("x is not present in data, but i %d is the index where it would be inserted AFTER.\n", i)
		fmt.Printf("Note that i can be -1 which does not exist in data.\n")
	}
	fmt.Println()

	fmt.Printf("(5) This time find an x that is missing above all items in data:\n")
	x = 31
	fmt.Printf("Searching for x: %d\n", x)
	i = SearchLast(len(data), func(i int) bool { return data[i] <= x } )
	fmt.Printf("x %d is, or would be, at index i: %d\n", x, i)

	// Check whether x is actually where SearchLast() said it is, or would be inserted.
	if i >= 0 && data[i] == x {
		fmt.Printf("x %d is present at data[%d]\n", x, i)
	} else {
		fmt.Printf("x is not present in data, but i %d is the index where it would be inserted AFTER.\n", i)
		fmt.Printf("Note that i can be -1 which does not exist in data.\n")
	}
	fmt.Println()

	// Output:
	// data: [4 8 10 10 10 20 23 29]
	// index: 0 1  2  3  4  5  6  7
	// 
	// (1) Find an element that is present:
	// Searching for x: 23
	// x 23 is, or would be, at index i: 6
	// x 23 is present at data[6]
	// 
	// (2) This time find an x that is present multiple times:
	// Searching for x: 10
	// x 10 is, or would be, at index i: 4
	// x 10 is present at data[4]
	// 
	// (3) This time find an x that is missing between items in data:
	// Searching for x: 15
	// x 15 is, or would be, at index i: 4
	// x is not present in data, but i 4 is the index where it would be inserted AFTER.
	// Note that i can be -1 which does not exist in data.
	// 
	// (4) This time find an x that is missing below all items in data:
	// Searching for x: 3
	// x 3 is, or would be, at index i: -1
	// x is not present in data, but i -1 is the index where it would be inserted AFTER.
	// Note that i can be -1 which does not exist in data.
	// 
	// (5) This time find an x that is missing above all items in data:
	// Searching for x: 31
	// x 31 is, or would be, at index i: 7
	// x is not present in data, but i 7 is the index where it would be inserted AFTER.
	// Note that i can be -1 which does not exist in data.
}

func TestTable_SearchFirst_by_user(t *testing.T) {
	tableString :=
	`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)
	}

	err = table.SetSortKeys("user")
	if err != nil {
		t.Error(err)
	}
	err = table.Sort()
	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		searchValue string
		expecting int
	}{
		{"dmr",    0},
		{"glenda", 1},
		{"gri",    2},
		{"ken",    4},
		{"r",      6},
		{"rsc",    8},
		{"NOT",   -1},
	}

	for _, test := range tests {
		found, err := table.SearchFirst(test.searchValue)
		if (found != test.expecting) {
			t.Errorf("Expecting SearchFirst(%q) = %d but found: %d, err: %v", test.searchValue, test.expecting, found, err)
		}
	}
}

func TestTable_SearchLast_by_user(t *testing.T) {
	tableString :=
	`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)
	}

	err = table.SetSortKeys("user")
	if err != nil {
		t.Error(err)
	}
	err = table.Sort()
	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		searchValue string
		expecting int
	}{
		{"dmr",    0},
		{"glenda", 1},
		{"gri",    3},
		{"ken",    5},
		{"r",      7},
		{"rsc",    8},
		{"NOT",   -1},
	}

	for _, test := range tests {
		found, err := table.SearchLast(test.searchValue)
		if (found != test.expecting) {
			t.Errorf("Expecting SearchLast(%q) = %d but found: %d, err: %v", test.searchValue, test.expecting, found, err)
		}
	}
}

func TestTable_SearchRange_by_user(t *testing.T) {
	tableString :=
	`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)
	}

	err = table.SetSortKeys("user")
	if err != nil {
		t.Error(err)
	}
	err = table.Sort()
	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		searchValue string
		expectingFirst int
		expectingLast int
	}{
		{"dmr",    0, 0},
		{"glenda", 1, 1},
		{"gri",    2, 3},
		{"ken",    4, 5},
		{"r",      6, 7},
		{"rsc",    8, 8},
		{"NOT",   -1,-1},
	}

	for _, test := range tests {
		foundFirst, foundLast, err := table.SearchRange(test.searchValue)
		if (foundFirst != test.expectingFirst || foundLast != test.expectingLast) {
			t.Errorf("Expecting SearchRange(%q) = %d, %d but found: %d, %d err: %v",
				test.searchValue, test.expectingFirst, test.expectingLast, foundFirst, foundLast, err)
		}
	}
}

func TestTable_SearchRange_by_user_lines(t *testing.T) {
	tableString :=
	`[changes]
	user     language    lines index
	string   string        int   int
	"rsc"    "Go"          200     0
	"r"      "Go"          100     0
	"r"      "C"           150     0
	"ken"    "C"           150     0
	"ken"    "Go"          200     0
	"ken"    "Go"          200     0
	"gri"    "Smalltalk"    80     0
	"gri"    "Go"          100     0
	"gri"    "Go"          100     0
	"gri"    "Go"          100     0
	"glenda" "Go"          200     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)
	}

	err = table.SetSortKeys("user", "lines")
	if err != nil {
		t.Error(err)
	}
	err = table.Sort()
	if err != nil {
		t.Error(err)
	}

	// To eye-ball errors.
	for i := 0; i < table.RowCount(); i++ {
		err = table.SetInt("index", i, i)
		if err != nil {
			t.Fatal(err)
		}
	}

	var tests = []struct {
		searchName string
		searchLines int
		expectingFirst int
		expectingLast int
	}{
		{"dmr",    100,  0,  4},
		{"glenda", 200,  5,  5},
		{"gri",    100,  7,  9},
		{"ken",    200, 11, 12},
		{"r",      150, 14, 14},
		{"rsc",    200, 15, 15},
		{"NOT",    500, -1, -1},
		{"NOT",    200, -1, -1},
		{"rsc",    100, -1, -1},
	}

	for _, test := range tests {
		foundFirst, foundLast, err := table.SearchRange(test.searchName, test.searchLines)
		if (foundFirst != test.expectingFirst || foundLast != test.expectingLast) {
			t.Errorf("Expecting SearchRange(%q, %d) = %d, %d but found: %d, %d err: %v",
				test.searchName, test.searchLines, test.expectingFirst, test.expectingLast, foundFirst, foundLast, err)
			fmt.Println(table)
		}
	}
}

func TestTable_SearchRange_by_user_lines_reverse_lines(t *testing.T) {
	tableString :=
	`[changes]
	user     language    lines index
	string   string        int   int
	"rsc"    "Go"          200     0
	"r"      "Go"          100     0
	"r"      "C"           150     0
	"ken"    "C"           150     0
	"ken"    "Go"          200     0
	"ken"    "Go"          200     0
	"gri"    "Smalltalk"    80     0
	"gri"    "Go"          100     0
	"gri"    "Go"          100     0
	"gri"    "Go"          100     0
	"glenda" "Go"          200     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Error(err)
	}

	err = table.SetSortKeys("user", "lines")
	if err != nil {
		t.Error(err)
	}
	err = table.SetSortKeysReverse("lines")
	if err != nil {
		t.Error(err)
	}
	err = table.Sort()
	if err != nil {
		t.Error(err)
	}

	// To eye-ball errors.
	for i := 0; i < table.RowCount(); i++ {
		err = table.SetInt("index", i, i)
		if err != nil {
			t.Fatal(err)
		}
	}

	var tests = []struct {
		searchName string
		searchLines int
		expectingFirst int
		expectingLast int
	}{
		{"dmr",    100,  0,  4},
		{"glenda", 200,  5,  5},
		{"gri",    100,  6,  8},
		{"ken",    200, 10, 11},
		{"r",      150, 13, 13},
		{"rsc",    200, 15, 15},
		{"NOT",    500, -1, -1},
	}

	for _, test := range tests {
		foundFirst, foundLast, err := table.SearchRange(test.searchName, test.searchLines)
		if (foundFirst != test.expectingFirst || foundLast != test.expectingLast) {
			t.Errorf("Expecting SearchRange(%q, %d) = %d, %d but found: %d, %d err: %v",
				test.searchName, test.searchLines, test.expectingFirst, test.expectingLast, foundFirst, foundLast, err)
			fmt.Println(table)
		}
	}
}

func ExampleTable_Merge() {

	t1string :=
    `[Table1]
    XYZ     y   s       f       i   diff
    string  int string  float64 int int
    "X"     1   "abc"   1.11    1   7
    "Y"     3   "ghi"   7.8910  3   8
    "Z"     2   "def"   NaN     2   9
    "A"     4   "jkl"   0       4   6
    "B"     5   "mno"   0       5   4
    "C"     8   "pqr"   NaN     6   45
    `
    table1, err := NewTableFromString(t1string)
    if err != nil {
		log.Println(err)
    }

    fmt.Println(table1)

    t2string :=
    `[Table2]
    s       b       diff    ui      f		i
    string  bool    int     uint    float64	int
    "abc"   true    55      99      2.22	1
    "def"   false   66      88      0		2
    "ghi"   false   66      0       42		3
    "jkl"   false   66      88      NaN		4
    "mno"   false   77      95      0		5
    "pqr"   true    88      97      0		6
    "pqr"   true    88      97      0		6
    `
    table2, err := NewTableFromString(t2string)
    if err != nil {
        log.Println(err)
    }

    fmt.Println(table2)

	// These tables share sort keys i and s.

	// Note that there is a duplicate row,
	// which will be removed during merging.

	// At least one of the tables must have these sort keys set.

	err = table1.SetSortKeys("i", "s")
    if err != nil {
        log.Println(err)
    }

	merged, err := table1.Merge(table2)
    if err != nil {
        log.Println(err)
    }

    fmt.Println(merged)

	// Output:
	// [Table1]
	// XYZ      y s            f   i diff
	// string int string float64 int  int
	// "X"      1 "abc"    1.11    1    7
	// "Y"      3 "ghi"    7.891   3    8
	// "Z"      2 "def"      NaN   2    9
	// "A"      4 "jkl"    0.0     4    6
	// "B"      5 "mno"    0.0     5    4
	// "C"      8 "pqr"      NaN   6   45
	// 
	// [Table2]
	// s      b     diff   ui       f   i
	// string bool   int uint float64 int
	// "abc"  true    55   99    2.22   1
	// "def"  false   66   88    0.0    2
	// "ghi"  false   66    0   42.0    3
	// "jkl"  false   66   88     NaN   4
	// "mno"  false   77   95    0.0    5
	// "pqr"  true    88   97    0.0    6
	// "pqr"  true    88   97    0.0    6
	// 
	// [Merged]
	//   i s      XYZ      y       f diff b       ui
	// int string string int float64  int bool  uint
	//   1 "abc"  "X"      1   1.11     7 true    99
	//   2 "def"  "Z"      2   0.0      9 false   88
	//   3 "ghi"  "Y"      3   7.891    8 false    0
	//   4 "jkl"  "A"      4   0.0      6 false   88
	//   5 "mno"  "B"      5   0.0      4 false   95
	//   6 "pqr"  "C"      8   0.0     45 true    97
}

func ExampleTable_SortUnique() {

	tableString :=
	`[Uniqueness]
	KeyCol number   s
	int float32 string
	2   0       "two point two"
	2   2.2     ""
	1   1.1     "one point one"
	3   3.3     "three point three"
	3   3.3     ""
	3   NaN     "three point three"
	4   0.0     "neither zero nor same X"
	4   NaN     "neither zero nor same Y"
	4   4.4     "neither zero nor same Z"
	4   NaN     "neither zero nor same A"
	5   NaN     "minus 5"
	5   -0      "minus 5"
	5   -5      "minus 5"
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
        log.Println(err)
	}

	fmt.Println("Before SortUnique() ...")
	fmt.Println(table)

	err = table.SetSortKeys("KeyCol")
	if err != nil {
        log.Println(err)
	}

	tableUnique, err := table.SortUnique()
	if err != nil {
        log.Println(err)
	}

	fmt.Println("After SortUnique() ...")
	fmt.Println(tableUnique)

	// Output:
	// Before SortUnique() ...
	// [Uniqueness]
	// KeyCol  number s
	//    int float32 string
	//      2     0.0 "two point two"
	//      2     2.2 ""
	//      1     1.1 "one point one"
	//      3     3.3 "three point three"
	//      3     3.3 ""
	//      3     NaN "three point three"
	//      4     0.0 "neither zero nor same X"
	//      4     NaN "neither zero nor same Y"
	//      4     4.4 "neither zero nor same Z"
	//      4     NaN "neither zero nor same A"
	//      5     NaN "minus 5"
	//      5    -0.0 "minus 5"
	//      5    -5.0 "minus 5"
	// 
	// After SortUnique() ...
	// [Uniqueness]
	// KeyCol  number s
	//    int float32 string
	//      1     1.1 "one point one"
	//      2     2.2 "two point two"
	//      3     3.3 "three point three"
	//      4     4.4 "neither zero nor same A"
	//      5    -5.0 "minus 5"
}

func ExampleTable_GetTableAsCSV() {

	tableString :=
	`[ForCSV]
	first_name  last_name   username    i   f64     b       f32     commas  quotes
	string      string      string      int float64 bool    float32 string  string
	"Rob"       "Pike"      "rob"       1   1.1     true    NaN     ",end"  "\"xyz\""
	"Ken"       "Thompson"  "ken"       3   NaN     true    3.3     "beg,"  "'abc'"
	"Robert"    "Griesemer" "gri"       5   5.5     true    NaN     "md,l"  " \"\" "
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
        log.Println(err)
	}

	fmt.Println("Before table.GetTableAsCSV() ...")
	fmt.Println(table)

	var csv string
	csv, err = table.GetTableAsCSV()
	if err != nil {
        log.Println(err)
	}

	fmt.Println("After table.GetTableAsCSV() ...")
	fmt.Println(csv)

	substituteHeadingNames := []string{"First Name", "Last Name", "username", "i", "f64", "bool", "f32", "Commas", "Quotes"}

	csv, err = table.GetTableAsCSV(substituteHeadingNames...)
	if err != nil {
        log.Println(err)
	}

	fmt.Println("After table.GetTableAsCSV(substituteHeadingNames) ...")
	fmt.Println(csv)

	// Output:
	// Before table.GetTableAsCSV() ...
	// [ForCSV]
	// first_name last_name   username   i     f64 b        f32 commas quotes
	// string     string      string   int float64 bool float32 string string
	// "Rob"      "Pike"      "rob"      1     1.1 true     NaN ",end" "\\\"xyz\\\""
	// "Ken"      "Thompson"  "ken"      3     NaN true     3.3 "beg," "'abc'"
	// "Robert"   "Griesemer" "gri"      5     5.5 true     NaN "md,l" " \\\"\\\" "
	// 
	// After table.GetTableAsCSV() ...
	// first_name,last_name,username,i,f64,b,f32,commas,quotes
	// Rob,Pike,rob,1,1.1,true,,",end","\""xyz\"""
	// Ken,Thompson,ken,3,,true,3.3,"beg,",'abc'
	// Robert,Griesemer,gri,5,5.5,true,,"md,l"," \""\"" "
	// 
	// After table.GetTableAsCSV(substituteHeadingNames) ...
	// First Name,Last Name,username,i,f64,bool,f32,Commas,Quotes
	// Rob,Pike,rob,1,1.1,true,,",end","\""xyz\"""
	// Ken,Thompson,ken,3,,true,3.3,"beg,",'abc'
	// Robert,Griesemer,gri,5,5.5,true,,"md,l"," \""\"" "
}
