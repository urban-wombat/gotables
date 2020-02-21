// Copyright (c) 2017 Malcolm Gorman

// Golang tabular data format for configs and channels, with a rich set of helper functions.

// This is to test the gotables table type as a valid column type.

package gotables

import (
	"fmt"
	"log"
)

/*
Copyright (c) 2020 Malcolm Gorman

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

/* Note: Leading lowercase in 'cellTableInStruct' is required for it to be recognised as an Example! */

func ExampleTable_GetTable_cellTableInStruct() {
	// A table literal. Sometimes easier than constructing a table programmatically.
	tableString := `[MyTable]
		MyBool bool = true
		MyString string = "The answer to life, the universe and everything."
		MyInt int = 42
		MyTable *Table = [CellTable]
		MyTable2 *gotables.Table = [CellTable2]
		`
	// Note 1: The only string form of a table cell containing a *Table is its table name in square brackets.
	// Note 2: To get a table cell *Table as a string, first retrieve it to a variable.
	// Note 3: It is parsed into an empty table with the name specified.

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	myTable, err := table.GetTable("MyTable", 0)
	if err != nil {
		log.Println(err)
	}

	err = myTable.AppendRow()
	if err != nil {
		log.Println(err)
	}

	err = myTable.AppendCol("msg", "string")
	if err != nil {
		log.Println(err)
	}

	err = myTable.SetString("msg", 0, "I am in a table in a cell!")
	if err != nil {
		log.Println(err)
	}

	err = myTable.SetStructShape(true)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(myTable)

	// Note: The struct/tabular shape is for readability and has no impact on its internal structure.

	// Output:
	// [MyTable]
	// MyBool bool = true
	// MyString string = "The answer to life, the universe and everything."
	// MyInt int = 42
	// MyTable *Table = [CellTable]
	// MyTable2 *gotables.Table = [CellTable2]
	//
	// [CellTable]
	// msg string = "I am in a table in a cell!"
}

/* Note: Leading lowercase in table is required for it to be recognised as an Example! */

func ExampleTable_SetTable_cellTableInStructSetToNil() {
	// A table literal. Sometimes easier than constructing a table programmatically.
	tableString := `[MyTable]
		MyBool bool = true
		MyString string = "The answer to life, the universe and everything."
		MyInt int = 42
		MyTable *Table = [CellTable]
		MyNilTable *Table = []
		`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	var nilTable *Table = nil
	err = table.SetTable("MyTable", 0, nilTable)
	if err != nil {
		log.Println(err)
	}

	// Print the table with MyTable cell set to nil.
	// MyNilTable will have the !nil value of an empty and unnamed table.
	fmt.Println(table)

	// Print the individual cells.

	// Here the table cell *Table is nil.

	MyTable, err := table.GetTable("MyTable", 0)
	if MyTable == nil {
		fmt.Println("MyTable == nil")
	} else {
		fmt.Println("MyTable != nil")
	}
	fmt.Printf("MyTable: %#v\n", MyTable)

	fmt.Println()

	// Here the table cell *Table is set to a kind of nil *Table table (with no name) that's not actually nil.

	MyNilTable, err := table.GetTable("MyNilTable", 0)
	if MyNilTable == nil {
		fmt.Println("MyNilTable == nil")
	} else {
		fmt.Println("MyNilTable != nil")
	}
	fmt.Printf("MyNilTable: %s", MyNilTable)
	isValidTable, err := MyNilTable.IsValidTable()
	fmt.Printf("MyNilTable.isValidTable() == %t\n", isValidTable)
	fmt.Println(err)

	// Output:
	// [MyTable]
	// MyBool bool = true
	// MyString string = "The answer to life, the universe and everything."
	// MyInt int = 42
	// MyTable *Table = [CellTable]
	// MyNilTable *Table = []
	//
	// [MyTable]
	// MyBool bool = true
	// MyString string = "The answer to life, the universe and everything."
	// MyInt int = 42
	// MyTable *Table = []
	// MyNilTable *Table = []
	//
	// MyTable == nil
	// MyTable: (*gotables.Table)(nil)
	//
	// MyNilTable != nil
	// MyNilTable: []
	// MyNilTable.isValidTable() == false
	// ERROR IsValidTable(): table has no name
}

func ExampleNewNilTable_createAndUse() {

	// We expect this to print a NilTable with syntax: []
	var nilTable *Table = NewNilTable()
	fmt.Println(nilTable)

	// We expect this to be an invalid table
	isValid, err := nilTable.IsValidTable()
	fmt.Printf("isValid = %t\n", isValid)
	fmt.Printf("err = %v\n", err)

	// We expect any Set, Get or Append operation on the table to fail
	err = nilTable.AppendRow()
	fmt.Printf("err = %v\n", err)
	err = nilTable.AppendCol("my_col", "string")
	fmt.Printf("err = %v\n", err)

	// This fails
	err = nilTable.SetString("missing_col", 0, "my_string_value")
	fmt.Printf("err = %v\n", err)
	missing_col, err := nilTable.GetInt("missing_col", 0)
	fmt.Printf("missing_col = %d\n", missing_col)
	fmt.Printf("err = %v\n", err)

	// Okay, now we will un-NilTable the NilTable
	err = nilTable.SetName("NilTableNoLonger")
	if err != nil {
		log.Println(err)
	}
	// Expecting: [NilTableNoLonger]
	fmt.Println(nilTable)

	// Now we can add a row and a col and set the cell value

	err = nilTable.AppendCol("my_col", "string")
	if err != nil {
		log.Println(err)
	}

	err = nilTable.AppendRow()
	if err != nil {
		log.Println(err)
	}

	err = nilTable.SetString("my_col", 0, "my_string_value")
	if err != nil {
		log.Println(err)
	}

	nilTable.SetStructShape(true)
	fmt.Println(nilTable)

	// Output:
	// []
	//
	// isValid = false
	// err = ERROR IsValidTable(): table has no name
	// err = table.AppendRow(): table is an unnamed NilTable. Call table.SetName() to un-Nil it
	// err = table.AppendCol(): table is an unnamed NilTable. Call table.SetName() to un-Nil it
	// err = table [] col does not exist: missing_col
	// missing_col = 0
	// err = table [] col does not exist: missing_col
	// [NilTableNoLonger]
	//
	// [NilTableNoLonger]
	// my_col string = "my_string_value"
}
