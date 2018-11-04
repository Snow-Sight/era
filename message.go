package era

// Message returns the human-readable message of the error, if available.
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
