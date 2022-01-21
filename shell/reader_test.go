package vish

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	r := NewReader(strings.NewReader("hello\\\nworld"))
	input, err := r.Next()
	assert.NoError(t, err)
	assert.Equal(t, "hello world", input)

	r = NewReader(strings.NewReader("hello\\\nworld\n"))
	input, err = r.Next()
	assert.NoError(t, err)
	assert.Equal(t, "hello world", input)
}
