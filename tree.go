package gotables

import (
	"fmt"
)

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

			isTable, _ := IsTableColType(cell.ColType)

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
