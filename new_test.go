package era

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := assert.New(t)

	e := New("Error and message")

	a.Equal("Error and message", e.Err.Error(), "Error.Err should be passed string")
}
