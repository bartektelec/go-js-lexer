package parser

type NodeKind string

const (
	Program        NodeKind = "Program"
	VarDeclaration NodeKind = "VarDeclaration"
)

type IAstNode interface {
	ToString() string
}

type AstNode struct {
	kind NodeKind
}
