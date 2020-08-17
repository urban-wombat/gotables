package gotables

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
	//	"gopkg.in/mgo.v2/bson"
)

/*
Copyright (c) 2018 Malcolm Gorman
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

func TestGetTableSetAsJSON(t *testing.T) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
	const verbose bool = false

	var err error
	var tableSet1 *TableSet
	var tableSet2 *TableSet

	var tableSet1String string = `
	[TypesGalore12]
    i   s      f       f32     t     b    ui    bb            uu8
    int string float64 float32 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     6.6     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     7.7     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     8.8     false 33   2     [33 34 35 36] [37 38 39]
    4   "xxxx" 5.9     9.9     true  44   3     []            []

	[AnotherTable]
	fred int = 33
	wilma int = 29
    `
	tableSet1, err = NewTableSetFromString(tableSet1String)
	if err != nil {
		t.Fatal(err)
	}

	err = tableSet1.SetName("MySet")
	if err != nil {
		t.Fatal(err)
	}

	var jsonTableSet string
	jsonTableSet, err = tableSet1.GetTableSetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	_ = jsonTableSet

	if verbose {
		fmt.Println(jsonTableSet)
		var out bytes.Buffer
		// For readability.
		err = json.Indent(&out, []byte(jsonTableSet), "", "\t")
		if err != nil {
			t.Fatal(err)
		}

		_, _ = out.WriteTo(os.Stdout)
		fmt.Println()
	}

	// Now turn it back into a TableSet.
	tableSet2, err = NewTableSetFromJSON(jsonTableSet)
	if err != nil {
		t.Fatal(err)
	}

	_, err = tableSet1.Equals(tableSet2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTableMetadataAsJSON(t *testing.T) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
	var err error
	var table *Table

	var tableString string = `
	[TypesGalore14]
    i   s      f       f32     t     b    ui    bb            uu8
    int string float64 float32 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     6.6     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     7.7     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     8.8     false 33   2     [33 34 35 36] [37 38 39]
    4   "xxxx" 5.9     9.9     true  44   3     []            []
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	var jsonString string
	jsonString, err = table.GetTableAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	_ = jsonString

	/*
		fmt.Println(jsonString)
		var out bytes.Buffer
		// For readability.
		err = json.Indent(&out, []byte(jsonString), "", "\t")
		if err != nil {
			t.Fatal(err)
		}
		_, _ = out.WriteTo(os.Stdout)
		fmt.Println()
	*/
}

func TestNewTableSetFromJSON_bothDirectionsRecursive(t *testing.T) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	const verbose = false

	var err error
	var tableSetInput *TableSet  // Input TableSet
	var tableSetOutput *TableSet // Output TableSet
	var tableInput *Table        // Input table
	var tableOutput *Table       // Output table

	var tableString string = `
		[[MyTableSet]]
		[TypesGalore16]
	    i   s      f       f32     t     b    ui    bb            uu8			table
	    int string float64 float32 bool  byte uint8 []byte        []uint8		*Table
	    0   "abc"  2.3     6.6     true  11   0     [11 12 13 14] [15 16 17]	[]
	    1   "xyz"  4.5     7.7     false 22   1     [22 23 24 25] [26 27 28]	[]
	    2   "ssss" 4.9     8.8     false 33   2     [33 34 35 36] [37 38 39]	[]
	    3   "xxxx" 5.9     9.9     true  44   3     []            []			[]
	    4   "yyyy" 6.9    10.9     false 55   4     [0]           [2]			[]

		[AnotherTable]
		i	j	k
		int	int	int
		1	3	9

		[YetAnotherTable]
		love bool = true
		hate bool = false
	    `
	tableSetInput, err = NewTableSetFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	tableInput, err = tableSetInput.GetTable("TypesGalore16")
	if err != nil {
		t.Fatal(err)
	}

	// Nest a table and see what happens.
	table4, err := NewTableFromString(
		`[Table4]
	i int = 3
	`)
	if err != nil {
		t.Fatal(err)
	}
	err = tableInput.SetTable("table", 4, table4)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where(table4)
		where(tableInput)
	}

	if verbose {
		fmt.Printf("\n\n%v\n", tableInput)
	}

	var jsonString string

	if verbose {
		where("calling GetTableAsJSON()")
	}

	jsonString, err = tableSetInput.GetTableSetAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where(jsonString)
	}

	if verbose {
		var buf bytes.Buffer
		// For readability.
		err = json.Indent(&buf, []byte(jsonString), "", "\t")
		if err != nil {
			t.Fatal(err)
		}
		_, _ = buf.WriteTo(os.Stdout)
	}

	// where("***CALLING** NewTableSetFromJSON() ...")
	tableSetOutput, err = NewTableSetFromJSON(jsonString)
	if err != nil {
		t.Fatal(err)
	}

	// where("***CALLING** NewTableFromJSON() ...")
	tableOutput, err = NewTableFromJSONByTableName(jsonString, "YetAnotherTable")
	if err != nil {
		t.Fatal(err)
	}

	_, err = tableSetInput.Equals(tableSetOutput)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		fmt.Printf("\n\n%v\n", tableOutput)
	}
}

func TestNewTableFromJSON_bothDirections(t *testing.T) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	const verbose = false

	var err error
	var tableInput *Table  // Input table
	var tableOutput *Table // Output table

	var tableString string = `
		[TypesGalore16]
	    i   s      f       f32     t     b    ui    bb            uu8			table
	    int string float64 float32 bool  byte uint8 []byte        []uint8		*Table
	    0   "abc"  2.3     6.6     true  11   0     [11 12 13 14] [15 16 17]	[]
	    1   "xyz"  4.5     7.7     false 22   1     [22 23 24 25] [26 27 28]	[]
	    2   "ssss" 4.9     8.8     false 33   2     [33 34 35 36] [37 38 39]	[]
	    3   "xxxx" 5.9     9.9     true  44   3     []            []			[]
	    4   "yyyy" 6.9    10.9     false 55   4     [0]           [2]			[]
	    `
	tableInput, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	// Nest a table and see what happens.
	table4, err := NewTableFromString(
		`[Table4]
	i int = 3
	`)
	if err != nil {
		t.Fatal(err)
	}
	err = tableInput.SetTable("table", 4, table4)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where(table4)
		where(tableInput)
	}

	if verbose {
		fmt.Printf("\n\n%v\n", tableInput)
	}

	var jsonString string

	if verbose {
		where("calling GetTableAsJSON()")
	}
	jsonString, err = tableInput.GetTableAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where(jsonString)
	}

	if verbose {
		var buf bytes.Buffer
		// For readability.
		err = json.Indent(&buf, []byte(jsonString), "", "\t")
		if err != nil {
			t.Fatal(err)
		}
		_, _ = buf.WriteTo(os.Stdout)
	}

	// where("***CALLING** NewTableFromJSON() ...")
	tableOutput, err = NewTableFromJSON(jsonString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = tableInput.Equals(tableOutput)
	if err != nil {
		t.Fatal(err)
	}

	// Does table.Equals() check row order?
	for rowIndex := 0; rowIndex < tableOutput.RowCount(); rowIndex++ {
		i, err := tableOutput.GetInt("i", rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if i != rowIndex {
			t.Fatalf("expecting decoded table rows in order, but found row %d at rowIndex %d", i, rowIndex)
		}
	}

	if verbose {
		fmt.Printf("\n\n%v\n", tableOutput)
	}
}

func TestNewTableFromJSONZeroRows(t *testing.T) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	const verbose = false

	var err error
	var tableInput *Table  // Input table
	var tableOutput *Table // Output table

	var tableString string = `
		[TypesGalore16]
	    i   s      f       f32     t     b    ui    bb            uu8			table
	    int string float64 float32 bool  byte uint8 []byte        []uint8		*Table
	    `
	tableInput, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		fmt.Printf("\n\n%v\n", tableInput)
	}

	var jsonString string

	if verbose {
		where("calling GetTableAsJSON()")
	}
	jsonString, err = tableInput.GetTableAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where(jsonString)
	}

	if verbose {
		var buf bytes.Buffer
		// For readability.
		err = json.Indent(&buf, []byte(jsonString), "", "\t")
		if err != nil {
			t.Fatal(err)
		}
		_, _ = buf.WriteTo(os.Stdout)
	}

	// where("***CALLING** NewTableFromJSON() ...")
	tableOutput, err = NewTableFromJSON(jsonString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = tableInput.Equals(tableOutput)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		fmt.Printf("\n\n%v\n", tableOutput)
	}
}

func TestNewTableFromJSONZeroCols(t *testing.T) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	const verbose = false

	var err error
	var tableInput *Table  // Input table
	var tableOutput *Table // Output table

	var tableString string = `
		[TypesGalore16]
	    `
	tableInput, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		fmt.Printf("\n\n%v\n", tableInput)
	}

	var jsonString string

	if verbose {
		where("calling GetTableAsJSON()")
	}
	jsonString, err = tableInput.GetTableAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where(jsonString)
	}

	if verbose {
		var buf bytes.Buffer
		// For readability.
		err = json.Indent(&buf, []byte(jsonString), "", "\t")
		if err != nil {
			t.Fatal(err)
		}
		_, _ = buf.WriteTo(os.Stdout)
	}

	// where("***CALLING** NewTableFromJSON() ...")
	tableOutput, err = NewTableFromJSON(jsonString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = tableInput.Equals(tableOutput)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		fmt.Printf("\n\n%v\n", tableOutput)
	}
}

func TestNewTableSetFromJSON(t *testing.T) {
	var verbose bool = false
	var err error
	var tableSet1 *TableSet // Input tableSet
	var tableSet2 *TableSet // Output tableSet

	var tableSetString string = `
		[[LetsNameIt]]
		[TypesGalore17]
	    i   s      f       f32     t     b    ui    bb            uu8
	    int string float64 float32 bool  byte uint8 []byte        []uint8
	    1   "abc"  2.3     6.6     true  11   0     [11 12 13 14] [15 16 17]
	    2   "xyz"  4.5     7.7     false 22   1     [22 23 24 25] [26 27 28]
	    3   "ssss" 4.9     8.8     false 33   2     [33 34 35 36] [37 38 39]
	    4   "xxxx" 5.9     9.9     true  44   3     []            []
	    5   "yyyy" 6.9    10.9     false 55   4     [0]           [2]

	[AnotherTable]
	fred int = 33
	wilma int = 29
    `
	tableSet1, err = NewTableSetFromString(tableSetString)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where(tableSet1)
	}

	var jsonString string
	jsonString, err = tableSet1.GetTableSetAsJSONIndent()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where(jsonString)
	}

	tableSet2, err = NewTableSetFromJSON(jsonString)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where(tableSet2)
	}

	_, err = tableSet1.Equals(tableSet2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAllJSON(t *testing.T) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
	const verbose = false

	var err error

	tableString :=
		`[MyTable]
	x int = 1
	y int = 2
	z int = 3
	`
	table1, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where("\n" + table1.String())
		fmt.Println()
	}

	var jsonString string
	var buf bytes.Buffer

	jsonString, err = table1.GetTableAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		// For readability.
		err = json.Indent(&buf, []byte(jsonString), "", "\t")
		if err != nil {
			t.Fatal(err)
		}
		_, _ = buf.WriteTo(os.Stdout)

		fmt.Println()
		fmt.Println()
	}

	table2, err := NewTableFromJSON(jsonString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = table2.Equals(table1)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where("\n" + table2.String())
		fmt.Println()
	}

	tableSetString :=
		`[[MyTableSet]]
	[MyTable]
	x int = 1
	y int = 2
	z int = 3

	[YourTable]
	a float32 = 1
	b float32 = 2
	c float32 = 3
	`
	tableSet1, err := NewTableSetFromString(tableSetString)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where("\n" + tableSet1.String())
		fmt.Println()
	}

	var jsonTableSet string

	jsonTableSet, err = tableSet1.GetTableSetAsJSONIndent()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		fmt.Println(jsonTableSet)
		fmt.Println()
	}

	tableSet2, err := NewTableSetFromJSON(jsonTableSet)
	if err != nil {
		t.Fatal(err)
	}

	_, err = tableSet2.Equals(tableSet1)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where("\n" + tableSet2.String())
		fmt.Println()
	}
}

func ExampleNewTableFromJSON() {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
	var err error

	tableString :=
		`[MyTable]
	i    u    f       t
	int  uint float32 time.Time
	1    2    3.3     2020-03-15T14:22:30Z
	4    5    6.6     2020-03-15T14:22:30.123456789+17:00
	`
	table1, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(table1)

	var jsonString string
	var buf bytes.Buffer

	jsonString, err = table1.GetTableAsJSON()
	if err != nil {
		fmt.Println(err)
	}

	// For readability.
	err = json.Indent(&buf, []byte(jsonString), "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	_, _ = buf.WriteTo(os.Stdout)
	fmt.Println()
	fmt.Println()

	table2, err := NewTableFromJSON(jsonString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(table2)

	equals, err := table2.Equals(table1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("table2.Equals(table1) == %t\n", equals)

	// Output:
	// [MyTable]
	//   i    u       f t
	// int uint float32 time.Time
	//   1    2     3.3 2020-03-15T14:22:30Z
	//   4    5     6.6 2020-03-15T14:22:30.123456789+17:00
	//
	// {
	//   "tableSetName": "",
	//   "tables": [
	//     {
	//       "tableName": "MyTable",
	//       "metadata": [
	//         {
	//           "i": "int"
	//         },
	//         {
	//           "u": "uint"
	//         },
	//         {
	//           "f": "float32"
	//         },
	//         {
	//           "t": "time.Time"
	//         }
	//       ],
	//       "data": [
	//         [
	//           {
	//             "i": 1
	//           },
	//           {
	//             "u": 2
	//           },
	//           {
	//             "f": 3.3
	//           },
	//           {
	//             "t": "2020-03-15T14:22:30Z"
	//           }
	//         ],
	//         [
	//           {
	//             "i": 4
	//           },
	//           {
	//             "u": 5
	//           },
	//           {
	//             "f": 6.6
	//           },
	//           {
	//             "t": "2020-03-15T14:22:30.123456789+17:00"
	//           }
	//         ]
	//       ]
	//     }
	//   ]
	// }
	//
	// [MyTable]
	//   i    u       f t
	// int uint float32 time.Time
	//   1    2     3.3 2020-03-15T14:22:30Z
	//   4    5     6.6 2020-03-15T14:22:30.123456789+17:00
	//
	// table2.Equals(table1) == true
}

/*
	It is permitted to have zero rows.
*/
func ExampleNewTableFromJSON_zeroRows() {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
	var err error

	tableString :=
		`[MyTable]
	i    u    f
	int  uint float32
	`
	table1, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table1)

	var jsonString string
	var buf bytes.Buffer

	jsonString, err = table1.GetTableAsJSON()
	if err != nil {
		fmt.Println(err)
	}

	// For readability.
	err = json.Indent(&buf, []byte(jsonString), "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	_, _ = buf.WriteTo(os.Stdout)

	fmt.Println()

	table2, err := NewTableFromJSON(jsonString)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(table2)

	equals, err := table2.Equals(table1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("table2.Equals(table1) == %t\n", equals)

	// Output:
	// [MyTable]
	//   i    u       f
	// int uint float32
	//
	// {
	//   "tableSetName": "",
	//   "tables": [
	//     {
	//       "tableName": "MyTable",
	//       "metadata": [
	//         {
	//           "i": "int"
	//         },
	//         {
	//           "u": "uint"
	//         },
	//         {
	//           "f": "float32"
	//         }
	//       ],
	//       "data": []
	//     }
	//   ]
	// }
	// [MyTable]
	//   i    u       f
	// int uint float32
	//
	// table2.Equals(table1) == true
}

/*
	It is permitted to have zero cols.
*/
func ExampleNewTableFromJSON_zeroCols() {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
	var err error

	tableString :=
		`[MyTable]
	`
	table1, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table1)

	var jsonString string
	jsonString, err = table1.GetTableAsJSONIndent()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(jsonString)
	fmt.Println()

	table2, err := NewTableFromJSON(jsonString)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(table2)

	equals, err := table2.Equals(table1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("table2.Equals(table1) == %t\n", equals)

	// Output:
	// [MyTable]
	//
	// {
	// 	"tableSetName": "",
	// 	"tables": [
	// 		{
	// 			"tableName": "MyTable",
	// 			"metadata": [],
	// 			"data": []
	// 		}
	// 	]
	// }
	//
	// [MyTable]
	//
	// table2.Equals(table1) == true
}

func ExampleNewTableSetFromJSON() {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
	var err error

	tableSetString := `
	[[TwoTablesComingUp]]

	[MyTable]
	x int = 1
	y int = 2
	z int = 3

	[YourTable]
	a float32 = 1
	b float32 = 2
	c float32 = 3
	`
	tableSet1, err := NewTableSetFromString(tableSetString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(tableSet1)

	var jsonString string
	jsonString, err = tableSet1.GetTableSetAsJSONIndent()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(jsonString)
	fmt.Println()

	tableSet2, err := NewTableSetFromJSON(jsonString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(tableSet2)

	equals, err := tableSet2.Equals(tableSet1)

	fmt.Printf("table2.Equals(table1) == %t\n", equals)

	// Output:
	// [[TwoTablesComingUp]]
	//
	// [MyTable]
	// x int = 1
	// y int = 2
	// z int = 3
	//
	// [YourTable]
	// a float32 = 1
	// b float32 = 2
	// c float32 = 3
	//
	// {
	// 	"tableSetName": "TwoTablesComingUp",
	// 	"tables": [
	// 		{
	// 			"tableName": "MyTable",
	// 			"isStructShape": true,
	// 			"metadata": [
	// 				{
	// 					"x": "int"
	// 				},
	// 				{
	// 					"y": "int"
	// 				},
	// 				{
	// 					"z": "int"
	// 				}
	// 			],
	// 			"data": [
	// 				[
	// 					{
	// 						"x": 1
	// 					},
	// 					{
	// 						"y": 2
	// 					},
	// 					{
	// 						"z": 3
	// 					}
	// 				]
	// 			]
	// 		},
	// 		{
	// 			"tableName": "YourTable",
	// 			"isStructShape": true,
	// 			"metadata": [
	// 				{
	// 					"a": "float32"
	// 				},
	// 				{
	// 					"b": "float32"
	// 				},
	// 				{
	// 					"c": "float32"
	// 				}
	// 			],
	// 			"data": [
	// 				[
	// 					{
	// 						"a": 1
	// 					},
	// 					{
	// 						"b": 2
	// 					},
	// 					{
	// 						"c": 3
	// 					}
	// 				]
	// 			]
	// 		}
	// 	]
	// }
	//
	// [[TwoTablesComingUp]]
	//
	// [MyTable]
	// x int = 1
	// y int = 2
	// z int = 3
	//
	// [YourTable]
	// a float32 = 1
	// b float32 = 2
	// c float32 = 3
	//
	// table2.Equals(table1) == true
}

/*
	It is permitted to have zero rows.
*/
func ExampleNewTableSetFromJSON_zeroRows() {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
	var err error

	tableSetString :=
		`[[ZeroRowsTableSet]]

	[MyTable]
	x	y	z
	int	int	int

	[YourTable]
	a		b		c
	float32	float32	float32
	`
	tableSet1, err := NewTableSetFromString(tableSetString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(tableSet1)

	var jsonString string
	jsonString, err = tableSet1.GetTableSetAsJSONIndent()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(jsonString)
	fmt.Println()

	tableSet2, err := NewTableSetFromJSON(jsonString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(tableSet2)

	equals, err := tableSet2.Equals(tableSet1)

	fmt.Printf("tableSet2.Equals(tableSet1) == %t\n", equals)

	// Output:
	// [[ZeroRowsTableSet]]
	//
	// [MyTable]
	//   x   y   z
	// int int int
	//
	// [YourTable]
	//       a       b       c
	// float32 float32 float32
	//
	// {
	// 	"tableSetName": "ZeroRowsTableSet",
	// 	"tables": [
	// 		{
	// 			"tableName": "MyTable",
	// 			"metadata": [
	// 				{
	// 					"x": "int"
	// 				},
	// 				{
	// 					"y": "int"
	// 				},
	// 				{
	// 					"z": "int"
	// 				}
	// 			],
	// 			"data": []
	// 		},
	// 		{
	// 			"tableName": "YourTable",
	// 			"metadata": [
	// 				{
	// 					"a": "float32"
	// 				},
	// 				{
	// 					"b": "float32"
	// 				},
	// 				{
	// 					"c": "float32"
	// 				}
	// 			],
	// 			"data": []
	// 		}
	// 	]
	// }
	//
	// [[ZeroRowsTableSet]]
	//
	// [MyTable]
	//   x   y   z
	// int int int
	//
	// [YourTable]
	//       a       b       c
	// float32 float32 float32
	//
	// tableSet2.Equals(tableSet1) == true
}

// TODO: Uncomment after 16.08.2020
/*
func ExampleTable_GetTableAsJSON_nestedTablesCircularReference() {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
	var err error
	var table *Table

	//	A table with value [] will result in a NilTable with no table name.
	//	To make the table usable, give it a table name.

	var tableString string
	tableString = `
	[SameTableReference]
    left	i   s      right
    *Table	int string *Table
    []		42  "abc"  [] 
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("This should fail: We are assigning the same table as the parent.")
	err = table.SetTable("right", 0, table) // table already exists (at the top level)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s", table)
	_, err = table.GetTableAsJSON()
	if err != nil {
		// Error prints here.
		fmt.Println(err)
	}
	fmt.Println()

// TODO: Restore Uncomment after 16.08.2020
//	hasCircularReference, err := table.HasCircularReference()
//	if hasCircularReference {
//where(err)
//		os.Exit(1)
//	}
//
//	isValidTableNesting, err := table.IsValidTableNesting()
//	if !isValidTableNesting {
//		panic(err)
//	}

	fmt.Println("Now try again with a COPY of the same table, which will have a new reference.")
	var jsonString string
	var tableCopy *Table
	tableCopy, err = table.Copy(true)
	if err != nil {
		// No error to print here.
		fmt.Println(err)
where(err)
	}
	err = tableCopy.SetName("TableCopy")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("By the way, don't try to set table 'right' to <nil>. Not allowed. Must use an actual *Table reference.")
	err = tableCopy.SetTable("right", 0, nil) // Not allowed. Must use an actual *Table reference.
	if err != nil {
		fmt.Println(err)
where(err)
	}
	err = tableCopy.SetTable("right", 0, NewNilTable()) // Otherwise this is another circular reference.
	if err != nil {
		fmt.Println(err)
where(err)
	}
	fmt.Printf("%s", tableCopy)
	fmt.Println()

	fmt.Println("This should succeed: We are assigning a DIFFERENT table (same contents doesn't matter).")
	table.SetTableMustSet("right", 0, tableCopy) // Different table reference.
	fmt.Printf("%s", table)
	jsonString, err = table.GetTableAsJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	fmt.Println("Print as is:")
	jsonString, err = table.GetTableAsJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonString)
	fmt.Println()

	fmt.Println("Print indented for readability:")
	jsonString, err = table.GetTableAsJSONIndent()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonString)

	fmt.Println()

	fmt.Println("(1) This should (NOT?) fail: We are assigning the same table to multiple cells.")
	table.SetTableMustSet("left", 0, tableCopy) // Different table reference.
	fmt.Printf("%s", table)
	jsonString, err = table.GetTableAsJSON()
	if err != nil {
		// Prints error.
		fmt.Println(err)
	}
	fmt.Println()

	fmt.Println("(2) This should fail: We are assigning the same table to multiple cells.")
	valid, err := table.IsValidTableNesting()
	fmt.Printf("table.IsValidTableNesting(): valid = %t\n", valid)
	if err != nil {
		// Prints error.
		fmt.Println(err)
	}
	fmt.Println()

	fmt.Println("(3) This should fail: We are assigning the same table to multiple cells.")
	valid, err = table.IsValidTableNesting()
	fmt.Printf("table.IsValidTableNesting(): valid = %t\n", valid)
	// where(fmt.Sprintf("table.IsValidTableNesting(): valid = %t\n", valid))
	if err != nil {
		// Prints error.
		fmt.Println(err)

		// Now check to see if there is a wrapped CircRefError inside err.
		// Method 1: Just see if it is there.
		var has bool = HasCircRefError(err)
		if has {
			var circError *CircRefError = GetCircRefError(err)
			// Now Get the error.
			fmt.Println("Yes, there is a wrapped CircRefError inside err:")
			fmt.Printf("circError.Error(): %s\n", circError.Error())
			fmt.Printf("circError.RootTable(): %s\n", circError.RootTable())
			fmt.Printf("circError.CircTable(): %s\n", circError.CircTable())
		}
		// Method 2: Just try to get the error and test if it is there.
		var circError *CircRefError = GetCircRefError(err)
		if circError != nil {
			fmt.Println("Yes, there is a wrapped CircRefError inside err:")
			fmt.Printf("circError.Error(): %s\n", circError.Error())
			fmt.Printf("circError.RootTable(): %s\n", circError.RootTable())
			fmt.Printf("circError.CircTable(): %s\n", circError.CircTable())
		}
	}

	// Output:
	// This should fail: We are assigning the same table as the parent.
	// [SameTableReference]
	// left     i s      right
	// *Table int string *Table
	// []      42 "abc"  [SameTableReference]
	// getTableAsJSON_recursive(): circular reference in table [SameTableReference]: a reference to table [SameTableReference] already exists
	//
	// Now try again with a COPY of the same table, which will have a new reference.
	// By the way, don't try to set table 'right' to <nil>. Not allowed. Must use an actual *Table reference.
	// SetTable(right, 0, val): table [TableCopy] col right expecting val of type *Table, not: <nil> [use NewNilTable() instead of <nil>]
	// [TableCopy]
	// left     i s      right
	// *Table int string *Table
	// []      42 "abc"  []
	//
	// This should succeed: We are assigning a DIFFERENT table (same contents doesn't matter).
	// [SameTableReference]
	// left     i s      right
	// *Table int string *Table
	// []      42 "abc"  [TableCopy]
	//
	// Print as is:
	// {"tableSetName":"","tables":[{"tableName":"SameTableReference","metadata":[{"left":"*Table"},{"i":"int"},{"s":"string"},{"right":"*Table"}],"data":[[{"left":null},{"i":42},{"s":"abc"},{"right":{"tableName":"TableCopy","metadata":[{"left":"*Table"},{"i":"int"},{"s":"string"},{"right":"*Table"}],"data":[[{"left":null},{"i":42},{"s":"abc"},{"right":null}]]}}]]}]}
	//
	// Print indented for readability:
	// {
	// 	"tableSetName": "",
	// 	"tables": [
	// 		{
	// 			"tableName": "SameTableReference",
	// 			"metadata": [
	// 				{
	// 					"left": "*Table"
	// 				},
	// 				{
	// 					"i": "int"
	// 				},
	// 				{
	// 					"s": "string"
	// 				},
	// 				{
	// 					"right": "*Table"
	// 				}
	// 			],
	// 			"data": [
	// 				[
	// 					{
	// 						"left": null
	// 					},
	// 					{
	// 						"i": 42
	// 					},
	// 					{
	// 						"s": "abc"
	// 					},
	// 					{
	// 						"right": {
	// 							"tableName": "TableCopy",
	// 							"metadata": [
	// 								{
	// 									"left": "*Table"
	// 								},
	// 								{
	// 									"i": "int"
	// 								},
	// 								{
	// 									"s": "string"
	// 								},
	// 								{
	// 									"right": "*Table"
	// 								}
	// 							],
	// 							"data": [
	// 								[
	// 									{
	// 										"left": null
	// 									},
	// 									{
	// 										"i": 42
	// 									},
	// 									{
	// 										"s": "abc"
	// 									},
	// 									{
	// 										"right": null
	// 									}
	// 								]
	// 							]
	// 						}
	// 					}
	// 				]
	// 			]
	// 		}
	// 	]
	// }
	//
	// (1) This should (NOT?) fail: We are assigning the same table to multiple cells.
	// [SameTableReference]
	// left          i s      right
	// *Table      int string *Table
	// [TableCopy]  42 "abc"  [TableCopy]
	// getTableAsJSON_recursive(): circular reference in table [SameTableReference]: a reference to table [TableCopy] already exists
	//
	// (2) This should fail: We are assigning the same table to multiple cells.
	// table.IsValidTableNesting(): valid = false
	// IsValidTableNesting(): CircRefError: circular reference in table [SameTableReference]: a reference to table [TableCopy] already exists
	//
	// (3) This should fail: We are assigning the same table to multiple cells.
	// table.IsValidTableNesting(): valid = false
	// IsValidTableNesting(): CircRefError: circular reference in table [SameTableReference]: a reference to table [TableCopy] already exists
	// Yes, there is a wrapped CircRefError inside err:
	// circError.Error(): CircRefError: circular reference in table [SameTableReference]: a reference to table [TableCopy] already exists
	// circError.RootTable(): [SameTableReference]
	// left          i s      right
	// *Table      int string *Table
	// [TableCopy]  42 "abc"  [TableCopy]
	//
	// circError.CircTable(): [TableCopy]
	// left     i s      right
	// *Table int string *Table
	// []      42 "abc"  []
	//
	// Yes, there is a wrapped CircRefError inside err:
	// circError.Error(): CircRefError: circular reference in table [SameTableReference]: a reference to table [TableCopy] already exists
	// circError.RootTable(): [SameTableReference]
	// left          i s      right
	// *Table      int string *Table
	// [TableCopy]  42 "abc"  [TableCopy]
	//
	// circError.CircTable(): [TableCopy]
	// left     i s      right
	// *Table int string *Table
	// []      42 "abc"  []
}
*/

func ExampleTable_GetTableAsJSON_nestedTables() {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
	var err error
	var table1 *Table

	var tableString string
	tableString = `

	[[MyTableSetName]]

	[TypesGalore22]
    i   s      right
    int string *Table
    0   "abc"  []
    1   "xyz"  []
    2   "ssss" []
    3   "xxxx" []
    4   "yyyy" []
    `
	table1, err = NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	// Now create and set some table cell tables.
	right0 := `
	[right0]
	i int = 32`

	right1 := `
	[right1]
	s string = "thirty-two"`

	right2 := `
	[right2]
	x	y	z
	int	int	int
	1	2	3
	4	5	6
	7	8	9`

	right3 := `
	[right3]
	f float32 = 88.8`

	right4 := `
	[right4]
	t1 *Table = []`

	table1.SetTableMustSet("right", 0, NewTableFromStringMustMake(right0))
	table1.SetTableMustSet("right", 1, NewTableFromStringMustMake(right1))
	table1.SetTableMustSet("right", 2, NewTableFromStringMustMake(right2))
	table1.SetTableMustSet("right", 3, NewTableFromStringMustMake(right3))
	table1.SetTableMustSet("right", 4, NewTableFromStringMustMake(right4))

	var jsonString string
	//where("***CALLING** NewTableFromJSON() ...")
	jsonString, err = table1.GetTableAsJSON()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Print as is:")
	fmt.Println()
	fmt.Println(jsonString)
	fmt.Println()

	fmt.Println("Print indented for readability:")
	fmt.Println()
	var out bytes.Buffer
	err = json.Indent(&out, []byte(jsonString), "", "\t")
	if err != nil {
		log.Println(err)
	}
	_, _ = out.WriteTo(os.Stdout)

	// Now let's get it back from JSON into *Table
	table2, err := NewTableFromJSON(jsonString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println(table2)

	// Output:
	// Print as is:
	//
	// {"tableSetName":"","tables":[{"tableName":"TypesGalore22","metadata":[{"i":"int"},{"s":"string"},{"right":"*Table"}],"data":[[{"i":0},{"s":"abc"},{"right":{"tableName":"right0","isStructShape":true,"metadata":[{"i":"int"}],"data":[[{"i":32}]]}}],[{"i":1},{"s":"xyz"},{"right":{"tableName":"right1","isStructShape":true,"metadata":[{"s":"string"}],"data":[[{"s":"thirty-two"}]]}}],[{"i":2},{"s":"ssss"},{"right":{"tableName":"right2","metadata":[{"x":"int"},{"y":"int"},{"z":"int"}],"data":[[{"x":1},{"y":2},{"z":3}],[{"x":4},{"y":5},{"z":6}],[{"x":7},{"y":8},{"z":9}]]}}],[{"i":3},{"s":"xxxx"},{"right":{"tableName":"right3","isStructShape":true,"metadata":[{"f":"float32"}],"data":[[{"f":88.8}]]}}],[{"i":4},{"s":"yyyy"},{"right":{"tableName":"right4","isStructShape":true,"metadata":[{"t1":"*Table"}],"data":[[{"t1":null}]]}}]]}]}
	//
	// Print indented for readability:
	//
	// {
	// 	"tableSetName": "",
	// 	"tables": [
	// 		{
	// 			"tableName": "TypesGalore22",
	// 			"metadata": [
	// 				{
	// 					"i": "int"
	// 				},
	// 				{
	// 					"s": "string"
	// 				},
	// 				{
	// 					"right": "*Table"
	// 				}
	// 			],
	// 			"data": [
	// 				[
	// 					{
	// 						"i": 0
	// 					},
	// 					{
	// 						"s": "abc"
	// 					},
	// 					{
	// 						"right": {
	// 							"tableName": "right0",
	// 							"isStructShape": true,
	// 							"metadata": [
	// 								{
	// 									"i": "int"
	// 								}
	// 							],
	// 							"data": [
	// 								[
	// 									{
	// 										"i": 32
	// 									}
	// 								]
	// 							]
	// 						}
	// 					}
	// 				],
	// 				[
	// 					{
	// 						"i": 1
	// 					},
	// 					{
	// 						"s": "xyz"
	// 					},
	// 					{
	// 						"right": {
	// 							"tableName": "right1",
	// 							"isStructShape": true,
	// 							"metadata": [
	// 								{
	// 									"s": "string"
	// 								}
	// 							],
	// 							"data": [
	// 								[
	// 									{
	// 										"s": "thirty-two"
	// 									}
	// 								]
	// 							]
	// 						}
	// 					}
	// 				],
	// 				[
	// 					{
	// 						"i": 2
	// 					},
	// 					{
	// 						"s": "ssss"
	// 					},
	// 					{
	// 						"right": {
	// 							"tableName": "right2",
	// 							"metadata": [
	// 								{
	// 									"x": "int"
	// 								},
	// 								{
	// 									"y": "int"
	// 								},
	// 								{
	// 									"z": "int"
	// 								}
	// 							],
	// 							"data": [
	// 								[
	// 									{
	// 										"x": 1
	// 									},
	// 									{
	// 										"y": 2
	// 									},
	// 									{
	// 										"z": 3
	// 									}
	// 								],
	// 								[
	// 									{
	// 										"x": 4
	// 									},
	// 									{
	// 										"y": 5
	// 									},
	// 									{
	// 										"z": 6
	// 									}
	// 								],
	// 								[
	// 									{
	// 										"x": 7
	// 									},
	// 									{
	// 										"y": 8
	// 									},
	// 									{
	// 										"z": 9
	// 									}
	// 								]
	// 							]
	// 						}
	// 					}
	// 				],
	// 				[
	// 					{
	// 						"i": 3
	// 					},
	// 					{
	// 						"s": "xxxx"
	// 					},
	// 					{
	// 						"right": {
	// 							"tableName": "right3",
	// 							"isStructShape": true,
	// 							"metadata": [
	// 								{
	// 									"f": "float32"
	// 								}
	// 							],
	// 							"data": [
	// 								[
	// 									{
	// 										"f": 88.8
	// 									}
	// 								]
	// 							]
	// 						}
	// 					}
	// 				],
	// 				[
	// 					{
	// 						"i": 4
	// 					},
	// 					{
	// 						"s": "yyyy"
	// 					},
	// 					{
	// 						"right": {
	// 							"tableName": "right4",
	// 							"isStructShape": true,
	// 							"metadata": [
	// 								{
	// 									"t1": "*Table"
	// 								}
	// 							],
	// 							"data": [
	// 								[
	// 									{
	// 										"t1": null
	// 									}
	// 								]
	// 							]
	// 						}
	// 					}
	// 				]
	// 			]
	// 		}
	// 	]
	// }
	//
	// [TypesGalore22]
	//   i s      right
	// int string *Table
	//   0 "abc"  [right0]
	//   1 "xyz"  [right1]
	//   2 "ssss" [right2]
	//   3 "xxxx" [right3]
	//   4 "yyyy" [right4]
}

func BenchmarkGetTableSetAsJSON(b *testing.B) {
	// Set up for benchmark.
	tableSetString :=
		`[[MySet]]
	[sable_fur]
    i   s       f           b
    int string  float64     bool
    1   "abc"   2.3         true
    2   "xyz"   4.5         false
    3   "ssss"  4.9         false

	[Struct_With_Data]
	Fred int = 42
	Wilma int = 39
	Pebbles int = 2

	[Empty_Struct]
	Fred int

	[Empty_Table]
	Fred
	int
	`
	tableSet, err := NewTableSetFromString(tableSetString)
	if err != nil {
		b.Error(err)
	}

	var jsonTableSet string
	for i := 0; i < b.N; i++ {
		jsonTableSet, err = tableSet.GetTableSetAsJSON()
		if err != nil {
			b.Error(err)
		}
	}
	_ = jsonTableSet
}

func BenchmarkNewTableSetFromJSON(b *testing.B) {
	// Set up for benchmark.
	tableSetString :=
		`[[MyTableSet]]
	[sable_fur]
    i   s       f           b
    int string  float64     bool
    1   "abc"   2.3         true
    2   "xyz"   4.5         false
    3   "ssss"  4.9         false

	[Struct_With_Data]
	Fred int = 42
	Wilma int = 39
	Pebbles int = 2

	[Empty_Struct]
	Fred int

	[Empty_Table]
	Fred
	int
	`
	tableSet, err := NewTableSetFromString(tableSetString)
	if err != nil {
		b.Error(err)
	}

	// Set up for benchmark.
	jsonString, err := tableSet.GetTableSetAsJSON()
	if err != nil {
		b.Error(err)
	}

	var tableSet2 *TableSet
	for i := 0; i < b.N; i++ {
		_, err := NewTableSetFromJSON(jsonString)
		if err != nil {
			b.Error(err)
		}
	}
	_ = tableSet2
}

/*
func TestTable_GetTableAsBinary_nestedTable(t *testing.T) {
	//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
	var err error
	var table1 *Table
	var tableString string
	tableString = `
	[TypesGalore22]
    i   s      right
    int string *Table
    0   "abc"  []
    1   "xyz"  []
    2   "ssss" []
    3   "xxxx" []
    4   "yyyy" []
    `
	table1, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}
	// Now create and set some table cell tables.
	right0 := `
	[right0]
	i int = 32`
	right1 := `
	[right1]
	s string = "thirty-two"`
	right2 := `
	[right2]
	x	y	z
	int	int	int
	1	2	3
	4	5	6
	7	8	9`
	right3 := `
	[right3]
	f float32 = 88.8`
	right4 := `
	[right4]
	t1 *Table = []`
	table1.SetTableMustSet("right", 0, NewTableFromStringMustMake(right0))
	table1.SetTableMustSet("right", 1, NewTableFromStringMustMake(right1))
	table1.SetTableMustSet("right", 2, NewTableFromStringMustMake(right2))
	table1.SetTableMustSet("right", 3, NewTableFromStringMustMake(right3))
	table1.SetTableMustSet("right", 4, NewTableFromStringMustMake(right4))
	fmt.Printf("table1:\n%s\n", table1)
	encoded, err := bson.Marshal(table1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("encoded type: %T\n", encoded)
	fmt.Printf("len(encoded) = %d\n", len(encoded))
	// Now let's get it back from JSON into *Table
	var table2 *Table
	err = bson.Unmarshal(encoded, &table2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("table2:\n%s\n", table2)
}
*/

func TestTable_isCircularReference(t *testing.T) {

	var err error

	t0string := `
	[T0]
	t0 *Table = [t0]`
	table0, err := NewTableFromString(t0string)
	if err != nil {
		t.Fatal(err)
	}
	t1string := `
	[T1]
	t1 *Table = [t1]`
	table1, err := NewTableFromString(t1string)
	if err != nil {
		t.Fatal(err)
	}
	t2string := `
	[T2]
	t2 *Table = [t2]`
	table2, err := NewTableFromString(t2string)
	if err != nil {
		t.Fatal(err)
	}
	t3string := `
	[T3]
	t3 *Table = [t3]`
	table3, err := NewTableFromString(t3string)
	if err != nil {
		t.Fatal(err)
	}

	// Assign to table-cell.
	err = table0.SetTableByColIndex(0, 0, table1)
	if err != nil {
		t.Fatal(err)
	}
	err = table1.SetTableByColIndex(0, 0, table2)
	if err != nil {
		t.Fatal(err)
	}
	err = table2.SetTableByColIndex(0, 0, table3)
	if err != nil {
		t.Fatal(err)
	}
	err = table3.SetTableByColIndex(0, 0, table0)
	if err != nil {
		t.Fatal(err)
	}

	/*
	   where(fmt.Sprintf("%p", table0))
	   where(fmt.Sprintf("\n%s", table0))

	   where(fmt.Sprintf("%p", table1))
	   where(fmt.Sprintf("\n%s", table1))

	   where(fmt.Sprintf("%p", table2))
	   where(fmt.Sprintf("\n%s", table2))

	   where(fmt.Sprintf("%p", table3))
	   where(fmt.Sprintf("\n%s", table3))
	*/
}

//	func TestGetTableSetAsJSON_int64_float64_conversion(t *testing.T) {
//		//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
//		const verbose bool = false
//
//		var err error
//		var tableSet1 *TableSet
//		var tableSet2 *TableSet
//
//		var tableSet1String string = `
//		[TypesGalore12]
//	    i   s      f       f32     t     b    ui    bb            uu8
//	    int string float64 float32 bool  byte uint8 []byte        []uint8
//	    1   "abc"  2.3     6.6     true  11   0     [11 12 13 14] [15 16 17]
//	    2   "xyz"  4.5     7.7     false 22   1     [22 23 24 25] [26 27 28]
//	    3   "ssss" 4.9     8.8     false 33   2     [33 34 35 36] [37 38 39]
//	    4   "xxxx" 5.9     9.9     true  44   3     []            []
//
//		[AnotherTable]
//		fred int = 33
//		wilma int = 29
//
//		[Tminus1]
//		f32 float32 = 28
//		f64 float64 = 3.402823e+38
//		bt byte = 65
//		u8 uint8 = 99
//		u16 uint16 = 116
//		u32 uint32 = 500
//		u64 uint64 = 900
//		maxint int = 9223372036854775807
//		iii2 int = 13
//		iii3 int = -20
//		uInt4 uint = 4294967295
//		maxuint uint = 18446744073709551615
//		i8 int8 = -128
//		i16 int16 = -32768
//		i32 int32 = 66
//		i64 int64 = 900
//		s string = "something"
//		bo bool = true
//		r rune = 'A'
//		bta []byte = [65 66 67]
//		u8a []uint8 = [97 98 99]
//		t time.Time = 2020-03-15T14:22:30.123456789+17:00
//	    `
//		tableSet1, err = NewTableSetFromString(tableSet1String)
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		err = tableSet1.SetName("MySet")
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		var jsonTableSet string
//		jsonTableSet, err = tableSet1.GetTableSetAsJSON()
//		if err != nil {
//			t.Fatal(err)
//		}
//		_ = jsonTableSet
//
//		if verbose {
//			fmt.Println(jsonTableSet)
//			var out bytes.Buffer
//			// For readability.
//			err = json.Indent(&out, []byte(jsonTableSet), "", "\t")
//			if err != nil {
//				t.Fatal(err)
//			}
//
//			_, _ = out.WriteTo(os.Stdout)
//			fmt.Println()
//		}
//
//		// Now turn it back into a TableSet.
//		tableSet2, err = NewTableSetFromJSON(jsonTableSet)
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		_, err = tableSet1.Equals(tableSet2)
//		if err != nil {
//			t.Fatal(err)
//		}
//	}
