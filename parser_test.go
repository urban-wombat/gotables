package gotables_test

// Note: This is a black box test (different package name: not gotables).

// Also, it's not REALLY a parser test. I thought it was going to be.

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

var typesMap = map[int]string{
	0:  "bool",
	1:  "byte",
	2:  "float32",
	3:  "float64",
	4:  "int",
	5:  "int16",
	6:  "int32",
	7:  "rune",
	8:  "int64",
	9:  "int8",
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

	// Set false to test testCount random tables with each test.
	const deterministic bool = false

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

	const testCount int = 50
	const MaxCols int = 15
	const MaxRows int = 5

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

		colCount := r.Intn(MaxCols + 1)
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

		rowCount := r.Intn(MaxRows + 1)
		// fmt.Printf("rowCount = %d\n", rowCount)
		err = table.AppendRows(rowCount)
		if err != nil {
			t.Fatal(err)
		}
		_, err = table.IsValidTable()
		if err != nil {
			t.Fatal(err)
		}

		// Randomly set to struct shape.
		if rowCount <= 1 && (colCount%2) == 0 {
			err = table.SetStructShape(true)
			if err != nil {
				t.Fatal(err)
			}
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

		// fmt.Println(table.String())
		tableString := table.String()
		_, err = gotables.NewTableFromString(tableString)
		if err != nil {
			where(fmt.Sprintf("\n\n%s\n", tableString))
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

func TestTable_Visit(t *testing.T) {
	var visitTable func(cell *gotables.Table) (err error)
	visitTable = func(table *gotables.Table) (err error) {
	/*
		fmt.Printf("***** table.Name() = %s *****\n", table.Name())
		fmt.Printf("table.ColCount() = %d\n", table.ColCount())
		fmt.Printf("table.RowCount() = %d\n", table.RowCount())
		fmt.Printf("table.String() = \n%s\n", table.String())

		if table.ParentTable() != nil {
			fmt.Printf("*** table.ParentTable.Name() = %s ***\n", table.ParentTable().Name())
		} else {
			fmt.Printf("*** table.ParentTable.Name() = NOTHING ***\n")
		}
	*/

		// Test whether parentTable has been populated.
		if table.Name() == "RootTable" {
			// This table is the root table and must not have a parent.
			if table.ParentTable() != nil {
				t.Fatalf("expecting root-table [%s] parentTable to be nil, but found: %v",
					table.Name(), table.ParentTable())
			}
		} else {
			// This table is a child (nested) table and must have a parent.
			if table.ParentTable() == nil {
				t.Fatalf("expecting NON-root-table [%s] parentTable to be NOT nil, but found: %v",
					table.Name(), table.ParentTable())
			}
		}

		if table.Name() == "RootTable" {
			// Top-level table which we know has a "nested" *Table col.
			for i := 0; i < 3; i++ {
				err = table.AppendRow()
				if err != nil {
					return err
				}
	
				lastRowIndex := table.RowCount()-1
	
				nested, err := table.GetTable("nested", lastRowIndex)
				if err != nil {
					return err
				}
				if nested.ParentTable() == nil {
					t.Fatalf("expecting NON-root-table [%s] parentTable to be NOT nil, but found: %v",
						table.Name(), table.ParentTable())
				}

				err = nested.SetName("AnyNameYouLike")
				if err != nil {
					return err
				}
			}
		}

		//fmt.Printf("\n%s\n", table)
		return
	}

	var visitCell func(cell gotables.Cell) (err error)
	visitCell = func(cell gotables.Cell) (err error) {
	/*
		fmt.Printf("cell.Table.Name() = %s\n", cell.Table.Name())
		fmt.Printf("cell.ColName = %s\n", cell.ColName)
		fmt.Printf("cell.ColIndex = %d\n", cell.ColIndex)
		fmt.Printf("cell.ColType = %s\n", cell.ColType)
	*/
		return
	}

	tableString :=
	`[RootTable]
	i int = 22
	j int = 33
	k int = 44
	nested *Table = []

	[MyName2]
	x		y		z
	float32	float64	uint
	1		3		9
	`
	table, err := gotables.NewTableFromStringByTableIndex(tableString, 0)
	if err != nil {
		panic(err)
	}

	nestedTableString :=
	`[Nested]
	t bool = true
	f bool = false
	`
	nestedTable, err := gotables.NewTableFromString(nestedTableString)
	if err != nil {
		panic(err)
	}

	err = table.SetTable("nested", 0, nestedTable)
	if err != nil {
		panic(err)
	}

	err = table.Walk(visitTable, visitCell)
	if err != nil {
		panic(err)
	}
}
