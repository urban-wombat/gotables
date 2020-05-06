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
	var visitTable = func(table *gotables.Table) (err error) {
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

				lastRowIndex := table.RowCount() - 1

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

	var visitCell = func(cell gotables.Cell) (err error) {
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

	err = table.Walk(visitTable, nil, visitCell)
	if err != nil {
		panic(err)
	}
}

func ExampleTable_Walk() {

	// Define the visitTable function.

	var visitTable = func(table *gotables.Table) (err error) {

		fmt.Printf("[%s].visitTable()\n", table.Name())

		fmt.Printf("table.Name() = %s\n", table.Name())
		fmt.Printf("table.ColCount() = %d\n", table.ColCount())
		fmt.Printf("table.RowCount() = %d\n", table.RowCount())
		fmt.Printf("table.String():n%s", table.String())

		if table.ParentTable() != nil { // I.e., this is a child table.
			fmt.Printf("table.ParentTable.Name() = %s\n", table.ParentTable().Name())
		}

		if table.Name() == "RootTable" {
			// Top-level table which we know has a "nested" *Table col.
			for i := 0; i < 3; i++ {
				err = table.AppendRow()
				if err != nil {
					return err
				}

				// Here we are avoiding the gotcha of a nilTable (no name) being unusable.
				// We need to at least give it a name.
				lastRowIndex := table.RowCount() - 1
				nested, err := table.GetTable("nested", lastRowIndex)
				if err != nil {
					return err
				}
				err = nested.SetName("AnyNameYouLike")
				if err != nil {
					return err
				}
			}
		}

		fmt.Printf("\n%s\n", table)

		return
	}

	// Define the visitCell function.

	var visitCell = func(cell gotables.Cell) (err error) {

		fmt.Printf("[%s].visitCell(colName=%s, colIndex=%d, rowIndex=%d)\n",
			cell.Table.Name(), cell.ColName, cell.ColIndex, cell.RowIndex)

		fmt.Printf("cell.Table.Name() = %s\n", cell.Table.Name())
		fmt.Printf("cell.ColName = %s\n", cell.ColName)
		fmt.Printf("cell.ColIndex = %d\n", cell.ColIndex)
		fmt.Printf("cell.RowIndex = %d\n", cell.RowIndex)
		fmt.Printf("cell.ColType = %s\n", cell.ColType)
		fmt.Println()

		return
	}

	tableString := `
	[RootTable]
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

	nestedTableString := `
	[Nested]
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

	err = table.Walk(visitTable, nil, visitCell)
	if err != nil {
		panic(err)
	}

	// Output:
	// [RootTable].visitTable()
	// table.Name() = RootTable
	// table.ColCount() = 4
	// table.RowCount() = 1
	// table.String():n[RootTable]
	// i int = 22
	// j int = 33
	// k int = 44
	// nested *Table = [Nested]
	//
	// [RootTable]
	//   i   j   k nested
	// int int int *Table
	//  22  33  44 [Nested]
	//   0   0   0 [AnyNameYouLike]
	//   0   0   0 [AnyNameYouLike]
	//   0   0   0 [AnyNameYouLike]
	//
	// [RootTable].visitCell(colName=i, colIndex=0, rowIndex=0)
	// cell.Table.Name() = RootTable
	// cell.ColName = i
	// cell.ColIndex = 0
	// cell.RowIndex = 0
	// cell.ColType = int
	//
	// [RootTable].visitCell(colName=i, colIndex=0, rowIndex=1)
	// cell.Table.Name() = RootTable
	// cell.ColName = i
	// cell.ColIndex = 0
	// cell.RowIndex = 1
	// cell.ColType = int
	//
	// [RootTable].visitCell(colName=i, colIndex=0, rowIndex=2)
	// cell.Table.Name() = RootTable
	// cell.ColName = i
	// cell.ColIndex = 0
	// cell.RowIndex = 2
	// cell.ColType = int
	//
	// [RootTable].visitCell(colName=i, colIndex=0, rowIndex=3)
	// cell.Table.Name() = RootTable
	// cell.ColName = i
	// cell.ColIndex = 0
	// cell.RowIndex = 3
	// cell.ColType = int
	//
	// [RootTable].visitCell(colName=j, colIndex=1, rowIndex=0)
	// cell.Table.Name() = RootTable
	// cell.ColName = j
	// cell.ColIndex = 1
	// cell.RowIndex = 0
	// cell.ColType = int
	//
	// [RootTable].visitCell(colName=j, colIndex=1, rowIndex=1)
	// cell.Table.Name() = RootTable
	// cell.ColName = j
	// cell.ColIndex = 1
	// cell.RowIndex = 1
	// cell.ColType = int
	//
	// [RootTable].visitCell(colName=j, colIndex=1, rowIndex=2)
	// cell.Table.Name() = RootTable
	// cell.ColName = j
	// cell.ColIndex = 1
	// cell.RowIndex = 2
	// cell.ColType = int
	//
	// [RootTable].visitCell(colName=j, colIndex=1, rowIndex=3)
	// cell.Table.Name() = RootTable
	// cell.ColName = j
	// cell.ColIndex = 1
	// cell.RowIndex = 3
	// cell.ColType = int
	//
	// [RootTable].visitCell(colName=k, colIndex=2, rowIndex=0)
	// cell.Table.Name() = RootTable
	// cell.ColName = k
	// cell.ColIndex = 2
	// cell.RowIndex = 0
	// cell.ColType = int
	//
	// [RootTable].visitCell(colName=k, colIndex=2, rowIndex=1)
	// cell.Table.Name() = RootTable
	// cell.ColName = k
	// cell.ColIndex = 2
	// cell.RowIndex = 1
	// cell.ColType = int
	//
	// [RootTable].visitCell(colName=k, colIndex=2, rowIndex=2)
	// cell.Table.Name() = RootTable
	// cell.ColName = k
	// cell.ColIndex = 2
	// cell.RowIndex = 2
	// cell.ColType = int
	//
	// [RootTable].visitCell(colName=k, colIndex=2, rowIndex=3)
	// cell.Table.Name() = RootTable
	// cell.ColName = k
	// cell.ColIndex = 2
	// cell.RowIndex = 3
	// cell.ColType = int
	//
	// [RootTable].visitCell(colName=nested, colIndex=3, rowIndex=0)
	// cell.Table.Name() = RootTable
	// cell.ColName = nested
	// cell.ColIndex = 3
	// cell.RowIndex = 0
	// cell.ColType = *Table
	//
	// [Nested].visitTable()
	// table.Name() = Nested
	// table.ColCount() = 2
	// table.RowCount() = 1
	// table.String():n[Nested]
	// t bool = true
	// f bool = false
	// table.ParentTable.Name() = RootTable
	//
	// [Nested]
	// t bool = true
	// f bool = false
	//
	// [Nested].visitCell(colName=t, colIndex=0, rowIndex=0)
	// cell.Table.Name() = Nested
	// cell.ColName = t
	// cell.ColIndex = 0
	// cell.RowIndex = 0
	// cell.ColType = bool
	//
	// [Nested].visitCell(colName=f, colIndex=1, rowIndex=0)
	// cell.Table.Name() = Nested
	// cell.ColName = f
	// cell.ColIndex = 1
	// cell.RowIndex = 0
	// cell.ColType = bool
	//
	// [RootTable].visitCell(colName=nested, colIndex=3, rowIndex=1)
	// cell.Table.Name() = RootTable
	// cell.ColName = nested
	// cell.ColIndex = 3
	// cell.RowIndex = 1
	// cell.ColType = *Table
	//
	// [AnyNameYouLike].visitTable()
	// table.Name() = AnyNameYouLike
	// table.ColCount() = 0
	// table.RowCount() = 0
	// table.String():n[AnyNameYouLike]
	// table.ParentTable.Name() = RootTable
	//
	// [AnyNameYouLike]
	//
	// [RootTable].visitCell(colName=nested, colIndex=3, rowIndex=2)
	// cell.Table.Name() = RootTable
	// cell.ColName = nested
	// cell.ColIndex = 3
	// cell.RowIndex = 2
	// cell.ColType = *Table
	//
	// [AnyNameYouLike].visitTable()
	// table.Name() = AnyNameYouLike
	// table.ColCount() = 0
	// table.RowCount() = 0
	// table.String():n[AnyNameYouLike]
	// table.ParentTable.Name() = RootTable
	//
	// [AnyNameYouLike]
	//
	// [RootTable].visitCell(colName=nested, colIndex=3, rowIndex=3)
	// cell.Table.Name() = RootTable
	// cell.ColName = nested
	// cell.ColIndex = 3
	// cell.RowIndex = 3
	// cell.ColType = *Table
	//
	// [AnyNameYouLike].visitTable()
	// table.Name() = AnyNameYouLike
	// table.ColCount() = 0
	// table.RowCount() = 0
	// table.String():n[AnyNameYouLike]
	// table.ParentTable.Name() = RootTable
	//
	// [AnyNameYouLike]
}
