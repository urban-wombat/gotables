package gotables

import (
	"fmt"
	"log"
)

/*
	Note:

	For these examples to compile and run for you, you need to import "github.com/urban-wombat/gotables"
	and prefix function and method calls with gotables.
*/

func ExampleTable_SetByteSlice_complete() {
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
