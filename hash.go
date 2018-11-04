package era

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

// Hash returns a five character hexadecimal hash of the error stack.
// This is useful for customer service and log analytics usecases as errors from the same cause can be quickly correlated.
// The hash function traverses down the error stack if the error is an era.Error,
// calculating the hash based on the Code, Message and Op.
// If the passed error is a regular error, Hash returns five zero characters.
func Hash(err error) string {
	e, ok := err.(*Error)

	if !ok {
		return "00000"
	}

	h := sha256.New()

	calchash(h, e)

	b := h.Sum(nil)

	return hex.EncodeToString(b)[0:5]
}

func calchash(h hash.Hash, err *Error) {
	if e, ok := err.Err.(*Error); ok {
		calchash(h, e)
	}

	h.Write([]byte(err.Message))
	h.Write([]byte(err.Op))
	h.Write([]byte(err.Code))

	return
}
