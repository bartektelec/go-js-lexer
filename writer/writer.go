package writer

import "github.com/bartektelec/go-js-lexer/lexer"

type TokenWriter struct {
	indent  int
	pos     int
	tokens  []lexer.Token
	current *lexer.Token
}

func (w *TokenWriter) Gen() string {
	return ""
}

func (w *TokenWriter) Step() {
}

func (w *TokenWriter) PeekBack() *lexer.Token {
	if w.pos-1 > 0 {
		return &w.tokens[w.pos-1]
	}

	return nil
}

func (w *TokenWriter) PeekNext() *lexer.Token {
	if w.pos+1 < len(w.tokens) {
		return &w.tokens[w.pos+1]
	}

	return nil
}
