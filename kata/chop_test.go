package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChop(t *testing.T) {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	i := Chop(5, ints)
	assert.Equal(t, 5, i)

	i = Chop(-17, ints)
	assert.Equal(t, 0, i)

	i = Chop(10, ints)
	assert.Equal(t, 10, i)

	i = Chop(11, ints)
	assert.Equal(t, 11, i)
}
