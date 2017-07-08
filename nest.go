// Copyright (c) 2017 Malcolm Gorman

// Functions and methods for processing Table tables.
package gotable

import (
	"fmt"
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


/*
	Split nestable by keys (which must match table) into separate tables.

	Nest each separate table into the matching (by keys) cell in nestColName col in table.

	Each nested table is given table name nestColName.

	Each nested table is stored as a string. To retrieve it:

		var nestedString string
		var nestedTable *gotable.Table
		nestedString, _ = table.GetString(nestColName, rowIndex)
		nestedTable,  _ = gotable.NewTableFromString(nestedString)
*/
func (table *Table) Nest(nestable *Table, nestColName string) error {

	var err error

	if table == nil {
		return fmt.Errorf("table.%s(nestable, nestColName): table is <nil>", funcName())
	}

	if nestable == nil {
		return fmt.Errorf("table.%s(nestable, nestColName): nestable is <nil>", funcName())
	}

	// Use whichever table has sort keys for the shared keys. table is dominant over nestable.
	if table.SortKeyCount() == 0 {
		err = nestable.SetSortKeysFromTable(table)
		if err != nil {
			return err
		}
	}
	err = nestable.SetSortKeysFromTable(table)
	if err != nil {
		return err
	}

	err = table.Sort()
	if err != nil {
		return err
	}

	err = nestable.Sort()
	if err != nil {
		return err
	}

	keysTable, err := table.GetSortKeysAsTable()
	if err != nil {
		return err
	}

	// Make a slice of cols.
	var cols []string = make([]string, keysTable.RowCount())
	for rowIndex := 0; rowIndex < keysTable.RowCount(); rowIndex++ {
		cols[rowIndex], err = keysTable.GetString("colName", rowIndex)
		if err != nil {
			return err
		}
	}

	// Make an empty slice of keys from table, which are search terms for nestable.
	var keys []interface{} = make([]interface{}, keysTable.RowCount())

	// Create a string col for nested table, for now. Later we may make this a *Table pointer.
	err = table.AppendCol(nestColName, "string")
	if err != nil {
		return err
	}

	// Loop through rows of table, slicing matching rows (if any) from nestable.
	for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {

		// Get the search key(s) for this row.
		for colIndex := 0; colIndex < len(cols); colIndex++ {
			keys[colIndex], err = table.GetVal(cols[colIndex], rowIndex)
			if err != nil {
				return err
			}
		}

		firstRow, lastRow, err := nestable.SearchRange(keys...)
		if err != nil {
			return err
		}

		nestTableName := nestColName
		newTable, err := NewTableFromRows(nestable, nestTableName, firstRow, lastRow)
		if err != nil {
			return err
		}

		err = table.SetString(nestColName, rowIndex, newTable.StringUnpadded())
		if err != nil {
			return err
		}
	}

	return nil
}
