package util

import "testing"
import "log"
import "github.com/bmizerany/assert"

func TestReverseInt(t *testing.T) {
	a := []int{3, 2, 1, 4}
	ReverseInt(a)
	b := []int{4, 1, 2, 3}
	assert.Equal(t, a, b)
}
