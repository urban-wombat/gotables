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

// Note: Leading lowercase in 'cellTableInStruct' is required for it to be recognised as an Example!

func ExampleNewTableFromString_cellTableInStruct() {
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

	// Print the table in its original struct shape.
	fmt.Println(table)

	// Now change its shape to tabular.
	err = table.SetStructShape(false)
	if err != nil {
		log.Println(err)
	}

	// The table is now printed as a single row of data.
	fmt.Println(table)

	// Note: The struct/tabular shape is for readability and has no impact on its internal structure.

	// Output:
	// [MyTable]
	// MyBool bool = true
	// MyString string = "The answer to life, the universe and everything."
	// MyInt int = 42
	// MyTable *Table = [CellTable]
	// MyTable2 *gotables.Table = [CellTable2]
	//
	// [MyTable]
	// MyBool MyString                                           MyInt MyTable     MyTable2
	// bool   string                                               int *Table      *gotables.Table
	// true   "The answer to life, the universe and everything."    42 [CellTable] [CellTable2]
}

// Note: Leading lowercase in table is required for it to be recognised as an Example!

func ExampleNewTableFromString_cellTableInStructSetToNil() {
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
	err = table.SetVal("MyTable", 0, nilTable)
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
