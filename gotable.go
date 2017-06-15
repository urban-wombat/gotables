// Copyright (c) 2017 Malcolm Gorman

// Functions and methods for processing Table tables.
package gotable

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"text/tabwriter"
	"unicode"
)

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

func init() {
	log.SetFlags(log.Lshortfile)
}

/*
#####################################################################################
TableSet
#####################################################################################
2016.12.16  Malcolm Gorman  Use bytes.Buffer to construct string() string strings.
#####################################################################################
*/

// ##########
// TableSet
// ##########

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
	newTables.tables = make([]*Table, 0) // An empty slice of tables.
	return newTables, nil
}

// Read and parse a gotable file into a TableSet.
func NewTableSetFromFile(fileName string) (*TableSet, error) {
	var p parser
	//	fmt.Printf("ReadFile(%q)\n", fileName)
	p.SetFileName(fileName) // Needed for printing file and line diagnostics.

	tables, err := p.parseFile(fileName)
	if err != nil {
		return nil, err
	}

	return tables, nil
}

// Write a TableSet to a text file.
func (tableSet *TableSet) WriteFile(fileName string, mode os.FileMode) error {
	var err error
	var tableSet_String string
	var tableSet_Bytes []byte

	tableSet_String = tableSet.String()
	tableSet_Bytes = []byte(tableSet_String)
	if mode == 0 { // No permissions set.
		mode = 0666
	}
	// where(fmt.Sprintf("mode = %v\n", mode))
	err = ioutil.WriteFile(fileName, tableSet_Bytes, mode)

	return err
}

// Write a Table to a text file.
func (table *Table) WriteFile(fileName string, mode os.FileMode) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	var err error
	var table_String string
	var table_Bytes []byte

	table_String = table.String()
	table_Bytes = []byte(table_String)
	if mode == 0 { // No permissions set.
		mode = 0666
	}
//	where(fmt.Sprintf("mode = %v\n", mode))
	err = ioutil.WriteFile(fileName, table_Bytes, mode)

	return err
}

// Read and parse a gotable string into a TableSet.
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
		return nil, fmt.Errorf("%s() expecting string to contain 1 table but found %d table%s", funcName(), tableCount, plural(tableCount))
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
		return nil, fmt.Errorf("%s() expecting file to contain 1 table but found %d table%s: %s",
			funcName(), tableCount, plural(tableCount), fileName)
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
Returns a set of parsable elastic tabbed tables as a string.
*/
//func (tableSet *TableSet) String() string {
//	var verticalSep string = ""
//	var s string
//
//	var tableSetName string = tableSet.Name()
//	if tableSetName != "" {
//		s += fmt.Sprintf("# %s\n\n", tableSet.Name())
//	}
//
//	var table *Table
//	for i := 0; i < len(tableSet.tables); i++ {
//		table = tableSet.tables[i]
//		s += verticalSep
//		s += table.String()
//		verticalSep = "\n"
//	}
//	return s
// }

/*
Returns a set of parsable tables with format right-aligned (numbers) as a string.
*/
func (tableSet *TableSet) String() string {
	if tableSet == nil {
		os.Stderr.WriteString(fmt.Sprintf("ERROR: *TableSet.%s() *TableSet is <nil>\n", funcName()))
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
		s += table.String()
		verticalSep = "\n"
	}
	return s
}

func (tableSet *TableSet) StringUnpadded() string {
	var horizontalSeparator byte = ' '
	return tableSet._String(horizontalSeparator)
}

// Return parsable set of tables as a string.
func (tableSet *TableSet) _String(horizontalSeparator byte) string {
	var s string
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
	s = buf.String()
	return s
}

/*
Returns a set of parsable right aligned tables as a string.
*/
//func (tableSet *TableSet) StringAligned() string {
//	var verticalSep string = ""
//	var s string
//
//	var tableSetName string = tableSet.Name()
//	if tableSetName != "" {
//		s += fmt.Sprintf("# %s\n\n", tableSet.Name())
//	}
//
//	var table *Table
//	for i := 0; i < len(tableSet.tables); i++ {
//		table = tableSet.tables[i]
//		s += verticalSep
//		s += table.StringAligned()
//		verticalSep = "\n"
//	}
//	return s
// }

func (tableSet *TableSet) Name() string {
	return tableSet.tableSetName
}

func (tableSet *TableSet) SetName(tableSetName string) {
	tableSet.tableSetName = tableSetName
}

// The file name if this TableSet has been created from a file. Otherwise ""
func (tableSet *TableSet) FileName() string {
	return tableSet.fileName
}

func (tableSet *TableSet) SetFileName(fileName string) {
	tableSet.fileName = fileName
}

func (tableSet *TableSet) TableCount() int {
	return len(tableSet.tables)
}

// Add a table to a table set.
func (tableSet *TableSet) AppendTable(newTable *Table) error {
	if tableSet == nil {
		return fmt.Errorf("%s(*TableSet) *Table is <nil>", funcName())
	}

	// Note: Could maintain a map in parallel for rapid lookup of table names.
	for _, existingTable := range tableSet.tables {
		//where(fmt.Sprintf("existingTable.Name() = %s\n", existingTable.Name()))
		//where(fmt.Sprintf("newTable.Name() = %s\n", newTable.Name()))
		if existingTable.Name() == newTable.Name() {
			return fmt.Errorf("table [%s] already exists: %s", newTable.tableName, newTable.tableName)
		}
	}

	tableSet.tables = append(tableSet.tables, newTable)

	return nil
}

// Checks whether table exists
func (tableSet *TableSet) HasTable(tableName string) (bool, error) {
	//where(fmt.Sprintf("HasTable(%q)\n", tableName))
	for _, table := range tableSet.tables {
		if table.Name() == tableName {
			return true, nil
		}
	}
	return false, fmt.Errorf("table [%s] does not exist", tableName)
}

func (tableSet *TableSet) Table(tableName string) (*Table, error) {
	for _, table := range tableSet.tables {
		if table.Name() == tableName {
			return table, nil
		}
	}
	return nil, fmt.Errorf("table [%s] does not exist", tableName)
}

func (tableSet *TableSet) TableByTableIndex(tableIndex int) (*Table, error) {
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
	tableName      string
	colNames       []string
	colTypes       []string
	colNamesLookup map[string]int // To look up a colNames index from a col name.
	rows           tableRows
	sortKeys       []sortKey
	structShape    bool
}
type TableExported struct {
	TableName      string
	ColNames       []string
	ColTypes       []string
	ColNamesLookup map[string]int // To look up a colNames index from a col name.
	Rows           tableRows
	SortKeys       []SortKeyExported
	StructShape    bool
}

func (table *Table) getColTypes() []string {
	if table == nil {
		os.Stderr.WriteString(fmt.Sprintf("ERROR: *Table.%s() *Table is <nil>\n", funcName()))
		return nil
	}
	return table.colTypes
}

type tableRow map[string]interface{}
type tableRows []tableRow

// Note: Reimplement this as a slice of byte for each row and a master map and/or slice to track offset.

// Factory function to generate a *Table pointer.
/*
	var myTable *gotable.Table
	myTable, err = gotable.NewTable("My_Table")
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
	newTable.colNames = make([]string, 0)
	newTable.colTypes = make([]string, 0)
	newTable.colNamesLookup = map[string]int{}
	newTable.rows = make([]tableRow, 0)
	return newTable, nil
}

func newTableExported(tableName string) (*TableExported, error) {
	var err error
	var NewTableExported *TableExported = new(TableExported)
	err = NewTableExported.setTableNameExported(tableName)
	if err != nil {
		return nil, err
	}
	NewTableExported.ColNames = make([]string, 0)
	NewTableExported.ColTypes = make([]string, 0)
	NewTableExported.ColNamesLookup = map[string]int{}
	NewTableExported.Rows = make([]tableRow, 0)
	return NewTableExported, nil
}

/*
	table, err := gotable.newTableFromCols("Moviegoers", []string{"Age", "Mothballs"}, []string{"int", "bool"})

	This is a "helper" function that currently doesn't provide any help. But best not public. Would rather users use other functions.
*/
func newTableFromCols(tableName string, colNames []string, colTypes []string) (*Table, error) {
	var newTable *Table
	var err error

	// Check for invalid input.
	if len(colNames) != len(colTypes) {
		return nil, fmt.Errorf("%s(colNames, colTypes) len(colNames)=%d != len(colTypes)=%d",
			funcName(), len(colNames), len(colTypes))
	}

	newTable, err = NewTable(tableName)
	if err != nil {
		return nil, err
	}

	err = newTable.appendColNames(colNames)
	if err != nil {
		return nil, err
	}

	err = newTable.appendColTypes(colTypes)
	if err != nil {
		return nil, err
	}

	_, err = newTable.IsValidTable()
	if err != nil {
		return nil, err
	}

	return newTable, nil
}

/*
Add (append) a new blank row to this table. This does NOT initialise the cell values. They will be set to nil.

Note: This is used by the parser. Not for use by end-users.
*/
func (table *Table) appendRowOfNil() error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	newRow := make(tableRow)
	table.rows = append(table.rows, newRow)
	return nil
}

// Note: Can append rows to an empty (no columns) table, and later append columns.
func (table *Table) AppendRows(howMany int) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	if howMany <= 0 {
		return fmt.Errorf("table [%s] AppendRows(%d) cannot append %d rows (must be 1 or more)", table.Name(), howMany, howMany)
	}
	for i := 0; i < howMany; i++ {
		err := table.AppendRow()
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
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	err := table.appendRowOfNil()
	if err != nil {
		return err
	}

	var rowIndex int
	rowIndex, _ = table.lastRowIndex()
	err = table.SetRowCellsToZeroValue(rowIndex)
	if err != nil {
		return err
	}

	return nil
}

// Set all float cells in this row to NaN. This is a convenience function to use NaN as a proxy for a missing value.
func (table *Table) SetRowFloatCellsToNaN(rowIndex int) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
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

// Set all cells in this row to their zero value, such as 0, "", or false.
func (table *Table) SetRowCellsToZeroValue(rowIndex int) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	var err error

	for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
		err = table.SetCellToZeroValueByColIndex(colIndex, rowIndex)
		if err != nil {
			return err
		}
	}

	return nil
}

// Set all cells in this col to their zero value, such as 0, "", or false.
func (table *Table) SetColCellsToZeroValue(colName string) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
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
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
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
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
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

func (table *Table) SetCellToZeroValueByColIndex(colIndex int, rowIndex int) error {
	// TODO: Test for colIndex or rowIndex out of range? Or is this done by underlying functions?
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	var err error
	var colType string

	colType, err = table.ColTypeByColIndex(colIndex)
	if err != nil {
		return err
	}

	switch colType {
	case "bool":
		err = table.SetBoolByColIndex(colIndex, rowIndex, false)
	case "float32":
		err = table.SetFloat32ByColIndex(colIndex, rowIndex, 0.0)
	case "float64":
		err = table.SetFloat64ByColIndex(colIndex, rowIndex, 0.0)
	case "uint":
		err = table.SetUintByColIndex(colIndex, rowIndex, 0)
	case "int":
		err = table.SetIntByColIndex(colIndex, rowIndex, 0)
	case "int16":
		err = table.SetInt16ByColIndex(colIndex, rowIndex, 0)
	case "int32":
		err = table.SetInt32ByColIndex(colIndex, rowIndex, 0)
	case "int64":
		err = table.SetInt64ByColIndex(colIndex, rowIndex, 0)
	case "int8":
		err = table.SetInt8ByColIndex(colIndex, rowIndex, 0)
	case "string":
		err = table.SetStringByColIndex(colIndex, rowIndex, "")
	case "uint16":
		err = table.SetUint16ByColIndex(colIndex, rowIndex, 0)
	case "uint32":
		err = table.SetUint32ByColIndex(colIndex, rowIndex, 0)
	case "uint64":
		err = table.SetUint64ByColIndex(colIndex, rowIndex, 0)
	case "uint8":
		err = table.SetUint8ByColIndex(colIndex, rowIndex, 0)
	}
	if err != nil {
		return err
	}

	return nil
}

/*
This is for adding an entire new row of data to a table in bulk, so to speak.

	var row gotable.tableRow = make(gotable.tableRow)
	row["Manager"] = "JC"
	row["Apostles"] = 12
	err = table.appendRowMap(row)
	if err != nil {
	    panic(err)
	}
*/
func (table *Table) appendRowMap(newRow tableRow) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	// Check types match what the table is expecting.
	var err error
	var colName string
	var colType string
	var valuePossiblyUpdated interface{}
	var exists bool
	var valType string
	var missingValue interface{}

	// Loop through all the cols defined in the table.
	for _, colName = range table.colNames {
		colType, err = table.ColType(colName)
		if err != nil {
			return err
		}
		// where(fmt.Sprintf("colName[%d] = %q\n", i, colName))

		// Check that a col has been provided for each corresponding col in the table.
		// (We don't [yet] check to see if excess cols have been provided.)
		_, exists = newRow[colName]
		if !exists {
			// For some types (float32 and float64) there is a missing value: NaN
			missingValue, exists = missingValueForType(colType) // Only for float32 and float64
			if !exists {
				// Don't permit a misleading missing value to be present for ints, bools, strings.
				return fmt.Errorf("%s(): Table [%s] col %s type %s is missing. Only types float32 and float64 NaN missing are allowed.",
					funcName(), table.tableName, colName, colType)
			}
			newRow[colName] = missingValue
		}

		// Check that the new value col type is the same as the table col type.
		valuePossiblyUpdated = newRow[colName]
		valType = fmt.Sprintf("%T", valuePossiblyUpdated)
		if valType != colType {
			return fmt.Errorf("%s(): table [%s] col %s expecting type %s but found type %s",
				funcName(), table.tableName, colName, colType, valType)
		}
	}

	// Append the thoroughly checked and complete row to existing rows.
	table.rows = append(table.rows, newRow)

	return nil
}

func (table *Table) DeleteRow(rowIndex int) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	if rowIndex < 0 || rowIndex > table.RowCount()-1 {
		/*
			err := errors.New(fmt.Sprintf("in table [%s] with %d rows, row index %d does not exist",
				table.tableName, table.RowCount(), rowIndex))
			return err
		*/
		return fmt.Errorf("in table [%s] with %d rows, row index %d does not exist",
			table.tableName, table.RowCount(), rowIndex)
	}

	/*
		copy(table.rows[rowIndex:], table.rows[rowIndex+1:])
		table.rows = table.rows[:len(table.rows)-1]
	*/
	// From Ivo Balbaert p182 for deleting a single element.
	table.rows = append(table.rows[:rowIndex], table.rows[rowIndex+1:]...)

	return nil
}

// Delete rows from firstRowIndex to lastRowIndex inclusive. This means lastRowIndex will be deleted.
func (table *Table) DeleteRows(firstRowIndex int, lastRowIndex int) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	if firstRowIndex < 0 || firstRowIndex > table.RowCount()-1 {
		return fmt.Errorf("in table [%s] with %d rows, firstRowIndex %d does not exist",
			table.tableName, table.RowCount(), firstRowIndex)
	}
	if lastRowIndex < 0 || lastRowIndex > table.RowCount()-1 {
		return fmt.Errorf("in table [%s] with %d rows, lastRowIndex %d does not exist",
			table.tableName, table.RowCount(), lastRowIndex)
	}
	if firstRowIndex > lastRowIndex {
		return fmt.Errorf("invalid row index range: firstRowIndex %d > lastRowIndex %d", firstRowIndex, lastRowIndex)
	}

	// From Ivo Balbaert p182 for deleting a range of elements.
	table.rows = append(table.rows[:firstRowIndex], table.rows[lastRowIndex+1:]...)

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

/*
Returns a parsable elastic tabbed table as a string.
*/
func (table *Table) stringTabWriter() (string, error) {
	if table == nil {
		return "", fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	var buf bytes.Buffer
	bufWriter := bufio.NewWriter(&buf) // Implements Writer interface. Instead of using os.Stdout.
	const minwidth = 0                 // No effect?
	const tabwidth = 0                 // No effect?
	const padding = 1                  // Space beween cells. This works.
	const padchar = ' '
	const flags = uint(0) // ?
	//	const flags    = uint(tabwriter.AlignRight)	// Right aligns ALL columns!
	//	const flags    = uint(tabwriter.Debug)		// Prints vertical bar between columns.
	//	tabWriter := new(tabwriter.Writer).Init(bufWriter, minwidth, tabwidth, padding, padchar, flags)
	//	tabWriter := new(tabwriter.Writer)	// 18.01.2017 temporarily commented out
	tabWriter := tabwriter.NewWriter(bufWriter, minwidth, tabwidth, padding, padchar, flags) // 18.01.2017 trying this
	//	tabWriter.Init(bufWriter, minwidth, tabwidth, padding, padchar, flags)	// 18.01.2017 temporarily commented out
	//	fmt.Fprintf(tabWriter, table._String())	// Write this table to tabWriter.
	fmt.Fprintf(tabWriter, table._String('\t')) // Write this table to tabWriter.
	err := tabWriter.Flush()
	if err != nil {
		return "", err
	}
	err = bufWriter.Flush() // Necessary!
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (table *Table) StringUnpadded() string {

	if table == nil {
		os.Stderr.WriteString(fmt.Sprintf("ERROR: *Table.%s() *Table is <nil>\n", funcName()))
		return ""
	}

	return table._String(' ')
}

/*
Return a parsable table as a string. Intended for internal library use.
*/
func (table *Table) _String(horizontalSeparator byte) string {
	if table == nil {
		os.Stderr.WriteString(fmt.Sprintf("ERROR: *Table.%s() *Table is <nil>\n", funcName()))
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
		for rowIndex := 0; rowIndex < len(table.rows); rowIndex++ {
			var rowMap tableRow
			rowMap, err := table.rowMap(rowIndex)
			if err != nil {
				// Admittedly, a rowIndex error can't happen here. This is paranoid.
				os.Stderr.WriteString(err.Error())
				return ""
			}
			horizontalSep = ""
			for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
				var sVal string
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
				var exists bool
				buf.WriteString(horizontalSep)
				switch table.colTypes[colIndex] {
				case "string":
					sVal, exists = rowMap[table.colNames[colIndex]].(string)
					if !exists {
						sVal = ""
					}
					// Replicate "%" chars in strings so they won't be interpreted by Sprintf()
					var replicatedPercentChars string
					replicatedPercentChars = strings.Replace(sVal, "%", "%%", -1)
					buf.WriteString(fmt.Sprintf("%q", replicatedPercentChars))
				case "bool":
					tVal, exists = rowMap[table.colNames[colIndex]].(bool)
					if !exists {
						tVal = false
					}
					buf.WriteString(fmt.Sprintf("%t", tVal))
				case "uint8":
					ui8Val, exists = rowMap[table.colNames[colIndex]].(uint8)
					if !exists {
						ui8Val = 0
					}
					buf.WriteString(fmt.Sprintf("%d", ui8Val))
				case "uint16":
					ui16Val, exists = rowMap[table.colNames[colIndex]].(uint16)
					if !exists {
						ui16Val = 0
					}
					buf.WriteString(fmt.Sprintf("%d", ui16Val))
				case "uint32":
					ui32Val, exists = rowMap[table.colNames[colIndex]].(uint32)
					if !exists {
						ui32Val = 0
					}
					buf.WriteString(fmt.Sprintf("%d", ui32Val))
				case "uint64":
					ui64Val, exists = rowMap[table.colNames[colIndex]].(uint64)
					if !exists {
						ui64Val = 0
					}
					buf.WriteString(fmt.Sprintf("%d", ui64Val))
				case "uint":
					uiVal, exists = rowMap[table.colNames[colIndex]].(uint)
					if !exists {
						uiVal = 0
					}
					buf.WriteString(fmt.Sprintf("%d", uiVal))
				case "int":
					iVal, exists = rowMap[table.colNames[colIndex]].(int)
					if !exists {
						iVal = 0
					}
					buf.WriteString(fmt.Sprintf("%d", iVal))
				case "int8":
					i8Val, exists = rowMap[table.colNames[colIndex]].(int8)
					if !exists {
						i8Val = 0
					}
					buf.WriteString(fmt.Sprintf("%d", i8Val))
				case "int16":
					i16Val, exists = rowMap[table.colNames[colIndex]].(int16)
					if !exists {
						i16Val = 0
					}
					buf.WriteString(fmt.Sprintf("%d", i16Val))
				case "int32":
					i32Val, exists = rowMap[table.colNames[colIndex]].(int32)
					if !exists {
						i32Val = 0
					}
					buf.WriteString(fmt.Sprintf("%d", i32Val))
				case "int64":
					i64Val, exists = rowMap[table.colNames[colIndex]].(int64)
					if !exists {
						i64Val = 0
					}
					buf.WriteString(fmt.Sprintf("%d", i64Val))
				case "float32":
					f32Val, exists = rowMap[table.colNames[colIndex]].(float32)
					if !exists {
						f32Val = 0.0
					}
					var f64ValForFormatFloat float64 = float64(f32Val)
					buf.WriteString(strconv.FormatFloat(f64ValForFormatFloat, 'f', -1, 32)) // -1 strips off excess decimal places.
				case "float64":
					f64Val, exists = rowMap[table.colNames[colIndex]].(float64)
					if !exists {
						f64Val = 0.0
					}
					buf.WriteString(strconv.FormatFloat(f64Val, 'f', -1, 64)) // -1 strips off excess decimal places.
				default:
					log.Printf("ERROR IN func %s(): Unknown type: %s\n", funcName(), table.colTypes[colIndex])
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

// 18.01.2017 M Gorman
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

	//	where(fmt.Sprintf("matrix = %v", matrix))
	for row := 0; row < len(matrix[0]); row++ {
		sep = "" // No separator before first column.
		for col := 0; col < len(matrix); col++ {
			if alignRight[col] { // Right-aligned col.
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
					//					where(fmt.Sprintf("About to parse %s: %s (bits=%d)", colTypes[col], matrix[col][row], bits))
					float64Val, err := strconv.ParseFloat(matrix[col][row], bits)
					//					where(fmt.Sprintf("float64Val = %f from %q\n", float64Val, matrix[col][row]))
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
				//				where(fmt.Sprintf("width[%d] = %d\n", col, width[col]))
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
		os.Stderr.WriteString(fmt.Sprintf("ERROR: *Table.%s() *Table is <nil>\n", funcName()))
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
	//	where(fmt.Sprintf("len(table.colNames) = %d", len(table.colNames)))
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

	//	where(fmt.Sprintf("matrix before printMatrix(): %v", matrix))

	// Rows of data
	for rowIndex := 0; rowIndex < len(table.rows); rowIndex++ {
		var rowMap tableRow
		rowMap, err := table.rowMap(rowIndex)
		if err != nil {
			// Admittedly, a rowIndex error can't happen here. This is paranoid.
			os.Stderr.WriteString(err.Error())
			return ""
		}
		horizontalSep = "" // No gap on left of first column.
		for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
			var sVal string
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
			var exists bool
			buf.WriteString(horizontalSep)
			var s string
			switch table.colTypes[colIndex] {
			case "string":
				sVal, exists = rowMap[table.colNames[colIndex]].(string)
				if !exists {
					sVal = ""
				}
				// Replicate "%" chars in strings so they won't be interpreted by Sprintf()
				var replicatedPercentChars string
				replicatedPercentChars = strings.Replace(sVal, "%", "%%", -1)
				s = fmt.Sprintf("%q", replicatedPercentChars)
			case "bool":
				tVal, exists = rowMap[table.colNames[colIndex]].(bool)
				if !exists {
					tVal = false
				}
				s = fmt.Sprintf("%t", tVal)
			case "uint8":
				ui8Val, exists = rowMap[table.colNames[colIndex]].(uint8)
				if !exists {
					ui8Val = 0
				}
				s = fmt.Sprintf("%d", ui8Val)
			case "uint16":
				ui16Val, exists = rowMap[table.colNames[colIndex]].(uint16)
				if !exists {
					ui16Val = 0
				}
				s = fmt.Sprintf("%d", ui16Val)
			case "uint32":
				ui32Val, exists = rowMap[table.colNames[colIndex]].(uint32)
				if !exists {
					ui32Val = 0
				}
				s = fmt.Sprintf("%d", ui32Val)
			case "uint64":
				ui64Val, exists = rowMap[table.colNames[colIndex]].(uint64)
				if !exists {
					ui64Val = 0
				}
				s = fmt.Sprintf("%d", ui64Val)
			case "uint":
				uiVal, exists = rowMap[table.colNames[colIndex]].(uint)
				if !exists {
					uiVal = 0
				}
				s = fmt.Sprintf("%d", uiVal)
			case "int":
				iVal, exists = rowMap[table.colNames[colIndex]].(int)
				if !exists {
					iVal = 0
				}
				s = fmt.Sprintf("%d", iVal)
			case "int8":
				i8Val, exists = rowMap[table.colNames[colIndex]].(int8)
				if !exists {
					i8Val = 0
				}
				s = fmt.Sprintf("%d", i8Val)
			case "int16":
				i16Val, exists = rowMap[table.colNames[colIndex]].(int16)
				if !exists {
					i16Val = 0
				}
				s = fmt.Sprintf("%d", i16Val)
			case "int32":
				i32Val, exists = rowMap[table.colNames[colIndex]].(int32)
				if !exists {
					i32Val = 0
				}
				s = fmt.Sprintf("%d", i32Val)
			case "int64":
				i64Val, exists = rowMap[table.colNames[colIndex]].(int64)
				if !exists {
					i64Val = 0
				}
				s = fmt.Sprintf("%d", i64Val)
			case "float32":
				f32Val, exists = rowMap[table.colNames[colIndex]].(float32)
				if !exists {
					f32Val = 0.0
				}
				var f64ValForFormatFloat float64 = float64(f32Val)
				s = strconv.FormatFloat(f64ValForFormatFloat, 'f', -1, 32) // -1 strips off excess decimal places.
				//					precis[colIndex] = max(precis[colIndex], precisionOf(s))
				setWidths(s, colIndex, prenum, points, precis, width)
			case "float64":
				f64Val, exists = rowMap[table.colNames[colIndex]].(float64)
				if !exists {
					f64Val = 0.0
				}
				s = strconv.FormatFloat(f64Val, 'f', -1, 64) // -1 strips off excess decimal places.
				//					precis[colIndex] = max(precis[colIndex], precisionOf(s))
				setWidths(s, colIndex, prenum, points, precis, width)
			default:
				log.Printf("ERROR IN func %s(): Unknown type: %s\n", funcName(), table.colTypes[colIndex])
				return ""
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
			//			where(fmt.Sprintf("len(%q) = %d\n", s, len(s)))
			//			where(fmt.Sprintf("width[%d] = %d\n", colIndex, width[colIndex]))

			//			horizontalSep = string(horizontalSeparator)
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
		os.Stderr.WriteString(fmt.Sprintf("ERROR: *Table.%s() *Table is <nil>\n", funcName()))
	}

	var asString string
	var s string
	var structHasRowData bool = table.RowCount() > 0

	s = fmt.Sprintf("[%s]\n", table.tableName)
	for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
		s += table.colNames[colIndex] + " " + table.colTypes[colIndex]
		if structHasRowData {
			const RowIndexZero = 0
			asString, _ = table.GetValAsStringByColIndex(colIndex, RowIndexZero)
			if table.colTypes[colIndex] == "string" {
				// Note: GetValAsStringByColIndex() doesn't include delimiters around strings.
				s += " = " + fmt.Sprintf("%q", asString)
			} else {
				s += " = " + asString
			}
		}
		s += "\n"
	}

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
	// where(fmt.Sprintf("prenumber of %q = %d\n", s, prenumber))
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
	// where(fmt.Sprintf("points of %q = %d\n", s, points))
	return points
}

func precisionOf(s string) (precision int) {
	index := strings.Index(s, ".")
	if index >= 0 {
		precision = (len(s) - index) - 1
	} else {
		precision = 0
	}
	// where(fmt.Sprintf("precision of %q = %d\n", s, precision))
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

Note: This does not (yet) implement handling of commas and quotation marks inside strings.
See: https://en.wikipedia.org/wiki/Comma-separated_values
*/
func (table *Table) StringCSV() string {
	if table == nil {
		os.Stderr.WriteString(fmt.Sprintf("ERROR: *Table.%s() *Table is <nil>\n", funcName()))
		return ""
	}
	const comma = ","
	var s string
	var buf bytes.Buffer

	// Col names
	var colNameSep = ""
	for _, colName := range table.colNames {
		buf.WriteString(colNameSep)
		buf.WriteString(colName)
		colNameSep = comma
	}
	buf.WriteByte('\n')

	// Rows of data
	for rowIndex := 0; rowIndex < len(table.rows); rowIndex++ {
		var rowSep = "\n"
		var rowMap tableRow
		rowMap, _ = table.rowMap(rowIndex)
		var rowColSep = ""
		for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
			var sVal string
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
			var exists bool
			buf.WriteString(rowColSep)
			switch table.colTypes[colIndex] {
			case "string":
				sVal, exists = rowMap[table.colNames[colIndex]].(string)
				buf.WriteString(sVal)
			case "bool":
				tVal, exists = rowMap[table.colNames[colIndex]].(bool)
				buf.WriteString(fmt.Sprintf("%t", tVal))
			case "uint8":
				ui8Val, exists = rowMap[table.colNames[colIndex]].(uint8)
				buf.WriteString(fmt.Sprintf("%d", ui8Val))
			case "uint16":
				ui16Val, exists = rowMap[table.colNames[colIndex]].(uint16)
				buf.WriteString(fmt.Sprintf("%d", ui16Val))
			case "uint32":
				ui32Val, exists = rowMap[table.colNames[colIndex]].(uint32)
				buf.WriteString(fmt.Sprintf("%d", ui32Val))
			case "uint64":
				ui64Val, exists = rowMap[table.colNames[colIndex]].(uint64)
				buf.WriteString(fmt.Sprintf("%d", ui64Val))
			case "uint":
				uiVal, exists = rowMap[table.colNames[colIndex]].(uint)
				buf.WriteString(fmt.Sprintf("%d", uiVal))
			case "int":
				iVal, exists = rowMap[table.colNames[colIndex]].(int)
				buf.WriteString(fmt.Sprintf("%d", iVal))
			case "int8":
				i8Val, exists = rowMap[table.colNames[colIndex]].(int8)
				buf.WriteString(fmt.Sprintf("%d", i8Val))
			case "int16":
				i16Val, exists = rowMap[table.colNames[colIndex]].(int16)
				buf.WriteString(fmt.Sprintf("%d", i16Val))
			case "int32":
				i32Val, exists = rowMap[table.colNames[colIndex]].(int32)
				buf.WriteString(fmt.Sprintf("%d", i32Val))
			case "int64":
				i64Val, exists = rowMap[table.colNames[colIndex]].(int64)
				buf.WriteString(fmt.Sprintf("%d", i64Val))
			case "float32":
				f32Val, exists = rowMap[table.colNames[colIndex]].(float32)
				if exists {
					var f64ValForFormatFloat float64 = float64(f32Val)
					if !math.IsNaN(f64ValForFormatFloat) {
						buf.WriteString(strconv.FormatFloat(f64ValForFormatFloat, 'f', -1, 32))
					} else {
						buf.WriteString(fmt.Sprintf("")) // Missing value.
					}
				} else {
					buf.WriteString(fmt.Sprintf("")) // Missing value.
					exists = true                    // To not trigger following test of missing value for other types.
				}
			case "float64":
				f64Val, exists = rowMap[table.colNames[colIndex]].(float64)
				if exists {
					if !math.IsNaN(f64Val) {
						buf.WriteString(strconv.FormatFloat(f64Val, 'f', -1, 64))
					} else {
						buf.WriteString(fmt.Sprintf("")) // Missing value.
					}
				} else {
					buf.WriteString(fmt.Sprintf("")) // Missing value.
					exists = true                    // To not trigger following test of missing value for other types.
				}
			default:
				log.Printf("ERROR IN func %s(): unknown type: %s\n", funcName(), table.colTypes[colIndex])
				return ""
			}

			if !exists {
				log.Printf("ERROR IN func (table *Table) StringCSV() Missing a value: table [%s] col %q row %d type %q value: %v\n",
					table.Name(), table.colNames[colIndex], rowIndex, table.colTypes[colIndex],
					rowMap[table.colNames[colIndex]])
				return ""
			}

			rowColSep = comma
		}
		buf.WriteString(rowSep)
	}

	s = buf.String()

	return s
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
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	if isValid, err := IsValidColName(colName); !isValid {
		return err
	}
	if isValid, err := IsValidColType(colType); !isValid {
		return err
	}

	// Make sure this col name doesn't already exist.
	_, exists := table.colNamesLookup[colName]
	if exists {
		err := fmt.Errorf("table [%s] col already exists: %s", table.Name(), colName)
		return err
	}

	table.colNames = append(table.colNames, colName)
	table.colTypes = append(table.colTypes, colType)
	colIndex := len(table.colNames) - 1

	table.colNamesLookup[colName] = colIndex

	err := table.SetColCellsToZeroValue(colName)
	if err != nil {
		return err
	}

	return nil
}

func (table *Table) DeleteColByColIndex(colIndex int) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	if colIndex < 0 || colIndex > table.ColCount()-1 {
		err := fmt.Errorf("in table [%s] with %d cols, col index %d does not exist",
			table.tableName, table.ColCount(), colIndex)
		return err
	}

	colName, err := table.ColName(colIndex)
	if err != nil {
		return err
	}
	delete(table.colNamesLookup, colName)

	copy(table.colNames[colIndex:], table.colNames[colIndex+1:])
	table.colNames = table.colNames[:len(table.colNames)-1]

	copy(table.colTypes[colIndex:], table.colTypes[colIndex+1:])
	table.colTypes = table.colTypes[:len(table.colTypes)-1]

	return nil
}

func (table *Table) DeleteCol(colName string) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	colIndex, err := table.ColIndex(colName)
	if err != nil {
		return err
	}
	return table.DeleteColByColIndex(colIndex)
}

// This is a fundamental method called by all type-specific methods.
// Requires a val of valid type for the col in the table.
func (table *Table) SetVal(colName string, rowIndex int, val interface{}) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	hasCell, err := table.HasCell(colName, rowIndex)
	if !hasCell {
		return err
	}

	colType, err := table.ColType(colName)
	if err != nil {
		return err
	}
	valType := fmt.Sprintf("%T", val)
	if valType != colType {
		return fmt.Errorf("table [%s] col %q expecting type %q not type %q", table.Name(), colName, colType, valType)
	}

	// Set the val
	rowMap := table.rows[rowIndex]
	rowMap[colName] = val

	return nil
}

// This is a fundamental method called by all type-specific methods.
// Requires a val of valid type for the col in the table.
func (table *Table) SetValByColIndex(colIndex int, rowIndex int, val interface{}) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	hasCell, err := table.HasCellByColIndex(colIndex, rowIndex)
	if !hasCell {
		return err
	}

	colName := table.colNames[colIndex]

	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return err
	}
	valType := fmt.Sprintf("%T", val)
	if valType != colType {
		return fmt.Errorf("table [%s] col index %d col name %q expecting type %q not type %q",
			table.Name(), colIndex, colName, colType, valType)
	}

	// Set the val
	rowMap := table.rows[rowIndex]
	rowMap[colName] = val

	return nil
}

/*
Initialise a freshly created *Table (see NewTable()) with a list of column names.
The column sequence is maintained.

The list of colNames and colTypes are parallel and the lists must be of equal length to each other.
*/
func (table *Table) appendColNames(colNames []string) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	var lenNames int = len(colNames)
	var lenTypes int = len(table.colTypes)
	if lenTypes != 0 && lenNames != lenTypes {
		return fmt.Errorf("table [%s] col names %d != col types %d", table.tableName, lenNames, lenTypes)
	}

	for _, colName := range colNames {
		if isValid, err := IsValidColName(colName); !isValid {
			return err
		}
	}

	/*	MOVED DOWN
		table.colNames = append(table.colNames, colNames...)	// Explode slice with ... notation.
	*/

	for index, colName := range colNames {
		// Make sure this col name doesn't already exist.
		_, exists := table.colNamesLookup[colName]
		if exists {
			err := fmt.Errorf("table [%s] col already exists: %s", table.Name(), colName)
			return err
		}

		table.colNamesLookup[colName] = index
	}

	table.colNames = append(table.colNames, colNames...) // Explode slice with ... notation.

	return nil
}

/*
Initialise a freshly created *Table (see NewTable()) with a list of column types.
The column sequence is maintained.

The list of colNames and colTypes are parallel and the lists must be of equal length to each other.
*/
func (table *Table) appendColTypes(colTypes []string) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	var lenNames int = len(table.colNames)
	var lenTypes int = len(colTypes)
	if lenNames != 0 && lenTypes != lenNames {
		return fmt.Errorf("table [%s] col types %d != col names %d", table.tableName, lenTypes, lenNames)
	}

	for _, colType := range colTypes {
		if isValid, err := IsValidColType(colType); !isValid {
			return err
		}
	}

	table.colTypes = append(table.colTypes, colTypes...) // Explode slice with ... notation.

	return nil
}

type colInfo struct {
	colName string
	colType string
}

// Checks whether col exists
func (table *Table) HasCol(colName string) (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	_, err := table.colInfo(colName)
	var exists bool = err == nil
	return exists, err
}

// Checks whether col exists
func (table *Table) HasColByColIndex(colIndex int) (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	if colIndex < 0 || colIndex > table.ColCount()-1 {
		err := fmt.Errorf("in table [%s] with %d col%s, col index %d does not exist",
			table.tableName, table.ColCount(), plural(table.ColCount()), colIndex)
		return false, err
	}

	return true, nil
}

func (table *Table) colInfo(colName string) (colInfo, error) {
	var cInfo colInfo
	if table == nil {
		return cInfo, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	var index int
	var exists bool
	index, exists = table.colNamesLookup[colName]
	if !exists {
		err := fmt.Errorf("table [%s] col does not exist: %s", table.tableName, colName)
		return cInfo, err
	}
	cInfo.colName = colName
	cInfo.colType = table.colTypes[index]
	return cInfo, nil
}

func (table *Table) ColType(colName string) (string, error) {
	if table == nil {
		return "", fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	index, exists := table.colNamesLookup[colName]
	if !exists {
		err := fmt.Errorf("table [%s] col does not exist: %s", table.tableName, colName)
		return "", err
	}
	var colType string = table.colTypes[index]
	return colType, nil
}

func (table *Table) ColTypeByColIndex(colIndex int) (string, error) {
	if table == nil {
		return "", fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
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
		return -1, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	index, exists := table.colNamesLookup[colName]
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
		return -1, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
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
		os.Stderr.WriteString(fmt.Sprintf("ERROR: *Table.%s() *Table is <nil>\n", funcName()))
		return ""
	}
	return table.tableName
}

func (table *Table) ColCount() int {
	if table == nil {
		os.Stderr.WriteString(fmt.Sprintf("ERROR: *Table.%s() *Table is <nil>\n", funcName()))
		return -1
	}
	return len(table.colTypes)
}

/*
	Return the number of rows in this table.
	Returns -1 if there is an error (namely: the table variable is nil).
*/
func (table *Table) RowCount() int {
	if table == nil {
		// os.Stderr.WriteString(fmt.Sprintf("ERROR: *Table.%s() *Table is <nil>\n", funcName()))
		return -1
	}
	return len(table.rows)
}

// This bulk data method that returns a RowMap which is the data for a given table row.
func (table *Table) rowMap(rowIndex int) (tableRow, error) {
	if table == nil {
		return nil, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	if rowIndex < 0 || rowIndex > table.RowCount()-1 {
		return nil, fmt.Errorf("table [%s] has %d row%s. Row index out of range: %d",
			table.Name(), table.RowCount(), plural(table.RowCount()), rowIndex)
	}
	return table.rows[rowIndex], nil
}

func (table *Table) SetString(colName string, rowIndex int, newValue string) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetVal(colName, rowIndex, newValue)
}

func (table *Table) SetStringByColIndex(colIndex int, rowIndex int, newValue string) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetValByColIndex(colIndex, rowIndex, newValue)
}

func (table *Table) SetBool(colName string, rowIndex int, newValue bool) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetVal(colName, rowIndex, newValue)
}

func (table *Table) SetBoolByColIndex(colIndex int, rowIndex int, newValue bool) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetValByColIndex(colIndex, rowIndex, newValue)
}

func (table *Table) SetUint(colName string, rowIndex int, newValue uint) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetVal(colName, rowIndex, newValue)
}

func (table *Table) SetInt(colName string, rowIndex int, newValue int) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetVal(colName, rowIndex, newValue)
}

func (table *Table) SetUint8(colName string, rowIndex int, newValue uint8) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetVal(colName, rowIndex, newValue)
}

func (table *Table) SetUint16(colName string, rowIndex int, newValue uint16) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetVal(colName, rowIndex, newValue)
}

func (table *Table) SetUint32(colName string, rowIndex int, newValue uint32) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetVal(colName, rowIndex, newValue)
}

func (table *Table) SetUint64(colName string, rowIndex int, newValue uint64) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetVal(colName, rowIndex, newValue)
}

func (table *Table) SetInt8(colName string, rowIndex int, newValue int8) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetVal(colName, rowIndex, newValue)
}

func (table *Table) SetInt16(colName string, rowIndex int, newValue int16) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetVal(colName, rowIndex, newValue)
}

func (table *Table) SetInt32(colName string, rowIndex int, newValue int32) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetVal(colName, rowIndex, newValue)
}

func (table *Table) SetInt64(colName string, rowIndex int, newValue int64) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetVal(colName, rowIndex, newValue)
}

func (table *Table) SetUintByColIndex(colIndex int, rowIndex int, newValue uint) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetValByColIndex(colIndex, rowIndex, newValue)
}

func (table *Table) SetIntByColIndex(colIndex int, rowIndex int, newValue int) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetValByColIndex(colIndex, rowIndex, newValue)
}

func (table *Table) SetUint8ByColIndex(colIndex int, rowIndex int, newValue uint8) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetValByColIndex(colIndex, rowIndex, newValue)
}

func (table *Table) SetUint16ByColIndex(colIndex int, rowIndex int, newValue uint16) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetValByColIndex(colIndex, rowIndex, newValue)
}

func (table *Table) SetUint32ByColIndex(colIndex int, rowIndex int, newValue uint32) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetValByColIndex(colIndex, rowIndex, newValue)
}

func (table *Table) SetUint64ByColIndex(colIndex int, rowIndex int, newValue uint64) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetValByColIndex(colIndex, rowIndex, newValue)
}

func (table *Table) SetInt8ByColIndex(colIndex int, rowIndex int, newValue int8) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetValByColIndex(colIndex, rowIndex, newValue)
}

func (table *Table) SetInt16ByColIndex(colIndex int, rowIndex int, newValue int16) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetValByColIndex(colIndex, rowIndex, newValue)
}

func (table *Table) SetInt32ByColIndex(colIndex int, rowIndex int, newValue int32) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetValByColIndex(colIndex, rowIndex, newValue)
}

func (table *Table) SetInt64ByColIndex(colIndex int, rowIndex int, newValue int64) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetValByColIndex(colIndex, rowIndex, newValue)
}

func (table *Table) SetFloat32(colName string, rowIndex int, newValue float32) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetVal(colName, rowIndex, newValue)
}

func (table *Table) SetFloat32ByColIndex(colIndex int, rowIndex int, newValue float32) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetValByColIndex(colIndex, rowIndex, newValue)
}

func (table *Table) SetFloat64(colName string, rowIndex int, newValue float64) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetVal(colName, rowIndex, newValue)
}

func (table *Table) SetFloat64ByColIndex(colIndex int, rowIndex int, newValue float64) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetValByColIndex(colIndex, rowIndex, newValue)
}

// This is a fundamental method called by all type-specific methods.
// Returns an interface{} value which may contain any valid gotable data type or NaN.
func (table *Table) GetVal(colName string, rowIndex int) (interface{}, error) {
	// Why don't we simply call GetValByColIndex() ???
	if table == nil {
		return nil, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	// Sadly, slice doesn't return a boolean to test whether a retrieval is in range.
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return nil, err
	}
	rowMap := table.rows[rowIndex]

	val, exists := rowMap[colName]
	if !exists {
		// This is purely to get a nicely formatted error message.
		// Only AFTER attempt to retrieve, for optimistic efficiency.
		// This call to HasCol() will always return false and return an error.
		_, err := table.HasCol(colName)
		return nil, err
	}

	return val, nil
}

// This is a fundamental method called by all type-specific methods.
// Returns an interface{} value which may contain any valid gotable data type or NaN.
func (table *Table) GetValByColIndex(colIndex int, rowIndex int) (interface{}, error) {
	if table == nil {
		return nil, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	// Sadly, slice doesn't return a boolean to test whether a retrieval is in range.
	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return nil, err
	}
	rowMap := table.rows[rowIndex]

	hasColIndex, err := table.HasColByColIndex(colIndex)
	if !hasColIndex {
		return nil, err
	}
	colName := table.colNames[colIndex]

	val, exists := rowMap[colName]
	if !exists {
		// This is purely to get a nicely formatted error message.
		// Only AFTER attempt to retrieve, for optimistic efficiency.
		// This call to HasCol() will always return false and return an error.
		_, err := table.HasCol(colName)
		return nil, err
	}

	return val, nil
}

// Returns true if this table has colName and has rowIndex.
func (table *Table) HasCell(colName string, rowIndex int) (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
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
		return false, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	hasColIndex, err := table.HasColByColIndex(colIndex)
	if !hasColIndex {
		return false, err
	}

	hasRow, err := table.HasRow(rowIndex)
	if !hasRow {
		return false, err
	}

	return true, nil
}

func (table *Table) HasRow(rowIndex int) (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	if rowIndex < 0 || rowIndex > table.RowCount()-1 {
		return false, fmt.Errorf("table [%s] has %d row%s. Row index out of range: %d",
			table.Name(), table.RowCount(), plural(table.RowCount()), rowIndex)
	}
	return true, nil
}

func (table *Table) GetString(colName string, rowIndex int) (string, error) {

	var err error
	const zeroVal = ""

	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	var interfaceType interface{}
	interfaceType, err = table.GetVal(colName, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(string)
	if !valid {
		_, err = table.IsColType(colName, "string")	// Get an error message.
		return zeroVal, err
	}

	return val, err
}

func (table *Table) GetStringByColIndex(colIndex int, rowIndex int) (string, error) {
	const zeroVal = ""
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(string)
	if !valid {
		_, err := table.IsColTypeByColIndex(colIndex, "string")
		return zeroVal, err
	}

	return val, nil
}

func (table *Table) GetFloat32(colName string, rowIndex int) (float32, error) {
	const zeroVal = 0.0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetVal(colName, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(float32)
	if !valid {
		_, err = table.IsColType(colName, "float32")
		return zeroVal, err
	}

	return val, err
}

func (table *Table) GetFloat32ByColIndex(colIndex int, rowIndex int) (float32, error) {
	const zeroVal = 0.0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(float32)
	if !valid {
		_, err := table.IsColTypeByColIndex(colIndex, "float32")
		return zeroVal, err
	}

	return val, nil
}

func (table *Table) GetFloat64(colName string, rowIndex int) (float64, error) {
	const zeroVal = 0.0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetVal(colName, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(float64)
	if !valid {
		_, err = table.IsColType(colName, "float64")
		return zeroVal, err
	}

	return val, err
}

func (table *Table) GetFloat64ByColIndex(colIndex int, rowIndex int) (float64, error) {
	const zeroVal = 0.0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(float64)
	if !valid {
		_, err := table.IsColTypeByColIndex(colIndex, "float64")
		return zeroVal, err
	}

	return val, nil
}

func (table *Table) GetUint(colName string, rowIndex int) (uint, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetVal(colName, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(uint)
	if !valid {
		_, err = table.IsColType(colName, "uint")
		return zeroVal, err
	}

	return val, err
}

func (table *Table) GetUintByColIndex(colIndex int, rowIndex int) (uint, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(uint)
	if !valid {
		_, err := table.IsColTypeByColIndex(colIndex, "uint")
		return zeroVal, err
	}

	return val, nil
}

func (table *Table) GetInt(colName string, rowIndex int) (int, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetVal(colName, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(int)
	if !valid {
		_, err = table.IsColType(colName, "int")
		return zeroVal, err
	}

	return val, err
}

func (table *Table) GetIntByColIndex(colIndex int, rowIndex int) (int, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(int)
	if !valid {
		_, err := table.IsColTypeByColIndex(colIndex, "int")
		return zeroVal, err
	}

	return val, nil
}

func (table *Table) GetUint8(colName string, rowIndex int) (uint8, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetVal(colName, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(uint8)
	if !valid {
		_, err = table.IsColType(colName, "uint8")
		return zeroVal, err
	}

	return val, err
}

func (table *Table) GetUint8ByColIndex(colIndex int, rowIndex int) (uint8, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(uint8)
	if !valid {
		_, err := table.IsColTypeByColIndex(colIndex, "uint8")
		return zeroVal, err
	}

	return val, nil
}

func (table *Table) GetUint16(colName string, rowIndex int) (uint16, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetVal(colName, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(uint16)
	if !valid {
		_, err = table.IsColType(colName, "uint16")
		return zeroVal, err
	}

	return val, err
}

func (table *Table) GetUint16ByColIndex(colIndex int, rowIndex int) (uint16, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(uint16)
	if !valid {
		_, err := table.IsColTypeByColIndex(colIndex, "uint16")
		return zeroVal, err
	}

	return val, nil
}

func (table *Table) GetUint32(colName string, rowIndex int) (uint32, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetVal(colName, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(uint32)
	if !valid {
		_, err = table.IsColType(colName, "uint32")
		return zeroVal, err
	}

	return val, err
}

func (table *Table) GetUint32ByColIndex(colIndex int, rowIndex int) (uint32, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(uint32)
	if !valid {
		_, err := table.IsColTypeByColIndex(colIndex, "uint32")
		return zeroVal, err
	}

	return val, nil
}

func (table *Table) GetUint64(colName string, rowIndex int) (uint64, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetVal(colName, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(uint64)
	if !valid {
		_, err = table.IsColType(colName, "uint64")
		return zeroVal, err
	}

	return val, err
}

func (table *Table) GetUint64ByColIndex(colIndex int, rowIndex int) (uint64, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(uint64)
	if !valid {
		_, err := table.IsColTypeByColIndex(colIndex, "uint64")
		return zeroVal, err
	}

	return val, nil
}

func (table *Table) GetInt8(colName string, rowIndex int) (int8, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetVal(colName, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(int8)
	if !valid {
		_, err = table.IsColType(colName, "int8")
		return zeroVal, err
	}

	return val, err
}

func (table *Table) GetInt8ByColIndex(colIndex int, rowIndex int) (int8, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(int8)
	if !valid {
		_, err := table.IsColTypeByColIndex(colIndex, "int8")
		return zeroVal, err
	}

	return val, nil
}

func (table *Table) GetInt16(colName string, rowIndex int) (int16, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetVal(colName, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(int16)
	if !valid {
		_, err = table.IsColType(colName, "int16")
		return zeroVal, err
	}

	return val, err
}

func (table *Table) GetInt16ByColIndex(colIndex int, rowIndex int) (int16, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(int16)
	if !valid {
		_, err := table.IsColTypeByColIndex(colIndex, "int16")
		return zeroVal, err
	}

	return val, nil
}

func (table *Table) GetInt32(colName string, rowIndex int) (int32, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetVal(colName, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(int32)
	if !valid {
		_, err = table.IsColType(colName, "int32")
		return zeroVal, err
	}

	return val, err
}

func (table *Table) GetInt32ByColIndex(colIndex int, rowIndex int) (int32, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(int32)
	if !valid {
		_, err = table.IsColTypeByColIndex(colIndex, "int32")
		return zeroVal, err
	}

	return val, nil
}

func (table *Table) GetInt64(colName string, rowIndex int) (int64, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetVal(colName, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(int64)
	if !valid {
		_, err = table.IsColType(colName, "int64")
		return zeroVal, err
	}

	return val, err
}

func (table *Table) GetInt64ByColIndex(colIndex int, rowIndex int) (int64, error) {
	const zeroVal = 0
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(int64)
	if !valid {
		_, err = table.IsColTypeByColIndex(colIndex, "int64")
		return zeroVal, err
	}

	return val, nil
}

func (table *Table) GetBool(colName string, rowIndex int) (bool, error) {
	const zeroVal = false
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetVal(colName, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(bool)
	if !valid {
		_, err = table.IsColType(colName, "bool")
		return zeroVal, err
	}

	return val, err
}

func (table *Table) GetBoolByColIndex(colIndex int, rowIndex int) (bool, error) {
	const zeroVal = false
	if table == nil {
		return zeroVal, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	interfaceType, err := table.GetValByColIndex(colIndex, rowIndex)
	if err != nil {
		return zeroVal, err
	}

	val, valid := interfaceType.(bool)
	if !valid {
		_, err = table.IsColTypeByColIndex(colIndex, "bool")
		return zeroVal, err
	}

	return val, nil
}

func (table *Table) IsColType(colName string, typeNameQuestioning string) (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	colType, _ := table.ColType(colName)
	if colType != typeNameQuestioning {
		err := fmt.Errorf("table [%s] col name %q type is %q, not %q",
			table.tableName, colName, colType, typeNameQuestioning)
		return false, err
	}
	return true, nil
}

func (table *Table) IsColTypeByColIndex(colIndex int, typeNameQuestioning string) (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	hasColIndex, err := table.HasColByColIndex(colIndex)
	if !hasColIndex {
		return false, err
	}

	colName := table.colNames[colIndex]

	isColType, err := table.IsColType(colName, typeNameQuestioning)
	if !isColType {
		colType, _ := table.ColType(colName)
		err := fmt.Errorf("table [%s] col %q col index %d type is %q, not %q",
			table.tableName, colName, colIndex, colType, typeNameQuestioning)
		return false, err
	}

	return true, nil
}

// ###
// Row
// ###

// See: http://blog.golang.org/json-and-go
// The empty interface serves as a general container type.
// A type assertion accesses the underlying concrete type.
// Or, if the underlying type is unknown, a type switch determines the type.
type _RowAsInterface []interface{}

// But for now we will use a map to store a Row for simplicity, even though it will take up more space.

// alias of RenameTable()
func (table *Table) SetName(tableName string) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	if len(tableName) < 1 {
		return errors.New("invalid table name has zero length")
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
*/
func (table *Table) RenameTable(tableName string) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.SetName(tableName)
}

func (tableSet *TableSet) RenameTable(renameFrom string, renameTo string) error {
	if exists, err := tableSet.HasTable(renameFrom); !exists {
		return err
	}

	if exists, _ := tableSet.HasTable(renameTo); exists {
		return fmt.Errorf("table [%s] already exists.", renameTo)
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
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
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
	delete(table.colNamesLookup, oldName)    // Delete the old one.
	table.colNamesLookup[newName] = colIndex // Add the new one.

	// Rename each row.
	// table.renameColCells()
	for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
		// Get the row
		var rowMap tableRow
		rowMap, err := table.rowMap(rowIndex)
		if err != nil {
			return nil
		}
		// Rename col in map of row cell values.
		// Save the cell value so it doesn't get deleted.
		var cellValue interface{}
		cellValue, ok := rowMap[oldName]
		if !ok {
			msg := fmt.Sprintf("SYSTEM ERROR: Table [%s] row %d col %q cell value does not exist!",
				table.Name(), rowIndex, oldName)
			err := errors.New(msg)
			return err
		}
		delete(rowMap, oldName)     // Delete the old name and value.
		rowMap[newName] = cellValue // Add the new name and saved cell value.
	}

	return nil
}

func (table *Table) ColName(colIndex int) (string, error) {
	if table == nil {
		return "", fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	if colIndex < 0 || colIndex > table.ColCount()-1 {
		return "", fmt.Errorf("table [%s] has %d col%s. Col index out of range: %d",
			table.Name(), table.ColCount(), plural(table.ColCount()), colIndex)
	}
	colName := table.colNames[colIndex]
	return colName, nil
}

/*
	Check for missing values in this row.
	That means completely missing values, not just empty strings or NaN floats.
*/
func (table *Table) IsValidRow(rowIndex int) (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	var err error
	var rowMap tableRow

	rowMap, err = table.rowMap(rowIndex)
	if err != nil {
		return false, err
	}

	var colNames []string = table.getColNames()
	for colIndex := 0; colIndex < len(colNames); colIndex++ {
		var ok bool
		var colName string = colNames[colIndex]
		_, ok = rowMap[colName]
		if !ok {
			msg := fmt.Sprintf("table [%s] col %q row %d cell value is missing", table.Name(), colName, rowIndex)
			err := errors.New(msg)
			return false, err
		}
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
		return false, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	var err error
	var isValid bool

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

	if len(table.colNamesLookup) != colNamesCount {
		err = fmt.Errorf("table [%s] with %d cols names expecting %d col names lookup entries but found: %d",
			tableName, colNamesCount, colNamesCount, len(table.colNamesLookup))
		return false, err
	}

	for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
		if isValid, err = table.IsValidRow(rowIndex); !isValid {
			return false, err
		}
		var rowMap tableRow
		rowMap, err = table.rowMap(rowIndex)
		if err != nil {
			return false, err
		}
		if len(rowMap) != colNamesCount {
			err = fmt.Errorf("table [%s] with %d cols expecting %d values per row but in row %d found: %d",
				tableName, colNamesCount, colNamesCount, rowIndex, len(rowMap))
			return false, err
		}
	}

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

/*
Round is a custom implementation for rounding values as Golang does not include a round function in the standard math package.

Round up if fraction is >= 0.5 otherwise round down.

From: https://medium.com/@edoardo849/mercato-how-to-build-an-effective-web-scraper-with-golang-e2918d279f49#.istjzw4nl
*/
func Round(val float64, places int) (rounded float64) {
	const roundOn = 0.5 // Round up if fraction is >= 0.5 otherwise round down.
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, frac := math.Modf(digit) // Modf(f) returns integer and fractional floating-point numbers that sum to f
	if frac >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	rounded = round / pow
	return
}

/* Type to encode:
type Table struct {
	tableName   string
	colNames  []string
	colTypes  []string
	colNamesLookup map[string]int	// To look up a colNames index from a col name.
	rows        tableRows
	sortKeys  []sortKey
}
*/

// Prepare table for GOB encoding, by copying its contents to an exportable (public) table data structure.
func (table *Table) exportTable() (*TableExported, error) {
	if table == nil {
		return nil, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	var err error
	var elementCount int
	var tableExported *TableExported

	tableExported, err = newTableExported(table.Name())
	if err != nil {
		return nil, err
	}

	var colCount int = table.ColCount()

	tableExported.ColNames = make([]string, colCount)
	if len(tableExported.ColNames) != colCount {
		err = fmt.Errorf("exportTable() [%s] Could not make col names slice of size %d",
			table.Name(), colCount)
		return nil, err
	}
	elementCount = copy(tableExported.ColNames, table.colNames)
	if elementCount != colCount {
		err = fmt.Errorf("exportTable() [%s] expecting to export %d col names but exported: %d",
			table.Name(), colCount, elementCount)
		return nil, err
	}

	tableExported.ColTypes = make([]string, colCount)
	if len(tableExported.ColTypes) != colCount {
		err = fmt.Errorf("exportTable() [%s] Could not make col types slice of size %d",
			table.Name(), colCount)
		return nil, err
	}
	elementCount = copy(tableExported.ColTypes, table.colTypes)
	if elementCount != colCount {
		err = fmt.Errorf("exportTable() [%s] expecting to export %d col types but exported: %d",
			table.Name(), colCount, elementCount)
		return nil, err
	}

	tableExported.ColNamesLookup = map[string]int{}
	for key, val := range table.colNamesLookup {
		tableExported.ColNamesLookup[key] = val
	}

	var rowCount int = table.RowCount()

	tableExported.Rows = make(tableRows, rowCount)
	if len(tableExported.Rows) != rowCount {
		err = fmt.Errorf("exportTable() [%s] Could not make rows slice of size %d",
			table.Name(), rowCount)
		return nil, err
	}
	elementCount = copy(tableExported.Rows, table.rows)
	if elementCount != rowCount {
		err = fmt.Errorf("exportTable() [%s] expecting to export %d rows but exported: %d",
			table.Name(), rowCount, elementCount)
		return nil, err
	}

	tableExported.SortKeys = []SortKeyExported{}
	for keyIndex, _ := range table.sortKeys {
		tableExported.SortKeys[keyIndex] = SortKeyExported{}
		tableExported.SortKeys[keyIndex].ColName = table.sortKeys[keyIndex].colName
		tableExported.SortKeys[keyIndex].ColType = table.sortKeys[keyIndex].colType
		tableExported.SortKeys[keyIndex].Reverse = table.sortKeys[keyIndex].reverse
		tableExported.SortKeys[keyIndex].SortFunc = table.sortKeys[keyIndex].sortFunc
	}

	tableExported.StructShape = table.structShape

	return tableExported, nil
}

// Reconstitute table from GOB decoding.
func (tableExported *TableExported) importTable() (*Table, error) {
	if tableExported == nil {
		return nil, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	var err error
	var elementCount int
	var table *Table

	var tableName string = tableExported.TableName
	table, err = NewTable(tableName)
	if err != nil {
		return nil, err
	}
	var isValid bool

	var colCount int = len(tableExported.ColNames)

	table.colNames = make([]string, colCount)
	elementCount = copy(table.colNames, tableExported.ColNames)
	if elementCount != colCount {
		err = fmt.Errorf("importTable() [%s] expecting to import %d col names but imported: %d",
			tableName, colCount, elementCount)
		return nil, err
	}

	table.colTypes = make([]string, colCount)
	elementCount = copy(table.colTypes, tableExported.ColTypes)
	if elementCount != colCount {
		err = fmt.Errorf("importTable() [%s] expecting to import %d col types but imported: %d",
			tableName, colCount, elementCount)
		return nil, err
	}

	table.colNamesLookup = map[string]int{}
	for key, val := range tableExported.ColNamesLookup {
		table.colNamesLookup[key] = val
	}

	var rowCount int = len(tableExported.Rows)

	table.rows = make(tableRows, rowCount)
	elementCount = copy(table.rows, tableExported.Rows)
	if elementCount != rowCount {
		err = fmt.Errorf("importTable() [%s] expecting to import %d rows but imported: %d",
			table.Name(), rowCount, elementCount)
		return nil, err
	}

	table.sortKeys = []sortKey{}
	for keyIndex, _ := range table.sortKeys {
		table.sortKeys[keyIndex] = sortKey{}
		table.sortKeys[keyIndex].colName = tableExported.SortKeys[keyIndex].ColName
		table.sortKeys[keyIndex].colType = tableExported.SortKeys[keyIndex].ColType
		table.sortKeys[keyIndex].reverse = tableExported.SortKeys[keyIndex].Reverse
		table.sortKeys[keyIndex].sortFunc = tableExported.SortKeys[keyIndex].SortFunc
	}

	table.structShape = tableExported.StructShape

	isValid, err = table.IsValidTable()
	if !isValid {
		return nil, err
	}

	return table, nil
}

func (table *Table) GobEncode() (bytes.Buffer, error) {
	if table == nil {
		var blankBuffer bytes.Buffer
		return blankBuffer, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	var err error
	var buffer bytes.Buffer
	var enc *gob.Encoder = gob.NewEncoder(&buffer)
	var tableExported *TableExported

	tableExported, err = table.exportTable()
	if err != nil {
		var blankBuffer bytes.Buffer
		return blankBuffer, err
	}

	err = enc.Encode(tableExported)
	if err != nil {
		var blankBuffer bytes.Buffer
		return blankBuffer, err
	}

	return buffer, nil
}

//func (table *Table) GobEncode() ([]byte, error) {
//	var err error
//	var buffer bytes.Buffer
//	var enc *gob.Encoder = gob.NewEncoder(&buffer)
//	var tableExported *TableExported
//
//	tableExported, err = table.exportTable()
//	if err != nil {
//		return buffer.Bytes(), err
//	}
//
//	err = enc.Encode(tableExported)
//	if err != nil {
//		return buffer.Bytes(), err
//	}
//
//	return buffer.Bytes(), nil
//}

func (tableSet *TableSet) GobEncode() ([]bytes.Buffer, error) {
	var emptyBuffer []bytes.Buffer
	var encodedTableSet []bytes.Buffer
	var err error

	for tableIndex := 0; tableIndex < len(tableSet.tables); tableIndex++ {
		var table *Table = tableSet.tables[tableIndex]
		var encodedTable bytes.Buffer
		encodedTable, err = table.GobEncode()
		if err != nil {
			return emptyBuffer, err
		}
		encodedTableSet = append(encodedTableSet, encodedTable)
		if len(encodedTableSet) != tableIndex+1 {
			err = fmt.Errorf("GobEncode(): table [%s] Error appending table to table set",
				table.Name())
			return emptyBuffer, err
		}
	}

	// Add header information to the tail end of the buffer array.
	var tableSetHeader TableSetExported
	tableSetHeader.TableSetName = tableSet.tableSetName
	tableSetHeader.FileName = tableSet.fileName
	var encodedHeader bytes.Buffer
	var enc *gob.Encoder = gob.NewEncoder(&encodedHeader)
	err = enc.Encode(tableSetHeader)
	if err != nil {
		return emptyBuffer, err
	}
	encodedTableSet = append(encodedTableSet, encodedHeader)
	var headerIndex int = len(tableSet.tables)
	if len(encodedTableSet) != headerIndex+1 {
		err = fmt.Errorf("GobEncode(): error appending table set header to table set")
		return emptyBuffer, err
	}

	return encodedTableSet, nil
}

//func (tableSet *TableSet) GobEncode() ([][]byte, error) {
////	var emptyBuffer []bytes.Buffer
//	var emptyBuffer [][]byte
////	var encodedTableSet []bytes.Buffer
//	var encodedTableSet [][]byte
//	var err error
//
//	for tableIndex := 0; tableIndex < len(tableSet.tables); tableIndex++ {
//		var table *Table
//		table = tableSet.tables[tableIndex]
////		var encodedTable bytes.Buffer
//		var encodedTable []byte
//		encodedTable, err = table.GobEncode()
//		if err != nil {
//			return emptyBuffer, err
//		}
//		encodedTableSet = append(encodedTableSet, encodedTable)
//		if len(encodedTableSet) != tableIndex+1 {
//			err = fmt.Errorf("GobEncode(): table [%s] Error appending table to table set",
//				table.Name())
//			return emptyBuffer, err
//		}
//	}
//
//	// Add header information to the tail end of the buffer array.
//	var tableSetHeader TableSetExported
//	tableSetHeader.TableSetName = tableSet.tableSetName
//	tableSetHeader.FileName = tableSet.fileName
//	var encodedHeader bytes.Buffer
//	var enc *gob.Encoder = gob.NewEncoder(&encodedHeader)
//	err = enc.Encode(tableSetHeader)
//	if err != nil {
//		return emptyBuffer, err
//	}
//	encodedTableSet = append(encodedTableSet, encodedHeader.Bytes())
//	var headerIndex int = len(tableSet.tables)
//	if len(encodedTableSet) != headerIndex+1 {
//		err = fmt.Errorf("GobEncode(): error appending table set header to table set")
//		return emptyBuffer, err
//	}
//
//	return encodedTableSet, nil
//}

func GobDecodeTableSet(buffer []bytes.Buffer) (*TableSet, error) {
	var tableSet *TableSet
	var err error
	tableSet, err = NewTableSet("")
	if err != nil {
		return nil, err
	}

	var table *Table
	var tableCount = len(buffer) - 1 // The tail end buffer element is the header.
	for tableIndex := 0; tableIndex < tableCount; tableIndex++ {
		table, err = GobDecodeTable(buffer[tableIndex])
		if err != nil {
			return nil, err
		}
		err = tableSet.AppendTable(table)
		if err != nil {
			return nil, err
		}
	}

	// Decode and restore the header.
	var headerIndex int = len(buffer) - 1
	var encodedHeader bytes.Buffer = buffer[headerIndex]
	var dec *gob.Decoder = gob.NewDecoder(&encodedHeader)
	var tableSetHeader TableSetExported
	err = dec.Decode(&tableSetHeader)
	if err != nil {
		return nil, err
	}
	tableSet.tableSetName = tableSetHeader.TableSetName
	tableSet.fileName = tableSetHeader.FileName

	return tableSet, nil
}

//func GobDecodeTableSet(buffer [][]byte) (*TableSet, error) {
//	var tableSet *TableSet
//	var err error
//	tableSet, err = NewTableSet("")
//	if err != nil {
//		return nil, err
//	}
//
//	var table *Table
//	var tableCount = len(buffer)-1	// The tail end buffer element is the header.
//	for tableIndex := 0; tableIndex < tableCount; tableIndex++ {
//		table, err = GobDecodeTable(buffer[tableIndex])
//		if err != nil {
//			return nil, err
//		}
//		err = tableSet.AppendTable(table)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	// Decode and restore the header.
//	var headerIndex int = len(buffer)-1
////	var encodedHeader bytes.Buffer = buffer[headerIndex]
//	var encodedHeader []byte = buffer[headerIndex]
//	var dec *gob.Decoder = gob.NewDecoder(&encodedHeader)
//	var tableSetHeader TableSetExported
//	err = dec.Decode(&tableSetHeader)
//	if err != nil {
//		return nil, err
//	}
//	tableSet.tableSetName = tableSetHeader.TableSetName
//	tableSet.fileName = tableSetHeader.FileName
//
//	return tableSet, nil
//}

func GobDecodeTable(buffer bytes.Buffer) (*Table, error) {
	var err error
	var tableDecoded *Table
	var dec *gob.Decoder = gob.NewDecoder(&buffer)
	var tableExported *TableExported

	err = dec.Decode(&tableExported)
	if err != nil {
		return nil, err
	}

	tableDecoded, err = tableExported.importTable()
	if err != nil {
		return nil, err
	}

	return tableDecoded, nil
}

//func GobDecodeTable(buffer []byte) (*Table, error) {
//	var err error
//	var tableDecoded *Table
//	var dec *gob.Decoder = gob.NewDecoder(&buffer)
//	var tableExported *TableExported
//
//	err = dec.Decode(&tableExported)
//	if err != nil {
//		return nil, err
//	}
//
//	tableDecoded, err = tableExported.importTable()
//	if err != nil {
//		return nil, err
//	}
//
//	return tableDecoded, nil
//}

func funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	nameFull := runtime.FuncForPC(pc).Name() // main.foo
	nameEnd := filepath.Ext(nameFull)        // .foo
	name := strings.TrimPrefix(nameEnd, ".") // foo
	return name
}

func funcNameFull() string {
	pc, sourceFile, lineNumber, ok := runtime.Caller(1)
	if !ok {
		return "Could not obtain func name and source file information."
	}
	nameFull := runtime.FuncForPC(pc).Name() // main.foo
	return fmt.Sprintf("%s[%d] %s", sourceFile, lineNumber, nameFull)
}

func (table *Table) GetValAsStringByColIndex(colIndex int, rowIndex int) (string, error) {
	if table == nil {
		return "", fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	var sVal string
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
	case "uint8":
		ui8Val = interfaceType.(uint8)
		buf.WriteString(fmt.Sprintf("%d", ui8Val))
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
		err = fmt.Errorf("func %s(): unknown type: %s\n", funcName(), table.colTypes[colIndex])
		return "", err
	}

	s = buf.String()

	return s, nil
}

func (table *Table) GetValAsString(colName string, rowIndex int) (string, error) {
	var colIndex int
	var err error

	if table == nil {
		return "", fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	colIndex, err = table.ColIndex(colName)
	if err != nil {
		return "", err
	}

	return table.GetValAsStringByColIndex(colIndex, rowIndex)
}

func (table *Table) IsStructShape() (bool, error) {
	if table == nil {
		return false, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	return table.structShape, nil
}

// Will be ignored (when writing table as string) if table RowCount() is more than 1
func (table *Table) SetStructShape(isStructShape bool) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	table.structShape = isStructShape

	return nil
}

// Join together a column of values. To compose a set of commands into a single command.
func (table *Table) JoinColVals(colName string, separator string) (string, error) {
	if table == nil {
		return "", fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	rowCount := table.RowCount()
	sVals := make([]string, rowCount)
	for rowIndex := 0; rowIndex < rowCount; rowIndex++ {
		s, err := table.GetValAsString(colName, rowIndex)
		if err != nil {
			return "", err
		}
		sVals[rowIndex] = s
	}

	joined := strings.Join(sVals, separator)

	return joined, nil
}

// Join together a column of values. To compose a set of commands into a single command.
func (table *Table) JoinColValsByColIndex(colIndex int, separator string) (string, error) {
	if table == nil {
		return "", fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	colName, err := table.ColName(colIndex)
	if err != nil {
		return "", err
	}

	return table.JoinColVals(colName, separator)
}

// Test that this value is a valid type for this column.
func (table *Table) IsValidColValue(colName string, value interface{}) (bool, error) {
	valueType := reflect.TypeOf(value)
	valueTypeName := valueType.Name()

	colType, err := table.ColType(colName)
	if err != nil {
		return false, err
	}

	if valueTypeName == colType {
		return true, nil
	} else {
		return false, fmt.Errorf("table[%s] col=%s type=%s invalid value: %v", table.Name(), colName, colType, value)
	}
}

/*
	- Pad trailing zeros on a string which is a floating point number.

	- Pad trailing zeros on the fractional part of a floating point number (which looks like an integer).
*/
func padTrailingZeros(s string) string {
	hasPoint := strings.Index(s, ".") >= 0
	b := []byte(s)
	for i := len(b)-1; i > 0; i-- {
		if (b[i-1] == '.' || b[i] != '0') {
			return string(b)
		}
		if hasPoint {
			// We don't want to remove zeros from floats with no decimal places (that look like an ints).
			if (b[i] == '0') {
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
	hasPoint := strings.Index(s, ".") >= 0
	b := []byte(s)
	for i := len(b)-1; i > 0; i-- {
		if (b[i-1] == '.' || b[i] != '0') {
			return string(b)
		}
		if hasPoint {
			// We don't want to remove zeros from floats with no decimal places (that look like an ints).
			if (b[i] == '0') {
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
	var typeOfCol reflect.Type
	colType, err := table.ColTypeByColIndex(colIndex)
	if err != nil {
		return nil, err
	}

	switch (colType) {
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
		err = fmt.Errorf("func %s(%q): unknown type: %s\n", funcName(), colType, table.colTypes[colIndex])
		return nil, err
	}

	return typeOfCol, nil
}

/*
// Doesn't work.
func (table *Table) AsSliceOfStruct(instanceOfStruct interface{}) (interface{}, error) {
	// See https://golang.org/pkg/reflect (search for: package path)
	var err error

	var val reflect.Value = reflect.ValueOf(instanceOfStruct)
//	fmt.Printf("val %T = %v\n", val, val)
	kind := val.Kind()
	fmt.Printf("kind = %T %v\n", kind, kind)
	if kind != reflect.Struct {
		err = fmt.Errorf("[%s].%s() is expecting a variable of type struct, not type %v", table.Name(), funcName(), kind)
		return instanceOfStruct, err
	}

//	valType := reflect.TypeOf(val)
	var sliceOfStruct []interface{}
	sliceOfStruct = reflect.MakeSlice(reflect.TypeOf(val), table.RowCount())

//	var emptySlice []reflect.Value
//
//	// Each column cell becomes a struct field.
//    fields := make([]reflect.StructField, table.ColCount())
//
//    for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
//
//        colName, err := table.ColName(colIndex)
//        if err != nil {
//            return emptySlice, err
//        }
//
//		if !IsExportableName(colName) {
//			err = fmt.Errorf("[%s].%s() cannot make struct: col name %s is an unexported (lowercase) identifier: %s",
//				table.Name(), funcName(), colName, colName)
//			return emptySlice, err
//		}
//
//        fields[colIndex].Name = colName
//
//        colReflectType, err := table.ReflectTypeOfColByColIndex(colIndex)
//        if err != nil {
//            return emptySlice, err
//        }
//        fields[colIndex].Type = colReflectType
//
//		// Add json tags for good measure. For later use.
//		tag := fmt.Sprintf("`json:%q`", colName)
//		fields[colIndex].Tag = reflect.StructTag(tag)
//    }
//
//	typ := reflect.StructOf(fields)
//	sliceOfStructOfFields := make([]reflect.Value, table.RowCount())
//
//	for rowIndex := 0; rowIndex < table.RowCount(); rowIndex++ {
//		structOfFields := reflect.New(typ).Elem()
//		for colIndex := 0; colIndex < table.ColCount(); colIndex++ {
//			var value interface{}
//			value, err = table.GetValByColIndex(colIndex, rowIndex)
//			if err != nil {
//				return emptySlice, err
//			}
//			var reflectValue reflect.Value
//			reflectValue = reflect.ValueOf(value)
//			structOfFields.Field(colIndex).Set(reflectValue)
//		}
//	}

	return instanceOfStruct, err
}
*/

// Helper function not used.
func isExportableName(name string) bool {
	rune0 := rune(name[0])
	if unicode.IsUpper(rune0) {
		return true
	} else {
		return false
	}
}

/*
	A strict table comparison:
	Column count must match.
	Row count must match.
	Column names must match.
	Column types must match.
	Rows order must match.
	Cell values must match.
	Column order need NOT match.

	Useful for testing.
*/
func (table1 *Table) Equals(table2 *Table) (bool, error) {
	var err error

	if table1 == nil {
		err = fmt.Errorf("func (table1 *Table) %s(table2 *Table): table1 is <nil>\n", funcName())
		return false, err
	}

	if table2 == nil {
		err = fmt.Errorf("func (table1 *Table) %s(table2 *Table): table2 is <nil>\n", funcName())
		return false, err
	}

	if table1.RowCount() != table2.RowCount() {
		err = fmt.Errorf("[%s].%s([%s]): [%s].RowCount() %d != [%s].RowCount() %d\n",
			table1.Name(), funcName(), table2.Name(), table1.Name(), table1.RowCount(), table2.Name(), table2.RowCount())
		return false, err
	}

	if table1.ColCount() != table2.ColCount() {
		err = fmt.Errorf("[%s].%s([%s]): [%s].ColCount() %d != [%s].ColCount() %d\n",
			table1.Name(), funcName(), table2.Name(), table1.Name(), table1.ColCount(), table2.Name(), table2.ColCount())
		return false, err
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
			err = fmt.Errorf("[%s].%s([%s]): [%s].ColType(%q) %s != [%s].ColType(%q) %s\n",
				table1.Name(), funcName(), table2.Name(), table1.Name(), colName, type1, table2.Name(), colName, type2)
			return false, err
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
				err = fmt.Errorf("[%s].%s([%s]): [%s].ColType(%q) %v != [%s].ColType(%q) %v\n",
					table1.Name(), funcName(), table2.Name(), table1.Name(), colName, val1, table2.Name(), colName, val2)
				return false, err
			}
		}
	}

	return true, nil
}
