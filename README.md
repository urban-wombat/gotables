# gotable

Table data format and utilities


## What Is A gotable.Table?

A gotable.Table is a table of data with the following sections:
1. A table name in square brackets.
2. A row of 1 or more column names and data types.
3. Rows of data.

Here's an example:

    [my_table]
    Flintstone Age Married Height
    string     int bool    float32
    "Fred"     33  true    1.8
    "Wilma"    31  true    1.6
    "Pebbles"   4  false   0.8

Many of the Go data types can be used. (Not yet implemented: complex64, complex128, rune, byte.)

Here is a simple program that parses the table into a gotable.TableSet (a set of tables -- in this case one table) and echoes it back out:

    package main

    import (
        "github.com/urban-wombat/gotable"
        "fmt"
    )

    var myTable string = `
        [my_table]
        Flintstone Age Married Height
        string     int bool    float32
        "Fred"     33  true    1.8
        "Wilma"    31  true    1.6
        "Pebbles"   4  false   0.8
    `

    func main() {
        tables, err := gotable.NewTableSetFromString(myTable)
        if err != nil {
            panic(err)
        }
        fmt.Println(tables)
    }

The output is:

    [my_table]
    Flintstone Age Married  Height
    string     int bool    float32
    "Fred"      33 true        1.8
    "Wilma"     31 true        1.6
    "Pebbles"    4 false       0.8
