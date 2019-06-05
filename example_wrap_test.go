package errbase_test

import (
	"fmt"

	"github.com/KoharaKazuya/errbase"
	"golang.org/x/xerrors"
)

type MyCustomErrorWithCause struct {
	errbase.Err
}

func Example_wrap() {
	err := doSomething2()
	if err != nil {
		fmt.Printf("%+v", err) // print stack trace, error cause also
	}
}

func doSomething2() error {
	err := xerrors.New("lower error")
	if err != nil {
		// lower layer's error causes error
		myErr := &MyCustomErrorWithCause{}
		myErr.Wrap("failed to xxx", err) // here, generates call stack info
		return myErr
	}
	return nil
}
