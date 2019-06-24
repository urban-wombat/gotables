package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
	"testing"

	"github.com/urban-wombat/util"
)

/*
Copyright (c) 2018 Malcolm Gorman

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

func TestCmdGotecho(t *testing.T) {

	installed, err := util.IsCommandInstalled("go")
	if !installed {
		t.Fatal(err)
	}

	const tables_got = "test_files/tables.got"
	goArgs := []string{"run", "gotecho.go"}

	var tests = []struct {
		desc             string
		stdinPipe        bool   // Will "cat" (via stdin) file to gotecho.
		fileOfExpected   string // Name of file containing expected data.
		exitCodeExpected int    // 0 means success expected, 1 means failure expected.
		args             string // The arguments passed to "go run gotables.go".
	}{
		//  Description						Piped?	Expected						ExitVal?	Arguments
		/*
		 */
		{"echo all", false, "test_files/tables.got", 0, tables_got},
		{"only Table", false, "test_files/Table.got", 0, "-t Table  " + tables_got},
		{"only Struct", false, "test_files/Struct.got", 0, "-t Struct " + tables_got},
		{"missing table", false, "test_files/empty.got", 1, "-t MissingTable " + tables_got},
		{"rotate struct", false, "test_files/rotateStructHasTable.got", 0, "-r Struct " + tables_got},
		{"rotate struct no table", false, "test_files/rotateStructNoTable.got", 0, "-r Struct -t Struct " + tables_got},
		{"ignores rotate table", false, "test_files/Table.got", 0, "-r Table -t Table " + tables_got},
		{"pipe echo all", true, "test_files/tables.got", 0, ""},
		{"pipe only Table", true, "test_files/Table.got", 0, "-t Table"},
		{"pipe only Struct", true, "test_files/Struct.got", 0, "-t Struct"},
		{"pipe missing table", true, "test_files/empty.got", 1, "-t MissingTable"},
		{"pipe rotate struct", true, "test_files/rotateStructHasTable.got", 0, "-r Struct"},
		{"pipe rotate struct no table", true, "test_files/rotateStructNoTable.got", 0, "-r Struct -t Struct"},
		{"pipe ignores rotate table", true, "test_files/Table.got", 0, "-r Table -t Table"},
		{"echo missing <file>", false, "test_files/empty.got", 1, ""},
		{"echo -r -t missing <file>", false, "test_files/empty.got", 1, "-r Table -t Table"},
		{"echo -r missing <file>", false, "test_files/empty.got", 1, "-r Table"},
		{"echo -t missing <file>", false, "test_files/empty.got", 1, "-t Table"},
		{"<nil>| echo missing <file>", true, "test_files/empty.got", 1, ""}, // echo "" | gotecho
		/*
		 */
	}

	whiteSpace := regexp.MustCompile(`\s+`)

	var cmd *exec.Cmd
	fmt.Printf("%-8s %-30s %s\n", "Test#", "Description", "Command")
	for i, test := range tests {
		const verbose = true
		if verbose {
			// fmt.Println("---------------------------------------------------")
			fmt.Printf("test[%2d] %-28s   go run gotecho.go %s\n", i, test.desc, test.args)
		}
		contents, err := ioutil.ReadFile(test.fileOfExpected)
		if err != nil {
			t.Error(err)
		}
		expected := string(contents)

		args := goArgs
		slicedArgs := whiteSpace.Split(test.args, -1)
		args = append(args, slicedArgs...)
		cmd = exec.Command("go", args...)

		var stdoutByteBuffer bytes.Buffer
		var stderrByteBuffer bytes.Buffer
		cmd.Stdout = &stdoutByteBuffer
		cmd.Stderr = &stderrByteBuffer

		var stdin io.WriteCloser
		if test.stdinPipe {
			stdin, err = cmd.StdinPipe()
			if err != nil {
				t.Error(err)
			}

			var tablesTxt string
			if strings.HasPrefix(test.desc, "<nil>|") {
				tablesTxt = ""
			} else {
				tablesBytes, err := ioutil.ReadFile(tables_got)
				if err != nil {
					t.Error(err)
				}
				tablesTxt = string(tablesBytes)
			}
			go func() {
				defer stdin.Close()
				io.WriteString(stdin, tablesTxt)
			}()
		}

		err = cmd.Run()

		var exitCode int
		exitError, hasError := err.(*exec.ExitError)
		if hasError {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
			/*
				where(fmt.Sprintf("exitError = %v", exitError))
				where(fmt.Sprintf("hasError = %v", hasError))
				where(fmt.Sprintf("exitCode = %v", exitCode))
			*/
		}

		stdout := stdoutByteBuffer.String()
		stderr := stderrByteBuffer.String()

		/*
			where(fmt.Sprintf("expected =\n%s", expected))
			where(fmt.Sprintf("stdout = \n%s", stdout))
			where(fmt.Sprintf("stderr = \n%s", stderr))
		*/

		if stdout != expected {
			t.Errorf("test[%2d] %s: %s OUTPUT != EXPECTED:-\nOUTPUT:\n%s\nEXPECTED:\n%s\nSTDERR:\n%s",
				i, test.desc, cmd.Args, stdout, expected, stderr)
		}

		if exitCode != test.exitCodeExpected {
			t.Errorf("test[%2d] %s: %s exitCode %d != exitCodeExpected %d\nSTDERR:\n%s",
				i, test.desc, cmd.Args, exitCode, test.exitCodeExpected, stderr)
		}
	}
}
