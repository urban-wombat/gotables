package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
//	"os"
	"os/exec"
//	"path/filepath"
//	"runtime/debug"
	"strings"
	"syscall"
	"testing"
//	"time"

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
	if  !installed {
		t.Fatal(err)
	}

	const tables_got = "gotecho_files/tables.got"
	goArgs := []string{"run", "gotecho.go"}

	var tests = []struct {
		desc string
		stdinPipe bool			// Will "cat" (via stdin) file to gotecho.
		fileOfExpected string	// Name of file containing expected data.
		exitCodeExpected int
		args string
	}{
	//  Description						Piped?	Expected								Exit?	Arguments
	/*
		{ "echo all",                    false, "gotecho_files/tables.got",               0, "-f "+tables_got+"" },
		{ "only Table",                  false, "gotecho_files/Table.got",                0, "-f "+tables_got+" -t Table" },
		{ "only Struct",                 false, "gotecho_files/Struct.got",               0, "-f "+tables_got+" -t Struct" },
		{ "missing table",               false, "gotecho_files/empty.got",                1, "-f "+tables_got+" -t MissingTable" },
		{ "rotate struct",               false, "gotecho_files/rotateStructHasTable.got", 0, "-f "+tables_got+" -r Struct" },
		{ "rotate struct no table",      false, "gotecho_files/rotateStructNoTable.got",  0, "-f "+tables_got+" -r Struct -t Struct" },
		{ "ignores rotate table",        false, "gotecho_files/Table.got",                0, "-f "+tables_got+" -r Table -t Table" },
		{ "pipe echo all",               true,  "gotecho_files/tables.got",               0, "" },
		{ "pipe only Table",             true,  "gotecho_files/Table.got",                0, "-t Table" },
		{ "pipe only Struct",            true,  "gotecho_files/Struct.got",               0, "-t Struct" },
		{ "pipe missing table",          true,  "gotecho_files/empty.got",                1, "-t MissingTable" },
		{ "pipe rotate struct",          true,  "gotecho_files/rotateStructHasTable.got", 0, "-r Struct" },
		{ "pipe rotate struct no table", true,  "gotecho_files/rotateStructNoTable.got",  0, "-r Struct -t Struct" },
		{ "pipe ignores rotate table",   true,  "gotecho_files/Table.got",                0, "-r Table -t Table" },
		{ "echo all - missing filename", false, "gotecho_files/empty.got",                1, "-f" },
	*/
		{ "echo all - missing -f",       false, "gotecho_files/empty.got",                1, ""+tables_got+"" },
	}

	var cmd *exec.Cmd
	for i, test := range tests {
		const verbose = true
		if verbose {
			// fmt.Println("---------------------------------------------------")
			fmt.Printf("test[%d] %s\n", i, test.desc)
		}
		contents, err := ioutil.ReadFile(test.fileOfExpected)
		if err != nil { t.Error(err) }
		expected := string(contents)

		args := goArgs
		slicedArgs := strings.Split(test.args, " ")
		args = append(args, slicedArgs...)
		cmd = exec.Command("go", args...)

		var stdoutByteBuffer bytes.Buffer
		var stderrByteBuffer bytes.Buffer
		cmd.Stdout = &stdoutByteBuffer
		cmd.Stderr = &stderrByteBuffer

		var stdin io.WriteCloser
		if test.stdinPipe {
			stdin, err = cmd.StdinPipe()
			if err != nil { t.Error(err) }

			tablesBytes, err := ioutil.ReadFile(tables_got)
			if err != nil { t.Error(err) }
			tables_txt := string(tablesBytes)
			go func() {
				defer stdin.Close()
				io.WriteString(stdin, tables_txt)
			}()
		}

		err = cmd.Run()

		var exitCode int
		exitError, hasError := err.(*exec.ExitError)
		if hasError {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
			//where(fmt.Sprintf("exitError = %v", exitError))
			//where(fmt.Sprintf("hasError = %v", hasError))
			//where(fmt.Sprintf("exitCode = %v", exitCode))
		}

		stdout := stdoutByteBuffer.String()
		stderr := stderrByteBuffer.String()

/*
where(fmt.Sprintf("expected =\n%s", expected))
where(fmt.Sprintf("stdout = \n%s", stdout))
where(fmt.Sprintf("stderr = \n%s", stderr))
*/

		if stdout != expected {
			t.Errorf("test[%d] %s: ( gotecho %s ) OUTPUT != EXPECTED:-\nOUTPUT:\n%s\nEXPECTED:\n%s\nSTDERR:\n%s",
				i, test.desc, test.args, stdout, expected, stderr)
		}

		if exitCode != test.exitCodeExpected {
			t.Errorf("test[%d] %s: ( gotecho %s ) exitCode %d != exitCodeExpected %d\nSTDERR:\n%s",
				i, test.desc, test.args, exitCode, test.exitCodeExpected, stderr)
		}
	}
}
