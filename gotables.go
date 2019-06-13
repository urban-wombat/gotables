// Copyright (c) 2017 Malcolm Gorman

// Golang tabular data format for configs and channels, with a rich set of helper functions.

package gotables

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/urban-wombat/util"
)

//	"golang.org/x/text/width"
// Would like to be able to call width.Narrow() on wide glyphs.
//	Seems unavailable: http://docs.activestate.com/activego/1.8/pkg/golang.org/x/text/width
//	                   https://godoc.org/golang.org/x/text/width#example-Transformer--Narrow

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

const debugging bool = false
const printcallers = false
const printstack bool = false
const todo bool = false

func init() {
	/*
		if debugging {
			log.SetFlags(log.Lshortfile)
			log.SetOutput(os.Stderr)
		} else {
			log.SetOutput(ioutil.Discard)
		}
	*/
	log.SetFlags(log.Lshortfile)
	//	log.SetOutput(os.Stderr)
}

var where = log.Print

/*
#####################################################################################
TableSet
#####################################################################################
2016.12.16  Malcolm Gorman  Use bytes.Buffer to construct string() string strings.
#####################################################################################
*/

// ########
// TableSet
// ########

/*
TableSet is an ordered set of *Table pointers.
*/
type TableSet struct {
	tableSetName string
	fileName     string
	tables       []*Table
}

// Selected header information for exporting.
type TableSetExported struct {
	TableSetName string
	FileName     string
}

// Factory function to return an initialised *TableSet pointer.
func NewTableSet(tableSetName string) (*TableSet, error) {
	var newTables *TableSet = new(TableSet)
	newTables.tableSetName = tableSetName
	newTables.tables = []*Table{} // An empty slice of tables.
	return newTables, nil
}

// Read and parse a gotables file into a TableSet.
func NewTableSetFromFile(fileName string) (*TableSet, error) {
	var p parser
	p.SetFileName(fileName) // Needed for printing file and line diagnostics.

	tables, err := p.parseFile(fileName)
	if err != nil {
		return nil, err
	}

	return tables, nil
}

// Write a TableSet to a text file.
func (tableSet *TableSet) WriteFile(fileName string, mode os.FileMode) error {
	if tableSet == nil {
		return fmt.Errorf("%s ERROR: tableSet.%s tableSet is <nil>", util.FuncSource(), util.FuncName())
	}

	var err error
	var tableSet_String string
	var tableSet_Bytes []byte

	tableSet_String = tableSet.String()
	tableSet_Bytes = []byte(tableSet_String)
	if mode == 0 {
		// No permissions set.
		mode = 0666
	}
	err = ioutil.WriteFile(fileName, tableSet_Bytes, mode)

	return err
}

// Write a Table to a text file.
func (table *Table) WriteFile(fileName string, mode os.FileMode) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s(%q, mode) table is <nil>", util.FuncSource(), util.FuncName(), fileName)
	}
	var err error
	var table_String string
	var table_Bytes []byte

	table_String = table.String()
	table_Bytes = []byte(table_String)
	if mode == 0 { // No permissions set.
		mode = 0666
	}
	err = ioutil.WriteFile(fileName, table_Bytes, mode)

	return err
}

// Read and parse a gotables string into a TableSet.
func NewTableSetFromString(s string) (*TableSet, error) {
	var p parser

	tables, err := p.parseString(s)
	if err != nil {
		return nil, err
	}

	return tables, nil
}

/*
	This function expects exactly ONE table in the string. Otherwise it's an error.
	If there's more than one table in the string, use NewTableFromStringByTableName() instead.
*/
func NewTableFromString(s string) (*Table, error) {
	tableSet, err := NewTableSetFromString(s)
	if err != nil {
		return nil, err
	}

	tableCount := tableSet.TableCount()
	if tableCount != 1 {
		return nil, fmt.Errorf("%s expecting string to contain 1 table but found %d table%s", util.FuncName(), tableCount, plural(tableCount))
	}

	table, err := tableSet.TableByTableIndex(0)
	if err != nil {
		return nil, err
	}

	return table, nil
}

func NewTableFromStringByTableName(s string, tableName string) (*Table, error) {
	tableSet, err := NewTableSetFromString(s)
	if err != nil {
		return nil, err
	}

	table, err := tableSet.Table(tableName)
	if err != nil {
		return nil, err
	}

	return table, nil
}

/*
	This function expects exactly ONE table in the file. Otherwise it's an error.
	If there's more than one table in the file, use NewTableFromFileByTableName() instead.
*/
func NewTableFromFile(fileName string) (*Table, error) {
	tableSet, err := NewTableSetFromFile(fileName)
	if err != nil {
		return nil, err
	}

	tableCount := tableSet.TableCount()
	if tableCount != 1 {
		return nil, fmt.Errorf("%s expecting file to contain 1 table but found %d table%s: %s",
			util.FuncName(), tableCount, plural(tableCount), fileName)
	}

	table, err := tableSet.TableByTableIndex(0)
	if err != nil {
		return nil, err
	}

	return table, nil
}

func NewTableFromFileByTableName(fileName string, tableName string) (*Table, error) {
	tableSet, err := NewTableSetFromFile(fileName)
	if err != nil {
		return nil, err
	}

	table, err := tableSet.Table(tableName)
	if err != nil {
		return nil, err
	}

	return table, nil
}

/*
Returns a set of parsable tables with format right-aligned (numbers) as a string.
*/
func (tableSet *TableSet) String() string {
	if tableSet == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: tableSet.%s tableSet is <nil>\n", util.FuncSource(), util.FuncName()))
		return ""
	}
	return tableSet.StringPadded()
}

func (tableSet *TableSet) StringPadded() string {
	if tableSet == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: tableSet.%s tableSet is <nil>\n", util.FuncSource(), util.FuncName()))
		return ""
	}
	var verticalSep string = ""
	var s string

	var tableSetName string = tableSet.Name()
	if tableSetName != "" {
		s += fmt.Sprintf("# %s\n\n", tableSet.Name())
	}

	var table *Table
	for i := 0; i < len(tableSet.tables); i++ {
		table = tableSet.tables[i]
		s += verticalSep
		s += table.StringPadded()
		verticalSep = "\n"
	}
	return s
}

func (tableSet *TableSet) StringUnpadded() string {
	if tableSet == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: tableSet.%s tableSet is <nil>\n", util.FuncSource(), util.FuncName()))
		return ""
	}

	var horizontalSeparator byte = ' '
	return tableSet._String(horizontalSeparator)
}

// Return parsable set of tables as a string.
func (tableSet *TableSet) _String(horizontalSeparator byte) string {
	if tableSet == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: tableSet.%s tableSet is <nil>\n", util.FuncSource(), util.FuncName()))
		return ""
	}

	var buf bytes.Buffer
	//	buf.WriteString("# From file: \"" + tableSet.name + "\"\n\n")
	var tableSep = ""
	var table *Table
	for i := 0; i < len(tableSet.tables); i++ {
		table = tableSet.tables[i]
		buf.WriteString(tableSep)
		//		buf.WriteString(fmt.Sprintf("%v", table))
		buf.WriteString(fmt.Sprintf("%v", table._String(horizontalSeparator)))
		tableSep = "\n"
	}

	var s string = buf.String()
	return s
}

func (tableSet *TableSet) Name() string {
	if tableSet == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: tableSet.%s tableSet is <nil>\n", util.FuncSource(), util.FuncName()))
		return ""
	}

	return tableSet.tableSetName
}

func (tableSet *TableSet) SetName(tableSetName string) {
	if tableSet == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: tableSet.%s tableSet is <nil>\n", util.FuncSource(), util.FuncName()))
		return
	}

	tableSet.tableSetName = tableSetName
}

// The file name if this TableSet has been created from a file. Otherwise ""
func (tableSet *TableSet) FileName() string {
	if tableSet == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: tableSet.%s tableSet is <nil>\n", util.FuncSource(), util.FuncName()))
		return ""
	}

	return tableSet.fileName
}

func (tableSet *TableSet) SetFileName(fileName string) {
	if tableSet == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: tableSet.%s tableSet is <nil>\n", util.FuncSource(), util.FuncName()))
		return
	}

	tableSet.fileName = fileName
}

func (tableSet *TableSet) TableCount() int {
	if tableSet == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: tableSet.%s tableSet is <nil>\n", util.FuncSource(), util.FuncName()))
		return -1
	}

	return len(tableSet.tables)
}

// Add a table to a table set.
func (tableSet *TableSet) AppendTable(newTable *Table) error {
	if tableSet == nil {
		return fmt.Errorf("%s ERROR: tableSet.%s tableSet is <nil>", util.FuncSource(), util.FuncName())
	}

	// Note: Could maintain a map in parallel for rapid lookup of table names.
	for _, existingTable := range tableSet.tables {
		if existingTable.Name() == newTable.Name() {
			return fmt.Errorf("table [%s] already exists: [%s]", newTable.tableName, newTable.tableName)
		}
	}

	tableSet.tables = append(tableSet.tables, newTable)

	return nil
}

// Checks whether table exists
func (tableSet *TableSet) HasTable(tableName string) (bool, error) {
	if tableSet == nil {
		return false, fmt.Errorf("%s ERROR: tableSet.%s tableSet is <nil>", util.FuncSource(), util.FuncName())
	}

	for _, table := range tableSet.tables {
		if table.Name() == tableName {
			return true, nil
		}
	}
	return false, fmt.Errorf("table [%s] does not exist", tableName)
}

func (tableSet *TableSet) Table(tableName string) (*Table, error) {
	if tableSet == nil {
		return nil, fmt.Errorf("%s ERROR: tableSet.%s tableSet is <nil>", util.FuncSource(), util.FuncName())
	}

	for _, table := range tableSet.tables {
		if table.Name() == tableName {
			return table, nil
		}
	}

	return nil, fmt.Errorf("table [%s] does not exist", tableName)
}

func (tableSet *TableSet) TableByTableIndex(tableIndex int) (*Table, error) {
	if tableSet == nil {
		return nil, fmt.Errorf("%s ERROR: tableSet.%s tableSet is <nil>", util.FuncSource(), util.FuncName())
	}

	if tableIndex < 0 || tableIndex > tableSet.TableCount()-1 {
		err := fmt.Errorf("in *TableSet with %d tables, table index %d does not exist",
			tableSet.TableCount(), tableIndex)
		return nil, err
	}

	return tableSet.tables[tableIndex], nil
}

/*
#####################################################################################
Table
#####################################################################################
2016.02.03  Malcolm Gorman  Add sort keys to Table.
#####################################################################################
*/

type Table struct {
	tableName   string
	colNames    []string
	colTypes    []string
	colNamesMap map[string]int // To look up a colNames index from a col name.
	rows        []tableRow     // new_model
	sortKeys    []sortKey
	structShape bool
}

type TableExported struct {
	TableName      string
	ColNames       []string
	ColTypes       []string
	ColNamesLookup map[string]int // To look up a colNames index from a col name.
	Rows           []tableRow     // new_model
	SortKeys       []SortKeyExported
	StructShape    bool
}

func (table *Table) getColTypes() []string {
	if table == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: table.%s: table is <nil>\n", util.FuncSource(), util.FuncName()))
		return nil
	}
	return table.colTypes
}

type tableRow []interface{}

// Note: Reimplement this as a slice of byte for each row and a master map and/or slice to track offset.

// Factory function to generate a *Table pointer.
/*
	var myTable *gotables.Table
	myTable, err = gotables.NewTable("My_Table")
	if err != nil {
		panic(err)
	}
*/
func NewTable(tableName string) (*Table, error) {
	var err error
	var newTable *Table = new(Table)
	err = newTable.SetName(tableName)
	if err != nil {
		return nil, err
	}
	newTable.colNames = []string{}
	newTable.colTypes = []string{}
	newTable.colNamesMap = map[string]int{}
	newTable.rows = []tableRow{}

	return newTable, nil
}

func newTableExported(tableName string) (*TableExported, error) {
	var err error
	var NewTableExported *TableExported = new(TableExported)
	err = NewTableExported.setTableNameExported(tableName)
	if err != nil {
		return nil, err
	}
	NewTableExported.ColNames = []string{}
	NewTableExported.ColTypes = []string{}
	NewTableExported.ColNamesLookup = map[string]int{}
	NewTableExported.Rows = []tableRow{}
	return NewTableExported, nil
}

/*
	table, err := gotables.NewTableFromMetadata("Moviegoers", []string{"Age", "Mothballs"}, []string{"int", "bool"})
*/
func NewTableFromMetadata(tableName string, colNames []string, colTypes []string) (*Table, error) {
	var newTable *Table
	var err error

	// Check for invalid input.
	if len(colNames) != len(colTypes) {
		return nil, fmt.Errorf("%s(colNames, colTypes) len(colNames)=%d != len(colTypes)=%d",
			util.FuncName(), len(colNames), len(colTypes))
	}

	newTable, err = NewTable(tableName)
	if err != nil {
		return nil, err
	}

	err = newTable.appendCols(colNames, colTypes)
	if err != nil {
		return nil, err
	}

	if debugging {
		_, err = newTable.IsValidTable()
		if err != nil {
			return nil, err
		}
	}

	return newTable, nil
}

/*
Add (append) a new blank row to this table. This does NOT initialise the cell values. They will be set to nil.

Note: This is used by the parser. Not for use by end-users.
*/
func (table *Table) appendRowOfNil() error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	return nil
}

// Note: Can append rows to an empty (no columns) table, and later append columns - but not for long!
func (table *Table) AppendRows(howMany int) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var err error

	if debugging {
		_, err = table.IsValidTable()
		if err != nil {
			return err
		}
	}

	if howMany < 0 {
		return fmt.Errorf("table [%s] AppendRows(%d) cannot append %d rows (must be %d or more)", table.Name(), howMany, howMany, howMany)
	}

	for i := 0; i < howMany; i++ {
		err = table.AppendRow()
		if err != nil {
			return err
		}
	}

	if debugging {
		_, err = table.IsValidTable()
		if err != nil {
			return err
		}
	}

	return nil
}

// All cells in the new added row will be set to their zero value, such as 0, "", or false.
// Note: Can append rows to an empty (no columns) table, and later append columns.
func (table *Table) AppendRow() error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var err error

	if debugging {
		_, err = table.IsValidTable()
		if err != nil {
			return err
		}
	}

	//	if printcallers { util.PrintCaller() }

	/*
		// This is an interesting consideration. It sounds right, but it might make things less flexible unnecessarily.
		if table.ColCount() == 0 {
			err = fmt.Errorf("in table [%s]: cannot add rows to a table that has no columns", table.Name())
			return err
		}
	*/

	// Note: function make() sets slice values to <nil> and NOT to their zero value.
	var newRow tableRow = make(tableRow, len(table.colNames))
	table.rows = append(table.rows, newRow)

	// table.SetRowCellsToZeroValue() sets cells to their zero value. Otherwise cells would be left set to <nil>.
	var rowIndex int
	rowIndex, _ = table.lastRowIndex()
	err = table.SetRowCellsToZeroValue(rowIndex)
	if err != nil {
		return err
	}

	if debugging {
		_, err = table.IsValidTable()
		if err != nil {
			return err
		}
	}

	return nil
}

// Set all float cells in this row to NaN. This is a convenience function to use NaN as a proxy for a missing value.
func (table *Table) SetRowFloatCellsToNaN(rowIndex int) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var err error
	var colType string
	for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
		colType, err = table.ColTypeByColIndex(colIndex)
		if err != nil {
			return err
		}
		switch colType {
		case "float32":
			err = table.SetFloat32ByColIndex(colIndex, rowIndex, float32(math.NaN()))
		case "float64":
			err = table.SetFloat64ByColIndex(colIndex, rowIndex, math.NaN())
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// Set all float cells in this table to NaN. This is a convenience function to use NaN as a proxy for a missing value.
func (table *Table) SetAllFloatCellsToNaN() error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}
	var err error

	for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
		err = table.SetRowFloatCellsToNaN(rowIndex)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
// Set all cells in this row to their zero value, such as 0, "", or false.
//func (table *Table) SetRowCellsToZeroValue(rowIndex int) error {
//	var err error
//
//	if table == nil { return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName()) }
//
//	for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
//		err = table.SetCellToZeroValueByColIndex(colIndex, rowIndex)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
*/

// Set all cells in this col to their zero value, such as 0, "", or false.
func (table *Table) SetColCellsToZeroValue(colName string) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}
	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}
	return table.SetColCellsToZeroValueByColIndex(colIndex)
}

// Set all cells in this col to their zero value, such as 0, "", or false.
func (table *Table) SetColCellsToZeroValueByColIndex(colIndex int) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}
	var err error

	for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
		err = table.SetCellToZeroValueByColIndex(colIndex, rowIndex)
		if err != nil {
			return err
		}
	}

	return nil
}

func (table *Table) SetCellToZeroValue(colName string, rowIndex int) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}
	var err error
	var colIndex int

	colIndex, err = table.ColIndex(colName)
	if err != nil {
		return err
	}

	err = table.SetCellToZeroValueByColIndex(colIndex, rowIndex)
	if err != nil {
		return err
	}

	return nil
}

func (table *Table) appendRowSlice(rowSlice tableRow) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	// We're going to assume that all error checking was done in getRowSlice()

	// Append row to existing rows.
	if debugging {
		// where(fmt.Sprintf("BEFORE: table.rows = %v\n", table.rows))
		// where(fmt.Sprintf("DURING: rowSlice = %v\n", rowSlice))
		// where(fmt.Sprintf("append(%v, %v)\n", table.rows, rowSlice))
	}
	table.rows = append(table.rows, rowSlice)
	if debugging {
		// where(fmt.Sprintf("AFTER: table.rows = %v\n", table.rows))
		// where(fmt.Sprintf("\n"))
	}

	return nil
}

func (table *Table) DeleteRow(rowIndex int) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var err error

	if debugging {
		_, err := table.IsValidTable()
		if err != nil {
			return err
		}
	}

	if rowIndex < 0 || rowIndex > table.RowCount()-1 {
		return fmt.Errorf("%s: in table [%s] with %d rows, row index %d does not exist",
			util.FuncName(), table.tableName, table.RowCount(), rowIndex)
	}

	err = table.DeleteRows(rowIndex, rowIndex)
	if err != nil {
		return nil
	}

	if debugging {
		_, err = table.IsValidTable()
		if err != nil {
			return err
		}
	}

	return nil
}

// If table has any rows, delete them all. This is to deal simply with empty tables.
func (table *Table) DeleteRowsAll() error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var err error

	if debugging {
		_, err = table.IsValidTable()
		if err != nil {
			return err
		}
	}

	if table.RowCount() > 0 {
		err = table.DeleteRows(0, table.RowCount()-1)
		if err != nil {
			return err
		}
	}

	if debugging {
		_, err = table.IsValidTable()
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete rows from firstRowIndex to lastRowIndex inclusive. This means lastRowIndex will be deleted.
func (table *Table) DeleteRows(firstRowIndex int, lastRowIndex int) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var err error

	if debugging {
		_, err := table.IsValidTable()
		if err != nil {
			debug.PrintStack()
		}
		if err != nil {
			return err
		}
	}

	if firstRowIndex < 0 || firstRowIndex > table.RowCount()-1 {
		if err != nil {
			debug.PrintStack()
		}
		return fmt.Errorf("%s: in table [%s] with %d rows, firstRowIndex %d does not exist",
			util.FuncName(), table.tableName, table.RowCount(), firstRowIndex)
	}

	if lastRowIndex < 0 || lastRowIndex > table.RowCount()-1 {
		if err != nil {
			debug.PrintStack()
		}
		return fmt.Errorf("%s: in table [%s] with %d rows, lastRowIndex %d does not exist",
			util.FuncName(), table.tableName, table.RowCount(), lastRowIndex)
	}

	if firstRowIndex > lastRowIndex {
		if err != nil {
			debug.PrintStack()
		}
		return fmt.Errorf("%s: invalid row index range: firstRowIndex %d > lastRowIndex %d", util.FuncName(), firstRowIndex, lastRowIndex)
	}

	if debugging {
		_, err = table.IsValidTable()
		if err != nil {
			debug.PrintStack()
			return err
		}
	}

	// From Ivo Balbaert p182 for deleting a range of elements from a slice.
	table.rows = append(table.rows[:firstRowIndex], table.rows[lastRowIndex+1:]...)

	if debugging {
		_, err = table.IsValidTable()
		if err != nil {
			debug.PrintStack()
			return err
		}
	}

	return nil
}

/*
Return a missing value for a type. The only types that have a good enough missing value are float32 and float64 with NaN.
*/
func missingValueForType(typeName string) (missingValue interface{}, hasMissing bool) {
	switch typeName {
	case "float32", "float64":
		missingValue = math.NaN()
	default:
		return nil, false
	}
	return missingValue, true
}

func (table *Table) StringUnpadded() string {

	if table == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: table.%s: table is <nil>\n", util.FuncSource(), util.FuncName()))
		return ""
	}

	return table._String(' ')
}

/*
Return a parsable table as a string. Intended for internal library use.
*/
func (table *Table) _String(horizontalSeparator byte) string {

	if table == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: table.%s: table is <nil>\n", util.FuncSource(), util.FuncName()))
		return ""
	}
	const tabForTabwriter = '\t'
	if horizontalSeparator == 0 { // Null char.
		horizontalSeparator = tabForTabwriter
	}
	var horizontalSep string
	const verticalSep byte = '\n'
	var s string
	var buf bytes.Buffer

	// Print as struct shape or table shape.
	if table.structShape && table.RowCount() <= 1 {
		s = printStruct(table)
	} else {
		// Table name
		buf.WriteByte('[')
		buf.WriteString(table.tableName)
		buf.WriteString("]\n")

		// Col names
		if len(table.colNames) > 0 {
			horizontalSep = ""
			for _, colName := range table.colNames {
				buf.WriteString(horizontalSep)
				buf.WriteString(colName)
				horizontalSep = string(horizontalSeparator)
			}
			buf.WriteByte(verticalSep)
		}

		// Col types
		if len(table.colTypes) > 0 {
			horizontalSep = ""
			for _, colType := range table.colTypes {
				buf.WriteString(horizontalSep)
				buf.WriteString(colType)
				horizontalSep = string(horizontalSeparator)
			}
			buf.WriteByte(verticalSep)
		}

		// Rows of data
		for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
			var row tableRow = table.rows[rowIndex]
			horizontalSep = ""
			for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
				var sVal string
				var runeVal rune
				var tVal bool
				var ui8Val uint8
				var ui16Val uint16
				var ui32Val uint32
				var ui64Val uint64
				var uiVal uint
				var iVal int
				var i8Val int8
				var i16Val int16
				var i32Val int32
				var i64Val int64
				var f32Val float32
				var f64Val float64
				buf.WriteString(horizontalSep)
				switch table.colTypes[colIndex] {
				case "string":
					sVal = row[colIndex].(string)
					var quoted string = strconv.Quote(sVal)
					buf.WriteString(quoted)
				case "rune":
					runeVal = row[colIndex].(rune)
					s = strconv.QuoteRune(runeVal)
					buf.WriteString(fmt.Sprintf("%s", s))
				case "bool":
					tVal = row[colIndex].(bool)
					buf.WriteString(fmt.Sprintf("%t", tVal))
				case "uint8":
					ui8Val = row[colIndex].(uint8)
					buf.WriteString(fmt.Sprintf("%d", ui8Val))
				case "uint16":
					ui16Val = row[colIndex].(uint16)
					buf.WriteString(fmt.Sprintf("%d", ui16Val))
				case "uint32":
					ui32Val = row[colIndex].(uint32)
					buf.WriteString(fmt.Sprintf("%d", ui32Val))
				case "uint64":
					ui64Val = row[colIndex].(uint64)
					buf.WriteString(fmt.Sprintf("%d", ui64Val))
				case "uint":
					uiVal = row[colIndex].(uint)
					buf.WriteString(fmt.Sprintf("%d", uiVal))
				case "int":
					iVal = row[colIndex].(int)
					buf.WriteString(fmt.Sprintf("%d", iVal))
				case "int8":
					i8Val = row[colIndex].(int8)
					buf.WriteString(fmt.Sprintf("%d", i8Val))
				case "int16":
					i16Val = row[colIndex].(int16)
					buf.WriteString(fmt.Sprintf("%d", i16Val))
				case "int32":
					i32Val = row[colIndex].(int32)
					buf.WriteString(fmt.Sprintf("%d", i32Val))
				case "int64":
					i64Val = row[colIndex].(int64)
					buf.WriteString(fmt.Sprintf("%d", i64Val))
				case "float32":
					f32Val = row[colIndex].(float32)
					var f64ValForFormatFloat float64 = float64(f32Val)
					buf.WriteString(strconv.FormatFloat(f64ValForFormatFloat, 'f', -1, 32)) // -1 strips off excess decimal places.
				case "float64":
					f64Val = row[colIndex].(float64)
					buf.WriteString(strconv.FormatFloat(f64Val, 'f', -1, 64)) // -1 strips off excess decimal places.
				default:
					log.Printf("%s #1 ERROR IN %s: Unknown type: %s\n", util.FuncSource(), util.FuncName(), table.colTypes[colIndex])
					return ""
				}

				horizontalSep = string(horizontalSeparator)
			}
			buf.WriteByte(verticalSep)
		}

		s = buf.String()
	}

	return s
}

// For int type.
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// 18.01.2017 M Gorman
// This prints a gotable in tabular format. Called indirectly by table.String()
func printMatrix(tableName string, matrix [][]string, width []int, precis []int, alignRight []bool, colTypes []string) string {
	var buf bytes.Buffer
	var s string
	var sep string // Printed before each value.

	s = fmt.Sprintf("[%s]\n", tableName)
	buf.WriteString(s)

	// Avoid out of bounds error accessing matrix[0]
	if len(matrix) == 0 {
		return buf.String()
	}

	var rightmostCol int = len(matrix) - 1

	for row := 0; row < len(matrix[0]); row++ {
		sep = "" // No separator before first column.
		for col := 0; col < len(matrix); col++ {
			if alignRight[col] {
				// TODO: Move this functionality to where printMatrix is called.
				var toWrite string
				if row <= 1 { // It's a colName or typeName
					toWrite = matrix[col][row]
				} else { // It's numeric. Note: float conversion handles int conversion.
					var bits int
					switch colTypes[col] {
					case "float32":
						bits = 32
					case "float64":
						bits = 64
					default:
						bits = 64 // For int and other non-float integrals.
					}
					// Convert back to float so we can format it again in light of the maximum precision in the column.
					float64Val, err := strconv.ParseFloat(matrix[col][row], bits)
					// TODO: Remove this panic.
					if err != nil {
						panic(err)
					}
					toWrite = strconv.FormatFloat(float64Val, 'f', precis[col], bits)
					//					width[col] = max(width[col], len(toWrite))
				}
				//				s = fmt.Sprintf("%s%*s", sep, width[col], matrix[col][row])	// Align right
				if colTypes[col] == "float32" || colTypes[col] == "float64" {
					// Replace trailing zeros with space padding here.
					// The padding is to ensure the next column to the right is aligned along a straight edge.
					toWrite = padTrailingZeros(toWrite)
				}
				s = fmt.Sprintf("%s%*s", sep, width[col], toWrite) // Align right
				if col == rightmostCol {
					// Padding to the right of the rightmost column is unnecessary. Remove it here.
					// Remove any jagged space padding to the right of decimal point.
					s = strings.TrimRight(s, " ")
				}
				buf.WriteString(s)
			} else { // Left-aligned col. Cells in non-numeric cols are treated as left-aligned, eg string and bool.
				if col == rightmostCol {
					// Don't pad (unnecessarily) to the right of rightmost col if it is left-aligned.
					s = fmt.Sprintf("%s%s", sep, matrix[col][row]) // With no padding, doesn't need align left with -
				} else {
					s = fmt.Sprintf("%s%-*s", sep, width[col], matrix[col][row]) // Align left with -
				}
				buf.WriteString(s)
			}
			sep = " " // Separator before subsequent columns.
		}
		s = fmt.Sprintln()
		buf.WriteString(s)
	}

	return buf.String()
}

/*
Return a parsable table as a string with numbers format aligned right.
*/
func (table *Table) String() string {
	if table == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: table.%s: table is <nil>\n", util.FuncSource(), util.FuncName()))
		return ""
	}
	return table.StringPadded()
}

func (table *Table) StringPadded() string {
	var err error

	if table == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: table.%s: table is <nil>\n", util.FuncSource(), util.FuncName()))
		return ""
	}

	//	var horizontalSeparator byte = ' '	// Remove this later?
	var gapBetweenCols string = " "
	var horizontalSep string
	const verticalSep byte = '\n'
	const colNameRowIndex = 0
	const colTypeRowIndex = 1
	const headingRows = 2 // names row plus types row
	var s string
	var buf bytes.Buffer

	// Table name
	buf.WriteByte('[')
	buf.WriteString(table.tableName)
	buf.WriteString("]\n")

	// Write table headings, types and column values to parallel slices.
	// Keep track of the widest entry in each slice.
	var colCount int = table.ColCount()          // For efficiency
	var rowCountPlus2 int = table.RowCount() + 2 // Allow for col name and type.

	alignRight := make([]bool, colCount)

	width := make([]int, colCount)
	prenum := make([]int, colCount)
	points := make([]int, colCount)
	precis := make([]int, colCount)

	/*
		// Special case to align decimal points in float32 and float64.
		widthBeforePoint := make([]int, colCount)
		widthAfterPoint := make([]int, colCount)
	*/

	matrix := make([][]string, colCount)
	for colIndex := 0; colIndex < colCount; colIndex++ {
		matrix[colIndex] = make([]string, rowCountPlus2)
	}

	// Col names
	// Initialise width to width of colName.
	if len(table.colNames) > 0 { // Allow for empty table?
		for colIndex, colName := range table.colNames {
			matrix[colIndex][colNameRowIndex] = colName
			width[colIndex] = max(width[colIndex], len(colName))
		}
	}

	// Col types
	// Stretch width if colType is wider than colName.
	// Set alignRight true if col is numeric.
	if len(table.colTypes) > 0 { // Allow for empty table?
		for colIndex, colType := range table.colTypes {
			matrix[colIndex][colTypeRowIndex] = colType
			width[colIndex] = max(width[colIndex], len(colType))
			alignRight[colIndex], _ = IsNumericColType(colType)
		}
	}

	// Rows of data
	var row tableRow
	for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
		row = table.rows[rowIndex]
		horizontalSep = "" // No gap on left of first column.
		for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
			var sVal string
			var tVal bool
			var ui8Val uint8
			var ui8SliceVal []uint8
			var byteSliceVal []byte
			var ui16Val uint16
			var ui32Val uint32
			var ui64Val uint64
			var uiVal uint
			var iVal int
			var i8Val int8
			var i16Val int16
			var i32Val int32
			var runeVal rune
			var i64Val int64
			var f32Val float32
			var f64Val float64
			var iFaceVal interface{}
			buf.WriteString(horizontalSep)
			switch table.colTypes[colIndex] {
			case "string":
				sVal = row[colIndex].(string)
				s = strconv.Quote(sVal)
			case "bool":
				tVal = row[colIndex].(bool)
				s = fmt.Sprintf("%t", tVal)
			case "uint8", "byte":
				ui8Val = row[colIndex].(uint8)
				s = fmt.Sprintf("%d", ui8Val)
			case "[]uint8":
				ui8SliceVal = row[colIndex].([]uint8)
				s = fmt.Sprintf("%v", ui8SliceVal)
			case "[]byte":
				byteSliceVal = row[colIndex].([]byte)
				s = fmt.Sprintf("%v", byteSliceVal)
			case "uint16":
				ui16Val = row[colIndex].(uint16)
				s = fmt.Sprintf("%d", ui16Val)
			case "uint32":
				ui32Val = row[colIndex].(uint32)
				s = fmt.Sprintf("%d", ui32Val)
			case "uint64":
				ui64Val = row[colIndex].(uint64)
				s = fmt.Sprintf("%d", ui64Val)
			case "uint":
				uiVal = row[colIndex].(uint)
				s = fmt.Sprintf("%d", uiVal)
			case "int":
				iVal = row[colIndex].(int)
				s = fmt.Sprintf("%d", iVal)
			case "int8":
				i8Val = row[colIndex].(int8)
				s = fmt.Sprintf("%d", i8Val)
			case "int16":
				i16Val = row[colIndex].(int16)
				s = fmt.Sprintf("%d", i16Val)
			case "int32":
				i32Val = row[colIndex].(int32)
				s = fmt.Sprintf("%d", i32Val)
			case "rune":
				runeVal = row[colIndex].(rune)
				s = strconv.QuoteRune(runeVal)
				/*
					if len(s) >= 5 {
						s = width.Narrow(s)	// This package seems to be unavailable.
					}
				*/
				setWidths(s, colIndex, prenum, points, precis, width)
			case "int64":
				i64Val = row[colIndex].(int64)
				s = fmt.Sprintf("%d", i64Val)
			case "float32":
				f32Val = row[colIndex].(float32)
				var f64ValForFormatFloat float64 = float64(f32Val)
				s = strconv.FormatFloat(f64ValForFormatFloat, 'f', -1, 32) // -1 strips off excess decimal places.
				setWidths(s, colIndex, prenum, points, precis, width)
			case "float64":
				f64Val = row[colIndex].(float64)
				s = strconv.FormatFloat(f64Val, 'f', -1, 64) // -1 strips off excess decimal places.
				setWidths(s, colIndex, prenum, points, precis, width)
			default:
				// Is a user-defined interface value.
				iFaceVal = row[colIndex]
				s, err = interfaceValAsString(iFaceVal)
				if err != nil {
					log.Printf("#2 %s ERROR IN %s: %v\n", util.FuncSource(), util.FuncName(), err)
					return ""
				}
				setWidths(s, colIndex, prenum, points, precis, width)
/*
				log.Printf("#2 %s ERROR IN %s: Unknown type: %s\n", util.FuncSource(), util.FuncName(), table.colTypes[colIndex])
				return ""
*/
			}
			matrix[colIndex][headingRows+rowIndex] = s

			/*	MAL WAS HERE 14.02.2017
				Accumulate the widest number before the decimal point.
				Accumulate the widest number after the decimal point.
				Width is then the widest before plus decimal point plus widest after.
				The problem is handling floats with no decimal points.
				And what if some in a column have points and others don't?
				printMatrix() will have to use a format which places the decimal point in a uniform location. Tricky.
							fmt.Printf("table.colTypes[%d] = %s\n", colIndex, table.colTypes[colIndex])
							if table.colTypes[colIndex] == "float32" || table.colTypes[colIndex] == "float64" {
								fmt.Printf("s = %s\n", s)
								widthBeforePoint[colIndex] = len(strings.Split(s, ".")[0])
								fmt.Printf("widthBeforePoint[%d] = %d s=%s\n", colIndex, widthBeforePoint[colIndex], s)
								widthAfterPoint[colIndex] = len(strings.Split(s, ".")[1])
								fmt.Printf("widthAfterPoint[%d] = %d s=%s\n", colIndex, widthAfterPoint[colIndex], s)
								// os.Exit(44)
							}
			*/

			width[colIndex] = max(width[colIndex], len(s)) // Needed for non-numeric columns.
			horizontalSep = gapBetweenCols
		}
		buf.WriteByte(verticalSep)
	}

	// Print as struct shape or table shape.
	if table.structShape && table.RowCount() <= 1 {
		s = printStruct(table)
	} else {
		s = printMatrix(table.tableName, matrix, width, precis, alignRight, table.colTypes)
	}

	return s
}

func printStruct(table *Table) string {
	if table == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: table.%s: table is <nil>\n", util.FuncSource(), util.FuncName()))
	}

	var err error
	var asString string
	var s string
	var structHasRowData bool = table.RowCount() > 0

	s = fmt.Sprintf("[%s]\n", table.tableName)
	for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
		s += table.colNames[colIndex] + " " + table.colTypes[colIndex]
		if structHasRowData {
			const RowIndexZero = 0
			asString, err = table.GetValAsStringByColIndex(colIndex, RowIndexZero)
			if err != nil {
				_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: %s: %s\n", util.FuncSource(), util.FuncName(), err))
			}

			switch table.colTypes[colIndex] {
			case "string":
				// Note 1: GetValAsStringByColIndex() doesn't include delimiters around strings.
				// Note 2: Don't use %q. Use \"%s\" instead. Because %q replicates escape slashes.
				s += " = " + fmt.Sprintf("\"%s\"", asString)
			case "rune":
				// Note: GetValAsStringByColIndex() doesn't include delimiters around runes.
				s += " = " + fmt.Sprintf("'%s'", asString)
			default:
				// All other types don't have delimiters.
				s += " = " + asString
			}

		}
		s += "\n"
	}

	// where(s)

	return s
}

// How many chars before the decimal point (if any decimal point) does this string have?
// If no decimal point, that implies: the number of chars in the entire string.
// Pretends that there is a decimal point to the right of the string.
func preNumberOf(s string) (prenumber int) {
	index := strings.Index(s, ".")
	if index >= 0 {
		prenumber = index
	} else {
		prenumber = len(s)
	}
	return prenumber
}

// How many decimal points (zero or one) does this string have?
func pointsOf(s string) (points int) {
	index := strings.Index(s, ".")
	if index >= 0 {
		points = 1
	} else {
		points = 0
	}
	return points
}

func precisionOf(s string) (precision int) {
	index := strings.Index(s, ".")
	if index >= 0 {
		precision = (len(s) - index) - 1
	} else {
		precision = 0
	}
	return precision
}

func setWidths(s string, colIndex int, prenum []int, points []int, precis []int, width []int) {
	prenum[colIndex] = max(prenum[colIndex], preNumberOf(s))
	points[colIndex] = max(points[colIndex], pointsOf(s))
	precis[colIndex] = max(precis[colIndex], precisionOf(s))
	thisWidth := prenum[colIndex] + points[colIndex] + precis[colIndex]
	width[colIndex] = max(width[colIndex], thisWidth)
}

/*
Return a table as a comma separated variables for spreadsheets.

optionalSubstituteHeadingNames: Leave empty or provide a []string of names of length table.ColCount()

See: https://en.wikipedia.org/wiki/Comma-separated_values
*/
func (table *Table) GetTableAsCSV(optionalSubstituteHeadingNames ...string) (string, error) {
	var err error

	if table == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var buf *bytes.Buffer = new(bytes.Buffer)
	csvWriter := csv.NewWriter(buf)
	var record []string = make([]string, table.ColCount())

	// Col names
	if len(optionalSubstituteHeadingNames) > 0 {
		// Use user-provided headings.
		if len(optionalSubstituteHeadingNames) != table.ColCount() {
			return "", fmt.Errorf("[%s].%s(optionalSubstituteHeadingNames): expecting %d []optionalSubstituteHeadingNames, not %d",
				table.Name(), util.FuncName(), table.ColCount(), len(optionalSubstituteHeadingNames))
		}
		// should use copy() instead of a loop (S1001)
		for colIndex, substituteHeadingName := range optionalSubstituteHeadingNames {
			if len(optionalSubstituteHeadingNames[colIndex]) > 0 {
				// Use substitute heading name.
				record[colIndex] = substituteHeadingName
			} else {
				// Default to existing heading name.
				colName, err := table.ColNameByColIndex(colIndex)
				if err != nil {
					return "", err
				}
				record[colIndex] = colName
			}
		}
	} else {
		// Be default use table col names.
		// should use copy() instead of a loop (S1001)
		for colIndex, colName := range table.colNames {
			record[colIndex] = colName
		}
	}
	err = csvWriter.Write(record)
	if err != nil {
		return "", err
	}

	// Rows of data
	for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
		for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
			var sVal string
			sVal, err = table.GetValAsStringByColIndex(colIndex, rowIndex)
			if err != nil {
				return "", nil
			}

			// Handle special case of NaN which needs to be written as "" which is ,, in CSV.
			switch table.colTypes[colIndex] {
			case "float32", "float64":
				if sVal == "NaN" {
					record[colIndex] = "" // Empty value.
				} else {
					record[colIndex] = sVal
				}
			default:
				record[colIndex] = sVal // All the other types.
			}
		}

		err = csvWriter.Write(record)
		if err != nil {
			return "", err
		}
	}

	csvWriter.Flush()
	err = csvWriter.Error()
	if err != nil {
		return "", err
	}

	var s string = buf.String()

	return s, nil
}

// Append a column to this table.
/*
	err = myTable.AppendCol(headingName, headingType)
	if err != nil {
		panic(err)
	}
*/
func (table *Table) AppendCol(colName string, colType string) error {

	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	if isValid, err := IsValidColName(colName); !isValid {
		return err
	}

	if isValid, err := IsValidColType(colType); !isValid {
		return err
	}

	// Make sure this col name doesn't already exist.
	_, exists := table.colNamesMap[colName]
	if exists {
		err := fmt.Errorf("table [%s] col already exists: %s", table.Name(), colName)
		return err
	}

	table.colNames = append(table.colNames, colName)
	table.colTypes = append(table.colTypes, colType)

	colIndex := len(table.colNames) - 1
	table.colNamesMap[colName] = colIndex

	// Extend each row by 1 element. The new element will default to a zero value.
	for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
		table.rows[rowIndex] = append(table.rows[rowIndex], nil)
	}

	err := table.SetColCellsToZeroValue(colName)
	if err != nil {
		return err
	}

	return nil
}

func (table *Table) DeleteColByColIndex(colIndex int) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	if colIndex < 0 || colIndex > table.ColCount()-1 {
		err := fmt.Errorf("%s: in table [%s] with %d cols, col index %d does not exist",
			util.FuncName(), table.tableName, table.ColCount(), colIndex)
		return err
	}

	colName, err := table.ColName(colIndex)
	if err != nil {
		return err
	}

	// From Ivo Balbaert p182 for deleting a single element from a slice.
	table.colNames = append(table.colNames[:colIndex], table.colNames[colIndex+1:]...)

	// Delete colName from colNamesMap. This has to come AFTER deleting the colName from colNames.
	delete(table.colNamesMap, colName)
	for localColIndex := 0; localColIndex < len(table.colNames); localColIndex++ {
		table.colNamesMap[table.colNames[localColIndex]] = localColIndex
	}

	// From Ivo Balbaert p182 for deleting a single element from a slice.
	table.colTypes = append(table.colTypes[:colIndex], table.colTypes[colIndex+1:]...)

	for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
		//		if isValidRow, err := table.IsValidRow(rowIndex); !isValidRow { where(fmt.Sprintf("%s\n", err)) }
		var row []interface{} = table.rows[rowIndex]
		// From Ivo Balbaert p182 for deleting a single element from a slice.
		row = append(row[:colIndex], row[colIndex+1:]...)
		table.rows[rowIndex] = row
		//		if isValidRow, err := table.IsValidRow(rowIndex); !isValidRow { where(fmt.Sprintf("%s\n", err)) }
	}

	return nil
}

func (table *Table) DeleteCol(colName string) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	return table.DeleteColByColIndex(colIndex)
}

// Requires a val of valid type for the col in the table.
func (table *Table) SetVal(colName string, rowIndex int, val interface{}) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	if printcallers {
		util.PrintCaller()
	}

	// Removing unnecessary checking led to %40 speedup.

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: inlining SetValByColIndex() into SetVal() makes no noticeable difference.
	err = table.SetValByColIndex(colIndex, rowIndex, val)
	if err != nil {
		return err
	}

	return nil
}

// Requires a val of valid type for the col in the table.
func (table *Table) SetValByColIndex(colIndex int, rowIndex int, val interface{}) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	if printcallers {
		util.PrintCaller()
	}

	// No need to call HasCellByColIndex(). Col checked below.
	// No measurable speedup.
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return err
	}
	valType := fmt.Sprintf("%T", val)
	if valType != colType {
		if !isAlias(colType, valType) {
			colName := table.colNames[colIndex]
			return fmt.Errorf("%s: table [%s] col index %d col name %s expecting type %s not type %s: %v",
				util.FuncName(), table.Name(), colIndex, colName, colType, valType, val)
		}
	}

	// Set the val
	table.rows[rowIndex][colIndex] = val

	return nil
}

/*
	This is to avoid use of appendColNames() and appendColTypes() in parseString().
*/
func (table *Table) appendCols(colNames []string, colTypes []string) error {

	// Check for invalid input.
	if len(colNames) != len(colTypes) {
		return fmt.Errorf("%s(colNames, colTypes) len(colNames)=%d != len(colTypes)=%d",
			util.FuncName(), len(colNames), len(colTypes))
	}

	for colIndex := 0; colIndex < len(colNames); colIndex++ {
		err := table.AppendCol(colNames[colIndex], colTypes[colIndex])
		if err != nil {
			return err
		}
	}

	return nil
}

type colInfoStruct struct {
	colName string
	colType string
}

// Checks whether col exists
func (table *Table) HasCol(colName string) (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	if printcallers {
		util.PrintCaller()
	}

	_, err := table.getColInfo(colName)
	var exists bool = err == nil
	return exists, err
}

// Checks whether col exists
func (table *Table) HasColByColIndex(colIndex int) (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	if colIndex < 0 || colIndex > table.ColCount()-1 {
		err := fmt.Errorf("%s: in table [%s] with %d col%s, col index %d does not exist",
			util.FuncName(), table.tableName, table.ColCount(), plural(table.ColCount()), colIndex)
		return false, err
	}

	return true, nil
}

func (table *Table) getColInfo(colName string) (colInfoStruct, error) {
	var cInfo colInfoStruct
	if table == nil {
		return cInfo, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}
	var index int
	var exists bool
	index, exists = table.colNamesMap[colName]
	if !exists {
		err := fmt.Errorf("table [%s] col does not exist: %s", table.tableName, colName)
		return cInfo, err
	}
	cInfo.colName = colName
	cInfo.colType = table.colTypes[index]
	return cInfo, nil
}

func (table *Table) GetColInfoAsTable() (*Table, error) {
	if table == nil {
		return nil, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}
	var err error
	var colsTable *Table

	colsTable, err = NewTable("ColInfo")
	if err != nil {
		return nil, err
	}

	if err = colsTable.AppendCol("index", "int"); err != nil {
		return nil, err
	}
	if err = colsTable.AppendCol("colName", "string"); err != nil {
		return nil, err
	}
	if err = colsTable.AppendCol("colType", "string"); err != nil {
		return nil, err
	}

	for colIndex := 0; colIndex < table.ColCount(); colIndex++ {

		err = colsTable.AppendRow()
		if err != nil {
			return nil, err
		}

		rowIndex := colIndex // An output table row for each input table column.

		if err = colsTable.SetInt("index", rowIndex, rowIndex); err != nil {
			return nil, err
		}

		colName, err := table.ColName(colIndex)
		if err != nil {
			return nil, err
		}

		colInfo, err := table.getColInfo(colName)
		if err != nil {
			return nil, err
		}

		if err = colsTable.SetString("colName", rowIndex, colInfo.colName); err != nil {
			return nil, err
		}

		if err = colsTable.SetString("colType", rowIndex, colInfo.colType); err != nil {
			return nil, err
		}
	}

	return colsTable, nil
}

/*
	Return a slice of col names and a slice of col types:

	colNames []string
	colTypes []string
*/
func (table *Table) GetColInfoAsSlices() ([]string, []string, error) {
	if table == nil {
		return nil, nil, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var colNames []string = []string{}
	var colTypes []string = []string{}

	for colIndex := 0; colIndex < table.ColCount(); colIndex++ {

		colName, err := table.ColName(colIndex)
		if err != nil {
			return nil, nil, err
		}
		colNames = append(colNames, colName)

		colInfo, err := table.getColInfo(colName)
		if err != nil {
			return nil, nil, err
		}
		colTypes = append(colTypes, colInfo.colType)
	}

	return colNames, colTypes, nil
}

func (table *Table) ColType(colName string) (string, error) {
	if table == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	index, exists := table.colNamesMap[colName]
	if !exists {
		err := fmt.Errorf("table [%s] col does not exist: %s", table.tableName, colName)
		return "", err
	}
	var colType string = table.colTypes[index]
	return colType, nil
}

func (table *Table) ColTypeByColIndex(colIndex int) (string, error) {
	if table == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	if colIndex < 0 || colIndex > len(table.colTypes)-1 {
		err := fmt.Errorf("table [%s] col index does not exist: %d", table.tableName, colIndex)
		return "", err
	}
	var colType string = table.colTypes[colIndex]
	return colType, nil
}

func (table *Table) ColIndex(colName string) (int, error) {
	if table == nil {
		return -1, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	index, exists := table.colNamesMap[colName]
	if exists {
		return index, nil
	}

	err := fmt.Errorf("table [%s] col does not exist: %s", table.tableName, colName)
	return -1, err
}

/*
	Note: This will return -1 and an error if table.RowCount() == 0
	Safer to use table.RowCount() for looping.
	We might just remove LastRowIndex() from the library.
	We have made this private (23/07/2016)
*/
func (table *Table) lastRowIndex() (int, error) {
	if table == nil {
		return -1, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var err error
	var rowCount int = table.RowCount()
	if rowCount < 1 {
		err = fmt.Errorf("table [%s] has zero rows", table.Name())
		return -1, err
	}
	return table.RowCount() - 1, nil
}

func (table *Table) Name() string {
	if table == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR calling table.%s: table is <nil>\n", util.FuncSource(), util.FuncName()))
		return ""
	}
	return table.tableName
}

func (table *Table) ColCount() int {
	if table == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: table.%s: table is <nil>\n", util.FuncSource(), util.FuncName()))
		return -1
	}

	colNamesCount := len(table.colNames)

	return colNamesCount
}

/*
	Return the number of rows in this table.
	Returns -1 if there is an error (namely: the table variable is nil).
*/
func (table *Table) RowCount() int {
	if table == nil {
		return -1
	}

	return len(table.rows)
}

// Returns an interface{} value which may contain any valid gotables data type or NaN.
func (table *Table) GetVal(colName string, rowIndex int) (interface{}, error) {
	// Why don't we simply call GetValByColIndex() ???
	// Because old memory model makes it faster to look up colName than to lookup colIndex.

	if table == nil {
		return nil, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	if printcallers {
		util.PrintCaller()
	}

	var val interface{}

	// Sadly, slice doesn't return a boolean to test whether a retrieval is in range.
	/*	CHECKED IN table.GetValByColIndex()
	//	hasRow, err := table.HasRow(rowIndex)
	//	if !hasRow { return nil, err }
	*/

	/*	CHECKED IN table.GetValByColIndex()
	//	hasCol, err := table.HasCol(colName)
	//	if !hasCol { return nil, err }
	*/

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return nil, err
	}

	val, err = table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return nil, err
	}

	return val, nil
}

// Returns an interface{} value which may contain any valid gotables data type or NaN.
func (table *Table) GetValByColIndex(colIndex int, rowIndex int) (interface{}, error) {
	if table == nil {
		return nil, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	if printcallers {
		util.PrintCaller()
	}

	// Sadly, slice doesn't return a boolean to test whether a retrieval is in range.
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return nil, err
	}

	hasCol, err := table.HasColByColIndex(colIndex)
	if !hasCol {
		return nil, err
	}

	var val interface{}

	row := table.rows[rowIndex]
	val = row[colIndex]

	return val, nil
}

// Returns true if this table has colName and has rowIndex.
func (table *Table) HasCell(colName string, rowIndex int) (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	if printcallers {
		util.PrintCaller()
	}

	hasCol, err := table.HasCol(colName)
	if !hasCol {
		return false, err
	}

	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return false, err
	}

	return true, nil
}

// Returns true if this table has colIndex and has rowIndex.
func (table *Table) HasCellByColIndex(colIndex int, rowIndex int) (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var err error

	if printcallers {
		util.PrintCaller()
	}

	// The col exists (based on header info: name and type).
	hasColIndex, err := table.HasColByColIndex(colIndex)
	if !hasColIndex {
		return false, err
	}

	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return false, err
	}

	hasRow = len(table.rows) >= rowIndex+1
	if !hasRow {
		err = fmt.Errorf("%s: in table [%s] row %d does not exist",
			util.FuncName(),
			table.tableName,
			rowIndex)
		return false, err
	}

	// Does the cell in the row actually exist? Is the row long enough to contain cell colIndex?
	rowElementCount := len(table.rows[rowIndex])

	if rowElementCount != table.ColCount() {
		err = fmt.Errorf("%s ERROR %s table [%s] with %d cols expecting %d values (cells) per row but in row %d found: %d",
			util.FuncSource(), util.FuncName(), table.Name(), table.ColCount(), table.ColCount(), rowIndex, len(table.rows[rowIndex]))
		return false, err
	}

	hasCol := rowElementCount >= colIndex+1
	if !hasCol {
		err := fmt.Errorf("%s: in table [%s] in row %d with %d col element%s, col %d does not exist",
			util.FuncName(),
			table.tableName,
			rowIndex,
			rowElementCount,
			plural(rowElementCount),
			colIndex)
		return false, err
	}

	return true, nil
}

func (table *Table) HasRow(rowIndex int) (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	if printcallers {
		util.PrintCaller()
	}

	rowCount := len(table.rows)
	/*
		if rowCount == 0 {
			return false, fmt.Errorf("#2a table [%s] has %d row%s. Row index is out of range: %d",
				table.Name(), rowCount, plural(rowCount), rowIndex)
		} else if rowIndex < 0 || rowIndex > rowCount-1 {
			return false, fmt.Errorf("#2a table [%s] has %d row%s. Row index is out of range (0..%d): %d",
				table.Name(), rowCount, plural(rowCount), rowCount-1, rowIndex)
		}
	*/
	if rowIndex < 0 || rowIndex > rowCount-1 {
		return false, fmt.Errorf("table [%s] has %d row%s. Row index is out of range (0..%d): %d",
			table.Name(), rowCount, plural(rowCount), rowCount-1, rowIndex)
	}

	return true, nil
}

func (table *Table) IsColType(colName string, typeNameQuestioning string) (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	colType, err := table.ColType(colName)
	if err != nil {
		return false, err
	}

	if colType != typeNameQuestioning {
		return false, nil
	}

	return true, nil
}

func (table *Table) IsColTypeByColIndex(colIndex int, typeNameQuestioning string) (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	hasColIndex, err := table.HasColByColIndex(colIndex)
	if !hasColIndex {
		return false, err
	}

	colName := table.colNames[colIndex]

	isColType, err := table.IsColType(colName, typeNameQuestioning)
	if err != nil {
		return false, err
	}

	return isColType, nil
}

// ###
// Row
// ###

// See: http://blog.golang.org/json-and-go
// The empty interface serves as a general container type.
// A type assertion accesses the underlying concrete type.
// Or, if the underlying type is unknown, a type switch determines the type.
type rowAsInterface []interface{}

// But for now we will use a map to store a Row for simplicity, even though it will take up more space.

/*
	alias of RenameTable()

	Note: If this table is inside a TableSet, be sure to not set the table name the same as another table in the TableSet.
	To avoid this, use the TableSet.SetName() method.
*/
func (table *Table) SetName(tableName string) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	_, err := IsValidTableName(tableName)
	if err != nil {
		return err
	}

	table.tableName = tableName

	return nil
}

func (table *TableExported) setTableNameExported(tableName string) error {
	if len(tableName) < 1 {
		return errors.New("invalid table name has zero length")
	}

	_, err := IsValidTableName(tableName)
	if err != nil {
		return err
	}

	table.TableName = tableName

	return nil
}

/*
	Alias of SetName()

	Note: If this table is inside a TableSet, be sure to not set the table name the same as another table in the TableSet.
	To avoid this, use the TableSet.RenameTable() method.
*/
func (table *Table) RenameTable(tableName string) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}
	return table.SetName(tableName)
}

func (tableSet *TableSet) RenameTable(renameFrom string, renameTo string) error {
	if tableSet == nil {
		return fmt.Errorf("%s ERROR: tableSet.%s tableSet is <nil>", util.FuncSource(), util.FuncName())
	}

	if exists, err := tableSet.HasTable(renameFrom); !exists {
		return err
	}

	if exists, _ := tableSet.HasTable(renameTo); exists {
		return fmt.Errorf("table [%s] already exists", renameTo)
	}

	table, err := tableSet.Table(renameFrom)
	if err != nil {
		return err
	}

	err = table.RenameTable(renameTo)
	if err != nil {
		return err
	}

	return nil
}

// Note: This may leave the table in an invalid or unstable state if an error is returned.
func (table *Table) RenameCol(oldName string, newName string) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	// Make a copy of the table to reinstate in case there is an error that invalidates the table?
	// Need method: CopyTable()

	// Requires oldCol to be already in the table for renaming from.
	if hasCol, err := table.HasCol(oldName); !hasCol {
		return err
	}

	// Requires newCol to NOT be already in the table for renaming to.
	if hasCol, _ := table.HasCol(newName); hasCol {
		err := fmt.Errorf("table [%s] col already exists: %s", table.Name(), newName)
		return err
	}

	if isValid, err := IsValidColName(newName); !isValid {
		return err
	}

	// Rename col in array of col names.
	colIndex, err := table.ColIndex(oldName)
	if err != nil {
		return nil
	}
	table.colNames[colIndex] = newName

	// Rename col in map of col names to col indexes.
	delete(table.colNamesMap, oldName)    // Delete the old one.
	table.colNamesMap[newName] = colIndex // Add the new one.

	return nil
}

/*
	Return the name of the column at this column index.
	Same as table.ColNameByColIndex(colIndex)
*/
func (table *Table) ColName(colIndex int) (string, error) {
	if table == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}
	if colIndex < 0 || colIndex > table.ColCount()-1 {
		return "", fmt.Errorf("table [%s] has %d col%s. Col index out of range: %d",
			table.Name(), table.ColCount(), plural(table.ColCount()), colIndex)
	}
	colName := table.colNames[colIndex]
	return colName, nil
}

/*
	Return the name of the column at this column index.
	Same as table.ColName(colIndex)
*/
func (table *Table) ColNameByColIndex(colIndex int) (string, error) {
	return table.ColName(colIndex)
}

/*
	Check for missing values in this row.
	That means completely missing values, not just empty strings or NaN floats.
*/
func (table *Table) IsValidRow(rowIndex int) (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var err error

	if rowIndex < 0 || rowIndex > table.RowCount()-1 {
		return false, fmt.Errorf("%s: in table [%s] with %d rows, row index %d does not exist",
			util.FuncName(), table.tableName, table.RowCount(), rowIndex)
	}

	if len(table.rows[rowIndex]) != table.ColCount() {
		err = fmt.Errorf("%s ERROR %s table [%s] with %d cols expecting %d cell values per row but in row %d found: %d",
			util.FuncSource(), util.FuncName(), table.Name(), table.ColCount(), table.ColCount(), rowIndex, len(table.rows[rowIndex]))
		return false, err
	}

	return true, nil
}

/*
	Test internal consistency of this table:
		Valid table name?
		Valid col names?
		Valid col types?
		Valid (equal) lengths of internal slices of col names, col types?
		Valid data in each cell of each row?
		Valid sort keys (if any are set)?
*/
func (table *Table) IsValidTable() (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var err error
	var isValid bool

	if printcallers {
		util.PrintCaller()
	}

	// These are serious errors. Hence calls to debug.PrintStack()
	if table.tableName == "" {
		err = fmt.Errorf("%s ERROR %s: table has no name", util.FuncSource(), util.FuncName())
		return false, err
	}
	if table.colNames == nil {
		err = fmt.Errorf("%s ERROR %s: table [%s] colNames == nil", util.FuncSource(), util.FuncName(), table.tableName)
		return false, err
	}
	if table.colTypes == nil {
		err = fmt.Errorf("%s ERROR %s: table [%s] colTypes == nil", util.FuncSource(), util.FuncName(), table.tableName)
		return false, err
	}
	if table.colNamesMap == nil {
		err = fmt.Errorf("%s ERROR %s: table [%s] colNamesMap == nil", util.FuncSource(), util.FuncName(), table.tableName)
		return false, err
	}
	if table.rows == nil {
		err = fmt.Errorf("%s ERROR %s: table [%s] rows == nil", util.FuncSource(), util.FuncName(), table.tableName)
		return false, err
	}

	var tableName string = table.Name()
	if isValid, err = IsValidTableName(tableName); !isValid {
		return false, err
	}

	var colNames []string = table.getColNames()
	for colIndex := 0; colIndex < len(colNames); colIndex++ {
		var colName string = colNames[colIndex]
		if isValid, err = IsValidColName(colName); !isValid {
			return false, err
		}
	}

	var colTypes []string = table.getColTypes()
	for colIndex := 0; colIndex < len(colTypes); colIndex++ {
		var colType string = colTypes[colIndex]
		if isValid, err = IsValidColType(colType); !isValid {
			return false, err
		}
	}

	var colNamesCount int = len(colNames)
	if len(colTypes) != colNamesCount {
		err = fmt.Errorf("table [%s] with %d cols names expecting %d col types but found: %d",
			tableName, colNamesCount, colNamesCount, len(colTypes))
		return false, err
	}

	if len(table.colNamesMap) != colNamesCount {
		err = fmt.Errorf("table [%s] with %d cols names expecting %d col names lookup entries but found: %d",
			tableName, colNamesCount, colNamesCount, len(table.colNamesMap))
		return false, err
	}

	_, err = table.isValidColNamesMap()
	if err != nil {
		return false, nil
	}

	for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
		var isValidRow bool
		if isValidRow, err = table.IsValidRow(rowIndex); !isValidRow {
			return false, err
		}
	}

	// should omit value from range; this loop is equivalent to `for keyIndex := range ...` (S1005)
	for keyIndex, _ := range table.sortKeys {
		if isValid, err = IsValidColName(table.sortKeys[keyIndex].colName); !isValid {
			err = fmt.Errorf("table [%s].sortKeys[%d].colName: %v", tableName, keyIndex, err)
			return false, err
		}
		if isValid, err = IsValidColType(table.sortKeys[keyIndex].colType); !isValid {
			err = fmt.Errorf("table [%s].sortKeys[%d].colType: %v", tableName, keyIndex, err)
			return false, err
		}
		// Note: Not fully sure that a nil sortFunc is an invalid state.
		if table.sortKeys[keyIndex].sortFunc == nil {
			err = fmt.Errorf("table [%s].sortKeys[%d].sortFunc == nil", tableName, keyIndex)
			return false, err
		}
	}

	return true, nil
}

func (table *Table) isValidColNamesMap() (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	for colIndex := 0; colIndex < len(table.colNamesMap); colIndex++ {
		if table.colNamesMap[table.colNames[colIndex]] != colIndex {
			return false, fmt.Errorf("table [%s] inconsistent table.colNamesMap: colNamesMap[%q] != colIndex %d",
				table.Name(), table.colNames[colIndex], colIndex)
		}
	}

	return true, nil
}

func (tableSet *TableSet) IsValidTableSet() (bool, error) {
	if tableSet == nil {
		return false, fmt.Errorf(fmt.Sprintf("%s ERROR: tableSet.%s tableSet is <nil>", util.FuncSource(), util.FuncName()))
	}

	for i := 0; i < len(tableSet.tables); i++ {
		table := tableSet.tables[i]
		valid, err := table.IsValidTable()
		if !valid {
			return false, err
		}
	}

	return true, nil
}

//
///*
//Round is a custom implementation for rounding values.
//
//Round up if fraction is >= 0.5 otherwise round down.
//
//From: https://medium.com/@edoardo849/mercato-how-to-build-an-effective-web-scraper-with-golang-e2918d279f49#.istjzw4nl
//*/
//func Round(val float64, places int) (rounded float64) {
//	const roundOn = 0.5 // Round up if fraction is >= 0.5 otherwise round down.
//	var round float64
//	pow := math.Pow(10, float64(places))
//	digit := pow * val
//	_, frac := math.Modf(digit) // Modf(f) returns integer and fractional floating-point numbers that sum to f
//	if frac >= roundOn {
//		round = math.Ceil(digit)
//	} else {
//		round = math.Floor(digit)
//	}
//	rounded = round / pow
//	return
//}

/* Type to encode:
type Table struct {
	tableName   string
	colNames  []string
	colTypes  []string
	colNamesMap map[string]int	// To look up a colNames index from a col name.
	rows      []tableRow
	sortKeys  []sortKey
}
*/

func (table *Table) GetValAsStringByColIndex(colIndex int, rowIndex int) (string, error) {
	if table == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	if printcallers {
		util.PrintCaller()
	}

	var sVal string
	var tVal bool
	var ui8Val uint8
	var ui8SliceVal []uint8
	var byteSliceVal []byte
	var ui16Val uint16
	var ui32Val uint32
	var ui64Val uint64
	var uiVal uint
	var iVal int
	var i8Val int8
	var i16Val int16
	var i32Val int32
	var runeVal rune
	var i64Val int64
	var f32Val float32
	var f64Val float64

	var interfaceType interface{}
	var err error
	var buf bytes.Buffer
	var s string

	interfaceType, err = table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return "", err
	}

	switch table.colTypes[colIndex] {
	case "string":
		sVal = interfaceType.(string)
		// DON'T include string delimiters in string.
		buf.WriteString(sVal)
	case "bool":
		tVal = interfaceType.(bool)
		buf.WriteString(fmt.Sprintf("%t", tVal))
	case "uint8", "byte":
		ui8Val = interfaceType.(uint8)
		buf.WriteString(fmt.Sprintf("%d", ui8Val))
	case "[]uint8":
		ui8SliceVal = interfaceType.([]uint8)
		buf.WriteString(fmt.Sprintf("%v", ui8SliceVal))
	case "[]byte":
		byteSliceVal = interfaceType.([]byte)
		buf.WriteString(fmt.Sprintf("%v", byteSliceVal))
	case "uint16":
		ui16Val = interfaceType.(uint16)
		buf.WriteString(fmt.Sprintf("%d", ui16Val))
	case "uint32":
		ui32Val = interfaceType.(uint32)
		buf.WriteString(fmt.Sprintf("%d", ui32Val))
	case "uint64":
		ui64Val = interfaceType.(uint64)
		buf.WriteString(fmt.Sprintf("%d", ui64Val))
	case "uint":
		uiVal = interfaceType.(uint)
		buf.WriteString(fmt.Sprintf("%d", uiVal))
	case "int":
		iVal = interfaceType.(int)
		buf.WriteString(fmt.Sprintf("%d", iVal))
	case "int8":
		i8Val = interfaceType.(int8)
		buf.WriteString(fmt.Sprintf("%d", i8Val))
	case "int16":
		i16Val = interfaceType.(int16)
		buf.WriteString(fmt.Sprintf("%d", i16Val))
	case "rune":
		// DON'T include rune delimiters in rune.
		runeVal = interfaceType.(rune)
		buf.WriteString(fmt.Sprintf("%c", runeVal))
	case "int32":
		i32Val = interfaceType.(int32)
		buf.WriteString(fmt.Sprintf("%d", i32Val))
	case "int64":
		i64Val = interfaceType.(int64)
		buf.WriteString(fmt.Sprintf("%d", i64Val))
	case "float32":
		f32Val = interfaceType.(float32)
		var f64ValForFormatFloat float64 = float64(f32Val)
		buf.WriteString(strconv.FormatFloat(f64ValForFormatFloat, 'f', -1, 32)) // -1 strips off excess decimal places.
	case "float64":
		f64Val = interfaceType.(float64)
		buf.WriteString(strconv.FormatFloat(f64Val, 'f', -1, 64)) // -1 strips off excess decimal places.
	default:
		err = fmt.Errorf("%s ERROR IN %s: unknown type: %s", util.FuncSource(), util.FuncName(), table.colTypes[colIndex])
		return "", err
	}

	s = buf.String()

	return s, nil
}

func (table *Table) GetValAsString(colName string, rowIndex int) (string, error) {
	var colIndex int
	var err error

	if table == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	if printcallers {
		util.PrintCaller()
	}

	colIndex, err = table.ColIndex(colName)
	if err != nil {
		return "", err
	}

	return table.GetValAsStringByColIndex(colIndex, rowIndex)
}

func (table *Table) IsStructShape() (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	return table.structShape, nil
}

// Will be ignored (when writing table as string) if table RowCount() is more than 1
func (table *Table) SetStructShape(isStructShape bool) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	table.structShape = isStructShape

	return nil
}

// Join together a column of values. To compose a set of commands into a single command.
func (table *Table) JoinColVals(colName string, separator string) (string, error) {
	if table == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	sVals, err := table.GetColValsAsStrings(colName)
	if err != nil {
		return "", err
	}

	joined := strings.Join(sVals, separator)

	return joined, nil
}

// Join together a column of values. To compose a set of commands into a single command.
func (table *Table) JoinColValsByColIndex(colIndex int, separator string) (string, error) {
	if table == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	colName, err := table.ColName(colIndex)
	if err != nil {
		return "", err
	}

	return table.JoinColVals(colName, separator)
}

// Get column values (of any type) as a slice of strings.
func (table *Table) GetColValsAsStrings(colName string) ([]string, error) {
	if table == nil {
		return nil, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	rowCount := table.RowCount()
	sVals := make([]string, rowCount)
	for rowIndex := 0; rowIndex < rowCount; rowIndex++ {
		s, err := table.GetValAsString(colName, rowIndex)
		if err != nil {
			return nil, err
		}
		sVals[rowIndex] = s
	}

	return sVals, nil
}

/*
	Test that this value is a valid type for this column.

	Alias of IsValidCellValue()
*/
func (table *Table) IsValidColValue(colName string, value interface{}) (bool, error) {
	return table.IsValidCellValue(colName, value)
}

/*
	Test that this value is a valid type for this column.

	Alias of IsValidColValue()
*/
func (table *Table) IsValidCellValue(colName string, value interface{}) (bool, error) {
	colType, err := table.ColType(colName)
	if err != nil {
		return false, err
	}

	if value == nil {
		return false, fmt.Errorf("table[%s] col=%s type=%s invalid value: <nil>", table.Name(), colName, colType)
	}

	valueType := reflect.TypeOf(value)
	valueTypeName := valueType.Name()

	if valueTypeName != colType {
		return false, fmt.Errorf("table[%s] col=%s type=%s invalid value: %v", table.Name(), colName, colType, value)
	}

	return true, nil
}

/*
	- Pad trailing zeros on a string which is a floating point number.

	- Pad trailing zeros on the fractional part of a floating point number (which looks like an integer).
*/
func padTrailingZeros(s string) string {
	hasPoint := strings.Contains(s, ".")
	b := []byte(s)
	for i := len(b) - 1; i > 0; i-- {
		if b[i-1] == '.' || b[i] != '0' {
			return string(b)
		}
		if hasPoint {
			// We don't want to remove zeros from floats with no decimal places (that look like an ints).
			if b[i] == '0' {
				b[i] = ' '
			}
		}
	}

	return string(b)
}

/*
	- Trim trailing zeros on a string which is a floating point number.

	- Trim trailing zeros on the fractional part of a floating point number (which looks like an integer).
*/
func trimTrailingZeros(s string) string {
	hasPoint := strings.Contains(s, ".")
	b := []byte(s)
	for i := len(b) - 1; i > 0; i-- {
		if b[i-1] == '.' || b[i] != '0' {
			return string(b)
		}
		if hasPoint {
			// We don't want to remove zeros from floats with no decimal places (that look like an ints).
			if b[i] == '0' {
				b[i] = ' '
				b = bytes.TrimSuffix(b, []byte(" "))
			}
		}
	}

	return string(b)
}

// A helper function not used.
func (table *Table) reflectTypeOfColByColIndex(colIndex int) (reflect.Type, error) {

	var colType string
	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return nil, err
	}

	var typeOfCol reflect.Type

	switch colType {
	case "string":
		typeOfCol = reflect.TypeOf(string(""))
	case "bool":
		typeOfCol = reflect.TypeOf(bool(false))
	case "uint8":
		typeOfCol = reflect.TypeOf(uint8(0))
	case "uint16":
		typeOfCol = reflect.TypeOf(uint16(0))
	case "uint32":
		typeOfCol = reflect.TypeOf(uint32(0))
	case "uint64":
		typeOfCol = reflect.TypeOf(uint64(0))
	case "uint":
		typeOfCol = reflect.TypeOf(uint(0))
	case "int":
		typeOfCol = reflect.TypeOf(int(0))
	case "int8":
		typeOfCol = reflect.TypeOf(int8(0))
	case "int16":
		typeOfCol = reflect.TypeOf(int16(0))
	case "int32":
		typeOfCol = reflect.TypeOf(int32(0))
	case "int64":
		typeOfCol = reflect.TypeOf(int64(0))
	case "float32":
		typeOfCol = reflect.TypeOf(float32(0))
	case "float64":
		typeOfCol = reflect.TypeOf(float64(0))
	default:
		err = fmt.Errorf("%s ERROR IN %s(%q): unknown type: %s", util.FuncSource(), util.FuncName(), colType, table.colTypes[colIndex])
		return nil, err
	}

	return typeOfCol, nil
}

func isExportableName(name string) bool {
	rune0 := rune(name[0])
	if unicode.IsUpper(rune0) {
		return true
	} else {
		return false
	}
}

/*
		A fairy strict table comparison:
		1 Column count must match.
		2 Row count must match.
		3 Column names must match.
		4 Column types must match.
		5 Rows order must match.
		6 Cell values must match.
		7 Table name must match.
		8 Column order need NOT match.

	Useful for testing.
*/
func (table1 *Table) Equals(table2 *Table) (bool, error) {
	if table1 == nil {
		return false, fmt.Errorf("func (table1 *Table) Equals(table2 *Table): table1 is nil")
	}
	if table2 == nil {
		return false, fmt.Errorf("func (table1 *Table) Equals(table2 *Table): table2 is nil")
	}

	// Compare table names.
	if table1.Name() != table2.Name() {
		return false, fmt.Errorf("[%s].Equals([%s]): table names: %s != %s",
			table1.Name(), table2.Name(), table1.Name(), table2.Name())
	}

	// Compare number of rows.
	if table1.RowCount() != table2.RowCount() {
		return false, fmt.Errorf("[%s].Equals([%s]): row count: %d != %d",
			table1.Name(), table2.Name(), table1.RowCount(), table2.RowCount())
	}

	// Compare number of columns.
	if table1.ColCount() != table2.ColCount() {
		return false, fmt.Errorf("[%s].Equals([%s]): col count: %d != %d",
			table1.Name(), table2.Name(), table1.ColCount(), table2.ColCount())
	}

	// Compare column types.
	// This has the side-effect of comparing all column names.
	for colIndex := 0; colIndex < table1.ColCount(); colIndex++ {
		colName, err := table1.ColName(colIndex)
		if err != nil {
			return false, err
		}
		type1, err := table1.ColTypeByColIndex(colIndex)
		if err != nil {
			return false, err
		}
		type2, err := table2.ColType(colName)
		if err != nil {
			return false, err
		}
		if type1 != type2 {
			return false, fmt.Errorf("[%s].Equals([%s]): col %q type: %s != %s",
				table1.Name(), table2.Name(), colName, type1, type2)
		}
	}

	// Compare cell values.
	for colIndex := 0; colIndex < table1.ColCount(); colIndex++ {
		colName, err := table1.ColName(colIndex)
		if err != nil {
			return false, err
		}

		for rowIndex := 0; rowIndex < table1.RowCount(); rowIndex++ {
			val1, err := table1.GetValByColIndex(colIndex, rowIndex)
			if err != nil {
				return false, err
			}
			val2, err := table2.GetVal(colName, rowIndex)
			if err != nil {
				return false, err
			}
			if val1 != val2 {
				return false, fmt.Errorf("[%s].Equals([%s]): col %q row %d: %v != %v",
					table1.Name(), table2.Name(), colName, colIndex, val1, val2)
			}
		}
	}

	return true, nil
}

/*
	Append all columns from fromTable to table.

	Column order is ignored. Identical duplicate columns are ignored.
*/
func (table *Table) AppendColsFromTable(fromTable *Table) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s(fromTable): table is <nil>", util.FuncSource(), util.FuncName())
	}
	if fromTable == nil {
		return fmt.Errorf("fromTable.table.%s: table is <nil>", util.FuncName())
	}

	var err error

	colsTable, err := fromTable.GetColInfoAsTable()
	if err != nil {
		return err
	}

	for rowIndex := 0; rowIndex < colsTable.RowCount(); rowIndex++ {
		colName, err := colsTable.GetString("colName", rowIndex)
		if err != nil {
			return err
		}
		colType, err := colsTable.GetString("colType", rowIndex)
		if err != nil {
			return err
		}

		err = table.AppendCol(colName, colType)
		if err != nil {
			if hasCol, _ := table.HasCol(colName); hasCol {
				// Skip duplicate column name, but only if it has same type.
				var colInfo colInfoStruct
				colInfo, err = table.getColInfo(colName)
				if err != nil {
					return err
				}
				if colType != colInfo.colType {
					// Not the same type!
					return fmt.Errorf("[%s].%s([%s]): skipping duplicate colName %q (is okay), but expecting type %q, not %q",
						table.Name(), util.FuncName(), fromTable.Name(), colName, colInfo.colType, colType)
				}
			} else {
				// Must be some other error.
				return err
			}
		}
	}

	return nil
}

func (toTable *Table) AppendRowsFromTable(fromTable *Table, firstRow int, lastRow int) error {
	var err error

	if toTable == nil {
		return fmt.Errorf("toTable.table.%s: table is <nil>", util.FuncName())
	}

	if fromTable == nil {
		return fmt.Errorf("fromTable.table.%s: table is <nil>", util.FuncName())
	}

	if debugging {
		_, err = toTable.IsValidTable()
		if err != nil {
			return err
		}
	}

	if debugging {
		_, err = fromTable.IsValidTable()
		if err != nil {
			return err
		}
	}

	// Note: multiple assignment syntax in for loop.
	for fromRow, toRow := firstRow, toTable.RowCount(); fromRow <= lastRow; fromRow, toRow = fromRow+1, toRow+1 {

		err = toTable.AppendRow()
		if err != nil {
			return err
		}

		for fromCol := 0; fromCol < fromTable.ColCount(); fromCol++ {
			var colName string
			colName, err = fromTable.ColName(fromCol)
			if err != nil {
				return err
			}

			var cellVal interface{}
			cellVal, err = fromTable.GetValByColIndex(fromCol, fromRow)
			if err != nil {
				return err
			}

			err = toTable.SetVal(colName, toRow, cellVal)
			if err != nil {
				return err
			}
		}
	}

	if debugging {
		_, err = toTable.IsValidTable()
		if err != nil {
			return err
		}
	}

	if debugging {
		_, err = fromTable.IsValidTable()
		if err != nil {
			return err
		}
	}

	return nil
}

/*
	Create a new table from a range of rows in this table.
*/
func (table *Table) NewTableFromRows(newTableName string, firstRow int, lastRow int) (*Table, error) {
	var newTable *Table
	var err error

	newTable, err = NewTable(newTableName)
	if err != nil {
		return nil, err
	}

	err = newTable.AppendColsFromTable(table)
	if err != nil {
		return nil, err
	}

	err = newTable.AppendRowsFromTable(table, firstRow, lastRow)
	if err != nil {
		return nil, err
	}

	return newTable, nil
}

/*
	Create a copy of table, with or without copying its rows of data.

	To copy some but not all rows, use NewTableFromRows()
*/
func (table *Table) Copy(copyRows bool) (*Table, error) {
	var err error
	var tableCopy *Table
	const firstRow = 0
	var lastRow int = 0
	if copyRows {
		lastRow = table.RowCount() - 1
	}

	tableCopy, err = NewTable(table.Name())
	if err != nil {
		return nil, err
	}

	err = tableCopy.AppendColsFromTable(table)
	if err != nil {
		return nil, err
	}

	if copyRows {
		if table.RowCount() > 0 {
			err = tableCopy.AppendRowsFromTable(table, firstRow, lastRow)
			if err != nil {
				return nil, err
			}
		}
	}

	return tableCopy, nil
}

/*
	Create a copy of TableSet, with or without copying table rows of data.
*/
func (tableSet *TableSet) Copy(copyRows bool) (*TableSet, error) {
	var err error
	var tableSetCopy *TableSet

	tableSetCopy, err = NewTableSet(tableSet.Name())
	if err != nil {
		return nil, err
	}
	tableSetCopy.SetFileName(tableSet.FileName())

	for tableIndex := 0; tableIndex < tableSet.TableCount(); tableIndex++ {
		table, err := tableSet.TableByTableIndex(tableIndex)
		if err != nil {
			return nil, err
		}

		tableCopy, err := table.Copy(copyRows)
		if err != nil {
			return nil, err
		}

		err = tableSetCopy.AppendTable(tableCopy)
		if err != nil {
			return nil, err
		}
	}

	return tableSetCopy, nil
}

/*
	Create a new table from a range of rows in this table searched by keys.
*/
func (table *Table) NewTableFromRowsBySearchRange(newTableName string, searchValues ...interface{}) (*Table, error) {
	var newTable *Table
	var err error

	firstRow, lastRow, err := table.SearchRange(searchValues...)
	if err != nil {
		return nil, err
	}

	newTable, err = table.NewTableFromRows(newTableName, firstRow, lastRow)
	if err != nil {
		return nil, err
	}

	return newTable, nil
}

func isAlias(aliasTypeName string, typeName string) bool {
	alias, exists := typeAliasMap[typeName]
	if exists && alias == aliasTypeName {
		return true
	} else {
		return false
	}
}

func (tableSet *TableSet) DeleteTableByTableIndex(tableIndex int) error {
	if tableSet == nil {
		return fmt.Errorf("%s ERROR: tableSet.%s tableSet is <nil>", util.FuncSource(), util.FuncName())
	}
	if tableIndex < 0 || tableIndex > tableSet.TableCount()-1 {
		return fmt.Errorf("in tableSet %q with %d tables, table index %d does not exist",
			tableSet.Name(), tableSet.TableCount(), tableIndex)
	}

	// From Ivo Balbaert p182 for deleting a single element from a slice.
	tableSet.tables = append(tableSet.tables[:tableIndex], tableSet.tables[tableIndex+1:]...)

	return nil
}

func (tableSet *TableSet) DeleteTable(tableName string) error {
	if tableSet == nil {
		return fmt.Errorf("%s ERROR: tableSet.%s tableSet is <nil>", util.FuncSource(), util.FuncName())
	}

	tableIndex, err := tableSet.TableIndex(tableName)
	if err != nil {
		return err
	}

	err = tableSet.DeleteTableByTableIndex(tableIndex)
	if err != nil {
		return err
	}

	return nil
}

func (tableSet *TableSet) TableIndex(tableName string) (int, error) {
	if tableSet == nil {
		return -1, fmt.Errorf("%s ERROR: tableSet.%s tableSet is <nil>", util.FuncSource(), util.FuncName())
	}

	for tableIndex := 0; tableIndex < tableSet.TableCount(); tableIndex++ {
		table, err := tableSet.TableByTableIndex(tableIndex)
		if err != nil {
			return -1, err
		}
		if table.Name() == tableName {
			return tableIndex, nil
		}
	}

	err := fmt.Errorf("table [%s] does not exist: %s", tableName, tableName)
	return -1, err
}

// Compare slice1 with slice2
func Uint8SliceEquals(slice1, slice2 []uint8) (bool, error) {
	// UNUSED BUT retain this to repurpose as a slice comparison for other types.
	// Confirm logic by looking at bytes.Equal() code.
	if slice1 == nil && slice2 != nil {
		return false, fmt.Errorf("slice1 == <nil> && slice2 != <nil>")
	}

	if slice1 != nil && slice2 == nil {
		return false, fmt.Errorf("slice1 != <nil> && slice2 == <nil>")
	}

	if len(slice1) != len(slice2) {
		return false, fmt.Errorf("len(slice1) %d != len(slice2) %d", len(slice1), len(slice2))
	}

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false, fmt.Errorf("slice1[%d] %d != slice2[%d] %d", i, slice1[i], i, slice2[i])
		}
	}

	return true, nil
}

/*
	Mute all but one val from a multi-value function return.
	Returns an array of return values.

	Assume that RowCount() is defined like this:

	func (table *Table) RowCount() (int, error)

	To ignore the error and use the int:

	mu(table.RowCount())[0].(int)

	fmt.Sprintf("RowCount() = %d", mu(table.RowCount())[0].(int))

	0 is the index of the return value we want.
	And the return value in this case must be asserted to be type int: .(int)
*/
func mu(all ...interface{}) []interface{} {
	return all
}

func ZeroValue(typeName string) (interface{}, error) {
	switch typeName {
	case "bool":
		return false, nil
	case "float32":
		return 0.0, nil
	case "float64":
		return 0.0, nil
	case "uint":
		return 0, nil
	case "[]uint8":
		return []uint8{}, nil
	case "[]byte":
		return []byte{}, nil
	case "int":
		return 0, nil
	case "int16":
		return 0, nil
	case "int32":
		return 0, nil
	case "rune":
		return 0, nil
	case "int64":
		return 0, nil
	case "int8":
		return 0, nil
	case "string":
		return "", nil
	case "uint16":
		return 0, nil
	case "uint32":
		return 0, nil
	case "uint64":
		return 0, nil
	case "uint8":
		return 0, nil
	case "byte":
		return 0, nil
	default:
		msg := fmt.Sprintf("invalid type: %s (Valid types:", typeName)
		// Note: Because maps are not ordered, this (desirably) shuffles the order of valid col types with each call.
		for typeName, _ := range globalColTypesMap {
			msg += fmt.Sprintf(" %s", typeName)
		}
		msg += ")"
		err := errors.New(msg)
		return nil, err
	}
}

/*
	Check whether tableSet1 and tableSet2 are the same.
*/
func (tableSet1 *TableSet) Equals(tableSet2 *TableSet) (bool, error) {
	if tableSet1 == nil {
		return false, fmt.Errorf("%s ERROR: tableSet.%s tableSet1 is <nil>", util.FuncSource(), util.FuncName())
	}
	if tableSet2 == nil {
		return false, fmt.Errorf("%s ERROR: tableSet.%s tableSet2 is <nil>", util.FuncSource(), util.FuncName())
	}

	if tableSet2.TableCount() != tableSet1.TableCount() {
		return false, fmt.Errorf("tableSet2.TableCount() %d != tableSet1.TableCount() %d",
			tableSet2.TableCount(), tableSet1.TableCount())
	}

	for tableIndex := 0; tableIndex < tableSet1.TableCount(); tableIndex++ {
		table1, err := tableSet1.TableByTableIndex(tableIndex)
		if err != nil {
			return false, err
		}

		table2, err := tableSet2.TableByTableIndex(tableIndex)
		if err != nil {
			return false, err
		}

		_, err = table1.Equals(table2)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

/*
	Change the order of columns to the order provided by orderNames slice or arguments of col names.

	A new table is created and returned.
*/
func (table *Table) NewTableReorderCols(orderNames ...string) (reorderedTable *Table, err error) {
	if table == nil {
		return nil, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var colCount int = table.ColCount()
	if len(orderNames) != colCount {
		return nil, fmt.Errorf("[%s].%s(orderNames %v): expecting %d orderNames for table with colCount %d, not: %d",
			table.Name(), util.FuncNameNoParens(), orderNames, colCount, colCount, len(orderNames))
	}

	/*	In two minds about making this an error.
		if stringSliceEquals(table.colNames, orderNames) {
			return nil, fmt.Errorf("[%s].%s(orderNames %v): columns already ordered by orderNames",
				table.Name(), util.FuncNameNoParens(), orderNames)
		}
	*/

	// Translate (reordered) col names into (reordered) col indices.
	var orderIndices []int = make([]int, colCount)
	for colIndex := 0; colIndex < colCount; colIndex++ {
		orderIndices[colIndex], err = table.ColIndex(orderNames[colIndex])
		if err != nil {
			return nil, err
		}
	}

	reorderedTable, err = table.NewTableReorderColsByColIndex(orderIndices...)

	if err != nil {
		// Prepend info to returned err to make more sense of the indices.
		err = fmt.Errorf("[%s].%s(orderNames %v): %v", table.Name(), util.FuncNameNoParens(), orderNames, err)
	}

	return
}

/*
	Change the order of columns to the order provided by orderIndices slice or arguments of col indices.

	A new table is created and returned.
*/
func (table *Table) NewTableReorderColsByColIndex(orderIndices ...int) (reorderedTable *Table, err error) {
	if table == nil {
		return nil, fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var colCount int = table.ColCount()
	var rowCount int = table.RowCount()

	if len(orderIndices) != colCount {
		return nil, fmt.Errorf("[%s].%s(orderIndices %v): expecting %d orderIndices for table with colCount %d, not: %d",
			table.Name(), util.FuncNameNoParens(), orderIndices, colCount, colCount, len(orderIndices))
	}

	// Check that sorted copy of orderIndices are in colIndex range and are consecutive.
	var s []int = make([]int, colCount)
	copy(s, orderIndices)
	sort.Ints(s)
	for i := 0; i < colCount; i++ {
		if s[i] != i {
			return nil, fmt.Errorf("[%s].%s(orderIndices %v): invalid new col index: %d (valid unique range: 0..%d)",
				table.Name(), util.FuncNameNoParens(), orderIndices, s[i], colCount-1)
		}
	}

	/*	In two minds about making this an error.
		// Error if already in order.
		var alreadyInOrder = true
		for i := 0; i < colCount; i++ {
			if orderIndices[i] != i {
				alreadyInOrder = false
				break
			}
		}
		if alreadyInOrder {
			return nil, fmt.Errorf("[%s].%s(orderIndices %v): columns already ordered by orderIndices",
				table.Name(), util.FuncNameNoParens(), orderIndices)
		}
	*/

	reorderedTable, err = NewTable(table.Name())
	if err != nil {
		return nil, err
	}

	err = reorderedTable.AppendRows(rowCount)
	if err != nil {
		return nil, err
	}

	for oldIndex := 0; oldIndex < table.ColCount(); oldIndex++ {
		newIndex := orderIndices[oldIndex]

		colName, err := table.ColNameByColIndex(newIndex)
		if err != nil {
			return nil, err
		}

		colType, err := table.ColTypeByColIndex(newIndex)
		if err != nil {
			return nil, err
		}

		err = reorderedTable.AppendCol(colName, colType)
		if err != nil {
			return nil, err
		}

		for rowIndex := 0; rowIndex < rowCount; rowIndex++ {
			reorderedTable.rows[rowIndex][oldIndex] = table.rows[rowIndex][newIndex]
		}
	}

	return
}

/*
	Change the order of columns to the order provided by orderIndices slice or arguments of col indices.

	The table is reordered in-place.
*/
func (table *Table) ReorderColsByColIndex(orderIndices ...int) error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var colCount int = table.ColCount()
	var rowCount int = table.RowCount()

	if len(orderIndices) != colCount {
		return fmt.Errorf("[%s].%s(orderIndices %v): expecting %d orderIndices for table with colCount %d, not: %d",
			table.Name(), util.FuncNameNoParens(), orderIndices, colCount, colCount, len(orderIndices))
	}

	// Check that orderIndices are in colIndex range and are consecutive.
	var s []int = make([]int, colCount)
	copy(s, orderIndices)
	sort.Ints(s)
	for i := 0; i < colCount; i++ {
		if s[i] != i {
			return fmt.Errorf("[%s].%s(orderIndices %v): invalid new col index: %d (valid unique range: 0..%d)",
				table.Name(), util.FuncNameNoParens(), orderIndices, s[i], colCount-1)
		}
	}

	/*	In two minds about making this an error.
		// Error if already in order.
		var alreadyInOrder = true
		for i := 0; i < colCount; i++ {
			if orderIndices[i] != i {
				alreadyInOrder = false
				break
			}
		}
		if alreadyInOrder {
			return fmt.Errorf("[%s].%s(orderIndices %v): columns already ordered by orderIndices",
				table.Name(), util.FuncNameNoParens(), orderIndices)
		}
	*/

	// Type string (not interface{}) for colNames and colTypes to avoid type coercion.
	tempStrings := make([]string, colCount)

	// Type interface{} works for rows because the underlying cell values are stored as interface{}.
	tempInterfaces := make([]interface{}, colCount)

	// Swap col names.
	// Also update table.colNamesMap. Remember: we are not creating this table anew.
	copy(tempStrings, table.colNames)
	for colIndex := 0; colIndex < colCount; colIndex++ {
		table.colNames[colIndex] = tempStrings[orderIndices[colIndex]]
		table.colNamesMap[table.colNames[colIndex]] = colIndex
	}

	// Swap col types.
	copy(tempStrings, table.colTypes)
	for colIndex := 0; colIndex < colCount; colIndex++ {
		table.colTypes[colIndex] = tempStrings[orderIndices[colIndex]]
	}

	for rowIndex := 0; rowIndex < rowCount; rowIndex++ {
		// Swap cell values.
		copy(tempInterfaces, table.rows[rowIndex])
		for colIndex := 0; colIndex < colCount; colIndex++ {
			table.rows[rowIndex][colIndex] = tempInterfaces[orderIndices[colIndex]]
		}
	}

	return nil
}

/*
	Change the order of columns to the order provided by orderNames slice or arguments of col names.

	The table is reordered in-place.
*/
func (table *Table) ReorderCols(orderNames ...string) (err error) {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	var colCount int = table.ColCount()
	if len(orderNames) != colCount {
		return fmt.Errorf("[%s].%s(orderNames %v): expecting %d orderNames for table with colCount %d, not: %d",
			table.Name(), util.FuncNameNoParens(), orderNames, colCount, colCount, len(orderNames))
	}

	/*	In two minds about making this an error.
		if stringSliceEquals(table.colNames, orderNames) {
			return fmt.Errorf("[%s].%s(orderNames %v): columns already ordered by orderNames",
				table.Name(), util.FuncNameNoParens(), orderNames)
		}
	*/

	// Translate (reordered) col names into (reordered) col indices.
	var orderIndices []int = make([]int, colCount)
	for colIndex := 0; colIndex < colCount; colIndex++ {
		orderIndices[colIndex], err = table.ColIndex(orderNames[colIndex])
		if err != nil {
			return err
		}
	}

	err = table.ReorderColsByColIndex(orderIndices...)

	if err != nil {
		// Prepend info to returned err to make more sense of the indices.
		return fmt.Errorf("[%s].%s(orderNames %v): %v", table.Name(), util.FuncNameNoParens(), orderNames, err)
	}

	return
}

/*
	Reverse the order of rows in this table.
*/
func (table *Table) Reverse() error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	// Reversing algorithm from https://github.com/golang/go/wiki/SliceTricks
	for left, right := 0, len(table.rows)-1; left < right; left, right = left+1, right-1 {
		table.rows[left], table.rows[right] = table.rows[right], table.rows[left]
	}

	return nil
}

/*
	Shuffle the rows in this table "deterministically", meaning the same size table
	will have its rows shuffled into the same shuffled order each time. That's why
	this function has an Example, because it has repeatable output.
*/
func (table *Table) ShuffleDeterministic() error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	rand.Shuffle(len(table.rows), func(i, j int) {
		table.rows[i], table.rows[j] = table.rows[j], table.rows[i]
	})

	return nil
}

/*
	Shuffle the rows in this table "randomly", meaning you will be unable
	to predict the resulting shuffled order.
*/
func (table *Table) ShuffleRandom() error {
	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", util.FuncSource(), util.FuncName())
	}

	random := rand.New(rand.NewSource(time.Now().Unix()))
	random.Shuffle(len(table.rows), func(i, j int) {
		table.rows[i], table.rows[j] = table.rows[j], table.rows[i]
	})

	return nil
}

/*	Set table cell in colName at rowIndex to newVal interface{}

	For storage of user-defined types in a table.
*/
func (table *Table) SetInterfaceVal(colName string, rowIndex int, newVal interface{}) error {

	var err error

	if table == nil {
		return fmt.Errorf("table.%s(): table is <nil>", util.FuncName())
	}

	var valType string = fmt.Sprintf("%T", newVal)

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %s expecting val of type %s, not type %s: %v",
				util.FuncName(), table.Name(), colName, colType, valType, fmt.Sprintf("%#v", newVal))
		}
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 30% speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

/*	Set table cell in colIndex at rowIndex to newVal interface{}

	For storage of user-defined types in a table.
*/
func (table *Table) SetInterfaceValByColIndex(colIndex int, rowIndex int, newVal interface{}) error {

	// See: Set<type>ByColIndex() functions

	var err error

	if table == nil {
		return fmt.Errorf("table.%s(): table is <nil>", util.FuncName())
	}

	var valType string = fmt.Sprintf("%T", newVal)

	colType := table.colTypes[colIndex]

	if valType != colType {
		if !isAlias(colType, valType) {
			return fmt.Errorf("%s: table [%s] col %d expecting val of type %s, not type %s: %v",
				util.FuncName(), table.Name(), colIndex, colType, valType, fmt.Sprintf("%#v", newVal))
		}
	}

	// Note: hasCol was checked by ColTypeByColIndex() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return err
	}

	// Set the newVal
	// Note: This essentially inlines SetValByColIndex(): an average 5 times speedup.
	table.rows[rowIndex][colIndex] = newVal

	return nil
}

/*	Get interface table cell from colIndex at rowIndex

	For storage of user-defined types in a table.
	Retrieve to an interface{} type and then assert to your type.
*/
func (table *Table) GetInterfaceValByColIndex(colIndex int, rowIndex int) (val interface{}, err error) {

	if table == nil {
		err = fmt.Errorf("table.%s(): table is <nil>", util.FuncName())
		return
	}

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 25% speedup.
	val = table.rows[rowIndex][colIndex]

	var valType string = fmt.Sprintf("%T", val)

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col index %d is not type %s",
				util.FuncName(), table.Name(), colIndex, colType)
		}
	}

	return
}

/*	Get interface table cell from colName at rowIndex

	For storage of user-defined types in a table.
	Retrieve to an interface{} type and then assert to your type.
*/
func (table *Table) GetInterfaceVal(colName string, rowIndex int) (val interface{}, err error) {

	// See: Get<type>() functions

	if table == nil {
		return val, fmt.Errorf("table.%s(): table is <nil>", util.FuncName())
	}

	colType, err := table.ColType(colName)
	if err != nil {
		return val, err
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return val, err
	}

	// Note: hasCol was checked by ColType() above. No need to call HasCell()
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return val, err
	}

	// Get the val
	// Note: This essentially inlines GetVal(): an average 15% speedup.
	val = table.rows[rowIndex][colIndex]

	var valType string = fmt.Sprintf("%T", val)

	if valType != colType {
		if !isAlias(colType, valType) {
			return val, fmt.Errorf("%s: table [%s] col %s is not type %s",
				util.FuncName(), table.Name(), colName, colType)
		}
	}

	return
}

/*	Gob-encode the value as a string and quote it for safe parsing.

	The string consists of a contiguous pair of strings in square brackets: [human-readable][machine-readable]
	The [human-readable] part is for human readability only and is not parsed (it is skipped).
	The [machine-readable] part contains the machine-decodable text from which the original
	value is reconstituted.

	Whilst it's easy to generate a string (textual table) from a gotables.Table
	(using methods gotables.Table.String() or gotables.Table.StringUnpadded())
	it's not something that can be done by hand, unlike simple Go types such
	as string, int and bool. If you wish to hand-generate a textual table
	containing user-defined values, use
	InterfaceValAsString() to convert your user-defined values to strings.
*/
func InterfaceValAsString(val interface{}) (string, error) {
	var err error

	if val != nil {
		// Create quoted string representation of val, for humans to read, not for parsing.
		unQuotedValString := fmt.Sprintf("%#v", val)
		quotedValString := fmt.Sprintf("%q", unQuotedValString)

		// Created quoted string representation of GOB encoding of val, for parsing.
		var buffer bytes.Buffer
		var encoder *gob.Encoder = gob.NewEncoder(&buffer)
		err = encoder.Encode(val)
		if err != nil {
			return "", err
		}
		unQuotedEncoding := buffer.String()
		quotedEncoding := fmt.Sprintf("%q", unQuotedEncoding)

		// Enclose each in contiguous square-brackets for parsing.
		s := fmt.Sprintf("[%s][%s]", quotedValString, quotedEncoding)
		return s, nil
	}

	return "<nil>", nil
}
