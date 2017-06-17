package gotable


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


import (
	"fmt"
)

func (table1 *Table) Merge(table2 *Table) (merged *Table, err error) {

//	var err error
//	var merged *Table

	// Local function.
	// Make sort keys of both input tables the same.
	setSortKeysBetweenTables := func () error {
		if table1.SortKeyCount() > 0 {
			// Table1 is dominant.
			err = table2.SetSortKeysFromTable(table1)
			if err != nil {
				return err
			}
		} else if table2.SortKeyCount() > 0 {
			err = table1.SetSortKeysFromTable(table2)
			if err != nil {
				return err
			}
		} else {
			err = fmt.Errorf("[%s].Merge([%s]) needs at least [%s] or [%s] to have sort keys",
				table1.Name(), table2.Name(), table1.Name(), table2.Name())
			return err
		}
		return nil
	}

where()
	// Local function.
	sortMerged := func (localMerged *Table) (*Table, error) {
		// TODO: Copy sort keys from table1 or table2 to merged
		if localMerged.SortKeyCount() == 0 {
			setSortKeysBetweenTables()
		}

		err = localMerged.Sort()
		if err != nil {
			return nil, err
		}

where()
		return localMerged, nil
	}

where()
	if table1 == nil {
		err = fmt.Errorf("func (table1 *Table) %s(table2 *Table): table1 is <nil>\n", funcName())
		return merged, err
	}

where()
	if table2 == nil {
		err = fmt.Errorf("func (table1 *Table) %s(table2 *Table): table2 is <nil>\n", funcName())
		return merged, err
	}

where()
	if table1.RowCount() == 0 {
		merged, err = sortMerged(table2)
		if err != nil {
			return nil, err
		}
		return table2, nil
	}

where()
	if table2.RowCount() == 0 {
		merged, err = sortMerged(table1)
		if err != nil {
			return nil, err
		}
		return table1, nil
	}

where()
	// Check that table1 and table2 have the same sort columns.
	err = setSortKeysBetweenTables()
	if err != nil {
		return nil, err
	}

	// Okay. They're compatible, now set up for merging.

where()
	// Add all columns from table1 and table2 into merged.
	colsTable, err := table1.GetColInfoAsTable()
	if err != nil {
		return nil, err
	}
	fmt.Println(colsTable)

where()
	merged, err = NewTable("Merged")
	if err != nil {
		return nil, err
	}

where()
	err = merged.AppendColsFromTable(table1)
	if err != nil {
		return nil, err
	}
//	fmt.Println(merged)

	err = merged.AppendColsFromTable(table2)
	if err != nil {
		return nil, err
	}
//	fmt.Println(merged)

	err = merged.SetSortKeysFromTable(table1)
	if err != nil {
		return nil, err
	}

	err = merged.OrderColsBySortKeys()
	if err != nil {
		return nil, err
	}

where()
	return merged, nil
}
