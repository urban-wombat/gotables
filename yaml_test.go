package gotables_test

import (
	_ "encoding/binary"
	"fmt"
	"math"
	_ "os"
	"testing"

	"github.com/urban-wombat/gotables"
)

/*
Copyright (c) 2018 Malcolm Gorman

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

func Test_NewTableSetFromYAML(t *testing.T) {

	var err error
	var tableSet1 *gotables.TableSet
	var tableSet2 *gotables.TableSet
	var tableSetString string

	tableSetString = `
	[[TwoTables]]

	[Tminus1]
	maxint64  int64  =  9000000000000000000
	maxuint64 uint64 = 18000000000000000000
	maxint32  int32  =  2147483647
	maxuint32 uint32 = 4294967295
	f32 float32 = 28
	f64 float64 = 3.402823e+38
	bt byte = 65
	u8 uint8 = 99
	u16 uint16 = 116
	u32 uint32 = 500
	u64 uint64 = 900
	iii2 int = 13
	iii3 int = -20
	uInt4 uint = 4294967295
	i8 int8 = -128
	i16 int16 = -32768
	i32 int32 = 66
	i64 int64 = 900
	s string = "something"
	bo bool = true
	r rune = 'A'
	bta []byte = [65 66 67]
	u8a []uint8 = [97 98 99]
	t time.Time = 2020-03-15T14:22:30.123456789+17:00

	[T1]
	a int = 1
	y int = 4
	s []byte = [88 89 90]
	u []uint8 = [120 121 122 123 124]
	Y float32 = 66.666

	[T2]
	x		y		s				sss
	bool	byte	string			string
	true	44		"forty-four"	"sss0"
	false	55		"fifty-five"	"sss1"
	true	66		"sixty-six"		"sss3"

	[T3]
	a rune = 'b'
	t *Table = [WHAT]
	b string = "b"

	[T4]
	x1 bool = true
	x2 string = "true"
	y1 float32 = 1.1
	y2 string = "one-point-one"
	`
	tableSet1, err = gotables.NewTableSetFromString(tableSetString)
	if err != nil {
		t.Fatal(err)
	}

	var nestedString string = `
	[NestedTable]
	noByte []byte = [1 3 5]
	noUint8 []uint8 = [2 4 6]
	runeVal rune = 'A'
	float32Val float32 = 66.6
	`
	nestedTable, err := gotables.NewTableFromString(nestedString)
	if err != nil {
		t.Fatal(err)
	}

	t3, err := tableSet1.GetTable("T3")
	if err != nil {
		t.Fatal(err)
	}

	err = t3.SetTable("t", 0, nestedTable)
	if err != nil {
		t.Fatal(err)
	}

	var yamlString string
	yamlString, err = tableSet1.GetTableSetAsYAML()
	if err != nil {
		t.Fatal(err)
	}

	_, err = tableSet1.GetTableSetAsMap()
	if err != nil {
		t.Fatal(err)
	}

	tableSet2, err = gotables.NewTableSetFromYAML(yamlString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = tableSet1.Equals(tableSet2)
	if err != nil {
		t.Fatal(err)
	}

	//where(math.MaxInt64 == 9223372036854775807)
	{
		//	var i1 int = 9223372036854775807 // MaxInt64
		//	fmt.Printf("i1: %d\n", i1)
		//
		//	var f1 float64 = float64(i1)
		//	fmt.Printf("f1: %0f\n", f1)
		//
		//	var f2 float64 = 9223372036854775807
		//	fmt.Printf("f2: %0f\n", f2)

		//	var i2 int = int(f2)
		//	fmt.Printf("i2: %d\n", i2)
	}
	fmt.Println()
	{
		//	var i1 int64 = 9223372036854775807 // MaxInt64
		//	fmt.Printf("i1: %d\n", i1)
		//
		//	var f1 float64 = float64(i1)
		//	fmt.Printf("f1: %0f\n", f1)
		//
		//	var f2 float64 = 9223372036854775807
		//	fmt.Printf("f2: %0f\n", f2)

		//	var i2 int64 = int64(f2)
		//	fmt.Printf("i2: %d\n", i2)
	}

	/*
	   {
	   	var i1 int64
	   	var i2 int64
	   	var f float64
	   fmt.Println()
	   //where("working ...")
	   //	const start = 17000000000000000
	   	const start = 1152921504600000000
	   	const inc = 10000000
	   	for i1 = start; i1 <= math.MaxInt64; i1 += inc {
	   		f = float64(i1)
	   		i2 = int64(f)
	   		if i2 != i1 {
	   			fmt.Printf("i1 %d != i2 %d\n", i1, i2)
	   			if (i2-inc) == (i1-inc) {
	   				fmt.Printf("%d == %d\n", i1-inc, i2-inc)
	   			} else {
	   				fmt.Println("What th!")
	   			}
	   			os.Exit(43)
	   		}
	   	}
	   }
	*/
	//var maxint int = 9223372036854775807
	//fmt.Printf("%d\n", maxint)
	//fmt.Printf("%b\n", maxint)
	//fmt.Println()
	//
	//maxint = -9223372036854775808
	//fmt.Printf("%d\n", maxint)
	//fmt.Printf("%b\n", maxint)
	//fmt.Println()
	//
	//var f float64 = float64(9223372036854775807)
	//fmt.Printf("%f\n", f)
	//fmt.Printf("%b\n", f)
	//fmt.Println()
	//
	//maxint = 9223372036854775807
	//fmt.Printf("maxint bits: %b\n", maxint)
	//var b []byte = make([]byte, 8)
	//binary.LittleEndian.PutUint64(b, uint64(maxint))
	////where(fmt.Sprintf("%b\n", b))

	{
		var f64 float64

		// int64
		var i64 int64 = 9223372036854775807
		//where(i64)
		f64 = math.Float64frombits(uint64(i64))
		i64 = int64(math.Float64bits(f64))
		//where(i64)

		// uint64
		var ui64 uint64 = 18446744073709551615
		//where(ui64)
		f64 = math.Float64frombits(uint64(ui64))
		ui64 = math.Float64bits(f64)
		//where(ui64)
	}

	// os.Exit(44)

	var jsonString string
	jsonString, err = tableSet1.GetTableSetAsJSONIndent()
	if err != nil {
		t.Fatal(err)
	}
	//where(jsonString)

	tableSet2, err = gotables.NewTableSetFromJSON(jsonString)
	if err != nil {
		t.Fatal(err)
	}
	//where(tableSet2)

	_, err = tableSet1.Equals(tableSet2)
	if err != nil {
		t.Fatal(err)
	}
}
