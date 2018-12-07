/*
Functions and methods for sorting Table tables.
*/

package gotables

import (
	"fmt"
	"os"
	"sort"
	"strings"
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

type compareFunc func(i interface{}, j interface{}) int

var compareFuncs = map[string]compareFunc{
	"bool":    compare_bool,
	"float32": compare_float32,
	"float64": compare_float64,
	"uint":    compare_uint,
	"int":     compare_int,
	"int16":   compare_int16,
	"int32":   compare_int32,
	"int64":   compare_int64,
	"int8":    compare_int8,
	"string":  compare_Alphabetic_string,
	"uint16":  compare_uint16,
	"uint32":  compare_uint32,
	"uint64":  compare_uint64,
	"uint8":   compare_uint8,
}

type sortKey struct {
	colName  string
	colType  string
	reverse  bool
	sortFunc compareFunc
}

// For GOB encoding and GOB decoding, which requires items to be exported.
type SortKeyExported struct {
	ColName  string
	ColType  string
	Reverse  bool
	SortFunc compareFunc
}

func (key sortKey) String() string {
	return fmt.Sprintf("{colName:%q,colType:%q,reverse:%t}", key.colName, key.colType, key.reverse)
}

type SortKeys []sortKey

func (keys SortKeys) String() string {
	if keys == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: %s(SortKeys) SortKeys is <nil>\n", funcSource(), funcName()))
		return ""
	}
	var s string = "SortKeys["
	keySep := ""
	for _, key := range keys {
		s += keySep + key.String()
		keySep = ","
	}
	s += "]"
	return s
}

// Returns a copy of the sort keys as a Table. Useful for debugging.
func (thisTable *Table) GetSortKeysAsTable() (*Table, error) {
	if thisTable == nil {
		return nil, fmt.Errorf("table.%s() table is <nil>", funcName())
	}
	var keysTable *Table
	var err error
	keysTable, err = NewTable("SortKeys")
	if err != nil {
		return nil, err
	}
	if err = keysTable.AppendCol("index", "int"); err != nil {
		return nil, err
	}
	err = keysTable.AppendCol("colName", "string")
	if err != nil {
		return nil, err
	}
	err = keysTable.AppendCol("colType", "string")
	if err != nil {
		return nil, err
	}
	err = keysTable.AppendCol("reverse", "bool")
	if err != nil {
		return nil, err
	}
	for rowIndex := 0; rowIndex < len(thisTable.sortKeys); rowIndex++ {
		err = keysTable.AppendRow()
		if err != nil {
			return nil, err
		}
		if err = keysTable.SetInt("index", rowIndex, rowIndex); err != nil {
			return nil, err
		}
		err = keysTable.SetString("colName", rowIndex, thisTable.sortKeys[rowIndex].colName)
		if err != nil {
			return nil, err
		}
		err = keysTable.SetString("colType", rowIndex, thisTable.sortKeys[rowIndex].colType)
		if err != nil {
			return nil, err
		}
		err = keysTable.SetBool("reverse", rowIndex, thisTable.sortKeys[rowIndex].reverse)
		if err != nil {
			return nil, err
		}
	}

	return keysTable, nil
}

/*
Call with an argument list, or a slice of string followed by an ellipsis ...

(1) Pass sort keys as separate arguments:
	err = table.SetSortKeys("col1","col2","col3")

(2) Pass sort keys as a slice:
	err = table.SetSortKeys([]string{"col1","col2","col3"}...)

(3) Pass sort keys as a slice:
	sortColNames := []string{"col1","col2","col3"}
	err = table.SetSortKeys(sortColNames...)

(4) Clear sort keys (if any) by calling with empty argument list:
	err = table.SetSortKeys()
*/
func (table *Table) SetSortKeys(sortColNames ...string) error {

	if table == nil {
		return fmt.Errorf("table.%s() table is <nil>", funcName())
	}

	table.sortKeys = newSortKeys() // Replace any existing sort keys.

	for _, colName := range sortColNames {
		err := table.AppendSortKey(colName)
		if err != nil {
//			errSortKey := errors.New(fmt.Sprintf("SetSortKeys(%v): %v\n", sortColNames, err))
			errSortKey := fmt.Errorf("SetSortKeys(%v): %v\n", sortColNames, err)
			return errSortKey
		}
	}

	return nil
}

/*
Call with an argument list, or a slice of string followed by ...

Example 1: SetSortKeysReverse("col1","col3")

Example 2: SetSortKeysReverse([]string{"col1","col3"}...)
*/
func (table *Table) SetSortKeysReverse(reverseSortColNames ...string) error {
	if table == nil {
		return fmt.Errorf("table.%s() table is <nil>", funcName())
	}

	for _, colName := range reverseSortColNames {
		err := table.setSortKeyReverse(colName)
		if err != nil {
			errSortKey := fmt.Errorf("SetSortKeysReverse(%v): %v\n", reverseSortColNames, err)
			return errSortKey
		}
	}
	return nil
}

func (table *Table) setSortKeyReverse(colName string) error {
	if table == nil {
		return fmt.Errorf("table.%s() table is <nil>", funcName())
	}
	if len(table.sortKeys) == 0 {
		err := fmt.Errorf("must call SetSortKeys() before calling SetSortKeysReverse()")
		return err
	}
	var found bool = false
	for i, sortKey := range table.sortKeys {
		if sortKey.colName == colName {
			table.sortKeys[i].reverse = true
			found = true
		}
	}
	if !found {
		err := fmt.Errorf("sortKey not found: %q", colName)
		return err
	}

	return nil
}

func (table *Table) AppendSortKey(colName string) error {
	if table == nil {
		return fmt.Errorf("table.%s() table is <nil>", funcName())
	}
	colInfo, err := table.getColInfo(colName)
	if err != nil {
		// Col doesn't exist.
		return err
	}

	var key sortKey
	key.colName = colName

	var colType = colInfo.colType
	if len(colType) == 0 {
		return fmt.Errorf("table [%s]: unknown colType for col: %q", table.Name(), colName)
	}
	key.colType = colType

	sortFunc, exists := compareFuncs[colType]
	if !exists { // Error occurs only during software development if a type has not been handled.
		return fmt.Errorf("table [%s] col %q: compareFunc compare_%s has not been defined for colType: %q",
			table.Name(), colName, colType, colType)
	}

	key.sortFunc = sortFunc
	table.sortKeys = append(table.sortKeys, key)

	return nil
}

// Delete a sort key by name.
func (table *Table) DeleteSortKey(keyName string) error {
	if table == nil {
		return fmt.Errorf("table.%s() table is <nil>", funcName())
	}
	_, err := table.getColInfo(keyName)
	if err != nil {
		// Col doesn't exist.
		return err
	}

	for keyIndex := 0; keyIndex < len(table.sortKeys); keyIndex++ {
		if table.sortKeys[keyIndex].colName == keyName {
			// From Ivo Balbaert p182 for deleting a single element.
			table.sortKeys = append(table.sortKeys[:keyIndex], table.sortKeys[keyIndex+1:]...)
			return nil
		}
	}

	return fmt.Errorf("[%s].%s(%q) sort key not found: %q", table.Name(), funcName(), keyName, keyName)
}

func (table *Table) getColNames() []string {
	if table == nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s ERROR: table.%s() table is <nil>\n", funcSource(), funcName()))
		return nil
	}
	return table.colNames
}

// Sorting functions:

var compare_Alphabetic_string compareFunc = func(i, j interface{}) int {
	var si_string string = i.(string)
	var sj_string string = j.(string)
	var si_lower string = strings.ToLower(si_string)
	var sj_lower string = strings.ToLower(sj_string)
	if si_lower < sj_lower {
		return -1
	} else if si_lower > sj_lower {
		return +1
	} else { // si_lower == sj_lower
		if si_string < sj_string {
			return -1
		} else if si_string > sj_string {
			return +1
		} else {
			return 0
		}
	}
}

var compare_uint compareFunc = func(i, j interface{}) int {
	var inti uint = i.(uint)
	var intj uint = j.(uint)
	if inti < intj {
		return -1
	} else if inti > intj {
		return +1
	} else {
		return 0
	}
}

var compare_int compareFunc = func(i, j interface{}) int {
	var inti int = i.(int)
	var intj int = j.(int)
	if inti < intj {
		return -1
	} else if inti > intj {
		return +1
	} else {
		return 0
	}
}

var compare_int8 compareFunc = func(i, j interface{}) int {
	var int8i int8 = i.(int8)
	var int8j int8 = j.(int8)
	if int8i < int8j {
		return -1
	} else if int8i > int8j {
		return +1
	} else {
		return 0
	}
}

var compare_int16 compareFunc = func(i, j interface{}) int {
	var int16i int16 = i.(int16)
	var int16j int16 = j.(int16)
	if int16i < int16j {
		return -1
	} else if int16i > int16j {
		return +1
	} else {
		return 0
	}
}

var compare_int32 compareFunc = func(i, j interface{}) int {
	var int32i int32 = i.(int32)
	var int32j int32 = j.(int32)
	if int32i < int32j {
		return -1
	} else if int32i > int32j {
		return +1
	} else {
		return 0
	}
}

var compare_int64 compareFunc = func(i, j interface{}) int {
	var int64i int64 = i.(int64)
	var int64j int64 = j.(int64)
	if int64i < int64j {
		return -1
	} else if int64i > int64j {
		return +1
	} else {
		return 0
	}
}

var compare_uint8 compareFunc = func(i, j interface{}) int {
	var uint8i uint8 = i.(uint8)
	var uint8j uint8 = j.(uint8)
	if uint8i < uint8j {
		return -1
	} else if uint8i > uint8j {
		return +1
	} else {
		return 0
	}
}

var compare_uint16 compareFunc = func(i, j interface{}) int {
	var uint16i uint16 = i.(uint16)
	var uint16j uint16 = j.(uint16)
	if uint16i < uint16j {
		return -1
	} else if uint16i > uint16j {
		return +1
	} else {
		return 0
	}
}

var compare_uint32 compareFunc = func(i, j interface{}) int {
	var uint32i uint32 = i.(uint32)
	var uint32j uint32 = j.(uint32)
	if uint32i < uint32j {
		return -1
	} else if uint32i > uint32j {
		return +1
	} else {
		return 0
	}
}

var compare_uint64 compareFunc = func(i, j interface{}) int {
	var uint64i uint64 = i.(uint64)
	var uint64j uint64 = j.(uint64)
	if uint64i < uint64j {
		return -1
	} else if uint64i > uint64j {
		return +1
	} else {
		return 0
	}
}

var compare_float32 compareFunc = func(i, j interface{}) int {
	var float32i float32 = i.(float32)
	var float32j float32 = j.(float32)
	if float32i < float32j {
		return -1
	} else if float32i > float32j {
		return +1
	} else {
		return 0
	}
}

var compare_float64 compareFunc = func(i, j interface{}) int {
	var float64i float64 = i.(float64)
	var float64j float64 = j.(float64)
	if float64i < float64j {
		return -1
	} else if float64i > float64j {
		return +1
	} else {
		return 0
	}
}

var compare_bool compareFunc = func(i, j interface{}) int {
	var booli bool = i.(bool)
	var boolj bool = j.(bool)
	if !booli && boolj {
		return -1
	} else if booli && !boolj {
		return +1
	} else {
		return 0
	}
}

/*
	Sort this table by this table's currently-set sort keys.

	To see the currently-set sort keys use GetSortKeysAsTable()
*/
func (table *Table) Sort() error {

	if table == nil {
		return fmt.Errorf("table.%s() table is <nil>", funcName())
	}

	if len(table.sortKeys) == 0 {
		return fmt.Errorf("%s() cannot sort table that has 0 sort keys - use SetSortKeys()", funcName())
	}

	table.sortByKeys(table.sortKeys)

	return nil
}

/*
	Sort by one or more columns ascending-only.

	1. All column keys are set to ascending order.

	2. One or more column keys must be provided.

	3. To sort one or more columns in reverse (eg with "key2" reversed):

	table.SetSortKeys("key1", "key2", "key3")

	table.SetSortKeysReverse("key2")

	table.Sort()

	4. SortSimple() sets the table's sort keys, so subsequent calls to table.Sort() will have the same effect
	as calling table.SetSortKeys() and then table.Sort()
*/
func (table *Table) SortSimple(sortCols ...string) error {

	if table == nil {
		return fmt.Errorf("table.%s() table is <nil>", funcName())
	}

	if len(sortCols) == 0 {
		return fmt.Errorf("%s() cannot sort table using 0 sortCols", funcName())
	}

	err := table.SetSortKeys(sortCols...)
	if err != nil { return err }

	table.sortByKeys(table.sortKeys)

	return nil
}

type tableSortable struct {
	table *Table
	rows  []tableRow
	less  func(i tableRow, j tableRow) bool
}

func (table tableSortable) Len() int { return len(table.rows) }

func (table tableSortable) Swap(i int, j int) {
	table.rows[i], table.rows[j] = table.rows[j], table.rows[i]
}

func (table tableSortable) Less(i int, j int) bool {
	return table.less(table.rows[i], table.rows[j])
}

func (table *Table) sortByKeys(sortKeys SortKeys) {
	sort.Sort(tableSortable{table, table.rows, func(iRow, jRow tableRow) bool {
//		compareCount++
		for _, sortKey := range table.sortKeys {
			var colName string = sortKey.colName
			colIndex, _ := table.ColIndex(colName)
			var sortFunc compareFunc = sortKey.sortFunc
			var iVal interface{} = iRow[colIndex]
			var jVal interface{} = jRow[colIndex]
			var compared int = sortFunc(iVal, jVal)
			if sortKey.reverse {
				// Reverse the sign to reverse the sort.
				// Reverse is intended to be descending, and not a toggle between ascending and descending.
				compared *= -1
			}
			if compared != 0 {
				return compared < 0		// Less is true if compared < 0
			}
		}
		return false
	}})
}

func (table *Table) checkSearchArguments(searchValues ...interface{}) error {
	if table == nil {
		return fmt.Errorf("table.%s() table is <nil>", funcName())
	}

	if len(searchValues) == 0 {
		return fmt.Errorf("[%s].Search(...) expecting 1 or more search values, but found none", table.Name())
	}

	if len(table.sortKeys) == 0 {
		return fmt.Errorf("cannot search table that has 0 sort keys - use SetSortKeys()")
	}

	// Test for special case where Sort() has been passed a slice without ... instead of comma-separated args.
	if len(searchValues) == 1 && len(table.sortKeys) > 1 {
		return fmt.Errorf("%s() searchValues count %d != sort keys count %d  If passing a slice use ellipsis syntax: Search(mySliceOfKeys...)",
			funcName(), len(searchValues), len(table.sortKeys))
	}

	if len(searchValues) != len(table.sortKeys) {
		return fmt.Errorf("%s() searchValues count %d != sort keys count %d",
			funcName(), len(searchValues), len(table.sortKeys))
	}

	// Check that searchValues are the right type.
	for sortIndex, sortKey := range table.sortKeys {
		colName := sortKey.colName
		value := searchValues[sortIndex]
		isValid, err := table.IsValidColValue(colName, value)
		if !isValid {
			// Append key name and type information to end of err.
			var keyInfo string
			sep := ""
			for _, sortKey := range table.sortKeys {
				keyInfo += fmt.Sprintf("%s%s %s", sep, sortKey.colName, sortKey.colType)
				sep = ", "
			}
			return fmt.Errorf("%v (valid key type%s: %s)", err, plural(len(table.sortKeys)), keyInfo)
		}
	}

	return nil
}

/*
	Search this table by this table's currently-set sort keys.

	To see the currently-set sort keys use GetSortKeysAsTable()

	Note: This calls *Table.SearchFirst() which returns the first (if any) match in the table.
	Search first is what the Go sort.Search() function does.
*/
func (table *Table) Search(searchValues ...interface{}) (int, error) {
   	return table.SearchFirst(searchValues...)
}

/*
	Search this table by this table's currently-set sort keys.

	To see the currently-set sort keys use GetSortKeysAsTable()
*/
func (table *Table) SearchFirst(searchValues ...interface{}) (int, error) {

	err := table.checkSearchArguments(searchValues...)
	if err != nil {
		return -1, err
	}

	rowIndex, err := table.searchByKeysFirst(searchValues)

	return rowIndex, err
}

func (table *Table) searchByKeysFirst(searchValues []interface{}) (int, error) {

	var searchIndex int = -1

	// sort.Search() is enclosed (enclosure) here so it can access table values.
	searchIndex = SearchFirst(table.RowCount(), func(rowIndex int) bool {	// Locally-defined Search() function
		var keyCount = len(table.sortKeys)
		var keyLast = keyCount - 1
		var compared int
		for keyIndex, sortKey := range table.sortKeys {
			var colName string = sortKey.colName
			var sortFunc compareFunc = sortKey.sortFunc
			var searchVal interface{} = searchValues[keyIndex]
			var cellVal interface{}
			cellVal, err := table.GetVal(colName, rowIndex)
			if err != nil {
				// Should never happen. Hasn't been tested.
				break	// Out to searchByKeys() enclosing function.
			}
			compared = sortFunc(cellVal, searchVal)

			if sortKey.reverse {
				// Reverse the sign to reverse the sort.
				compared *= -1
			}

			// Most searches will be single-key searches, so last key is the most common.
			if keyIndex == keyLast {	// Last key is the deciding key because all previous keys matched.
				return compared >= 0
			} else {
				if compared > 0 {	// Definite result regardless of subsequent keys: no match.
					return true
				} else if compared < 0 {
					return false	// Definite result regardless of subsequent keys: no match.
				}
			}
			// Otherwise the first keys are equal, so keep looping through keys.
		}

		// Should never be reached. Hasn't been tested.
		return false
	})

	// See logic at: https://golang.org/pkg/sort/#Search
	// See Search() source code at: https://golang.org/src/sort/search.go?s=2247:2287#L49
	if searchIndex < table.RowCount() && searchValuesMatchRowValues(table, searchValues, searchIndex) {
		return searchIndex, nil
	} else {
		return -1, fmt.Errorf("[%s].Search(%v) search values not in table: %v",
			table.Name(), searchValues, searchValues)
	}
}

// Compare search values with row values to determine if search was successful or not.
func searchValuesMatchRowValues(table *Table, searchValues []interface{}, searchIndex int) bool {
	// Loop through the parallel lists of sort keys and search values.
	for i := 0; i < len(table.sortKeys); i++ {
		colName    := table.sortKeys[i].colName
		sortFunc   := table.sortKeys[i].sortFunc
		cellVal, _ := table.GetVal(colName, searchIndex)
		searchVal  := searchValues[i]
		compared   := sortFunc(cellVal, searchVal)
		if compared != 0 {
			// At least one search value doesn't match a cell value.
			return false
		}
	}

	// They all match.
	return true
}

/*
	Compare two rows using table sort keys.

		Return -1 if rowIndex1 is less than rowIndex2.
		Return  0 if rowIndex1 equals rowIndex2.
		Return  1 if rowIndex1 is greater than rowIndex2.
		Return -2 if error.
*/
func (table *Table) CompareRows(rowIndex1 int, rowIndex2 int) (int, error) {

	if table == nil {
		return -2, fmt.Errorf("table.%s() table is <nil>", funcName())
	}

	if rowIndex1 < 0 || rowIndex1 > table.RowCount()-1 {
		return -2, fmt.Errorf("[%s].%s(%d, %d) in table [%s] with %d rows, row index %d does not exist",
			table.Name(), funcName(), rowIndex1, rowIndex2, table.Name(), table.RowCount(), rowIndex1)
	}

	if rowIndex2 < 0 || rowIndex2 > table.RowCount()-1 {
		return -2, fmt.Errorf("[%s].%s(%d, %d) in table [%s] with %d rows, row index %d does not exist",
			table.Name(), funcName(), rowIndex1, rowIndex2, table.Name(), table.RowCount(), rowIndex2)
	}

	if len(table.sortKeys) == 0 {
		return -2, fmt.Errorf("[%s].%s(%d, %d) table has 0 sort keys - use table.SetSortKeys()",
			table.Name(), funcName(), rowIndex1, rowIndex2)
	}

	// Loop through the parallel lists of sort keys and search values.
	for i := 0; i < len(table.sortKeys); i++ {
		sortFunc    := table.sortKeys[i].sortFunc
		colName     := table.sortKeys[i].colName
		cellVal1, _ := table.GetVal(colName, rowIndex1)
		cellVal2, _ := table.GetVal(colName, rowIndex2)
		compared    := sortFunc(cellVal1, cellVal2)

		if compared != 0 {
			// We have a decision: At least one search value doesn't match a cell value.
			return compared, nil
		}
	}

	// They all match. Means they're equal.
	return 0, nil	// Equal.
}

// Factory function to generate a slice of SortKeys.
func newSortKeys() SortKeys {
	return []sortKey{}
}

func (table *Table) SortKeyCount() int {
	return len(table.sortKeys)
}

// Copy sort keys into table from fromTable.
func (table *Table) SetSortKeysFromTable(fromTable *Table) error {
	if table == nil {
		return fmt.Errorf("table.%s() table is <nil>", funcName())
	}
	if fromTable == nil {
		return fmt.Errorf("fromTable.%s() fromTable is <nil>", funcName())
	}
	if fromTable.SortKeyCount() == 0 {
		return fmt.Errorf("table.%s(fromTable): fromTable.SortKeyCount() == 0", funcName())
	}

	var err error
	var ascending []string	// They default to ascending, and may be later reversed.
	var descending []string

	keysTable, err := fromTable.GetSortKeysAsTable()
	if err != nil {
		return err
	}

	for rowIndex := 0; rowIndex < keysTable.RowCount(); rowIndex++ {

		var colName string
		colName, err = keysTable.GetString("colName", rowIndex)
		if err != nil {
			return err
		}

		var reverse bool
		reverse, err = keysTable.GetBool("reverse", rowIndex)
		if err != nil {
			return err
		}

		ascending = append(ascending, colName)
		if reverse {
			descending = append(descending, colName)
		}
	}

	err = table.SetSortKeys(ascending...)
	if err != nil {
		return err
	}

	err = table.SetSortKeysReverse(descending...)
	if err != nil {
		return err
	}

	return nil
}

/*
	Move sort key columns to the left of the table, and into sort key order.

	Note: This is purely for human readability. It is not required for sorting.
*/
func (table *Table) OrderColsBySortKeys() error {
	if table == nil {
		return fmt.Errorf("table.%s() table is <nil>", funcName())
	}

	var err error
	var newOrder []string = make([]string, table.ColCount())	// List of new colNames.
	var key int

	// Populate new order ...

	// First slots with key col names.
	for key = 0; key < table.SortKeyCount(); key++ {
		keyName := table.sortKeys[key].colName
		newOrder[key] = keyName
	}

	// Subsequent slots with non-key col names.
	row := table.SortKeyCount()
	for col := 0; col < table.ColCount(); col++ {
		colName := table.colNames[col]
		var isSortKey bool
		isSortKey, err = table.IsSortKey(colName)
		if err != nil {
			return err
		}
		if !isSortKey {
			newOrder[row] = colName
			row++
		}
	}

	err = table.ReorderCols(newOrder...)
	if err != nil {
		return err
	}

	return err
}

// True if colName is a sort key in table. False if not. Error if colName not in table.
func (table *Table) IsSortKey(colName string) (bool, error) {
	if table == nil { return false, fmt.Errorf("table.%s() table is <nil>", funcName()) }

	hasCol, err := table.HasCol(colName)
	if err != nil {
		return hasCol, err
	}

	for keyIndex := 0; keyIndex < len(table.sortKeys); keyIndex++ {
		if table.sortKeys[keyIndex].colName == colName {
			return true, nil
		}
	}

	return false, nil
}

/*
	Swap these two columns with each other.
*/
func (table *Table) SwapColsByColIndex(colIndex1 int, colIndex2 int) error {
	// This sets out the relationship between table.colNames, table.colTypes and table.colNamesMap.
	if table == nil { return fmt.Errorf("table.%s() table is <nil>", funcName()) }

	var err error

	if colIndex1 < 0 || colIndex1 > table.ColCount()-1 {
		return fmt.Errorf("[%s].%s: in table [%s] with %d cols, colIndex1 %d does not exist",
			table.tableName, funcName(), table.tableName, table.ColCount(), colIndex1)
	}

	if colIndex2 < 0 || colIndex2 > table.ColCount()-1 {
		return fmt.Errorf("[%s].%s: in table [%s] with %d cols, colIndex2 %d does not exist",
			table.tableName, funcName(), table.tableName, table.ColCount(), colIndex2)
	}

	if colIndex1 == colIndex2 {
		return fmt.Errorf("[%s].%s: [%s] colIndex1 %d == colIndex2 %d",
			table.tableName, funcName(), table.tableName, colIndex1, colIndex2)
	}

	table.colNames[colIndex1], table.colNames[colIndex2] = table.colNames[colIndex2], table.colNames[colIndex1]

	table.colTypes[colIndex1], table.colTypes[colIndex2] = table.colTypes[colIndex2], table.colTypes[colIndex1]

	colName1, err := table.ColName(colIndex1)
	if err != nil { return err }

	colName2, err := table.ColName(colIndex2)
	if err != nil { return err }

	table.colNamesMap[colName1], table.colNamesMap[colName2] = table.colNamesMap[colName2], table.colNamesMap[colName1]

	return nil
}

/*
	Swap these two columns with each other.
*/
func (table *Table) SwapCols(colName1 string, colName2 string) error {
	// This sets out the relationship between table.colNames, table.colTypes and table.colNamesMap.
	if table == nil { return fmt.Errorf("table.%s() table is <nil>", funcName()) }

	var err error

	col1, err := table.ColIndex(colName1)
	if err != nil { return err }

	col2, err := table.ColIndex(colName2)
	if err != nil { return err }

	if colName1 == colName2 {
		fmt.Errorf("[%s].%s: [%s] colName1 %q == colName2 %q",
			table.tableName, funcName(), table.tableName, colName1, colName2)
	}

	table.colNames[col1], table.colNames[col2] = table.colNames[col2], table.colNames[col1]

	table.colTypes[col1], table.colTypes[col2] = table.colTypes[col2], table.colTypes[col1]

	table.colNamesMap[colName1], table.colNamesMap[colName2] = table.colNamesMap[colName2], table.colNamesMap[colName1]

	return nil
}

/*
	Search this table by this table's currently-set sort keys.

	To see the currently-set sort keys use GetSortKeysAsTable()
*/
   func (table *Table) SearchLast(searchValues ...interface{}) (int, error) {

	err := table.checkSearchArguments(searchValues...)
	if err != nil {
		return -1, err
	}

	rowIndex, err := table.searchByKeysLast(searchValues)

	return rowIndex, err
}

func (table *Table) searchByKeysLast(searchValues []interface{}) (int, error) {

	var searchIndex int = -1

	// sort.Search() is enclosed (enclosure) here so it can access table values.
	searchIndex = SearchLast(table.RowCount(), func(rowIndex int) bool {	// Locally-defined Search() function
		var keyCount = len(table.sortKeys)
		var keyLast = keyCount - 1
		var compared int
		for keyIndex, sortKey := range table.sortKeys {
			var colName string = sortKey.colName
			var sortFunc compareFunc = sortKey.sortFunc
			var searchVal interface{} = searchValues[keyIndex]
			var cellVal interface{}
			cellVal, err := table.GetVal(colName, rowIndex)
			if err != nil {
				// Should never happen. Hasn't been tested.
				break	// Out to searchByKeys() enclosing function.
			}
			compared = sortFunc(cellVal, searchVal)

			if sortKey.reverse {
				// Reverse the sign to reverse the sort.
				compared *= -1
			}

			// Most searches will be single-key searches, so last key is the most common.
			if keyIndex == keyLast {	// Last key is the deciding key because all previous keys matched.
				return compared <= 0
			} else {
				if compared < 0 {	// Definite result regardless of subsequent keys: no match.
					return true
				} else if compared > 0 {
					return false	// Definite result regardless of subsequent keys: no match.
				}
			}
			// Otherwise the first keys are equal, so keep looping through keys.
		}

		// Should never be reached. Hasn't been tested.
		return false
	})

	// See logic at: https://golang.org/pkg/sort/#Search
	// See Search() source code at: https://golang.org/src/sort/search.go?s=2247:2287#L49
	if searchIndex < table.RowCount() && searchValuesMatchRowValues(table, searchValues, searchIndex) {
		return searchIndex, nil
	} else {
		return -1, fmt.Errorf("[%s].Search(%v) search values not in table: %v",
			table.Name(), searchValues, searchValues)
	}
}

// This is lifted straight from See https://golang.org/pkg/sort/#Search
// Used to write the mirror version.
func search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	var h int
	for i < j {
		h = i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
//	fmt.Printf("%2d %2d %2d\n", i, h, j)
	return i
}

/*
	gotables.SearchLast() mirrors the Go library function sort.Search()

    	Comparison of sort.Search() and gotables.SearchLast() describing their mirrored relationship.
		Assume data is a zero-based array/slice of elements sorted in ascending order.
		Each search function returns an index into data.
		----------------------------------------------------------------------------------------------------------
		Go library sort.Search()                         |  gotables.SearchLast()
		----------------------------------------------------------------------------------------------------------
		Index to greater than or equal to search term.   |  Index to less than or equal to search term.
		Finds index of FIRST instance equal.             |  Finds index of LAST instance equal.
		Multiple instances will be at and AFTER index.   |  Multiple instances will be at and BEFORE index.
		if term is missing from data, where it WOULD be  |  If term is missing from data, where it WOULD be
		  is insert BEFORE index.                        |    is insert AFTER index.
		Missing at high end of data returns              |  Missing at low end of data returns
		  index 1-greater than last element, len(data),  |    index 1-less than first element, -1,
		  which means it would insert BEFORE len(data),  |    which means it would insert AFTER -1 data,
		  which would be an append to data array.        |    which would be an insert to before start of data array.
		  Check index to avoid out of bounds errors.     |    Check index != -1 to avoid out of bounds errors.
		Example: multiple search terms present           |  Example: multiple search terms present
		  data: [4 8 10 10 10 20 23 29]                  |    data: [4 8 10 10 10 20 23 29]
		  index: 0 1  2  3  4  5  6  7                   |    index: 0 1  2  3  4  5  6  7
		  x: 10                                          |    x: 10
		  sort.Search(x, func) = 2 (finds FIRST)         |    gotables.SearchLast(x, func) = 4 (finds LAST)
		----------------------------------------------------------------------------------------------------------

		This binary search has two steps: (1) binary search for x, and (2) check if x was found.

		Strange, huh? Go library sort.Search() works the same way, except in the opposite (mirror image) direction.
		See https://golang.org/pkg/sort/#Search

		(1) Binary search for x.
		x := 23
		i := gotables.SearchLast(len(data), func(i int) bool { return data[i] <= x })

		(2) Check that x was found.
		if i >= 0 && data[i] == x {
			// x is present at data[i]
		} else {
			// x is not present in data,
			// but i is the index where it would be inserted.
			// Note that i can be -1 which does not exist in data.
		}
*/
func SearchLast(n int, f func(int) bool) int {
	i, j := -1, n
	var h int
	for i < j-1 {
		h = i + (j-i)/2
		if !f(h) {
			j = h
		} else {
			i = h
		}
	}
	return i
}

/*
	gotables.SearchFirst() calls the Go library function sort.Search()

    	Comparison of sort.Search() and gotables.SearchLast() describing their mirrored relationship.
		Assume data is a zero-based array/slice of elements sorted in ascending order.
		Each search function returns an index into data.
		----------------------------------------------------------------------------------------------------------
		Go library sort.Search()                         |  gotables.SearchLast()
		----------------------------------------------------------------------------------------------------------
		Index to greater than or equal to search term.   |  Index to less than or equal to search term.
		Finds index of FIRST instance equal.             |  Finds index of LAST instance equal.
		Multiple instances will be at and AFTER index.   |  Multiple instances will be at and BEFORE index.
		if term is missing from data, where it WOULD be  |  If term is missing from data, where it WOULD be
		  is insert BEFORE index.                        |    is insert AFTER index.
		Missing at high end of data returns              |  Missing at low end of data returns
		  index 1-greater than last element, len(data),  |    index 1-less than first element, -1,
		  which means it would insert BEFORE len(data),  |    which means it would insert AFTER -1 data,
		  which would be an append to data array.        |    which would be an insert to before start of data array.
		  Check index to avoid out of bounds errors.     |    Check index != -1 to avoid out of bounds errors.
		Example: multiple search terms present           |  Example: multiple search terms present
		  data: [4 8 10 10 10 20 23 29]                  |    data: [4 8 10 10 10 20 23 29]
		  index: 0 1  2  3  4  5  6  7                   |    index: 0 1  2  3  4  5  6  7
		  x: 10                                          |    x: 10
		  sort.Search(x, func) = 2 (finds FIRST)         |    gotables.SearchLast(x, func) = 4 (finds LAST)
		----------------------------------------------------------------------------------------------------------

		This binary search has two steps: (1) binary search for x, and (2) check if x was found.

		Strange, huh? Go library sort.Search() works the same way, except in the opposite (mirror image) direction.
		See https://golang.org/pkg/sort/#Search

		(1) Binary search for x.
		x := 23
		i := gotables.SearchFirst(len(data), func(i int) bool { return data[i] >= x })

		(2) Check that x was found.
		if i < len(data) && data[i] == x {
			// x is present at data[i]
		} else {
			// x is not present in data,
			// but i is the index where it would be inserted.
			// Note that i can be len(data) which does not exist in data.
		}
*/
func SearchFirst(n int, f func(int) bool) int {
	return sort.Search(n, f)
}

/*
	Return the first and last index of 1 or more instances of searchValues in table.

	Note: searchValues must match, not merely GT and LT where match would be.
*/
func (table *Table) SearchRange(searchValues ...interface{}) (firstRow int, lastRow int, err error) {

	firstRow, lastRow = -1, -1

	err = table.checkSearchArguments(searchValues...)
	if err != nil {
		return
	}

	firstRow, err = table.SearchFirst(searchValues...)
	if err != nil {
		return
	}

	lastRow, err = table.SearchLast(searchValues...)
	if err != nil {
		return
	}

	return
}

/*
	Sort this table UNIQUE by this table's currently-set sort keys.
	Non-key column cell values are not used for uniqueness but are evaluated and merged.

	To see the currently-set sort keys use GetSortKeysAsTable()

	Note: the sorted table is returned, not sorted in place.

	SortUnique() uses table.Merge() which resolves missing or conflicting cell values.
*/
func (table *Table) SortUnique() (tableUnique *Table, err error) {

	if table == nil {
		return nil, fmt.Errorf("table.%s() table is <nil>", funcName())
	}

	if len(table.sortKeys) == 0 {
		return nil, fmt.Errorf("%s() cannot sort table that has 0 sort keys - use SetSortKeys()", funcName())
	}

	// Merge() eliminates duplicates (based on keys) and fills in zero and missing values where available.
	tableUnique, err = table.Merge(table)
	if err != nil {
		return nil, err
	}

	// Merge() calls the merged table "merged". Rename it.
	err = tableUnique.SetName(table.Name())
	if err != nil {
		return nil, err
	}

	return tableUnique, nil
}
