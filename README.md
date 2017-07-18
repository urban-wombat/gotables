# gotable

Table data format and utilities

## Why Use gotable?

1. Sometimes the data you want to represent is intrinsically tabular.
2. You want tables of data to be more readable by human beings. Be able to look at the data and spot any problems.
3. You want to eliminate repetitive metadata such as tags, and reduce the size of each tabular chunk of data.
   Data name and type are mentioned only once in a gotable Table.
4. XML and JSON are great -- especially for tree shaped data or irregular data with twigs and leaves that may or may not need to be present.
   But sometimes the data you want to represent is intrinsically tabular, and really you don't want any elements to be missing.
   And if they are, you want it to be obvious.
5. It feels like overkill to set up a relational database table (or tables) to store (and modify) your software configurations,
   or to use a database as a conduit for sharing messages or data flows between processes or threads or goroutines.
6. If you are sending messages between goroutines in Go, you can use a gotable Table or a set of Tables (a TableSet) to send
   data through your channels. A Table can be sent and received as an object or as a string.
7. gotable has methods and functions to perform tasks in these broad categories:

   a. Get and Set values. Most Go types are supported.

   b. Sort and Search a table. Multiple keys and reverse sort and search are supported.

   c. Merge two tables (with shared key(s)) into one. NaN and zero values are handled consistently.

   d. SortUnique to remove NaN and zero values.

## What Is A gotable.Table?

A gotable.Table is a table of data with the following sections:
1. A table name in square brackets.
2. A row of 1 or more column names and data types.
3. Rows of data.

It's a bit like a slice of struct.

Here's an example:

    [planets]
    name         mass distance moons index mnemonic
    string    float64  float64   int   int string
    "Mercury"   0.055      0.4     0     0 "my"
    "Venus"     0.815      0.7     0     1 "very"
    "Earth"     1.000      1.0     1     2 "elegant"
    "Mars"      0.107      1.5     2     3 "mother"
    "Jupiter" 318.000      5.2    67     4 "just"
    "Saturn"   95.000     29.4    62     5 "sat"
    "Uranus"   15.000     84.0    27     6 "upon"
    "Neptune"  17.000    164.0    13     7 "nine ... porcupines"

Most of the Go builtin data types can be used. (But not yet: complex64, complex128, rune, byte.)

Here is a simple program that parses the table into a gotable.Table and echoes it back out:

    // main_echo.go
    
    package main
    
    import (
        "github.com/urban-wombat/gotable"
        "fmt"
        "log"
    )
    
    func main() {
        tableString :=
        `[planets]
        name         mass distance moons index mnemonic
        string    float64  float64   int   int string
        "Mercury"   0.055      0.4     0     0 "my"
        "Venus"     0.815      0.7     0     1 "very"
        "Earth"     1.000      1.0     1     2 "elegant"
        "Mars"      0.107      1.5     2     3 "mother"
        "Jupiter" 318.000      5.2    67     4 "just"
        "Saturn"   95.000     29.4    62     5 "sat"
        "Uranus"   15.000     84.0    27     6 "upon"
        "Neptune"  17.000    164.0    13     7 "nine ... porcupines"
        `
    
        table, err := gotable.NewTableFromString(tableString)
        if err != nil {
            log.Println(err)
        }
    
        fmt.Println("Default String() padded output\n")
        fmt.Println(table)
    
        // Notice that the columns of data are padded with spaces, and numeric types are right-aligned.
        // This reflects the opinion that human readability is important.
    
        fmt.Println("For unpadded output use StringUnpadded()\n")
        fmt.Println(table.StringUnpadded())
    }

Here's the output:

    Default String() padded output
    
    [planets]
    name         mass distance moons index mnemonic
    string    float64  float64   int   int string
    "Mercury"   0.055      0.4     0     0 "my"
    "Venus"     0.815      0.7     0     1 "very"
    "Earth"     1.0        1.0     1     2 "elegant"
    "Mars"      0.107      1.5     2     3 "mother"
    "Jupiter" 318.0        5.2    67     4 "just"
    "Saturn"   95.0       29.4    62     5 "sat"
    "Uranus"   15.0       84.0    27     6 "upon"
    "Neptune"  17.0      164.0    13     7 "nine ... porcupines"
    
    For unpadded output use StringUnpadded()
    
    [planets]
    name mass distance moons index mnemonic
    string float64 float64 int int string
    "Mercury" 0.055 0.4 0 0 "my"
    "Venus" 0.815 0.7 0 1 "very"
    "Earth" 1 1 1 2 "elegant"
    "Mars" 0.107 1.5 2 3 "mother"
    "Jupiter" 318 5.2 67 4 "just"
    "Saturn" 95 29.4 62 5 "sat"
    "Uranus" 15 84 27 6 "upon"
    "Neptune" 17 164 13 7 "nine ... porcupines"
