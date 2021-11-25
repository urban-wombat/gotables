package gotables

import (
	"fmt"
	"os"
	"strings"
)

var CURRENT_TESTS_VERBOSE bool = false

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

		var walkSafe WalkSafe = make(WalkSafe)
		err = table.Walk(walkNestedTables, walkSafe, visitTable, visitRow, visitCell)
		if err != nil {
			return
		}
	}

	return
}

type WalkSafe map[*Table]struct{}

/*
	Visit the root table, each of its cells, and any nested child tables (inside cells):

		* run visitTable() on each table
		* run visitRow() on each row
		* run visitCell() on each cell

	Define one or more func variables visitTable(), visitRow() and visitCell() (see Example).

	If visitTable(), visitRow() or visitCell() are nil, no action will be taken.
*/
func (table *Table) Walk(
	walkNestedTables bool,
	walkSafe WalkSafe,
	visitTable func(*Table) error,
	visitRow func(Row) error,
	visitCell func(bool, CellInfo) error) (err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is nil", UtilFuncNameNoParens())
		return
	}

	if table == nil {
		err = fmt.Errorf("table.%s(): walkSafe is nil", UtilFuncNameNoParens())
		return
	}

	// where(fmt.Sprintf("Walk() walkSafe %p = %v", walkSafe, walkSafe))

	// Visit table.
	if visitTable != nil {
		// where("visitTable")
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
		// where(fmt.Sprintf("visiting cell col %d row %d", colIndex, rowIndex))

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

					var isNilTable bool
					isNilTable, err = nestedTable.IsNilTable()
					if err != nil {
						return
					}

					if !isNilTable {
						// Only worry about tables that may have nested tables.

						// Have we already seen this table?
// where(fmt.Sprintf("walkSafe %p = %v", walkSafe, walkSafe))
						_, exists := walkSafe[nestedTable]
// where(fmt.Sprintf("[%s] exists = %t", nestedTable.Name(), exists))
						if exists { // Invalid table with circular reference!
							// Construct CircRefError.
							circError := NewCircRefError(table, nestedTable, "")
							err = fmt.Errorf("visitCell(): %s", circError) // Wrap circError in err.
							return err
						} else {
							// Add this nested table to the map.
// where(fmt.Sprintf("ADDING [%s] to %p %v", nestedTable.Name(), walkSafe, walkSafe))
							walkSafe[nestedTable] = EmptyStruct
// where(fmt.Sprintf("walkSafe %p = %v", walkSafe, walkSafe))
						}
					}

					// Down into nested table.
					nestedTable.depth++

					// Recursive call to visit nested tables.
// where(fmt.Sprintf("walkSafe %p = %v", walkSafe, walkSafe))
// where("calling nestedTable.Walk()")
					err = nestedTable.Walk(walkNestedTables, walkSafe, visitTable, visitRow, visitCell)
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

	var walkSafe WalkSafe = make(WalkSafe)
	err = rootTable.Walk(walkNestedTables, walkSafe, nil, nil, visitCell)
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

func (table *Table) CopyDeep() (tableCopy *Table, err error) {
	// where(UtilFuncName())
	if table == nil {
		return nil, fmt.Errorf("table.%s: table is nil", UtilFuncName())
	}

	// This may seem odd, but this is the point in CopyDeep() where it's possible
	// to set tableCopy before all the recursive calls have occurred.
	tableCopy = table

	var visitTable = func(t *Table) (err error) {
		// where(fmt.Sprintf("[%s]", t.Name()))
		return
	}

	var visitCell = func(walkDeep bool, cell CellInfo) (err error) {
		if IsTableColType(cell.ColType) {
			// where(fmt.Sprintf("[%s] IsTableColType", cell.Table.Name()))

			var nestedTable *Table
			nestedTable, err = cell.Table.GetTableByColIndex(cell.ColIndex, cell.RowIndex)
			if err != nil {
				return
			}
			// where(fmt.Sprintf("nestedTable = %p", nestedTable))

			var isNilTable bool
			isNilTable, err = nestedTable.IsNilTable()
			if err != nil {
				return
			}

			var nestedTableCopy *Table
			if isNilTable {
				nestedTableCopy = NewNilTable()	// New reference just to be on the safe side.
			} else {
				nestedTableCopy, err = nestedTable.Copy()
				if err != nil {
					return
				}
			}

			err = cell.Table.SetTableByColIndex(cell.ColIndex, cell.RowIndex, nestedTableCopy)
			if err != nil {
			// where(fmt.Sprintf("nestedTableCopy = %p", nestedTableCopy))
				return
			}
		}
		tableCopy = cell.Table

		return
	}

	const walkDeep = true
	var walkSafe WalkSafe = make(WalkSafe)
	// where(fmt.Sprintf("declaring walkSafe = %p", walkSafe))
	// where("BEGIN Walk()")
	err = table.Walk(walkDeep, walkSafe, visitTable, nil, visitCell)
	// where("END Walk()")
	if err != nil {
		return nil, err
	}

	return
}

/*
	Build nested treeTable from table.

	If table contains 1 or more cols of type *Tree then
	replicate tree to depth levels, where depth 1 means
	create 1 new level below table.

	If there are no type *Table cols in table, or if tablesDepth is set to 0, then
	return table unchanged.
*/
func (rootTable *Table) NewTreeTable1(tablesDepth int) (treeTable *Table, err error) {
	// where("caller: " + UtilFuncCaller())
	// where(UtilFuncName())
	const DontWalkNestedTables = true
	var depth int = -1	// Allow that visitTable pre-increments depth.
	var cellCount int = 0

	// This may be an unnecessary precaution to avoid repeated table references.
	// TODO: test this with circular reference checker.
	treeTable, err = rootTable.Copy()
	if err != nil {
		return nil, fmt.Errorf("%s %v", UtilFuncSource(), err)
	}

	var visitTable = func(treeTable *Table) (err error) {
		// where("visitTable")
		depth++
		// where(fmt.Sprintf("depth++ = %d", depth))
		return
	}

	var visitCell = func(DontWalkNestedTables bool, cell CellInfo) (err error) {
		// where("visitCell")
		cellCount++
		// where(fmt.Sprintf("cellCount = %d", cellCount))
		// where(fmt.Sprintf("depth=%d", depth))
		if depth >= tablesDepth {
			// where(depth)
			// where(fmt.Sprintf("visit return: depth=%d >= tablesDepth=%d", depth, tablesDepth))
			return
		}

		if cell.ColType == "*Table" {
			// where(fmt.Sprintf(`cell.ColType == "*Table"`))
			var nestedTable *Table
			nestedTable, err = cell.Table.GetTable(cell.ColName, cell.RowIndex)
			if err != nil {
				return fmt.Errorf("%s %v", UtilFuncSource(), err)
			}
			where(fmt.Sprintf("TABLE nestedTable colName=%s rowIndex=%d:\n%s", cell.ColName, cell.RowIndex, nestedTable.String()))

			// Replace only NilTable cell entries, and not cells that are already populated.
			if nestedTable.Name() == "" {	// NilTable doesn't have a name.
				where(fmt.Sprintf("[%s] col=%s row=%d is NilTable", cell.Table.Name(), cell.ColName, cell.RowIndex))
				var parentTable *Table = nestedTable.ParentTable()
				where(fmt.Sprintf("TABLE parentTable:\n%s", parentTable.String()))
				parentTableCopy, err := parentTable.Copy()
				if err != nil {
					return fmt.Errorf("%s %v", UtilFuncSource(), err)
				}
				tableCopyName := parentTableCopy.Name()
//				parentTableCopy.SetName(tableCopyName + "_" + string('a' + cellCount-1))
				newTableName := fmt.Sprintf("%s%c", tableCopyName, rune('A' + cellCount-1))
				// where(fmt.Sprintf("newTableName: %s", newTableName))
				err = parentTableCopy.SetName(newTableName)
				if err != nil {
					return fmt.Errorf("%s %v", UtilFuncSource(), err)
				}
				// where(fmt.Sprintf("parentTableCopy:\n%s", parentTableCopy))

				newColName := fmt.Sprintf("%s%02d", cell.ColName, cellCount)
				// where(fmt.Sprintf("newColName: %s", newColName))
				err = parentTableCopy.RenameCol(cell.ColName, newColName)
				if err != nil {
					return fmt.Errorf("%s %v", UtilFuncSource(), err)
				}

				// where(fmt.Sprintf("BEFORE treeTable:\n%s", treeTable.String()))
				err = treeTable.SetTable(cell.ColName, cell.RowIndex, parentTableCopy)
				if err != nil {
					return fmt.Errorf("%s %v", UtilFuncSource(), err)
				}
				// where(fmt.Sprintf("AFTER  treeTable:\n%s", treeTable.String()))
			}
		}
		return
	}

	var walkSafe WalkSafe = make(WalkSafe)
	err = treeTable.Walk(DontWalkNestedTables, walkSafe, visitTable, nil, visitCell)
	if err != nil {
		// where("visit return")
		return
	}

	// where("visit final return")
	return
}

/*
	Build and populate a nested treeTable from table.

	The returned table is not strictly new. It is the receiver table
	populated with as many nested tables as allowed by depth and
	the existence of nil tables to populate with non-nil tables.

	If table contains 1 or more cols of type *Tree then
	replicate tree to depth levels, where depth 1 means
	create 1 new level below table.

	If there are no type *Table cols in table, or if depth is set to 0, then
	return table unchanged.
*/
var globalTreeCalls int
var globalTreeDepth int
func (table *Table) NewTreeTable(depth int) (treeTable *Table, err error) {
globalTreeCalls++
if CURRENT_TESTS_VERBOSE {
where(fmt.Sprintf("NewTreeTable(globalTreeCalls=%d)\n", globalTreeCalls))
}
	// where("caller: " + UtilFuncCaller())
	// where(UtilFuncName())

	if table == nil {
		return nil, fmt.Errorf("%s table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
	}

	// where("\n" + table.String())
	nestedTable, err := table.Copy()
	if err != nil {
		return nil, fmt.Errorf("%s %v", UtilFuncSource(), err)
	}
	// where("\n" + nestedTable.String())

	err = newTreeTable_recursive(table, nestedTable, depth)
	if err != nil {
		return nil, fmt.Errorf("%s %v", UtilFuncSource(), err)
	}

	// where("visit final return")
	return table, nil
}

func newTreeTable_recursive(originalTable *Table, table *Table, depth int) (err error) {
globalTreeDepth++
if CURRENT_TESTS_VERBOSE {
where(fmt.Sprintf("newTreeTable_recursive(depth=%d)\n", depth))
where(fmt.Sprintf("newTreeTable_recursive(globalTreeCalls=%d)\n", globalTreeCalls))
where(fmt.Sprintf("newTreeTable_recursive(globalTreeDepth=%d)\n", globalTreeDepth))
}

	if originalTable == nil {
		return fmt.Errorf("%s originalTable.%s: originalTable is <nil>", UtilFuncSource(), UtilFuncName())
	}

	if table == nil {
		return fmt.Errorf("%s table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
	}

	if depth < 0 {
		return fmt.Errorf("%s table.%s: depth %d is negative", UtilFuncSource(), UtilFuncCaller(), depth)
	}

	if depth <= 0 {
if CURRENT_TESTS_VERBOSE {
		where()
		fmt.Printf("newTreeTable_recursive depth %d <= 0 return\n", depth)
}
		return nil
	}

//where(fmt.Sprintf("BEFORE processing isTableCol table:\n%s", table.StringNested()))
	// where(fmt.Sprintf("for rowIndex := 0; rowIndex < table.RowCount()=%d; rowIndex++", table.RowCount()))
	for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {	// Row Major order.
		for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
			isTableCol, err := table.IsTableColByColIndex(colIndex)
			if err != nil {
				return fmt.Errorf("%s %v", UtilFuncSource(), err)
			}

//where("\n" + table.String())
			if isTableCol {
	// where(fmt.Sprintf("cell.ColType == \"*Table\""))
				nestedTable, err := table.GetTableByColIndex(colIndex, rowIndex)
				if err != nil {
					return fmt.Errorf("%s %v", UtilFuncSource(), err)
				}

				// Replace only NilTable cell entries, and not cells that are already populated.
				isNilTable, err := nestedTable.IsNilTable()
				if err != nil {
					return fmt.Errorf("%s %v", UtilFuncSource(), err)
				}
				if isNilTable {
if CURRENT_TESTS_VERBOSE {
where("isNilTable")
}
					colName, err := table.ColName(colIndex)
					if err != nil {
						return fmt.Errorf("%s %v", UtilFuncSource(), err)
					}
if CURRENT_TESTS_VERBOSE {
					where(fmt.Sprintf("table [%s] colName %s rowIndex %d\n", table.Name(), colName, rowIndex))
					where(fmt.Sprintf("\noriginalTable: >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n"))
					where(fmt.Sprintf("%s", originalTable.StringNested()))
					where(fmt.Sprintf("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n\n"))
}

					// Copy the originalTable into this cell.
if CURRENT_TESTS_VERBOSE {
where("Copy the originalTable into this cell.")
where(fmt.Sprintf("[%s] col=%s row=%d IsNilTable()", table.Name(), colName, rowIndex))
}
					tableCopy, err := originalTable.Copy()
					if err != nil {
						return fmt.Errorf("%s %v", UtilFuncSource(), err)
					}
//where(fmt.Sprintf("tableCopy:>>>>>>\n%s\n<<<<<<", tableCopy.StringNested()))
//	
//						tableCopyName := tableCopy.Name()
//	//					tableCopy.SetName(tableCopyName + "_" + string('a' + cellCount-1))
//						newTableName := fmt.Sprintf("%s%c", tableCopyName, rune('A' + depth-1))
//	where(fmt.Sprintf("newTableName: %s", newTableName))
//						err = tableCopy.SetName(newTableName)
//						if err != nil {
//							return fmt.Errorf("%s %v", UtilFuncSource(), err)
//						}
//	
//						newColName := fmt.Sprintf("%s%02d", colName, depth)
//	where(fmt.Sprintf("newColName: %s", newColName))
//						err = tableCopy.RenameCol(colName, newColName)
//						if err != nil {
//							return fmt.Errorf("%s %v", UtilFuncSource(), err)
//						}
//	

//where(fmt.Sprintf("BEFORE *originalTable:\n%s", (*originalTable).StringNested()))
//where(fmt.Sprintf("BEFORE *table:\n%s", table.StringNested()))
//					err = (*originalTable).SetTable(colName, rowIndex, tableCopy)
					err = table.SetTable(colName, rowIndex, tableCopy)
					if err != nil {
						return fmt.Errorf("%s %v", UtilFuncSource(), err)
					}
//where(fmt.Sprintf("AFTER  *originalTable:\n%s", (*originalTable).StringNested()))
//where(fmt.Sprintf("AFTER  *table:\n%s", table.StringNested()))
				
					if depth >= 1 {
if CURRENT_TESTS_VERBOSE {
						where()
						fmt.Printf("calling newTreeTable_recursive(%d)\n", depth-1)
}
						err = newTreeTable_recursive(originalTable, nestedTable, depth-1)
						if err != nil {
							return fmt.Errorf("%s %v", UtilFuncSource(), err)
						}
					}
				} else {
if CURRENT_TESTS_VERBOSE {
					where(fmt.Sprintf("not a NilTable: [%s]", nestedTable.Name()))
}
				}
			}
		}
	}

//where("return")
	return
}
