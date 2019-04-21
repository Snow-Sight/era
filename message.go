package era

// Message returns the human-readable message of the error, if available.
// The message will be the lowest message in the stack.
// Otherwise returns a generic error message.
func Message(err error) string {
	if err == nil {
		return ""
	}

	m := message(err)

	if m != "" {
		return m
	}

	return "An internal error has occurred."
}

// private message returns the lowest message found or no message if no message is found
func message(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Err == nil {
		return e.Message
	} else if ok && e.Err != nil {
		// Get the message of the nested error
		m := message(e.Err)

		// If the nested error has no message, return the current error message
		if m != "" {
			return m
		}

		return e.Message
	}

	return ""
}

// WithMessage adds a message to the error
func WithMessage(err error, message string) *Error {
	e, ok := err.(*Error)

	if !ok {
		return &Error{
			Err:     err,
			Message: message,
		}
	}
	if e.Code == "" {
		e.Message = message
		return e
	} else if e.Message == message {
		return e
	}

	return &Error{
		Err:     err,
		Message: message,
	}
}

// WithMessage adds a message to the Error
func (e *Error) WithMessage(message string) *Error {
	if e.Message == "" {
		e.Message = message
		return e
	} else if e.Message == message {
		return e
	}

	return &Error{
		Err:     e,
		Message: message,
	}
}
