// Check the syntax of gotables files.
package main

import (
	"fmt"
	"os"
	"github.com/urban-wombat/gotables"
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

func main() {
	var err error
	var fileName string
	var tables *gotables.TableSet
	var exitVal int = 0

	if len(os.Args) <= 1 {
		// No fileName arguments provided.
		printUsage()
		exitVal = 4
		os.Exit(exitVal)
	}

/*	Let's remember how to do this.
	var file *os.File
	var fileInfo os.FileInfo
	var fileMode os.FileMode
	var isRegular bool
*/
	for i := 1; i < len(os.Args); i++ {
		fileName = os.Args[i]
/*		Let's remember how to do this.
		file, err = os.Open(fileName)
		fileInfo, err = file.Stat()
		fileMode = fileInfo.Mode()
		isRegular = fileMode.IsRegular()
*/
		tables, err = gotables.NewTableSetFromFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error    %v\n", err)
			exitVal = 3
		} else {
			tableCount := tables.TableCount()
			if tableCount == 1 {
				table0, err := tables.TableByTableIndex(0)
				if err != nil {
					fmt.Fprintf(os.Stderr, "SYSTEM ERROR: %v\n", err)
					exitVal = 2
				}
				rowCount := table0.RowCount()
				fmt.Printf("ok       %s (%d table%s with %d row%s)\n",
					fileName, tableCount, plural(tableCount), rowCount, plural(rowCount))
			} else {
				fmt.Printf("ok       %s      (%d table%s)\n", fileName, tableCount, plural(tableCount))
			}
			if tableCount == 0 {
				fmt.Fprintf(os.Stderr, "warning  %s (WARNING: file has 0 tables)\n", fileName)
				exitVal = 1
			}
		}
	}

	os.Exit(exitVal)
}

// TODO: List exitVal meanings.
func printUsage() {
	fmt.Fprintf(os.Stderr, "usage: gotsyntax <gotables-files>\n")
}

func plural(items int) string {
	if items == 1 || items == -1 {
		// Singular
		return ""
	} else {
		// Plural
		return "s"
	}
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
