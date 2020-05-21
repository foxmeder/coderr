package coderr_test

import (
	"errors"
	"fmt"

	"github.com/foxmeder/coderr"
)

func Example() {
	err := coderr.Error(30, "error with code")
	code := coderr.Code(err)
	fmt.Println(code)
	fmt.Println(err)
	errFromGo := errors.New("golang error")
	fmt.Println(coderr.Code(errFromGo))
	err = coderr.FromError(80, errFromGo)
	fmt.Println(err)
	err = coderr.UnknownError(errFromGo)
	fmt.Println(coderr.Code(err))
	fmt.Println(coderr.Code(nil))
	// Output:30
	// 30:error with code
	// 5000
	// 80:golang error
	// 5000
	// 0
}
