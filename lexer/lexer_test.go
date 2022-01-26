package lexer

import (
	"io"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"

	"github.com/sharpvik/vish/token"
)

func TestLexer(t *testing.T) {
	testEat(t)
	testEatIf(t)
	testEatWhile(t)

	t.SkipNow()
	assertLexerOutput(
		t, "name", []*token.Token{token.New(token.Name, "name")})
}

func testEat(t *testing.T) {
	lex := NewFromString("h\n")

	one, err := lex.eat()
	assert.NoError(t, err)
	assert.Equal(t, 'h', one)
	assert.Equal(t, 0, lex.position.Index)
	assert.Equal(t, 1, lex.position.Line)
	assert.Equal(t, 1, lex.position.Column)

	one, err = lex.eat()
	assert.NoError(t, err)
	assert.Equal(t, '\n', one)
	assert.Equal(t, 1, lex.position.Index)
	assert.Equal(t, 2, lex.position.Line)
	assert.Equal(t, 0, lex.position.Column)

	one, err = lex.eat()
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, rune(0), one)
	assert.Equal(t, 1, lex.position.Index)
	assert.Equal(t, 2, lex.position.Line)
	assert.Equal(t, 0, lex.position.Column)
}

func testEatIf(t *testing.T) {
	lex := NewFromString("h\n")

	one, err := lex.eatIf(unicode.IsLetter)
	assert.NoError(t, err)
	assert.Equal(t, 'h', one)
	assert.Equal(t, 0, lex.position.Index)
	assert.Equal(t, 1, lex.position.Line)
	assert.Equal(t, 1, lex.position.Column)

	one, err = lex.eatIf(unicode.IsLetter)
	assert.Equal(t, ErrNoMatch, err)
	assert.Equal(t, rune(0), one)

	one, err = lex.eatIf(unicode.IsSpace)
	assert.NoError(t, err)
	assert.Equal(t, '\n', one)
	assert.Equal(t, 1, lex.position.Index)
	assert.Equal(t, 2, lex.position.Line)
	assert.Equal(t, 0, lex.position.Column)
}

func testEatWhile(t *testing.T) {
	lex := NewFromString("hello\n")

	ate, n := lex.eatWhile(unicode.IsSpace)
	assert.Equal(t, 0, n)
	assert.Equal(t, "", ate)
	assert.Equal(t, -1, lex.position.Index)
	assert.Equal(t, 1, lex.position.Line)
	assert.Equal(t, 0, lex.position.Column)

	ate, n = lex.eatWhile(unicode.IsLetter)
	assert.Equal(t, 5, n)
	assert.Equal(t, "hello", ate)
	assert.Equal(t, 4, lex.position.Index)
	assert.Equal(t, 1, lex.position.Line)
	assert.Equal(t, 5, lex.position.Column)
}

func assertLexerOutput(
	t *testing.T,
	input string,
	expect []*token.Token,
) {
	result, err := NewFromString(input).Lex()
	assert.NoError(t, err)
	assert.Equal(t, expect, result)
}
