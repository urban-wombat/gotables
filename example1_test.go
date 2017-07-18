package gotable

import (
//	"bytes"
	"fmt"
//	"io/ioutil"
	"log"
//	"math"
//	"math/rand"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//	"testing"
//	"time"
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

	For these examples to compile and run for you, you need to import "github.com/urban-wombat/gotable"
	and prefix function and method calls with gotable.
*/

func Example() {
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

	var err error

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(table)

	var rowIndex int = 0

	// Get the name and mass of the first planet.
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

	// Get and Set the mnemonic of the second planet.
	rowIndex = 1
	name, err = table.GetString("name", rowIndex)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("name = %s\n", name)

	mnemonic, err := table.GetString("mnemonic", rowIndex)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("mnemonic = %s\n", mnemonic)

	err = table.SetString("mnemonic", rowIndex, "VERY")
	if err != nil {
		log.Println(err)
	}

	mnemonic, err = table.GetString("mnemonic", rowIndex)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("mnemonic = %s\n", mnemonic)
	fmt.Println()

	// Sort and Search.
	err = table.SetSortKeys("name")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}

	planet := "Mars"
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

	// Sort and Search Range.
	err = table.SetSortKeys("moons")
	if err != nil {
		log.Println(err)
	}

	err = table.Sort()
	if err != nil {
		log.Println(err)
	}

	moons = 0
	firstRowIndex, lastRowIndex, err := table.SearchRange(moons)
	if err != nil {
		log.Println(err)
	}

	planets := lastRowIndex - firstRowIndex + 1
	fmt.Println(table)
	fmt.Printf("%d planets have %d moons.\n", planets, moons)
	fmt.Println()


	// Sort Unique.

	tableString =
	`[Unique]
	KeyCol number   s
	int float32 string
	2   0       "two point two"
	2   2.2     ""
	1   1.1     "one point one"
	3   3.3     "three point three"
	3   3.3     ""
	3   NaN     "three point three"
	4   0.0     "neither zero nor same X"
	4   NaN     "neither zero nor same Y"
	4   4.4     "neither zero nor same Z"
	4   NaN     "neither zero nor same A"
	5   NaN     "minus 5"
	5   -0      "minus 5"
	5   -5      "minus 5"
	`
	table, err = NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	err = table.SetSortKeys("KeyCol")
	if err != nil {
		log.Println(err)
	}

	tableUnique, err := table.SortUnique()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(tableUnique)

	// Output:
	// [planets]
	// name         mass distance moons index mnemonic
	// string    float64  float64   int   int string
	// "Mercury"   0.055      0.4     0     0 "my"
	// "Venus"     0.815      0.7     0     1 "very"
	// "Earth"     1.0        1.0     1     2 "elegant"
	// "Mars"      0.107      1.5     2     3 "mother"
	// "Jupiter" 318.0        5.2    67     4 "just"
	// "Saturn"   95.0       29.4    62     5 "sat"
	// "Uranus"   15.0       84.0    27     6 "upon"
	// "Neptune"  17.0      164.0    13     7 "nine ... porcupines"
	// 
	// rowIndex = 0
	// name = Mercury
	// mass = 0.055000
	// 
	// name = Venus
	// mnemonic = very
	// mnemonic = VERY
	// 
	// [planets]
	// name         mass distance moons index mnemonic
	// string    float64  float64   int   int string
	// "Earth"     1.0        1.0     1     2 "elegant"
	// "Jupiter" 318.0        5.2    67     4 "just"
	// "Mars"      0.107      1.5     2     3 "mother"
	// "Mercury"   0.055      0.4     0     0 "my"
	// "Neptune"  17.0      164.0    13     7 "nine ... porcupines"
	// "Saturn"   95.0       29.4    62     5 "sat"
	// "Uranus"   15.0       84.0    27     6 "upon"
	// "Venus"     0.815      0.7     0     1 "VERY"
	// 
	// Mars has 2 moons.
	// 
	// [planets]
	// name         mass distance moons index mnemonic
	// string    float64  float64   int   int string
	// "Venus"     0.815      0.7     0     1 "VERY"
	// "Mercury"   0.055      0.4     0     0 "my"
	// "Earth"     1.0        1.0     1     2 "elegant"
	// "Mars"      0.107      1.5     2     3 "mother"
	// "Neptune"  17.0      164.0    13     7 "nine ... porcupines"
	// "Uranus"   15.0       84.0    27     6 "upon"
	// "Saturn"   95.0       29.4    62     5 "sat"
	// "Jupiter" 318.0        5.2    67     4 "just"
	// 
	// 2 planets have 0 moons.
	// 
	// [Unique]
	// KeyCol  number s
	//    int float32 string
	//      2     0.0 "two point two"
	//      2     2.2 ""
	//      1     1.1 "one point one"
	//      3     3.3 "three point three"
	//      3     3.3 ""
	//      3     NaN "three point three"
	//      4     0.0 "neither zero nor same X"
	//      4     NaN "neither zero nor same Y"
	//      4     4.4 "neither zero nor same Z"
	//      4     NaN "neither zero nor same A"
	//      5     NaN "minus 5"
	//      5    -0.0 "minus 5"
	//      5    -5.0 "minus 5"
	// 
	// [Unique]
	// KeyCol  number s
	//    int float32 string
	//      1     1.1 "one point one"
	//      2     2.2 "two point two"
	//      3     3.3 "three point three"
	//      4     4.4 "neither zero nor same A"
	//      5    -5.0 "minus 5"
}
