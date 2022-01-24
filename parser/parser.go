package parser

import (
	"strings"

	"github.com/sharpvik/vish/ast"
)

type Parser struct {
	astree ast.AST
}

func New() *Parser {
	return &Parser{
		astree: ast.New(),
	}
}

func (p *Parser) Parse(input string) (astree ast.AST, err error) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		p.parseCommand(line)
	}
	return p.astree, nil
}

func (p *Parser) parseCommand(input string) {
	split := strings.Fields(input)
	if len(split) == 0 {
		return
	}
	p.astree = append(p.astree, ast.NewCommand(split[0], split[1:]))
}
