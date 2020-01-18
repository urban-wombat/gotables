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

func TestGetTableDataAsJSON(t *testing.T) {
	var err error
	var table *Table

	var tableString string = `
	[TypesGalore11]
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
	jsonString, err = table.GetTableDataAsJSON()
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

		out.WriteTo(os.Stdout)
		fmt.Println()
	*/
}

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

	var jsonMetadataString []string
	var jsonDataString []string
	jsonMetadataString, jsonDataString, err = tableSet.GetTableSetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	_ = jsonMetadataString
	_ = jsonDataString

	/*
		fmt.Println(jsonMetadataString)
		fmt.Println(jsonDataString)
	*/

	/*
		var out bytes.Buffer
		// For readability.
		err = json.Indent(&out, []byte(jsonString), "", "\t")
		if err != nil {
			t.Fatal(err)
		}

		out.WriteTo(os.Stdout)
		fmt.Println()
	*/
}

func TestGetTableMetadataAsJSON(t *testing.T) {
	var err error
	var table *Table

	var tableString string = `
	[TypesGalore13]
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

		out.WriteTo(os.Stdout)
		fmt.Println()
	*/
}

//func TestGetTableSetMetadataAsJSON(t *testing.T) {
//	var err error
//	var tableSet *TableSet
//
//	var tableSetString string = `
//	[TypesGalore14]
//    i   s      f       f32     t     b    ui    bb            uu8
//    int string float64 float32 bool  byte uint8 []byte        []uint8
//    1   "abc"  2.3     6.6     true  11   0     [11 12 13 14] [15 16 17]
//    2   "xyz"  4.5     7.7     false 22   1     [22 23 24 25] [26 27 28]
//    3   "ssss" 4.9     8.8     false 33   2     [33 34 35 36] [37 38 39]
//    4   "xxxx" 5.9     9.9     true  44   3     []            []
//
//	[AnotherTable]
//	fred int = 33
//	wilma int = 29
//    `
//	tableSet, err = NewTableSetFromString(tableSetString)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	var jsonString string
//	jsonString, err = tableSet.GetTableSetMetadataAsJSON()
//	if err != nil {
//		t.Fatal(err)
//	}
//	_ = jsonString
//
//	/*
//		fmt.Println(jsonString)
//
//		var out bytes.Buffer
		// For readability.
//		err = json.Indent(&out, []byte(jsonString), "", "\t")
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		out.WriteTo(os.Stdout)
//		fmt.Println()
//	*/
//}

func TestNewTableFromJSON(t *testing.T) {

	const verbose = false

	var err error
	var table1 *Table // Input table
	var table2 *Table // Output table

	var tableString string = `
		[TypesGalore15]
	    i   s      f       f32     t     b    ui    bb            uu8
	    int string float64 float32 bool  byte uint8 []byte        []uint8
	    0   "abc"  2.3     6.6     true  11   0     [11 12 13 14] [15 16 17]
	    1   "xyz"  4.5     7.7     false 22   1     [22 23 24 25] [26 27 28]
	    2   "ssss" 4.9     8.8     false 33   2     [33 34 35 36] [37 38 39]
	    3   "xxxx" 5.9     9.9     true  44   3     []            []
	    4   "yyyy" 6.9    10.9     false 55   4     [0]           [2]
	    `
	table1, err = NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		fmt.Printf("\n\n%v\n", table1)
	}

	var jsonMetadataString string
	jsonMetadataString, err = table1.GetTableMetadataAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		var buf bytes.Buffer
		// For readability.
		err = json.Indent(&buf, []byte(jsonMetadataString), "", "\t")
		if err != nil {
			t.Fatal(err)
		}
		buf.WriteTo(os.Stdout)
	}

	var jsonString string
	jsonString, err = table1.GetTableDataAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		var buf bytes.Buffer
		// For readability.
		err = json.Indent(&buf, []byte(jsonString), "", "\t")
		if err != nil {
			t.Fatal(err)
		}
		buf.WriteTo(os.Stdout)
	}

	table2, err = NewTableFromJSON(jsonMetadataString, jsonString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = table1.Equals(table2)
	if err != nil {
		t.Fatal(err)
	}

	// Does table.Equals() check row order?
	for rowIndex := 0; rowIndex < table2.RowCount(); rowIndex++ {
		i, err := table2.GetInt("i", rowIndex)
		if err != nil {
			t.Fatal(err)
		}

		if i != rowIndex {
			t.Fatalf("expecting decoded table rows in order, but found row %d at rowIndex %d", i, rowIndex)
		}
	}

	if verbose {
		fmt.Printf("\n\n%v\n", table2)
	}
}

func TestNewTableSetFromJSON(t *testing.T) {
	var err error
	var tableSet1 *TableSet // Input tableSet
	var tableSet2 *TableSet // Output tableSet

	var tableSetString string = `
		[TypesGalore15]
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

	var jsonMetadataStrings []string
	var jsonDataStrings []string
	jsonMetadataStrings, jsonDataStrings, err = tableSet1.GetTableSetAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	/*
		var buf bytes.Buffer
		// For readability.
		err = json.Indent(&buf, []byte(jsonMetadataString), "", "\t")
		if err != nil {
			t.Fatal(err)
		}
		buf.WriteTo(os.Stdout)
	*/

	tableSet2, err = NewTableSetFromJSON(jsonMetadataStrings, jsonDataStrings)
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

	var jsonMetadata string
	var jsonData string
	var buf bytes.Buffer

	jsonMetadata, err = table1.GetTableMetadataAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		// For readability.
		err = json.Indent(&buf, []byte(jsonMetadata), "", "\t")
		if err != nil {
			t.Fatal(err)
		}
		buf.WriteTo(os.Stdout)

		fmt.Println()
		fmt.Println()
	}

	jsonData, err = table1.GetTableDataAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		// For readability.
		err = json.Indent(&buf, []byte(jsonData), "", "\t")
		if err != nil {
			t.Fatal(err)
		}
		buf.WriteTo(os.Stdout)

		fmt.Println()
		fmt.Println()
	}

	table2, err := NewTableFromJSON(jsonMetadata, jsonData)
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

	var jsonMetadataSlice []string
	var jsonDataSlice []string

	jsonMetadataSlice, jsonDataSlice, err = tableSet1.GetTableSetAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		for i := 0; i < len(jsonMetadataSlice); i++ {
			// For readability.
			err = json.Indent(&buf, []byte(jsonMetadataSlice[i]), "", "\t")
			if err != nil {
				t.Fatal(err)
			}
			buf.WriteTo(os.Stdout)

			// For readability.
			err = json.Indent(&buf, []byte(jsonDataSlice[i]), "", "\t")
			if err != nil {
				t.Fatal(err)
			}
			buf.WriteTo(os.Stdout)
		}

		fmt.Println()
		fmt.Println()
	}

	tableSet2, err := NewTableSetFromJSON(jsonMetadataSlice, jsonDataSlice)
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

func TestAllJSONZeroRows(t *testing.T) {
	const verbose = false

	var err error

	tableString :=
		`[MyTable]
	x int
	y int
	z int
	`
	table1, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where("\n" + table1.String())
		fmt.Println()
	}

	var jsonMetadata string
	var jsonData string
	var buf bytes.Buffer

	jsonMetadata, err = table1.GetTableMetadataAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		// For readability.
		err = json.Indent(&buf, []byte(jsonMetadata), "", "\t")
		if err != nil {
			t.Fatal(err)
		}
		buf.WriteTo(os.Stdout)

		fmt.Println()
		fmt.Println()
	}

	jsonData, err = table1.GetTableDataAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		// For readability.
		err = json.Indent(&buf, []byte(jsonData), "", "\t")
		if err != nil {
			t.Fatal(err)
		}
		buf.WriteTo(os.Stdout)

		fmt.Println()
		fmt.Println()
	}

	table2, err := NewTableFromJSON(jsonMetadata, jsonData)
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
	x int
	y int
	z int

	[YourTable]
	a float32
	b float32
	c float32
	`
	tableSet1, err := NewTableSetFromString(tableSetString)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where("\n" + tableSet1.String())
		fmt.Println()
	}

	var jsonMetadataSlice []string
	var jsonDataSlice []string

	jsonMetadataSlice, jsonDataSlice, err = tableSet1.GetTableSetAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		for i := 0; i < len(jsonMetadataSlice); i++ {
			// For readability.
			err = json.Indent(&buf, []byte(jsonMetadataSlice[i]), "", "\t")
			if err != nil {
				t.Fatal(err)
			}
			buf.WriteTo(os.Stdout)

			// For readability.
			err = json.Indent(&buf, []byte(jsonDataSlice[i]), "", "\t")
			if err != nil {
				t.Fatal(err)
			}
			buf.WriteTo(os.Stdout)
		}

		fmt.Println()
		fmt.Println()
	}

	tableSet2, err := NewTableSetFromJSON(jsonMetadataSlice, jsonDataSlice)
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

// Essentially, all marshalling of metadata into json fails for tables with zero metadata.
// Marshalling data (rows) into json works, as with GetTableDataAsJSON()
func TestAllJSONZeroMetadata(t *testing.T) {
	const verbose = false

	var err error

	tableString :=
		`[MyTable]`
	table1, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where("\n" + table1.String())
		fmt.Println()
	}

	var jsonData string
	var buf bytes.Buffer

	_, err = table1.GetTableMetadataAsJSON()
	if err == nil {
		t.Fatalf("expecting error (cannot marshal json metadata from a table with zero columns) but got err == nil")
	}

	jsonData, err = table1.GetTableDataAsJSON()
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		// For readability.
		err = json.Indent(&buf, []byte(jsonData), "", "\t")
		if err != nil {
			t.Fatal(err)
		}
		buf.WriteTo(os.Stdout)

		fmt.Println()
		fmt.Println()
	}

	tableSetString :=
		`[MyTable]

	[YourTable]
	`
	tableSet1, err := NewTableSetFromString(tableSetString)
	if err != nil {
		t.Fatal(err)
	}

	if verbose {
		where("\n" + tableSet1.String())
		fmt.Println()
	}

	_, _, err = tableSet1.GetTableSetAsJSON()
	if err == nil {
		t.Fatalf("expecting error (cannot marshal json metadata from a table with zero columns) but got err == nil")
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

	var jsonMetadata string
	var jsonData string
	var buf bytes.Buffer

	jsonMetadata, err = table1.GetTableMetadataAsJSON()
	if err != nil {
		log.Println(err)
	}

	// For readability.
	err = json.Indent(&buf, []byte(jsonMetadata), "", "  ")
	if err != nil {
		log.Println(err)
	}
	buf.WriteTo(os.Stdout)

	fmt.Println()
	fmt.Println()

	jsonData, err = table1.GetTableDataAsJSON()
	if err != nil {
		log.Println(err)
	}

	// For readability.
	err = json.Indent(&buf, []byte(jsonData), "", "  ")
	if err != nil {
		log.Println(err)
	}
	buf.WriteTo(os.Stdout)

	fmt.Println()
	fmt.Println()

	table2, err := NewTableFromJSON(jsonMetadata, jsonData)
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

	var jsonMetadata string
	var jsonData string
	var buf bytes.Buffer

	jsonMetadata, err = table1.GetTableMetadataAsJSON()
	if err != nil {
		log.Println(err)
	}

	// For readability.
	err = json.Indent(&buf, []byte(jsonMetadata), "", "  ")
	if err != nil {
		log.Println(err)
	}
	buf.WriteTo(os.Stdout)

	fmt.Println()
	fmt.Println()

	jsonData, err = table1.GetTableDataAsJSON()
	if err != nil {
		log.Println(err)
	}

	// For readability.
	err = json.Indent(&buf, []byte(jsonData), "", "  ")
	if err != nil {
		log.Println(err)
	}
	buf.WriteTo(os.Stdout)

	fmt.Println()
	fmt.Println()

	table2, err := NewTableFromJSON(jsonMetadata, jsonData)
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

	var jsonMetadataSlice []string
	var jsonDataSlice []string
	var buf bytes.Buffer

	jsonMetadataSlice, jsonDataSlice, err = tableSet1.GetTableSetAsJSON()
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(jsonMetadataSlice); i++ {
		// For readability.
		err = json.Indent(&buf, []byte(jsonMetadataSlice[i]), "", "  ")
		if err != nil {
			log.Println(err)
		}
		buf.WriteTo(os.Stdout)

		// For readability.
		err = json.Indent(&buf, []byte(jsonDataSlice[i]), "", "  ")
		if err != nil {
			log.Println(err)
		}
		buf.WriteTo(os.Stdout)
	}

	fmt.Println()
	fmt.Println()

	tableSet2, err := NewTableSetFromJSON(jsonMetadataSlice, jsonDataSlice)
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

	var jsonMetadataSlice []string
	var jsonDataSlice []string
	var buf bytes.Buffer

	jsonMetadataSlice, jsonDataSlice, err = tableSet1.GetTableSetAsJSON()
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(jsonMetadataSlice); i++ {
		// For readability.
		err = json.Indent(&buf, []byte(jsonMetadataSlice[i]), "", "  ")
		if err != nil {
			log.Println(err)
		}
		buf.WriteTo(os.Stdout)

		// For readability.
		err = json.Indent(&buf, []byte(jsonDataSlice[i]), "", "  ")
		if err != nil {
			log.Println(err)
		}
		buf.WriteTo(os.Stdout)
	}

	fmt.Println()
	fmt.Println()

	tableSet2, err := NewTableSetFromJSON(jsonMetadataSlice, jsonDataSlice)
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
