package token

type Token struct {
	Type  Type
	Value interface{}
}

type Type int

const (
	EOF Type = iota
	Name
	String
)

func New(t Type, v interface{}) *Token {
	return &Token{
		Type:  t,
		Value: v,
	}
}

func NewEOF() *Token {
	return New(EOF, nil)
}
