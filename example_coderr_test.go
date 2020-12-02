package coderr_test

import (
	"fmt"

	"github.com/foxmeder/coderr"
	"github.com/pkg/errors"
)

func Example() {
	// normal error with code
	err := coderr.Error(30, "error with code")
	code := coderr.Code(err)
	fmt.Println(code)
	fmt.Println(err)
	// create from error
	errFromGo := errors.New("golang error")
	fmt.Println(coderr.Code(errFromGo))
	err = coderr.FromError(80, errFromGo)
	fmt.Println(err)
	// wrap error
	err = errors.Wrap(err, "wrapped error")
	fmt.Println(coderr.Code(err))
	err = coderr.UnknownError(errFromGo)
	fmt.Println(coderr.Code(err))
	fmt.Println(coderr.Code(nil))
	// Output:30
	// 30:error with code
	// 5000
	// 80:golang error
	// 80
	// 5000
	// 0
}
