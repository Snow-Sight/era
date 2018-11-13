package era

const (
	// EConflict Action cannot be performed
	EConflict = "conflict"
	// EInternal internal error
	EInternal = "internal"
	// EInvalid validation failed
	EInvalid = "invalid"
	// ENotFound entity not found/doesn't exist
	ENotFound = "not_found"
	// EAlreadyExists entity already exists
	EAlreadyExists = "already_exists"
	// EPermissionDenied Authenticated user does not have permission
	EPermissionDenied = "permission_denied"
	// EUnauthenticated Requestor does not have valid authentication to perform to operation
	EUnauthenticated = "unauthenticated"
	// ECannotDecode Data could not be decoded
	ECannotDecode = "cannot_decode"
	// ECannotEncode Data could not be encoded
	ECannotEncode = "cannot_encode"
	// ECannotParse Data could not be parsed
	ECannotParse = "cannot_parse"
)

// WithCode adds an error code to the provided error.
// If the err is an era.Error && err.Code is undefined, the code is applied to the era.Error;
// If the err is an era.Error && err.Code is defined, the err is wrapped and the code is applied;
// If the err is a regular error, the error is wrapped and the code is applied.
func WithCode(err error, code string) *Error {
	e, ok := err.(*Error)

	if !ok {
		return &Error{
			Err:  err,
			Code: code,
		}
	}
	if e.Code == "" {
		e.Code = code
		return e
	} else if e.Code == code {
		return e
	}

	return &Error{
		Err:  err,
		Code: code,
	}
}

// WithCode adds an error code to the provided custom Error.
// The withcode method is useful as opposed to the era.WithCode() function,
// as it allows for errors to be composed in a chain like format.
// If WithCode is applied to a custom Error with a defined code, it is wrapped.
func (e *Error) WithCode(code string) *Error {
	if e.Code == "" {
		e.Code = code
		return e
	} else if e.Code == code {
		return e
	}

	return &Error{
		Err:  e,
		Code: code,
	}
}

// Code returns the error code of the uppermost error object if type is of era.Error;
// If the error is a regular error or no code is specified, Code returns era.EInternal;
// Providing the code of the highest error object encourages codes to be applied at each level of the stack.
func Code(err error) string {
	if err == nil {
		return ""
	}

	if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	}

	return EInternal
}

// HighestCode returns the highest defined error code of the error object if type is of era.Error;
// If the error is a regular error or no code is specified in the entire error stack, Code returns era.EInternal;
// Unlike Code(), HighestCode() will recurse down in order to find the first defined code.
func HighestCode(err error) string {
	if err == nil {
		return ""
	}

	code := highestCode(err)

	if code != "" {
		return code
	}

	return EInternal
}

func highestCode(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		// Get the code of the nested error
		return highestCode(e.Err)
	}

	return ""
}

// LowestCode returns the lowest error code of the error, if available. Otherwise returns EInternal.
func LowestCode(err error) string {
	if err == nil {
		return ""
	}

	c := lowestCode(err)

	if c != "" {
		return c
	}

	return EInternal
}

// lowestCode returns the lowest code found or an empty string if no code is found
func lowestCode(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Err == nil {
		return e.Code
	} else if ok && e.Err != nil {
		// Get the code of the nested error
		c := lowestCode(e.Err)

		// If the nested error has no code, return the current error code
		if c != "" {
			return c
		}

		return e.Code
	}

	return ""
}
