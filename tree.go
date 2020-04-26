package gotables

import (
	"fmt"
)

/*
	Visit each table in tableSet.

	Visit the root table and any nested child tables: run visitTable on each table.

	Visit all cells in each table: run visitCell on each cell.

	Define func variables visitTable and visitCell (see Example).

	If visitTable or visitCell are nil, no action will be taken in the nil case.
*/
func (tableSet *TableSet) Walk(visitTable func(*Table) error, visitCell func(Cell) error) (err error) {

	if tableSet == nil {
		return fmt.Errorf("TableSet.%s(): tableSet is nil", UtilFuncNameNoParens())
	}

	for tableIndex := 0; tableIndex < tableSet.TableCount(); tableIndex++ {

		table, err := tableSet.TableByTableIndex(tableIndex)
		if err != nil {
			return err
		}

		err = table.Walk(visitTable, visitCell)
		if err != nil {
			return err
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
func (table *Table) Walk(visitTable func(*Table) error, visitCell func(Cell) error) (err error) {

	if table == nil {
		return fmt.Errorf("table.%s(): table is nil", UtilFuncNameNoParens())
	}

	// Visit table.
	if visitTable != nil {
		err = visitTable(table)
		if err != nil {
			return err
		}
	}

	// Visit cell values.
	for colIndex := 0; colIndex < table.ColCount(); colIndex++ {

		for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {

			var cell Cell
			cell, err := table.CellByColIndex(colIndex, rowIndex)
			if err != nil {
				return err
			}

			// Visit cell.
			if visitCell != nil {
				err = visitCell(cell)
				if err != nil {
					return err
				}
			}

			isTable := IsTableColType(cell.ColType)

			if isTable {

				nestedTable, err := table.GetTableByColIndex(colIndex, rowIndex)
				if err != nil {
					return err
				}

				// Recursive call to visit nested tables.
				err = nestedTable.Walk(visitTable, visitCell)
				if err != nil {
					return err
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

func (cell Cell) TableName() string {
	return cell.Table.Name()
}

/*
func (table *Table) IsValidTableNesting2() (valid bool, err error) {

	if table == nil {
		return false, fmt.Errorf("table.%s(): table is nil", UtilFuncNameNoParens())
	}

	var refMap circRefMap = map[*Table]struct{}{}

	// Add the root table to the map.
	refMap[table] = struct{}	// empty struct

	var visitCell func (cell Cell) (err error)
	visitCell = func (cell Cell) (err error) {
		if IsTableColType(cell.ColType) {
		}
	}
}
*/