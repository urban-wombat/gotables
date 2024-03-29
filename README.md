[![Build Status](https://travis-ci.org/urban-wombat/gotables.svg?branch=master)](https://travis-ci.org/urban-wombat/gotables)

See [gotables godoc](https://godoc.org/github.com/urban-wombat/gotables)

# Gotables & FlatBuffers

The focus of recent work is on support for Google `FlatBuffers`.
If you don't want to get into the weeds of the FlatBuffers API, `gotables` will save you a lot of time.
But a warning up front: this implementation doesn't support complex data structures such as Monsters. Only rectangular tables.

To distinguish the `gotables` FlatBuffers support from something that would support Monsters, I have called it `FlatTables`.

* See [FlatBuffers Monster](https://google.github.io/flatbuffers/flatbuffers_guide_tutorial.html) for FlatBuffers Monster documentation
to see what you may be missing out on. It's suitable for gaming.

* See [The `gotables.FlatTables` implementation of Google FlatBuffers](https://github.com/urban-wombat/flattables).
It contains all the installation instructions and how to get started generating FlatBuffers code, and reading and writing FlatTables data.

In a nutshell, `gotables.FlatTables` does the following:
* Makes it very easy to create a FlatBuffers schema and associated table-type-specific Go code to access FlatBuffers data.
* Makes it a piece of cake (rectangular cake) to populate a set of one or more FlatBuffers tables as a []byte array.
* Gives you easy methods (generated specifically for your table data types) to read from a FlatBuffers []byte array.
* Provides a worked example (in the form of a main program) specific to your schema that is a tutorial you can use directly in your code.
* [This particular tutorial](https://github.com/urban-wombat/flattables_sample_main/blob/master/flattables_sample_main.go)
  is the `flattables_sample_main.go` generated for the `flattables_sample` package I created from
  [my gotables file of tables](https://github.com/urban-wombat/flattables_sample/blob/master/tables.got).
  Your own generated code will produce a tutorial main program _specific to **your** own gotables file of tables_.
  Can you believe it?

* See [https://github.com/urban-wombat/flattables](https://github.com/urban-wombat/flattables) for instructions on how to
  install `FlatTables` and run the flattables_sample package, or create and run your own.

* [https://urban-wombat.github.io](https://urban-wombat.github.io) has an introduction to `gotables`.


# gotables

`gotables` reinvents the wheel as a table.
Go (golang) Table data format is simple and self-describing.
Often data and configurations can be expressed best as tables, or can be flattened or normalised into tables.
And events which are simple structs (or records) can be batched into tables for transmission.

	`go get -u github.com/urban-wombat/gotables`

## Here's the `gotables` API ...

[gotables API](https://godoc.org/github.com/urban-wombat/gotables) - a rich set of helper functions and methods

## Here are the Release Notes ...

https://github.com/urban-wombat/gotables/releases

<a id="cli" />
## Command line utilities ...

| CLI | Purpose |
| --- | ------- |
| [gotecho](https://github.com/urban-wombat/gotables/tree/master/cmd/gotecho)           | echo 1 or all tables from a gotables file
| [gotsyntax](https://github.com/urban-wombat/gotables/tree/master/cmd/gotsyntax)       | check syntax of 1 or more gotables files
| [flattablesc](https://github.com/urban-wombat/flattables/tree/master/cmd/flattablesc) | generate Google FlatBuffers code

## Contact:

* _email:_ urban.wombat.burrow@gmail.com

* _Twitter:_ @UrbanWombat

## Why Use gotables?

1. Often enough the data you want to represent is intrinsically tabular, or should be.
2. You want tables of data to be more readable by human beings. Be able to look at the data and spot any problems.
3. You want to eliminate repetitive metadata such as tags, and reduce the size of each tabular chunk of data.
   Data name and type are mentioned only once in a gotables Table.
4. Sometimes the data you want to represent is intrinsically tabular, and really you don't want any elements to be missing.
   And if they are, you want it to be obvious.
5. It feels like overkill to set up a relational database table (or tables) to store (and modify) your software configurations,
   or to use a database as a conduit for sharing messages or data flows between processes or goroutines.
6. If you are sending messages between goroutines in Go, you can use a gotables Table or a set of Tables (a TableSet) to send
   data through your channels. A Table can be sent and received as an object or as a string.
7. gotables has methods and functions to perform tasks in these broad categories:

   1. Instantiate a gotables.Table or gotables.TableSet from a file, a string, a Go binary object, or a segment of an existing table.

   2. Get and Set values. Most Go types are supported.

   3. Sort and Search a table. Multiple keys and reverse sort and search are supported.

   4. Merge two tables (with shared key(s)) into one. NaN and zero values are handled consistently.

   5. SortUnique to remove NaN and zero values.

8. **gotables Table is _simple_**. For instance, sorting (and searching) a table is probably as easy as it can possibly be.
   And that can mean multiple sort/search keys, and even reverse keys. It's very simple. And if a wrong column name is
   used, or you forget to set sort keys before sorting, the gotables error handling system will notice and return to you
   a meaningful error message.

9. Some advantages are subtle. For instance, **versioning** is easier. Your program can test for the presence of particular
   columns (and their types) before accessing potentially new columns of data. And sending a table with additional columns
   will not break downstream code.

10. gotables is written in the Go language, using purely Go libraries. No third-party libraries will be used. If down the track
	non-Go libraries are needed, a separate distribution will be created (to contain any third-party dependency issues) and
	will itself use gotables as a third-party library to maintain separation.
	The core gotables library will **not* use third-party libraries.

## Go Doc for gotables

https://urban-wombat.github.io

## What is a gotables.Table?

A gotables.Table is a table of data with the following sections:
1. A table name in square brackets:   **[planets]**
2. A row of one or more column names: **name      mass**
3. A row of one or more column types: **string    float64**
4. Zero or more rows of data:         **"Mercury" 0.055**
5. One or more blank lines before any subsequent table(s)

It's a bit like a slice of struct.

Here's an example:

    [planets]
    name         mass distance moons index mnemonic
    string    float64  float64   int   int string
    "Mercury"   0.055      0.4     0     0 "my"
    "Venus"     0.815      0.7     0     1 "very"
    "Earth"     1.0        1.0     1     2 "elegant"
    "Mars"      0.107      1.5     2     3 "mother"
    "Jupiter" 318.0        5.2    79     4 "just"
    "Saturn"   95.0        9.5    62     5 "sat"
    "Uranus"   15.0       19.2    27     6 "upon"
    "Neptune"  17.0       30.6    13     7 "nine"
    "Pluto"     0.002     39.4     5     8 "porcupines"

Most of the Go builtin data types can be used (except complex64 and complex128).

Here is a simple program that parses the table into a gotables.Table and echoes it back out:

```
// main_echo.go

package main

import (
    "github.com/urban-wombat/gotables"
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
    "Earth"     1.0        1.0     1     2 "elegant"
    "Mars"      0.107      1.5     2     3 "mother"
    "Jupiter" 318.0        5.2    79     4 "just"
    "Saturn"   95.0        9.5    62     5 "sat"
    "Uranus"   15.0       19.2    27     6 "upon"
    "Neptune"  17.0       30.6    13     7 "nine"
    "Pluto"     0.002     39.4     5     8 "porcupines"
    `
    
    table, err := gotables.NewTableFromString(tableString)
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
```

Here's the output:

```
Default String() padded output

[planets]
name         mass distance moons index mnemonic
string    float64  float64   int   int string
"Mercury"   0.055      0.4     0     0 "my"
"Venus"     0.815      0.7     0     1 "very"
"Earth"     1.0        1.0     1     2 "elegant"
"Mars"      0.107      1.5     2     3 "mother"
"Jupiter" 318.0        5.2    79     4 "just"
"Saturn"   95.0        9.5    62     5 "sat"
"Uranus"   15.0       19.2    27     6 "upon"
"Neptune"  17.0       30.6    13     7 "nine"
"Pluto"     0.002     39.4     5     8 "porcupines"

For unpadded output use StringUnpadded()

[planets]
name mass distance moons index mnemonic
string float64 float64 int int string
"Mercury" 0.055 0.4 0 0 "my"
"Venus" 0.815 0.7 0 1 "very"
"Earth" 1 1 1 2 "elegant"
"Mars" 0.107 1.5 2 3 "mother"
"Jupiter" 318 5.2 79 4 "just"
"Saturn" 95 9.5 62 5 "sat"
"Uranus" 15 19.2 27 6 "upon"
"Neptune" 17 30.6 13 7 "nine"
"Pluto" 0.002 39.4 5 8 "porcupines"

```

## Can you show me some worked examples?

For these examples to compile and run for you, you need to go get and import "github.com/urban-wombat/gotables"
and prefix function and method calls with gotables.

```
// example1.go

package main

import (
    "fmt"
    "log"
    "github.com/urban-wombat/gotables"
)

// Copyright (c) 2017 Malcolm Gorman

func main() {
    tableString :=
    `[planets]
    name         mass distance moons index mnemonic
    string    float64  float64   int   int string
    "Mercury"   0.055      0.4     0     0 "my"
    "Venus"     0.815      0.7     0     1 "very"
    "Earth"     1.0        1.0     1     2 "elegant"
    "Mars"      0.107      1.5     2     3 "mother"
    "Jupiter" 318.0        5.2    79     4 "just"
    "Saturn"   95.0        9.5    62     5 "sat"
    "Uranus"   15.0       19.2    27     6 "upon"
    "Neptune"  17.0       30.6    13     7 "nine"
    "Pluto"     0.002     39.4     5     8 "porcupines"
    `

    var err error

    table, err := gotables.NewTableFromString(tableString)
    if err != nil {
    	log.Println(err)
    }
    fmt.Println("Table [planets] already in distance order.")
    fmt.Println("Distance is in AU: Earth units from the Sun")
    fmt.Println(table)

    var rowIndex int

    fmt.Println("Get the name and mass of the first planet.")
    rowIndex = 0
    fmt.Printf("rowIndex = %d\n", rowIndex)
    name, err := table.GetString("name", rowIndex)
    if err != nil {
    	log.Println(err)
    }
    fmt.Printf("name = %s\n", name)

    mass, err := table.GetFloat64("mass", rowIndex)
    if err != nil {
    	log.Println(err)
    }
    fmt.Printf("mass = %f\n", mass)
    fmt.Println()

    fmt.Println("Get and Set the mnemonic of the second planet: index 1")
    rowIndex = 1
    fmt.Printf("rowIndex = %d\n", rowIndex)
    name, err = table.GetString("name", rowIndex)
    if err != nil {
    	log.Println(err)
    }
    fmt.Printf("name = %s\n", name)

    mnemonic, err := table.GetString("mnemonic", rowIndex)
    if err != nil {
    	log.Println(err)
    }
    fmt.Printf("mnemonic = %q\n", mnemonic)

    err = table.SetString("mnemonic", rowIndex, "*VERY*")
    if err != nil {
    	log.Println(err)
    }

    mnemonic, err = table.GetString("mnemonic", rowIndex)
    if err != nil {
    	log.Println(err)
    }
    fmt.Printf("mnemonic = %q\n", mnemonic)
    fmt.Println()

    fmt.Println("Sort and Search.")
    sortKey := "name"
    fmt.Printf("sortKey = %q\n", sortKey)
    err = table.SetSortKeys(sortKey)
    if err != nil {
    	log.Println(err)
    }
    err = table.Sort()
    if err != nil {
    	log.Println(err)
    }

    planet := "Saturn"
    fmt.Printf("search value: planet = %q\n", planet)
    rowIndex, err = table.Search(planet)
    if err != nil {
    	log.Println(err)
    }

    moons, err := table.GetInt("moons", rowIndex)
    if err != nil {
    	log.Println(err)
    }

    fmt.Println(table)
    fmt.Printf("%s has %d moons.\n", planet, moons)
    fmt.Println()

    fmt.Println("Sort and Search Range.")
    err = table.SetSortKeys("moons")
    if err != nil {
    	log.Println(err)
    }

    err = table.Sort()
    if err != nil {
    	log.Println(err)
    }

    moons = 2
    firstRowIndex, lastRowIndex, err := table.SearchRange(moons)
    if err != nil {
    	log.Println(err)
    }
    var planets int
    if err == nil {
    	fmt.Println("Found at least 1 row with 2 moons.")
    	planets = lastRowIndex - firstRowIndex + 1
    } else {
    	// moons = 3: [planets].Search([3]) search values not in table: [3]
    	fmt.Println(err)
    	planets = 0
    }
    fmt.Println(table)
    fmt.Printf("%d planets have %d moons.\n", planets, moons)
    fmt.Println()


    // Sort Unique.

    tableString =
    `[Unique]
    key n       s
    int float32 string
    2   0       "two point two"
    2   2.2     ""
    1   1.1     "one point one"
    3   3.3     "three point three"
    3   3.3     ""
    3   NaN     "three point three"
    4   0.0     "neither zero nor same XX"
    4   NaN     "neither zero nor same YY"
    4   4.4     "neither zero nor same ZZ"
    4   NaN     "neither zero nor same AA"
    5   NaN     "minus 5"
    5   -0      "minus 5"
    5   -5      "minus 5"
    `
    table, err = gotables.NewTableFromString(tableString)
    if err != nil {
    	log.Println(err)
    }

    fmt.Println("Table [Unique] in no particular order, contains duplicate key values and zero and NaN values.")
    fmt.Println(table)

    sortKey = "key"
    fmt.Printf("sortKey = %q\n", sortKey)
    err = table.SetSortKeys(sortKey)
    if err != nil {
    	log.Println(err)
    }

    tableUnique, err := table.SortUnique()
    if err != nil {
    	log.Println(err)
    }
    fmt.Printf("table [%s] sorted unique by key %q\n", tableUnique.Name(), sortKey)
    fmt.Println(tableUnique)
}
```

Output:

```
Table [planets] already in distance order.
Distance is in AU: Earth units from the Sun
[planets]
name         mass distance moons index mnemonic
string    float64  float64   int   int string
"Mercury"   0.055      0.4     0     0 "my"
"Venus"     0.815      0.7     0     1 "very"
"Earth"     1.0        1.0     1     2 "elegant"
"Mars"      0.107      1.5     2     3 "mother"
"Jupiter" 318.0        5.2    79     4 "just"
"Saturn"   95.0        9.5    62     5 "sat"
"Uranus"   15.0       19.2    27     6 "upon"
"Neptune"  17.0       30.6    13     7 "nine"
"Pluto"     0.002     39.4     5     8 "porcupines"

Get the name and mass of the first planet.
rowIndex = 0
name = Mercury
mass = 0.055000

Get and Set the mnemonic of the second planet: index 1
rowIndex = 1
name = Venus
mnemonic = "very"
mnemonic = "*VERY*"

Sort and Search.
sortKey = "name"
search value: planet = "Saturn"
[planets]
name         mass distance moons index mnemonic
string    float64  float64   int   int string
"Earth"     1.0        1.0     1     2 "elegant"
"Jupiter" 318.0        5.2    79     4 "just"
"Mars"      0.107      1.5     2     3 "mother"
"Mercury"   0.055      0.4     0     0 "my"
"Neptune"  17.0       30.6    13     7 "nine"
"Pluto"     0.002     39.4     5     8 "porcupines"
"Saturn"   95.0        9.5    62     5 "sat"
"Uranus"   15.0       19.2    27     6 "upon"
"Venus"     0.815      0.7     0     1 "*VERY*"

Saturn has 62 moons.

Sort and Search Range.
Found at least 1 row with 2 moons.
[planets]
name         mass distance moons index mnemonic
string    float64  float64   int   int string
"Venus"     0.815      0.7     0     1 "*VERY*"
"Mercury"   0.055      0.4     0     0 "my"
"Earth"     1.0        1.0     1     2 "elegant"
"Mars"      0.107      1.5     2     3 "mother"
"Pluto"     0.002     39.4     5     8 "porcupines"
"Neptune"  17.0       30.6    13     7 "nine"
"Uranus"   15.0       19.2    27     6 "upon"
"Saturn"   95.0        9.5    62     5 "sat"
"Jupiter" 318.0        5.2    79     4 "just"

1 planets have 2 moons.

Table [Unique] in no particular order, contains duplicate key values and zero and NaN values.
[Unique]
key       n s
int float32 string
  2     0.0 "two point two"
  2     2.2 ""
  1     1.1 "one point one"
  3     3.3 "three point three"
  3     3.3 ""
  3     NaN "three point three"
  4     0.0 "neither zero nor same XX"
  4     NaN "neither zero nor same YY"
  4     4.4 "neither zero nor same ZZ"
  4     NaN "neither zero nor same AA"
  5     NaN "minus 5"
  5    -0.0 "minus 5"
  5    -5.0 "minus 5"

sortKey = "key"
table [Unique] sorted unique by key "key"
[Unique]
key       n s
int float32 string
  1     1.1 "one point one"
  2     2.2 "two point two"
  3     3.3 "three point three"
  4     4.4 "neither zero nor same AA"
  5    -5.0 "minus 5"
```
