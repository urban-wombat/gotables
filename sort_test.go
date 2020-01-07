package gotables

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"testing"
)

func ExampleTable_Sort() {
	tableString :=
		`[planets]
	name         mass distance
	string    float64  float64
	"Mercury"   0.055      0.4
	"Venus"     0.815      0.7
	"Earth"     1.000      1.0
	"Mars"      0.107      1.5
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(1) Unsorted table:")
	fmt.Println(table)

	// First let's sort the table by name.
	err = table.SetSortKeys("name")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(2) Sorted table by name:")
	fmt.Println(table)

	// Now let's sort the table by name but this time in reverse.
	err = table.SetSortKeys("name")
	if err != nil {
		log.Println(err)
	}
	err = table.SetSortKeysReverse("name")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(3) Sorted table by name in reverse:")
	fmt.Println(table)

	// Output:
	// (1) Unsorted table:
	// [planets]
	// name         mass distance
	// string    float64  float64
	// "Mercury"   0.055      0.4
	// "Venus"     0.815      0.7
	// "Earth"     1.0        1.0
	// "Mars"      0.107      1.5
	//
	// (2) Sorted table by name:
	// [planets]
	// name         mass distance
	// string    float64  float64
	// "Earth"     1.0        1.0
	// "Mars"      0.107      1.5
	// "Mercury"   0.055      0.4
	// "Venus"     0.815      0.7
	//
	// (3) Sorted table by name in reverse:
	// [planets]
	// name         mass distance
	// string    float64  float64
	// "Venus"     0.815      0.7
	// "Mercury"   0.055      0.4
	// "Mars"      0.107      1.5
	// "Earth"     1.0        1.0
}

func ExampleTable_SetSortKeys() {
	tableString :=
		`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(1) Unsorted table:")
	fmt.Println(table)

	// Sort the table by user.
	if err = table.SetSortKeys("user"); err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(2) Sorted by user:")
	fmt.Println(table)

	// Sort by user and lines.
	err = table.SetSortKeys("user", "lines")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(3) Sorted by user and lines:")
	fmt.Println(table)

	// Sort the table by user but reverse lines.
	err = table.SetSortKeys("user", "lines")
	if err != nil {
		log.Println(err)
	}
	err = table.SetSortKeysReverse("lines")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(4) Sort by user but reverse lines:")
	fmt.Println(table)

	// Sort the table by language and lines.
	err = table.SetSortKeys("language", "lines")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(5) Sort by language and lines:")
	fmt.Println(table)

	// Sort the table by language and lines and user.
	err = table.SetSortKeys("language", "lines", "user")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(6) Sort by language and lines and user:")
	fmt.Println(table)

	keysTable, err := table.GetSortKeysAsTable()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(7) SortKeys as a table:")
	fmt.Println(keysTable)

	// Output:
	// (1) Unsorted table:
	// [changes]
	// user     language    lines
	// string   string        int
	// "gri"    "Go"          100
	// "ken"    "C"           150
	// "glenda" "Go"          200
	// "rsc"    "Go"          200
	// "r"      "Go"          100
	// "ken"    "Go"          200
	// "dmr"    "C"           100
	// "r"      "C"           150
	// "gri"    "Smalltalk"    80
	//
	// (2) Sorted by user:
	// [changes]
	// user     language    lines
	// string   string        int
	// "dmr"    "C"           100
	// "glenda" "Go"          200
	// "gri"    "Go"          100
	// "gri"    "Smalltalk"    80
	// "ken"    "C"           150
	// "ken"    "Go"          200
	// "r"      "Go"          100
	// "r"      "C"           150
	// "rsc"    "Go"          200
	//
	// (3) Sorted by user and lines:
	// [changes]
	// user     language    lines
	// string   string        int
	// "dmr"    "C"           100
	// "glenda" "Go"          200
	// "gri"    "Smalltalk"    80
	// "gri"    "Go"          100
	// "ken"    "C"           150
	// "ken"    "Go"          200
	// "r"      "Go"          100
	// "r"      "C"           150
	// "rsc"    "Go"          200
	//
	// (4) Sort by user but reverse lines:
	// [changes]
	// user     language    lines
	// string   string        int
	// "dmr"    "C"           100
	// "glenda" "Go"          200
	// "gri"    "Go"          100
	// "gri"    "Smalltalk"    80
	// "ken"    "Go"          200
	// "ken"    "C"           150
	// "r"      "C"           150
	// "r"      "Go"          100
	// "rsc"    "Go"          200
	//
	// (5) Sort by language and lines:
	// [changes]
	// user     language    lines
	// string   string        int
	// "dmr"    "C"           100
	// "ken"    "C"           150
	// "r"      "C"           150
	// "r"      "Go"          100
	// "gri"    "Go"          100
	// "ken"    "Go"          200
	// "glenda" "Go"          200
	// "rsc"    "Go"          200
	// "gri"    "Smalltalk"    80
	//
	// (6) Sort by language and lines and user:
	// [changes]
	// user     language    lines
	// string   string        int
	// "dmr"    "C"           100
	// "ken"    "C"           150
	// "r"      "C"           150
	// "gri"    "Go"          100
	// "r"      "Go"          100
	// "glenda" "Go"          200
	// "ken"    "Go"          200
	// "rsc"    "Go"          200
	// "gri"    "Smalltalk"    80
	//
	// (7) SortKeys as a table:
	// [SortKeys]
	// index colName    colType  reverse
	//   int string     string   bool
	//     0 "language" "string" false
	//     1 "lines"    "int"    false
	//     2 "user"     "string" false
}

func TestTable_SortZeroKeys(t *testing.T) {

	table, err := NewTable("HasZeroSortKeys")
	if err != nil {
		t.Fatal(err)
	}

	err = table.Sort()
	if err == nil {
		t.Fatalf("Expecting table.Sort() err because of 0 sort keys")
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

/*
	Test Sort() in two modes:-
		(1) Sort by columns specified in argument list.
		(2) Sort by columns specified in table sort keys.
*/
func TestTable_Sort(t *testing.T) {

	unsortedString := `
	[table]
	i	b		s
	int	bool	string
	1	true	"Z"
	2	false	"Y"
	3	true	"X"
	4	false	"W"
	5	true	"V"
	6	false	"U"
	`
	table, err := NewTableFromString(unsortedString)
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Printf("%v\n", table)

	sortedString := `
	[table]
	  i b     s
	int bool  string
	  6 false "U"
	  4 false "W"
	  2 false "Y"
	  5 true  "V"
	  3 true  "X"
	  1 true  "Z"
	`
	tableSorted, err := NewTableFromString(sortedString)
	if err != nil {
		t.Fatal(err)
	}

	// (1) Sort by columns specified in argument list.

	err = table.Sort("b", "s")
	if err != nil {
		t.Fatal(err)
	}

	// fmt.Printf("%v\n", table)
	if equals, err := table.Equals(tableSorted); !equals {
		t.Fatal(err)
	}


	// Set sort keys to zero before the next test case.

	err = table.SetSortKeys()
	if err != nil {
		t.Fatal(err)
	}

	if table.SortKeyCount() != 0 {
		t.Fatalf("expecting SortKeyCount() to be zero because we just now cleared the keys")
	}


	// (2) Sort by columns specified in table sort keys.

	err = table.SetSortKeys("b", "s")
	if err != nil {
		t.Fatal(err)
	}

	err = table.Sort()
	if err != nil {
		t.Fatal(err)
	}

	// fmt.Printf("%v\n", table)
	if equals, err := table.Equals(tableSorted); !equals {
		t.Fatal(err)
	}
}

func TestTable_SortSimple(t *testing.T) {

	table, err := NewTable("HasZeroSortKeys")
	if err != nil {
		t.Fatal(err)
	}

	err = table.Sort()
	if err == nil {
		t.Fatalf("Expecting table.Sort() err because of 0 sort keys")
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func TestSearch(t *testing.T) {

	tableString :=
		`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = table.Search()
	if err == nil {
		t.Fatalf("Expecting table.Search() err because of 0 sort keys")
	}

	// Clear sort keys (if any) by calling with empty argument list.
	err = table.SetSortKeys() // Note: sort keys count 0
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetSortKeys("user") // Note: sort keys count 1
	if err != nil {
		t.Fatal(err)
	}

	err = table.Sort()
	if err != nil {
		t.Fatal(err)
	}

	_, err = table.Search() // Note: 0 search values passed to Search()
	if err == nil {
		t.Fatalf("Expecting searchValues count 0 != sort keys count 1")
	}

	_, err = table.Search("glenda")
	if err != nil {
		t.Fatal(err)
	}

	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func ExampleTable_Search_keys1() {
	// mass:     Earth = 1 (relative to Earth)
	// distance: Earth = 1 (relative to Earth - AU)
	// http://www.windows2universe.org/our_solar_system/planets_table.html
	// http://www.space.com/17001-how-big-is-the-sun-size-of-the-sun.html
	tableString :=
		`[planets]
	name         mass distance moons index mnemonic
	string    float64   float64   int   int string
	"Sun"      333333        0     0    -1 ""
	"Mercury"   0.055      0.4     0     0 "my"
	"Venus"     0.815      0.7     0     1 "very"
	"Earth"     1.000      1.0     1     2 "elegant"
	"Mars"      0.107      1.5     2     3 "mother"
	"Jupiter" 318.000      5.2    79     4 "just"
	"Saturn"   95.000      9.5    82     5 "sat"
	"Uranus"   15.000     19.2    27     6 "upon"
	"Neptune"  17.000     30.6    13     7 "nine"
	"Pluto"     0.002     39.4     5     8 "porcupines"
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(1) Unsorted table:")
	fmt.Println(table)

	// First let's sort the table by name.
	err = table.SetSortKeys("name")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(2) Sorted table by name:")
	fmt.Println(table)

	searchValue := "Mars" // 2
	fmt.Printf("(3) Search for name: %s\n", searchValue)
	rowIndex, err := table.Search(searchValue)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Found %s at rowIndex = %d\n", searchValue, rowIndex)
	fmt.Println()

	searchValue = "Ceres" // -1
	fmt.Printf("(4) Search for name: %s\n", searchValue)
	rowIndex, _ = table.Search(searchValue)
	fmt.Printf("Found %s at rowIndex = %d (missing)\n", searchValue, rowIndex)

	// Output:
	// (1) Unsorted table:
	// [planets]
	// name            mass distance moons index mnemonic
	// string       float64  float64   int   int string
	// "Sun"     333333.0        0.0     0    -1 ""
	// "Mercury"      0.055      0.4     0     0 "my"
	// "Venus"        0.815      0.7     0     1 "very"
	// "Earth"        1.0        1.0     1     2 "elegant"
	// "Mars"         0.107      1.5     2     3 "mother"
	// "Jupiter"    318.0        5.2    79     4 "just"
	// "Saturn"      95.0        9.5    82     5 "sat"
	// "Uranus"      15.0       19.2    27     6 "upon"
	// "Neptune"     17.0       30.6    13     7 "nine"
	// "Pluto"        0.002     39.4     5     8 "porcupines"
	//
	// (2) Sorted table by name:
	// [planets]
	// name            mass distance moons index mnemonic
	// string       float64  float64   int   int string
	// "Earth"        1.0        1.0     1     2 "elegant"
	// "Jupiter"    318.0        5.2    79     4 "just"
	// "Mars"         0.107      1.5     2     3 "mother"
	// "Mercury"      0.055      0.4     0     0 "my"
	// "Neptune"     17.0       30.6    13     7 "nine"
	// "Pluto"        0.002     39.4     5     8 "porcupines"
	// "Saturn"      95.0        9.5    82     5 "sat"
	// "Sun"     333333.0        0.0     0    -1 ""
	// "Uranus"      15.0       19.2    27     6 "upon"
	// "Venus"        0.815      0.7     0     1 "very"
	//
	// (3) Search for name: Mars
	// Found Mars at rowIndex = 2
	//
	// (4) Search for name: Ceres
	// Found Ceres at rowIndex = -1 (missing)
}

func ExampleTable_Search_keys1Reverse() {
	// mass:     Earth = 1 (relative to Earth)
	// distance: Earth = 1 (relative to Earth - AU)
	// http://www.windows2universe.org/our_solar_system/planets_table.html
	tableString :=
		`[planets]
	name         mass distance moons index mnemonic
	string    float64  float64   int   int string
	"Mercury"   0.055      0.4     0     0 "my"
	"Venus"     0.815      0.7     0     1 "very"
	"Earth"     1.000      1.0     1     2 "elegant"
	"Mars"      0.107      1.5     2     3 "mother"
	"Jupiter" 318.000      5.2    79     4 "just"
	"Saturn"   95.000      9.5    82     5 "sat"
	"Uranus"   15.000     19.2    27     6 "upon"
	"Neptune"  17.000     30.6    13     7 "nine"
	"Pluto"     0.002     39.4     5     8 "porcupines"
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(1) Unsorted table:")
	fmt.Println(table)

	// First let's sort the table by name in reverse.
	err = table.SetSortKeys("name")
	if err != nil {
		log.Println(err)
	}
	err = table.SetSortKeysReverse("name")
	if err != nil {
		log.Println(err)
	}
	err = table.Sort()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("(2) Sorted table by name in reverse order:")
	fmt.Println(table)

	searchValue := "Mars" // 5
	fmt.Printf("(3) Search for name: %s\n", searchValue)
	rowIndex, err := table.Search(searchValue)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Found %s at rowIndex = %d\n", searchValue, rowIndex)

	searchValue = "Ceres" // -1
	fmt.Printf("(4) Search for name: %s\n", searchValue)
	rowIndex, _ = table.Search(searchValue)
	fmt.Printf("Found %s at rowIndex = %d (missing)\n", searchValue, rowIndex)

	// Output:
	// (1) Unsorted table:
	// [planets]
	// name         mass distance moons index mnemonic
	// string    float64  float64   int   int string
	// "Mercury"   0.055      0.4     0     0 "my"
	// "Venus"     0.815      0.7     0     1 "very"
	// "Earth"     1.0        1.0     1     2 "elegant"
	// "Mars"      0.107      1.5     2     3 "mother"
	// "Jupiter" 318.0        5.2    79     4 "just"
	// "Saturn"   95.0        9.5    82     5 "sat"
	// "Uranus"   15.0       19.2    27     6 "upon"
	// "Neptune"  17.0       30.6    13     7 "nine"
	// "Pluto"     0.002     39.4     5     8 "porcupines"
	//
	// (2) Sorted table by name in reverse order:
	// [planets]
	// name         mass distance moons index mnemonic
	// string    float64  float64   int   int string
	// "Venus"     0.815      0.7     0     1 "very"
	// "Uranus"   15.0       19.2    27     6 "upon"
	// "Saturn"   95.0        9.5    82     5 "sat"
	// "Pluto"     0.002     39.4     5     8 "porcupines"
	// "Neptune"  17.0       30.6    13     7 "nine"
	// "Mercury"   0.055      0.4     0     0 "my"
	// "Mars"      0.107      1.5     2     3 "mother"
	// "Jupiter" 318.0        5.2    79     4 "just"
	// "Earth"     1.0        1.0     1     2 "elegant"
	//
	// (3) Search for name: Mars
	// Found Mars at rowIndex = 6
	// (4) Search for name: Ceres
	// Found Ceres at rowIndex = -1 (missing)
}

func TestTable_Search_1key(t *testing.T) {
	// mass:     Earth = 1 (relative to Earth)
	// distance: Earth = 1 (relative to Earth - AU)
	// http://www.windows2universe.org/our_solar_system/planets_table.html
	tableString :=
		`[planets]
	name         mass distance moons index mnemonic
	string    float64  float64   int   int string
	"Mercury"   0.055      0.4     0     0 "my"
	"Venus"     0.815      0.7     0     1 "very"
	"Earth"     1.000      1.0     1     2 "elegant"
	"Mars"      0.107      1.5     2     3 "mother"
	"Jupiter" 318.000      5.2    79     4 "just"
	"Saturn"   95.000      9.5    82     5 "sat"
	"Uranus"   15.000     19.2    27     6 "upon"
	"Neptune"  17.000     30.6    13     7 "nine"
	"Pluto"     0.002     39.4     5     8 "porcupines"
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	// First let's sort the table by name.
	err = table.SetSortKeys("name")
	if err != nil {
		t.Fatal(err)
	}
	err = table.Sort()
	if err != nil {
		t.Fatal(err)
	}

	var searchValue string
	var expecting int
	var rowIndex int

	// Search for entries that exist in the table.
	for i := 0; i < table.RowCount(); i++ {
		searchValue, err = table.GetString("name", i)
		if err != nil {
			t.Fatal(err)
		}
		expecting = i
		rowIndex, err = table.Search(searchValue)
		if err != nil {
			t.Fatal(err)
		}
		if rowIndex != expecting {
			t.Fatalf("Expecting Search(%q) = %d but found: %d", searchValue, expecting, rowIndex)
		}
	}

	// Search for entries that don't exist.
	dontExist := []string{
		"Sun",
		"Moon",
		"Ceres",
	}
	for _, item := range dontExist {
		searchValue = item
		expecting = -1
		rowIndex, err = table.Search(searchValue)
		if err == nil {
			t.Fatalf("Expecting an error with Search(%v)", searchValue)
		}
		if rowIndex != expecting {
			t.Fatalf("Expecting Search(%q) = %d but found: %d", searchValue, expecting, rowIndex)
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func TestTable_Search_1key_reverse(t *testing.T) {
	// mass:     Earth = 1 (relative to Earth)
	// distance: Earth = 1 (relative to Earth - AU)
	// http://www.windows2universe.org/our_solar_system/planets_table.html
	tableString :=
		`[planets]
	name         mass distance moons index mnemonic
	string    float64  float64   int   int string
	"Mercury"   0.055      0.4     0     0 "my"
	"Venus"     0.815      0.7     0     1 "very"
	"Earth"     1.000      1.0     1     2 "elegant"
	"Mars"      0.107      1.5     2     3 "mother"
	"Jupiter" 318.000      5.2    79     4 "just"
	"Saturn"   95.000      9.5    82     5 "sat"
	"Uranus"   15.000     19.2    27     6 "upon"
	"Neptune"  17.000     30.6    13     7 "nine"
	"Pluto"     0.002     39.4     5     8 "porcupines"
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	// First let's sort the table by name - in reverse order.
	err = table.SetSortKeys("name")
	if err != nil {
		t.Fatal(err)
	}
	err = table.SetSortKeysReverse("name")
	if err != nil {
		t.Fatal(err)
	}
	err = table.Sort()
	if err != nil {
		t.Fatal(err)
	}

	var searchValue string
	var expecting int
	var rowIndex int

	// Search for entries that exist in the table.
	for i := 0; i < table.RowCount(); i++ {
		searchValue, err = table.GetString("name", i)
		if err != nil {
			t.Fatal(err)
		}
		expecting = i

		rowIndex, err = table.Search(searchValue)
		if err != nil {
			t.Fatal(err)
		}

		if rowIndex != expecting {
			t.Fatalf("Expecting Search(%q) = %d but found: %d", searchValue, expecting, rowIndex)
		}
	}
	//	log.Printf("%q expecting %d found %d", searchValue, expecting, rowIndex)

	// Search for entries that don't exist.
	dontExist := []string{
		"Sun",
		"Moon",
		"Ceres",
	}
	for _, item := range dontExist {
		searchValue = item
		expecting = -1

		rowIndex, err = table.Search(searchValue)
		if err == nil {
			t.Fatalf("Expecting an error with Search(%v)", searchValue)
		}
		if rowIndex != expecting {
			t.Fatalf("Expecting Search(%q) = %d but found: %d", searchValue, expecting, rowIndex)
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func TestTable_Search_2keys(t *testing.T) {
	tableString :=
		`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	// First let's sort the table by user and lines.
	err = table.SetSortKeys("user", "lines")
	if err != nil {
		t.Fatal(err)
	}
	err = table.Sort()
	if err != nil {
		t.Fatal(err)
	}

	var searchValues []interface{} = make([]interface{}, 2)
	var expecting int
	var found int

	// Search for entries that exist in the table.
	for i := 0; i < table.RowCount(); i++ {
		searchValues[0], err = table.GetString("user", i)
		if err != nil {
			t.Fatal(err)
		}
		searchValues[1], err = table.GetInt("lines", i)
		if err != nil {
			t.Fatal(err)
		}
		expecting = i
		found, err = table.Search(searchValues...)
		if err != nil {
			t.Fatal(err)
		}
		if found != expecting {
			t.Fatalf("Expecting Search(%v) = %d but found: %d", searchValues, expecting, found)
		}
	}
	//	log.Printf("%q expecting %d found %d", searchValues, expecting, found)

	// Search for entries that don't exist.
	dontExist := [][]interface{}{
		{"steve", 42},
		{"bill", 42},
		{"larry", 42},
	}
	for _, item := range dontExist {
		searchValues = item
		expecting = -1
		found, _ = table.Search(searchValues...)
		if found != expecting {
			t.Fatalf("Expecting Search(%q) = %d but found: %d", searchValues, expecting, found)
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func TestTable_Search_2keys_reverse2nd(t *testing.T) {
	tableString :=
		`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	// First let's sort the table by user and lines.
	err = table.SetSortKeys("user", "lines")
	if err != nil {
		t.Fatal(err)
	}
	err = table.SetSortKeysReverse("lines")
	if err != nil {
		t.Fatal(err)
	}
	err = table.Sort()
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Printf("here:\n%s", table)

	var searchValues []interface{} = make([]interface{}, 2)
	var expecting int
	var found int

	// Search for entries that exist in the table.
	for i := 0; i < table.RowCount(); i++ {
		searchValues[0], err = table.GetString("user", i)
		if err != nil {
			t.Fatal(err)
		}
		searchValues[1], err = table.GetInt("lines", i)
		if err != nil {
			t.Fatal(err)
		}
		expecting = i
		found, err = table.Search(searchValues...)
		if err != nil {
			t.Fatal(err)
		}
		if found != expecting {
			t.Fatalf("Expecting Search(%v) = %d but found: %d", searchValues, expecting, found)
		}
	}
	//	log.Printf("%q expecting %d found %d", searchValues, expecting, found)

	// Search for entries that don't exist.
	dontExist := [][]interface{}{
		{"steve", 42},
		{"bill", 42},
		{"larry", 42},
	}
	for _, item := range dontExist {
		searchValues = item
		expecting = -1
		found, _ = table.Search(searchValues...)
		if found != expecting {
			t.Fatalf("Expecting Search(%q) = %d but found: %d", searchValues, expecting, found)
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func TestTable_Search_2keys_reverseBoth(t *testing.T) {
	tableString :=
		`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetSortKeys("user", "lines")
	if err != nil {
		t.Fatal(err)
	}
	err = table.SetSortKeysReverse("user", "lines")
	if err != nil {
		t.Fatal(err)
	}
	err = table.Sort()
	if err != nil {
		t.Fatal(err)
	}

	var searchValues []interface{} = make([]interface{}, 2)
	var expecting int
	var found int

	// Search for entries that exist in the table.
	for i := 0; i < table.RowCount(); i++ {
		searchValues[0], err = table.GetString("user", i)
		if err != nil {
			t.Fatal(err)
		}
		searchValues[1], err = table.GetInt("lines", i)
		if err != nil {
			t.Fatal(err)
		}
		expecting = i
		found, err = table.Search(searchValues...)
		if err != nil {
			t.Fatal(err)
		}
		if found != expecting {
			t.Fatalf("Expecting Search(%v) = %d but found: %d", searchValues, expecting, found)
		}
	}

	// Search for entries that don't exist.
	dontExist := [][]interface{}{
		{"steve", 42},
		{"bill", 42},
		{"larry", 42},
	}
	for _, item := range dontExist {
		searchValues = item
		expecting = -1
		found, err = table.Search(searchValues...)
		if err == nil {
			t.Fatalf("Expecting an error with Search(%v)", searchValues)
		}
		if found != expecting {
			t.Fatalf("Expecting Search(%q) = %d but found: %d", searchValues, expecting, found)
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func ExampleTable_GetSortKeysAsTable() {
	tableString :=
		`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	// Sort the table by user but reverse lines.
	err = table.SetSortKeys("user", "lines")
	if err != nil {
		log.Println(err)
	}

	err = table.SetSortKeysReverse("lines")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("(1) GetSortKeysAsTable():")
	sortKeysTable, err := table.GetSortKeysAsTable()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(sortKeysTable)

	err = table.Sort()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("(2) Sort by user but reverse lines:")
	fmt.Println(table)

	// Output:
	// (1) GetSortKeysAsTable():
	// [SortKeys]
	// index colName colType  reverse
	//   int string  string   bool
	//     0 "user"  "string" false
	//     1 "lines" "int"    true
	//
	// (2) Sort by user but reverse lines:
	// [changes]
	// user     language    lines
	// string   string        int
	// "dmr"    "C"           100
	// "glenda" "Go"          200
	// "gri"    "Go"          100
	// "gri"    "Smalltalk"    80
	// "ken"    "Go"          200
	// "ken"    "C"           150
	// "r"      "C"           150
	// "r"      "Go"          100
	// "rsc"    "Go"          200
}

func TestTable_SortKeyCount(t *testing.T) {
	tableString :=
		`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	// First let's sort the table by user and lines.
	err = table.SetSortKeys("user", "lines")
	if err != nil {
		t.Fatal(err)
	}

	expecting := 2
	count := table.SortKeyCount()
	if count != expecting {
		t.Fatalf("Expecting table.SortKeyCount() = %d but found %d", expecting, count)
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func TestTable_SetSortKeysFromTable(t *testing.T) {
	fromTableString :=
		`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	fromTable, err := NewTableFromString(fromTableString)
	if err != nil {
		t.Fatal(err)
	}

	// First let's sort the table by user and lines.
	err = fromTable.SetSortKeys("user", "lines")
	if err != nil {
		t.Fatal(err)
	}

	err = fromTable.SetSortKeysReverse("lines")
	if err != nil {
		t.Fatal(err)
	}

	toTableString :=
		`[ToTable]
	user	string
	lines	int
	`
	toTable, err := NewTableFromString(toTableString)
	if err != nil {
		t.Fatal(err)
	}

	err = toTable.SetSortKeysFromTable(fromTable)
	if err != nil {
		t.Fatal(err)
	}

	keysTable1, err := fromTable.GetSortKeysAsTable()
	if err != nil {
		t.Fatal(err)
	}

	keysTable2, err := toTable.GetSortKeysAsTable()
	if err != nil {
		t.Fatal(err)
	}

	expecting := true

	equals, err := keysTable1.Equals(keysTable2)
	if err != nil {
		t.Fatal(err)
	}

	if equals != expecting {
		t.Fatalf("Expecting table1.Equals(table2) = %t but found %t", expecting, equals)
	}
	if isValid, err := keysTable1.IsValidTable(); !isValid {
		t.Fatal(err)
	}
	if isValid, err := keysTable2.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func ExampleTable_OrderColsBySortKeys() {
	tableString :=
		`[MyTable]
	ColA   ColB Key2      ColC Key1 ColD ColE
	string  int string float64  int  int bool
	`

	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	err = table.SetSortKeys("Key1", "Key2")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("(1) Here is the table in its original column order:")
	fmt.Println(table)

	fmt.Println("(2) Here are the keys:")
	sortKeysTable, err := table.GetSortKeysAsTable()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(sortKeysTable)

	fmt.Println("(3) Order the sort key columns to the left:")
	err = table.OrderColsBySortKeys()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(table)

	// Output:
	// (1) Here is the table in its original column order:
	// [MyTable]
	// ColA   ColB Key2      ColC Key1 ColD ColE
	// string  int string float64  int  int bool
	//
	// (2) Here are the keys:
	// [SortKeys]
	// index colName colType  reverse
	//   int string  string   bool
	//     0 "Key1"  "int"    false
	//     1 "Key2"  "string" false
	//
	// (3) Order the sort key columns to the left:
	// [MyTable]
	// Key1 Key2   ColA   ColB    ColC ColD ColE
	//  int string string  int float64  int bool
}

/*
	This tests a copy gotables.Search() of sort.Search()
	to confirm that SearchLast() is a mirror image in
	behaviour: Search() is GE and SearchLast is LE.
*/
func Test_Search(t *testing.T) {

	/*
		sliceToString := func(slice []int) string {
			var s string
			for i := 0; i < len(slice); i++ {
				s += fmt.Sprintf("%3d", slice[i])
			}
			return s
		}
	*/

	const tests = 40 // Make this 20 for realism.
	const elements = 10
	const intRange = 10
	slice := make([]int, elements)
	indices := make([]int, elements)

	for i := 0; i < elements; i++ {
		indices[i] = i
	}

	//	rand.Seed(time.Now().UnixNano())

	for i := 0; i < tests; i++ {
		for j := 0; j < elements; j++ {
			slice[j] = rand.Intn(intRange)
		}
		sort.Ints(slice)
		// fmt.Println()
		// fmt.Printf("%s\n", util.FuncName())
		// fmt.Printf("test[%2d] %s\n", i, sliceToString(slice))
		// fmt.Printf("test[%2d] %s\n", i, sliceToString(indices))
		var index int
		for searchFor := -1; searchFor <= intRange; searchFor++ {
			index = sort.Search(elements, func(element int) bool {
				return slice[element] >= searchFor
			})

			// fmt.Printf("index for %d is %2d\n", searchFor, index)

			if index >= elements {
				// fmt.Printf("%d is missing but would be at (nonexistent) index %d (insert before %d)\n", searchFor, index, index)
			} else {
				if slice[index] != searchFor {
					// Have we found at the very least A right element, or if it is missing, an element less than it.
					if slice[index] < searchFor {
						t.Fatal(fmt.Sprintf("test[%d] Expecting Search() slice[%d] = %d or more than %d, but found %d",
							i, index, searchFor, searchFor, slice[index]))
					} else {
						// fmt.Printf("%d is missing but would be at index %d (insert before %d)\n", searchFor, index, index)
					}
				}
			}

			if index > 0 && slice[index-1] == searchFor {
				// Have we found THE right element.
				t.Fatal(fmt.Sprintf("test[%d] Expecting Search() slice[%d] = %d to be lowest index, but found slice[%d-1] = %d lower",
					i, index, searchFor, index, slice[index-1]))
			}
		}
	}
}

// LE: Less than or equal.
func TestSearchLast(t *testing.T) {

	/*
		// Inner function to convert a slice to a string.
		sliceToString := func(slice []int) string {
			var s string
			for i := 0; i < len(slice); i++ {
				s += fmt.Sprintf("%3d", slice[i])
			}
			return s
		}
	*/

	const tests = 40 // Make this 20 for realism.
	const elements = 10
	const intRange = 10
	slice := make([]int, elements)
	indices := make([]int, elements)

	for i := 0; i < elements; i++ {
		indices[i] = i
	}

	//	rand.Seed(time.Now().UnixNano())

	for i := 0; i < tests; i++ {
		for j := 0; j < elements; j++ {
			slice[j] = rand.Intn(intRange)
		}
		sort.Ints(slice)
		// fmt.Println()
		// fmt.Printf("%s\n", util.FuncName())
		// fmt.Printf("test[%2d] %s\n", i, sliceToString(slice))
		// fmt.Printf("test[%2d] %s\n", i, sliceToString(indices))
		var index int
		for searchFor := -1; searchFor <= intRange; searchFor++ {
			index = SearchLast(elements, func(element int) bool {
				return slice[element] <= searchFor
			})

			// fmt.Printf("index for %d is %2d\n", searchFor, index)

			if index < 0 {
				// fmt.Printf("%d is missing but would be at (nonexistent) index %d (insert after %d)\n", searchFor, index, index)
			} else {
				if slice[index] != searchFor {
					// Have we found at the very least A right element, or if it is missing, an element less than it.
					if slice[index] > searchFor {
						t.Fatal(fmt.Sprintf("test[%d] Expecting SearchLast() slice[%d] = %d or less than %d, but found %d",
							i, index, searchFor, searchFor, slice[index]))
					} else {
						// fmt.Printf("%d is missing but would be at index %d (insert after %d)\n", searchFor, index, index)
					}
				}
			}

			if index < elements-1 && slice[index+1] == searchFor {
				// Have we found THE right element.
				t.Fatal(fmt.Sprintf("test[%d] Expecting SearchLast() slice[%d] = %d to be greatest index, but found slice[%d+1] = %d greater",
					i, index, searchFor, index, slice[index+1]))
			}
		}
	}
}

func ExampleTable_SearchLast() {

	var data []int = []int{4, 8, 10, 10, 10, 20, 23, 29}
	fmt.Printf("data: %v\n", data)
	fmt.Println("index: 0 1  2  3  4  5  6  7")
	fmt.Println()

	fmt.Printf("(1) Find an element that is present:\n")
	x := 23
	fmt.Printf("Searching for x: %d\n", x)
	i := SearchLast(len(data), func(i int) bool { return data[i] <= x })
	fmt.Printf("x %d is, or would be, at index i: %d\n", x, i)

	// Check whether x is actually where SearchLast() said it is, or would be inserted.
	if i >= 0 && data[i] == x {
		fmt.Printf("x %d is present at data[%d]\n", x, i)
	} else {
		fmt.Printf("x is not present in data, but i %d is the index where it would be inserted AFTER.\n", i)
		fmt.Printf("Note that i can be -1 which does not exist in data.\n")
	}
	fmt.Println()

	fmt.Printf("(2) This time find an x that is present multiple times:\n")
	x = 10
	fmt.Printf("Searching for x: %d\n", x)
	i = SearchLast(len(data), func(i int) bool { return data[i] <= x })
	fmt.Printf("x %d is, or would be, at index i: %d\n", x, i)

	// Check whether x is actually where SearchLast() said it is, or would be inserted.
	if i >= 0 && data[i] == x {
		fmt.Printf("x %d is present at data[%d]\n", x, i)
	} else {
		fmt.Printf("x is not present in data, but i %d is the index where it would be inserted AFTER.\n", i)
		fmt.Printf("Note that i can be -1 which does not exist in data.\n")
	}
	fmt.Println()

	fmt.Printf("(3) This time find an x that is missing between items in data:\n")
	x = 15
	fmt.Printf("Searching for x: %d\n", x)
	i = SearchLast(len(data), func(i int) bool { return data[i] <= x })
	fmt.Printf("x %d is, or would be, at index i: %d\n", x, i)

	// Check whether x is actually where SearchLast() said it is, or would be inserted.
	if i >= 0 && data[i] == x {
		fmt.Printf("x %d is present at data[%d]\n", x, i)
	} else {
		fmt.Printf("x is not present in data, but i %d is the index where it would be inserted AFTER.\n", i)
		fmt.Printf("Note that i can be -1 which does not exist in data.\n")
	}
	fmt.Println()

	fmt.Printf("(4) This time find an x that is missing below all items in data:\n")
	x = 3
	fmt.Printf("Searching for x: %d\n", x)
	i = SearchLast(len(data), func(i int) bool { return data[i] <= x })
	fmt.Printf("x %d is, or would be, at index i: %d\n", x, i)

	// Check whether x is actually where SearchLast() said it is, or would be inserted.
	if i >= 0 && data[i] == x {
		fmt.Printf("x %d is present at data[%d]\n", x, i)
	} else {
		fmt.Printf("x is not present in data, but i %d is the index where it would be inserted AFTER.\n", i)
		fmt.Printf("Note that i can be -1 which does not exist in data.\n")
	}
	fmt.Println()

	fmt.Printf("(5) This time find an x that is missing above all items in data:\n")
	x = 31
	fmt.Printf("Searching for x: %d\n", x)
	i = SearchLast(len(data), func(i int) bool { return data[i] <= x })
	fmt.Printf("x %d is, or would be, at index i: %d\n", x, i)

	// Check whether x is actually where SearchLast() said it is, or would be inserted.
	if i >= 0 && data[i] == x {
		fmt.Printf("x %d is present at data[%d]\n", x, i)
	} else {
		fmt.Printf("x is not present in data, but i %d is the index where it would be inserted AFTER.\n", i)
		fmt.Printf("Note that i can be -1 which does not exist in data.\n")
	}
	fmt.Println()

	// Output:
	// data: [4 8 10 10 10 20 23 29]
	// index: 0 1  2  3  4  5  6  7
	//
	// (1) Find an element that is present:
	// Searching for x: 23
	// x 23 is, or would be, at index i: 6
	// x 23 is present at data[6]
	//
	// (2) This time find an x that is present multiple times:
	// Searching for x: 10
	// x 10 is, or would be, at index i: 4
	// x 10 is present at data[4]
	//
	// (3) This time find an x that is missing between items in data:
	// Searching for x: 15
	// x 15 is, or would be, at index i: 4
	// x is not present in data, but i 4 is the index where it would be inserted AFTER.
	// Note that i can be -1 which does not exist in data.
	//
	// (4) This time find an x that is missing below all items in data:
	// Searching for x: 3
	// x 3 is, or would be, at index i: -1
	// x is not present in data, but i -1 is the index where it would be inserted AFTER.
	// Note that i can be -1 which does not exist in data.
	//
	// (5) This time find an x that is missing above all items in data:
	// Searching for x: 31
	// x 31 is, or would be, at index i: 7
	// x is not present in data, but i 7 is the index where it would be inserted AFTER.
	// Note that i can be -1 which does not exist in data.
}

func TestTable_SearchFirst_by_user(t *testing.T) {
	tableString :=
		`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetSortKeys("user")
	if err != nil {
		t.Fatal(err)
	}
	err = table.Sort()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		searchValue string
		expecting   int
	}{
		{"dmr", 0},
		{"glenda", 1},
		{"gri", 2},
		{"ken", 4},
		{"r", 6},
		{"rsc", 8},
		{"NOT", -1},
	}

	for _, test := range tests {
		found, err := table.SearchFirst(test.searchValue)
		if found != test.expecting {
			t.Fatalf("Expecting SearchFirst(%q) = %d but found: %d, err: %v", test.searchValue, test.expecting, found, err)
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func TestTable_SearchLast_by_user(t *testing.T) {
	tableString :=
		`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetSortKeys("user")
	if err != nil {
		t.Fatal(err)
	}
	err = table.Sort()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		searchValue string
		expecting   int
	}{
		{"dmr", 0},
		{"glenda", 1},
		{"gri", 3},
		{"ken", 5},
		{"r", 7},
		{"rsc", 8},
		{"NOT", -1},
	}

	for _, test := range tests {
		found, err := table.SearchLast(test.searchValue)
		if found != test.expecting {
			t.Fatalf("Expecting SearchLast(%q) = %d but found: %d, err: %v", test.searchValue, test.expecting, found, err)
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func TestTable_SearchRange_by_user(t *testing.T) {
	tableString :=
		`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetSortKeys("user")
	if err != nil {
		t.Fatal(err)
	}
	err = table.Sort()
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		searchValue    string
		expectingFirst int
		expectingLast  int
	}{
		{"dmr", 0, 0},
		{"glenda", 1, 1},
		{"gri", 2, 3},
		{"ken", 4, 5},
		{"r", 6, 7},
		{"rsc", 8, 8},
		{"NOT", -1, -1},
	}

	for _, test := range tests {
		foundFirst, foundLast, err := table.SearchRange(test.searchValue)
		if foundFirst != test.expectingFirst || foundLast != test.expectingLast {
			t.Fatalf("Expecting SearchRange(%q) = %d, %d but found: %d, %d err: %v",
				test.searchValue, test.expectingFirst, test.expectingLast, foundFirst, foundLast, err)
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func TestTable_SearchRange_by_user_lines(t *testing.T) {
	tableString :=
		`[changes]
	user     language    lines index
	string   string        int   int
	"rsc"    "Go"          200     0
	"r"      "Go"          100     0
	"r"      "C"           150     0
	"ken"    "C"           150     0
	"ken"    "Go"          200     0
	"ken"    "Go"          200     0
	"gri"    "Smalltalk"    80     0
	"gri"    "Go"          100     0
	"gri"    "Go"          100     0
	"gri"    "Go"          100     0
	"glenda" "Go"          200     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetSortKeys("user", "lines")
	if err != nil {
		t.Fatal(err)
	}
	err = table.Sort()
	if err != nil {
		t.Fatal(err)
	}

	// To eye-ball errors.
	for i := 0; i < table.RowCount(); i++ {
		err = table.SetInt("index", i, i)
		if err != nil {
			t.Fatal(err)
		}
	}

	var tests = []struct {
		searchName     string
		searchLines    int
		expectingFirst int
		expectingLast  int
	}{
		{"dmr", 100, 0, 4},
		{"glenda", 200, 5, 5},
		{"gri", 100, 7, 9},
		{"ken", 200, 11, 12},
		{"r", 150, 14, 14},
		{"rsc", 200, 15, 15},
		{"NOT", 500, -1, -1},
		{"NOT", 200, -1, -1},
		{"rsc", 100, -1, -1},
	}

	for _, test := range tests {
		foundFirst, foundLast, err := table.SearchRange(test.searchName, test.searchLines)
		if foundFirst != test.expectingFirst || foundLast != test.expectingLast {
			t.Fatalf("Expecting SearchRange(%q, %d) = %d, %d but found: %d, %d err: %v",
				test.searchName, test.searchLines, test.expectingFirst, test.expectingLast, foundFirst, foundLast, err)
			fmt.Println(table)
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func TestTable_SearchRange_by_user_lines_reverse_lines(t *testing.T) {
	tableString :=
		`[changes]
	user     language    lines index
	string   string        int   int
	"rsc"    "Go"          200     0
	"r"      "Go"          100     0
	"r"      "C"           150     0
	"ken"    "C"           150     0
	"ken"    "Go"          200     0
	"ken"    "Go"          200     0
	"gri"    "Smalltalk"    80     0
	"gri"    "Go"          100     0
	"gri"    "Go"          100     0
	"gri"    "Go"          100     0
	"glenda" "Go"          200     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	"dmr"    "C"           100     0
	`
	table, err := NewTableFromString(tableString)
	if err != nil {
		t.Fatal(err)
	}

	err = table.SetSortKeys("user", "lines")
	if err != nil {
		t.Fatal(err)
	}
	err = table.SetSortKeysReverse("lines")
	if err != nil {
		t.Fatal(err)
	}
	err = table.Sort()
	if err != nil {
		t.Fatal(err)
	}

	// To eye-ball errors.
	for i := 0; i < table.RowCount(); i++ {
		err = table.SetInt("index", i, i)
		if err != nil {
			t.Fatal(err)
		}
	}

	var tests = []struct {
		searchName     string
		searchLines    int
		expectingFirst int
		expectingLast  int
	}{
		{"dmr", 100, 0, 4},
		{"glenda", 200, 5, 5},
		{"gri", 100, 6, 8},
		{"ken", 200, 10, 11},
		{"r", 150, 13, 13},
		{"rsc", 200, 15, 15},
		{"NOT", 500, -1, -1},
	}

	for _, test := range tests {
		foundFirst, foundLast, err := table.SearchRange(test.searchName, test.searchLines)
		if foundFirst != test.expectingFirst || foundLast != test.expectingLast {
			t.Fatalf("Expecting SearchRange(%q, %d) = %d, %d but found: %d, %d err: %v",
				test.searchName, test.searchLines, test.expectingFirst, test.expectingLast, foundFirst, foundLast, err)
			fmt.Println(table)
		}
	}
	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
}

func ExampleTable_SortUnique() {

	tableString :=
		`[Uniqueness]
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
	table, err := NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Before SortUnique() ...")
	fmt.Println(table)

	err = table.SetSortKeys("KeyCol")
	if err != nil {
		log.Println(err)
	}

	tableUnique, err := table.SortUnique()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("After SortUnique() ...")
	fmt.Println(tableUnique)

	// Output:
	// Before SortUnique() ...
	// [Uniqueness]
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
	// After SortUnique() ...
	// [Uniqueness]
	// KeyCol  number s
	//    int float32 string
	//      1     1.1 "one point one"
	//      2     2.2 "two point two"
	//      3     3.3 "three point three"
	//      4     4.4 "neither zero nor same A"
	//      5    -5.0 "minus 5"
}

func ExampleTable_SortSimple() {

	var tableString string
	var table *Table
	var err error

	tableString =
		`[planets]
	name         mass distance moons index mnemonic
	string    float64   float64   int   int string
	"Earth"     1.000      1.0     1     2 "elegant"
	"Jupiter" 318.000      5.2    79     4 "just"
	"Mars"      0.107      1.5     2     3 "mother"
	"Mercury"   0.055      0.4     0     0 "my"
	"Neptune"  17.000     30.6    13     7 "nine"
	"Pluto"     0.002     39.4     5     8 "porcupines"
	"Saturn"   95.000      9.5    82     5 "sat"
	"Sun"      333333        0     0    -1 ""
	"Uranus"   15.000     19.2    27     6 "upon"
	"Venus"     0.815      0.7     0     1 "very"
	`

	table, err = NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	// Single column sort. Sort the planets in order from the Sun.
	err = table.SortSimple("distance")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	tableString =
		`[changes]
	user     language    lines
	string   string        int
	"gri"    "Go"          100
	"ken"    "C"           150
	"glenda" "Go"          200
	"rsc"    "Go"          200
	"r"      "Go"          100
	"ken"    "Go"          200
	"dmr"    "C"           100
	"r"      "C"           150
	"gri"    "Smalltalk"    80
	`

	table, err = NewTableFromString(tableString)
	if err != nil {
		log.Println(err)
	}

	// Multiple column sort. Sort users by lines, language and user name.
	err = table.SortSimple("lines", "language", "user")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(table)

	// Output:
	// [planets]
	// name            mass distance moons index mnemonic
	// string       float64  float64   int   int string
	// "Sun"     333333.0        0.0     0    -1 ""
	// "Mercury"      0.055      0.4     0     0 "my"
	// "Venus"        0.815      0.7     0     1 "very"
	// "Earth"        1.0        1.0     1     2 "elegant"
	// "Mars"         0.107      1.5     2     3 "mother"
	// "Jupiter"    318.0        5.2    79     4 "just"
	// "Saturn"      95.0        9.5    82     5 "sat"
	// "Uranus"      15.0       19.2    27     6 "upon"
	// "Neptune"     17.0       30.6    13     7 "nine"
	// "Pluto"        0.002     39.4     5     8 "porcupines"
	//
	// [changes]
	// user     language    lines
	// string   string        int
	// "gri"    "Smalltalk"    80
	// "dmr"    "C"           100
	// "gri"    "Go"          100
	// "r"      "Go"          100
	// "ken"    "C"           150
	// "r"      "C"           150
	// "glenda" "Go"          200
	// "ken"    "Go"          200
	// "rsc"    "Go"          200
}
