package gotable

import (
	"fmt"
	"testing"
	//	"strconv"
	//	"math"
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

func TestNewGoTableSetFromHtmlString1(t *testing.T) {

	tests := []struct {
		tableHtml string
		tableName string
		colCount  int
		rowCount  int
		succeeds  bool
		testCell  bool
	}{
		{ // test0 Well-formed HTML table
			`<table>
			<tr> <td>r0 c0</td> <td>r0 c1</td> <td>r0 c2</td> </tr>
			<tr> <td>test0</td> <td>r1 c1</td> <td>r1 c2</td> </tr>
			<tr> <td>r2 c0</td> <td>r2 c1</td> <td>r2 c2</td> </tr>
			</table>`,
			"table_0", 3, 3, true, true},
		{ // test1 Missing </table>
			`<table>
			<tr> <td>r0 c0</td> <td>r0 c1</td> <td>r0 c2</td> </tr>
			<tr> <td>test1</td> <td>r1 c1</td> <td>r1 c2</td> </tr>
			<tr> <td>r2 c0</td> <td>r2 c1</td> <td>r2 c2</td> </tr>`,
			"table_0", 3, 3, true, true},
		{ // test2 Missing </tr> </table>
			`<table>
			<tr> <td>r0 c0</td> <td>r0 c1</td> <td>r0 c2</td> </tr>
			<tr> <td>test2</td> <td>r1 c1</td> <td>r1 c2</td> </tr>
			<tr> <td>r2 c0</td> <td>r2 c1</td> <td>r2 c2</td>`,
			"table_0", 3, 3, true, true},
		{ // test3 Missing </td> </tr> </table>
			`<table>
			<tr> <td>r0 c0</td> <td>r0 c1</td> <td>r0 c2</td> </tr>
			<tr> <td>test3</td> <td>r1 c1</td> <td>r1 c2</td> </tr>
			<tr> <td>r2 c0</td> <td>r2 c1</td> <td>r2 c2`,
			"table_0", 3, 3, true, true},
		{ // test4 Missing <td> </td> </tr> </table>
			`<table>
			<tr> <td>r0 c0</td> <td>r0 c1</td> <td>r0 c2</td> </tr>
			<tr> <td>test4</td> <td>r1 c1</td> <td>r1 c2</td> </tr>
			<tr> <td>r2 c0</td> <td>r2 c1</td>     r2 c2`,
			"table_0", 3, 3, true, true},
		{ // test5 Missing col in row 0. Equivalent to additional col in row 1.
			`<table>
			<tr> <td>r0 c0</td> <td>r0 c1</td>                </tr>
			<tr> <td>test5</td> <td>r1 c1</td> <td>r1 c2</td> </tr>
			<tr> <td>r2 c0</td> <td>r2 c1</td> <td>r2 c2</td> </tr>
			</table>`,
			"table_0", 3, 3, true, true},
		{ // test6 Progressively more cols in additional rows.
			`<table>
			<tr> <td>r0 c0</td>                               </tr>
			<tr> <td>test6</td> <td>r1 c1</td>                </tr>
			<tr> <td>r2 c0</td> <td>r2 c1</td> <td>r2 c2</td> </tr>
			</table>`,
			"table_0", 3, 3, true, false},
		{ // test7 Empty first row. Progressively more cols in additional rows.
			`<table>
			<tr>                                              </tr>
			<tr> <td>r0 c0</td> <td>r0 c1</td> <td>r0 c2</td> </tr>
			<tr> <td>test7</td> <td>r1 c1</td> <td>r1 c2</td> </tr>
			<tr> <td>r2 c0</td> <td>r2 c1</td> <td>r2 c2</td> </tr>
			</table>`,
			"table_0", 3, 3, true, true},
		{ // test8 Add more than one col in subsequent row.
			`<table>
			<tr> <td>r0 c0</td>                               </tr>
			<tr> <td>test8</td> <td>r1 c1</td> <td>r1 c2</td> </tr>
			<tr> <td>r2 c0</td> <td>r2 c1</td> <td>r2 c2</td> </tr>
			</table>`,
			"table_0", 3, 3, true, false},
	}

	for i, test := range tests {

		sourceTag := fmt.Sprintf("tests[%d].tableHtml", i)

		goTableSet, err := NewGoTableSetFromHtmlString(sourceTag, test.tableHtml)
		if err != nil {
			t.Fatal(err)
		}

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

		var expecting string

		if test.testCell {
			// Pick a cell with a predictable value and see if it's correct.
			row0col1, err := table.GetStringByColIndex(1, 0)
			if err != nil {
				t.Errorf("test[%d]: %v", i, err)
			}
			expecting = "r0 c1"
			if row0col1 != expecting {
				t.Errorf("test[%d]: Cell contents wrong %q (expecting %q)", i, row0col1, expecting)
			}
		}

		// test1 etc are intended to help identify the erroneous table.
		testCellValue, err := table.GetStringByColIndex(0, 1)
		if err != nil {
			t.Errorf("test[%d]: %v", i, err)
		}
		expecting = "test" + fmt.Sprintf("%d", i)
		if testCellValue != expecting {
			t.Errorf("test[%d]: Cell contents wrong %q (expecting %q)", i, testCellValue, expecting)
		}

		fmt.Printf("\nRESULT %s\n%s\n", sourceTag, table)
	}
}

//func TestReadString2(t *testing.T) {
//	_, err := ReadString(
//		`[EmptyTable1]
//
//		# Should be a syntax error. Table should have both names AND types.
//		[TableWithColNamesOnly]
//		A	B	C
//
//		[EmptyTable2]
//	`)
//	if err == nil {
//		t.Errorf("Should return a syntax error. Table should have both names AND types.")
//	}
//}
//
//func TestReadString3(t *testing.T) {
//	_, err := ReadString(
//		`[TableWithRow]
//		D	E	F
//		int	int	int
//		1	2	3
//
//		A	B	C
//	`)
//	if err == nil {
//		t.Errorf("Should return a syntax error. Col names should not follow blank lines.")
//	}
//}
//
//func TestReadString4(t *testing.T) {
//	_, err := ReadString(
//		`[TableWithRow]
//		D	E	F
//		int	int	int
//		1	2	3
//
//		4	5	6
//		`)
//	if err == nil {
//		t.Errorf("Should return a syntax error. Col values should not follow blank lines.")
//	}
//}
//
//func TestReadString5(t *testing.T) {
//	goTableSet, err := ReadString(
//		`[TableEmpty]
//
//		`)
//	if err != nil {
//		t.Error(err)
//	}
//
//	tests := []struct {
//		tableName string
//		colCount int
//		rowCount int
//		succeeds bool
//	}{
//		{ "TableEmpty", 0, 0, true },
//	}
//
//	for i, test := range tests {
//		table, err := goTableSet.Table(test.tableName)
//		if err != nil {
//			t.Errorf("[%d] %v", i, err)
//		}
//		if (err == nil) != test.succeeds {
//			t.Errorf("test[%d]: err == %v but expecting succeeds == %t", i, err, test.succeeds)
//		}
//
//		colCount := table.ColCount()
//		if colCount != test.colCount {
//			tableName := table.TableName()
//			t.Errorf("test[%d]: Expecting [%s] colCount %d, not %d\n", i, tableName, test.colCount, colCount)
//		}
//
//		rowCount := table.RowCount()
//		if rowCount != test.rowCount {
//			tableName := table.TableName()
//			t.Errorf("test[%d]: Expecting [%s] rowCount %d, not %d\n", i, tableName, test.rowCount, rowCount)
//		}
//	}
//}
//
//func TestReadString6(t *testing.T) {
//	goTableSet, err := ReadString(
//		`[TableStruct]
//		i int = 42
//		j int = 44
//
//		[Empty]
//
//		`)
//	if err != nil {
//		t.Error(err)
//	}
//
//	tests := []struct {
//		tableName string
//		colCount int
//		rowCount int
//		succeeds bool
//	}{
//		{ "TableStruct", 2, 1, true },
//	}
//
//	for i, test := range tests {
//		table, err := goTableSet.Table(test.tableName)
//		if err != nil {
//			t.Errorf("[%d] %v", i, err)
//		}
//		if (err == nil) != test.succeeds {
//			t.Errorf("test[%d]: err == %v but expecting succeeds == %t", i, err, test.succeeds)
//		}
//
//		colCount := table.ColCount()
//		if colCount != test.colCount {
//			tableName := table.TableName()
//			t.Errorf("test[%d]: Expecting [%s] colCount %d, not %d\n", i, tableName, test.colCount, colCount)
//		}
//
//		rowCount := table.RowCount()
//		if rowCount != test.rowCount {
//			tableName := table.TableName()
//			t.Errorf("test[%d]: Expecting [%s] rowCount %d, not %d\n", i, tableName, test.rowCount, rowCount)
//		}
//	}
//}
//
//func TestReadString7(t *testing.T) {
//	_, err := ReadString(
//		`[TableStruct]
//		i int = 42
//		j int = 44
//		# Expecting more structs or a blank line.
//		X Y Z
//
//		[Empty]
//
//		`)
//	if err == nil {
//		t.Error(err)
//	}
//}
//
//func TestReadString8(t *testing.T) {
//	_, err := ReadString(
//		`[TableShaped]
//		X Y Z
//		# Expecting col types, not structs.
//		i int = 42
//		j int = 44
//
//		[Empty]
//
//		`)
//	if err == nil {
//		t.Error(err)
//	}
//}
//
//func ExampleNewGoTableSetFromHtml() {
//	url := "http://www.bom.gov.au/vic/observations/melbourne.shtml"
//	var tableSet *GoTableSet
//	tableSet, err := NewGoTableSetFromHtml(url)
//	if err != nil {
//		panic(err)
//	}
//	tableCount := tableSet.TableCount()
//	fmt.Println(tableCount)
//	// Output:
//	// 1
//}
//
//func ExampleNewGoTableSet() {
//	tableSetName := "MyTableSet"
//	tableSet, err := NewGoTableSet(tableSetName)
//	if err != nil {
//		panic(err)
//	}
//	tableCount := tableSet.TableCount()
//	name := tableSet.GoTableSetName()
//	fmt.Println(tableCount)
//	fmt.Println(name)
//	// Output:
//	// 0
//	// MyTableSet
//}
//
//func ExampleRound() {
//	numberToRound := 12.326
//	places := 2		// The rounded fractional part will have 2 decimal places.
//
//	rounded := Round(numberToRound, places)
//	fmt.Println(rounded)
//	// Output:
//	// 12.33
//}
//
//func TestRound(t *testing.T) {
//	tests := []struct {
//		val     float64
//		places  int
//		rounded float64
//	}{
//		{ 12.326, 2, 12.33 },
//		{ 12.325, 2, 12.33 },
//		{ 12.324, 2, 12.32 },
//		{ 12.32,  2, 12.32 },
//		{ 12.3,   2, 12.3  },
//	}
//
//	for i, test := range tests {
//		rounded := Round(test.val, test.places)
//		if rounded != test.rounded {
//			t.Errorf("test[%d]: Expecting %f, not %f", i, test.rounded, rounded)
//		}
//	}
//}
//
//func TestSetAndGetFunctions(t *testing.T) {
//	var bVal   bool
////	byte 			// alias for uint8
////	complex128 		// The set of all complex numbers with float64 real and imaginary parts
////	complex64		// The set of all complex numbers with float32 real and imaginary parts
///*
//	var f32Val float32	// The set of all IEEE-754 32-bit floating-point numbers
//	var f64Val float64	// The set of all IEEE-754 64-bit floating-point numbers
//	var iVal   int		// Machine-dependent
//	var i16Val int16	// The set of all signed 16-bit integers (-32768 to 32767)
//	var i32Val int32	// The set of all signed 32-bit integers (-2147483648 to 2147483647)
//	var i64Val int64	// The set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
//	var i8Val  int8		// The set of all signed  8-bit integers (-128 to 127)
////	rune 			// alias for int32
//	var sVal    string
//*/
//	var uiVal   uint	// Machine-dependent
///*
//	var ui16Val uint16	// The set of all unsigned 16-bit integers (0 to 65535)
//	var ui32Val uint32	// The set of all unsigned 32-bit integers (0 to 4294967295)
//	var ui64Val uint64	// The set of all unsigned 64-bit integers (0 to 18446744073709551615)
//	var ui8Val  uint8	// The set of all unsigned  8-bit integers (0 to 255)
//*/
//
//	var err error
//	var table *GoTable
//	const rowIndex = 0
//	var colIndex = 0
//
//	table, err = NewGoTable("SetAndGet")
//	if err != nil {
//		t.Error(err)
//	}
//
//	table.AddRow()
//
//
//	// Bool tests
//
//	err = table.AddCol("bVal", "bool")
//	if err != nil {
//		t.Error(err)
//	}
//
//	err = table.SetBool("bVal", rowIndex, true)
//	if err != nil {
//		t.Error(err)
//	}
//	bVal, err = table.GetBool("bVal", rowIndex)
//	if err != nil {
//		t.Error(err)
//	}
//	if bVal != true {
//		t.Errorf("Expecting GetBool() value %t, not %t\n", true, bVal)
//	}
//
//	err = table.SetBoolByColIndex(colIndex, rowIndex, false)
//	if err != nil {
//		t.Error(err)
//	}
//	bVal, err = table.GetBoolByColIndex(colIndex, rowIndex)
//	if err != nil {
//		t.Error(err)
//	}
//	if bVal != false {
//		t.Errorf("Expecting GetBoolByColIndex() value %t, not %t\n", true, bVal)
//	}
//
//
//	// Uint tests
//
//	err = table.AddCol("uiVal", "uint")
//	if err != nil {
//		t.Error(err)
//	}
//	colIndex += 1
//
//	err = table.SetUint("uiVal", rowIndex, 55)
//	if err != nil {
//		t.Error(err)
//	}
//	uiVal, err = table.GetUint("uiVal", rowIndex)
//	if err != nil {
//		t.Error(err)
//	}
//	if uiVal != 55 {
//		t.Errorf("Expecting GetUint() value %d, not %d\n", 55, uiVal)
//	}
//
//	err = table.SetUintByColIndex(colIndex, rowIndex, 66)
//	if err != nil {
//		t.Error(err)
//	}
//	uiVal, err = table.GetUintByColIndex(colIndex, rowIndex)
//	if err != nil {
//		t.Error(err)
//	}
//	if uiVal != 66 {
//		t.Errorf("Expecting GetUintByColIndex() value %d, not %d\n", 66, uiVal)
//	}
//}
//
//func TestSetIntegerMinAndMax(t *testing.T) {
//	var err error
//
//	// For testing machine-dependent types
//	var intBits int = strconv.IntSize	// uint and int are the same bit size.
//	var intMinVal int64
//	var intMaxVal uint64
//	var uintMaxVal uint64
//	switch intBits {
//		case 32:
//			intMinVal = math.MinInt32
//			intMaxVal = math.MaxInt32
//			uintMaxVal = math.MaxUint32
//		case 64:
//			intMinVal = math.MinInt64
//			intMaxVal = math.MaxInt64
//			uintMaxVal = math.MaxUint64
//		default:
//			msg := fmt.Sprintf("CHECK int or uint ON THIS SYSTEM: Unknown int size: %d bits", intBits)
//			t.Errorf(msg)
//	}
//
//	var tests = []struct {
//		input string
//		valid bool
//	}{
//
//		// uint8
//		{`
//			[uint8]
//			i uint8 = 0`,
//			true,
//		},
//		{`
//			[uint8]
//			i uint8 = 255`,
//			true,
//		},
//		{`
//			[uint8]
//			i uint8 = -1`,
//			false,
//		},
//		{`
//			[uint8]
//			i uint8 = 256`,
//			false,
//		},
//
//		// uint16
//		{`
//			[uint16]
//			i uint16 = 0`,
//			true,
//		},
//		{`
//			[uint16]
//			i uint16 = 65535`,
//			true,
//		},
//		{`
//			[uint16]
//			i uint16 = -1`,
//			false,
//		},
//		{`
//			[uint16]
//			i uint16 = 65536`,
//			false,
//		},
//
//		// uint32
//		{`
//			[uint32]
//			i uint32 = 0`,
//			true,
//		},
//		{`
//			[uint32]
//			i uint32 = 4294967295`,
//			true,
//		},
//		{`
//			[uint32]
//			i uint32 = -1`,
//			false,
//		},
//		{`
//			[uint32]
//			i uint32 = 4294967296`,
//			false,
//		},
//
//		// uint64
//		{`
//			[uint64]
//			i uint64 = 0`,
//			true,
//		},
//		{`
//			[uint64]
//			i uint64 = 18446744073709551615`,
//			true,
//		},
//		{`
//			[uint64]
//			i uint64 = -1`,
//			false,
//		},
//		{`
//			[uint64]
//			i uint64 = 18446744073709551616`,
//			false,
//		},
//
//		// uint
//		{`
//			[uint]
//			i uint = 0`,
//			true,
//		},
//		{fmt.Sprintf(`
//			[uint]
//			i uint = %d`, uintMaxVal),
//			true,
//		},
//		{`
//			[uint]
//			i uint = -1`,
//			false,
//		},
//		// Note: Cannot easily test machine-dependent types outside range values (except for uint 0 and -1), so skipping them.
//
//		// int8
//		{`
//			[int8]
//			i int8 = -128`,
//			true,
//		},
//		{`
//			[int8]
//			i int8 = 127`,
//			true,
//		},
//		{`
//			[int8]
//			i int8 = -129`,
//			false,
//		},
//		{`
//			[int8]
//			i int8 = 128`,
//			false,
//		},
//
//		// int16
//		{`
//			[int16]
//			i int16 = -32768`,
//			true,
//		},
//		{`
//			[int16]
//			i int16 = 32767`,
//			true,
//		},
//		{`
//			[int16]
//			i int16 = -32769`,
//			false,
//		},
//		{`
//			[int16]
//			i int16 = 32768`,
//			false,
//		},
//
//		// int32
//		{`
//			[int32]
//			i int32 = -2147483648`,
//			true,
//		},
//		{`
//			[int32]
//			i int32 = 2147483647`,
//			true,
//		},
//		{`
//			[int32]
//			i int32 = -2147483649`,
//			false,
//		},
//		{`
//			[int32]
//			i int32 = 2147483648`,
//			false,
//		},
//
//		// int64
//		{`
//			[int64]
//			i int64 = -9223372036854775808`,
//			true,
//		},
//		{`
//			[int64]
//			i int64 = 9223372036854775807`,
//			true,
//		},
//		{`
//			[int64]
//			i int64 = -9223372036854775809`,
//			false,
//		},
//		{`
//			[int64]
//			i int64 = 9223372036854775808`,
//			false,
//		},
//
//		// int
//		{fmt.Sprintf(`
//			[int]
//			i int = %d`, intMinVal),
//			true,
//		},
//		{fmt.Sprintf(`
//			[int]
//			i int = %d`, intMaxVal),
//			true,
//		},
//		// Note: Cannot easily test machine-dependent types outside range values here, so skipping them.
//		//       See TestSetIntegerMinAndMaxMachineDependent()
//	}
//
//	for i, test := range tests {
//		_, err = NewGoTableSetFromString(test.input)
//		if err == nil != test.valid {
//			switch test.valid {
//				case true:
//					t.Errorf("test[%d]: %v", i, err)
//				case false:
//					t.Errorf("test[%d]: NewGoTableSetFromString(): Expecting this input to fail with a range error: %s", i, test.input)
//			}
//		}
//	}
//}
//
//func TestSetIntegerMinAndMaxMachineDependent(t *testing.T) {
//	var err error
//
//	type testCase struct {
//		input string
//		valid bool
//	}
//	var tests []testCase
//
//	// All of these tests are of out-of-range values (1 too small or 1 too large) which should fail when parsed.
//	// NOTE: Only half of these tests will be executed. They are machine dependent: 32-bit OR 64-bit machines.
//
//	// For testing machine-dependent types
//	var intBits int = strconv.IntSize	// uint and int are the same bit size.
//	switch intBits {
//		case 32:	// NOTE: This will be executed on 32-bit machines ONLY.
//			tests = append(tests, testCase{`
//					[uint]
//					i uint = 4294967296`,
//					false,
//				},
//			)
//			tests = append(tests, testCase{`
//					[int]
//					i int = -2147483649`,
//					false,
//				},
//			)
//			tests = append(tests, testCase{`
//					[int]
//					i int = 2147483648`,
//					false,
//				},
//			)
//		case 64:	// NOTE: This will be executed on 32-bit machines ONLY.
//			tests = append(tests, testCase{`
//					[uint]
//					i uint = 18446744073709551616`,
//					false,
//				},
//			)
//			tests = append(tests, testCase{`
//					[int]
//					i int = -9223372036854775809`,
//					false,
//				},
//			)
//			tests = append(tests, testCase{`
//					[int]
//					i int = 9223372036854775808`,
//					false,
//				},
//			)
//		default:
//			msg := fmt.Sprintf("CHECK int or uint ON THIS SYSTEM: Unknown int size: %d bits", intBits)
//			t.Errorf(msg)
//	}
//
//	for i, test := range tests {
//		_, err = NewGoTableSetFromString(test.input)
//		if err == nil != test.valid {
//			switch test.valid {
//				case true:
//					t.Errorf("test[%d]: %v", i, err)
//				case false:
//					t.Errorf("test[%d]: NewGoTableSetFromString(): Expecting this input to fail with a range error: %s", i, test.input)
//			}
//		}
//	}
//}
//
//var goTableSetString string = `
//	[sable_fur]
//    i   s       f           b
//    int string  float64     bool
//    1   "abc"   2.3         true
//    2   "xyz"   4.5         false
//    3   "ssss"  4.9         false
//
//    [my_struct_table]
//    i int    = 9223372036854775807
//    i2 int64 = 9223372036854775807
//    s string = "forty-two"
//    f int8 = 42
//    u uint8  = 255
//    i81 int8 = 127
//    i82 int8 = -128
//    i161 int16 = 32767
//    i162 int16 = -32768
//    i321 int8 = 127
//    i322 int8 = -128
//    ui uint16 = 65535
//    `
//
//func BenchmarkNewGoTableSetFromString(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		NewGoTableSetFromString(goTableSetString)
//	}
//}
//
//func BenchmarkGoTableSetToString(b *testing.B) {
//	// Set up for benchmark.
//	goTableSet, err := NewGoTableSetFromString(goTableSetString)
//	if err != nil {
//		b.Error(err)
//	}
//
//	for i := 0; i < b.N; i++ {
//		goTableSet.String()
//	}
//}
//
//func BenchmarkGobEncode(b *testing.B) {
//	// Set up for benchmark.
//	goTableSet, err := NewGoTableSetFromString(goTableSetString)
//	if err != nil {
//		b.Error(err)
//	}
//
//	for i := 0; i < b.N; i++ {
//		_, err := goTableSet.GobEncode()
//		if err != nil {
//			b.Error(err)
//		}
//	}
//}
//
//func BenchmarkGobDecode(b *testing.B) {
//	// Set up for benchmark.
//	goTableSet, err := NewGoTableSetFromString(goTableSetString)
//	if err != nil {
//		b.Error(err)
//	}
//
//	// Set up for benchmark.
//	gobEncodedTableSet, err := goTableSet.GobEncode()
//	if err != nil {
//		b.Error(err)
//	}
//
//	for i := 0; i < b.N; i++ {
//		_, err := GobDecodeTableSet(gobEncodedTableSet)
//		if err != nil {
//			b.Error(err)
//		}
//	}
//}
