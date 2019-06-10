package errbase_test

import (
	"fmt"

	"github.com/KoharaKazuya/errbase"
	"golang.org/x/xerrors"
)

type MyCustomError struct {
	errbase.Err
	CustomField string
}

func Example() {
	err := doSomething1()
	if err != nil {
		var x *MyCustomError
		if xerrors.As(err, &x) {
			fmt.Printf("%v", x.CustomField)
		}
	}

	// Output:
	// additional information
}

func doSomething1() error {
	// generate custom error
	err := &MyCustomError{
		CustomField: "additional information",
	}
	err.Build("failed to xxx") // here, generates call stack info
	return err
}
