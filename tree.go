package gotables

import (
	"fmt"
	"os"
)

/*
	Visit each table in tableSet.

	Visit the root table and any nested child tables: run visitTable on each table.

	Visit all cells in each table: run visitCell on each cell.

	Define func variables visitTable and visitCell (see Example).

	If visitTable or visitCell are nil, no action will be taken in the nil case.
*/
func (tableSet *TableSet) Walk(visitTable func(*Table) error, visitCell func(Cell) error, in interface{}) (out interface{}, err error) {

	if tableSet == nil {
		err = fmt.Errorf("TableSet.%s(): tableSet is nil", UtilFuncNameNoParens())
		return
	}

	for tableIndex := 0; tableIndex < tableSet.TableCount(); tableIndex++ {

		var table *Table
		table, err = tableSet.TableByTableIndex(tableIndex)
		if err != nil {
			return
		}

		out, err = table.Walk(visitTable, visitCell, in)
		if err != nil {
			return
		}
	}

	return
}

/*
	Visit the root table and any nested child tables: run visitTable on each table.

	Visit all cells in each table: run visitCell on each cell.

	Define func variables visitTable and visitCell (see Example).

	If visitTable or visitCell are nil, no action will be taken in the nil case.
*/
func (table *Table) Walk(visitTable func(*Table) error, visitCell func(Cell) error, in interface{}) (out interface{}, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is nil", UtilFuncNameNoParens())
		return
	}

	// Visit table.
	if visitTable != nil {
		err = visitTable(table)
		if err != nil {
			return
		}
	}

	// Visit cell values.
	for colIndex := 0; colIndex < table.ColCount(); colIndex++ {

		for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {

			var cell Cell
			cell, err = table.CellByColIndex(colIndex, rowIndex)
			if err != nil {
				return
			}

			// Visit cell.
			if visitCell != nil {
				err = visitCell(cell)
				if err != nil {
					return
				}
			}

			isTable := IsTableColType(cell.ColType)

			if isTable {

				var nestedTable *Table
				nestedTable, err = table.GetTableByColIndex(colIndex, rowIndex)
				if err != nil {
					return
				}

				// Recursive call to visit nested tables.
				out, err = nestedTable.Walk(visitTable, visitCell, in)
				if err != nil {
					return
				}
			}
		}
	}

	return
}

func (table *Table) CellByColIndex(colIndex int, rowIndex int) (cell Cell, err error) {

	if table == nil {
		return cell, fmt.Errorf("table.%s(): table is nil", UtilFuncNameNoParens())
	}

	cell.Table = table
	cell.ColIndex = colIndex
	cell.RowIndex = rowIndex
	cell.ColType = table.colTypes[colIndex]

	cell.ColName, err = table.ColName(colIndex)
	if err != nil {
		return cell, err
	}

	return cell, nil
}

func (table *Table) Cell(colName string, rowIndex int) (cell Cell, err error) {

	if table == nil {
		return cell, fmt.Errorf("table.%s(): table is nil", UtilFuncNameNoParens())
	}

	var colIndex int
	colIndex, err = table.ColIndex(colName)
	if err != nil {
		return cell, err
	}

	cell, err = table.CellByColIndex(colIndex, rowIndex)
	if err != nil {
		return cell, err
	}

	return cell, nil
}

/*
func (cell Cell) TableName() string {
	return cell.Table.Name()
}
*/

func (rootTable *Table) IsValidTableNesting() (valid bool, err error) {

	const funcName = "IsValidTableNesting()"

	if rootTable == nil {
		return false, fmt.Errorf("rootTable.%s(): rootTable is nil", UtilFuncNameNoParens())
	}

	var refMap circRefMap = map[*Table]struct{}{}
	refMap[rootTable] = EmptyStruct // Add the root table to the map.

	var visitCell func(cell Cell) (err error)
	visitCell = func(cell Cell) (err error) {
		if IsTableColType(cell.ColType) {
			var nestedTable *Table
			nestedTable, err = rootTable.GetTableByColIndex(cell.ColIndex, cell.RowIndex)
			if err != nil {
				return err
			}

			// Have we already seen this table?
			_, exists := refMap[nestedTable]
			if exists { // Invalid table with circular reference!
				// Construct CircRefError.
				circError := NewCircRefError(rootTable, nestedTable)
				err = fmt.Errorf("%s: %w", funcName, circError) // Wrap circError in err.
				return err
			} else {
				// Add this nested table to the map.
				refMap[nestedTable] = EmptyStruct
			}
		}
		return nil
	}

	_, err = rootTable.Walk(nil, visitCell, nil)
	if err != nil {
		// Found a circular reference!
		return false, err
	}

	return true, nil
}

func (parentTable *Table) isCircularReference(candidateChildTable *Table) (isCircular bool, depth int) {
	UtilPrintCaller()
	// where(fmt.Sprintf("isCircularReference(parentTable:%p, candidateChildTable:%p)", parentTable, candidateChildTable))
	depth = 1 // Can only have a circular reference with depth 1 or more.
	if parentTable == candidateChildTable {
		// where(fmt.Sprintf("depth:%d parentTable == candidateChildTable", depth))
		return true, depth
	}
	// where(fmt.Sprintf("parentTable.parentTable:%p", parentTable.parentTable))
	for depth = 1; parentTable.parentTable != nil; depth++ {
		// where(fmt.Sprintf("fff depth:%d parentTable:%p candidateChildTable:%p", depth, parentTable, candidateChildTable))
		if parentTable == candidateChildTable {
			// where(fmt.Sprintf("depth:%d parentTable == candidateChildTable", depth))
			return true, depth
		}
		parentTable = parentTable.parentTable
		// where(parentTable.String())
		// where(fmt.Sprintf("*** depth:%d parentTable ref:%p parentTable name:%s parentTable ref:%p", depth, parentTable, parentTable.Name(), parentTable.parentTable))
		if depth >= 9 {
			os.Exit(44)
		}
	}

	// where(fmt.Sprintf("depth:%d return -1", depth))
	return false, -1
}
