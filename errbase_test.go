package errbase_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/KoharaKazuya/errbase"
	"golang.org/x/xerrors"
)

// check interfaces
var (
	_ error             = &errbase.Err{}
	_ fmt.Formatter     = &errbase.Err{}
	_ xerrors.Formatter = &errbase.Err{}
)

type SomethingError struct {
	errbase.Err
}

func TestNoWrapMessage(t *testing.T) {
	err := new(SomethingError)
	err.Build("test msg")

	expected := "test msg"
	actual := err.Error()
	if actual != expected {
		t.Errorf("expected: %#v, got: %#v", expected, actual)
	}
}

func TestWrapMessage(t *testing.T) {
	err := new(SomethingError)
	err.Wrap("test msg", xerrors.New("cause msg"))

	expected := "test msg: cause msg"
	actual := err.Error()
	if actual != expected {
		t.Errorf("expected: %#v, got: %#v", expected, actual)
	}
}

func TestTypeAssertion(t *testing.T) {
	err := new(SomethingError)
	err.Build("test msg")

	var x *SomethingError
	if !xerrors.As(err, &x) {
		t.Errorf("type assertion error")
	}
}

func TestWrap(t *testing.T) {
	cause := xerrors.New("cause msg")

	err := new(SomethingError)
	err.Wrap("test msg", cause)

	actual := fmt.Sprintf("%+v", err)
	if !strings.Contains(actual, "cause msg") {
		t.Errorf("error message doesn't contain wrapped error message (actual: %s)", actual)
	}
}
