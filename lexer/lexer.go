package lexer

import (
	"io"

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

func (lex *Lexer) eatIf(ok func(rune) bool) (one rune, err error) {
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
