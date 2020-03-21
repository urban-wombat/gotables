package gotables

import (
	"fmt"
	"testing"
	"time"
)

/*
Copyright (c) 2020 Malcolm Gorman

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

func TestGetTime(t *testing.T) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	var err error
	var table *Table

	// ISO 8601 (RFC 3339)
	var tableString string = `[TimeTable]
	t1 time.Time = 2020-03-15T14:22:30Z
	t2 time.Time = 2020-03-15T14:22:30+17:00
	t3 time.Time = 2020-03-15T14:22:30-17:00
	t4 time.Time = 2020-03-15T14:22:30.12345Z
	t5 time.Time = 2020-03-15T14:22:30.12345+17:00
	t6 time.Time = 2020-03-15T14:22:30.12345-17:00
	`
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	_ = table
}

func ExampleTable_GetTime() {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	const rowIndex int = 0 // Always row 0 for struct shaped tables

	var err error
	var table *Table
	var colIndex int

	// RFC 3339
	var tableString string = `[TimeTable]
	t0 time.Time = 2020-03-15T14:22:30Z
	t1 time.Time = 2020-03-15T14:22:30+17:00
	t2 time.Time = 2020-03-15T14:22:30-17:00
	t3 time.Time = 2020-03-15T14:22:30.12345Z
	t4 time.Time = 2020-03-15T14:22:30.12345+17:00
	t5 time.Time = 2020-03-15T14:22:30.12345-17:00
	`
	table, err = NewTableFromString(tableString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(table)

	fmt.Println("AppendCol() initialises new col with the time.Time zero val: MinTime")
	err = table.AppendCol("t6", "time.Time")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(table)

	// Here are the time.Date() function arguments:
	// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time

	fmt.Println("Set it to 2020 last day at 10pm")
	// 2020 last day at 10pm
	err = table.SetTime("t6", rowIndex, time.Date(2020, time.December, 31, 22, 0, 0, 0, time.UTC))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(table)

	fmt.Println("Add an hour")
	var t time.Time
	t, err = table.GetTime("t6", rowIndex)
	if err != nil {
		fmt.Println(err)
	}
	t = t.Add(time.Hour)
	fmt.Printf("t = %v\n", t)
	err = table.SetTime("t6", rowIndex, t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(table)

	fmt.Println("Append col t7 and set it to 2020 last day at 11:59pm and 1 nanosecond before midnight")
	// 2020 last day at 11:59pm and 1 nanosecond before midnight
	// There are 1,000,000,000 nanoseconds in a second
	err = table.AppendCol("t7", "time.Time")
	if err != nil {
		fmt.Println(err)
	}
	colIndex = 7
	err = table.SetTimeByColIndex(colIndex, rowIndex, time.Date(2020, time.December, 31, 23, 59, 59, 999999999, time.UTC))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(table)

	fmt.Println("Add a nanosecond")
	t, err = table.GetTimeByColIndex(colIndex, rowIndex)
	if err != nil {
		fmt.Println(err)
	}
	t = t.Add(time.Nanosecond)
	fmt.Printf("t = %v\n", t)
	err = table.SetTimeByColIndex(colIndex, rowIndex, t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(table)

	fmt.Println("AppendCol() and set it to gotables.MinTime")
	// MinTime is a global variable defined in gotables.go
	err = table.AppendCol("minTime", "time.Time")
	if err != nil {
		fmt.Println(err)
	}
	err = table.SetTime("minTime", 0, MinTime)

	fmt.Println("AppendCol() and set it to gotables.MaxTime")
	// MaxTime is a global variable defined in gotables.go
	err = table.AppendCol("maxTime", "time.Time")
	if err != nil {
		fmt.Println(err)
	}
	err = table.SetTime("maxTime", 0, MaxTime)
	fmt.Println(table)

	// Output:
	// [TimeTable]
	// t0 time.Time = 2020-03-15T14:22:30Z
	// t1 time.Time = 2020-03-15T14:22:30+17:00
	// t2 time.Time = 2020-03-15T14:22:30-17:00
	// t3 time.Time = 2020-03-15T14:22:30.12345Z
	// t4 time.Time = 2020-03-15T14:22:30.12345+17:00
	// t5 time.Time = 2020-03-15T14:22:30.12345-17:00
	//
	// AppendCol() initialises new col with the time.Time zero val: MinTime
	// [TimeTable]
	// t0 time.Time = 2020-03-15T14:22:30Z
	// t1 time.Time = 2020-03-15T14:22:30+17:00
	// t2 time.Time = 2020-03-15T14:22:30-17:00
	// t3 time.Time = 2020-03-15T14:22:30.12345Z
	// t4 time.Time = 2020-03-15T14:22:30.12345+17:00
	// t5 time.Time = 2020-03-15T14:22:30.12345-17:00
	// t6 time.Time = 0001-01-01T00:00:00Z
	//
	// Set it to 2020 last day at 10pm
	// [TimeTable]
	// t0 time.Time = 2020-03-15T14:22:30Z
	// t1 time.Time = 2020-03-15T14:22:30+17:00
	// t2 time.Time = 2020-03-15T14:22:30-17:00
	// t3 time.Time = 2020-03-15T14:22:30.12345Z
	// t4 time.Time = 2020-03-15T14:22:30.12345+17:00
	// t5 time.Time = 2020-03-15T14:22:30.12345-17:00
	// t6 time.Time = 2020-12-31T22:00:00Z
	//
	// Add an hour
	// t = 2020-12-31 23:00:00 +0000 UTC
	// [TimeTable]
	// t0 time.Time = 2020-03-15T14:22:30Z
	// t1 time.Time = 2020-03-15T14:22:30+17:00
	// t2 time.Time = 2020-03-15T14:22:30-17:00
	// t3 time.Time = 2020-03-15T14:22:30.12345Z
	// t4 time.Time = 2020-03-15T14:22:30.12345+17:00
	// t5 time.Time = 2020-03-15T14:22:30.12345-17:00
	// t6 time.Time = 2020-12-31T23:00:00Z
	//
	// Append col t7 and set it to 2020 last day at 11:59pm and 1 nanosecond before midnight
	// [TimeTable]
	// t0 time.Time = 2020-03-15T14:22:30Z
	// t1 time.Time = 2020-03-15T14:22:30+17:00
	// t2 time.Time = 2020-03-15T14:22:30-17:00
	// t3 time.Time = 2020-03-15T14:22:30.12345Z
	// t4 time.Time = 2020-03-15T14:22:30.12345+17:00
	// t5 time.Time = 2020-03-15T14:22:30.12345-17:00
	// t6 time.Time = 2020-12-31T23:00:00Z
	// t7 time.Time = 2020-12-31T23:59:59.999999999Z
	//
	// Add a nanosecond
	// t = 2021-01-01 00:00:00 +0000 UTC
	// [TimeTable]
	// t0 time.Time = 2020-03-15T14:22:30Z
	// t1 time.Time = 2020-03-15T14:22:30+17:00
	// t2 time.Time = 2020-03-15T14:22:30-17:00
	// t3 time.Time = 2020-03-15T14:22:30.12345Z
	// t4 time.Time = 2020-03-15T14:22:30.12345+17:00
	// t5 time.Time = 2020-03-15T14:22:30.12345-17:00
	// t6 time.Time = 2020-12-31T23:00:00Z
	// t7 time.Time = 2021-01-01T00:00:00Z
	//
	// AppendCol() and set it to gotables.MinTime
	// AppendCol() and set it to gotables.MaxTime
	// [TimeTable]
	// t0 time.Time = 2020-03-15T14:22:30Z
	// t1 time.Time = 2020-03-15T14:22:30+17:00
	// t2 time.Time = 2020-03-15T14:22:30-17:00
	// t3 time.Time = 2020-03-15T14:22:30.12345Z
	// t4 time.Time = 2020-03-15T14:22:30.12345+17:00
	// t5 time.Time = 2020-03-15T14:22:30.12345-17:00
	// t6 time.Time = 2020-12-31T23:00:00Z
	// t7 time.Time = 2021-01-01T00:00:00Z
	// minTime time.Time = 0001-01-01T00:00:00Z
	// maxTime time.Time = 292277024627-12-07T02:30:07.999999999+11:00
}
