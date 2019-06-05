# errbase

[![license widget]][license] [![godoc widget]][godoc] [![circleci widget]][circleci]

Custom error type generation helper for Go.

```go
package main

import (
	"fmt"

	"github.com/KoharaKazuya/errbase"
	"golang.org/x/xerrors"
)

type MyCustomError struct {
	errbase.Err
	CustomField string
}

func main() {
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
```

## Motivation

To leverage error handling, we often use new standard error interface in Go 1.13+ or with xerrors package. However, we feel bored to implement for each error types.

See below example. This may be what you want to implement for **each error**.

```go
type MyCustomError struct {
	err         error
	msg         string
	frame       xerrors.Frame
	CustomField string
}

func (e *MyCustomError) Error() string {
	return msg
}

func (e *MyCustomError) Format(f fmt.State, c rune) {
	xerrors.FormatError(e, f, c)
}

func (e *MyCustomError) FormatError(p xerrors.Printer) error {
	p.Print(e.msg)
	e.frame.Format(p)
	return e.err
}
```

errbase package provides embed-able implementation for each error types. It lets you to define the specialized error easily.

## Basic Design

- Used by embeded
- Identified by `errors.As`
- Provides call stack printing
- Provides error chain printing
- Not `errors.Wrapper`

### Used by embeded

errbase package provides a base type (`errbase.Err`). By embeded, it works.

```go
type MyCustomError struct {
	errbase.Err // embed
}
```

### Identified by `errors.As`

Your custom type embeded `errbase.Err` must be identified by `errors.As`. See [Go 2 Draft Designs](https://go.googlesource.com/proposal/+/master/design/go2draft-error-inspection.md) for details.

```go
if err != nil {
	var x *MyCustomError
	if xerrors.As(err, &x) {
		// x is a custom error
	}
}
```

### Provides call stack printing

You can print call stack if you need to know where error occured. Use `fmt.Printf` and `%+v`.

```go
// error occur
err := &MyCustomError{
	CustomField: "additional information",
}
err.Build("failed to xxx") // here, generates call stack info
return err

// print call stack
if err != nil {
	fmt.Printf("%+v", err)
}
```

### Provides error chain printing

You can print error chain if you need to know why error occured. Use `fmt.Printf` and `%+v`.

```go
// error occur
cause := ... // other error
if cause != nil {
	err := &MyCustomError{
		CustomField: "additional information",
	}
	err.Wrap("failed to xxx", cause) // wrap error cause
	return err
}

// print error with cause
if err != nil {
	fmt.Printf("%+v", err)
}
```

### Not `errors.Wrapper`

errbase package does not provide implementation for `errors.Wrapper`. It means that you cannot identify error cause of error.
This is because I think that implementators of error should translate errors into their own errors, not just expose error cause. Errors are also part of API.

If you want to expose error cause as your own error, just use `fmt.Errorf` and `: %w`.

```go
return fmt.Errorf("error found: %w", err)
```

or, you can implement `errors.Wrapper` yourself.

```go
type MyCustomError struct {
	errbase.Err
	cause error
}

func (e *MyCustomError) Unwrap() error {
	return e.cause
}
```

[license]: https://github.com/KoharaKazuya/errbase/blob/master/LICENSE
[license widget]: https://img.shields.io/github/license/KoharaKazuya/errbase.svg
[godoc]: https://godoc.org/github.com/KoharaKazuya/errbase
[godoc widget]: https://godoc.org/github.com/KoharaKazuya/errbase?status.svg
[circleci]: https://circleci.com/gh/KoharaKazuya/errbase
[circleci widget]: https://circleci.com/gh/KoharaKazuya/errbase.svg?style=svg
