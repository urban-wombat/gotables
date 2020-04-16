package gotables_test

// Note: This is a black box test (different package name: not gotables).

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/urban-wombat/gotables"
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

var where = log.Print

var typesMap = map[int]string {
	 0: "bool",
	 1: "byte",
	 2: "float32",
	 3: "float64",
	 4: "int",
	 5: "int16",
	 6: "int32",
	 7: "rune",
	 8: "int64",
	 9: "int8",
	10: "string",
	11: "uint",
	12: "uint16",
	13: "uint32",
	14: "uint64",
	15: "uint8",
	16: "[]byte",
	17: "[]uint8",
	18: "*Table",
	19: "time.Time",
}

func TestNewTableFromString_random(t *testing.T) {

	const deterministic bool = true

	var err error
	var table *gotables.Table

	var random *rand.Rand
	if deterministic {
		random = rand.New(rand.NewSource(0))
	} else {
		// Make it non-deterministic.
		random = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	var r *rand.Rand = rand.New(random)

	const testCount int = 10
	const MaxCols int = 20
	const MaxRows int = 10

	for testIndex := 0; testIndex < testCount; testIndex++ {

		tableName := fmt.Sprintf("t%d", testIndex)
		// where(tableName)
		table, err = gotables.NewTable(tableName)
		if err != nil {
			t.Fatal(err)
		}
		_, err = table.IsValidTable()
		if err != nil {
			t.Fatal(err)
		}

		colCount := r.Intn(MaxCols)
		// fmt.Printf("colCount = %d\n", colCount)
		for colIndex := 0; colIndex < colCount; colIndex++ {
			colName := fmt.Sprintf("c%d", colIndex)	
			colTypeIndex := r.Intn(len(typesMap))
			colType := typesMap[colTypeIndex]
			err = table.AppendCol(colName, colType)
			if err != nil {
				t.Fatal(err)
			}
			_, err = table.IsValidTable()
			if err != nil {
				t.Fatal(err)
			}
		}

		rowCount := r.Intn(MaxRows)
		// fmt.Printf("rowCount = %d\n", rowCount)
		err = table.AppendRows(rowCount)
		if err != nil {
			t.Fatal(err)
		}
		_, err = table.IsValidTable()
		if err != nil {
			t.Fatal(err)
		}

		for colIndex := 0; colIndex < colCount; colIndex++ {
			err = table.SetColCellsToZeroValueByColIndex(colIndex)
			if err != nil {
				t.Fatal(err)
			}
			_, err = table.IsValidTable()
			if err != nil {
				t.Fatal(err)
			}
		}

		_, err = table.IsValidTable()
		if err != nil {
			t.Fatal(err)
		}
		// fmt.Println(table.String())

		_, err = gotables.NewTableFromString(table.String())
		if err != nil {
			t.Fatal(err)
		}
	}

	/*
	//	Return a number within a range.
	min := 10
	max := 30
	fmt.Println(r.Intn(max - min) + min)
	*/
}
