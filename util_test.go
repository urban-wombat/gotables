package gotables

import (
	"fmt"
	"strings"
	"testing"
)

func TestFormatSource(t *testing.T) {
	source := `
	if err != nil { return nil }
	`

	expected := `
	if err != nil {
		return nil
	}
	`

	formatted, err := UtilFormatSource(source)
	if err != nil {
		t.Error(err)
	}

	if formatted != expected {
		t.Fatalf("expecting %s but got %s", expected, formatted)
	}
}

func ExampleFuncName() {
	// Called from inside func ExampleFuncName()
	fmt.Println(UtilFuncName())

	// Output:
	// ExampleFuncName()
}

func ExampleFuncNameNoParens() {
	// Called from inside func ExampleFuncNameNoParens()
	fmt.Println(UtilFuncNameNoParens())

	// Output:
	// ExampleFuncNameNoParens
}

// Output can vary, so don't use as an example, such as:
// c:/golang/src/github.com/urban-wombat/util/util_test.go[40] github.com/urban-wombat/gotables.TestFuncNameFull
func TestFuncNameFull(t *testing.T) {
	name := UtilFuncNameFull()

	split := strings.Split(name, " ")
	if len(split) != 2 {
		t.Fatalf("expecting 2 strings but got %d", len(split))
	}

	pkg := split[1]

	pkgExpected := "github.com/urban-wombat/gotables.TestFuncNameFull"
	if pkg != pkgExpected {
		t.Fatalf("expecting pkg %s but got pkg %s", pkgExpected, pkg)
	}
}

// Output can vary, so don't use as an example, such as:
// c:/golang/src/github.com/urban-wombat/util/util_test.go[72]
func TestFuncSource(t *testing.T) {
	source := UtilFuncSource()
	split := strings.Split(source, " ")
	if len(split) != 1 {
		t.Fatalf("expecting 1 strings but got %d", len(split))
	}
}
