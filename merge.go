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
			err = setSortKeysBetweenTables()
			if err != nil {
				return nil, err
			}
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

	// This is purely aesthetic for human readability.
	err = merged.OrderColsBySortKeys()
	if err != nil {
		return nil, err
	}

	// Add a column to keep track of which columns came from which table.
	const tableNumberColName = "_TNUM_"
	err = merged.AppendCol(tableNumberColName, "int")
	if err != nil {
		return nil, err
	}

	// Make space for both tables.
	err = merged.AppendRows(table1.RowCount() + table2.RowCount())
	if err != nil {
		return nil, err
	}

	// Set floats to NaN values.
	err = merged.SetAllFloatCellsToNaN()
	if err != nil {
		return nil, err
	}

fmt.Println(merged)

	// Copy table1 into merged.
	var beginRow = 0
	err = table1.copyTableCells(beginRow, merged)
	if err != nil {
		return nil, err
	}

	// Set table number in merged of table1.
	for rowIndex := beginRow; rowIndex < table1.RowCount(); rowIndex++ {
		err = merged.SetInt(tableNumberColName, rowIndex, 1)
		if err != nil {
			return nil, err
		}
	}

	// Copy table2 into merged.
	beginRow = table1.RowCount()
	err = table2.copyTableCells(beginRow, merged)
	if err != nil {
		return nil, err
	}

	// Set table number in merged of table2.
	for rowIndex := beginRow; rowIndex < merged.RowCount(); rowIndex++ {
		err = merged.SetInt(tableNumberColName, rowIndex, 2)
		if err != nil {
			return nil, err
		}
	}

	err = merged.Sort()
	if err != nil {
		return nil, err
	}

where()
	return merged, nil
}

// This copies from left (srcTable) to right (targTable) beginning at beginRow in targTable.
func (srcTable *Table) copyTableCells(beginRow int, targTable *Table) error {
	for srcCol := 0; srcCol < srcTable.ColCount(); srcCol++ {
		colName, err := srcTable.ColName(srcCol)
		where(fmt.Sprintf("srcTable.ColName(%d) = %q\n", srcCol, colName))
		if err != nil {
			return err
		}
		// Note: multiple assignment syntax in for loop.
		for srcRow, targRow := 0, beginRow; targRow < (beginRow + srcTable.RowCount()); srcRow, targRow = srcRow+1, targRow+1 {
			cellVal, err := srcTable.GetValByColIndex(srcCol, srcRow)
			where(fmt.Sprintf("srcTable.GetValByColIndex(%d, %d) = %v\n", srcCol, srcRow, cellVal))
			if err != nil {
				return err
			}
			err = targTable.SetVal(colName, targRow, cellVal)
			where(fmt.Sprintf("targTable.SetVal(%q, %d, %v)\n", colName, targRow, cellVal))
			if err != nil {
				return err
			}
			where(fmt.Sprintln())
		}
	}

	return nil
}
