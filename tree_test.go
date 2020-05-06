package gotables_test

import (
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

func TestTable_Walk_countInts(t *testing.T) {
	var err error
	var table1 *gotables.Table

	var tableString string
	tableString = `
	[TypesGalore22]
    i   s      right
    int string *Table
    0   "abc"  []
    1   "xyz"  []
    2   "ssss" []
    3   "xxxx" []
    4   "yyyy" []
    `
	table1, err = gotables.NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	// Now create and set some table cell tables.
	right0 := `
	[right0]
	i int = 32`

	right1 := `
	[right1]
	s string = "thirty-two"`

	right2 := `
	[right2]
	x	y	z
	int	int	int
	1	2	3
	4	5	6
	7	8	9`

	right3 := `
	[right3]
	f float32 = 88.8`

	right4 := `
	[right4]
	t1 *Table = []`

	table1.SetTableMustSet("right", 0, gotables.NewTableFromStringMustMake(right0))
	table1.SetTableMustSet("right", 1, gotables.NewTableFromStringMustMake(right1))
	table1.SetTableMustSet("right", 2, gotables.NewTableFromStringMustMake(right2))
	table1.SetTableMustSet("right", 3, gotables.NewTableFromStringMustMake(right3))
	table1.SetTableMustSet("right", 4, gotables.NewTableFromStringMustMake(right4))

	// fmt.Printf("table1:\n%s\n", table1)

	var tableCount int
	var visitTable = func(table *gotables.Table) (err error) {
		// fmt.Printf("visiting:\n%s\n", table)
		tableCount++
		return
	}

	var cellCount int
	var intCount int
	var intSum int
	var float32Count int
	var visitCell = func(cell gotables.Cell) (err error) {
		cellCount++
		if cell.ColType == "int" {
			intCount++
			intSum += cell.Table.GetIntByColIndexMustGet(cell.ColIndex, cell.RowIndex)
		}
		if cell.ColType == "float32" {
			float32Count++
		}
		return
	}

	err = table1.Walk(visitTable, nil, visitCell)
	if err != nil {
		t.Fatal(err)
	}

	var expecting int = 7

	expecting = 7
	if tableCount != expecting {
		t.Fatalf("expecting tableCount %d, not %d", expecting, tableCount)
	}

	expecting = 28
	if cellCount != expecting {
		t.Fatalf("expecting cellCount %d, not %d", expecting, cellCount)
	}

	expecting = 15
	if intCount != expecting {
		t.Fatalf("expecting intCount %d, not %d", expecting, intCount)
	}

	expecting = 87
	if intSum != expecting {
		t.Fatalf("expecting intSum %d, not %d", expecting, intSum)
	}

	expecting = 1
	if float32Count != expecting {
		t.Fatalf("expecting float32Count %d, not %d", expecting, float32Count)
	}
}
