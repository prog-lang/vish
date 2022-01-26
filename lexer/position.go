package lexer

type Position struct {
	Index  int
	Line   int
	Column int
}

func NewPosition() *Position {
	return &Position{
		Index:  -1,
		Line:   1,
		Column: 0,
	}
}

func (pos *Position) Receive(one rune) {
	if one == '\n' {
		pos.NewLine()
	} else {
		pos.Advance()
	}
}

func (pos *Position) Advance() {
	pos.Index++
	pos.Column++
}

func (pos *Position) NewLine() {
	pos.Index++
	pos.Line++
	pos.Column = 0
}
