package status

import (
	"github.com/Snow-Sight/era"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Error converts an error into a grpc status error.
// This allows for correct grpc status codes to be sent.
func Error(e error) error {
	var c codes.Code

	switch era.Code(e) {
	case era.EUnauthenticated:
		c = codes.Unauthenticated
	case era.EPermissionDenied:
		c = codes.PermissionDenied
	case era.EAlreadyExists:
		c = codes.AlreadyExists
	case era.EConflict:
		c = codes.AlreadyExists
	case era.ENotFound:
		c = codes.NotFound
	case era.EInternal:
		c = codes.Internal
	default:
		c = codes.Unknown
	}

	return status.Error(c, era.Safe(e).Error())
}
