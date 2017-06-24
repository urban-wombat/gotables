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
	"math"
	"reflect"
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

	// where()
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

	// where()
		return localMerged, nil
	}

	// where()
	if table1 == nil {
		err = fmt.Errorf("func (table1 *Table) %s(table2 *Table): table1 is <nil>\n", funcName())
		return merged, err
	}

	// where()
	if table2 == nil {
		err = fmt.Errorf("func (table1 *Table) %s(table2 *Table): table2 is <nil>\n", funcName())
		return merged, err
	}

	// where()
	if table1.RowCount() == 0 {
		merged, err = sortMerged(table2)
		if err != nil {
			return nil, err
		}
		return table2, nil
	}

	// where()
	if table2.RowCount() == 0 {
		merged, err = sortMerged(table1)
		if err != nil {
			return nil, err
		}
		return table1, nil
	}

	// where()
	// Check that table1 and table2 have the same sort columns.
	err = setSortKeysBetweenTables()
	if err != nil {
		return nil, err
	}

	// Okay. They're compatible, now set up for merging.

/*
	// where()
	colInfoTable, err := table1.GetColInfoAsTable()
	if err != nil {
		return nil, err
	}
	// where(fmt.Sprintf("colInfoTable =\n%s\n", colInfoTable))
*/

	// where()
	merged, err = NewTable("Merged")
	if err != nil {
		return nil, err
	}

	// where()
	// Add all columns (but not yet rows) from table1 and table2 into merged.

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

	// where(fmt.Sprintf("merged = %s\n", merged))

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

	err = merged.AppendSortKey("_TNUM_")
	if err != nil {
		return nil, err
	}
	err = merged.Sort()
	if err != nil {
		return nil, err
	}
	err = merged.DeleteSortKey("_TNUM_")
	if err != nil {
		return nil, err
	}

	// Add a column to keep track of which columns are duplicates, to be deleted.
	const deleteColName = "_DELETE_"
	err = merged.AppendCol(deleteColName, "bool")
	if err != nil {
		return nil, err
	}

where(fmt.Sprintf("BEFORE Merge()\n%s\n", merged))

	// Loop through to second-last row, comparing each row with the row after it.
	/*
		There are 4 possibilities (based on whether the values are zero values):
		Combination	val1		val2		Action
		(a)			zero		zero		do nothing
		(b)			zero		non-zero	copy val2 to val1	Assumes zero is a missing value
		(c)			non-zero	zero		copy val1 to val2	Assumes zero is a missing value
		(d)			non-zero	non-zero	copy val1 to val2	(table1 takes precedence)
		There are 2 further possibilities with float32 and float64:
		(e)			zero		NaN			copy val1 to val2	Assumes zero is NOT a missing value
		(f)			NaN			zero		copy val2 to val1	Assumes zero is NOT a missing value
	*/
	for rowIndex := 0; rowIndex < merged.RowCount()-1; rowIndex++ {
		comparison, err := merged.CompareRows(rowIndex, rowIndex+1)
		if err != nil {
			return nil, err
		}
//		where(fmt.Sprintf("[%s].CompareRows(%d, %d) = %d\n", merged.Name(), rowIndex, rowIndex+1, comparison))
		if comparison == 0 {
			// They are equal.
			// Loop through columns, one short of the last column which is the (temporary) table number column.
			for colIndex := 0; colIndex < merged.ColCount()-1; colIndex++ {
				colType, err := merged.ColTypeByColIndex(colIndex)
				if err != nil {
					return nil, err
				}
				switch colType {
					case "string":
						var val1 string
						var val2 string
						const zeroVal = ""
						val1, err = merged.GetStringByColIndex(colIndex, rowIndex)
						if err != nil {
							return nil, err
						}
						val2, err = merged.GetStringByColIndex(colIndex, rowIndex+1)
						if err != nil {
							return nil, err
						}
						if val1 != zeroVal {	// Covers combinations (c) and (d)
							err = merged.SetStringByColIndex(colIndex, rowIndex+1, val1)	// Use val1
							if err != nil {
								return nil, err
							}
						} else if val2 != zeroVal {	// Covers combination (b)
							err = merged.SetStringByColIndex(colIndex, rowIndex, val2)	// Use val2
							if err != nil {
								return nil, err
							}
						}
						// Otherwise both vals must be zero. Do nothing.
					case "bool":
						var val1 bool
						var val2 bool
						const zeroVal = false
						val1, err = merged.GetBoolByColIndex(colIndex, rowIndex)
						if err != nil {
							return nil, err
						}
						val2, err = merged.GetBoolByColIndex(colIndex, rowIndex+1)
						if err != nil {
							return nil, err
						}
						if val1 != zeroVal {	// Covers combinations (c) and (d)
							err = merged.SetBoolByColIndex(colIndex, rowIndex+1, val1)	// Use val1
							if err != nil {
								return nil, err
							}
						} else if val2 != zeroVal {	// Covers combination (b)
							err = merged.SetBoolByColIndex(colIndex, rowIndex, val2)	// Use val2
							if err != nil {
								return nil, err
							}
						}
						// Otherwise both vals must be zero. Do nothing.
					case "int8", "int16", "int32", "int64", "int":
						var tmp1 interface{}
						var tmp2 interface{}
						var val1 int64
						var val2 int64
						const zeroVal = 0
						tmp1, err = merged.GetValByColIndex(colIndex, rowIndex)
						if err != nil {
							return nil, err
						}
						val1 = reflect.ValueOf(tmp1).Int()
						tmp2, err = merged.GetValByColIndex(colIndex, rowIndex+1)
						if err != nil {
							return nil, err
						}
						val2 = reflect.ValueOf(tmp2).Int()
						if val1 != zeroVal {	// Covers combinations (c) and (d)
							err = merged.SetValByColIndex(colIndex, rowIndex+1, tmp1)	// Use val1
							if err != nil {
								return nil, err
							}
						} else if val2 != zeroVal {	// Covers combination (b)
							err = merged.SetValByColIndex(colIndex, rowIndex, tmp2)	// Use val2
							if err != nil {
								return nil, err
							}
						}
						// Otherwise both vals must be zero. Do nothing.
					case "uint8", "uint16", "uint32", "uint64", "uint":
						var tmp1 interface{}
						var tmp2 interface{}
						var val1 uint64
						var val2 uint64
						const zeroVal = 0
						tmp1, err = merged.GetValByColIndex(colIndex, rowIndex)
						if err != nil {
							return nil, err
						}
						val1 = reflect.ValueOf(tmp1).Uint()
						tmp2, err = merged.GetValByColIndex(colIndex, rowIndex+1)
						if err != nil {
							return nil, err
						}
						val2 = reflect.ValueOf(tmp2).Uint()
						if val1 != zeroVal {	// Covers combinations (c) and (d)
							err = merged.SetValByColIndex(colIndex, rowIndex+1, tmp1)	// Use val1
							if err != nil {
								return nil, err
							}
						} else if val2 != zeroVal {	// Covers combination (b)
							err = merged.SetValByColIndex(colIndex, rowIndex, tmp2)	// Use val2
							if err != nil {
								return nil, err
							}
						}
						// Otherwise both vals must be zero. Do nothing.
					case "float32", "float64":
						// Note: NaN is more zero than zero, so zero value 0.0 trumps NaN.
						var tmp1 interface{}
						var tmp2 interface{}
						var val1 float64
						var val2 float64
						const zeroVal = 0.0
						tmp1, err = merged.GetValByColIndex(colIndex, rowIndex)
						if err != nil {
							return nil, err
						}
						val1 = reflect.ValueOf(tmp1).Float()
						tmp2, err = merged.GetValByColIndex(colIndex, rowIndex+1)
						if err != nil {
							return nil, err
						}
						val2 = reflect.ValueOf(tmp2).Float()
						if val1 != zeroVal && !math.IsNaN(val1) {	// Covers combinations (c) and (d)
							where(fmt.Sprintf("val1 %f != zeroVal", val1))
							err = merged.SetValByColIndex(colIndex, rowIndex+1, tmp1)	// Use val1
							if err != nil {
								return nil, err
							}
						} else if val2 != zeroVal && !math.IsNaN(val2) {	// Covers combination (b)
							where(fmt.Sprintf("val2 %f != zeroVal", val2))
							err = merged.SetValByColIndex(colIndex, rowIndex, tmp2)	// Use val2
							if err != nil {
								return nil, err
							}
						} else if math.IsNaN(val1) { // Maybe one of them is NaN and the other is zero.
							where("math.IsNaN(val1)")
							err = merged.SetValByColIndex(colIndex, rowIndex, tmp2)	// Use val2
							if err != nil {
								return nil, err
							}
						} else if math.IsNaN(val2) { // Maybe one of them is NaN and the other is zero.
							where("math.IsNaN(val2)")
							err = merged.SetValByColIndex(colIndex, rowIndex+1, tmp1)	// Use val1
							if err != nil {
								return nil, err
							}
						}
						// Otherwise both vals must be zero. Do nothing.
					default:
						// Should never reach here.
						isValid, err := IsValidColType(colType)
						if !isValid {
							return nil, err
						} else {
							return nil, fmt.Errorf("What? We seem to have an unlisted type: %s", colType)
						}
				}
			}
			err = merged.SetBool(deleteColName, rowIndex+1, true)
			if err != nil {
				return nil, err
			}
		}
	}

where()
	return merged, nil
}

// This copies from left (srcTable) to right (targTable) beginning at beginRow in targTable.
func (srcTable *Table) copyTableCells(beginRow int, targTable *Table) error {
	for srcCol := 0; srcCol < srcTable.ColCount(); srcCol++ {
		colName, err := srcTable.ColName(srcCol)
		// where(fmt.Sprintf("srcTable.ColName(%d) = %q\n", srcCol, colName))
		if err != nil {
			return err
		}
		// Note: multiple assignment syntax in for loop.
		for srcRow, targRow := 0, beginRow; targRow < (beginRow + srcTable.RowCount()); srcRow, targRow = srcRow+1, targRow+1 {
			cellVal, err := srcTable.GetValByColIndex(srcCol, srcRow)
			// where(fmt.Sprintf("srcTable.GetValByColIndex(%d, %d) = %v\n", srcCol, srcRow, cellVal))
			if err != nil {
				return err
			}
			err = targTable.SetVal(colName, targRow, cellVal)
			// where(fmt.Sprintf("targTable.SetVal(%q, %d, %v)\n", colName, targRow, cellVal))
			if err != nil {
				return err
			}
			// where(fmt.Sprintln())
		}
	}

	return nil
}
