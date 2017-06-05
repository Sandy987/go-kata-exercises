package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnnJohn(t *testing.T) {
	assert.Equal(t, []int{0, 0, 1, 2, 2, 3, 4, 4, 5, 6, 6}, John(11))
	//assert.Equal(t, []int{1, 1, 2, 2, 3, 3}, Ann(6))
	//assert.Equal(t, 1720, SumJohn(75))
	//assert.Equal(t, 4070, SumAnn(115))
}
