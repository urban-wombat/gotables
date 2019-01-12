package gotables

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

/*
	Utility functions that don't fit in any particular package and may be needed by any package.

	These are not considered part of the `gotables` interface surface and may change at any time.
*/

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

func init() {
	log.SetFlags(log.Lshortfile) // For var where
}
var where = log.Print

type StringFlag struct {
	set bool
	val string
	exists bool
	err error
}

func (sf *StringFlag) Set(s string) error {
	sf.exists = true

	if strings.HasPrefix(s, "-") {
		sf.val = s
		sf.err = fmt.Errorf("flag needs a valid string argument, not %s", sf.val)
	} else {
		sf.set = true
		sf.val = s
	}

	return nil
}

func (sf *StringFlag) string() string {
	return sf.val
}

func (sf *StringFlag) isSet() bool {
	return sf.set
}

func (sf *StringFlag) Exists() bool {
	return sf.exists
}

func (sf *StringFlag) error() error {
	return sf.err
}

/*
	Utility function to test string flags.

	It avoids boilerplate code testing flags.

	It can be called and:-

	(1) Required flags can trust the existence of an argument.

		// Required flag.
		exists, err := util.CheckStringFlag("r", flags.r, util.FlagRequired)
		if !exists {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

	(2) Optional flags can test exists.

		// Optional flag.
		exists, err := util.CheckStringFlag("o", flags.o, util.FlagOptional)
		if exists {
			// Do something with flag.
		}
*/
const (
	FlagRequired = true
	FlagOptional = false
)
func checkStringFlag(name string, arg string, required bool) (exists bool, err error) {
    var hasValidLookingArg bool

    if arg != "" {
        exists = true
    }

    // Try to detect missing flag argument.
    // If an argument is another flag, argument has not been provided.
    if exists && !strings.HasPrefix(arg, "-") {
        // Option expecting an argument but has been followed by another flag.
        hasValidLookingArg = true
    }
/*
    where(fmt.Sprintf("-%s required           = %t", name, required))
    where(fmt.Sprintf("-%s exists             = %t", name, exists))
    where(fmt.Sprintf("-%s hasValidLookingArg = %t", name, hasValidLookingArg))
    where(fmt.Sprintf("-%s value              = %s", name, arg))
*/
    if required && !exists {
        err = fmt.Errorf("missing required flag: -%s", name)
        return false, err
    }

    if exists && !hasValidLookingArg {
        err = fmt.Errorf("flag -%s needs a valid argument (not: %s)", name, arg)
        return false, err
    }

    return
}

/*
	Handle Cygwin environment.

	The problem:
		cygwinPath := "/cygdrive/c/mypath/myfile"
		windowsPath := filepath.Abs(cygwinPath)

	returns: "C:/cygdrive/c/mypath/myfile"

	It should return: "C:/mypath/myfile"
*/
func FilepathAbs(inputPath string) (path string, err error) {
	var OSTYPE string
	const cygwinRootOfAllDrives = "/cygdrive/"
	if strings.HasPrefix(inputPath, cygwinRootOfAllDrives) {
		OSTYPE = "cygwin"
		// OSTYPE := os.Getenv("OSTYPE")	// Is not helpful (returns nothing on Windows 10)
	}
	if OSTYPE == "cygwin" {	// Atypical case: cygwin drive.
		// Use cygwin utility cygpath to convert cygwin path to windows path.
		const executable = "cygpath"
		const flag = "-w"
		var cmd *exec.Cmd = exec.Command(executable, flag, inputPath)
		var out bytes.Buffer
		cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {
			err = fmt.Errorf("%s exit code %v error: %s", executable, err, out.String())
			return
		}
		path = out.String()
		// cygpath or cygwin bash appends path with an unwelcome new line.
		path = strings.Replace(path, "\n", "", -1)
	} else {	// Typical case.
		path, err = filepath.Abs(inputPath)
	}

	return
}

/*
	Pipe a Go program file (as a string) through the Go tool gofmt and return its output.

	Use it to tidy up generated Go source code before writing it to file.

	On error the input string is returned unchanged, not an empty "" string.
	This is unusual but we do that here to avoid crunching goProgramString in the calling function
	if it happens to be called like this:

		goProgramString, err = GoFmtProgramString(goProgramString)
		if err != nil {
			// goProgramString is unchanged, not crunched
		}
		// goProgramString has been formatted by gofmt

	Because this function calls out to gofmt in the operating system, the potential
	for failure is possible on some machines (and not testable by me during development).
	Hence a more forgiving return of its input string so as to avoid crunching user data.
*/
func GoFmtProgramString(goProgramString string) (formattedGoProgramString string, err error) {
	// We return the input string even if error, so as to not crunch it in the calling function.
	formattedGoProgramString = goProgramString

	var cmd *exec.Cmd = exec.Command("gofmt")

	var fileBytes []byte = []byte(goProgramString)
	cmd.Stdin = bytes.NewBuffer(fileBytes)

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil { return }

	formattedGoProgramString = out.String()

	return
}

/*
	See 1: https://stackoverflow.com/questions/35212985/is-it-possible-get-information-about-caller-function-in-golang

	See 2: http://moazzam-khan.com/blog/golang-get-the-function-callers-name

	This is a blend of both (above URLs) examples. Provides:-

	(1) The function name called.

	(2) The function name of the caller.

	(2) The file name[line number] at the call.

	This is intentionally a print-only function because calling it from another function (other than the one being
	tracked) will change the calling information by nesting to an additional level.
*/
func printCaller() {
	var calledName string
	var callerFile string
	var callerName string
	var n int	// number of callers
	var lastIndex int

	// Remove package name from function name and append ().
	var funcBaseName = func(longName string) (baseName string) {
		lastIndex = strings.LastIndex(longName, ".")
		if lastIndex >= 0 {
			baseName = longName[lastIndex+1:] + "()"
		}
		return baseName
	}

	fpcs := make([]uintptr, 1)

	// Skip 1 level to get the called: the name of the function calling PrintCaller()
	n = runtime.Callers(2, fpcs)
	if n == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "%s ERROR: no called\n", funcName())
		return
	}
	called := runtime.FuncForPC(fpcs[0]-1)
	if called == nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s ERROR: called was nil\n", funcName())
		return
	}
	calledName = called.Name()
	calledName = funcBaseName(calledName)

	// Skip 2 levels to get the caller
	n = runtime.Callers(3, fpcs)
	if n == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "%s ERROR: no caller\n", funcName())
		return
	}
	caller := runtime.FuncForPC(fpcs[0]-1)
	if caller == nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s ERROR: caller was nil\n", funcName())
		return
	}
	callerName = caller.Name()
	callerName = funcBaseName(callerName)

	// Get the file name and line number
	fileName, lineNum := caller.FileLine(fpcs[0]-1)
	fileName = filepath.Base(fileName)
	callerFile = fmt.Sprintf("%s[%d]", fileName, lineNum)

	_, _ = fmt.Fprintf(os.Stderr, "%s called by %s at %s\n", calledName, callerName, callerFile)
}

func funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	nameFull := runtime.FuncForPC(pc).Name() // main.foo
	nameEnd := filepath.Ext(nameFull)        // .foo
	name := strings.TrimPrefix(nameEnd, ".") // foo
	return name + "()"
}

func funcNameNoParens() string {
	pc, _, _, _ := runtime.Caller(1)
	nameFull := runtime.FuncForPC(pc).Name() // main.foo
	nameEnd := filepath.Ext(nameFull)        // .foo
	name := strings.TrimPrefix(nameEnd, ".") // foo
	return name
}

func funcNameFull() string {
	pc, sourceFile, lineNumber, ok := runtime.Caller(1)
	if !ok {
		return "Could not obtain func name and source file information."
	}
	nameFull := runtime.FuncForPC(pc).Name() // main.foo
	return fmt.Sprintf("%s[%d] %s", sourceFile, lineNumber, nameFull)
}

func funcSource() string {
	_, sourceFile, lineNumber, ok := runtime.Caller(1)
	if !ok {
		return "Could not obtain func name and source file information."
	}
	sourceBase := filepath.Base(sourceFile)
	return fmt.Sprintf("%s[%d]", sourceBase, lineNumber)
}

/*
Round is a custom implementation for rounding values.

Round up if fraction is >= 0.5 otherwise round down.

From: https://medium.com/@edoardo849/mercato-how-to-build-an-effective-web-scraper-with-golang-e2918d279f49#.istjzw4nl
*/
func round(val float64, places int) (rounded float64) {
	const roundOn = 0.5 // Round up if fraction is >= 0.5 otherwise round down.
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, frac := math.Modf(digit) // Modf(f) returns integer and fractional floating-point numbers that sum to f
	if frac >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	rounded = round / pow
	return
}

/*
	The word size (in bits) of the machine we're now running on. Typically 64 or 32 bits.

	Or use: intBits := strconv.IntSize
*/
func wordSize() int {
	return 32 << (^uint(0) >> 32 & 1)
}

// Check to see if this program can read piped input.
func canReadFromPipe() (bool, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}

	if info.Mode() & os.ModeCharDevice != 0 || info.Size() <= 0 {
		return true, nil
	}

	return false, nil
}

func gulpFromPipeWithTimeout(timeout time.Duration) (input string, err error) {
where("GulpFromPipeWithTimeout")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
where("defer cancel()")
	defer cancel()

where("BEFORE go input, err = GulpFromPipe()")
	go func() {
		input, err = gulpFromPipe()
	}()
where("AFTER go input, err = GulpFromPipe()")
	select {
		case <-ctx.Done():
where("case <-ctx.Done():")
			return "", fmt.Errorf("(1) Didn't see any piped input before timeout: %v", timeout)
	}

	return "", fmt.Errorf("(2) Didn't see any piped input before timeout: %v", timeout)
}

// Read and return piped input as a string.
func gulpFromPipe() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	var output []rune
	for {
		inputRune, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		output = append(output, inputRune)
	}

	return string(output), nil
}

func isCommandInstalled(commandName string) (bool, error) {
	path, err := exec.LookPath(commandName)
	if err != nil {
		return false, fmt.Errorf("%v: command %s is not installed in path %s", err, commandName, path)
	}

	return true, nil
}
