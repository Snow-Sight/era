package era

import (
	"fmt"
)

// New returns a new Error, with the just the Err set to the m as an error.
func New(err string) *Error {
	return &Error{
		Err: fmt.Errorf(err),
	}
}
