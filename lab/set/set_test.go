package set

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

func TestSet(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(2)
	assert.Equal(t, s.Len(), 3)
	assert.Equal(t, s.Contains(2), true)
	s2 := New()
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	assert.Equal(t, fmt.Sprint(s.Elements()), "[1 2 3]")
	s.Remove(2)
	assert.Equal(t, s.Len(), 2)
	s.Clear()
	assert.Equal(t, s.Len(), 0)

}
