/*
Package errbase is custom error type generation helper.
*/
package errbase

import (
	"fmt"

	"golang.org/x/xerrors"
)

// Err is a type assumed embedded in custom error type.
type Err struct {
	msg   string
	err   error
	frame xerrors.Frame
}

// Build sets error message and caller stack to error.
func (e *Err) Build(msg string) {
	e.msg = msg
	e.frame = xerrors.Caller(1)
}

// Wrap sets error message, error cause and caller stack to error.
func (e *Err) Wrap(msg string, err error) {
	e.msg = msg
	e.err = err
	e.frame = xerrors.Caller(1)
}

// Error returns an error message string.
func (e *Err) Error() string {
	cause := e.err
	if cause == nil {
		return e.msg
	}
	return fmt.Sprintf("%s: %v", e.msg, cause)
}

// Format prints the stack as error detail.
func (e *Err) Format(f fmt.State, c rune) {
	xerrors.FormatError(e, f, c)
}

// FormatError prints the stack as error detail.
func (e *Err) FormatError(p xerrors.Printer) error {
	p.Print(e.msg)
	e.frame.Format(p)
	return e.err
}
