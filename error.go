/*

package coderr generate error with code

get error code

	code := coderr.Code(err)

create new error

	err := coderr.Error(30, "error with code 30")

create from error

	err := coderr.FromError(30, err)

*/
package coderr

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	CodeOk      = 0    // ok
	CodeUnknown = 5000 // unknown error
)

// CodeErr error with code
type CodeErr struct {
	code int
	err  error
}

// Error implement error
func (ce CodeErr) Error() string {
	return fmt.Sprintf("%d:%v", ce.code, ce.err)
}

// Code return code from the first error in err's chain that matches CodeErr .
// It returns CodeOk when err is nil.
// It returns CodeUnknown if no CodeErr matched.
func Code(err error) int {
	if err == nil {
		return CodeOk
	}
	c := CodeErr{}
	if errors.As(err, &c) {
		return c.code
	} else {
		return CodeUnknown
	}
}

// Error return new CodeErr
func Error(code int, msg string) error {
	return CodeErr{
		code: code,
		err:  errors.New(msg),
	}
}

// FromError create CodeErr with code from err
func FromError(code int, err error) error {
	return CodeErr{
		code: code,
		err:  err,
	}
}

// UnknownError return CodeErr with CodeUnknown
func UnknownError(err error) error {
	return FromError(CodeUnknown, err)
}
