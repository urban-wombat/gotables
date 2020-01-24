package gotables

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/urban-wombat/util"
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
			t.Fatal(err)
		}

		err = table.RenameTable(test.input)
		if (err == nil) != test.succeeds {
			t.Fatalf("Error renaming to %q: %s", test.output, err)
		}

		var tableName string = table.Name()
		if tableName != test.output {
			t.Fatalf("Expected %q, not %q", test.output, tableName)
		}

	}

	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
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
			t.Fatal(err)
		}
		err = tableSet.RenameTable(test.renameFrom, test.renameTo)
		if (err == nil) != test.succeeds {
			t.Fatalf("test[%d]: Error renaming from %q to %q: %v", i, test.renameFrom, test.renameTo, err)
		}
		if isValid, err := tableSet.IsValidTableSet(); !isValid {
			t.Fatal(err)
		}
	}
}

func TestNewTableFromString01(t *testing.T) {
	tableSet, err := NewTableSetFromString(
		`[EmptyTable1]

		[EmptyTable2]

		[TableWithColNamesAndTypes]
		A	B	C
		int	int	int

		[TableWithRow]
		D	E	F
		int	int	int
		1	2	31

		[TableWithRows]
		G	H	I
		int	int	int
		4	5	6
		7	8	9
	`)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := tableSet.IsValidTableSet(); !isValid {
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
			t.Fatalf("[%d] %v", i, err)
		}
		if (err == nil) != test.succeeds {
			t.Fatalf("test[%d]: err == %v but expecting succeeds == %t", i, err, test.succeeds)
		}

		colCount := table.ColCount()
		if colCount != test.colCount {
			t.Fatalf("test[%d]: expecting [%s] colCount %d, not %d\n", i, test.tableName, test.colCount, colCount)
		}
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		rowCount := table.RowCount()
		if rowCount != test.rowCount {
			t.Fatalf("test[%d]: expecting [%s] rowCount %d, not %d\n", i, test.tableName, test.rowCount, rowCount)
		}
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}
	}
}

func TestNewTableFromString02(t *testing.T) {
	_, err := NewTableSetFromString(
		`[EmptyTable1]

		# Should be a syntax error. Table should have both names AND types.
		[TableWithColNamesOnly]
		A	B	C

		[EmptyTable2]
	`)
	if err == nil {
		t.Fatalf("Should return a syntax error. Table should have both names AND types.")
	}
}

func TestNewTableFromString03(t *testing.T) {
	_, err := NewTableSetFromString(
		`[TableWithRow]
		D	E	F
		int	int	int
		1	2	32

		A	B	C
	`)
	if err == nil {
		t.Fatalf("Should return a syntax error. Col names should not follow blank lines.")
	}
}

func TestNewTableFromString04(t *testing.T) {
	_, err := NewTableSetFromString(
		`[TableWithRow]
		D	E	F
		int	int	int
		1	2	33

		4	5	6
		`)
	if err == nil {
		t.Fatalf("Should return a syntax error. Col values should not follow blank lines.")
	}
}

func TestNewTableFromString05(t *testing.T) {
	tableSet, err := NewTableSetFromString(
		`[TableEmpty]
		
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
		{"TableEmpty", 0, 0, true},
	}

	for i, test := range tests {
		table, err := tableSet.Table(test.tableName)
		if err != nil {
			t.Fatalf("[%d] %v", i, err)
		}
		if (err == nil) != test.succeeds {
			t.Fatalf("test[%d]: err == %v but expecting succeeds == %t", i, err, test.succeeds)
		}

		colCount := table.ColCount()
		if colCount != test.colCount {
			t.Fatalf("test[%d]: expecting [%s] colCount %d, not %d\n", i, test.tableName, test.colCount, colCount)
		}

		rowCount := table.RowCount()
		if rowCount != test.rowCount {
			t.Fatalf("test[%d]: expecting [%s] rowCount %d, not %d\n", i, test.tableName, test.rowCount, rowCount)
		}
	}
	if isValid, err := tableSet.IsValidTableSet(); !isValid {
		t.Fatal(err)
	}
}

func TestNewTableFromString06(t *testing.T) {
	tableSet, err := NewTableSetFromString(
		`[TableStruct]
		i int = 42
		j int = 44

		[Empty]

		`)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := tableSet.IsValidTableSet(); !isValid {
		t.Fatal(err)
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
			t.Fatalf("[%d] %v", i, err)
		}
		if (err == nil) != test.succeeds {
			t.Fatalf("test[%d]: err == %v but expecting succeeds == %t", i, err, test.succeeds)
		}

		colCount := table.ColCount()
		if colCount != test.colCount {
			t.Fatalf("test[%d]: expecting [%s] colCount %d, not %d\n", i, test.tableName, test.colCount, colCount)
		}

		rowCount := table.RowCount()
		if rowCount != test.rowCount {
			t.Fatalf("test[%d]: expecting [%s] rowCount %d, not %d\n", i, test.tableName, test.rowCount, rowCount)
		}
		if isValid, err := tableSet.IsValidTableSet(); !isValid {
			t.Fatal(err)
		}

	}
	if isValid, err := tableSet.IsValidTableSet(); !isValid {
		t.Fatal(err)
	}
}

func TestNewTableFromString07(t *testing.T) {
	_, err := NewTableSetFromString(
		`[TableStruct]
		i int = 42
		j int = 44
		# Expecting more structs or a blank line.
		X Y Z

		[Empty]

	`)
	if err == nil {
		t.Fatal(err)
	}
}

func TestNewTableFromString08(t *testing.T) {
	_, err := NewTableSetFromString(
		`[TableShaped]
		X Y Z
		# Expecting col types, not structs.
		i int = 42
		j int = 44

		[Empty]

		`)
	if err == nil {
		t.Fatal(err)
	}
}

// Testing struct using = with zero rows.
func TestNewTableFromString09(t *testing.T) {
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
			t.Fatalf("test[%d]: expecting [%s] colCount %d, not %d\n", i, test.tableName, test.colCount, colCount)
		}

		rowCount := table.RowCount()
		if rowCount != test.rowCount {
			t.Fatalf("test[%d]: expecting [%s] rowCount %d, not %d\n", i, test.tableName, test.rowCount, rowCount)
		}
	}

	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

// 02.05.2017
// Testing struct without = having zero rows.
// This is a struct format change to have = only if there is a value following it.
func TestNewTableFromString10(t *testing.T) {
	table, err := NewTableFromString(
		`[BlankTableStruct]
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
		{"BlankTableStruct", 2, 0, true},
	}

	for i, test := range tests {

		colCount := table.ColCount()
		if colCount != test.colCount {
			t.Fatalf("test[%d]: expecting [%s] colCount %d, not %d\n", i, test.tableName, test.colCount, colCount)
		}

		rowCount := table.RowCount()
		if rowCount != test.rowCount {
			t.Fatalf("test[%d]: expecting [%s] rowCount %d, not %d\n", i, test.tableName, test.rowCount, rowCount)
		}
	}

	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

// 02.05.2017
// Testing struct with name type = value
// This is a struct format change to have = only if there is a value following it.
func TestNewTableFromString11(t *testing.T) {
	table, err := NewTableFromString(
		`[ValuesTableStruct]
		i int = 1
		j int = 2
		s string = "ABC"
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
		{"ValuesTableStruct", 3, 1, true},
	}

	for i, test := range tests {

		colCount := table.ColCount()
		if colCount != test.colCount {
			t.Fatalf("test[%d]: expecting [%s] colCount %d, not %d\n", i, test.tableName, test.colCount, colCount)
		}

		rowCount := table.RowCount()
		if rowCount != test.rowCount {
			t.Fatalf("test[%d]: expecting [%s] rowCount %d, not %d\n", i, test.tableName, test.rowCount, rowCount)
		}
	}

	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

// 02.05.2017
// Testing struct with name type = value
// This is a struct format change to have = only if there is a value following it.
func TestNewTableFromString12(t *testing.T) {
	_, err := NewTableFromString(
		`[InvalidTableStruct]
		i int =
		j int =
		s string =
	`)
	if err == nil {
		t.Fatal(err)
	}
}

// 02/09.2017
// Testing table with slice of uint: []uint
func TestNewTableFromString13(t *testing.T) {
	var err error

	// Should pass: all values are in uint8 range
	_, err = NewTableFromString(
		`[TableWithUintSlice]
		uintNums
		[]uint8
		[]
		[0 1 255 3 4 1 5]
	`)
	if err != nil {
		t.Fatal(err)
	}

	// Should fail: uint8 -1
	_, err = NewTableFromString(
		`[TableWithUintSlice]
		uintNums
		[]uint8
		[0 -1 255 3 4]
	`)
	if err == nil {
		t.Fatal(err)
	}

	// Should fail: uint8 256
	_, err = NewTableFromString(
		`[TableWithByteSlice]
		uintNums
		[]byte
		[0 1 256 3 4 2 6]
	`)
	if err == nil {
		t.Fatal(err)
	}

	_, err = NewTableFromString(
		`[TableWithByteSlice]
		uintNums []byte = [0 1 256 3 4 3 7]
		i int = 42
		b []byte = [1 1 255]
		u []uint8 = [2 2 255 2]
		f float32 = 32
		x []byte = []
		s string = "In Between ..."
		y []uint8 = []
		j int64 = 99
	`)
	// Invalid uint 256
	if err == nil {
		t.Fatal(err)
	}
}

// 03/09.2017
// Testing table with slice of uint: []uint
func TestNewTableFromString14(t *testing.T) {
	var err error
	s := `[TableX]
	i   x                   f           bb                  s       b
	int []uint8             float64     []byte              string  byte
	1   [10 11 12 13 14]    1           [90 81 72 63]       "one"   15
	2   [20 21 22 23]       2           [90 81 72 63]       "two"   22
	3   [30 31 32]          3           [90 81 72]          "three" 33
	4   [40 41 42 43 44]    4           [90 81 72 63 255]   "four"  44
	`
	table, err := NewTableFromString(s)
	if err != nil {
		t.Fatal(err)
	}

	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
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
		t.Fatal(err)
	}

	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	err = table.AppendCol("bite", "[]byte")
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

//	// Testing table with escaped characters.
//	func TestNewTableFromString15(t *testing.T) {
//		var err error
//
//		s :=
//		`[Table]
//		s		chars
//		str   	[]byte
//		"Fred"	[22]
//		`
//	//	"\""	[34]
//
//		table, err := NewTableFromString(s)
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		if isValid, err := table.IsValidTable(); !isValid { t.Fatal(err) }
//
//	/*
//		for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
//			s, err := table.GetString("s", rowIndex)
//			if err != nil {
//				t.Fatal(err)
//			}
//
//			chars, err := table.GetByteSlice("chars", rowIndex)
//			if err != nil {
//				t.Fatal(err)
//			}
//
//			if len(s) != len(chars) {
//				t.Fatalf("len(%q)=%d != len(%v)=%d", s, len(s), chars, len(chars))
//			}
//
//			for i := 0; i < table.RowCount(); i++ {
//				if s[i] != chars[i] {
//					t.Fatalf("s[%d] != chars[%d]: '%c' != %d", i, i, s[i], chars[i])
//				}
//			}
//		}
//	*/
//	}

// Test surprise syntax error discovered using Go Playground.
// It seems we've never tested a 2-column table.
func TestNewTableFromString16(t *testing.T) {
	var err error
	var s string

	s = `[Table]
	i
	int
	`
	_, err = NewTableFromString(s)
	if err != nil {
		t.Fatal(err)
	}

	s = `[Table]
	i	j
	int int
	`
	_, err = NewTableFromString(s)
	if err != nil {
		//panic(err)
		t.Fatal(err)
	}

	s = `[Table]
	i	j	k
	int int int
	`
	_, err = NewTableFromString(s)
	if err != nil {
		t.Fatal(err)
	}

	s = `[Table]
	i	j	k	l
	int int int	int
	`
	_, err = NewTableFromString(s)
	if err != nil {
		t.Fatal(err)
	}

	s = `[Table]
	i	j	k	l	m
	int int int	int	int
	`
	_, err = NewTableFromString(s)
	if err != nil {
		t.Fatal(err)
	}

	s = `[Table]
	i   x
	int []uint8
	1   [10 11 12 13 20]
	2   [20 21 22 23]
	3   [30 31 32]
	4   [40 41 42 43 44]
	`
	_, err = NewTableFromString(s)
	if err != nil {
		t.Fatal(err)
	}

	s = `[Table]
	x		y
	float32	float64
	1.0		1.1
	2.0		2.2
	3.0		3.3
	`
	_, err = NewTableFromString(s)
	if err != nil {
		t.Fatal(err)
	}

	s = `[Table]
	i   x
	int []uint8
	`
	_, err = NewTableFromString(s)
	if err != nil {
		t.Fatal(err)
	}

	s = `[Table]
	x		y
	float32	float64
	`
	_, err = NewTableFromString(s)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewTableFromString(s)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewTableFromString(s)
	if err != nil {
		t.Fatal(err)
	}

	s = `[Table]
	x	y	=
	`
	// looks like struct but missing value after equals: x y =
	_, err = NewTableFromString(s)
	if err == nil {
		t.Fatal(err)
	}

	s = `[Table]
	x string = "one"
	`
	_, err = NewTableFromString(s)
	if err != nil {
		t.Fatal(err)
	}

	s = `[Table]
	x string = "one two"
	`
	_, err = NewTableFromString(s)
	if err != nil {
		t.Fatal(err)
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

	rounded := util.Round(numberToRound, places)
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
		rounded := util.Round(test.val, test.places)
		if rounded != test.rounded {
			t.Fatalf("test[%d]: expecting %f, not %f", i, test.rounded, rounded)
		}
	}
}

func TestSetAndGetFunctions(t *testing.T) {
	var bVal bool
	var byteVal byte // alias for uint8
	var byteSlice []byte
	var ui8Slice []uint8
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
		t.Fatal(err)
	}

	// Note: Tests are collected inside code blocks for human readability.

	{ // bool tests

		err = table.AppendCol("bVal", "bool")
		if err != nil {
			t.Fatal(err)
		}

		// After first col has been appended.
	}

	if table.RowCount() == 0 {
		err = table.AppendRow()
		if err != nil {
			t.Fatal(err)
		}
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		expected := true
		err = table.SetBool("bVal", rowIndex, expected)
		if err != nil {
			t.Fatal(err)
		}
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		bVal, err = table.GetBool("bVal", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		if bVal != expected {
			t.Fatalf("expecting GetBool(bVal) value %t, not %t\n", expected, bVal)
		}
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		expected = false
		err = table.SetBoolByColIndex(colIndex, rowIndex, expected)
		if err != nil {
			t.Fatal(err)
		}

		bVal, err = table.GetBoolByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if bVal != expected {
			t.Fatalf("expecting GetBoolByColIndex() value %t, not %t\n", expected, bVal)
		}
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}
	}

	{ // float32 tests

		err = table.AppendCol("f32Val", "float32")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetFloat32("f32Val", rowIndex, 55.1)
		if err != nil {
			t.Fatal(err)
		}
		f32Val, err = table.GetFloat32("f32Val", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if f32Val != 55.1 {
			t.Fatalf("expecting GetFloat32() value %f, not %f\n", 55.1, f32Val)
		}

		err = table.SetFloat32ByColIndex(colIndex, rowIndex, 66.1)
		if err != nil {
			t.Fatal(err)
		}
		f32Val, err = table.GetFloat32ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if f32Val != 66.1 {
			t.Fatalf("expecting GetFloat32ByColIndex() value %f, not %f\n", 66.1, f32Val)
		}
	}

	{ // float64 tests

		err = table.AppendCol("f64Val", "float64")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetFloat64("f64Val", rowIndex, 88.1)
		if err != nil {
			t.Fatal(err)
		}
		f64Val, err = table.GetFloat64("f64Val", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if f64Val != 88.1 {
			t.Fatalf("expecting GetFloat64() value %f, not %f\n", 88.1, f64Val)
		}

		err = table.SetFloat64ByColIndex(colIndex, rowIndex, 77.1)
		if err != nil {
			t.Fatal(err)
		}
		f64Val, err = table.GetFloat64ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if f64Val != 77.1 {
			t.Fatalf("expecting GetFloat64ByColIndex() value %f, not %f\n", 77.1, f64Val)
		}
	}

	{ // int tests

		err = table.AppendCol("iVal", "int")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetInt("iVal", rowIndex, 55)
		if err != nil {
			t.Fatal(err)
		}
		iVal, err = table.GetInt("iVal", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if iVal != 55 {
			t.Fatalf("expecting GetInt() value %d, not %d\n", 55, iVal)
		}

		err = table.SetIntByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Fatal(err)
		}
		iVal, err = table.GetIntByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if iVal != 66 {
			t.Fatalf("expecting GetIntByColIndex() value %d, not %d\n", 66, iVal)
		}
	}

	{ // int16 tests

		err = table.AppendCol("i16Val", "int16")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetInt16("i16Val", rowIndex, 55)
		if err != nil {
			t.Fatal(err)
		}
		i16Val, err = table.GetInt16("i16Val", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if i16Val != 55 {
			t.Fatalf("expecting GetInt16() value %d, not %d\n", 55, i16Val)
		}

		err = table.SetInt16ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Fatal(err)
		}
		i16Val, err = table.GetInt16ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if i16Val != 66 {
			t.Fatalf("expecting GetInt16ByColIndex() value %d, not %d\n", 66, i16Val)
		}
	}

	{ // int32 tests

		err = table.AppendCol("i32Val", "int32")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetInt32("i32Val", rowIndex, 55)
		if err != nil {
			t.Fatal(err)
		}
		i32Val, err = table.GetInt32("i32Val", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if i32Val != 55 {
			t.Fatalf("expecting GetInt32() value %d, not %d\n", 55, i32Val)
		}

		err = table.SetInt32ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Fatal(err)
		}
		i32Val, err = table.GetInt32ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if i32Val != 66 {
			t.Fatalf("expecting GetInt32ByColIndex() value %d, not %d\n", 66, i32Val)
		}
	}

	{ // int64 tests

		err = table.AppendCol("i64Val", "int64")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetInt64("i64Val", rowIndex, 55)
		if err != nil {
			t.Fatal(err)
		}
		i64Val, err = table.GetInt64("i64Val", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if i64Val != 55 {
			t.Fatalf("expecting GetInt64() value %d, not %d\n", 55, i64Val)
		}

		err = table.SetInt64ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Fatal(err)
		}
		i64Val, err = table.GetInt64ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if i64Val != 66 {
			t.Fatalf("expecting GetInt64ByColIndex() value %d, not %d\n", 66, i64Val)
		}
	}

	{ // int8 tests

		err = table.AppendCol("i8Val", "int8")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetInt8("i8Val", rowIndex, 55)
		if err != nil {
			t.Fatal(err)
		}
		i8Val, err = table.GetInt8("i8Val", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if i8Val != 55 {
			t.Fatalf("expecting GetInt8() value %d, not %d\n", 55, i8Val)
		}

		err = table.SetInt8ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Fatal(err)
		}
		i8Val, err = table.GetInt8ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if i8Val != 66 {
			t.Fatalf("expecting GetInt8ByColIndex() value %d, not %d\n", 66, i8Val)
		}
	}

	{ // uint tests

		err = table.AppendCol("uiVal", "uint")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetUint("uiVal", rowIndex, 55)
		if err != nil {
			t.Fatal(err)
		}
		uiVal, err = table.GetUint("uiVal", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if uiVal != 55 {
			t.Fatalf("expecting GetUint() value %d, not %d\n", 55, uiVal)
		}

		err = table.SetUintByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Fatal(err)
		}
		uiVal, err = table.GetUintByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if uiVal != 66 {
			t.Fatalf("expecting GetUintByColIndex() value %d, not %d\n", 66, uiVal)
		}
	}

	{ // uint16 tests

		err = table.AppendCol("ui16Val", "uint16")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetUint16("ui16Val", rowIndex, 55)
		if err != nil {
			t.Fatal(err)
		}
		ui16Val, err = table.GetUint16("ui16Val", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if ui16Val != 55 {
			t.Fatalf("expecting GetUint16() value %d, not %d\n", 55, ui16Val)
		}

		err = table.SetUint16ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Fatal(err)
		}
		ui16Val, err = table.GetUint16ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if ui16Val != 66 {
			t.Fatalf("expecting GetUint16ByColIndex() value %d, not %d\n", 66, ui16Val)
		}
	}

	{ // uint32 tests

		err = table.AppendCol("ui32Val", "uint32")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetUint32("ui32Val", rowIndex, 55)
		if err != nil {
			t.Fatal(err)
		}
		ui32Val, err = table.GetUint32("ui32Val", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if ui32Val != 55 {
			t.Fatalf("expecting GetUint32() value %d, not %d\n", 55, ui32Val)
		}

		err = table.SetUint32ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Fatal(err)
		}
		ui32Val, err = table.GetUint32ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if ui32Val != 66 {
			t.Fatalf("expecting GetUint32ByColIndex() value %d, not %d\n", 66, ui32Val)
		}
	}

	{ // uint64 tests

		err = table.AppendCol("ui64Val", "uint64")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetUint64("ui64Val", rowIndex, 55)
		if err != nil {
			t.Fatal(err)
		}
		ui64Val, err = table.GetUint64("ui64Val", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if ui64Val != 55 {
			t.Fatalf("expecting GetUint64() value %d, not %d\n", 55, ui64Val)
		}

		err = table.SetUint64ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Fatal(err)
		}
		ui64Val, err = table.GetUint64ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if ui64Val != 66 {
			t.Fatalf("expecting GetUint64ByColIndex() value %d, not %d\n", 66, ui64Val)
		}
	}

	{ // uint8 tests

		err = table.AppendCol("ui8Val", "uint8")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetUint8("ui8Val", rowIndex, 55)
		if err != nil {
			t.Fatal(err)
		}
		ui8Val, err = table.GetUint8("ui8Val", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if ui8Val != 55 {
			t.Fatalf("expecting GetUint8() value %d, not %d\n", 55, ui8Val)
		}

		err = table.SetUint8ByColIndex(colIndex, rowIndex, 66)
		if err != nil {
			t.Fatal(err)
		}
		ui8Val, err = table.GetUint8ByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if ui8Val != 66 {
			t.Fatalf("expecting GetUint8ByColIndex() value %d, not %d\n", 66, ui8Val)
		}
	}

	{ // uint8[] slice tests

		err = table.AppendCol("ui8Slice", "[]uint8")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetByteSlice("ui8Slice", rowIndex, []uint8{0, 1, 2})
		if err != nil {
			t.Fatal(err)
		}
		ui8Slice, err = table.GetByteSlice("ui8Slice", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(ui8Slice, []uint8{0, 1, 2}) { // Slices not equal.
			t.Fatalf("expecting GetByteSlice() value %d, not %d\n", []uint8{0, 1, 2}, ui8Slice)
		}

		err = table.SetByteSliceByColIndex(colIndex, rowIndex, []uint8{2, 4, 6})
		if err != nil {
			t.Fatal(err)
		}
		ui8Slice, err = table.GetByteSliceByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(ui8Slice, []uint8{2, 4, 6}) { // Slices not equal.
			t.Fatalf("expecting GetByteSliceByColIndex() value %d, not %d\n", []uint8{2, 4, 6}, ui8Slice)
		}
	}

	{ // byte tests

		err = table.AppendCol("byteVal", "byte")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetByte("byteVal", rowIndex, 56)
		if err != nil {
			t.Fatal(err)
		}
		byteVal, err = table.GetByte("byteVal", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if byteVal != 56 {
			t.Fatalf("expecting GetByte() value %d, not %d\n", 56, byteVal)
		}

		err = table.SetByteByColIndex(colIndex, rowIndex, 67)
		if err != nil {
			t.Fatal(err)
		}
		byteVal, err = table.GetByteByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if byteVal != 67 {
			t.Fatalf("expecting GetByteByColIndex() value %d, not %d\n", 67, byteVal)
		}
	}

	{ // byte[] slice tests

		err = table.AppendCol("byteSlice", "[]byte")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetByteSlice("byteSlice", rowIndex, []byte{4, 5, 6})
		if err != nil {
			t.Fatal(err)
		}
		byteSlice, err = table.GetByteSlice("byteSlice", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(byteSlice, []byte{4, 5, 6}) { // Slices not equal.
			t.Fatalf("expecting GetByteSlice() value %d, not %d\n", []byte{4, 5, 6}, byteSlice)
		}

		err = table.SetByteSliceByColIndex(colIndex, rowIndex, []byte{7, 8, 9})
		if err != nil {
			t.Fatal(err)
		}
		byteSlice, err = table.GetByteSliceByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(byteSlice, []byte{7, 8, 9}) { // Slices not equal.
			t.Fatalf("expecting GetByteSliceByColIndex() value %d, not %d\n", []byte{7, 8, 9}, byteSlice)
		}
	}

	{ // string tests

		err = table.AppendCol("sVal", "string")
		if err != nil {
			t.Fatal(err)
		}
		colIndex += 1
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.SetString("sVal", rowIndex, "55")
		if err != nil {
			t.Fatal(err)
		}
		sVal, err = table.GetString("sVal", rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if sVal != "55" {
			t.Fatalf("expecting GetString() value %q, not %q\n", "55", sVal)
		}

		err = table.SetStringByColIndex(colIndex, rowIndex, "66")
		if err != nil {
			t.Fatal(err)
		}
		sVal, err = table.GetStringByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}
		if sVal != "66" {
			t.Fatalf("expecting GetStringByColIndex() value %q, not %q\n", "66", sVal)
		}
	}

	//	fmt.Println(table)

	var shape bool
	var expected bool = true
	err = table.SetStructShape(expected)
	if err != nil {
		t.Fatal(err)
	}
	shape, err = table.IsStructShape()
	if err != nil {
		t.Fatal(err)
	}
	if shape != expected {
		t.Fatalf("expecting [%s].IsStructShape() value %t, not %t\n", table.Name(), expected, shape)
	}

	//	fmt.Println(table)

	expected = false
	err = table.SetStructShape(expected)
	if err != nil {
		t.Fatal(err)
	}
	shape, err = table.IsStructShape()
	if err != nil {
		t.Fatal(err)
	}
	if shape != expected {
		t.Fatalf("expecting [%s].IsStructShape() value %t, not %t\n", table.Name(), expected, shape)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func TestSetIntegerMinAndMax(t *testing.T) {
	var err error

	// For testing machine-dependent types
	var intBits int = strconv.IntSize // uint and int are the same bit size. Like sizeof()
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
		t.Fatalf(msg)
	}

	var tests = []struct {
		input string
		valid bool
	}{

		// uint8
		{`
			[uint8_]
			i uint8 = 0`,
			true,
		},
		{`
			[uint8_]
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
			[uint16_]
			i uint16 = 0`,
			true,
		},
		{`
			[uint16_]
			i uint16 = 65535`,
			true,
		},
		{`
			[uint16_]
			i uint16 = -1`,
			false,
		},
		{`
			[uint16_]
			i uint16 = 65536`,
			false,
		},

		// uint32
		{`
			[uint32_]
			i uint32 = 0`,
			true,
		},
		{`
			[uint32_]
			i uint32 = 4294967295`,
			true,
		},
		{`
			[uint32_]
			i uint32 = -1`,
			false,
		},
		{`
			[uint32_]
			i uint32 = 4294967296`,
			false,
		},

		// uint64
		{`
			[uint64_]
			i uint64 = 0`,
			true,
		},
		{`
			[uint64_]
			i uint64 = 18446744073709551615`,
			true,
		},
		{`
			[uint64_]
			i uint64 = -1`,
			false,
		},
		{`
			[uint64_]
			i uint64 = 18446744073709551616`,
			false,
		},

		// uint
		{`
			[uint_]
			i uint = 0`,
			true,
		},
		{fmt.Sprintf(`
			[uint_]
			i uint = %d`, uintMaxVal),
			true,
		},
		{`
			[uint_]
			i uint = -1`,
			false,
		},
		// Note: Cannot easily test machine-dependent types outside range values (except for uint 0 and -1), so skipping them.

		// int8
		{`
			[int8_]
			i int8 = -128`,
			true,
		},
		{`
			[int8_]
			i int8 = 127`,
			true,
		},
		{`
			[int8_]
			i int8 = -129`,
			false,
		},
		{`
			[int8_]
			i int8 = 128`,
			false,
		},

		// int16
		{`
			[int16_]
			i int16 = -32768`,
			true,
		},
		{`
			[int16_]
			i int16 = 32767`,
			true,
		},
		{`
			[int16_]
			i int16 = -32769`,
			false,
		},
		{`
			[int16_]
			i int16 = 32768`,
			false,
		},

		// int32
		{`
			[int32_]
			i int32 = -2147483648`,
			true,
		},
		{`
			[int32_]
			i int32 = 2147483647`,
			true,
		},
		{`
			[int32_]
			i int32 = -2147483649`,
			false,
		},
		{`
			[int32_]
			i int32 = 2147483648`,
			false,
		},

		// int64
		{`
			[int64_]
			i int64 = -9223372036854775808`,
			true,
		},
		{`
			[int64_]
			i int64 = 9223372036854775807`,
			true,
		},
		{`
			[int64_]
			i int64 = -9223372036854775809`,
			false,
		},
		{`
			[int64_]
			i int64 = 9223372036854775808`,
			false,
		},

		// int
		{fmt.Sprintf(`
			[int_]
			i int = %d`, intMinVal),
			true,
		},
		{fmt.Sprintf(`
			[int_]
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
				t.Fatalf("test[%d]: %v", i, err)
			case false:
				t.Fatalf("test[%d]: NewTableSetFromString(): expecting this input to fail with a range error: %s", i, test.input)
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
	var intBits int = strconv.IntSize // uint and int are the same bit size. Like sizeof()
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
		t.Fatalf(msg)
	}

	for i, test := range tests {
		_, err = NewTableSetFromString(test.input)
		if err == nil != test.valid {
			switch test.valid {
			case true:
				t.Fatalf("test[%d]: %v", i, err)
			case false:
				t.Fatalf("test[%d]: NewTableSetFromString(): expecting this input to fail with a range error: %s", i, test.input)
			}
		}
	}
}

var globalTableSetString string = `
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
	r rune = '世'
    `

func BenchmarkNewTableSetFromString(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		_, err = NewTableSetFromString(globalTableSetString)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkTableSetToString(b *testing.B) {
	// Set up for benchmark.
	tableSet, err := NewTableSetFromString(globalTableSetString)
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
"Jupiter" 318.0        5.2    79     4 "just"
"Saturn"   95.0        9.5    82     5 "sat"
"Uranus"   15.0       19.2    27     6 "upon"
"Neptune"  17.0       30.6    13     7 "nine"
"Pluto"     0.002     39.4     5     8 "porcupines"
`

var planets_unpadded = `[planets_unpadded]
name mass distance moons index mnemonic
string float64 float64 int int string
"Mercury" 0.055 0.4 0 0 "my"
"Venus" 0.815 0.7 0 1 "very"
"Earth" 1 1 1 2 "elegant"
"Mars" 0.107 1.5 2 3 "mother"
"Jupiter" 318 5.2 79 4 "just"
"Saturn" 95 9.5 82 5 "sat"
"Uranus" 15 19.2 27 6 "upon"
"Neptune" 17 30.6 13 7 "nine"
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

/*
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
*/

/*
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
*/

func BenchmarkNewTableSetFromFile(b *testing.B) {
	// Set up for benchmark.
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

	for i := 0; i < b.N; i++ {
		_, err := NewTableSetFromFile(fileName)
		if err != nil {
			b.Error(err)
		}
	}
}

func TestIsNumericColType(t *testing.T) {
	tableString := `
    [tbl]
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

	table, err := tableSet.Table("tbl")
	if err != nil {
		t.Fatal(err)
	}

	for colIndex := 0; colIndex < table.ColCount(); colIndex++ {

		colName, err := table.ColName(colIndex)
		if err != nil {
			t.Fatal(err)
		}

		colType, err := table.ColTypeByColIndex(colIndex)
		if err != nil {
			t.Fatal(err)
		}

		isNumeric, _ := IsNumericColType(colType)

		hasPrefixT := strings.HasPrefix(colName, "T_")
		hasPrefixF := strings.HasPrefix(colName, "F_")
		if !hasPrefixT && !hasPrefixF {
			t.Fatalf("expecting col name %s to have prefix \"T_\" or \"F_\" but found: %q", colName, colName)
		}

		if isNumeric != hasPrefixT {
			err := fmt.Errorf("col %s type %s unexpected IsNumeric: %t", colName, colType, isNumeric)
			t.Fatal(err)
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func TestTable_AppendRow_DeleteRow(t *testing.T) {
	tableString := `
    [my_table]
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
	if isValid, err := tableSet.IsValidTableSet(); !isValid {
		t.Fatal(err)
	}

	table, err := tableSet.Table("my_table")
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	rowCount := table.RowCount()
	if rowCount != 0 {
		t.Fatalf("expecting row count of 0, not: %d", rowCount)
	}

	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	rowCount = table.RowCount()
	if rowCount != 1 {
		t.Fatalf("expecting row count of 1, not: %d", rowCount)
	}

	err = table.DeleteRow(0)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	rowCount = table.RowCount()
	if rowCount != 0 {
		t.Fatalf("expecting row count of 0, not: %d", rowCount)
	}
}

func TestColCount(t *testing.T) {
	tableString := `
    [tbl]
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

	table, err := tableSet.Table("tbl")
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	const initialColCount = 14

	colCount := table.ColCount()
	if colCount != initialColCount {
		t.Fatalf("expecting col count of %d, not: %d", initialColCount, colCount)
	}

	err = table.AppendCol("ExtraCol", "bool")
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount+1 {
		t.Fatalf("expecting col count of %d, not: %d", initialColCount+1, colCount)
	}

	lastCol := colCount - 1
	err = table.DeleteColByColIndex(lastCol)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount {
		t.Fatalf("expecting col count of %d, not: %d", initialColCount, colCount)
	}

	err = table.AppendCol("AnotherCol", "string")
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount+1 {
		t.Fatalf("expecting col count of %d, not: %d", initialColCount+1, colCount)
	}

	err = table.DeleteCol("AnotherCol")
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount {
		t.Fatalf("expecting col count of %d, not: %d", initialColCount, colCount)
	}
}

func TestDeleteCol(t *testing.T) {
	// Note: the table must have at least 1 row of data to fully
	//       test DeleteCol(). Column of row cells must be deleted too.
	tableString := `
    [tbl]
	F_bool bool = true
	F_string string = "yes"
	T_float32 float32 = 33.3
	T_float64 float64 = 64.4
	T_int int = 1
	T_int16 int16 = 16
	T_int32 int32 = 32
	T_int64 int64 = 64
	T_int8 int8 = 8
	T_uint uint = 99
	T_uint16 uint16 = 111
	T_uint32 uint32 = 44
	T_uint64 uint64 = 22
	T_uint8 uint8 = 33
    `

	tableSet, err := NewTableSetFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := tableSet.IsValidTableSet(); !isValid {
		t.Fatal(err)
	}

	table, err := tableSet.Table("tbl")
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	err = table.DeleteCol("T_uint32")
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func TestDeleteRow(t *testing.T) {
	tableString := `
	[tbl]
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
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	initialRowCount := table.RowCount()

	const deleteRow = 4

	err = table.DeleteRow(deleteRow)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	rowCount := table.RowCount()
	if rowCount != initialRowCount-1 {
		t.Fatalf("expecting 1 row less than %d after DeleteRow(%d) but found %d", initialRowCount, deleteRow, rowCount)
	}

	// fmt.Println(table)

	for i := 0; i < table.RowCount(); i++ {
		item, err := table.GetInt("item", i)
		if err != nil {
			t.Fatal(err)
		}
		if item == deleteRow {
			t.Fatalf("expecting to NOT find item %d after DeleteRow(%d) but found %d", deleteRow, deleteRow, deleteRow)
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
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	initialRowCount := table.RowCount()

	// Test invalid row index range: first greater than last
	err = table.DeleteRows(5, 4)
	if err == nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	// Test invalid row index range: first below zero
	err = table.DeleteRows(-1, 4)
	if err == nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	// Test invalid row index range: last above initialRowCount-1
	err = table.DeleteRows(0, initialRowCount)
	if err == nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
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
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	// fmt.Println(table)
	rowCount = table.RowCount()
	if rowCount != initialRowCount-items {
		t.Fatalf("expecting %d row less than %d after DeleteRows(%d, %d) but found %d", items, initialRowCount, first, last, rowCount)
	}
	for i := 0; i < table.RowCount(); i++ {
		var item int
		item, err = table.GetInt("item", i)
		if err != nil {
			t.Fatal(err)
		}
		if item == first {
			t.Fatalf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d", first, first, last, first)
		}
	}

	// Test delete 2 items: 4 to 5
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	first = 4
	last = 5
	items = last - first + 1
	err = table.DeleteRows(first, last)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	// fmt.Println(table)
	rowCount = table.RowCount()
	if rowCount != initialRowCount-items {
		t.Fatalf("expecting %d row less than %d after DeleteRows(%d, %d) but found %d", items, initialRowCount, first, last, rowCount)
	}
	for i := 0; i < table.RowCount(); i++ {
		var item int
		item, err = table.GetInt("item", i)
		if err != nil {
			t.Fatal(err)
		}
		if item == first {
			t.Fatalf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d", first, first, last, first)
		}
		if item == last {
			t.Fatalf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d", last, first, last, last)
		}
	}

	// Test delete 6 items: 4 to 9
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	first = 6
	last = 9
	items = last - first + 1
	err = table.DeleteRows(first, last)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	// fmt.Println(table)
	rowCount = table.RowCount()
	if rowCount != initialRowCount-items {
		t.Fatalf("expecting %d row less than %d after DeleteRows(%d, %d) but found %d", items, initialRowCount, first, last, rowCount)
	}
	for i := 0; i < table.RowCount(); i++ {
		var item int
		item, err = table.GetInt("item", i)
		if err != nil {
			t.Fatal(err)
		}
		if item == first {
			t.Fatalf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d", first, first, last, first)
		}
		if item == last {
			t.Fatalf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d", last, first, last, last)
		}
	}

	// Test delete 3 items: 0 to 9
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	first = 0
	last = 2
	items = last - first + 1
	err = table.DeleteRows(first, last)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	// fmt.Println(table)
	rowCount = table.RowCount()
	if rowCount != initialRowCount-items {
		t.Fatalf("expecting %d row less than %d after DeleteRows(%d, %d) but found %d", items, initialRowCount, first, last, rowCount)
	}
	for i := 0; i < table.RowCount(); i++ {
		item, err := table.GetInt("item", i)
		if err != nil {
			t.Fatal(err)
		}
		if item == first {
			t.Fatalf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d", first, first, last, first)
		}
		if item == last {
			t.Fatalf("expecting to NOT find item %d after DeleteRows(%d, %d) but found %d", last, first, last, last)
		}
	}
}

// Note: Leading lowercase in struct is required for it to be recognised as an Example!
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
	if err != nil {
		log.Println(err)
	}

	if isValid, err := table.IsValidTable(); !isValid {
		log.Println(err)
	}

	err = table.DeleteRows(4, 6)
	if isValid, err := table.IsValidTable(); !isValid {
		log.Println(err)
	}

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

	tableString :=
		`[commands]
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

	tableString :=
		`[commands]
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
	[tbl]
	s string = "Fred"
	t bool = true
	i int = 23
	f float64 = 55.5
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	var expecting string
	var found string

	expecting = "Fred"
	found, err = table.GetValAsString("s", 0)
	if err != nil {
		t.Fatal(err)
	}
	if found != expecting {
		t.Fatalf("expecting %s but found: %s", expecting, found)
	}

	expecting = "true"
	found, err = table.GetValAsString("t", 0)
	if err != nil {
		t.Fatal(err)
	}
	if found != expecting {
		t.Fatalf("expecting %s but found: %s", expecting, found)
	}

	expecting = "23"
	found, err = table.GetValAsString("i", 0)
	if err != nil {
		t.Fatal(err)
	}
	if found != expecting {
		t.Fatalf("expecting %s but found: %s", expecting, found)
	}

	expecting = "55.5"
	found, err = table.GetValAsString("f", 0)
	if err != nil {
		t.Fatal(err)
	}
	if found != expecting {
		t.Fatalf("expecting %s but found: %s", expecting, found)
	}
}

func TestTableSet_FileName(t *testing.T) {
	tableString := `
		[tbl]
		s string = "Fred"
		t bool = true
		i int = 23
		f float64 = 55.5
	`

	// For testing, we need to write this out to a file so we can read it back.
	actualFileName := util.FuncNameNoParens() + ".txt"
	err := ioutil.WriteFile(actualFileName, []byte(tableString), 0644)
	if err != nil {
		t.Fatal(err)
	}

	tableSet, err := NewTableSetFromFile(actualFileName)
	if err != nil {
		t.Fatal(err)
	}

	fileName := tableSet.FileName()
	if fileName != actualFileName {
		t.Fatalf("Expecting FileName() = %q but found %q", actualFileName, fileName)
	}
	if isValid, err := tableSet.IsValidTableSet(); !isValid {
		t.Fatal(err)
	}
}

func TestTableSet_SetName(t *testing.T) {
	expected := ""
	tableSet, err := NewTableSet(expected)
	if err != nil {
		t.Fatal(err)
	}

	tableSetName := tableSet.Name()
	if tableSetName != expected {
		t.Fatalf("Expecting tableSetName = %q but found %q", expected, tableSetName)
	}

	expected = "Musk"
	tableSet.SetName(expected)
	tableSetName = tableSet.Name()
	if tableSetName != expected {
		t.Fatalf("Expecting tableSetName = %q but found %q", expected, tableSetName)
	}
	if isValid, err := tableSet.IsValidTableSet(); !isValid {
		t.Fatal(err)
	}
}

func TestTable_SetName(t *testing.T) {
	expected := "InitialName"
	table, err := NewTable(expected)
	if err != nil {
		t.Fatal(err)
	}

	tableName := table.Name()
	if tableName != expected {
		t.Fatalf("Expecting tableName = %q but found %q", expected, tableName)
	}

	expected = "Elon"
	if err = table.SetName(expected); err != nil {
		t.Fatal(err)
	}
	tableName = table.Name()
	if tableName != expected {
		t.Fatalf("Expecting tableName = %q but found %q", expected, tableName)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
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
		{"float32", true}, // Missing value is math.NaN()
		{"float64", true}, // Missing value is math.NaN()
	}

	for _, test := range tests {
		_, hasMissing := missingValueForType(test.typeName)
		if hasMissing != test.expected {
			t.Fatalf("Expecting missingValueForType(%q) = %t but found %t", test.typeName, test.expected, hasMissing)
		}
	}
}

func TestPreNumberOf(t *testing.T) {
	var tests = []struct {
		sNumber  string
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
			t.Fatalf("Expecting preNumberOf(%q) = %d but found %d", test.sNumber, test.expected, preNumber)
		}
	}
}

func TestPointsOf(t *testing.T) {
	var tests = []struct {
		sNumber  string
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
			t.Fatalf("Expecting pointsOf(%q) = %d but found %d", test.sNumber, test.expected, points)
		}
	}
}

func TestPrecisionOf(t *testing.T) {
	var tests = []struct {
		sNumber  string
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
			t.Fatalf("Expecting precisionOf(%q) = %d but found %d", test.sNumber, test.expected, precision)
		}
	}
}

func TestPadTrailingZeros(t *testing.T) {
	var tests = []struct {
		trailing string
		expected string
	}{
		{"0.0", "0.0"},       // Leave as is.
		{"0.00", "0.0 "},     // Pad with space.
		{"0.0000", "0.0   "}, // Pad with spaces.
		{".0", ".0"},         // Pad.
		{".10", ".1 "},       // Pad.
		{".100", ".1  "},     // Pad.
		{"0", "0"},           // Integer. Don't pad trailing zeros in int-like floats.
		{"10", "10"},         // Integer. Don't pad trailing zeros in int-like floats.
		{"100", "100"},       // Integer. Don't pad trailing zeros in int-like floats.
	}

	for _, test := range tests {

		trimmed := padTrailingZeros(test.trailing)
		if trimmed != test.expected {
			t.Fatalf("Expecting TrimTrailingZeros(%q) = %q but found %q", test.trailing, test.expected, trimmed)
		}
	}
}

func TestTrimTrailingZeros(t *testing.T) {
	var tests = []struct {
		trailing string
		expected string
	}{
		{"0.0", "0.0"},    // Leave as is.
		{"0.00", "0.0"},   // Trim zeros.
		{"0.0000", "0.0"}, // Trim zeros.
		{".0", ".0"},      // Trim.
		{".10", ".1"},     // Trim.
		{".100", ".1"},    // Trim.
		{"0", "0"},        // Integer. Don't pad trailing zeros in int-like floats.
		{"10", "10"},      // Integer. Don't pad trailing zeros in int-like floats.
		{"100", "100"},    // Integer. Don't pad trailing zeros in int-like floats.
	}

	for _, test := range tests {

		trimmed := trimTrailingZeros(test.trailing)
		if trimmed != test.expected {
			t.Fatalf("Expecting TrimTrailingZeros(%q) = %q but found %q", test.trailing, test.expected, trimmed)
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
		t.Fatal(err) // We're not testing this function.
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	var tests = []struct {
		colIndex int
		colType  string
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
			t.Fatalf("Expecting table.IsColTypeByColIndex(%d, %q) = %t but found %t", test.colIndex, test.colType, test.expected, isColType)
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
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
		t.Fatal(err) // We're not testing this function.
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	var tests = []struct {
		colName  string
		colType  string
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
			t.Fatalf("Expecting table.IsColType(%s, %q) = %t but found %t", test.colName, test.colType, test.expected, isColType)
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

/*
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
*/

/*
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
*/

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
	s :=
		`[sable_fur]
	  i s            f b          ff
	int string float32 bool  float64
	  1 "abc"     2.34 true    7.899
	  2 "xyz"     4.5  false   6
	  3 "s  s"    4.9  false   5.5
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
	i   f       u    s
	int float32 uint string
	3   44.55   2    "Here I am!"
	4   22.99   255  "And now I'm not ..."
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

func TestNewTableFromMetadata(t *testing.T) {

	var table *Table
	var err error

	var tests = []struct {
		colNames []string
		colTypes []string
		expected bool
	}{
		{[]string{"Age", "Mothballs", "delims", "tags"}, []string{"int", "bool", "string", "string"}, true},
		{[]string{"Age", "Mothballs", "delims"}, []string{"int", "bool", "string", "string"}, false}, // Missing name
		{[]string{"Age", "Mothballs", "delims", "tags"}, []string{"int", "bool", "string"}, false},   // Missing type
		{[]string{}, []string{"int", "bool", "string", "string"}, false},                             // Empty name slice
		{[]string{"Age", "Mothballs", "delims", "tags"}, []string{}, false},                          // Empty type slice
		{[]string{}, []string{}, true}, // Empty table is allowed
	}

	for _, test := range tests {

		table, err = NewTableFromMetadata("Moviegoers", test.colNames, test.colTypes)
		if (err == nil) != test.expected {
			if err != nil {
				t.Fatal(err)
			} else {
				t.Fatalf("Expecting fail: NewTableFromMetadata(\"Moviegoers\", %v, %v)", test.colNames, test.colTypes)
			}
		}

		_, err = table.IsValidTable()
		if (err == nil) != test.expected {
			t.Fatal(err)
		}

		err = table.AppendRows(1)
		if (err == nil) != test.expected {
			t.Fatal(err)
		}

		if table != nil {
			rowCount := table.RowCount()
			if (rowCount == 1) != test.expected {
				t.Fatal(err)
			}
		}

		_, err = table.IsValidRow(0)
		if (err == nil) != test.expected {
			t.Fatal(err)
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func ExampleTable_SetRowFloatCellsToNaN() {
	s :=
		`[three_rows]
	i    s      f       b     f2
	int  string float32 bool  float64
	0    "abc"  2.3     true  42.0
	1    "xyz"  4.5     false 43.0
	2    "s  s" 4.9     false 44.0
	`
	table, err := NewTableFromString(s)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Before ...")
	fmt.Println(table)

	rowIndex := 1 // The middle row.
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
	s :=
		`[three_rows]
	i    s      f1      b     f2
	int  string float32 bool  float64
	0    "abc"  2.3     true  42.0
	1    "xyz"  4.5     false 43.0
	2    "s  s" 4.9     false 44.0
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
		from     string
		to       string
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
			t.Fatal(err)
		}

		err = table.RenameCol(test.from, test.to)
		if (err == nil) != test.expected {
			t.Fatalf("Expecting table.RenameCol(%q, %q) %s but found err = %v",
				test.from, test.to, ternString(test.expected, "SUCCESS", "FAILURE"), err)
		}
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}
	}
}

func ternString(itIs bool, ifTrue string, ifFalse string) string {
	if itIs {
		return ifTrue
	} else {
		return ifFalse
	}
}

func TestPlural(t *testing.T) {

	var tests = []struct {
		in       int
		expected string
	}{
		{1, ""},
		{-1, ""},
		{-2, "s"},
		{2, "s"},
	}

	for _, test := range tests {

		var result string = plural(test.in)
		if result != test.expected {
			t.Fatalf("Expecting plural(%d) = %q but found: %q", test.in, test.expected, result)
		}
	}
}

// Note: IsValidColValue() is an alias to IsValidCellValue()
func TestIsValidCellValue(t *testing.T) {
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
		t.Fatal(err)
	}

	var tests = []struct {
		col       string
		val       interface{}
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
		{"f32", 23.4, false},         // Floating point constant is float64
		{"f32", float32(23.4), true}, // It's now a float32
		{"s", nil, false},
		{"i", nil, false},
		{"f32", nil, false},
	}

	for _, test := range tests {

		result, err := table.IsValidCellValue(test.col, test.val)
		if result != test.expecting {
			t.Fatalf("Expecting IsValidCellValue(%q, %v) = %t but found: %t, err: %v", test.col, test.val, test.expecting, result, err)
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

// Note: Leading lowercase in planets is required for it to be recognised as an Example!
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
	"Jupiter" 318.000      5.2    79     4 "just"
	"Saturn"   95.000      9.5    82     5 "sat"
	"Uranus"   15.000     19.2    27     6 "upon"
	"Neptune"  17.000     30.6    13     7 "nine"
	"Pluto"     0.002     39.4     5     8 "porcupines"
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	// Simply echo it back out.
	fmt.Println(table)

	//	REINSTATE OPEN AND CLOSE COMMENTS
	//		Notice that by default the columns of data are padded with spaces and numeric types
	//		are right-aligned.
	//		This reflects the opinion that human readability is important.
	//		*Table.String() and *TableSet.String() call their underlying StringPadded() methods.
	//		Where human readability is not important (with messaging or as a wire format) use:
	//		*Table.StringUnpadded()
	//		*TableSet.StringUnpadded()
	//		StringUnpadded() is 3 to 4 times faster.
	//		Reading a padded table string is only slightly slower (about 2.7% slower).
	//	REINSTATE OPEN AND CLOSE COMMENTS

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
	// "Jupiter" 318.0        5.2    79     4 "just"
	// "Saturn"   95.0        9.5    82     5 "sat"
	// "Uranus"   15.0       19.2    27     6 "upon"
	// "Neptune"  17.0       30.6    13     7 "nine"
	// "Pluto"     0.002     39.4     5     8 "porcupines"
	//
	// [planets]
	// name mass distance moons index mnemonic
	// string float64 float64 int int string
	// "Mercury" 0.055 0.4 0 0 "my"
	// "Venus" 0.815 0.7 0 1 "very"
	// "Earth" 1 1 1 2 "elegant"
	// "Mars" 0.107 1.5 2 3 "mother"
	// "Jupiter" 318 5.2 79 4 "just"
	// "Saturn" 95 9.5 82 5 "sat"
	// "Uranus" 15 19.2 27 6 "upon"
	// "Neptune" 17 30.6 13 7 "nine"
	// "Pluto" 0.002 39.4 5 8 "porcupines"
}

// This is not a comprehensive test.
func TestTable_Equals(t *testing.T) {
	var err error
	var table1, table2 *Table

	_, err = table1.Equals(table2)
	if err == nil {
		t.Fatalf("Expecting an error calling Equals() on nil table")
	}
	//	fmt.Println(err)

	t1string :=
		`[MyTable]
    i   s       f       ui
    int string  float64 uint
    1   "abc"   5.50    50
    2   "def"   6.66    60
    `
	table1, err = NewTableFromString(t1string)
	if err != nil {
		t.Fatal(err)
	}

	_, err = table1.Equals(table2)
	if err == nil {
		t.Fatalf("Expecting an error calling Equals() with nil table")
	}
	//	fmt.Println(err)

	t2string :=
		`[MyTable]
    ui      i   s       f
    uint    int string  float64
    50      1   "abc"   5.5
    60      2   "def"   6.6600
    `
	table2, err = NewTableFromString(t2string)
	if err != nil {
		t.Fatal(err)
	}

	equals, err := table1.Equals(table2)
	if !equals {
		t.Fatalf("Expecting table1.Equals(table2) = true but found %t", equals)
	}
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table1.IsValidTable(); !isValid {
		t.Fatal(err)
	}
	if isValid, err := table2.IsValidTable(); !isValid {
		t.Fatal(err)
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

	// These tables share sort keys i and s

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

func ExampleTable_GetTableAsCSV() {

	tableString :=
		`[ForCSV]
	first_name  last_name   username    i   f64     b       f32     commas  quotes    runes end
	string      string      string      int float64 bool    float32 string  string    rune  string
	"Rob"       "Pike"      "rob"       1   1.1     true    NaN     ",end"  "\"xyz\"" '本'  "end"
	"Ken"       "Thompson"  "ken"       3   NaN     true    3.3     "beg,"  "'abc'"   '\''  "end"
	"Robert"    "Griesemer" "gri"       5   5.5     true    NaN     "m,d"   " \"\" "  ' '   "end"
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("gotables table we wish to convert to CSV ...")
	fmt.Println(table)

	var csv string
	csv, err = table.GetTableAsCSV()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("gotables table converted to CSV ...")
	fmt.Println(csv)

	optionalSubstituteHeadingNames := []string{"F Name", "L Name", "", "i", "f64", "bool", "f32", "Commas", "Quotes", "Runes", "end"}
	csv, err = table.GetTableAsCSV(optionalSubstituteHeadingNames...)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("gotables table converted to CSV with user-provided optional heading names ...")
	fmt.Println(csv)

	// Output:
	// gotables table we wish to convert to CSV ...
	// [ForCSV]
	// first_name last_name   username   i     f64 b        f32 commas quotes    runes end
	// string     string      string   int float64 bool float32 string string    rune  string
	// "Rob"      "Pike"      "rob"      1     1.1 true     NaN ",end" "\"xyz\"" '本'   "end"
	// "Ken"      "Thompson"  "ken"      3     NaN true     3.3 "beg," "'abc'"   '\''  "end"
	// "Robert"   "Griesemer" "gri"      5     5.5 true     NaN "m,d"  " \"\" "  ' '   "end"
	//
	// gotables table converted to CSV ...
	// first_name,last_name,username,i,f64,b,f32,commas,quotes,runes,end
	// Rob,Pike,rob,1,1.1,true,,",end","""xyz""",本,end
	// Ken,Thompson,ken,3,,true,3.3,"beg,",'abc',',end
	// Robert,Griesemer,gri,5,5.5,true,,"m,d"," """" "," ",end
	//
	// gotables table converted to CSV with user-provided optional heading names ...
	// F Name,L Name,username,i,f64,bool,f32,Commas,Quotes,Runes,end
	// Rob,Pike,rob,1,1.1,true,,",end","""xyz""",本,end
	// Ken,Thompson,ken,3,,true,3.3,"beg,",'abc',',end
	// Robert,Griesemer,gri,5,5.5,true,,"m,d"," """" "," ",end
}

func TestTable_Copy(t *testing.T) {

	tableStringRows0 :=
		`[Types]
	i int
	b bool
	f64 float64
	f32 float32
	bb []byte
	s string
	`
	tableStringRows1 :=
		`[Table]
	x float64 = 44.4
	b bool = true
	`
	tableStringRows2 :=
		` [Table]
	a	b	c
	int	int	string
	22	23	"Another"
	43	47	"Yet more"
	`

	var tests = []struct {
		tableString string
		rowCount    int
		copyRows    bool
	}{
		{tableStringRows0, 0, false},
		{tableStringRows0, 0, true},
		{tableStringRows1, 1, false},
		{tableStringRows1, 1, true},
		{tableStringRows2, 2, false},
		{tableStringRows2, 2, true},
	}

	for _, test := range tests {

		table, err := NewTableFromString(test.tableString)
		if err != nil {
			t.Fatal(err)
		}
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		tableCopy, err := table.Copy(test.copyRows)
		if err != nil {
			t.Fatalf("table.Copy(%t) with rowCount=%d: %s", test.copyRows, test.rowCount, err)
		}
		if isValid, err := tableCopy.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		if test.copyRows {
			// Expecting same rowCount in each.
			_, err = tableCopy.Equals(table)
			if err != nil {
				t.Fatalf("table.Copy(%t) with rowCount=%d: %s", test.copyRows, test.rowCount, err)
			}
		} else {
			// Expecting zero rowCount in tableCopy.
			// Need to delete rowCount in source table for Equals to compare empty with empty.
			err = table.DeleteRowsAll()
			if err != nil {
				t.Fatalf("table.Copy(%t) with rowCount=%d: %s", test.copyRows, test.rowCount, err)
			}
		}
	}
}

func TestGetColInfoAsSlices(t *testing.T) {

	table, err := NewTableFromString(
		`[TableWithByteSlice]
		uintNums []byte = [0 1 255 3 4 8]
		i int = 42
		b []byte = [1 1 255]
		u []uint8 = [2 2 255 2]
		f float32 = 32
		x []byte = []
		s string = "In Between ..."
		y []uint8 = []
		j int64 = 99
	`)
	if err != nil {
		t.Fatal(err)
	}

	var colNames []string
	var colTypes []string
	colNames, colTypes, err = table.GetColInfoAsSlices()
	if err != nil {
		t.Fatal(err)
	}

	for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
		colName, err := table.ColName(colIndex)
		if err != nil {
			t.Fatal(err)
		}

		colType, err := table.ColTypeByColIndex(colIndex)
		if err != nil {
			t.Fatal(err)
		}

		if colNames[colIndex] != colName {
			t.Fatalf("expecting colName %s at colNames[%d], not %s", colName, colIndex, colNames[colIndex])
		}

		if colTypes[colIndex] != colType {
			t.Fatalf("expecting colType %s at colTypes[%d], not %s", colType, colIndex, colTypes[colIndex])
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func TestTableSet_TableIndex(t *testing.T) {

	tableSet, err := NewTableSetFromString(
		`[Table0]

		[Table1]

		[Table2]

		[Table3]
	`)
	if err != nil {
		t.Fatal(err)
	}

	// Get table index.
	tableName := "Table2"
	tableIndex, err := tableSet.TableIndex(tableName)
	if err != nil {
		t.Fatal(err)
	}

	expecting := 2
	if tableIndex != expecting {
		t.Fatalf("expecting tableIndex %d from tableSet.TableIndex(%q), not %d", expecting, tableName, tableIndex)
	}
	if isValid, err := tableSet.IsValidTableSet(); !isValid {
		t.Fatal(err)
	}
}

func TestTableSet_DeleteTableByTableIndex(t *testing.T) {

	tableSet, err := NewTableSetFromString(
		`[Table0]

		[Table1]

		[Table2]

		[Table3]
	`)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := tableSet.IsValidTableSet(); !isValid {
		t.Fatal(err)
	}

	tableCount := tableSet.TableCount()
	expecting := 4
	if tableCount != expecting {
		t.Fatalf("expecting tableSet.TableCount() = %d, not %d", expecting, tableCount)
	}

	tableName := "Table2"
	_, err = tableSet.Table(tableName)
	if err != nil {
		t.Fatal(err)
	}

	// Delete table.
	tableIndex := 2
	err = tableSet.DeleteTableByTableIndex(tableIndex)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := tableSet.IsValidTableSet(); !isValid {
		t.Fatal(err)
	}

	tableCount = tableSet.TableCount()
	expecting = 3
	if tableCount != expecting {
		t.Fatalf("expecting tableSet.TableCount() = %d, not %d", expecting, tableCount)
	}

	_, err = tableSet.Table(tableName)
	if err == nil {
		t.Fatalf("expecting table [%s] to be deleted from tableSet, but it's still there", tableName)
	}
}

func TestTableSet_DeleteTable(t *testing.T) {

	tableSet, err := NewTableSetFromString(
		`[Table0]

		[Table1]

		[Table2]

		[Table3]
	`)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := tableSet.IsValidTableSet(); !isValid {
		t.Fatal(err)
	}

	tableCount := tableSet.TableCount()
	expecting := 4
	if tableCount != expecting {
		t.Fatalf("expecting tableSet.TableCount() = %d, not %d", expecting, tableCount)
	}

	tableName := "Table2"
	_, err = tableSet.Table(tableName)
	if err != nil {
		t.Fatal(err)
	}

	// Delete table.
	err = tableSet.DeleteTable(tableName)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := tableSet.IsValidTableSet(); !isValid {
		t.Fatal(err)
	}

	tableCount = tableSet.TableCount()
	expecting = 3
	if tableCount != expecting {
		t.Fatalf("expecting tableSet.TableCount() = %d, not %d", expecting, tableCount)
	}

	_, err = tableSet.Table(tableName)
	if err == nil {
		t.Fatalf("expecting table [%s] to be deleted from tableSet, but it's still there", tableName)
	}
}

var allTypesZeroVals string = `
    [AllTypes]
	_bool    bool    = false
	_byte    byte    = 0
	_byte_   []byte  = []
	_float32 float32 = 0.0
	_float64 float64 = 0.0
	_int     int     = 0
	_int16   int16   = 0
	_int32   int32   = 0
	_int64   int64   = 0
	_int8    int8    = 0
	_string  string  = ""
	_uint    uint    = 0
	_uint16  uint16  = 0
	_uint32  uint32  = 0
	_uint64  uint64  = 0
	_uint8   uint8   = 0
	_uint8_  []uint8 = []
    `

func TestTable_AppendRow(t *testing.T) {

	table, err := NewTableFromString(allTypesZeroVals)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	rowCount := table.RowCount()
	expecting := 1
	if rowCount != expecting {
		t.Fatalf("expecting table.RowCount() = %d, not %d", expecting, rowCount)
	}

	// All cells in new row should be zero values.
	err = table.AppendRow()
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
		var rowIndex int = 0
		expecting, err := table.GetValByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		rowIndex = 1
		value, err := table.GetValByColIndex(colIndex, rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		// invalid operation: value.([]uint8) != expecting.([]uint8) (slice can only be compared to nil)
		var colType string = fmt.Sprintf("%T", value)
		switch colType {
		case "[]uint8":
			if equals := bytes.Equal(value.([]uint8), expecting.([]uint8)); !equals {
				t.Fatalf("expecting table.GetValByColIndex(%d, %d) = %v, not %v", colIndex, 1, expecting, value)
			}
		case "[]byte":
			// Note: case "[]byte" seems to be never reached.
			if equals := bytes.Equal(value.([]byte), expecting.([]byte)); !equals {
				t.Fatalf("expecting table.GetValByColIndex(%d, %d) = %v, not %v", colIndex, 1, expecting, value)
			}
		default:
			if value != expecting {
				t.Fatalf("expecting table.GetValByColIndex(%d, %d) = %v, not %v", colIndex, 1, expecting, value)
			}
		}
	}
	//	table.PrintCols()
}

/*
	Surprisingly, Go actually executes the method on a nil object, and allows test of nil in the method.
*/
func TestTableSet_WriteFile_NilTableSet(t *testing.T) {

	var nilTableSet *TableSet

	err := nilTableSet.WriteFile("", 0644)
	if err == nil {
		t.Fatalf("expecting error tableSet is <nil>, but no error was returned")
	}

	// fmt.Println(err)
}

func TestByteSliceEquals(t *testing.T) {

	var tests = []struct {
		index    int // To simplify locating test case in tests.
		succeeds bool
		slice1   []byte
		slice2   []byte
	}{
		{0, true, []byte{1, 2, 3}, []byte{1, 2, 3}},
		{1, false, nil, []byte{1, 2, 3}},
		{2, false, []byte{1, 2, 3}, nil},
		{3, true, nil, nil},
		{4, false, []byte{4, 5, 6, 7}, []byte{4, 5, 6}},
		{5, true, []byte{'A', 'N', 'M', 'O', 'P', 'Q'}, []byte{'A', 'N', 'M', 'O', 'P', 'Q'}},
		{6, false, []byte{'a', 'g', 't', 'e', 'q', 'm'}, []byte{'A', 'n', 'M', 'o', 'p', 'Q'}},
		{7, false, []byte{4, 5, 6, 7}, []byte{4, 5, 6, 8}},
	}

	var equals bool

	for i, test := range tests {
		equals, err := Uint8SliceEquals(test.slice1, test.slice2)
		if equals != test.succeeds {
			t.Fatalf("test[%d]: Uint8SliceEquals(): equals == %t but expecting succeeds == %t", i, equals, test.succeeds)
		}

		// Check index is correct.
		if test.index != i {
			t.Fatalf("test[%d]: test.index == %d but expecting index == %d", i, test.index, i)
		}

		_ = err
		/* Output error messages
		if err != nil {
			fmt.Printf("test[%d]: Uint8SliceEquals(): %v\n", i, err)
		}
		*/
	}

	for i, test := range tests {
		// This is the official Equal function. Confirms validity. It doesn't tell you which element was unequal.
		equals = bytes.Equal(test.slice1, test.slice2)
		if equals != test.succeeds {
			t.Fatalf("test[%d]: bytes.Equal(): equals == %t but expecting succeeds == %t", i, equals, test.succeeds)
		}
	}

	for i, test := range tests {
		// This is the official Compare function. Confirms validity. It doesn't tell you which element was unequal.
		var comparison int = bytes.Compare(test.slice1, test.slice2)
		equals = comparison == 0
		if equals != test.succeeds {
			t.Fatalf("test[%d]: bytes.Compare(): equals == %t but expecting succeeds == %t", i, equals, test.succeeds)
		}
	}
}

func TestIsValidTableName(t *testing.T) {
	var tests = []struct {
		name     string
		validity bool
	}{
		{"Fred", true},
		{"fred", true},
		{"_", true},
		{"_1", true},
		{"1", false},
		{"", false},
	}

	for i, test := range tests {
		validity, err := IsValidTableName(test.name)
		if validity != test.validity {
			if validity {
				t.Fatalf("test[%d]: IsValidTable(%q): validity=%t but expecting validity=%t",
					i, test.name, validity, test.validity)
			} else {
				t.Fatalf("test[%d]: IsValidTable(%q): validity=%t but expecting validity=%t (%v)",
					i, test.name, validity, test.validity, err)
			}
		}
	}
}

func TestRune(t *testing.T) {

	const IGNORE_RUNE = 'M' // For the rune literals that the compiler won't allow.

	// Test using rune literals.
	// Rune literals (and invalid literals) are from: https://golang.org/ref/spec#Rune_literals

	var tests = []struct {
		runeVal   rune   // rune literal.
		stringVal string // string containing a rune literal surrounded by single quotes.
		int32Val  int32  // int32 equivalent of a rune. Calculate in loop.
		validity  bool
		index     int
	}{
		{'a', "'a'", 97, true, 0},                         //  0
		{'\a', "'\a'", 7, true, 1},                        //  1	\a   U+0007 alert or bell
		{'\b', "'\b'", 8, true, 2},                        //  2	\b   U+0008 backspace
		{'\f', "'\f'", 12, true, 3},                       //  3	\f   U+000C form feed
		{'\n', "'\n'", 10, true, 4},                       //  4	\n   U+000A line feed or newline
		{'\r', "'\r'", 13, true, 5},                       //  5	\r   U+000D carriage return
		{'\t', "'\t'", 9, true, 6},                        //  6	\t   U+0009 horizontal tab
		{'\v', "'\v'", 11, true, 7},                       //  7	\v   U+000b vertical tab
		{'\\', "'\\\\'", 92, true, 8},                     //  8	\\   U+005c backslash 92	strconv.UnquoteChar() doesn't like '\\'
		{'\'', `'\''`, 39, true, 9},                       //  9	\'   U+0027 single quote  (valid escape only within rune literals)
		{'a', "'a'", 97, true, 10},                        // 10
		{'ä', "'ä'", 228, true, 11},                       // 11
		{'本', "'本'", 26412, true, 12},                     // 12
		{'\t', "'\t'", 9, true, 13},                       // 13
		{'\000', "'\000'", 0, true, 14},                   // 14
		{'\007', "'\007'", 7, true, 15},                   // 15
		{IGNORE_RUNE, "'\377'", 255, true, 16},            // 16	// This octal is 65533 (255?) the unicode "replacement character". Doesn't parse.
		{'\x07', "'\x07'", 7, true, 17},                   // 17
		{IGNORE_RUNE, "'\xff'", 255, true, 18},            // 18	// This   hex is 65533 (255?) the unicode "replacement character". Doesn't parse.
		{'\u12e4', "'\u12e4'", 4836, true, 19},            // 19
		{'\U00101234', "'\U00101234'", 1053236, true, 20}, // 20
		{'\u2318', "'\u2318'", 8984, true, 21},            // 21
		{'ऎ', "'ऎ'", 2318, true, 22},                      // 22	// Place of Interest Sign apparently interchangable 2318 8984
		{'B', "'B'", 66, true, 23},                        // 23
		{'\u0000', "'\u0000'", 0, true, 24},               // 24	// Zero value of a rune.
		/*		These literals are caught by the Go compiler. Not possible to check them here.
				{ IGNORE_RUNE,	"'aa'", 		-1, false },// 25	// illegal: too many characters
				{ IGNORE_RUNE,	"'\xa'", 		-1, false },// 26	// illegal: too few hexadecimal digits
				{ IGNORE_RUNE,	"'\0'", 		-1, false },// 27	// illegal: too few octal digits
				{ IGNORE_RUNE,	"'\uDFFF'", 	-1, false },// 28	// illegal: surrogate half
				{ IGNORE_RUNE,	"'\U00110000'",	-1, false },// 29	// illegal: invalid Unicode code point
		*/
		{'D', "'\x44'", 68, true, 25},        // 25
		{IGNORE_RUNE, "'\x00'", 0, true, 26}, // 26
	}

	// Note: runeRegexpString is defined in parser.go
	var runeRegexp *regexp.Regexp = regexp.MustCompile(runeRegexpString)

	for i, test := range tests {
		if tests[i].int32Val == -1 { // Skip the zero value rune.
			tests[i].int32Val = test.runeVal // Initialise field int32Val to the correct rune value.
		}
		// fmt.Printf("'%c' = %d\n", test.runeVal, test.runeVal)
	}

	for i, test := range tests {

		// where(fmt.Sprintf("*** test[%d]", i))

		if test.index != i {
			t.Fatalf("test[%d]: index %d should be %d not %d", i, test.index, i, test.index)
		}

		// Basic sanity test of literals used in the tests.
		if test.runeVal != IGNORE_RUNE {
			validity := utf8.ValidRune(test.runeVal)
			if validity != test.validity {
				t.Fatalf("test[%d]: invalid rune: '%c'", i, test.runeVal)
			}
		}

		// Test the regular expression.
		// The match is of the full string which includes single quote delimiters. Not trimmed.
		matched := runeRegexp.MatchString(test.stringVal)
		// where(fmt.Sprintf("test[%d]: matched = %t on %v", i, matched, test.stringVal))
		if matched != test.validity {
			if test.validity {
				// We will not expect runeRegexp to reject invalid rune literals. Leave that to DecodeRuneInString()
				t.Fatalf("test[%d/%d]: runeRegexp match %s failed on string: %s", i, len(tests)-1, runeRegexp, test.stringVal)
			}
		} else {
			//			fmt.Printf("test[%d]: runeRegexp SUCCEEDED on string: %s\n", i, test.stringVal)
		}

		// Now that we have successfully parsed a rune with runeRegexp, see if it's a valid rune.

		// where(fmt.Sprintf("test[%d] parseRune(%q)", i, test.stringVal))
		// Trim off first ' quote.
		// where(fmt.Sprintf("BEFORE trim:%s", test.stringVal))
		test.stringVal = trimDelims(test.stringVal, "'")
		// where(fmt.Sprintf("AFTER  trim:%s", test.stringVal))
		rune2, err := parseRune(test.stringVal)
		if err != nil {
			t.Fatalf("test[%d]: %v", i, err)
		}

		validRune := utf8.ValidRune(rune2)
		if !validRune {
			t.Fatalf("test[%d]: invalid rune: '%c'", i, test.runeVal)
		}

		if test.runeVal != IGNORE_RUNE {
			if rune2 != test.int32Val {
				/*
					var rs []rune = make([]rune, 4)
					for j := 0; j < len(test.stringVal); j++ {
						rs[j] = rune(test.stringVal[j])
					}
				*/
				size := utf8.RuneLen(rune2)
				t.Fatalf("test[%d/%d]: expecting rune %q %d but got %q %d size=%d Decode failed on string %q len=%d",
					i, len(tests)-1, test.int32Val, test.int32Val, rune2, rune2, size, test.stringVal, len(test.stringVal))
			}
		}
		// fmt.Println()
	}
}

func runeToString(r rune) string {
	var s string = fmt.Sprintf("%c", r)
	//	fmt.Printf("%s\n", s)
	return s
}

func printBytes(b []byte) (s string) {
	for i := 0; i < len(b); i++ {
		s += fmt.Sprintf("%c ", i)
	}
	return
}

func TestRuneTable(t *testing.T) {

	var runes string = `[Runes]
	code     glyph  dec str
	rune     rune   int string
	'\u0000' '\x00'   0 ""
	'\x61'   'a'     97 "a"
	'\u0061' 'a'     97 "a"
	'\u0007' '\a'     7 "\a"
	'\u0008' '\b'     8 "\b"
	'\u0009' '\t'     9 "\t"
	'\u000A' '\n'    10 "\n"
	'\u000B' '\v'    11 "\v"
	'\u000C' '\f'    12 "\f"
	'\u000D' '\r'    13 "\r"
	'\u005C' '\\'    92 "\\"
	'\u4E16' '世' 19990 "世"
	'\u754C' '界' 30028 "界"
	'\u0041' 'A'     65 "A"
	'\u0042' 'B'     66 "A"
	'\u0043' 'C'     67 "A"
	'\u0044' 'D'     68 "D"
	'\x6D'   'm'    109 "m"
	'z'      'z'    122 "z"
	`

	table, err := NewTableFromString(runes)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < table.RowCount(); i++ {
		var code rune
		code, err = table.GetRune("code", i)
		if err != nil {
			t.Fatal(err)
		}

		var glyph rune
		glyph, err = table.GetRune("glyph", i)
		if err != nil {
			t.Fatal(err)
		}

		if glyph != code {
			t.Fatalf("table [%s] row %d: code %d != glyph %c (U+%04X) %d", table.Name(), i, code, glyph, glyph, glyph)
		}

		var dec int
		dec, err = table.GetInt("dec", i)
		if err != nil {
			t.Fatal(err)
		}

		if int(code) != dec {
			t.Fatalf("code %d != dec %d", int32(code), dec)
		}
	}

	// fmt.Printf("%v", table)
}

func TestRuneStruct(t *testing.T) {

	var runes string = `[Runes]
	c1     rune = 'a'
	numval int  = 97
	c2     rune = '\x22'
    `

	_, err := NewTableFromString(runes)
	if err != nil {
		t.Fatal(err)
	}

	//	fmt.Printf("%v", table)
}

// This is a time-consuming test. We may want to skip it normally.
func TestManyUnicodes(t *testing.T) {

	const skip bool = false

	var table *Table
	var err error

	if skip {
		// where(fmt.Sprintf("skip = %t", skip))
		// This is a token tiny test with speed the priority.
		var tableString string = `[Runes]
		code     glyph decimal
		rune     rune  int32
		'\u0000' '0' 0
		'\u0001' '0' 1
		'\u0002' '0' 2
		'\u0003' '0' 3
		'\u0004' '0' 4
		'\u0005' '0' 5
		'\u0006' '0' 6
		'\u0007' '\a' 7
		'\u0008' '\b' 8
		'\u0009' '\t' 9
		'\u000A' '\n' 10
		'\u000B' '\v' 11
		'\u000C' '\f' 12
		'\u000D' '\r' 13
		'\u000E' '0' 14
		'\u000F' '0' 15
		'\u0010' '0' 16
		`
		table, err = NewTableFromString(tableString)
		if err != nil {
			t.Fatal(err)
		}
	} else {
		// where(fmt.Sprintf("skip = %t", skip))
		table, err = NewTableFromFile("rune_test.got")
		if err != nil {
			t.Fatal(err)
		}
	}

	/*
		err = table.SetSortKeys("decimal")
		if err != nil { t.Fatal(err) }

		err = table.Sort()
		if err != nil { t.Fatal(err) }
	*/

	for i := 0; i < table.RowCount(); i++ {
		var glyph rune
		glyph, err = table.GetRune("glyph", i)
		if err != nil {
			t.Fatal(err)
		}

		var code rune
		code, err = table.GetRune("code", i)
		if err != nil {
			t.Fatal(err)
		}

		var decimal int32
		decimal, err = table.GetInt32("decimal", i)
		if err != nil {
			t.Fatal(err)
		}

		var specialChars = "\a\b\f\n\r\t\v"
		var isSpecialChar bool = strings.Contains(specialChars, string(glyph))
		// where(fmt.Sprintf("%c isSpecialChar? = %t", glyph, isSpecialChar))
		if (code >= 32 && code < 127) || code > 159 || isSpecialChar {
			// Printable characters: glyphs are set to themselves (and not '0').
			// where(fmt.Sprintf("row[%d]: decimal = %d", i, decimal))
			if glyph == '0' && decimal != 48 {
				t.Fatalf("row[%d]: glyph %q int32(glyph) %d == '0'", i, glyph, int32(glyph))
			}

			if int32(code) != int32(glyph) {
				t.Fatalf("row[%d]: glyph %q int32(glyph) %d != code %d U+0%X", i, glyph, glyph, int32(code), code)
			}

			// Allows for decimal being 0 where decimal wasn't present in the Wikipedia entry
			if int32(code) != decimal && decimal != 0 {
				t.Fatalf("row[%d]: int32(code) %d != decimal %d", i, int32(code), decimal)
			}

			// Allows for decimal being 0 where decimal wasn't present in the Wikipedia entry
			if int32(glyph) != decimal && decimal != 0 {
				t.Fatalf("row[%d]: glyph %q int32(glyph) %d != decimal %d", i, glyph, int32(glyph), decimal)
			}
		} else {
			// where(fmt.Sprintf("row[%d]: decimal = %d", i, decimal))
			// Non-printing characters. Should be set to '0'.
			if glyph != '0' {
				t.Fatalf("row[%d]: glyph %q int32(glyph) %d != 0", i, glyph, int32(glyph))
			}
		}
	}

	// fmt.Printf("%v", table)
}

func TestSomeUnicodes(t *testing.T) {
	tableString :=
		`[Literals1]
	code     glyph dec s
	rune     rune  int string
	'\u0000' '\x00'  0 ""
	'\u0061' 'a'    97 "a"
	'\u0007' '\a'    7 "\a"
	'\u0008' '\b'    8 "\b"
	'\u0009' '\t'    9 "\t"
	'\u000A' '\n'   10 "\n"
	'\u000B' '\v'   11 "\v"
	'\u000C' '\f'   12 "\f"
	'\u000D' '\r'   13 "\r"
	'\u000D' '\\'   13 "\\"
	'\u000D' '世'   13 "世"
	'\u000D' '界'   13 "界"
	`

	// where(tableString)
	_, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	// if err == nil { fmt.Printf("\n%v\n", table) }
}

// Note: Leading lowercase in unicodeRuneLiterals is required for it to be recognised as an Example!
// Note: Glyph width is (in my experience) difficult to manage. Hence the uneven columns. May revisit.
func ExampleNewTableFromString_unicodeRuneLiterals() {

	var runesEqual string = `
[RunesEqual]
code     glyph  dec str
rune     rune   int string
'\u0000' '\x00'   0 ""
'\x61'   'a'     97 "a"
'\u0061' 'a'     97 "a"
'\u0061' 'a'     97 "a"
'\u0061' 'a'     97 "a"
'\u0007' '\a'     7 "\a"
'\u0008' '\b'     8 "\b"
'\u0009' '\t'     9 "\t"
'\u000A' '\n'    10 "\n"
'\u000B' '\v'    11 "\v"
'\u000C' '\f'    12 "\f"
'\u000D' '\r'    13 "\r"
'\u005C' '\\'    92 "\\"
'\u4E16' '世' 19990 "世"
'\u754C' '界' 30028 "界"
'\u0041' 'A'     65 "A"
'\u0042' 'B'     66 "A"
'\u0043' 'C'     67 "A"
'\x44'   'D'     68 "D"
'\u006D' 'm'    109 "m"
'z'      'z'    122 "z"
`

	table, err := NewTableFromString(runesEqual)
	if err != nil {
		log.Println(err)
	}

	fmt.Print("(1) Runes source table:")
	fmt.Printf("%s\n", runesEqual)

	fmt.Println("(2) Runes output table:")
	fmt.Printf("%s", table)

	// Output:
	// (1) Runes source table:
	// [RunesEqual]
	// code     glyph  dec str
	// rune     rune   int string
	// '\u0000' '\x00'   0 ""
	// '\x61'   'a'     97 "a"
	// '\u0061' 'a'     97 "a"
	// '\u0061' 'a'     97 "a"
	// '\u0061' 'a'     97 "a"
	// '\u0007' '\a'     7 "\a"
	// '\u0008' '\b'     8 "\b"
	// '\u0009' '\t'     9 "\t"
	// '\u000A' '\n'    10 "\n"
	// '\u000B' '\v'    11 "\v"
	// '\u000C' '\f'    12 "\f"
	// '\u000D' '\r'    13 "\r"
	// '\u005C' '\\'    92 "\\"
	// '\u4E16' '世' 19990 "世"
	// '\u754C' '界' 30028 "界"
	// '\u0041' 'A'     65 "A"
	// '\u0042' 'B'     66 "A"
	// '\u0043' 'C'     67 "A"
	// '\x44'   'D'     68 "D"
	// '\u006D' 'm'    109 "m"
	// 'z'      'z'    122 "z"
	//
	// (2) Runes output table:
	// [RunesEqual]
	// code   glyph    dec str
	// rune   rune     int string
	// '\x00' '\x00'     0 ""
	// 'a'    'a'       97 "a"
	// 'a'    'a'       97 "a"
	// 'a'    'a'       97 "a"
	// 'a'    'a'       97 "a"
	// '\a'   '\a'       7 "\a"
	// '\b'   '\b'       8 "\b"
	// '\t'   '\t'       9 "\t"
	// '\n'   '\n'      10 "\n"
	// '\v'   '\v'      11 "\v"
	// '\f'   '\f'      12 "\f"
	// '\r'   '\r'      13 "\r"
	// '\\'   '\\'      92 "\\"
	// '世'    '世'    19990 "世"
	// '界'    '界'    30028 "界"
	// 'A'    'A'       65 "A"
	// 'B'    'B'       66 "A"
	// 'C'    'C'       67 "A"
	// 'D'    'D'       68 "D"
	// 'm'    'm'      109 "m"
	// 'z'    'z'      122 "z"
}

func printStringBytes(s string) {
	fmt.Printf("s = %s\n", s)
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("s bytes[%d] = %v\n", i, bytes[i])
	}
}

func printRuneBytes(r rune) {
	var s string = string(r)
	fmt.Printf("r = %c", r)
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("r bytes[%d] = %v\n", i, bytes[i])
	}
}

// Note: Leading lowercase in unicodeRuneLiterals is required for it to be recognised as an Example!
func ExampleNewTableFromString_unicodeRuneLiteralsUnpadded() {

	var runesEqual string = `
[RunesEqual]
code     glyph  dec str
rune     rune   int string
'\u0000' '\x00'   0 ""
'\x61'   'a'     97 "a"
'\u0061' 'a'     97 "a"
'\u0007' '\a'     7 "\a"
'\u0008' '\b'     8 "\b"
'\u0009' '\t'     9 "\t"
'\u000A' '\n'    10 "\n"
'\u000B' '\v'    11 "\v"
'\u000C' '\f'    12 "\f"
'\u000D' '\r'    13 "\r"
'\u005C' '\\'    92 "\\"
'\u4E16' '世' 19990 "世"
'\u754C' '界' 30028 "界"
'\u0041' 'A'     65 "A"
'\u0042' 'B'     66 "A"
'\u0043' 'C'     67 "A"
'\x44'   'D'     68 "D"
'\u006D' 'm'    109 "m"
'z'      'z'    122 "z"
`

	table, err := NewTableFromString(runesEqual)
	if err != nil {
		log.Println(err)
	}

	fmt.Print("(1) Runes source table:")
	fmt.Printf("%s\n", runesEqual)

	fmt.Println("(2) Runes output table:")
	fmt.Printf("%s", table.StringUnpadded())

	// Output:
	// (1) Runes source table:
	// [RunesEqual]
	// code     glyph  dec str
	// rune     rune   int string
	// '\u0000' '\x00'   0 ""
	// '\x61'   'a'     97 "a"
	// '\u0061' 'a'     97 "a"
	// '\u0007' '\a'     7 "\a"
	// '\u0008' '\b'     8 "\b"
	// '\u0009' '\t'     9 "\t"
	// '\u000A' '\n'    10 "\n"
	// '\u000B' '\v'    11 "\v"
	// '\u000C' '\f'    12 "\f"
	// '\u000D' '\r'    13 "\r"
	// '\u005C' '\\'    92 "\\"
	// '\u4E16' '世' 19990 "世"
	// '\u754C' '界' 30028 "界"
	// '\u0041' 'A'     65 "A"
	// '\u0042' 'B'     66 "A"
	// '\u0043' 'C'     67 "A"
	// '\x44'   'D'     68 "D"
	// '\u006D' 'm'    109 "m"
	// 'z'      'z'    122 "z"
	//
	// (2) Runes output table:
	// [RunesEqual]
	// code glyph dec str
	// rune rune int string
	// '\x00' '\x00' 0 ""
	// 'a' 'a' 97 "a"
	// 'a' 'a' 97 "a"
	// '\a' '\a' 7 "\a"
	// '\b' '\b' 8 "\b"
	// '\t' '\t' 9 "\t"
	// '\n' '\n' 10 "\n"
	// '\v' '\v' 11 "\v"
	// '\f' '\f' 12 "\f"
	// '\r' '\r' 13 "\r"
	// '\\' '\\' 92 "\\"
	// '世' '世' 19990 "世"
	// '界' '界' 30028 "界"
	// 'A' 'A' 65 "A"
	// 'B' 'B' 66 "A"
	// 'C' 'C' 67 "A"
	// 'D' 'D' 68 "D"
	// 'm' 'm' 109 "m"
	// 'z' 'z' 122 "z"
}

func BenchmarkSetVal(b *testing.B) {
	var err error

	var planetsString = `[planets]
	name         mass distance moons index mnemonic
	string    float64  float64   int   int string
	"Mercury"   0.055      0.4     0     0 "my"
	"Venus"     0.815      0.7     0     1 "very"
	"Earth"     1.0        1.0     1     2 "elegant"
	"Mars"      0.107      1.5     2     3 "mother"
	"Jupiter" 318.0        5.2    79     4 "just"
	"Saturn"   95.0        9.5    82     5 "sat"
	"Uranus"   15.0       19.2    27     6 "upon"
	"Neptune"  17.0       30.6    13     7 "nine"
	"Pluto"     0.002     39.4     5     8 "porcupines"
	`

	table, err := NewTableFromString(planetsString)
	if err != nil {
		b.Error(err)
	}

	type planets struct {
		name     string
		mass     float64
		distance float64
		moons    int
		index    int
		mnemonic string
	}

	var tests []planets = make([]planets, table.RowCount())

	for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
		tests[rowIndex].name, err = table.GetString("name", rowIndex)
		if err != nil {
			b.Error(err)
		}

		tests[rowIndex].mass, err = table.GetFloat64("mass", rowIndex)
		if err != nil {
			b.Error(err)
		}

		tests[rowIndex].distance, err = table.GetFloat64("distance", rowIndex)
		if err != nil {
			b.Error(err)
		}

		tests[rowIndex].moons, err = table.GetInt("moons", rowIndex)
		if err != nil {
			b.Error(err)
		}

		tests[rowIndex].index, err = table.GetInt("index", rowIndex)
		if err != nil {
			b.Error(err)
		}

		tests[rowIndex].mnemonic, err = table.GetString("mnemonic", rowIndex)
		if err != nil {
			b.Error(err)
		}
	}

	for i := 0; i < b.N; i++ {
		for rowIndex, test := range tests {
			err = table.SetVal("name", rowIndex, test.name)
			if err != nil {
				b.Error(err)
			}

			err = table.SetVal("mass", rowIndex, test.mass)
			if err != nil {
				b.Error(err)
			}

			err = table.SetVal("distance", rowIndex, test.distance)
			if err != nil {
				b.Error(err)
			}

			err = table.SetVal("moons", rowIndex, test.moons)
			if err != nil {
				b.Error(err)
			}

			err = table.SetVal("index", rowIndex, test.index)
			if err != nil {
				b.Error(err)
			}

			err = table.SetVal("mnemonic", rowIndex, test.mnemonic)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkSetValByColIndex(b *testing.B) {
	var err error

	var planetsString = `[planets]
	name         mass distance moons index mnemonic
	string    float64  float64   int   int string
	"Mercury"   0.055      0.4     0     0 "my"
	"Venus"     0.815      0.7     0     1 "very"
	"Earth"     1.0        1.0     1     2 "elegant"
	"Mars"      0.107      1.5     2     3 "mother"
	"Jupiter" 318.0        5.2    79     4 "just"
	"Saturn"   95.0        9.5    82     5 "sat"
	"Uranus"   15.0       19.2    27     6 "upon"
	"Neptune"  17.0       30.6    13     7 "nine"
	"Pluto"     0.002     39.4     5     8 "porcupines"
	`

	fromTable, err := NewTableFromString(planetsString)
	if err != nil {
		b.Error(err)
	}

	type planet struct {
		name     string
		mass     float64
		distance float64
		moons    int
		index    int
		mnemonic string
	}
	var planets []planet = make([]planet, fromTable.RowCount())
	for rowIndex := 0; rowIndex < fromTable.RowCount(); rowIndex++ {
		planets[rowIndex].name, err = fromTable.GetString("name", rowIndex)
		if err != nil {
			b.Error(err)
		}

		planets[rowIndex].mass, err = fromTable.GetFloat64("mass", rowIndex)
		if err != nil {
			b.Error(err)
		}

		planets[rowIndex].distance, err = fromTable.GetFloat64("distance", rowIndex)
		if err != nil {
			b.Error(err)
		}

		planets[rowIndex].moons, err = fromTable.GetInt("moons", rowIndex)
		if err != nil {
			b.Error(err)
		}

		planets[rowIndex].index, err = fromTable.GetInt("index", rowIndex)
		if err != nil {
			b.Error(err)
		}

		planets[rowIndex].mnemonic, err = fromTable.GetString("mnemonic", rowIndex)
		if err != nil {
			b.Error(err)
		}
	}

	var toTable *Table
	toTable, err = fromTable.Copy(true)
	if err != nil {
		b.Error(err)
	}
	for rowIndex := 0; rowIndex < toTable.RowCount(); rowIndex++ {
		err = toTable.SetRowCellsToZeroValue(rowIndex)
		if err != nil {
			b.Error(err)
		}
	}
	for i := 0; i < b.N; i++ {
		for rowIndex, planet := range planets {
			err = toTable.SetValByColIndex(0, rowIndex, planet.name)
			if err != nil {
				b.Error(err)
			}

			err = toTable.SetValByColIndex(1, rowIndex, planet.mass)
			if err != nil {
				b.Error(err)
			}

			err = toTable.SetValByColIndex(2, rowIndex, planet.distance)
			if err != nil {
				b.Error(err)
			}

			err = toTable.SetValByColIndex(3, rowIndex, planet.moons)
			if err != nil {
				b.Error(err)
			}

			err = toTable.SetValByColIndex(4, rowIndex, planet.index)
			if err != nil {
				b.Error(err)
			}

			err = toTable.SetValByColIndex(5, rowIndex, planet.mnemonic)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

func TestSetVal(t *testing.T) {
	// Note: SetVal() calls SetValByColIndex() and so mostly tests it.
	var err error
	var table *Table
	var tableString string = `
	[sable_fur]
    i   s      f       t     b    ui    bb            uu8
    int string float64 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     false 33   2     [33 34 35 36] [37 38 39]
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	// Test data type / col type mismatch.
	err = table.SetVal("i", 0, 23.4)
	if err == nil {
		t.Fatal("SetVal() Expecting type error")
	}

	// Test data type / col type mismatch.
	err = table.SetVal("b", 0, []uint8{22})
	if err == nil {
		t.Fatal("SetVal() Expecting type error")
	}

	// Test byte <> uint8 alias in both directions.
	err = table.SetVal("b", 0, uint8(2))
	if err != nil {
		t.Fatal("SetVal() Expecting type error")
	}
	err = table.SetVal("ui", 0, byte(2))
	if err != nil {
		t.Fatal("SetVal() Expecting type error")
	}

	// Test []byte <> []uint8 alias in both directions.
	err = table.SetVal("bb", 0, []uint8{2})
	if err != nil {
		t.Fatal("SetVal() Expecting type error")
	}
	err = table.SetVal("uu8", 0, []byte{2})
	if err != nil {
		t.Fatal("SetVal() Expecting type error")
	}

	// Test col missing.
	err = table.SetVal("MISSING_COL", 0, 23.4)
	if err == nil {
		t.Fatal("SetVal() Expecting col does not exist error")
	}

	// Test row missing.
	err = table.SetVal("t", 3, false)
	if err == nil {
		t.Fatal("SetVal() Expecting row index out of range error")
	}
	err = table.SetVal("f", -1, 3.3)
	if err == nil {
		t.Fatal("SetVal() Expecting row index out of range error")
	}

	// Test col missing.
	err = table.SetValByColIndex(8, 2, false)
	if err == nil {
		t.Fatal("SetValByColIndex() Expecting col index does not exist error")
	}
	err = table.SetValByColIndex(-1, 1, 3.3)
	if err == nil {
		t.Fatal("SetValByColIndex() Expecting col index does not exist error")
	}
}

func TestGetVal(t *testing.T) {
	// Note: GetVal() calls GetValByColIndex() and so mostly tests it.
	var err error
	var table *Table
	var tableString string = `
	[sable_fur]
    i   s      f       t     b    ui    bb            uu8
    int string float64 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     false 33   2     [33 34 35 36] [37 38 39]
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	var val interface{}

	// Test get of valid value.
	val, err = table.GetVal("i", 2)
	if err != nil {
		t.Fatal(err)
	}
	if val.(int) != 3 {
		t.Fatalf("GetVal() expecting 3, not %v", val)
	}

	// Note: data type mismatch is impossible. It gets what it gets.

	// Test col missing.
	_, err = table.GetVal("MISSING_COL", 0)
	if err == nil {
		t.Fatal("GetVal() Expecting col does not exist error")
	}

	// Test row missing.
	_, err = table.GetVal("t", 3)
	if err == nil {
		t.Fatal("GetVal() Expecting row index out of range error")
	}
	_, err = table.GetVal("f", -1)
	if err == nil {
		t.Fatal("GetVal() Expecting row index out of range error")
	}

	// Test col missing.
	_, err = table.GetValByColIndex(8, 2)
	if err == nil {
		t.Fatal("GetValByColIndex() Expecting col index does not exist error")
	}
	_, err = table.GetValByColIndex(-1, 1)
	if err == nil {
		t.Fatal("GetValByColIndex() Expecting col index does not exist error")
	}
}

func BenchmarkGetVal(b *testing.B) {
	var err error

	var planetsString = `[planets]
	name         mass distance moons index mnemonic
	string    float64  float64   int   int string
	"Mercury"   0.055      0.4     0     0 "my"
	"Venus"     0.815      0.7     0     1 "very"
	"Earth"     1.0        1.0     1     2 "elegant"
	"Mars"      0.107      1.5     2     3 "mother"
	"Jupiter" 318.0        5.2    79     4 "just"
	"Saturn"   95.0        9.5    82     5 "sat"
	"Uranus"   15.0       19.2    27     6 "upon"
	"Neptune"  17.0       30.6    13     7 "nine"
	"Pluto"     0.002     39.4     5     8 "porcupines"
	`

	table, err := NewTableFromString(planetsString)
	if err != nil {
		b.Error(err)
	}

	var val interface{}

	for i := 0; i < b.N; i++ {
		for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
			val, err = table.GetVal("name", rowIndex)
			_ = val
			if err != nil {
				b.Error(err)
			}

			val, err = table.GetVal("mass", rowIndex)
			_ = val
			if err != nil {
				b.Error(err)
			}

			val, err = table.GetVal("distance", rowIndex)
			_ = val
			if err != nil {
				b.Error(err)
			}

			val, err = table.GetVal("moons", rowIndex)
			_ = val
			if err != nil {
				b.Error(err)
			}

			val, err = table.GetVal("index", rowIndex)
			_ = val
			if err != nil {
				b.Error(err)
			}

			val, err = table.GetVal("mnemonic", rowIndex)
			_ = val
			if err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkGetValByColIndex(b *testing.B) {
	var err error

	var planetsString = `[planets]
	name         mass distance moons index mnemonic
	string    float64  float64   int   int string
	"Mercury"   0.055      0.4     0     0 "my"
	"Venus"     0.815      0.7     0     1 "very"
	"Earth"     1.0        1.0     1     2 "elegant"
	"Mars"      0.107      1.5     2     3 "mother"
	"Jupiter" 318.0        5.2    79     4 "just"
	"Saturn"   95.0        9.5    82     5 "sat"
	"Uranus"   15.0       19.2    27     6 "upon"
	"Neptune"  17.0       30.6    13     7 "nine"
	"Pluto"     0.002     39.4     5     8 "porcupines"
	`

	table, err := NewTableFromString(planetsString)
	if err != nil {
		b.Error(err)
	}

	var val interface{}

	for i := 0; i < b.N; i++ {
		for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
			for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
				val, err = table.GetValByColIndex(colIndex, rowIndex)
				_ = val
				if err != nil {
					b.Error(err)
				}
			}
		}
	}
}

func TestGetBoolVal(t *testing.T) {
	var err error
	var table *Table
	var tableString string = `
	[sable_fur]
    i   s      f       t     b    ui    bb            uu8
    int string float64 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     false 33   2     [33 34 35 36] [37 38 39]
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	colNames, _, err := table.GetColInfoAsSlices()
	if err != nil {
		t.Fatal(err)
	}

	for _, colName := range colNames {
		for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
			_, err := table.GetBool(colName, rowIndex)
			if colName == "t" {
				if err != nil {
					t.Fatal(err)
				}
			} else {
				if err == nil {
					t.Fatalf("Expecting GetBool(%q, %d) to fail on non-bool col %q", colName, rowIndex, colName)
				}
			}
		}
	}
}

func TestHasRow(t *testing.T) {
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
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	var tests = []struct {
		row       int
		expecting bool
	}{
		{-1, false},
		{0, true},
		{15, true},
		{16, false},
	}

	for _, test := range tests {
		hasRow, _ := table.HasRow(test.row)
		if hasRow != test.expecting {
			t.Fatalf("expecting table.HasRow(%d) = %v, not %v", test.row, test.expecting, hasRow)
		}
	}
}

func ExampleTable_NewTableReorderColsByColIndex() {
	var err error
	var table *Table
	var tableString string = `
	[TypesGalore01]
    i   s      f       t     b    ui    bb            uu8
    int string float64 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     false 33   2     [33 34 35 36] [37 38 39]
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	// This numeric sequence reverses the column order.
	reorderedTable, err := table.NewTableReorderColsByColIndex(7, 6, 5, 4, 3, 2, 1, 0)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(reorderedTable)

	// Output:
	// [TypesGalore01]
	//   i s            f t        b    ui bb            uu8
	// int string float64 bool  byte uint8 []byte        []uint8
	//   1 "abc"      2.3 true    11     0 [11 12 13 14] [15 16 17]
	//   2 "xyz"      4.5 false   22     1 [22 23 24 25] [26 27 28]
	//   3 "ssss"     4.9 false   33     2 [33 34 35 36] [37 38 39]
	//
	// [TypesGalore01]
	// uu8        bb               ui    b t           f s        i
	// []uint8    []byte        uint8 byte bool  float64 string int
	// [15 16 17] [11 12 13 14]     0   11 true      2.3 "abc"    1
	// [26 27 28] [22 23 24 25]     1   22 false     4.5 "xyz"    2
	// [37 38 39] [33 34 35 36]     2   33 false     4.9 "ssss"   3
}

func ExampleTable_ReorderColsByColIndex() {
	var err error
	var table *Table
	var tableString string = `
	[TypesGalore02]
    i   s      f       t     b    ui    bb            uu8
    int string float64 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     false 33   2     [33 34 35 36] [37 38 39]
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	// This method reorders the table cols in-place, and does not return a new table.

	// This numeric sequence reorders col names to alphabetic order.
	err = table.ReorderColsByColIndex(4, 6, 2, 0, 1, 3, 5, 7)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	// Output:
	// [TypesGalore02]
	//    b bb                  f   i s      t        ui uu8
	// byte []byte        float64 int string bool  uint8 []uint8
	//   11 [11 12 13 14]     2.3   1 "abc"  true      0 [15 16 17]
	//   22 [22 23 24 25]     4.5   2 "xyz"  false     1 [26 27 28]
	//   33 [33 34 35 36]     4.9   3 "ssss" false     2 [37 38 39]
}

func ExampleTable_NewTableReorderCols() {
	var err error
	var table *Table
	var tableString string = `
	[TypesGalore03]
    i   s      f       t     b    ui    bb            uu8
    int string float64 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     false 33   2     [33 34 35 36] [37 38 39]
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	// This method reorders the table cols in-place, and does not return a new table.

	// Let's reorder the cols in alphabetic order.
	colsOrder := []string{"i", "s", "f", "t", "b", "ui", "bb", "uu8"}
	sort.Strings(colsOrder)
	reorderedTable, err := table.NewTableReorderCols(colsOrder...)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(reorderedTable)

	// Output:
	// [TypesGalore03]
	//   i s            f t        b    ui bb            uu8
	// int string float64 bool  byte uint8 []byte        []uint8
	//   1 "abc"      2.3 true    11     0 [11 12 13 14] [15 16 17]
	//   2 "xyz"      4.5 false   22     1 [22 23 24 25] [26 27 28]
	//   3 "ssss"     4.9 false   33     2 [33 34 35 36] [37 38 39]
	//
	// [TypesGalore03]
	//    b bb                  f   i s      t        ui uu8
	// byte []byte        float64 int string bool  uint8 []uint8
	//   11 [11 12 13 14]     2.3   1 "abc"  true      0 [15 16 17]
	//   22 [22 23 24 25]     4.5   2 "xyz"  false     1 [26 27 28]
	//   33 [33 34 35 36]     4.9   3 "ssss" false     2 [37 38 39]
}

func ExampleTable_ReorderCols() {
	var err error
	var table *Table
	var tableString string = `
	[TypesGalore04]
    i   s      f       t     b    ui    bb            uu8
    int string float64 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     false 33   2     [33 34 35 36] [37 38 39]
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	// Let's reorder the cols in alphabetic order.
	colsOrder := []string{"i", "s", "f", "t", "b", "ui", "bb", "uu8"}
	sort.Strings(colsOrder)
	// Let's reorder the cols in reverse order. Which here means reverse alphabetic order.
	// Reversing algorithm from https://github.com/golang/go/wiki/SliceTricks
	for left, right := 0, len(colsOrder)-1; left < right; left, right = left+1, right-1 {
		colsOrder[left], colsOrder[right] = colsOrder[right], colsOrder[left]
	}
	err = table.ReorderCols(colsOrder...)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	// Output:
	// [TypesGalore04]
	//   i s            f t        b    ui bb            uu8
	// int string float64 bool  byte uint8 []byte        []uint8
	//   1 "abc"      2.3 true    11     0 [11 12 13 14] [15 16 17]
	//   2 "xyz"      4.5 false   22     1 [22 23 24 25] [26 27 28]
	//   3 "ssss"     4.9 false   33     2 [33 34 35 36] [37 38 39]
	//
	// [TypesGalore04]
	// uu8           ui t     s        i       f bb               b
	// []uint8    uint8 bool  string int float64 []byte        byte
	// [15 16 17]     0 true  "abc"    1     2.3 [11 12 13 14]   11
	// [26 27 28]     1 false "xyz"    2     4.5 [22 23 24 25]   22
	// [37 38 39]     2 false "ssss"   3     4.9 [33 34 35 36]   33
}

func TestNewTableReorderCols(t *testing.T) {
	var err error
	var table *Table
	var tableString string = `
	[TypesGalore05]
    i   s      f       t     b    ui    bb            uu8
    int string float64 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     false 33   2     [33 34 35 36] [37 38 39]
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	var tests = []struct {
		newColsOrderNames []string
		valid             bool
	}{
		{[]string{"t", "s", "f", "i", "b", "ui", "bb", "uu8"}, true}, // just right
		{[]string{"i", "s", "f", "t", "b", "bb", "ui", "uu8"}, true}, // just right
		// In two minds about making this an error:
		//		{ []string{"i", "s", "f", "t", "b", "ui", "bb", "uu8"     }, false },	// Already in reorder sequence
		{[]string{"x", "s", "f", "t", "b", "ui", "bb", "uu8"}, false},      // no col "x"
		{[]string{"i", "s", "f", "t", "b", "ui", "bb"}, false},             // too few
		{[]string{"i", "s", "f", "t", "b", "ui", "bb", "uu8", "y"}, false}, // too many
		{[]string{"i", "s", "f", "t", "b", "ui", "ui", "uu8"}, false},      // duplicate col name
	}

	for i, test := range tests {
		_, err := table.NewTableReorderCols(test.newColsOrderNames...)
		if err == nil != test.valid {
			t.Fatalf("test[%d]: newColsOrderNames: %v (%v)", i, test.newColsOrderNames, err)
		}
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}
	}
}

func TestReorderCols(t *testing.T) {
	var err error
	var table *Table
	var tableString string = `
	[TypesGalore06]
    i   s      f       t     b    ui    bb            uu8
    int string float64 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     false 33   2     [33 34 35 36] [37 38 39]
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	var tests = []struct {
		newColsOrderNames []string
		valid             bool
	}{
		{[]string{"t", "s", "f", "i", "b", "ui", "bb", "uu8"}, true}, // just right
		{[]string{"i", "s", "f", "t", "b", "bb", "ui", "uu8"}, true}, // just right
		// In two minds about making this an error:
		//		{ []string{"i", "s", "f", "t", "b", "ui", "bb", "uu8"     }, false },	// Already in reorder sequence
		{[]string{"x", "s", "f", "t", "b", "ui", "bb", "uu8"}, false},      // no col "x"
		{[]string{"i", "s", "f", "t", "b", "ui", "bb"}, false},             // too few
		{[]string{"i", "s", "f", "t", "b", "ui", "bb", "uu8", "y"}, false}, // too many
		{[]string{"i", "s", "f", "t", "b", "ui", "ui", "uu8"}, false},      // duplicate col name
	}

	for i, test := range tests {
		err := table.ReorderCols(test.newColsOrderNames...)
		if err == nil != test.valid {
			t.Fatalf("test[%d]: newColsOrderNames: %v (%v)", i, test.newColsOrderNames, err)
		}
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}
	}
}

func TestNewTableReorderColsByColIndex(t *testing.T) {
	var err error
	var table *Table
	var tableString string = `
	[TypesGalore07]
    i   s      f       t     b    ui    bb            uu8
    int string float64 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     false 33   2     [33 34 35 36] [37 38 39]
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	var tests = []struct {
		newColsOrderIndices []int
		valid               bool
	}{
		{[]int{7, 1, 2, 3, 4, 5, 6, 0}, true}, // Just right
		{[]int{0, 1, 5, 3, 4, 2, 6, 7}, true}, // Just right
		// In two minds about making this an error:
		//		{ []int{0, 1, 2, 3, 4, 5, 6, 7},	false },	// Already in reorder sequence
		{[]int{0, 1, 2, 3, 4, 5, 7, 7}, false},  // Duplicate 7
		{[]int{}, false},                        // This puts [] into error
		{[]int{7}, false},                       // This puts [7] into error
		{[]int{0, 1, 2, 3, 4, 5, 6, 8}, false},  // Gap in sequence
		{[]int{0, -1, 2, 3, 4, 5, 6, 8}, false}, // Out of range index
	}

	for i, test := range tests {
		_, err := table.NewTableReorderColsByColIndex(test.newColsOrderIndices...)
		if err == nil != test.valid {
			t.Fatalf("test[%d]: newColsOrderIndices %v %t: %v", i, test.newColsOrderIndices, test.valid, err)
		}
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}
	}
}

func TestReorderColsByColIndex(t *testing.T) {
	var err error
	var table *Table
	var tableString string = `
	[TypesGalore08]
    i   s      f       t     b    ui    bb            uu8
    int string float64 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     false 33   2     [33 34 35 36] [37 38 39]
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	var tests = []struct {
		newColsOrderIndices []int
		valid               bool
	}{
		{[]int{7, 1, 2, 3, 4, 5, 6, 0}, true}, // Just right
		{[]int{0, 1, 5, 3, 4, 2, 6, 7}, true}, // Just right
		// In two minds about making this an error:
		//		{ []int{0, 1, 2, 3, 4, 5, 6, 7},	false },	// Already in reorder sequence
		{[]int{0, 1, 2, 3, 4, 5, 7, 7}, false},  // Duplicate 7
		{[]int{}, false},                        // This puts [] into error
		{[]int{7}, false},                       // This puts [7] into error
		{[]int{0, 1, 2, 3, 4, 5, 6, 8}, false},  // Gap in sequence
		{[]int{0, -1, 2, 3, 4, 5, 6, 8}, false}, // Out of range index
	}

	for i, test := range tests {
		err := table.ReorderColsByColIndex(test.newColsOrderIndices...)
		if err == nil != test.valid {
			t.Fatalf("test[%d]: newColsOrderIndices %v %t: %v", i, test.newColsOrderIndices, test.valid, err)
		}
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}
	}
}

func ExampleReverse() {
	var err error
	var table *Table

	var tableString string = `[planets]
	name            mass distance moons index mnemonic
	string       float64  float64   int   int string
	"Sun"     333333.0        0.0     0     0 ""
	"Mercury"      0.055      0.4     0     1 "my"
	"Venus"        0.815      0.7     0     2 "very"
	"Earth"        1.0        1.0     1     3 "elegant"
	"Mars"         0.107      1.5     2     4 "mother"
	"Jupiter"    318.0        5.2    79     5 "just"
	"Saturn"      95.0        9.5    82     6 "sat"
	"Uranus"      15.0       19.2    27     7 "upon"
	"Neptune"     17.0       30.6    13     8 "nine"
	"Pluto"        0.002     39.4     5     9 "porcupines"
	`

	table, err = NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	err = table.Reverse()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	// Output:
	// [planets]
	// name            mass distance moons index mnemonic
	// string       float64  float64   int   int string
	// "Sun"     333333.0        0.0     0     0 ""
	// "Mercury"      0.055      0.4     0     1 "my"
	// "Venus"        0.815      0.7     0     2 "very"
	// "Earth"        1.0        1.0     1     3 "elegant"
	// "Mars"         0.107      1.5     2     4 "mother"
	// "Jupiter"    318.0        5.2    79     5 "just"
	// "Saturn"      95.0        9.5    82     6 "sat"
	// "Uranus"      15.0       19.2    27     7 "upon"
	// "Neptune"     17.0       30.6    13     8 "nine"
	// "Pluto"        0.002     39.4     5     9 "porcupines"
	//
	// [planets]
	// name            mass distance moons index mnemonic
	// string       float64  float64   int   int string
	// "Pluto"        0.002     39.4     5     9 "porcupines"
	// "Neptune"     17.0       30.6    13     8 "nine"
	// "Uranus"      15.0       19.2    27     7 "upon"
	// "Saturn"      95.0        9.5    82     6 "sat"
	// "Jupiter"    318.0        5.2    79     5 "just"
	// "Mars"         0.107      1.5     2     4 "mother"
	// "Earth"        1.0        1.0     1     3 "elegant"
	// "Venus"        0.815      0.7     0     2 "very"
	// "Mercury"      0.055      0.4     0     1 "my"
	// "Sun"     333333.0        0.0     0     0 ""
}

func ExampleTable_ShuffleDeterministic() {
	var err error
	var table *Table

	var tableString string = `[planets]
	name            mass distance moons index mnemonic
	string       float64  float64   int   int string
	"Sun"     333333.0        0.0     0     0 ""
	"Mercury"      0.055      0.4     0     1 "my"
	"Venus"        0.815      0.7     0     2 "very"
	"Earth"        1.0        1.0     1     3 "elegant"
	"Mars"         0.107      1.5     2     4 "mother"
	"Jupiter"    318.0        5.2    79     5 "just"
	"Saturn"      95.0        9.5    82     6 "sat"
	"Uranus"      15.0       19.2    27     7 "upon"
	"Neptune"     17.0       30.6    13     8 "nine"
	"Pluto"        0.002     39.4     5     9 "porcupines"
	`

	table, err = NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	err = table.ShuffleDeterministic()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	// Output:
	// [planets]
	// name            mass distance moons index mnemonic
	// string       float64  float64   int   int string
	// "Sun"     333333.0        0.0     0     0 ""
	// "Mercury"      0.055      0.4     0     1 "my"
	// "Venus"        0.815      0.7     0     2 "very"
	// "Earth"        1.0        1.0     1     3 "elegant"
	// "Mars"         0.107      1.5     2     4 "mother"
	// "Jupiter"    318.0        5.2    79     5 "just"
	// "Saturn"      95.0        9.5    82     6 "sat"
	// "Uranus"      15.0       19.2    27     7 "upon"
	// "Neptune"     17.0       30.6    13     8 "nine"
	// "Pluto"        0.002     39.4     5     9 "porcupines"
	//
	// [planets]
	// name            mass distance moons index mnemonic
	// string       float64  float64   int   int string
	// "Neptune"     17.0       30.6    13     8 "nine"
	// "Sun"     333333.0        0.0     0     0 ""
	// "Mars"         0.107      1.5     2     4 "mother"
	// "Venus"        0.815      0.7     0     2 "very"
	// "Earth"        1.0        1.0     1     3 "elegant"
	// "Pluto"        0.002     39.4     5     9 "porcupines"
	// "Mercury"      0.055      0.4     0     1 "my"
	// "Jupiter"    318.0        5.2    79     5 "just"
	// "Uranus"      15.0       19.2    27     7 "upon"
	// "Saturn"      95.0        9.5    82     6 "sat"
}

func ExampleTable_ShuffleRandom() {
	var err error
	var table *Table

	var tableString string = `[planets]
	name            mass distance moons index mnemonic
	string       float64  float64   int   int string
	"Sun"     333333.0        0.0     0     0 ""
	"Mercury"      0.055      0.4     0     1 "my"
	"Venus"        0.815      0.7     0     2 "very"
	"Earth"        1.0        1.0     1     3 "elegant"
	"Mars"         0.107      1.5     2     4 "mother"
	"Jupiter"    318.0        5.2    79     5 "just"
	"Saturn"      95.0        9.5    82     6 "sat"
	"Uranus"      15.0       19.2    27     7 "upon"
	"Neptune"     17.0       30.6    13     8 "nine"
	"Pluto"        0.002     39.4     5     9 "porcupines"
	`

	table, err = NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	err = table.ShuffleRandom()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("ShuffleRandom() is \"truly\" random, so no predictable output.")

	// Output:
	// [planets]
	// name            mass distance moons index mnemonic
	// string       float64  float64   int   int string
	// "Sun"     333333.0        0.0     0     0 ""
	// "Mercury"      0.055      0.4     0     1 "my"
	// "Venus"        0.815      0.7     0     2 "very"
	// "Earth"        1.0        1.0     1     3 "elegant"
	// "Mars"         0.107      1.5     2     4 "mother"
	// "Jupiter"    318.0        5.2    79     5 "just"
	// "Saturn"      95.0        9.5    82     6 "sat"
	// "Uranus"      15.0       19.2    27     7 "upon"
	// "Neptune"     17.0       30.6    13     8 "nine"
	// "Pluto"        0.002     39.4     5     9 "porcupines"
	//
	// ShuffleRandom() is "truly" random, so no predictable output.
}

func TestTable_SwapColsByColIndex(t *testing.T) {
	var err error
	var table *Table
	var tableString string = `
	[TypesGalore09]
    i   s      f       t     b    ui    bb            uu8
    int string float64 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     false 33   2     [33 34 35 36] [37 38 39]
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	var tests = []struct {
		swapCols []int
		colNames []string
		colTypes []string
	}{
		{[]int{0, 1}, []string{"s", "i", "f", "t", "b", "ui", "bb", "uu8"}, []string{"string", "int", "float64", "bool", "byte", "uint8", "[]byte", "[]uint8"}},
		{[]int{6, 7}, []string{"i", "s", "f", "t", "b", "ui", "uu8", "bb"}, []string{"int", "string", "float64", "bool", "byte", "uint8", "[]uint8", "[]byte"}},
		{[]int{0, 7}, []string{"uu8", "s", "f", "t", "b", "ui", "bb", "i"}, []string{"[]uint8", "string", "float64", "bool", "byte", "uint8", "[]byte", "int"}},
	}

	for i, test := range tests {
		err := table.SwapColsByColIndex(test.swapCols[0], test.swapCols[1])
		if err != nil {
			t.Fatal(err)
		}

		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		colNames := table.getColNames()
		if !stringSliceEquals(colNames, test.colNames) {
			t.Fatalf("test[%d]: table.SwapColsByColIndex(%v) swapped colNames %v != expected %v", i, test.swapCols, colNames, test.colNames)
		}

		colTypes := table.getColTypes()
		if !stringSliceEquals(colTypes, test.colTypes) {
			t.Fatalf("test[%d]: table.SwapColsByColIndex(%v) swapped colTypes %v != expected %v", i, test.swapCols, colTypes, test.colTypes)
		}

		// Return table to original data for next iteration.
		table, err = NewTableFromString(tableString)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestTable_SwapCols(t *testing.T) {
	var err error
	var table *Table
	var tableString string = `
	[TypesGalore10]
    i   s      f       t     b    ui    bb            uu8
    int string float64 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     false 33   2     [33 34 35 36] [37 38 39]
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}

	var tests = []struct {
		swapCols []string
		colNames []string
		colTypes []string
	}{
		{[]string{"i", "s"}, []string{"s", "i", "f", "t", "b", "ui", "bb", "uu8"}, []string{"string", "int", "float64", "bool", "byte", "uint8", "[]byte", "[]uint8"}},
		{[]string{"bb", "uu8"}, []string{"i", "s", "f", "t", "b", "ui", "uu8", "bb"}, []string{"int", "string", "float64", "bool", "byte", "uint8", "[]uint8", "[]byte"}},
		{[]string{"i", "uu8"}, []string{"uu8", "s", "f", "t", "b", "ui", "bb", "i"}, []string{"[]uint8", "string", "float64", "bool", "byte", "uint8", "[]byte", "int"}},
	}

	for i, test := range tests {
		err := table.SwapCols(test.swapCols[0], test.swapCols[1])
		if err != nil {
			t.Fatal(err)
		}

		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		colNames := table.getColNames()
		if !stringSliceEquals(colNames, test.colNames) {
			t.Fatalf("test[%d]: table.SwapColsByColIndex(%v) swapped colNames %v != expected %v", i, test.swapCols, colNames, test.colNames)
		}

		colTypes := table.getColTypes()
		if !stringSliceEquals(colTypes, test.colTypes) {
			t.Fatalf("test[%d]: table.SwapColsByColIndex(%v) swapped colTypes %v != expected %v", i, test.swapCols, colTypes, test.colTypes)
		}

		// Return table to original data for next iteration.
		table, err = NewTableFromString(tableString)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func stringSliceEquals(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

/*
	NOTE: ExampleGoFmtProgramString() required some
	careful tweaking to get it right. The test output
	reported spaces expected, but the actual output
	of GoFmtProgramString() has 5 tabs, which needed
	to be placed in the Output area. A larger example
	would have been very tedious to get right.
*/
func ExampleGoFmtProgramString() {
	var err error

	goProgramString :=
		`package main
import "os"
func main() {
i := 42
if i != 42 { os.Exit(1) }
}`

	formattedGoProgramString, err := util.GoFmtProgramString(goProgramString)
	if err != nil {
		log.Println(err)
	}

	hasTabs := strings.Contains(formattedGoProgramString, "\t")
	fmt.Printf("hasTabs = %t\n", hasTabs)

	numTabs := strings.Count(formattedGoProgramString, "\t")
	fmt.Printf("numTabs = %d\n", numTabs)

	fmt.Println()

	fmt.Println(formattedGoProgramString)

	// Output:
	// hasTabs = true
	// numTabs = 5
	//
	// package main
	//
	// import "os"
	//
	// func main() {
	// 	i := 42
	// 	if i != 42 {
	// 		os.Exit(1)
	// 	}
	// }
}

/*
 7 \a   U+0007 alert or bell
 8 \b   U+0008 backspace
 9 \t   U+0009 horizontal tab
10 \n   U+000A line feed or newline
11 \v   U+000b vertical tab
12 \f   U+000C form feed
13 \r   U+000D carriage return
34 \"   U+0022 double quote  (valid escape only within string literals) CAN    use in string literal
44 \'   U+0027 single quote  (valid escape only within rune literals)   CANNOT use in string literal
92 \\   U+005c backslash
*/
// This is a test of my undestanding.
func TestUnquote(t *testing.T) {
	var tests = []struct {
		quoted   string
		expected []byte
	}{
		{`"abc"`, []byte{97, 98, 99}},
		{"`abc`", []byte{97, 98, 99}},
		{"\"abc\"", []byte{97, 98, 99}},
		{"'a'", []byte{97}},
		{`"\a\b\t\n\v\f\r\"\\Y'Z"`, []byte{7, 8, 9, 10, 11, 12, 13, 34, 92, 89, 39, 90}},
		{`"\a"`, []byte{7}},
		{`"\a\b"`, []byte{7, 8}},
		{`"\a\b\f\n\r\t\v"`, []byte{7, 8, 12, 10, 13, 9, 11}},
		{`"\a\b\f\n\r\t\v'"`, []byte{7, 8, 12, 10, 13, 9, 11, 39}},
		{`"'"`, []byte{39}},  // \'   U+0027 single quote  (valid escape only within rune   literals)
		{`"\""`, []byte{34}}, // \"   U+0022 double quote  (valid escape only within string literals)
	}

	for i, test := range tests {

		// Unquote
		// Test that Unquote returns a string consisting of the expected bytes.

		var unquoted string
		unquoted, err := strconv.Unquote(test.quoted)
		if err != nil {
			t.Fatalf("strconv.Unquote(%q) error: %v", test.quoted, err)
		}

		// Check for len() error.
		if len(unquoted) != len(test.expected) {
			t.Fatalf("len(unquoted)=%d != len(test.expected)=%d", len(unquoted), len(test.expected))
		}

		for j := 0; j < min(len(unquoted), len(test.expected)); j++ { // min() in case of len() error above.
			if unquoted[j] != test.expected[j] {
				t.Fatalf("test[%d]: expecting quoted[%d] = %d, not %d", i, j, test.expected[j], unquoted[j])
			}
		}

		// Quote
		// Test that Quote reverses the string back to the original quoted string.

		const printable = 32
		var quoted string = strconv.Quote(unquoted)
		// Ignore first and last char which may be ' or ` rather than "
		for j := 1; j < min(len(quoted)-1, len(test.quoted)); j++ { // min() in case of len() error above.
			if quoted[j] != test.quoted[j] {
				if test.quoted[j] >= printable && quoted[j] >= printable {
					t.Fatalf("test[%d]: expecting quoted[%d] = %d '%c', not %d '%c'",
						i, j, test.quoted[j], test.quoted[j], quoted[j], quoted[j])
				} else {
					t.Fatalf("test[%d]: expecting quoted[%d] = %d, not %d", i, j, test.quoted[j], quoted[j])
				}
			}
		}
	}
}

func (table *Table) multiplyInt(factorCol1 string, factorCol2 string, productCol string, row int) (err error) {
	var factor1, factor2, product int

	factor1, err = table.GetInt(factorCol1, row)
	if err != nil {
		return err
	}

	factor2, err = table.GetInt(factorCol2, row)
	if err != nil {
		return err
	}

	product = factor1 * factor2

	err = table.SetInt(productCol, row, product)
	if err != nil {
		return err
	}

	return nil
}

var goKeywords = map[string]string{
	"break":       "break",
	"default":     "default",
	"func":        "func",
	"interface":   "interface",
	"select":      "select",
	"case":        "case",
	"defer":       "defer",
	"go":          "go",
	"map":         "map",
	"struct":      "struct",
	"chan":        "chan",
	"else":        "else",
	"goto":        "goto",
	"package":     "package",
	"switch":      "switch",
	"const":       "const",
	"fallthrough": "fallthrough",
	"if":          "if",
	"range":       "range",
	"type":        "type",
	"continue":    "continue",
	"for":         "for",
	"import":      "import",
	"return":      "return",
	"var":         "var",
}

// Transplanted from flattables.go for comparison with Go equivalent.
func isGoKeyword(name string) bool {
	nameLower := strings.ToLower(name)
	_, exists := goKeywords[nameLower]
	return exists
}

// Note: Leading lowercase in go113NumericLiteralsBinOctHex is required for it to be recognised as an Example!
func ExampleNewTableSetFromString_go113NumericLiteralsBinOctHex() {
	tableString :=
		`[int_literals]
	decimal binary octal hex single zeros
	int     int    int   int int    int
	10      0b1010 0o12  0xA 0      0
	10      0B1010 0O12  0XA 1      00
	10      0b1010 012   0xa 2      000
	10      0B1010 0012  0xA 3      0000

	[int8_literals]
	decimal binary octal hex  single zeros
	int8    int8   int8  int8 int8   int8
	10      0b1010 0o12  0xA  0      0
	10      0B1010 0O12  0XA  1      00
	10      0b1010 012   0xa  2      000
	10      0B1010 0012  0xA  3      0000

	[uint8_literals]
	decimal binary octal hex   single zeros
	uint8   uint8  uint8 uint8 uint8  uint8
	10      0b1010 0o12  0xA   0      0
	10      0B1010 0O12  0XA   1      00
	10      0b1010 012   0xa   2      000
	10      0B1010 0012  0xA   3      0000

	[any_literals]
	dec int = 11
	bin int = 0b1011
	oct int = 0o13
	hex int = 0xB

	[slice_literals]
	xx                       bb
	[]uint8                  []byte
	[0b1010 0o12 0B1010 0xA] [0xa 0b1010 0o12 012]
	[10 0o12 0b1010 0xA]     [0Xa 0B1010 0o12 012]
	[30]                     [90 81 72]
	[]                       []
	[10]                     [90 81 72 63 255]
	`

	tableSet, err := NewTableSetFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(tableSet)

	// Output:
	// [int_literals]
	// decimal binary octal hex single zeros
	//     int    int   int int    int   int
	//      10     10    10  10      0     0
	//      10     10    10  10      1     0
	//      10     10    10  10      2     0
	//      10     10    10  10      3     0
	//
	// [int8_literals]
	// decimal binary octal  hex single zeros
	//    int8   int8  int8 int8   int8  int8
	//      10     10    10   10      0     0
	//      10     10    10   10      1     0
	//      10     10    10   10      2     0
	//      10     10    10   10      3     0
	//
	// [uint8_literals]
	// decimal binary octal   hex single zeros
	//   uint8  uint8 uint8 uint8  uint8 uint8
	//      10     10    10    10      0     0
	//      10     10    10    10      1     0
	//      10     10    10    10      2     0
	//      10     10    10    10      3     0
	//
	// [any_literals]
	// dec int = 11
	// bin int = 11
	// oct int = 11
	// hex int = 11
	//
	// [slice_literals]
	// xx            bb
	// []uint8       []byte
	// [10 10 10 10] [10 10 10 10]
	// [10 10 10 10] [10 10 10 10]
	// [30]          [90 81 72]
	// []            []
	// [10]          [90 81 72 63 255]
}

func TestParseInt(t *testing.T) {

	var tests = []struct {
		s             string
		bitSize       int
		expected      int64
		expectedError error
	}{
		/*  0 */ {"10", 64, 10, nil},
		/*  1 */ {"0b1010", 64, 10, nil},
		/*  2 */ {"0B1010", 64, 10, nil},
		/*  3 */ {"0o12", 64, 10, nil},
		/*  4 */ {"0O12", 64, 10, nil},
		/*  5 */ {"012", 64, 10, nil},
		/*  6 */ {"0xA", 64, 10, nil},
		/*  7 */ {"0XA", 64, 10, nil},
		/*  8 */ {"0X", 64, 10, errors.New("hex missing digits")},
		/*  9 */ {"", 64, 10, errors.New("empty string")},
		/* 10 */ {"0B1012", 64, 10, errors.New("invalid bin digit: 2")},
		/* 11 */ {"0", 64, 0, nil},
		/* 12 */ {"00", 64, 0, nil},
		/* 13 */ {"000", 64, 0, nil},
		/* 14 */ {"ob1010", 0, 0, errors.New("invalid prefix: ob")},
		/* 15 */ {"oxa", 0, 0, errors.New("invalid prefix: ox")},
	}

	for testIndex, test := range tests {

		result, err := parseInt(test.s, test.bitSize)
		if err != nil && test.expectedError == nil {
			t.Fatalf("[%d] parseInt(%q) %v", testIndex, test.s, err)
		}

		if result != test.expected && test.expectedError == nil {
			t.Fatalf("[%d] parseInt(%q) expecting %d but got %d", testIndex, test.s, test.expected, result)
		}

		if err == nil && test.expectedError != nil {
			t.Fatalf("[%d] parseInt(%q) expecting error %v but got err == nil", testIndex, test.s, test.expectedError)
		}
	}
}

var globalTableSetStringIntDecimal string = `
	[ints]
	i	i8		i16		i32		i64		ui		ui8		ui16	ui32	ui64
	int	int8	int16	int32	int64	uint	uint8	uint16	uint32	uint64
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	-10	-10		-10		-10		-10		10		10		10		10		10
	`

var globalTableSetStringIntNonDecimal string = `
	[ints]
	i		i8			i16			i32			i64			ui		ui8		ui16	ui32	ui64
	int		int8		int16		int32		int64		uint	uint8	uint16	uint32	uint64
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	0b1010	0b1010		0b1010		0b1010		0b1010		0xA		0xA		0xA		0xA		0o12
	`

func BenchmarkNewTableSetFromStringIntDecimal_padded(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		_, err = NewTableSetFromString(globalTableSetStringIntDecimal)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkNewTableSetFromStringIntNonDecimal_padded(b *testing.B) {
	if go_1_13_number_literals {
		var err error
		for i := 0; i < b.N; i++ {
			_, err = NewTableSetFromString(globalTableSetStringIntNonDecimal)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

// Testing table with slice of uint: []uint with BIN, OCT and HEX
func TestNewTableFromStringGo_1_3_NonDecimalLiterals(t *testing.T) {
	if go_1_13_number_literals {
		var err error
		s :=
			`[TableX]
	i   x                   		f           bb                  	s       b
	int []uint8             		float64     []byte              	string  byte
	1   [0b1010 0o12 0B1010 0xA]    1           [0xa 0b1010 0o12 012]   "one"   11
	2   [10 0o12 0b1010 0xA]       	2           [0Xa 0B1010 0o12 012]   "two"   22
	3   [30]          				3           [90 81 72]          	"three" 33
	4   []    						4           [90 81 72 63 255]   	"four"  44
	4   [10]    					5           [90 81 72 63 255]   	"four"  44
	`
		table, err := NewTableFromString(s)
		if err != nil {
			t.Fatal(err)
		}

		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
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
			t.Fatal(err)
		}

		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		err = table.AppendCol("bite", "[]byte")
		if err != nil {
			t.Fatal(err)
		}
		if isValid, err := table.IsValidTable(); !isValid {
			t.Fatal(err)
		}

		table, err = NewTableFromString(globalTableSetStringIntNonDecimal)
		if err != nil {
			t.Fatal(err)
		}

		table, err = NewTableFromString(globalTableSetStringIntDecimal)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestParserGlobalVars(t *testing.T) {

	const nameWidth = 30

	var stringTests = []struct {
		name   string
		global *string
	}{
		//		{"runeRegexpString", &runeRegexpString},			// now a const
		{"intRegexpString", &intRegexpString},
		{"uintRegexpString", &uintRegexpString},
		{"uintSliceRegexpString", &uintSliceRegexpString},
		//		{"namePattern", &namePattern},						// now a const
		//		{"tableNamePattern", &tableNamePattern},			// now a const
	}

	for testIndex, test := range stringTests {
		if *test.global == "" {
			t.Fatalf("parser.go: global string test[%d] var %s string == \"\"", testIndex, test.name)
			// fmt.Printf("[%2d] %-*s: %q\n", testIndex, nameWidth, test.name, *test.global)
		}
	}

	var regexpTests = []struct {
		name   string
		global *regexp.Regexp
	}{
		{"stringRegexp", stringRegexp},
		{"boolRegexp", boolRegexp},
		{"runeRegexp", runeRegexp},
		{"floatRegexp", floatRegexp},
		{"tableNameRegexp", tableNameRegexp},
		{"colNameRegexp", colNameRegexp},
		{"whiteRegexp", whiteRegexp},
		{"equalsRegexp", equalsRegexp},
	}

	for testIndex, test := range regexpTests {
		if test.global == nil {
			t.Fatalf("parser.go: global regexp test[%d] var %s *regexp.Regexp == nil", testIndex, test.name)
			// fmt.Printf("[%2d] %-*s: %v\n", testIndex, nameWidth, test.name, *test.global)
		}
	}
}

// reflect.TypeOf(val) is MUCH faster than fmt.Sprintf("%T", val)
func BenchmarkSprintfType(b *testing.B) {
	var val interface{}
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%T", val)
	}
}
