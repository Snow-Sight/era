package era

import (
	"fmt"
	"testing"
)

func TestSafe(t *testing.T) {
	tt := []struct {
		err      error
		output   string
		scenario string
	}{
		{
			&Error{
				Code:    EConflict,
				Err:     fmt.Errorf("error"),
				Message: "A helpful message.",
			},
			"<5fe41> conflict: A helpful message.",
			"Provided message and code should ",
		},
		{
			fmt.Errorf("test"),
			"<00000> internal: An internal error has occurred.",
			"Regular error should return default code, message and zero hash",
		},
		{
			nil,
			"",
			"Nil error should return nothing /empty string",
		},
		{
			&Error{
				Code:    EInternal,
				Message: "A message.",
				Err: &Error{
					Code:    EConflict,
					Message: "A lower message.",
				},
			},
			"<4e0fc> conflict: A lower message.",
			"Nested custom error should return lowest error code and message along with hash",
		},
		{
			&Error{
				Code:    EInternal,
				Message: "A message.",
				Err: &Error{
					Code: EConflict,
				},
			},
			"<83047> conflict: A message.",
			"Nested custom error should return lowest error code and message even if on different levels",
		},
		{
			&Error{
				Code:    EConflict,
				Message: "A message.",
				Err: &Error{
					Message: "A lower message.",
				},
			},
			"<66050> conflict: A lower message.",
			"Nested custom error should return lowest error code and message even if on different levels",
		},
	}

	for _, test := range tt {
		e := Safe(test.err)

		if e == nil && test.output != "" {
			t.Errorf("%s: expected %s got nil error", test.scenario, test.output)
		} else if e == nil {
			continue
		} else if e.Error() != test.output {
			t.Errorf("%s: expected %s got %s", test.scenario, test.output, Safe(test.err))
		}
	}
}
