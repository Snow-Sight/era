package era

import (
	"fmt"
)

// Safe should be used when errors leave the scope of the service, and enter into a consumers domain.
// Safe provides the consumer with the error code and message, neither of these components should expose sensitive information.
// Safe returns an error for compatability reasons, an http handler may expect an error to be returned in the case of failure.
func Safe(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("<%s> %s: %s", Hash(err), LowestCode(err), Message(err))
}
