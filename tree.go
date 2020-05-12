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
func (tableSet *TableSet) Walk(
	visitTableSet func(*TableSet) error, visitTable func(*Table) error, visitRow func(Row) error, visitCell func(Cell) error) (err error) {

	if tableSet == nil {
		err = fmt.Errorf("TableSet.%s(): tableSet is nil", UtilFuncNameNoParens())
		return
	}

	// Visit tableSet.
	if visitTableSet != nil {
		err = visitTableSet(tableSet)
		if err != nil {
			return
		}
	}

	for tableIndex := 0; tableIndex < tableSet.TableCount(); tableIndex++ {

		var table *Table
		table, err = tableSet.TableByTableIndex(tableIndex)
		if err != nil {
			return
		}

		err = table.Walk(visitTable, visitRow, visitCell)
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
func (table *Table) Walk(visitTable func(*Table) error, visitRow func(Row) error, visitCell func(Cell) error) (err error) {

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

	// Visit cell values, row by row.

	for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {

		var row Row
		row, err = table.RowByRowIndex(rowIndex)
		if err != nil {
			return
		}

		// Visit row.
		if visitRow != nil {
			err = visitRow(row)
			if err != nil {
				return
			}
		}

		for colIndex := 0; colIndex < table.ColCount(); colIndex++ {

			var cell Cell
			cell, err = table.CellByColIndex(colIndex, rowIndex)
			if err != nil {
				return
			}

			// Visit cell.
			if visitCell != nil {
				err = visitCell(cell)
				if err != nil {
//where()
//UtilPrintCallerCaller()
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

				// Down into nested table.
				nestedTable.depth++

				// Recursive call to visit nested tables.
				err = nestedTable.Walk(visitTable, visitRow, visitCell)
				if err != nil {
					return
				}

				// Back up from nested table.
				nestedTable.depth--
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

func (table *Table) RowByRowIndex(rowIndex int) (row Row, err error) {

	if table == nil {
		return row, fmt.Errorf("table.%s(): table is nil", UtilFuncNameNoParens())
	}

	row.Table = table
	row.RowIndex = rowIndex

	return row, nil
}

/*
func (cell Cell) TableName() string {
	return cell.Table.Name()
}
*/

func (rootTable *Table) IsValidTableNesting() (valid bool, err error) {
//UtilPrintCaller()

	const funcName = "IsValidTableNesting()"

	if rootTable == nil {
		return false, fmt.Errorf("rootTable.%s(): rootTable is nil", UtilFuncNameNoParens())
	}

	if rootTable.parentTable != nil {
		msg := fmt.Sprintf("rootTable.parentTable should be <nil> but found: %q", rootTable.parentTable.Name())
		circError := NewCircRefError(rootTable, nil, msg)
		err = fmt.Errorf("%s: %w", funcName, circError) // Wrap circError in err.
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

			if nestedTable.parentTable == nil {
				msg := fmt.Sprintf("nestedTable.parentTable should not be <nil>")
				circError := NewCircRefError(rootTable, nestedTable, msg)
				err = fmt.Errorf("%s: %w", funcName, circError) // Wrap circError in err.
			}

			// Have we already seen this table?
			_, exists := refMap[nestedTable]
			if exists { // Invalid table with circular reference!
				// Construct CircRefError.
				circError := NewCircRefError(rootTable, nestedTable, "")
				err = fmt.Errorf("%s: %w", funcName, circError) // Wrap circError in err.
				return err
			} else {
				// Add this nested table to the map.
				refMap[nestedTable] = EmptyStruct
			}
		}
		return nil
	}

	err = rootTable.Walk(nil, nil, visitCell)
	if err != nil {
		// Found a circular reference!
		return false, err
	}

	return true, nil
}

func (parentTable *Table) isCircularReference(candidateChildTable *Table) (isCircular bool, depth int) {
	//UtilPrintCaller()
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
