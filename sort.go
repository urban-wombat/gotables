/*
Functions and methods for sorting Table tables.
*/
package gotable

import (
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

// Sorting functions:

func (tableRows tableRows) Len() int {
	return len(tableRows)
}

// func (table *Table) Swap(i, j int) {
func (tableRows tableRows) Swap(i, j int) {
	tableRows[i], tableRows[j] = tableRows[j], tableRows[i]
}

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

var compareCount int

type tableSortable struct {
	table *Table
	rows  tableRows
	less  func(i, j TableRow) bool
}

func (table tableSortable) Len() int { return len(table.rows) }

func (table tableSortable) Swap(i, j int) {
	table.rows[i], table.rows[j] = table.rows[j], table.rows[i]
}

func (table tableSortable) Less(i, j int) bool {
	return table.less(table.rows[i], table.rows[j])
}

func (table *Table) Sort() {
	table.SortByKeys(table.sortKeys)
}

func (table *Table) SortByKeys(sortKeys SortKeys) {
	//	where(fmt.Sprintf("Calling SortByKeys(%v)\n", sortKeys))
	sort.Sort(tableSortable{table, table.rows, func(iRow, jRow TableRow) bool {
		compareCount++
		//where(fmt.Sprintf("len(sortKeys) = %d\n", len(sortKeys)))
		//where(fmt.Sprintf("table.sortKeys ... %v\n", table.sortKeys))
		for _, sortKey := range table.sortKeys {
			var colName string = sortKey.colName
			var sortFunc compareFunc = sortKey.sortFunc
			var iInterface interface{} = iRow[colName]
			var jInterface interface{} = jRow[colName]
			var compared int = sortFunc(iInterface, jInterface)
			//where(fmt.Sprintf("sortKey.reverse = %t\n", sortKey.reverse))
			//where(fmt.Sprintf("compared = %d ...\n", compared))
			if sortKey.reverse {
				// Reverse the sign to reverse the sort.
				compared *= -1
				//temp := compared
				//temp *= -1	// Reverse the sign.
				//				// Reverse the sign.
				//				if compared < 0 {
				//					compared = +1
				//				} else if compared > 0 {
				//					compared = -1
				//				}
				//if temp != compared {
				//	err := errors.New("compared = %d and temp = %d")
				//	panic(err)
				//}
			}
			//where(fmt.Sprintf("... compared = %d\n", compared))
			if compared != 0 {
				//	where(fmt.Sprintf("not equal"))
				//	where(fmt.Sprintf("Less = %v\n", compared < 0))
				return compared < 0
			}
			//	where(fmt.Sprintf("*** return false\n"))
		}
		return false
	}})
}

/*
func (tableRows tableRows) Less(i, j int) bool {
compareCount++
////	where(fmt.Sprintf("len(table.sortKeys) = %d\n", len(table.sortKeys)))
	var sortKeyCount = len(table.sortKeys)
	for i, sortKey := range table.sortKeys {
//	//	where(fmt.Sprintf("i = %d, sortKey = %v\n", i, sortKey))
		var colName string = sortKey.colName
		var sortFunc compareFunc = sortKey.sortFunc
		var compared int = sortFunc(table, colName, i, j)
		if compared != 0 {
		//	where(fmt.Sprintf("not equal\n"))
		//	where(fmt.Sprintf("Less = %v\n", compared < 0))
			return compared < 0
		}
		if i == sortKeyCount - 1 {	// Final iteration through sort keys.
		//	where(fmt.Sprintf("Final iteration through sort keys.\n"))
		//	where(fmt.Sprintf("Less = %v\n", compared < 0))
			return compared < 0
		}
	}

//	where(fmt.Sprintf("***** End of function.\n"))
//	where(fmt.Sprintf("Less = false\n"))
	return false
}
*/

func (tableRows tableRows) Less(i, j int) bool {
	compareCount++
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
