package gotable

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
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

/*
//func TestRenameTable2(t *testing.T) {
//	var err error
//	var testsTableSet *TableSet
//	var tests *Table
//	var setupName string = "Fred"
//
//	testsTableSet, err = NewTableSetFromString(
//		`[tests]
//		input		succeeds	output
//		string		bool		string
//		"Barney"	true		"Barney"
//		""			false		"Fred"
//		"$&*"		false		"Fred"
//		`)
//	if err != nil {
//		panic(err)
//	}
//	tests, err = testsTableSet.Table("tests")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("tests = \n%v", tests)
//
//	var table *Table
//
//	for row := 0; row < tests.RowCount(); row++ {
//		// Get test parameters for this row.
//		input, err := tests.GetString("input", row)
//		if err != nil {
//			t.Error(err)
//		}
//		succeeds, err := tests.GetBool("succeeds", row)
//		if err != nil {
//			t.Error(err)
//		}
//		output, err := tests.GetString("output", row)
//		if err != nil {
//			t.Error(err)
//		}
//
//		table, err = NewTable(setupName)
//		if err != nil {
//			t.Error(err)
//		}
//
//		err = table.RenameTable(input)
//		if (err == nil) != succeeds {
//			t.Errorf("Error renaming to %q: %s", output, err)
//		}
//
//		var tableName string = table.Name()
//		if tableName != output {
//			t.Errorf("Expected %q, not %q", output, tableName)
//		}
//	}
//}
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

func ExampleNewTableSet() {
	tableSetName := "MyTableSet"
	tableSet, err := NewTableSet(tableSetName)
	if err != nil {
		panic(err)
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
	//	byte 			// alias for uint8
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

		err = table.SetBool("bVal", rowIndex, true)
		if err != nil {
			t.Error(err)
		}
		bVal, err = table.GetBool("bVal", rowIndex)
		if err != nil {
			t.Error(err)
		}
		if bVal != true {
			t.Errorf("expecting GetBool() value %t, not %t\n", true, bVal)
		}

		err = table.SetBoolByColIndex(colIndex, rowIndex, false)
		if err != nil {
			t.Error(err)
		}
		bVal, err = table.GetBoolByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Error(err)
		}
		if bVal != false {
			t.Errorf("expecting GetBoolByColIndex() value %t, not %t\n", true, bVal)
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

	/*
		table.SetStructShape(true)
		where(table)
	*/
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
			t.Error(fmt.Errorf("expecting col name %s to have prefix \"T_\" or \"F_\" but found: %q", colName, colName))
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
		t.Error(fmt.Errorf("expecting row count of 0, not: %d", rowCount))
	}

	err = table.AppendRow()
	if err != nil {
		t.Error(err)
	}

	rowCount = table.RowCount()
	if rowCount != 1 {
		t.Error(fmt.Errorf("expecting row count of 1, not: %d", rowCount))
	}

	err = table.DeleteRow(0)
	if err != nil {
		t.Error(err)
	}

	rowCount = table.RowCount()
	if rowCount != 0 {
		t.Error(fmt.Errorf("expecting row count of 0, not: %d", rowCount))
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
		t.Error(fmt.Errorf("expecting col count of %d, not: %d", initialColCount, colCount))
	}

	err = table.AppendCol("ExtraCol", "bool")
	if err != nil {
		t.Error(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount+1 {
		t.Error(fmt.Errorf("expecting col count of %d, not: %d", initialColCount+1, colCount))
	}

	lastCol := colCount - 1
	err = table.DeleteColByColIndex(lastCol)
	if err != nil {
		t.Error(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount {
		t.Error(fmt.Errorf("expecting col count of %d, not: %d", initialColCount, colCount))
	}

	err = table.AppendCol("AnotherCol", "string")
	if err != nil {
		t.Error(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount+1 {
		t.Error(fmt.Errorf("expecting col count of %d, not: %d", initialColCount+1, colCount))
	}

	err = table.DeleteCol("AnotherCol")
	if err != nil {
		t.Error(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount {
		t.Error(fmt.Errorf("expecting col count of %d, not: %d", initialColCount, colCount))
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
		t.Error(fmt.Errorf("expecting 1 row less than %d after DeleteRow(%d) but found %d", initialRowCount, deleteRow, rowCount))
	}

	// fmt.Println(table)

	for i := 0; i < table.RowCount(); i++ {
		item, err := table.GetInt("item", i)
		if err != nil {
			t.Error(err)
		}
		if item == deleteRow {
			t.Error(fmt.Errorf("expecting to NOT find item %d after DeleteRow(%d) but found %d", deleteRow, deleteRow, deleteRow))
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
		t.Error(fmt.Errorf("expecting %d row less than %d after DeleteRows(%d, %d) but found %d",
			items, initialRowCount, first, last, rowCount))
	}
	for i := 0; i < table.RowCount(); i++ {
		item, err := table.GetInt("item", i)
		if err != nil {
			t.Error(err)
		}
		if item == first {
			t.Error(fmt.Errorf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d",
				first, first, last, first))
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
		t.Error(fmt.Errorf("expecting %d row less than %d after DeleteRows(%d, %d) but found %d",
			items, initialRowCount, first, last, rowCount))
	}
	for i := 0; i < table.RowCount(); i++ {
		item, err := table.GetInt("item", i)
		if err != nil {
			t.Error(err)
		}
		if item == first {
			t.Error(fmt.Errorf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d",
				first, first, last, first))
		}
		if item == last {
			t.Error(fmt.Errorf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d",
				last, first, last, last))
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
		t.Error(fmt.Errorf("expecting %d row less than %d after DeleteRows(%d, %d) but found %d",
			items, initialRowCount, first, last, rowCount))
	}
	for i := 0; i < table.RowCount(); i++ {
		item, err := table.GetInt("item", i)
		if err != nil {
			t.Error(err)
		}
		if item == first {
			t.Error(fmt.Errorf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d",
				first, first, last, first))
		}
		if item == last {
			t.Error(fmt.Errorf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d",
				last, first, last, last))
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
		t.Error(fmt.Errorf("expecting %d row less than %d after DeleteRows(%d, %d) but found %d",
			items, initialRowCount, first, last, rowCount))
	}
	for i := 0; i < table.RowCount(); i++ {
		item, err := table.GetInt("item", i)
		if err != nil {
			t.Error(err)
		}
		if item == first {
			t.Error(fmt.Errorf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d",
				first, first, last, first))
		}
		if item == last {
			t.Error(fmt.Errorf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d",
				last, first, last, last))
		}
	}
}

func ExampleNewTableFromString() {
	// A table literal. Sometimes easier than constructing a table programmatically.
	tableString := `[MyTable]
		MyBool bool = true
		MyString string = "The answer to life, the universe and everything is forty-two."
		MyInt int = 42`

	table, err := NewTableFromString(tableString)
	if err != nil {
		panic(err)
	}

	// Print the table in its original struct shape.
	fmt.Println(table)

	// Now change its shape to tabular.
	err = table.SetStructShape(false)
	if err != nil {
		panic(err)
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
		panic(err)
	}

	// For testing, we need to write this out to a file so we can read it back.
	fileName := "ExampleNewTableFromFile.txt"
	err = table1.WriteFile(fileName, 0644)
	if err != nil {
		panic(err)
	}

	table2, err := NewTableFromFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println(table2)

	err = table2.SetStructShape(false)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	fileName := "ExampleNewTableFromFileByTableName.txt"
	err = tableSet.WriteFile(fileName, 0644)
	if err != nil {
		panic(err)
	}

	table, err := NewTableFromFileByTableName(fileName, "MyTable")
	if err != nil {
		panic(err)
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
		panic(err)
	}

	fmt.Println(table)

	err = table.DeleteRows(4, 6)
	if err != nil {
		panic(err)
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
		panic(err)
	}

	fmt.Println(table)

	joined, err := table.JoinColVals("command", " | ")
	if err != nil {
		panic(err)
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
		panic(err)
	}

	fmt.Println(table)

	colIndex := 0
	joined, err := table.JoinColValsByColIndex(colIndex, " | ")
	if err != nil {
		panic(err)
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
		t.Error(fmt.Errorf("expecting %s but found: %s", expecting, found))
	}

	expecting = "true"
	found, err = table.GetValAsString("t", 0)
	if err != nil {
		t.Error(err)
	}
	if found != expecting {
		t.Error(fmt.Errorf("expecting %s but found: %s", expecting, found))
	}

	expecting = "23"
	found, err = table.GetValAsString("i", 0)
	if err != nil {
		t.Error(err)
	}
	if found != expecting {
		t.Error(fmt.Errorf("expecting %s but found: %s", expecting, found))
	}

	expecting = "55.5"
	found, err = table.GetValAsString("f", 0)
	if err != nil {
		t.Error(err)
	}
	if found != expecting {
		t.Error(fmt.Errorf("expecting %s but found: %s", expecting, found))
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
		panic(err)
	}

	tables, err := NewTableSetFromFile(actualFileName)
	if err != nil {
		panic(err)
	}

	fileName := tables.FileName()
	if fileName != actualFileName {
		t.Error(fmt.Errorf("Expecting FileName() = %q but found %q", actualFileName, fileName))
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
		t.Error(fmt.Errorf("Expecting tableSetName = %q but found %q", expected, tableSetName))
	}

	
	expected = "Musk"
	tableSet.SetName(expected)
	tableSetName = tableSet.Name()
	if tableSetName != expected {
		t.Error(fmt.Errorf("Expecting tableSetName = %q but found %q", expected, tableSetName))
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
		t.Error(fmt.Errorf("Expecting tableName = %q but found %q", expected, tableName))
	}

	
	expected = "Elon"
	table.SetName(expected)
	tableName = table.Name()
	if tableName != expected {
		t.Error(fmt.Errorf("Expecting tableName = %q but found %q", expected, tableName))
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
			t.Error(fmt.Errorf("Expecting missingValueForType(%q) = %t but found %t",
				test.typeName, test.expected, hasMissing))
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
			t.Error(fmt.Errorf("Expecting preNumberOf(%q) = %d but found %d",
				test.expected, preNumber, preNumber))
		}
	}
}
