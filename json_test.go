package gotables

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
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
	var err error
	var tableSet *TableSet

	var tableSetString string = `
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
	tableSet, err = NewTableSetFromString(tableSetString)
	if err != nil {
		t.Fatal(err)
	}

	tableSet.SetName("MySet")

	var jsonStrings []string
	jsonStrings, err = tableSet.GetTableSetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	_ = jsonStrings

	/*
		fmt.Println(jsonStrings)
	*/

	/*
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

func TestGetTableMetadataAsJSON(t *testing.T) {
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
	jsonString, err = table.GetTableMetadataAsJSON()
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

func TestNewTableFromJSON_bothDirections(t *testing.T) {

	const verbose = false

	var err error
	var tableInput *Table // Input table
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

where(jsonString)
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

	const verbose = false

	var err error
	var tableInput *Table // Input table
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

	const verbose = false

	var err error
	var tableInput *Table // Input table
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
	var verbose bool = true
	var err error
	var tableSet1 *TableSet // Input tableSet
	var tableSet2 *TableSet // Output tableSet

	var tableSetString string = `
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

	var jsonStrings []string
	jsonStrings, err = tableSet1.GetTableSetAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		var buf bytes.Buffer
		// For readability.
		for i := 0; i < len(jsonStrings); i++ {
			err = json.Indent(&buf, []byte(jsonStrings[i]), "", "\t")
			if err != nil {
				t.Fatal(err)
			}
			_, _ = buf.WriteTo(os.Stdout)
		}
	}

	tableSet2, err = NewTableSetFromJSON(jsonStrings)
	if err != nil {
		t.Fatal(err)
	}

	_, err = tableSet1.Equals(tableSet2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAllJSON(t *testing.T) {
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
		`[MyTable]
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

	var jsonSlice []string

	jsonSlice, err = tableSet1.GetTableSetAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		for i := 0; i < len(jsonSlice); i++ {
			// For readability.
			err = json.Indent(&buf, []byte(jsonSlice[i]), "", "\t")
			if err != nil {
				t.Fatal(err)
			}
			_, _ = buf.WriteTo(os.Stdout)
		}

		fmt.Println()
		fmt.Println()
	}

	tableSet2, err := NewTableSetFromJSON(jsonSlice)
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
	var err error

	tableString :=
		`[MyTable]
	i    u    f
	int  uint float32
	1    2    3.3
	4    5    6.6
	`
	table1, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table1)

	var jsonString string
	var buf bytes.Buffer

	jsonString, err = table1.GetTableMetadataAsJSON()
	if err != nil {
		log.Println(err)
	}

	// For readability.
	err = json.Indent(&buf, []byte(jsonString), "", "  ")
	if err != nil {
		log.Println(err)
	}
	_, _ = buf.WriteTo(os.Stdout)

	fmt.Println()
	fmt.Println()

	table2, err := NewTableFromJSON(jsonString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table2)

	equals, err := table2.Equals(table1)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("table2.Equals(table1) == %t\n", equals)

	// Output:
	// [MyTable]
	//   i    u       f
	// int uint float32
	//   1    2     3.3
	//   4    5     6.6
	//
	// {
	//   "MyTable": [
	//     {
	//       "i": "int"
	//     },
	//     {
	//       "u": "uint"
	//     },
	//     {
	//       "f": "float32"
	//     }
	//   ]
	// }
	//
	// {
	//   "MyTable": [
	//     {
	//       "i": 1,
	//       "u": 2,
	//       "f": 3.3
	//     },
	//     {
	//       "i": 4,
	//       "u": 5,
	//       "f": 6.6
	//     }
	//   ]
	// }
	//
	// [MyTable]
	//   i    u       f
	// int uint float32
	//   1    2     3.3
	//   4    5     6.6
	//
	// table2.Equals(table1) == true
}

/*
	It is permitted to have zero rows.

	There must always be columns defined -- name(s) and type(s).
*/
func ExampleNewTableFromJSON_zeroRows() {
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
	var jsonData string
	var buf bytes.Buffer

	jsonString, err = table1.GetTableMetadataAsJSON()
	if err != nil {
		log.Println(err)
	}

	// For readability.
	err = json.Indent(&buf, []byte(jsonString), "", "  ")
	if err != nil {
		log.Println(err)
	}
	_, _ = buf.WriteTo(os.Stdout)

	fmt.Println()
	fmt.Println()

	table2, err := NewTableFromJSON(jsonData)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table2)

	equals, err := table2.Equals(table1)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("table2.Equals(table1) == %t\n", equals)

	// Output:
	// [MyTable]
	//   i    u       f
	// int uint float32
	//
	// {
	//   "MyTable": [
	//     {
	//       "i": "int"
	//     },
	//     {
	//       "u": "uint"
	//     },
	//     {
	//       "f": "float32"
	//     }
	//   ]
	// }
	//
	// {
	//   "MyTable": []
	// }
	//
	// [MyTable]
	//   i    u       f
	// int uint float32
	//
	// table2.Equals(table1) == true
}

func ExampleNewTableSetFromJSON() {
	var err error

	tableSetString :=
		`[MyTable]
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

	var jsonSlice []string
	var buf bytes.Buffer

	jsonSlice, err = tableSet1.GetTableSetAsJSON()
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(jsonSlice); i++ {
		// For readability.
		err = json.Indent(&buf, []byte(jsonSlice[i]), "", "  ")
		if err != nil {
			log.Println(err)
		}
		_, _ = buf.WriteTo(os.Stdout)
	}

	fmt.Println()
	fmt.Println()

	tableSet2, err := NewTableSetFromJSON(jsonSlice)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(tableSet2)

	equals, err := tableSet2.Equals(tableSet1)

	fmt.Printf("table2.Equals(table1) == %t\n", equals)

	// Output:
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
	//   "MyTable": [
	//     {
	//       "x": "int"
	//     },
	//     {
	//       "y": "int"
	//     },
	//     {
	//       "z": "int"
	//     }
	//   ]
	// }{
	//   "MyTable": [
	//     {
	//       "x": 1,
	//       "y": 2,
	//       "z": 3
	//     }
	//   ]
	// }{
	//   "YourTable": [
	//     {
	//       "a": "float32"
	//     },
	//     {
	//       "b": "float32"
	//     },
	//     {
	//       "c": "float32"
	//     }
	//   ]
	// }{
	//   "YourTable": [
	//     {
	//       "a": 1,
	//       "b": 2,
	//       "c": 3
	//     }
	//   ]
	// }
	//
	// [MyTable]
	//   x   y   z
	// int int int
	//   1   2   3
	//
	// [YourTable]
	//       a       b       c
	// float32 float32 float32
	//       1       2       3
	//
	// table2.Equals(table1) == true
}

/*
	It is permitted to have zero rows.

	There must always be columns defined -- name(s) and type(s).
*/
func ExampleNewTableSetFromJSON_zeroRows() {
	var err error

	tableSetString :=
		`[MyTable]
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

	var jsonSlice []string
	var buf bytes.Buffer

	jsonSlice, err = tableSet1.GetTableSetAsJSON()
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(jsonSlice); i++ {
		// For readability.
		err = json.Indent(&buf, []byte(jsonSlice[i]), "", "  ")
		if err != nil {
			log.Println(err)
		}
		_, _ = buf.WriteTo(os.Stdout)
	}

	fmt.Println()
	fmt.Println()

	tableSet2, err := NewTableSetFromJSON(jsonSlice)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(tableSet2)

	equals, err := tableSet2.Equals(tableSet1)

	fmt.Printf("table2.Equals(table1) == %t\n", equals)

	// Output:
	// [MyTable]
	//   x   y   z
	// int int int
	//
	// [YourTable]
	//       a       b       c
	// float32 float32 float32
	//
	// {
	//   "MyTable": [
	//     {
	//       "x": "int"
	//     },
	//     {
	//       "y": "int"
	//     },
	//     {
	//       "z": "int"
	//     }
	//   ]
	// }{
	//   "MyTable": []
	// }{
	//   "YourTable": [
	//     {
	//       "a": "float32"
	//     },
	//     {
	//       "b": "float32"
	//     },
	//     {
	//       "c": "float32"
	//     }
	//   ]
	// }{
	//   "YourTable": []
	// }
	//
	// [MyTable]
	//   x   y   z
	// int int int
	//
	// [YourTable]
	//       a       b       c
	// float32 float32 float32
	//
	// table2.Equals(table1) == true
}

func ExampleTable_GetTableDataAsJSON() {
	var err error
	var table *Table

	var tableString string = `
	[TypesGalore18]
    i   s      f       f32     t     b    ui    bb            uu8
    int string float64 float32 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     6.6     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     7.7     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     8.8     false 33   2     [33 34 35 36] [37 38 39]
    4   "xxxx" 5.9     9.9     true  44   3     []            []
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	var jsonString string
	jsonString, err = table.GetTableAsJSON()
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

	// Output:
	// Print as is:
	//
	// {"TypesGalore18":[{"i":1,"s":"abc","f":2.3,"f32":6.6,"t":true,"b":11,"ui":0,"bb":[11,12,13,14],"uu8":[15,16,17]},{"i":2,"s":"xyz","f":4.5,"f32":7.7,"t":false,"b":22,"ui":1,"bb":[22,23,24,25],"uu8":[26,27,28]},{"i":3,"s":"ssss","f":4.9,"f32":8.8,"t":false,"b":33,"ui":2,"bb":[33,34,35,36],"uu8":[37,38,39]},{"i":4,"s":"xxxx","f":5.9,"f32":9.9,"t":true,"b":44,"ui":3,"bb":[],"uu8":[]}]}
	//
	// Print indented for readability:
	//
	// {
	// 	"TypesGalore18": [
	// 		{
	// 			"i": 1,
	// 			"s": "abc",
	// 			"f": 2.3,
	// 			"f32": 6.6,
	// 			"t": true,
	// 			"b": 11,
	// 			"ui": 0,
	// 			"bb": [
	// 				11,
	// 				12,
	// 				13,
	// 				14
	// 			],
	// 			"uu8": [
	// 				15,
	// 				16,
	// 				17
	// 			]
	// 		},
	// 		{
	// 			"i": 2,
	// 			"s": "xyz",
	// 			"f": 4.5,
	// 			"f32": 7.7,
	// 			"t": false,
	// 			"b": 22,
	// 			"ui": 1,
	// 			"bb": [
	// 				22,
	// 				23,
	// 				24,
	// 				25
	// 			],
	// 			"uu8": [
	// 				26,
	// 				27,
	// 				28
	// 			]
	// 		},
	// 		{
	// 			"i": 3,
	// 			"s": "ssss",
	// 			"f": 4.9,
	// 			"f32": 8.8,
	// 			"t": false,
	// 			"b": 33,
	// 			"ui": 2,
	// 			"bb": [
	// 				33,
	// 				34,
	// 				35,
	// 				36
	// 			],
	// 			"uu8": [
	// 				37,
	// 				38,
	// 				39
	// 			]
	// 		},
	// 		{
	// 			"i": 4,
	// 			"s": "xxxx",
	// 			"f": 5.9,
	// 			"f32": 9.9,
	// 			"t": true,
	// 			"b": 44,
	// 			"ui": 3,
	// 			"bb": [],
	// 			"uu8": []
	// 		}
	// 	]
	// }
}

func ExampleTable_GetTableMetadataAsJSON() {
	var err error
	var table *Table

	var tableString string = `
	[TypesGalore19]
    i   s      f       f32     t     b    ui    bb            uu8
    int string float64 float32 bool  byte uint8 []byte        []uint8
    1   "abc"  2.3     6.6     true  11   0     [11 12 13 14] [15 16 17]
    2   "xyz"  4.5     7.7     false 22   1     [22 23 24 25] [26 27 28]
    3   "ssss" 4.9     8.8     false 33   2     [33 34 35 36] [37 38 39]
    4   "xxxx" 5.9     9.9     true  44   3     []            []
    `
	table, err = NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	var jsonString string
	jsonString, err = table.GetTableMetadataAsJSON()
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
	fmt.Println()

	// Output:
	// Print as is:
	// 
	// {"TypesGalore19":[{"i":"int"},{"s":"string"},{"f":"float64"},{"f32":"float32"},{"t":"bool"},{"b":"byte"},{"ui":"uint8"},{"bb":"[]byte"},{"uu8":"[]uint8"}]}
	//
	// Print indented for readability:
	//
	// {
	// 	"TypesGalore19": [
	// 		{
	// 			"i": "int"
	// 		},
	// 		{
	// 			"s": "string"
	// 		},
	// 		{
	// 			"f": "float64"
	// 		},
	// 		{
	// 			"f32": "float32"
	// 		},
	// 		{
	// 			"t": "bool"
	// 		},
	// 		{
	// 			"b": "byte"
	// 		},
	// 		{
	// 			"ui": "uint8"
	// 		},
	// 		{
	// 			"bb": "[]byte"
	// 		},
	// 		{
	// 			"uu8": "[]uint8"
	// 		}
	// 	]
	// }
}

//func ExampleTable_GetTableMetadataAsJSON_nestedTables() {
//	var err error
//	var table *Table
//
//	var tableString string = `
//	[TypesGalore20]
//    i   s      f       f32     t     b    ui    bb            uu8        right
//    int string float64 float32 bool  byte uint8 []byte        []uint8    *Table
//    0   "abc"  2.3     6.6     true  11   0     [11 12 13 14] [15 16 17] []
//    1   "xyz"  4.5     7.7     false 22   1     [22 23 24 25] [26 27 28] []
//    2   "ssss" 4.9     8.8     false 33   2     [33 34 35 36] [37 38 39] []
//    3   "xxxx" 5.9     9.9     true  44   3     []            []         []
//    `
//	table, err = NewTableFromString(tableString)
//	if err != nil {
//		log.Println(err)
//	}
//
//	// Now create and set some table cell tables.
//	right0 := `
//	[right0]
//	i int = 32`
//
//	right1 := `
//	[right1]
//	s string = "thirty-two"`
//
//	right2 := `
//	[right2]
//	x	y	z
//	int	int	int
//	1	2	3
//	4	5	6
//	7	8	9`
//
//	right3 := `
//	[right3]
//	f float32 = 88.8`
//
//	table.SetTableMustSet("right", 0, NewTableFromStringMustMake(right0))
//	table.SetTableMustSet("right", 1, NewTableFromStringMustMake(right1))
//	table.SetTableMustSet("right", 2, NewTableFromStringMustMake(right2))
//	table.SetTableMustSet("right", 3, NewTableFromStringMustMake(right3))
//
//	var jsonString string
//	jsonString, err = table.GetTableMetadataAsJSON()
//	if err != nil {
//		log.Println(err)
//	}
//
//	fmt.Println()
//	fmt.Println("NOTE: The following output is not what you would at first expect.")
//	fmt.Println("      The metadata of only the top-level table is marshalled.")
//	fmt.Println("      Nested table metadata is marshalled WITH the data itself.")
//	fmt.Println()
//
//	fmt.Println("Print as is:")
//	fmt.Println()
//	fmt.Println(jsonString)
//	fmt.Println()
//
//	fmt.Println("Print indented for readability:")
//	fmt.Println()
//
//	var out bytes.Buffer
//	err = json.Indent(&out, []byte(jsonString), "", "\t")
//	if err != nil {
//		log.Println(err)
//	}
//
//	_, _ = out.WriteTo(os.Stdout)
//	fmt.Println()
//
//	// Output:
//	// NOTE: The following output is not what you would at first expect.
//	//       The metadata of only the top-level table is marshalled.
//	//       Nested table metadata is marshalled WITH the data itself.
//	//
//	// Print as is:
//	//
//	// {"TypesGalore20":[{"i":"int"},{"s":"string"},{"f":"float64"},{"f32":"float32"},{"t":"bool"},{"b":"byte"},{"ui":"uint8"},{"bb":"[]byte"},{"uu8":"[]uint8"},{"right":"*Table"}]}
//	//
//	// Print indented for readability:
//	//
//	// {
//	// 	"TypesGalore20": [
//	// 		{
//	// 			"i": "int"
//	// 		},
//	// 		{
//	// 			"s": "string"
//	// 		},
//	// 		{
//	// 			"f": "float64"
//	// 		},
//	// 		{
//	// 			"f32": "float32"
//	// 		},
//	// 		{
//	// 			"t": "bool"
//	// 		},
//	// 		{
//	// 			"b": "byte"
//	// 		},
//	// 		{
//	// 			"ui": "uint8"
//	// 		},
//	// 		{
//	// 			"bb": "[]byte"
//	// 		},
//	// 		{
//	// 			"uu8": "[]uint8"
//	// 		},
//	// 		{
//	// 			"right": "*Table"
//	// 		}
//	// 	]
//	// }
//}

//func ExampleTable_GetTableDataAsJSON_nestedTables() {
//	var err error
//	var table *Table
//
//	var tableString string = `
//	[TypesGalore21]
//    i   s      f       f32     t     b    ui    bb            uu8        right
//    int string float64 float32 bool  byte uint8 []byte        []uint8    *Table
//    0   "abc"  2.3     6.6     true  11   0     [11 12 13 14] [15 16 17] []
//    1   "xyz"  4.5     7.7     false 22   1     [22 23 24 25] [26 27 28] []
//    2   "ssss" 4.9     8.8     false 33   2     [33 34 35 36] [37 38 39] []
//    3   "xxxx" 5.9     9.9     true  44   3     []            []         []
//    `
//	table, err = NewTableFromString(tableString)
//	if err != nil {
//		log.Println(err)
//	}
//
//	// Now create and set some table cell tables.
//	right0 := `
//	[right0]
//	i int = 32`
//
//	right1 := `
//	[right1]
//	s string = "thirty-two"`
//
//	right2 := `
//	[right2]
//	x	y	z
//	int	int	int
//	1	2	3
//	4	5	6
//	7	8	9`
//
//	right3 := `
//	[right3]
//	f float32 = 88.8`
//
//	table.SetTableMustSet("right", 0, NewTableFromStringMustMake(right0))
//	table.SetTableMustSet("right", 1, NewTableFromStringMustMake(right1))
//	table.SetTableMustSet("right", 2, NewTableFromStringMustMake(right2))
//	table.SetTableMustSet("right", 3, NewTableFromStringMustMake(right3))
//
//	jsonString, err := table.GetTableDataAsJSON()
//	if err != nil {
//		log.Println(err)
//	}
//
//	fmt.Println("Print as is:")
//	fmt.Println()
//	fmt.Println(jsonString)
//	fmt.Println()
//
//	fmt.Println("Print indented for readability:")
//	fmt.Println()
//	var out bytes.Buffer
//	err = json.Indent(&out, []byte(jsonString), "", "\t")
//	if err != nil {
//		log.Println(err)
//	}
//	_, _ = out.WriteTo(os.Stdout)
//
//	// Now let's get it back from JSON into *Table
//
//	// First, we need the table metadata as JSON to get the data back
//	jsonMetadataString, err := table.GetTableMetadataAsJSON()
//	if err != nil {
//		log.Println(err)
//	}
//
//	tableFromJSON, err := NewTableFromJSON(jsonMetadataString, jsonString)
//	if err != nil {
//		log.Println(err)
//	}
//	fmt.Println(tableFromJSON)
//
//	// Output:
//	// Print as is:
//	// 
//	// {"TypesGalore21":[{"i":0,"s":"abc","f":2.3,"f32":6.6,"t":true,"b":11,"ui":0,"bb":[11,12,13,14],"uu8":[15,16,17],"right":{"right0":[{"i":32}]}},{"i":1,"s":"xyz","f":4.5,"f32":7.7,"t":false,"b":22,"ui":1,"bb":[22,23,24,25],"uu8":[26,27,28],"right":{"right1":[{"s":"thirty-two"}]}},{"i":2,"s":"ssss","f":4.9,"f32":8.8,"t":false,"b":33,"ui":2,"bb":[33,34,35,36],"uu8":[37,38,39],"right":{"right2":[{"x":1,"y":2,"z":3},{"x":4,"y":5,"z":6},{"x":7,"y":8,"z":9}]}},{"i":3,"s":"xxxx","f":5.9,"f32":9.9,"t":true,"b":44,"ui":3,"bb":[],"uu8":[],"right":{"right3":[{"f":88.8}]}}]}
//	// 
//	// Print indented for readability:
//	//
//	// {
//	// 	"TypesGalore21": [
//	// 		{
//	// 			"i": 0,
//	// 			"s": "abc",
//	// 			"f": 2.3,
//	// 			"f32": 6.6,
//	// 			"t": true,
//	// 			"b": 11,
//	// 			"ui": 0,
//	// 			"bb": [
//	// 				11,
//	// 				12,
//	// 				13,
//	// 				14
//	// 			],
//	// 			"uu8": [
//	// 				15,
//	// 				16,
//	// 				17
//	// 			],
//	// 			"right": {
//	// 				"right0": [
//	// 					{
//	// 						"i": 32
//	// 					}
//	// 				]
//	// 			}
//	// 		},
//	// 		{
//	// 			"i": 1,
//	// 			"s": "xyz",
//	// 			"f": 4.5,
//	// 			"f32": 7.7,
//	// 			"t": false,
//	// 			"b": 22,
//	// 			"ui": 1,
//	// 			"bb": [
//	// 				22,
//	// 				23,
//	// 				24,
//	// 				25
//	// 			],
//	// 			"uu8": [
//	// 				26,
//	// 				27,
//	// 				28
//	// 			],
//	// 			"right": {
//	// 				"right1": [
//	// 					{
//	// 						"s": "thirty-two"
//	// 					}
//	// 				]
//	// 			}
//	// 		},
//	// 		{
//	// 			"i": 2,
//	// 			"s": "ssss",
//	// 			"f": 4.9,
//	// 			"f32": 8.8,
//	// 			"t": false,
//	// 			"b": 33,
//	// 			"ui": 2,
//	// 			"bb": [
//	// 				33,
//	// 				34,
//	// 				35,
//	// 				36
//	// 			],
//	// 			"uu8": [
//	// 				37,
//	// 				38,
//	// 				39
//	// 			],
//	// 			"right": {
//	// 				"right2": [
//	// 					{
//	// 						"x": 1,
//	// 						"y": 2,
//	// 						"z": 3
//	// 					},
//	// 					{
//	// 						"x": 4,
//	// 						"y": 5,
//	// 						"z": 6
//	// 					},
//	// 					{
//	// 						"x": 7,
//	// 						"y": 8,
//	// 						"z": 9
//	// 					}
//	// 				]
//	// 			}
//	// 		},
//	// 		{
//	// 			"i": 3,
//	// 			"s": "xxxx",
//	// 			"f": 5.9,
//	// 			"f32": 9.9,
//	// 			"t": true,
//	// 			"b": 44,
//	// 			"ui": 3,
//	// 			"bb": [],
//	// 			"uu8": [],
//	// 			"right": {
//	// 				"right3": [
//	// 					{
//	// 						"f": 88.8
//	// 					}
//	// 				]
//	// 			}
//	// 		}
//	// 	]
//	// }
//}

func ExampleTable_GetTableAsJSON_nestedTablesCircularReference() {
	var err error
	var table *Table

	// Minor note: both nil and [] are acceptable empty table placeholders.
	//             Each will result in a NilTable with no table name.
	//             To make the table usable, give it a table name.

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
	table.SetTableMustSet("right", 0, table)	// table already exists (at the top level)
	fmt.Printf("%s", table)
	_, err = table.GetTableAsJSON()
	if err != nil {
		// Error prints here.
		fmt.Println(err)
	}
	fmt.Println()

	fmt.Println("Now try again with a COPY of the same table, which will have a new reference.")
	var jsonString string
	var tableCopy *Table
	tableCopy, err = table.Copy(true)
	if err != nil {
		// No error to print here.
		fmt.Println(err)
	}
	err = tableCopy.SetName("TableCopy")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("By the way, don't try to set table 'right' to <nil>. Not allowed. Must use an actual *Table reference.")
	err = tableCopy.SetTable("right", 0, nil)	// Not allowed. Must use an actual *Table reference.
	if err != nil {
		fmt.Println(err)
	}
	err = tableCopy.SetTable("right", 0, NewNilTable())	// Otherwise this is another circular reference.
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", tableCopy)
	fmt.Println()

	fmt.Println("This should succeed: We are assigning a DIFFERENT table (same contents doesn't matter).")
	table.SetTableMustSet("right", 0, tableCopy)	// Different table reference.
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
	jsonString, err = table.GetTableAsJSONIndent("", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonString)

	fmt.Println()

	fmt.Println("This should fail: We are assigning the same table to multiple cells.")
	table.SetTableMustSet("left", 0, tableCopy)	// Different table reference.
	fmt.Printf("%s", table)
	jsonString, err = table.GetTableAsJSON()
	if err != nil {
		// Error prints here.
		fmt.Println(err)
	}

	// Output:
	// This should fail: We are assigning the same table as the parent.
	// [SameTableReference]
	// left     i s      right
	// *Table int string *Table
	// []      42 "abc"  [SameTableReference]
	// getTableAsJSON_recursive(): circular reference: a reference to table [SameTableReference] already exists
	//
	// Now try again with a COPY of the same table, which will have a new reference.
	// By the way, don't try to set table 'right' to <nil>. Not allowed. Must use an actual *Table reference.
	// SetTable(): table [TableCopy] col right expecting val of type *Table, not: <nil> [use NewNilTable() instead of <nil>]
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
	// {"SameTableReference":{"metadata::SameTableReference":[{"left":"*Table"},{"i":"int"},{"s":"string"},{"right":"*Table"}]},"data::SameTableReference":[[{"left":null},{"i":42},{"s":"abc"},{"right":{"TableCopy":{"metadata::TableCopy":[{"left":"*Table"},{"i":"int"},{"s":"string"},{"right":"*Table"}]},"data::TableCopy":[[{"left":null},{"i":42},{"s":"abc"},{"right":null}]]}}]]}
	//
	// Print indented for readability:
	// {
	// 	"SameTableReference": {
	// 		"metadata::SameTableReference": [
	// 			{
	// 				"left": "*Table"
	// 			},
	// 			{
	// 				"i": "int"
	// 			},
	// 			{
	// 				"s": "string"
	// 			},
	// 			{
	// 				"right": "*Table"
	// 			}
	// 		]
	// 	},
	// 	"data::SameTableReference": [
	// 		[
	// 			{
	// 				"left": null
	// 			},
	// 			{
	// 				"i": 42
	// 			},
	// 			{
	// 				"s": "abc"
	// 			},
	// 			{
	// 				"right": {
	// 					"TableCopy": {
	// 						"metadata::TableCopy": [
	// 							{
	// 								"left": "*Table"
	// 							},
	// 							{
	// 								"i": "int"
	// 							},
	// 							{
	// 								"s": "string"
	// 							},
	// 							{
	// 								"right": "*Table"
	// 							}
	// 						]
	// 					},
	// 					"data::TableCopy": [
	// 						[
	// 							{
	// 								"left": null
	// 							},
	// 							{
	// 								"i": 42
	// 							},
	// 							{
	// 								"s": "abc"
	// 							},
	// 							{
	// 								"right": null
	// 							}
	// 						]
	// 					]
	// 				}
	// 			}
	// 		]
	// 	]
	// }
	//
	// This should fail: We are assigning the same table to multiple cells.
	// [SameTableReference]
	// left          i s      right
	// *Table      int string *Table
	// [TableCopy]  42 "abc"  [TableCopy]
	// getTableAsJSON_recursive(): circular reference: a reference to table [TableCopy] already exists
}

func ExampleTable_GetTableAsJSON_nestedTables() {
	var err error
	var table *Table

	var tableString string
/*
	tableString = `
	[TypesGalore22]
    i   s      f       f32     t     b    ui    bb            uu8        right
    int string float64 float32 bool  byte uint8 []byte        []uint8    *Table
    0   "abc"  2.3     6.6     true  11   0     [11 12 13 14] [15 16 17] []
    1   "xyz"  4.5     7.7     false 22   1     [22 23 24 25] [26 27 28] []
    2   "ssss" 4.9     8.8     false 33   2     [33 34 35 36] [37 38 39] []
    3   "xxxx" 5.9     9.9     true  44   3     []            []         []
    `
*/
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
/*
	tableString = `
	[TypesGalore22]
    i   s		ii	ss		iii
    int string	int	string	int
    0   "abc"	0	"def"	0
    1   "xyz"	1	"ghi"	1
    2   "ssss"	2	"jkl"	2
    3   "xxxx"	3	"mno"	3
    `
*/
/*
	tableString = `
	[SimpleTable]
    i   s		f
    int string	float32
    0   "abc"	0.0
    1   "xyz"	1.1
    `
*/
/*
	tableString = `
	[TypesGalore22]
	a	t		b
	int	*Table	int
	1	[]		3
	2	[]		4
	`
*/
	table, err = NewTableFromString(tableString)
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

/*
	right4 := `
	[right4]
	t1 *Table = []
	t2 *gotables.Table = []`
*/

	table.SetTableMustSet("right", 0, NewTableFromStringMustMake(right0))
	table.SetTableMustSet("right", 1, NewTableFromStringMustMake(right1))
	table.SetTableMustSet("right", 2, NewTableFromStringMustMake(right2))
	table.SetTableMustSet("right", 3, NewTableFromStringMustMake(right3))
	table.SetTableMustSet("right", 4, table)
/*
	tableRight4 := NewTableFromStringMustMake(right4)
	table.SetTableMustSet("right", 4, tableRight4)
	tableRight4.SetTableMustSet("t1", 0, NewTableFromStringMustMake(right2))
	tableRight4.SetTableMustSet("t2", 0, NewTableFromStringMustMake(right4))
*/

	var jsonString string
//	jsonString, err = table.GetTableMetadataAsJSON()	// okay
	if err != nil {
		log.Println(err)
	}
//	jsonString, err = table.GetTableDataAsJSON()		// okay
	if err != nil {
		log.Println(err)
	}
	jsonString, err = table.GetTableAsJSON()			// okay
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

	// Output:

}
