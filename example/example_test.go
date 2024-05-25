package enum_test

import (
	"testing"

	"github.com/PeerXu/enum-go"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	getByName := enum.GenerateGetByName(A, C)

	a, ok := getByName("A")
	assert.Equal(t, true, ok)
	assert.Equal(t, A, a)

	b, ok := getByName("B")
	assert.Equal(t, true, ok)
	assert.Equal(t, B, b)

	c, ok := getByName("C")
	assert.Equal(t, true, ok)
	assert.Equal(t, C, c)
}
