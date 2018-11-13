package era

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_Error(t *testing.T) {
	a := assert.New(t)

	// Test that single layer, op, code, error and hash are printed.
	e := &Error{
		Err:     fmt.Errorf("err"),
		Message: "Some message.",
		Op:      "SomeOp",
		Code:    EAlreadyExists,
	}

	m := e.Error()

	a.Equal("<81953> SomeOp, already_exists: err", m, "Single layer, op, code, error and hash are printed")

	// Highest code is printed
	e = &Error{
		Err: &Error{
			Err:  fmt.Errorf("err"),
			Code: ECannotDecode,
		},
		Message: "Some message.",
		Code:    EAlreadyExists,
	}

	m = e.Error()

	a.Equal("<bf9ff> already_exists: err", m, "Highest code is printed")

	// Op not printed if not part of highest error
	e = &Error{
		Err: &Error{
			Err:  fmt.Errorf("err"),
			Code: ECannotDecode,
			Op:   "SomeOp",
		},
		Message: "Some message.",
		Code:    EAlreadyExists,
	}

	m = e.Error()

	a.Equal("<240d7> already_exists: err", m, "Op not printed if not part of highest error")

	// If no err in stack, print message of top error
	e = &Error{
		Err: &Error{
			Code: ECannotDecode,
			Op:   "SomeOp",
		},
		Message: "Some message.",
		Code:    EAlreadyExists,
	}

	m = e.Error()

	a.Equal("<240d7> already_exists: Some message.", m, "If no err in stack, print message of top error")
}

func TestError_WithErr(t *testing.T) {
	a := assert.New(t)

	// When err is undefined on highest Error, err is set
	e := &Error{
		Err:     nil,
		Message: "A message",
	}

	e = e.WithErr(fmt.Errorf("Error"))

	a.NotNil(e.Err, "Err should be set")
	a.Equal("Error", e.Err.Error(), "Error should be set to correct standard value")

	// When err is defined as a standard error, the err is not overwritten
	e = &Error{
		Err:     fmt.Errorf("Specific E"),
		Message: "A message",
	}

	e = e.WithErr(fmt.Errorf("Overwritten error"))

	a.NotNil(e.Err, "Err is still defined")
	a.Equal("Specific E", e.Err.Error(), "Error is not overwritten when already defined")

	// When errors are nested, errors are not overwritten
	e = &Error{
		Err: &Error{
			Err: fmt.Errorf("Some err"),
		},
	}

	e = e.WithErr(fmt.Errorf("Overwritten error"))

	a.NotNil(e.Err, "Err still defined")
	a.Equal("<e3b0c> internal: Some err", e.Error(), "Error is not overwritten")
}
