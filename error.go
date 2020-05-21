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
	Code int
	Err  error
}

// Error implement error
func (ce CodeErr) Error() string {
	return fmt.Sprintf("%d:%v", ce.Code, ce.Err)
}

// Code return code from err.
// It returns CodeOk when err is nil.
// It returns CodeUnknown when err is not CodeErr.
func Code(err error) int {
	if err == nil {
		return CodeOk
	}
	if c, ok := err.(CodeErr); ok {
		return c.Code
	} else {
		return CodeUnknown
	}
}

// Error return new CodeErr
func Error(code int, msg string) error {
	return CodeErr{
		Code: code,
		Err:  errors.New(msg),
	}
}

// FromError create CodeErr with code from err
func FromError(code int, err error) error {
	return CodeErr{
		Code: code,
		Err:  err,
	}
}

// UnknownError return CodeErr with CodeUnknown
func UnknownError(err error) error {
	return FromError(CodeUnknown, err)
}
