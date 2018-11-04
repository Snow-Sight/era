package main

import (
	"fmt"

	"github.com/Snow-Sight/era"
)

func main() {
	// An example of era being used in username validation.
	err := validateUserName("bob")

	if err != nil {
		fmt.Printf("The HighestCode() of the err for username bob: %s\n", era.HighestCode(err))
		fmt.Printf("Internal representation of err for username bob: %s\n", err.Error())
		fmt.Printf("Error that may be sent to customer/external user for username bob: %s\n", era.Safe(err).Error())
	}

	// An example of era being used in username validation.
	err = validateUserName("alice")

	if err != nil {
		fmt.Printf("Internal representation of err for username alice: %s\n", err.Error())
		fmt.Printf("Error that may be sent to customer/external user for username alice: %s\n", era.Safe(err).Error())
	}
}

// validateUserName returns an error if the username is invalid and returns nil if the username is valid
func validateUserName(uname string) error {
	var err error
	if len(uname) < 5 {
		// NewSafe creates a new error where both the message and err are the same.
		// This is used when creating new errors, if the internal err and the external message
		// should be different use era.New() and manually set the message with era.WithMessage(err, message)
		err = era.New("The provided username is too short")

		// WithCode adds a code to the top error
		err = era.WithCode(err, era.EInvalid)

		// At this point err is the same as...
		// err = &era.Error{
		// 	Message: "The provided username is too short",
		// 	Err: fmt.Errorf("The provided username is too short"),
		// 	Code: era.EInvalid,
		// }

		return err
	}

	return nil
}
