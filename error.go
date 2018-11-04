package era

import (
	"bytes"
	"fmt"
)

// Error defines a standard application error.
type Error struct {
	// Machine-readable error code.
	// Error codes can be custom, but a set of predefined codes are defined in codes.go.
	Code string

	// Human-readable message, potentially consumer facing error.
	// This message can be delivered to a consumer as an indication of what happened and what they should do next.
	Message string

	// Logical operation
	Op string
	// Nested Error
	Err error
}

// Error returns the string representation of the error message.
// Error returns a string prefixed with the top
func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the error hash
	fmt.Fprintf(&buf, "<%s> ", Hash(e))

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s, ", e.Op)
	}

	// Print the highest code
	fc := HighestCode(e)
	if fc != "" {
		fmt.Fprintf(&buf, "%s: ", fc)
	}

	// One of these should be defined, but just in case we don't want to panic.
	// Err should be present if the error is being recieved from some other package.
	// If it is not a message should be specified.
	rErr := e.rootErr()
	if rErr != nil {
		buf.WriteString(rErr.Error())
	} else {
		buf.WriteString(e.Message)
	}

	return buf.String()
}

func (e *Error) rootErr() error {
	// If wrapping a custom error
	if err, ok := e.Err.(*Error); ok {
		return err.rootErr()
	}

	return e.Err
}
