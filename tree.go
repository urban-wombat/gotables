package gotables

import (
	"fmt"
	"os"
	"strings"
)

/*
	Visit each table in tableSet.

	To visit the root table and any nested child tables: run visitTable on each table.

	To visit all rows in each table: run visitRow on each row.

	To visit all cells in each table: run visitCell on each cell.

	Define func variables visitTable and visitCell (see Example).

	If visitTable or visitCell are nil, no action will be taken in the nil case.
*/
func (tableSet *TableSet) Walk(
	walkNestedTables bool,
	visitTableSet func(*TableSet) error,
	visitTable func(*Table) error,
	visitRow func(Row) error,
	visitCell func(bool, CellInfo) error) (err error) {

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
		table, err = tableSet.GetTableByTableIndex(tableIndex)
		if err != nil {
			return
		}

		err = table.Walk(walkNestedTables, visitTable, visitRow, visitCell)
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
func (table *Table) Walk(
	walkNestedTables bool,
	visitTable func(*Table) error,
	visitRow func(Row) error,
	visitCell func(bool, CellInfo) error) (err error) {

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
		//where(fmt.Sprintf("row = %#v", row))
		//where(fmt.Sprintf("row = %v", row))

		// Visit row.
		if visitRow != nil {
			err = visitRow(row)
			if err != nil {
				return
			}
		}

		for colIndex := 0; colIndex < table.ColCount(); colIndex++ {

			var cellInfo CellInfo
			cellInfo, err = table.GetCellInfoByColIndex(colIndex, rowIndex)
			if err != nil {
				return
			}

			// Visit cell.
			if visitCell != nil {
				err = visitCell(walkNestedTables, cellInfo)
				if err != nil {
					return
				}
			}

			if walkNestedTables {
				if IsTableColType(cellInfo.ColType) {

					var nestedTable *Table
					nestedTable, err = table.GetTableByColIndex(colIndex, rowIndex)
					if err != nil {
						return
					}

					// Down into nested table.
					nestedTable.depth++

					// Recursive call to visit nested tables.
					err = nestedTable.Walk(walkNestedTables, visitTable, visitRow, visitCell)
					if err != nil {
						return
					}

					// Back up from nested table.
					nestedTable.depth--
				}
			}
		}
	}

	return
}

/*
	Return a type CellInfo struct populated with *Table, col and row information:

		type CellInfo struct {
			Table    *Table
			ColName  string
			ColType  string
			ColIndex int
			RowIndex int
		}

	ColName, ColType, ColIndex, and RowIndex are read-only.

	Caution: Table is a mutable reference to the enclosing *Table.

		GetVal()
		GetValByColIndex()
		Get String()
		Get StringByColIndex()
*/
func (table *Table) GetCellInfoByColIndex(colIndex int, rowIndex int) (cellInfo CellInfo, err error) {

	if table == nil {
		return cellInfo, fmt.Errorf("table.%s(): table is nil", UtilFuncNameNoParens())
	}

	cellInfo.Table = table
	cellInfo.ColIndex = colIndex
	cellInfo.RowIndex = rowIndex
	cellInfo.ColType = table.colTypes[colIndex]

	cellInfo.ColName, err = table.ColName(colIndex)
	if err != nil {
		return cellInfo, err
	}

	return cellInfo, nil
}

/*
	Return a type CellInfo struct populated with *Table, col and row information:

		type CellInfo struct {
			Table    *Table
			ColName  string
			ColType  string
			ColIndex int
			RowIndex int
		}

	ColName, ColType, ColIndex, and RowIndex are read-only.

	Caution: Table is a mutable reference to the enclosing *Table.

	If you want the value of a cell (and not a type CellInfo struct) use:

		GetVal()
		GetValByColIndex()
		Get String()
		Get StringByColIndex()
*/
func (table *Table) GetCellInfo(colName string, rowIndex int) (cellInfo CellInfo, err error) {

	if table == nil {
		return cellInfo, fmt.Errorf("table.%s(): table is nil", UtilFuncNameNoParens())
	}

	var colIndex int
	colIndex, err = table.ColIndex(colName)
	if err != nil {
		return cellInfo, err
	}

	cellInfo, err = table.GetCellInfoByColIndex(colIndex, rowIndex)
	if err != nil {
		return cellInfo, err
	}

	return cellInfo, nil
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
func (cellInfo CellInfo) TableName() string {
	return cellInfo.Table.Name()
}
*/

/*
	This is a logical flip of table.IsValidTableNesting()
*/
func (rootTable *Table) HasCircularReference() (hasCircularReference bool, err error) {
	if rootTable == nil {
		return false, fmt.Errorf("rootTable.%s(): rootTable is nil", UtilFuncNameNoParens())
	}

	isValidTableNesting, err := rootTable.IsValidTableNesting()
	if !isValidTableNesting {
		// Say the error is from here, not from the called method.
		replacedMsg := strings.Replace(err.Error(), "IsValidTableNesting", "HasCircularReference", 1)
		err = fmt.Errorf("%s", replacedMsg)
		return !isValidTableNesting, err
	}

	return
}

/*
	This is a logical flip of table.HasCircularReference()
*/
func (rootTable *Table) IsValidTableNesting() (valid bool, err error) {
	//UtilPrintCaller()

	const funcName = "IsValidTableNesting()"
	const walkNestedTables = true

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

	var visitCell func(walkNestedTables bool, cellInfo CellInfo) (err error)
	visitCell = func(walkNestedTables bool, cellInfo CellInfo) (err error) {
		if IsTableColType(cellInfo.ColType) {
			var nestedTable *Table
			nestedTable, err = rootTable.GetTableByColIndex(cellInfo.ColIndex, cellInfo.RowIndex)
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

	err = rootTable.Walk(walkNestedTables, nil, nil, visitCell)
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
