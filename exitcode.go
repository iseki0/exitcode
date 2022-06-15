package exitcode

import (
	"errors"
	"fmt"
	"os"
)

type Code int

func (c Code) Error() string {
	return fmt.Sprintf("exit code %d", int(c))
}

type ExitError struct {
	Code  Code
	Cause error
}

func (e *ExitError) Error() string {
	if e.Cause == nil {
		return e.Code.Error()
	}
	return e.Cause.Error()
}

func (e *ExitError) Unwrap() error {
	return e.Cause
}

func (e *ExitError) Is(target error) bool {
	return target == e.Code
}

// Of returns attached error code. Return zero if the error is not wrapped.
func Of(e error) Code {
	var t *ExitError
	if errors.As(e, &t) {
		return t.Code
	}
	return 0
}

func Wrap(e error, code Code) error {
	return &ExitError{
		Code:  code,
		Cause: e,
	}
}

func Exit(e error) {
	if c := Of(e); c > 0 {
		os.Exit(int(c))
	}
	os.Exit(1)
}
