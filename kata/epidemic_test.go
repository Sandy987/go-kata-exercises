package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEpidemic(t *testing.T) {
	assert.Equal(t, 420, Epidemic(18, 432, 1004, 1, 0.00209, 0.51))
	assert.Equal(t, 461, Epidemic(12, 288, 1007, 2, 0.00206, 0.45))
	assert.Equal(t, 409, Epidemic(13, 312, 999, 1, 0.00221, 0.55))
}
