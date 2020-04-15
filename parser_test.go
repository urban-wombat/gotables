package gotables_test

// Note: This is a black box test (different package name: not gotables).

import (
	_ "fmt"
	"math/rand"
	"testing"

	_ "github.com/urban-wombat/gotables"
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

func TestNewTableFromString_random(t *testing.T) {
/*
	var err error
	var table *gotables.Table
	var setupName string = "Fred"

	var tests3 = []struct {
		input    string
		succeeds bool
		output   string
	}{
		{"Barney", true, "Barney"},
		{"", false, "Fred"},
		{"$&*", false, "Fred"},
	}

	const testCount int = 20
	const maxCols int = 5
	const maxRows int = 5

	for testIndex := 0; testIndex < testCount; testIndex++ {

		table, err = gotables.NewTable(setupName)
		if err != nil {
			t.Fatal(err)
		}

		for colIndex := 0; colIndex < rand.Intn(10); colIndex++ {
		}


	}

	if isValid, err := table.IsValidTable(); !isValid {
		t.Fatal(err)
	}
*/

	min := 10
	max := 30
//	fmt.Println(rand.Intn(max - min) + min)
	rand.Intn((max - min) + min)
}
