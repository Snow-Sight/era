package status

import (
	"fmt"
	"testing"

	"github.com/Snow-Sight/era"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	a := assert.New(t)

	e := &era.Error{
		Err:     fmt.Errorf("Some err"),
		Code:    era.EPermissionDenied,
		Message: "You do not have permission",
	}

	a.Equal(
		"rpc error: code = PermissionDenied desc = <4c38f> permission_denied: You do not have permission",
		Error(e).Error(),
		"era.Error should have correctly mapped code and safe message",
	)

	a.Equal(
		"rpc error: code = Internal desc = <00000> internal: An internal error has occurred.",
		Error(fmt.Errorf("Some err")).Error(),
		"Standard error should have Unknown code",
	)
}
