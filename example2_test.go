package gotables

import (
	"fmt"
	"log"
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

/*
	Note:

	For these examples to compile and run for you, you need to import "github.com/urban-wombat/gotables"
	and prefix function and method calls with gotables.
*/

func ExampleSetByteSlice() {
	// Create a struct-shaped table with two cells
	// for a string and a byte, just to show how it's done.
	var s string =
	`[myTableStruct]
	hello string = "Hello world!"
	b byte = 255
	`

	// Instantiate table from string.
	table, err := NewTableFromString(s)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(table)

	// Create a []byte (slice of byte) from another string.
	hello := "Hello slice!"
	var helloByteSlice []byte = make([]byte, len(hello))
	for i, val := range hello {
		fmt.Printf("%s = byte %d\n", string(val), byte(val))
		helloByteSlice[i] = byte(val)
	}
	fmt.Printf("helloByteSlice = %v\n", helloByteSlice)

	// Create a new column called helloSlice in the table.
	err = table.AppendCol("helloSlice", "[]byte")
	if err != nil {
		log.Println(err)
	}

	// Assign helloByteSlice to helloSlice cell at row 0
	err = table.SetByteSlice("helloSlice", 0, helloByteSlice)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("\n%s\n", table)

	// Retrieve it.
	var retrievedSlice []byte
	retrievedSlice, err = table.GetByteSlice("helloSlice", 0)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("retrievedSlice = %v\n", retrievedSlice)

	// Output:
	// [myTableStruct]
	// hello string = "Hello world!"
	// b byte = 255
	// 
	// H = byte 72
	// e = byte 101
	// l = byte 108
	// l = byte 108
	// o = byte 111
	//   = byte 32
	// s = byte 115
	// l = byte 108
	// i = byte 105
	// c = byte 99
	// e = byte 101
	// ! = byte 33
	// helloByteSlice = [72 101 108 108 111 32 115 108 105 99 101 33]
	// 
	// [myTableStruct]
	// hello string = "Hello world!"
	// b byte = 255
	// helloSlice []byte = [72 101 108 108 111 32 115 108 105 99 101 33]
	// 
	// retrievedSlice = [72 101 108 108 111 32 115 108 105 99 101 33]
}
