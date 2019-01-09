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
	"path/filepath"
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/urban-wombat/gotables"
	"github.com/urban-wombat/util"
//	"strings"
	"time"
)

type Flags struct {
//	f string	// gotables file
	// See: https://stackoverflow.com/questions/35809252/check-if-flag-was-provided-in-go#
	// See: https://golang.org/pkg/flag
	f util.StringFlag	// gotables file
	t string	// table
	r string	// rotate this table in one direction or the other (if possible)
	pipe bool	// pipe input
	h bool		// help
}
var flags Flags

func init() {
	log.SetFlags(log.Lshortfile)
}
var where = log.Print

func init() {
	log.SetFlags(log.Lshortfile)

	flag.Usage = printUsage	// Override the default flag.Usage variable.
	initFlags()
}

func initFlags() {
//	flag.StringVar(&flags.f,  "f", "",    "tables file")
	flag.Var(&flags.f,        "f",        "tables file")	// flag.Var() defaults to initial value of variable.
	flag.StringVar(&flags.t,  "t", "",    "this table")
	flag.StringVar(&flags.r,  "r", "",    "rotate table")
	flag.BoolVar(&flags.pipe, "-", false, "piped input")
	flag.BoolVar(&flags.h,    "h", false, "print gotecho usage")

	flag.Parse()

/*
	// Compulsory flag.
	exists, err := util.CheckStringFlag("f", flags.f, util.FlagRequired)
	if !exists {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		fmt.Fprintf(os.Stderr, "Expecting infile: -f <gotables-file>\n")
		printUsage()
		os.Exit(111)
	}
*/
/*
	// Optional flag.
	_, err := util.CheckStringFlag("f", flags.f, util.FlagOptional)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		printUsage()
		os.Exit(111)
	}
	// If we get past here, -f has been provided and has an argument.
*/

	if flags.h {
		printUsage()
		os.Exit(2)
	}

/*
	if len(os.Args) == 1 {
		// No args.
		printUsage()
		os.Exit(3)
	}
*/
}

func progName() string {
//	return strings.TrimSuffix(filepath.Base(os.Args[0]), ".exe")
	return filepath.Base(os.Args[0])
}

func printUsage() {
	var usageSlice = []string {
		"usage:   gotecho [-f <file>] [-t <this-table-only>] [-r <rotate-table>]",
		"         If no -f <file> is specified, gotecho searches standard input.",
		"purpose: echo a file of gotables tables to stdout",
		"flags:   -f  <gotables-file>  Input file text file containing a gotables.TableSet",
		"         -t  this-table-only  Echo this table only",
		"         -r  rotate-table     Rotate this table (from tabular-to-struct or struct-to-tabular)",
		"                              Note: Rotate tabular-to-struct is ignored if table has multiple rows",
		"         -h                   Help",
	}

	var usageString string
	for i := 0; i < len(usageSlice); i++ {
		usageString += usageSlice[i] + "\n"
	}

/*
	var progNameEndsWithExe bool = strings.HasSuffix(progName(), ".exe")
	if progNameEndsWithExe {
		// We are testing. Provide a useful example. Does not appear in final product.
		usageString += "example: go run gotecho.go -f ../flattablesmain/mytables.got -r AllTypes"
	}
*/

	fmt.Fprintf(os.Stderr, "%s\n", usageString)
}

func main() {
	var err error
	var file string
	var tables *gotables.TableSet

/*
	if len(os.Args) <= 1 {
		// No file arguments provided.
		printUsage()
		os.Exit(4)
	}
*/

	// Clumsy way to avoid multiple files being specified with -f
	// Only possible because we are sure what the max possible args can be.
	const maxArgsPossible = 7
	if len(os.Args) > maxArgsPossible {
		// No file arguments provided.
		fmt.Fprintf(os.Stderr, "Too many arguments on command line %s\n", os.Args[1:])
		printUsage()
		os.Exit(5)
	}

/*
where(fmt.Sprintf("flags.f.Exists() = %t", flags.f.Exists()))
where(fmt.Sprintf("flags.f.IsSet()  = %t", flags.f.IsSet()))
where(fmt.Sprintf("flags.f.String() = %s", flags.f.String()))
where(fmt.Sprintf("flags.f.Error()  = %v", flags.f.Error()))
*/

	if flags.f.Error() != nil {
		fmt.Fprintf(os.Stderr, "ERROR: -f %v\n", flags.f.Error())
		os.Exit(16)
	}
	if flags.f.Exists() && flags.f.IsSet() {
		file = flags.f.String()
		tables, err = gotables.NewTableSetFromFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(6)
		}
	} else {	// Pipe from Stdin.
		canPipe, err := util.CanReadFromPipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(6)
		}
		if canPipe {
where("BEFORE util.GulpFromPipeWithTimeout()")
			input, err := util.GulpFromPipeWithTimeout(3 * time.Second)
where("AFTER  util.GulpFromPipeWithTimeout()")
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
				os.Exit(6)
			}
			tables, err = gotables.NewTableSetFromString(input)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
				os.Exit(6)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Cannot pipe to gotecho (on this computer). Use -f <file> instead.\n", err)
			printUsage()
			os.Exit(6)
		}
	}

	if tables.TableCount() == 0 {
		fmt.Fprintf(os.Stderr, "%s (warning: file empty)\n", file)
		os.Exit(7)
	}

	var finalMsg string
	if flags.r != "" {
		table, err := tables.Table(flags.r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error finding -r %s: %s\n", flags.r, err)
			os.Exit(8)
		}

		isStructShape, _ := table.IsStructShape()

		if isStructShape {
			// Rotate table from struct to tabular.
			table.SetStructShape(false)
		} else {	// is tabular
			// Print this table as a struct (if possible). If more than 1 row, must be printed as tabular.
			if table.RowCount() > 1 {
				finalMsg = fmt.Sprintf("Warning: -r %s: table [%s] with multiple %d rows cannot be rotated from tabular to struct shape",
					table.Name(), table.Name(), table.RowCount())
			} else {
				// Rotate table from tabular to struct.
				table.SetStructShape(true)
			}
		}

	}

	if flags.t != "" {
		// Print just this table.
		table, err := tables.Table(flags.t)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error finding -t %s: %s\n", flags.t, err)
			os.Exit(9)
		}
		fmt.Println(table)
	} else {
		// Print all tables.
		fmt.Println(tables)
	}

	if len(finalMsg) > 0 {
		fmt.Fprintf(os.Stderr, "%s.\n", finalMsg)
	}
}
