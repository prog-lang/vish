package vish

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	assertReaderOutput(t, "hello\\\nworld", "hello world")
	assertReaderOutput(t, "hello\\\nworld\n", "hello world")
}

func assertReaderOutput(
	t *testing.T,
	input string,
	expect string,
) {
	r := NewReader(strings.NewReader(input))
	actual, err := r.Next()
	assert.NoError(t, err)
	assert.Equal(t, expect, actual)
}
