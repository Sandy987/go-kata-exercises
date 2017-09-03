package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrossover(t *testing.T) {
	xs, ys := Crossover([]int{3, 5, 1, 1, 3}, []int{1, 1, 1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2, 2, 2})
	assert.Equal(t, []int{1, 2, 2, 1, 1, 2, 2}, xs)
	assert.Equal(t, []int{2, 1, 1, 2, 2, 1, 1}, ys)
}
