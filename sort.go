/*
Functions and methods for sorting Table tables.
*/

package gotable

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
	"string":  compareAlphabetic_string,
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
		os.Stderr.WriteString(fmt.Sprintf("ERROR: %s(SortKeys) SortKeys is <nil>\n", funcName()))
		return ""
	}
	// where(fmt.Sprintf("len(keys) = %d\n", len(keys)))
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
		return nil, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	var keysTable *Table
	var err error
	keysTable, err = NewTable("sortKeys")
	if err != nil {
		return nil, err
	}
	if err = keysTable.AppendCol("key", "int"); err != nil {
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
		if err = keysTable.SetInt("key", rowIndex, rowIndex); err != nil {
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
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	table.sortKeys = newSortKeys() // Replace any existing sort keys.

	for _, colName := range sortColNames {
		err := table.AppendSortKey(colName)
		if err != nil {
//			errSortKey := errors.New(fmt.Sprintf("SetSortKeys(%v): %v\n", sortColNames, err))
			errSortKey := fmt.Errorf("SetSortKeys(%v): %v\n", sortColNames, err)
			// where(fmt.Sprintf("ERROR IN SetSortKeys(): %v", errSortKey))
			return errSortKey
		}
	}
	//where(fmt.Sprintf("table.sortKeys === %v\n", table.sortKeys))

	return nil
}

/*
Call with an argument list, or a slice of string followed by ...

Example 1: SetSortKeysReverse("col1","col3")

Example 2: SetSortKeysReverse([]string{"col1","col3"}...)
*/
func (table *Table) SetSortKeysReverse(reverseSortColNames ...string) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	for _, colName := range reverseSortColNames {
		err := table.setSortKeyReverse(colName)
		if err != nil {
			errSortKey := fmt.Errorf("SetSortKeysReverse(%v): %v\n", reverseSortColNames, err)
			return errSortKey
		}
	}
	//where(fmt.Sprintf("table.sortKeys === %v\n", table.sortKeys))
	return nil
}

func (table *Table) setSortKeyReverse(colName string) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	if len(table.sortKeys) == 0 {
		err := fmt.Errorf("must call SetSortKeys() before calling SetSortKeysReverse()")
		return err
	}
	var found bool = false
	// where(fmt.Sprintf("******** sortKeys = %v ...\n", table.sortKeys))
	for i, sortKey := range table.sortKeys {
		if sortKey.colName == colName {
			table.sortKeys[i].reverse = true
			found = true
		}
	}
	// where(fmt.Sprintf("******** ... sortKeys = %v\n", table.sortKeys))
	if !found {
		err := fmt.Errorf("sortKey not found: %q", colName)
		return err
	}

	return nil
}

func (table *Table) AppendSortKey(colName string) error {
	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	//	where(fmt.Sprintf("AppendSortKey(%q)\n", colName))
	colInfo, err := table.colInfo(colName)
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

/*
func (table *Table) getSortKeys() (SortKeys, error) {
	if table == nil {
		return nil, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}
	return table.sortKeys, nil
}
*/

func (table *Table) getColNames() []string {
	if table == nil {
		os.Stderr.WriteString(fmt.Sprintf("ERROR: %s(*Table) *Table is <nil>\n", funcName()))
		return nil
	}
	return table.colNames
}

// Sorting functions:

/*
func (tableRows tableRows) Len() int {
	return len(tableRows)
}
*/

/*
func (tableRows tableRows) Swap(i, j int) {
	tableRows[i], tableRows[j] = tableRows[j], tableRows[i]
}
*/

var compareAlphabetic_string compareFunc = func(i, j interface{}) int {
	var si_string string = i.(string)
	var sj_string string = j.(string)
	var si_lower string = strings.ToLower(si_string)
	var sj_lower string = strings.ToLower(sj_string)
	/*
		if si_lower == sj_lower {
			return si_string < sj_string
		}
		return si_lower < sj_lower
	*/
	if si_lower < sj_lower {
		//	//	where(fmt.Sprintf("%q < %q\n", si_string, sj_string))
		return -1
	} else if si_lower > sj_lower {
		//	where(fmt.Sprintf("%q > %q\n", si_string, sj_string))
		return +1
	} else { // si_lower == sj_lower
		if si_string < sj_string {
			//	where(fmt.Sprintf("%q < %q\n", si_string, sj_string))
			return -1
		} else if si_string > sj_string {
			//	where(fmt.Sprintf("%q > %q\n", si_string, sj_string))
			return +1
		} else {
			//	where(fmt.Sprintf("%q == %q\n", si_string, sj_string))
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

//var compareCount int

type tableSortable struct {
	table *Table
	rows  tableRows
	less  func(i tableRow, j tableRow) bool
}

func (table tableSortable) Len() int { return len(table.rows) }

func (table tableSortable) Swap(i int, j int) {
	table.rows[i], table.rows[j] = table.rows[j], table.rows[i]
}

func (table tableSortable) Less(i int, j int) bool {
	return table.less(table.rows[i], table.rows[j])
}

/*
	Sort this table by this table's currently-set sort keys.

	To see the currently-set sort keys use GetSortKeysAsTable()
*/
func (table *Table) Sort() error {

	if table == nil {
		return fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	if len(table.sortKeys) == 0 {
		return fmt.Errorf("cannot sort table that has 0 sort keys - use SetSortKeys()")
	}

	table.sortByKeys(table.sortKeys)

	return nil
}

func (table *Table) sortByKeys(sortKeys SortKeys) {
	//	where(fmt.Sprintf("Calling SortByKeys(%v)\n", sortKeys))
	sort.Sort(tableSortable{table, table.rows, func(iRow, jRow tableRow) bool {
//		compareCount++
		//where(fmt.Sprintf("len(sortKeys) = %d\n", len(sortKeys)))
		//where(fmt.Sprintf("table.sortKeys ... %v\n", table.sortKeys))
		for _, sortKey := range table.sortKeys {
			var colName string = sortKey.colName
			var sortFunc compareFunc = sortKey.sortFunc
			var iVal interface{} = iRow[colName]
			var jVal interface{} = jRow[colName]
			var compared int = sortFunc(iVal, jVal)
			//where(fmt.Sprintf("sortKey.reverse = %t\n", sortKey.reverse))
			//where(fmt.Sprintf("compared = %d ...\n", compared))
			if sortKey.reverse {
				// Reverse the sign to reverse the sort.
				// Reverse is intended to be descending, not a toggle between ascending and descending.
				compared *= -1
			}
			//where(fmt.Sprintf("... compared = %d\n", compared))
			if compared != 0 {
				//	where(fmt.Sprintf("not equal"))
				//	where(fmt.Sprintf("Less = %v\n", compared < 0))
				return compared < 0		// Less is true if compared < 0
			}
			//	where(fmt.Sprintf("*** return false\n"))
		}
		return false
	}})
}

/*
	Search this table by this table's currently-set sort keys.

	To see the currently-set sort keys use GetSortKeysAsTable()
*/
func (table *Table) Search(searchValues ...interface{}) (int, error) {

	if table == nil {
		return -1, fmt.Errorf("%s(*Table) *Table is <nil>", funcName())
	}

	if len(table.sortKeys) == 0 {
		return -1, fmt.Errorf("cannot search table that has 0 sort keys - use SetSortKeys()")
	}

	if len(searchValues) != len(table.sortKeys) {
		return -1, fmt.Errorf("%s(...) searchValues count %d != sort keys count %d",
			funcName(), len(searchValues), len(table.sortKeys))
	}

	// Check that searchValues are the right type.
	for sortIndex, sortKey := range table.sortKeys {
where(fmt.Sprintf("key[%d] = %v", sortIndex, sortKey))
where(fmt.Sprintf("val[%d] = %v", sortIndex, searchValues[sortIndex]))
	}

	table.searchByKeys(searchValues)

	return -1, nil
}

func (table *Table) searchByKeys(searchValues ...interface{}) {
where(fmt.Sprintf("Calling searchByKeys(%v)\n", searchValues))
	sort.Search(table.RowCount(), func(rowIndex int) bool {
//		compareCount++
		//where(fmt.Sprintf("len(sortKeys) = %d\n", len(sortKeys)))
		//where(fmt.Sprintf("table.sortKeys ... %v\n", table.sortKeys))
/*
		for _, sortKey := range table.sortKeys {
			var colName string = sortKey.colName
			var sortFunc compareFunc = sortKey.sortFunc
			var iVal interface{} = iRow[colName]
			var jVal interface{} = jRow[colName]
			var compared int = sortFunc(iVal, jVal)
			//where(fmt.Sprintf("sortKey.reverse = %t\n", sortKey.reverse))
			//where(fmt.Sprintf("compared = %d ...\n", compared))
			if sortKey.reverse {
				// Reverse the sign to reverse the sort.
				compared *= -1
			}
			//where(fmt.Sprintf("... compared = %d\n", compared))
			if compared != 0 {
				//	where(fmt.Sprintf("not equal"))
				//	where(fmt.Sprintf("Less = %v\n", compared < 0))
				return compared < 0
			}
			//	where(fmt.Sprintf("*** return false\n"))
		}
*/
		return false
	})
}


/*
func (tableRows tableRows) Less(i, j int) bool {
//	compareCount++
	sortFunc := compare_int
	colName := "SortOrder"
	var iInterface interface{} = tableRows[i][colName]
	var jInterface interface{} = tableRows[j][colName]
	var compared int = sortFunc(iInterface, jInterface)
	if compared != 0 {
		//	where(fmt.Sprintf("not equal"))
		//	where(fmt.Sprintf("Less = %v\n", compared < 0))
		return compared < 0
	}
	//	where(fmt.Sprintf("*** return false\n"))
	return false
}
*/

// Factory function to generate a slice of SortKeys.
func newSortKeys() SortKeys {
	return make([]sortKey, 0)
}
