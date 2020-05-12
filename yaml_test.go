package gotables_test

import (
	_ "fmt"
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

	[ColOrder]
	c0 int = 0
	c1 int = 1
	c2 int = 2
	c3 int = 3
	c4 int = 4
	c5 int = 5
	c6 int = 6
	c7 int = 7
	c8 int = 8
	c9 int = 9

	[T0]
	f float64 = 99.90001
	u uint8 = 8
	c byte = 100

	[T1]
	a int = 1
	b int = 2
	c int = 3
	y int = 4
	s []byte = [4 3 2 1]
	u []uint8 = [42 44 48 50 52]
	t time.Time = 2020-03-15T14:22:30.123456789+17:00
	Y float32 = 66.666
	e int = 1
	f int = 1
	g int = 1
	h int = 1
	i int = 1
	j int = 1
	k int = 1
	l int = 1
	m int = 1
	n int = 1
	N int = 1
	o int = 1
	p int = 1
	r int = 1
	v int = 1
	w int = 1
	x int = 1
	z int = 1

	[T2]
	x		y		s
	bool	byte	string
	true	42		"forty two"
	false	55		"fifty-five"
	true	66		"sixty six"

	[T3]
	t *Table = []
	`
	tableSet1, err = gotables.NewTableSetFromString(tableSetString)
	if err != nil {
		t.Fatal(err)
	}
where("\n" + tableSet1.String())
println()

	var nestedString string = `
	[NestedTable]
	empty *Table = []
	noByte []byte = []
	noUint8 []uint8 = []
	r rune = '^'
	`
	nestedTable, err := gotables.NewTableFromString(nestedString)
	if err != nil {
		t.Fatal(err)
	}

	t3, err := tableSet1.Table("T3")
	if err != nil {
		t.Fatal(err)
	}

	err = t3.SetTable("t", 0, nestedTable)
	if err != nil {
		t.Fatal(err)
	}

where()
	yamlString, err = tableSet1.GetTableSetAsYAML()
where(err)
	if err != nil {
		t.Fatal(err)
	}
where()
println()
where("\n" + yamlString)

	tableSet2, err = gotables.NewTableSetFromYAML(yamlString)
	if err != nil {
		t.Fatal(err)
	}
where("\n" + tableSet2.String())
	equals, err := tableSet1.Equals(tableSet2)
	if err != nil {
		t.Fatal(err)
	}
where(equals)
}
