package gotables

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
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

/*
UtilStringFlag implements the flag.Value interface https://golang.org/pkg/flag/#Value
	type Value interface {
		String() string
		Set(string) error
	}
*/
type UtilStringFlag struct {
	val    string // string field used by the flag.Value interface https://golang.org/pkg/flag/#Value
	exists bool
	set    bool
	err    error
}

// Set() implements part of the flag.Value interface https://golang.org/pkg/flag/#Value
func (sf *UtilStringFlag) Set(s string) error {
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

// String() implements part of the flag.Value interface https://golang.org/pkg/flag/#Value
func (sf *UtilStringFlag) String() string {
	return sf.val
}

// Exists() is specific to gotables.Util
func (sf *UtilStringFlag) Exists() bool {
	return sf.exists
}

// IsSet() is specific to gotables.Util
func (sf *UtilStringFlag) IsSet() bool {
	return sf.set
}

// Error() is specific to gotables.Util
func (sf *UtilStringFlag) Error() error {
	return sf.err
}

/*
AllOk() is specific to gotables.Util
It means:-
	(1) flag exists
	(2) flag is set
	(3) error is nil
*/
func (sf *UtilStringFlag) AllOk() bool {
	return (sf.Exists() && sf.IsSet() && sf.Error() == nil)
}

// Print to stdout UtilStringFlag field values and method results.
func (sf *UtilStringFlag) Print() {
	fmt.Fprintf(os.Stderr, "%#v\n", sf)
	fmt.Fprintf(os.Stderr, "&UtilStringFlag.String() = %q\n", sf.String())
	fmt.Fprintf(os.Stderr, "&UtilStringFlag.Exists() = %t\n", sf.Exists())
	fmt.Fprintf(os.Stderr, "&UtilStringFlag.IsSet()  = %t\n", sf.IsSet())
	fmt.Fprintf(os.Stderr, "&UtilStringFlag.Error()  = %v\n", sf.Error())
	fmt.Fprintf(os.Stderr, "&UtilStringFlag.AllOk()  = %v\n", sf.AllOk())
}

/*
	Utility function to test string flags.

	It avoids boilerplate code testing flags.

	It can be called and:-

	(1) Required flags can trust the existence of an argument.

		// Required flag.
		exists, err := UtilCheckStringFlag("r", flags.r, UtilFlagRequired)
		if !exists {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

	(2) Optional flags can test exists.

		// Optional flag.
		exists, err := UtilCheckStringFlag("o", flags.o, UtilFlagOptional)
		if exists {
			// Do something with flag.
		}
*/
const (
	FlagRequired = true
	FlagOptional = false
)

func UtilCheckStringFlag(name string, arg string, required bool) (exists bool, err error) {
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
func UtilFilepathAbs(inputPath string) (path string, err error) {
	var OSTYPE string
	const cygwinRootOfAllDrives = "/cygdrive/"
	if strings.HasPrefix(inputPath, cygwinRootOfAllDrives) {
		OSTYPE = "cygwin"
		// OSTYPE := os.Getenv("OSTYPE")	// Is not helpful (returns nothing on Windows 10)
	}
	if OSTYPE == "cygwin" { // Atypical case: cygwin drive.
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
	} else { // Typical case.
		path, err = filepath.Abs(inputPath)
	}

	return
}

func UtilFormatSource(source string) (formattedSource string, err error) {
	var formattedSourceBytes []byte
	formattedSourceBytes, err = format.Source([]byte(source))
	if err != nil {
		return "", err
	}

	formattedSource = string(formattedSourceBytes)

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
func UtilPrintCaller() {
	var calledName string
	var callerFile string
	var callerName string

	var n int // number of callers
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
		_, _ = fmt.Fprintf(os.Stderr, "%s ERROR: there was no called\n", UtilFuncName())
		return
	}
	called := runtime.FuncForPC(fpcs[0] - 1)
	if called == nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s ERROR: called was nil\n", UtilFuncName())
		return
	}
	calledName = called.Name()
	calledName = funcBaseName(calledName)

	// Skip 2 levels to get the caller
	n = runtime.Callers(3, fpcs)
	if n == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "%s ERROR: there was no caller\n", UtilFuncName())
		return
	}
	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s ERROR: caller was nil\n", UtilFuncName())
		return
	}
	callerName = caller.Name()
	callerName = funcBaseName(callerName)

	// Get the file name and line number
	fileName, lineNum := caller.FileLine(fpcs[0] - 1)
	fileName = filepath.Base(fileName)
//	callerFile = fmt.Sprintf("%s[%d]", fileName, lineNum)
	callerFile = fmt.Sprintf("%s:%d:", fileName, lineNum)

	_, _ = fmt.Fprintf(os.Stderr, "ttt UtilPrintCaller(): %s called by %s at %s\n", calledName, callerName, callerFile)
}

/*
	BEWARE: NOT PROPERLY TESTED!

	See 1: https://stackoverflow.com/questions/35212985/is-it-possible-get-information-about-caller-function-in-golang

	See 2: http://moazzam-khan.com/blog/golang-get-the-function-callers-name

	This is a blend of both (above URLs) examples. Provides:-

	(1) The function name called.

	(2) The function name of the caller's caller.

	(2) The file name[line number] at the call.

	This is intentionally a print-only function because calling it from another function (other than the one being
	tracked) will change the calling information by nesting to an additional level.
*/
func UtilPrintCallerCaller() {
	var calledName string
	var callerName string
	var callerCallerFile string
	var callerCallerName string

	var n int // number of callers
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
		_, _ = fmt.Fprintf(os.Stderr, "%s ERROR: there was no called\n", UtilFuncName())
		return
	}
	called := runtime.FuncForPC(fpcs[0] - 1)
	if called == nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s ERROR: called was nil\n", UtilFuncName())
		return
	}
	calledName = called.Name()
	calledName = funcBaseName(calledName)

	// Skip 2 levels to get the caller
	n = runtime.Callers(3, fpcs)
	if n == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "%s ERROR: there was no caller\n", UtilFuncName())
		return
	}
	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s ERROR: caller was nil\n", UtilFuncName())
		return
	}
	callerName = caller.Name()
	callerName = funcBaseName(callerName)

	// Skip 3 levels to get the caller's caller
	n = runtime.Callers(4, fpcs)
	if n == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "%s ERROR: there was no caller\n", UtilFuncName())
		return
	}
	callerCaller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s ERROR: caller was nil\n", UtilFuncName())
		return
	}
	callerCallerName = callerCaller.Name()
	callerCallerName = funcBaseName(callerCallerName)

	// Get the file name and line number
	fileName, lineNum := callerCaller.FileLine(fpcs[0] - 1)
	fileName = filepath.Base(fileName)
//	callerCallerFile = fmt.Sprintf("%s[%d]", fileName, lineNum)
	callerCallerFile = fmt.Sprintf("%s:%d:", fileName, lineNum)

	_, _ = fmt.Fprintf(os.Stderr, "UtilPrintCallerCaller(): %s called by %s at %s\n", calledName, callerCallerName, callerCallerFile)
}

/*
	Short function name with parentheses.

		pkgName.funcName

	becomes:

		funcName()
*/
func UtilFuncName() string {
	const stackFramesToSkip = 1
	pc, _, _, _ := runtime.Caller(stackFramesToSkip)
	nameFull := runtime.FuncForPC(pc).Name() // main.foo
	nameEnd := filepath.Ext(nameFull)        // .foo
	name := strings.TrimPrefix(nameEnd, ".") // foo
	return name + "()"
}

/*
	Return the name of the function that called this function.
*/
func UtilFuncCaller() string {
	const stackFramesToSkip = 2
	pc, _, _, _ := runtime.Caller(stackFramesToSkip)
	nameFull := runtime.FuncForPC(pc).Name() // main.foo
	nameEnd := filepath.Ext(nameFull)        // .foo
	name := strings.TrimPrefix(nameEnd, ".") // foo
	return name + "()"
}

/*
	Return the name of the function that called the caller of this function.
*/
func UtilFuncCallerCaller() string {
	const stackFramesToSkip = 3
	pc, _, _, _ := runtime.Caller(stackFramesToSkip)
	nameFull := runtime.FuncForPC(pc).Name() // main.foo
	nameEnd := filepath.Ext(nameFull)        // .foo
	name := strings.TrimPrefix(nameEnd, ".") // foo
	return name + "()"
}

/*
	Return the name of the function that called the caller of this function.
*/
func UtilFuncCallerCallerCaller() string {
	const stackFramesToSkip = 4
	pc, _, _, _ := runtime.Caller(stackFramesToSkip)
	nameFull := runtime.FuncForPC(pc).Name() // main.foo
	nameEnd := filepath.Ext(nameFull)        // .foo
	name := strings.TrimPrefix(nameEnd, ".") // foo
	return name + "()"
}

/*
	Return the name of the top-most function that ultimately called this function.

	Warning: Cannot control for relevant top caller.
*/
func utilFuncTopCaller() string {
	var nameFull string
	var pcSlice []uintptr
	var pc uintptr
	var ok bool
	var skip int

	pc, _, _, ok = runtime.Caller(skip)
	for skip = 0; ok; skip++ {
		pc, _, _, ok = runtime.Caller(skip)
		pcSlice = append(pcSlice, pc)
		nameFull = runtime.FuncForPC(pc).Name() // main.foo
		where(fmt.Sprintf("xxxx %d %s", skip, nameFull))
	}
	const minSliceLen = 7
	if len(pcSlice) >= minSliceLen {
		pc = pcSlice[len(pcSlice)-(minSliceLen-2)]
	}
	nameFull = runtime.FuncForPC(pc).Name() // main.foo
	where(nameFull)
	nameEnd := filepath.Ext(nameFull)        // .foo
	name := strings.TrimPrefix(nameEnd, ".") // foo
	return name + "()"
}

/*
	Short function name with NO parentheses.

		pkgName.funcName

	becomes:

		funcName
*/
func UtilFuncNameNoParens() string {
	pc, _, _, _ := runtime.Caller(1)
	nameFull := runtime.FuncForPC(pc).Name() // main.foo
	nameEnd := filepath.Ext(nameFull)        // .foo
	name := strings.TrimPrefix(nameEnd, ".") // foo
	return name
}

/*
	Full path of function source code with line number followed by full package name of function. Like this:

		<mydrive>/golang/src/github.com/urban-wombat/gotables/util_test.go[39] github.com/urban-wombat/UtilTestFuncNameFull
*/
func UtilFuncNameFull() string {
	pc, sourceFile, lineNumber, ok := runtime.Caller(1)
	if !ok {
		return "Could not obtain func name and source file information."
	}
	nameFull := runtime.FuncForPC(pc).Name() // main.foo
//	return fmt.Sprintf("%s[%d] %s", sourceFile, lineNumber, nameFull)
	return fmt.Sprintf("%s:%d: %s", sourceFile, lineNumber, nameFull)
}

/*
	Short source file name and line number. Like this:

		util_test.go[39]
*/
func UtilFuncSource() string {
	_, sourceFile, lineNumber, ok := runtime.Caller(1)
	if !ok {
		return "Could not obtain func name and source file information."
	}
	sourceBase := filepath.Base(sourceFile)
//	return fmt.Sprintf("%s[%d]", sourceBase, lineNumber)
	return fmt.Sprintf("%s:%d:", sourceBase, lineNumber)
}

/*
	Current source line number. A bit like the C preprocessor macro:
		__LINE__
*/
func UtilLineNumber() (lineNumber int) {
	_, _, lineNumber, ok := runtime.Caller(1)
	if !ok {
		// Could not obtain func name and source file information.
		return -1
	}
	return lineNumber
}

/*
Round is a custom implementation for rounding values.

Round up if fraction is >= 0.5 otherwise round down.

From: https://medium.com/@edoardo849/mercato-how-to-build-an-effective-web-scraper-with-golang-e2918d279f49#.istjzw4nl
*/
func UtilRound(val float64, places int) (rounded float64) {
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
func UtilWordSize() int {
	return 32 << (^uint(0) >> 32 & 1)
}

// Check to see if this program can read piped input on this machine.
func UtilCanReadFromPipe() (bool, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		return true, nil
	}

	return false, nil
}

/*
	Read and return piped input as a string.

	Beware: this blocks waiting for stdin.

		stdin, err := UtilGulpFromPipe()
*/
func UtilGulpFromPipe() (string, error) {

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

/*
	Read and return piped input as a string.

	This waits for stdin but only until timeout expires.

		stdin, err := UtilGulpFromPipe(1 * time.Second)
*/
func UtilGulpFromPipeWithTimeout(timeout time.Duration) (input string, err error) {

	c1 := make(chan string, 1)

	go func() {
		input, err = UtilGulpFromPipe()
		c1 <- input
	}()

	select {
	case result := <-c1:
		return result, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("did not read any piped input from stdin after waiting %v", timeout)
	}
}

/*
	Check whether commandName is installed on this machine.
*/
func UtilIsCommandInstalled(commandName string) (bool, error) {
	path, err := exec.LookPath(commandName)
	if err != nil {
		return false, fmt.Errorf("%v: command %s is not installed in path %s", err, commandName, path)
	}

	return true, nil
}

func UtilProgName() string {
	return filepath.Base(os.Args[0])
}

/*
	Return a string with the build date/time and (seconds-ago) of the executable and where it is installed.
*/
func UtilDateTimeOfBuild() (dateTimeOfBuild string) {
	executableName := os.Args[0]
	stat, err := os.Stat(executableName)
	if err == nil {
		var ago time.Duration = time.Now().Sub(stat.ModTime()).Truncate(time.Second)
		executableName = strings.Replace(executableName, ".exe", "", 1)
		executableName = filepath.Base(executableName)
		dateTimeOfBuild = fmt.Sprintf("%s.go built %s (%v ago) installed %s\n",
			executableName, stat.ModTime().Format(time.UnixDate), ago, os.Args[0])
	}
	return
}
