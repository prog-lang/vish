package ast

type AST []*Command

func New() AST {
	return make(AST, 0)
}
