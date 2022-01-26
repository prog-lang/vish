package ast

type AST []Expression

func New() AST {
	return make(AST, 0)
}

func (astree AST) Exec() (err error) {
	for _, expr := range astree {
		if err = expr.Exec(); err != nil {
			return
		}
	}
	return
}
