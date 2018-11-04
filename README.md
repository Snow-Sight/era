# Era: Go Errors


[![codecov](https://codecov.io/gh/Snow-Sight/era/branch/master/graph/badge.svg)](https://codecov.io/gh/Snow-Sight/era)
[![CircleCI](https://circleci.com/gh/Snow-Sight/era.svg?style=svg)](https://circleci.com/gh/Snow-Sight/era)

A reusable package aimed at simplifying errors and tracing their source.
Heavily inspired by [Failure is your Domain](https://middlemost.com/failure-is-your-domain/).

# Why?

Era was made to solve two primary issues faced in writing go code. Firstly, Era aims to make error checking easier through the use of error codes. When using Era, standard codes or custom codes can be easily applied to errors, checking these codes can then be done using `era.Code(err)` which will return the lowest error code 

The second and more pressing concern behind the creation of Era is the separation of internal and external error messages. Serving a single error message to both developers trying to debug their code, and customers wondering why their request failed is infeasible in many situations and makes for poorer less descriptive error messages for one party. Era aims to fix this issue by separating errors, the era.Error.Message is made to be external/customer friendly, this message should provide a safe to expose explanation of what went wrong. The root of an era.Error stack (stack in the essence that errors are nested) should point to a regular error, this error explains to the developer what wen't wrong and can delve into technical detail that external users would not want to see.
Using this format, errors that are shipped to the customer can be done so using the `era.Safe(e error)` function. 
This function returns a regular error that only includes details of the operation, the error hash and the message, these are all details that should be customer friendly.

Another useful extra provided by the Era package is the `era.Hash(err error)` function, this computes a 5 character hexadecimal hash of the passed error. If the error is a regular error the hash will be 00000, otherwise the hash is computed on all the code, message and operations of the error stack. This hash is part of the return string of the `era.Error.Error()` method and the `era.Safe(err error)` function. The hash is useful in correlating errors from customers, as a customer can send the error hash they recieved as part of a support request. The hash will always be the same for the same error stack.


## The struct

The Error struct is designed as follows.

```go
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
```
