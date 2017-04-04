package gotable

import (
	"fmt"
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
//	var testsTableSet *GoTableSet
//	var tests *GoTable
//	var setupName string = "Fred"
//
//	testsTableSet, err = NewGoTableSetFromString(
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
//	var table *GoTable
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
//		table, err = NewGoTable(setupName)
//		if err != nil {
//			t.Error(err)
//		}
//
//		err = table.RenameTable(input)
//		if (err == nil) != succeeds {
//			t.Errorf("Error renaming to %q: %s", output, err)
//		}
//
//		var tableName string = table.TableName()
//		if tableName != output {
//			t.Errorf("Expected %q, not %q", output, tableName)
//		}
//	}
//}
*/

func TestRenameTable(t *testing.T) {
	var err error
	var table *GoTable
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

		table, err = NewGoTable(setupName)
		if err != nil {
			t.Error(err)
		}

		err = table.RenameTable(test.input)
		if (err == nil) != test.succeeds {
			t.Errorf("Error renaming to %q: %s", test.output, err)
		}

		var tableName string = table.TableName()
		if tableName != test.output {
			t.Errorf("Expected %q, not %q", test.output, tableName)
		}
	}
}

func TestGoTableSetRenameTable(t *testing.T) {
	/*
		goTableSet, err := NewGoTableSetFromString(`[Wilma]`)
		if err != nil {
			panic(err)
		}
	*/
	//	fmt.Printf("goTableSet.TableCount() = %d\n", goTableSet.TableCount())

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
		goTableSet, err := NewGoTableSetFromString(`[Wilma]`)
		if err != nil {
			t.Error(err)
		}
		err = goTableSet.RenameTable(test.renameFrom, test.renameTo)
		if (err == nil) != test.succeeds {
			t.Errorf("test[%d]: Error renaming from %q to %q: %v", i, test.renameFrom, test.renameTo, err)
		}
	}
}

func TestReadString1(t *testing.T) {
	goTableSet, err := NewGoTableSetFromString(
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
		table, err := goTableSet.Table(test.tableName)
		if err != nil {
			t.Errorf("[%d] %v", i, err)
		}
		if (err == nil) != test.succeeds {
			t.Errorf("test[%d]: err == %v but expecting succeeds == %t", i, err, test.succeeds)
		}

		colCount := table.ColCount()
		if colCount != test.colCount {
			tableName := table.TableName()
			t.Errorf("test[%d]: Expecting [%s] colCount %d, not %d\n", i, tableName, test.colCount, colCount)
		}

		rowCount := table.RowCount()
		if rowCount != test.rowCount {
			tableName := table.TableName()
			t.Errorf("test[%d]: Expecting [%s] rowCount %d, not %d\n", i, tableName, test.rowCount, rowCount)
		}
	}
}

func TestReadString2(t *testing.T) {
	_, err := NewGoTableSetFromString(
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
	_, err := NewGoTableSetFromString(
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
	_, err := NewGoTableSetFromString(
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
	goTableSet, err := NewGoTableSetFromString(
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
		table, err := goTableSet.Table(test.tableName)
		if err != nil {
			t.Errorf("[%d] %v", i, err)
		}
		if (err == nil) != test.succeeds {
			t.Errorf("test[%d]: err == %v but expecting succeeds == %t", i, err, test.succeeds)
		}

		colCount := table.ColCount()
		if colCount != test.colCount {
			tableName := table.TableName()
			t.Errorf("test[%d]: Expecting [%s] colCount %d, not %d\n", i, tableName, test.colCount, colCount)
		}

		rowCount := table.RowCount()
		if rowCount != test.rowCount {
			tableName := table.TableName()
			t.Errorf("test[%d]: Expecting [%s] rowCount %d, not %d\n", i, tableName, test.rowCount, rowCount)
		}
	}
}

func TestReadString6(t *testing.T) {
	goTableSet, err := NewGoTableSetFromString(
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
		table, err := goTableSet.Table(test.tableName)
		if err != nil {
			t.Errorf("[%d] %v", i, err)
		}
		if (err == nil) != test.succeeds {
			t.Errorf("test[%d]: err == %v but expecting succeeds == %t", i, err, test.succeeds)
		}

		colCount := table.ColCount()
		if colCount != test.colCount {
			tableName := table.TableName()
			t.Errorf("test[%d]: Expecting [%s] colCount %d, not %d\n", i, tableName, test.colCount, colCount)
		}

		rowCount := table.RowCount()
		if rowCount != test.rowCount {
			tableName := table.TableName()
			t.Errorf("test[%d]: Expecting [%s] rowCount %d, not %d\n", i, tableName, test.rowCount, rowCount)
		}
	}
}

func TestReadString7(t *testing.T) {
	_, err := NewGoTableSetFromString(
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
	_, err := NewGoTableSetFromString(
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

func ExampleNewGoTableSet() {
	tableSetName := "MyTableSet"
	tableSet, err := NewGoTableSet(tableSetName)
	if err != nil {
		panic(err)
	}
	tableCount := tableSet.TableCount()
	name := tableSet.GoTableSetName()
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
			t.Errorf("test[%d]: Expecting %f, not %f", i, test.rounded, rounded)
		}
	}
}

func TestSetAndGetFunctions(t *testing.T) {
	var bVal bool
	//	byte 			// alias for uint8
	//	complex128 		// The set of all complex numbers with float64 real and imaginary parts
	//	complex64		// The set of all complex numbers with float32 real and imaginary parts
	/*
	   	var f32Val float32	// The set of all IEEE-754 32-bit floating-point numbers
	   	var f64Val float64	// The set of all IEEE-754 64-bit floating-point numbers
	   	var iVal   int		// Machine-dependent
	   	var i16Val int16	// The set of all signed 16-bit integers (-32768 to 32767)
	   	var i32Val int32	// The set of all signed 32-bit integers (-2147483648 to 2147483647)
	   	var i64Val int64	// The set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	   	var i8Val  int8		// The set of all signed  8-bit integers (-128 to 127)
	   //	rune 			// alias for int32
	   	var sVal    string
	*/
	var uiVal uint // Machine-dependent
	/*
		var ui16Val uint16	// The set of all unsigned 16-bit integers (0 to 65535)
		var ui32Val uint32	// The set of all unsigned 32-bit integers (0 to 4294967295)
		var ui64Val uint64	// The set of all unsigned 64-bit integers (0 to 18446744073709551615)
		var ui8Val  uint8	// The set of all unsigned  8-bit integers (0 to 255)
	*/

	var err error
	var table *GoTable
	const rowIndex = 0
	var colIndex = 0

	table, err = NewGoTable("SetAndGet")
	if err != nil {
		t.Error(err)
	}

	table.AddRow()

	// Bool tests

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
		t.Errorf("Expecting GetBool() value %t, not %t\n", true, bVal)
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
		t.Errorf("Expecting GetBoolByColIndex() value %t, not %t\n", true, bVal)
	}

	// Uint tests

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
		t.Errorf("Expecting GetUint() value %d, not %d\n", 55, uiVal)
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
		t.Errorf("Expecting GetUintByColIndex() value %d, not %d\n", 66, uiVal)
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
		_, err = NewGoTableSetFromString(test.input)
		if err == nil != test.valid {
			switch test.valid {
			case true:
				t.Errorf("test[%d]: %v", i, err)
			case false:
				t.Errorf("test[%d]: NewGoTableSetFromString(): Expecting this input to fail with a range error: %s", i, test.input)
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
		_, err = NewGoTableSetFromString(test.input)
		if err == nil != test.valid {
			switch test.valid {
			case true:
				t.Errorf("test[%d]: %v", i, err)
			case false:
				t.Errorf("test[%d]: NewGoTableSetFromString(): Expecting this input to fail with a range error: %s", i, test.input)
			}
		}
	}
}

var goTableSetString string = `
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

func BenchmarkNewGoTableSetFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewGoTableSetFromString(goTableSetString)
	}
}

func BenchmarkGoTableSetToString(b *testing.B) {
	// Set up for benchmark.
	goTableSet, err := NewGoTableSetFromString(goTableSetString)
	if err != nil {
		b.Error(err)
	}

	var _ string
	for i := 0; i < b.N; i++ {
		_ = goTableSet.String()
	}
}

func BenchmarkGobEncode(b *testing.B) {
	// Set up for benchmark.
	goTableSet, err := NewGoTableSetFromString(goTableSetString)
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		_, err := goTableSet.GobEncode()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGobDecode(b *testing.B) {
	// Set up for benchmark.
	goTableSet, err := NewGoTableSetFromString(goTableSetString)
	if err != nil {
		b.Error(err)
	}

	// Set up for benchmark.
	gobEncodedTableSet, err := goTableSet.GobEncode()
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
	F_bool bool =
	F_string string =
	T_float32 float32 =
	T_float64 float64 =
	T_int int =
	T_int16 int16 =
	T_int32 int32 =
	T_int64 int64 =
	T_int8 int8 =
	T_uint uint =
	T_uint16 uint16 =
	T_uint32 uint32 =
	T_uint64 uint64 =
	T_uint8 uint8 =
    `

	tableSet, err := NewGoTableSetFromString(tableString)
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
			t.Error(fmt.Errorf("Expecting col name %s to have prefix \"T_\" or \"F_\" but found: %q", colName, colName))
		}

		if isNumeric != hasPrefixT {
			err := fmt.Errorf("col %s type %s unexpected IsNumeric: %t", colName, colType, isNumeric)
			t.Error(err)
		}
	}
}

func TestAddRow(t *testing.T) {
	tableString := `
    [table]
	F_bool bool =
	F_string string =
	T_float32 float32 =
	T_float64 float64 =
	T_int int =
	T_int16 int16 =
	T_int32 int32 =
	T_int64 int64 =
	T_int8 int8 =
	T_uint uint =
	T_uint16 uint16 =
	T_uint32 uint32 =
	T_uint64 uint64 =
	T_uint8 uint8 =
    `

	tableSet, err := NewGoTableSetFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	table, err := tableSet.Table("table")
	if err != nil {
		t.Error(err)
	}

	rowCount := table.RowCount()
	if rowCount != 0 {
		t.Error(fmt.Errorf("Expecting row count of 0, not: %d", rowCount))
	}

	err = table.AddRow()
	if err != nil {
		t.Error(err)
	}

	rowCount = table.RowCount()
	if rowCount != 1 {
		t.Error(fmt.Errorf("Expecting row count of 1, not: %d", rowCount))
	}

	err = table.DeleteRow(0)
	if err != nil {
		t.Error(err)
	}

	rowCount = table.RowCount()
	if rowCount != 0 {
		t.Error(fmt.Errorf("Expecting row count of 0, not: %d", rowCount))
	}
}

func TestColCount(t *testing.T) {
	tableString := `
    [table]
	F_bool bool =
	F_string string =
	T_float32 float32 =
	T_float64 float64 =
	T_int int =
	T_int16 int16 =
	T_int32 int32 =
	T_int64 int64 =
	T_int8 int8 =
	T_uint uint =
	T_uint16 uint16 =
	T_uint32 uint32 =
	T_uint64 uint64 =
	T_uint8 uint8 =
    `

	tableSet, err := NewGoTableSetFromString(tableString)
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
		t.Error(fmt.Errorf("Expecting col count of %d, not: %d", initialColCount, colCount))
	}

	err = table.AppendCol("ExtraCol", "bool")
	if err != nil {
		t.Error(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount + 1 {
		t.Error(fmt.Errorf("Expecting col count of %d, not: %d", initialColCount + 1, colCount))
	}

	lastCol := colCount-1
	err = table.DeleteColByColIndex(lastCol)
	if err != nil {
		t.Error(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount {
		t.Error(fmt.Errorf("Expecting col count of %d, not: %d", initialColCount, colCount))
	}

	err = table.AppendCol("AnotherCol", "string")
	if err != nil {
		t.Error(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount + 1 {
		t.Error(fmt.Errorf("Expecting col count of %d, not: %d", initialColCount + 1, colCount))
	}

	err = table.DeleteCol("AnotherCol")
	if err != nil {
		t.Error(err)
	}

	colCount = table.ColCount()
	if colCount != initialColCount {
		t.Error(fmt.Errorf("Expecting col count of %d, not: %d", initialColCount, colCount))
	}
}
