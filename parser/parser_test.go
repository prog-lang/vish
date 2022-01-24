package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sharpvik/vish/ast"
)

func TestParser(t *testing.T) {
	assertParserOutput(t, "cd ..\necho hello", ast.AST{
		ast.NewCommand("cd", []string{".."}),
		ast.NewCommand("echo", []string{"hello"}),
	})
}

func assertParserOutput(
	t *testing.T,
	input string,
	expect ast.AST,
) {
	result, err := New().Parse(input)
	assert.NoError(t, err)
	assert.Equal(t, expect, result)
}
