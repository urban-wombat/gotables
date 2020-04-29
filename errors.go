package gotables

import (
	"errors"
	"fmt"
)

/*
	Basic gotables error handling:

	(1) Define a type struct: <type-name>Error

	(2) Define a method Error() string to implement error.

	(3) Define a factory function New<type-name>Error(...) *<type-name>Error.

	(4) Define a function HasGet<type-name>Error (bool, *<type-name>Error)

	(5) Define methods to get <type-name>Error struct private members.
*/

type CircRefError struct {
	rootTable string
	circTable string
	msg       string
}

func (circError *CircRefError) Error() string {
	return circError.msg
}

func NewCircRefError(rootTable string, circTable string) *CircRefError {
	var circError CircRefError
	circError.rootTable = rootTable
	circError.circTable = circTable
	circError.msg = fmt.Sprintf("circular reference in table [%s]: a reference to table [%s] already exists",
		circError.rootTable,
		circError.circTable)
	return &circError
}

// Check to see if err has a wrapped CircRefError inside.
func HasCircRefError(err error) (has bool) {
	// second argument to errors.As must be a pointer to an interface or a type implementing error
	var circError *CircRefError
	has = errors.As(err, &circError)
	return
}

// Check to see if err has a wrapped CircRefError inside, and get CircRefError if inside.
func GetCircRefError(err error) (circError *CircRefError) {
	// second argument to errors.As must be a pointer to an interface or a type implementing error
	errors.As(err, &circError)
	return
}

func (circError *CircRefError) RootTable() string {
	return circError.rootTable
}

func (circError *CircRefError) CircTable() string {
	return circError.circTable
}
