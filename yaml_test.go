package gotables_test

import (
	"os"
	_"fmt"
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
	var yamlString string

	tableSetString = `
	[[TipTopName]]

	[Tminus1]
	f32 float32 = 27
	f64 float64 = 3.402823e+38
	u8 uint8 = 99
	u16 uint16 = 116
	u32 uint32 = 500
	u64 uint64 = 900
	i int = 9223372036854775807
	i8 int8 = -128
	i16 int16 = -32768
	i32 int32 = 66
	i64 int64 = 900
	s string = "something"
	bo bool = true
	r rune = 'F'
	bt byte = 65
	bta []byte = [1 2 3]
	u8a []uint8 = [4 5 6]
	t time.Time = 2020-03-15T14:22:30.123456789+17:00

	[T0]
	f		u		c		k
	float64	uint16	rune	int
	11.1	2		'a'		3
	22.2	4		'b'		4
	33.3	6		'c'		5

	[T1]
	a int = 1
	y int = 4
	s []byte = [4 3 2 1]
	u []uint8 = [42 44 48 50 52]
	t time.Time = 2020-03-15T14:22:30.123456789+17:00
	Y float32 = 66.666

	[T2]
	x		y		s
	bool	byte	string
	true	42		"forty two"
	false	55		"fifty-five"
	`

	tableSetString = `
	[[TwoTables]]

	[T2]
	x		y		s				sss
	bool	byte	string			string
	true	42		"forty two"		"sss0"
	false	55		"fifty-five"	"sss1"
	`
/*
	[T3]
	t *Table = []
*/
	tableSet1, err = gotables.NewTableSetFromString(tableSetString)
	if err != nil {
		t.Fatal(err)
	}
where("tableSet1.String()\n" + tableSet1.String())
println()

//	empty *Table = [EmptyTable]
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
where("nestedTable.String()\n" + nestedTable.String())
println()

/*
	t3, err := tableSet1.GetTable("T3")
	if err != nil {
		t.Fatal(err)
	}
*/
//where("tableSet1.String()\n" + tableSet1.String())
//println()
/*
	err = t3.SetTable("t", 0, nestedTable)
	if err != nil {
		t.Fatal(err)
	}
*/
//where(tableSet1.TableCount())
//where("tableSet1.String()\n" + tableSet1.String())
println()

	yamlString, err = tableSet1.GetTableSetAsYAML()
	if err != nil {
		t.Fatal(err)
	}
where("tableSet1.String()\n" + tableSet1.String())
where("yamlString\n" + yamlString)
println()
//os.Exit(55)
_, _ = os.Hostname()

	var yamlMap map[string]interface{}
	yamlMap, err = tableSet1.GetTableSetAsMap()
	if err != nil {
		t.Fatal(err)
	}
// DATA PRESENT
//println()
//where(yamlMap)
_ = yamlMap

	tableSet2, err = gotables.NewTableSetFromYAML(yamlString)
	if err != nil {
where()
		t.Fatal(err)
	}
// DATA MISSING
where("DATA MISSING")
where(tableSet2.TableCount())
where("tableSet2.String()\n" + tableSet2.String())
println()

jsonStr1, err := tableSet1.GetTableSetAsJSONIndent()
if err != nil {
	t.Fatal(err)
}
where("tableSet1 jsonStr1:\n" + jsonStr1)
println()

jsonStr2, err := tableSet2.GetTableSetAsJSONIndent()
if err != nil {
	t.Fatal(err)
}
where("tableSet2 jsonStr2:\n" + jsonStr2)
println()

	equals, err := tableSet1.Equals(tableSet2)
	if err != nil {
where("\ntableSet1:\n------------\n" + tableSet1.String() + "------------\n")
where("\ntableSet2:\n------------\n" + tableSet2.String() + "------------\n")
where(err)
//os.Exit(22)
		t.Fatal(err)
	}
where(equals)
}
