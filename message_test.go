package era

import (
	"fmt"
	"testing"
)

func TestMessage(t *testing.T) {
	tt := []struct {
		err      error
		output   string
		scenario string
	}{
		{
			&Error{
				Message: "Here's a message",
			},
			"Here's a message",
			"Defined message should be returned",
		},
		{
			fmt.Errorf("test"),
			"An internal error has occurred.",
			"Regular error should return default statement",
		},
		{
			nil,
			"",
			"Nil error should return no message/empty string",
		},
		{
			&Error{
				Message: "Here's a message",
				Err: &Error{
					Message: "Here's a nested message",
					Err:     fmt.Errorf("error"),
				},
			},
			"Here's a nested message",
			"Lowest nested message should be returned",
		},
	}

	for _, test := range tt {
		if Message(test.err) != test.output {
			t.Errorf("%s: expected %s got %s", test.scenario, test.output, Message(test.err))
		}
	}

}
