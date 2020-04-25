package gotables

import (
	"fmt"
)

//	/*
//		Returns false if ANY table (including the top level table) exists more than once in the top level table
//		or in any nested table. In other words: no duplicate tables whatsoever.
//
//		The purpose is to completely eliminate the possibility of circular references.
//	*/
//	func (table *Table) IsValidTableNesting() (valid bool, err error) {
//
//		if table == nil {
//			return false, fmt.Errorf("table.%s(): table is nil", UtilFuncNameNoParens())
//		}
//
//		var refMap circRefMap = map[*Table]struct{}{}
//
//		valid, err = isValidTableNesting_recursive(table, table, refMap)
//		if err != nil {
//			return false, err
//		}
//
//		return
//	}

/*
func DISMEMBERING(topTable *Table, table *Table, refMap circRefMap) (bool, error) {

	if table == nil {
		return false, fmt.Errorf("table.%s(): table is nil", UtilFuncNameNoParens())
	}

	refMap[table] = empty // Add this table to the map.

	// Compare cell values.
	for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
		colName, err := table.ColName(colIndex)
		if err != nil {
			return false, err
		}

		colType, err := table.ColTypeByColIndex(colIndex)
		if err != nil {
			return false, err
		}

		isTable, _ := IsTableColType(colType)

		if isTable {
			for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {

				nestedTable, err := table.GetTableByColIndex(colIndex, rowIndex)
				if err != nil {
					return false, err
				}

				// Have we already seen this table?
				_, exists := refMap[nestedTable]
				if exists { // Invalid table with circular reference!
					err = fmt.Errorf("%s: circular reference in table [%s]: a reference to table [%s] already exists",
						UtilFuncName(), topTable.Name(), nestedTable.Name())
					return false, err
				} else { // No?
					refMap[nestedTable] = empty // Add this table to the map.
				}

				// Recursive call to check nested tables.
				valid, err := isValidTableNesting_recursive(topTable, nestedTable, refMap)
				if valid == false {
					// Hmmm. Maybe the error message will be composed with each recursive call?
					parentTableName := nestedTable.parentTable.Name()
					return false, fmt.Errorf("%v (parent table [%s] colIndex=%d colName=%q rowIndex=%d)",
						err, parentTableName, colIndex, colName, rowIndex)
				}
			}
		}
	}

	return true, nil
}
*/

func (table *Table) Walk(visit func(interface{})) (err error) {

	if table == nil {
		return fmt.Errorf("table.%s(): table is nil", UtilFuncNameNoParens())
	}

	err = table.walk(visit)
	if err != nil {
		return err
	}

	return
}

func (tableSet *TableSet) Walk(visit func(interface{})) (err error) {

	if tableSet == nil {
		return fmt.Errorf("TableSet.%s(): tableSet is nil", UtilFuncNameNoParens())
	}

	for tableIndex := 0; tableIndex < tableSet.TableCount(); tableIndex++ {
		table, err := tableSet.TableByTableIndex(tableIndex)
		if err != nil {
			return err
		}

		err = table.walk(visit)
		if err != nil {
			return err
		}
	}

	return
}

func (table *Table) walk(visit func(interface{})) (err error) {
	where()
	visit(table)
	where()
	return
}
