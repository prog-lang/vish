package lexer

import (
	"io"
	"strings"

	"github.com/sharpvik/vish/token"
)

type Lexer struct {
	input    []rune
	position *Position
}

func New(stream io.Reader) (lex *Lexer, err error) {
	input, err := io.ReadAll(stream)
	if err != nil {
		return
	}
	return NewFromString(string(input)), nil
}

func NewFromString(input string) *Lexer {
	return &Lexer{
		input:    []rune(input),
		position: NewPosition(),
	}
}

func (lex *Lexer) Lex() (toks []*token.Token, err error) {
	for {
		tok, err := lex.Next()
		if err != nil {
			return nil, err
		}
		toks = append(toks, tok)
		if tok.Type == token.EOF {
			break
		}
	}
	return
}

func (lex *Lexer) Next() (tok *token.Token, err error) {
	return
}

func (lex *Lexer) EOF() bool {
	return lex.position.Index+1 >= len(lex.input)
}

func (lex *Lexer) eatWhile(ok Condition) (ate string, n int) {
	var builder strings.Builder
	for {
		one, err := lex.eatIf(ok)
		if err != nil {
			ate = builder.String()
			break
		}
		builder.WriteRune(one)
		n++
	}
	return
}

func (lex *Lexer) eatIf(ok Condition) (one rune, err error) {
	one, err = lex.peek()
	if err != nil {
		return
	}
	if ok(one) {
		lex.position.Receive(one)
		return
	}
	return rune(0), ErrNoMatch
}

func (lex *Lexer) eat() (one rune, err error) {
	one, err = lex.peek()
	if err != nil {
		return
	}
	lex.position.Receive(one)
	return
}

func (lex *Lexer) peek() (one rune, err error) {
	if lex.EOF() {
		return 0, io.EOF
	}
	return lex.input[lex.position.Index+1], nil
}
