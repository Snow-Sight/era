package era

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithCode(t *testing.T) {
	a := assert.New(t)
	// Wrapping a standard error with a code
	// Should result in a single depth Error object, with the matching code.
	e := WithCode(fmt.Errorf("an error"), EInternal)

	a.Equal(EInternal, e.Code, "Expect code to be correct")

	a.NotNil(e.Err, "Error.Err should be defined")

	a.Equal("an error", e.Err.Error(), "Expect error to be defined and correct")

	// Wrapping a custom Error without a code, the code should be applied to the top level error wrap.
	e = &Error{
		Err:  fmt.Errorf("Error"),
		Code: "",
	}

	e = WithCode(e, EConflict)

	a.Equal(EConflict, e.Code, "Code on root error should be set")
	a.Equal("Error", e.Err.Error(), "Error is not wrapped when code was originally not set")

	// Provided a custom Error with a code, the error should be wrapped and the new code added.
	e = &Error{
		Err:  fmt.Errorf("Error"),
		Code: "some_code",
	}
	e = WithCode(e, ECannotDecode)

	a.Equal(ECannotDecode, e.Code, "Code is applied to new wrap")
	a.NotNil(e.Err, "Error is wrapped with nested error maintained")

	// Provided a custom Error with a code adding a code that is the same should not do anything.
	e = &Error{
		Err:  fmt.Errorf("Error"),
		Code: ECannotDecode,
	}
	e = WithCode(e, ECannotDecode)

	a.Equal(ECannotDecode, e.Code, "Code is applied")
	a.NotNil(e.Err, "Error is wrapped with nested error maintained")
	a.Equal("Error", e.Err.Error(), "Error is not wrapped again")
}

func TestCode(t *testing.T) {
	tt := []struct {
		err      error
		output   string
		scenario string
	}{
		{
			&Error{
				Code:    EConflict,
				Message: "irrelevant",
			},
			"conflict",
			"Conflict error code in high level era object",
		},
		{
			fmt.Errorf("test"),
			"internal",
			"Regular error should return internal code",
		},
		{
			nil,
			"",
			"Nil error should return no code/empty string",
		},
		{
			&Error{
				Code:    EInternal,
				Message: "irrelevant",
				Err: &Error{
					Code:    EConflict,
					Message: "Conflict",
				},
			},
			EInternal,
			"Nested custom error should return error code or top most object",
		},
	}

	for _, test := range tt {
		if Code(test.err) != test.output {
			t.Errorf("%s: expected %s got %s", test.scenario, test.output, Code(test.err))
		}
	}
}

func TestHighestCode(t *testing.T) {
	a := assert.New(t)

	a.Equal("", HighestCode(nil), "Nil error results in empty string code")

	e := &Error{
		Err:  fmt.Errorf("err"),
		Code: EUnauthenticated,
	}

	a.Equal(EUnauthenticated, HighestCode(e), "Single layer custom error with defined code returns defined code")

	e = &Error{
		Err: fmt.Errorf("err"),
	}

	a.Equal(EInternal, HighestCode(e), "Single layer custom error with no defined code returns EInternal")

	e = &Error{
		Err: &Error{
			Code: EInvalid,
			Err:  fmt.Errorf("err"),
		},
		Code: ECannotDecode,
	}

	a.Equal(ECannotDecode, HighestCode(e), "Nested custom error with two defined codes returns topmost code")

	e = &Error{
		Err: &Error{
			Code: EInvalid,
			Err:  fmt.Errorf("err"),
		},
		Message: "Message",
	}

	a.Equal(EInvalid, HighestCode(e), "Nested custom error with defined code within nested error returns first defined code")
}
