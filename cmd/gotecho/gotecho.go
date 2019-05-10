// Echo the tables of gotables files.
package main

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

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urban-wombat/gotables"
	"github.com/urban-wombat/util"
)

const pipeTimeout = 3 // seconds

type Flags struct {
	// See: https://stackoverflow.com/questions/35809252/check-if-flag-was-provided-in-go
	// See: https://golang.org/pkg/flag
	t util.StringFlag // table name
	r util.StringFlag // rotate this table in one direction or the other (if possible)
	h bool            // help
}

var flags Flags

func init() {
	log.SetFlags(log.Lshortfile)
}

var where = log.Print

func init() {
	log.SetFlags(log.Lshortfile)

	flag.Usage = printUsage // Override the default flag.Usage variable.
	initFlags()
}

func initFlags() {
	flag.Var(&flags.t, "t", "this table")   // flag.Var() defaults to initial value of variable.
	flag.Var(&flags.r, "r", "rotate table") // flag.Var() defaults to initial value of variable.
	flag.BoolVar(&flags.h, "h", false, "print gotecho usage")

	flag.Parse()

	if flags.h {
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	var usageSlice = []string{
		"usage 1  gotecho [-t <this-table-only>] [-r <rotate-table>] <gotables-file>",
		"usage 2  cat <gotables-file> | gotecho [-t <this-table-only>] [-r <rotate-table>]",
		"         If no <gotables-file> is specified, gotecho searches standard input for " +
			fmt.Sprintf("%d", pipeTimeout) + " seconds",
		"purpose: echo a file of gotables tables to stdout",
		"         -t  this-table-only  Echo this table only",
		"         -r  rotate-table     Rotate this table (from tabular-to-struct or struct-to-tabular)",
		"                              Note: Rotate tabular-to-struct is ignored if table has multiple rows",
		"         -h                   Help",
	}

	var usageString string
	for i := 0; i < len(usageSlice); i++ {
		usageString += usageSlice[i] + "\n"
	}

	fmt.Fprintf(os.Stderr, "%s\n", usageString)
}

func main() {
	var err error
	var file string
	var tables *gotables.TableSet

	// Clumsy way to avoid multiple files being specified
	// Only possible because we are sure what the max possible args can be.
	const maxArgsPossible = 7
	if len(os.Args) > maxArgsPossible {
		// No file arguments provided.
		fmt.Fprintf(os.Stderr, "Too many arguments on command line %s\n", os.Args[1:])
		printUsage()
		os.Exit(2)
	}

	if len(os.Args) > 2 && os.Args[1] != "-t" && os.Args[1] != "-r" {
		fmt.Fprintf(os.Stderr, "expecting flags to come before <gotables-file>\n")
		printUsage()
		os.Exit(3)
	}

	var nonFlagArgs []string = nonFlagArgs(flag.Args())
	if len(nonFlagArgs) >= 1 {
		file = nonFlagArgs[0]
		tables, err = gotables.NewTableSetFromFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(4)
		}
	} else { // Pipe from Stdin.
		canPipe, err := util.CanReadFromPipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(5)
		}
		if canPipe {
			stdin, err := util.GulpFromPipeWithTimeout(pipeTimeout * time.Second)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %s %v\n", util.ProgName(), err)
				printUsage()
				os.Exit(6)
			}
			tables, err = gotables.NewTableSetFromString(stdin)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
				os.Exit(7)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Cannot pipe to gotecho (on this computer). Use <gotables-file> instead: %v\n", err)
			printUsage()
			os.Exit(8)
		}
	}

	if tables.TableCount() == 0 {
		fmt.Fprintf(os.Stderr, "%s (warning: gotables file is empty)\n", file)
		os.Exit(9)
	}

	var finalMsg string
	if flags.r.AllOk() {
		if flags.t.AllOk() && flags.t.String() != flags.r.String() {
			finalMsg = fmt.Sprintf("warning: ignoring gotecho -r %s which is a different table from gotecho -t %s\n",
				flags.r.String(), flags.t.String())
		}
		table, err := tables.Table(flags.r.String())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error finding -r %s: %s\n", flags.r.String(), err)
			os.Exit(10)
		}

		isStructShape, _ := table.IsStructShape()

		if isStructShape {
			// Rotate table from struct to tabular.
			table.SetStructShape(false)
		} else { // is tabular
			// Print this table as a struct (if possible). If more than 1 row, must be printed as tabular.
			if table.RowCount() > 1 {
				finalMsg = fmt.Sprintf("warning: gotecho -r %s: table [%s] with multiple %d rows cannot be rotated from tabular to struct shape",
					table.Name(), table.Name(), table.RowCount())
			} else {
				// Rotate table from tabular to struct.
				table.SetStructShape(true)
			}
		}
	}

	if flags.t.AllOk() {
		// Print just this table.
		table, err := tables.Table(flags.t.String())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error finding -t %s: %s\n", flags.t.String(), err)
			os.Exit(11)
		}
		fmt.Println(table)
	} else {
		// Print all tables.
		fmt.Println(tables)
	}

	if len(finalMsg) > 0 {
		fmt.Fprintf(os.Stderr, "%s\n", finalMsg)
	}
}

func nonFlagArgs(nfa []string) (nonFlagArgs []string) {
	for i := 0; i < len(nfa); i++ {
		if len(nfa[i]) > 0 {
			nonFlagArgs = append(nonFlagArgs, nfa[i])
		} else { // Not sure why this could happen.
			// fmt.Printf("nfa[%d] = \"\"\n", i)
		}
	}
	return
}
