package lexer

import (
	"slices"
	"unicode"
)

type TokenType string

const EofRune = '\x00'

const (
	Illegal TokenType = "ILLEGAL"
	Eof     TokenType = "EOF"
	Ident   TokenType = "IDENT"
	Num     TokenType = "NUM"
	String  TokenType = "STRING"
	Symbol  TokenType = "SYMBOL"
	Comma   TokenType = "COMMA"
	Semi    TokenType = "SEMI"
	LParen  TokenType = "LPAREN"
	RParen  TokenType = "RPAREN"
	LCurly  TokenType = "LCURL"
	RCurly  TokenType = "RCURL"
	Comment TokenType = "COMMENT"
	Keyword TokenType = "KWD"
)

type Token struct {
	kind    TokenType
	literal string
}

type Tokenizer struct {
	pos   int
	rpos  int
	ch    rune
	input string
}

var Keywords = []string{
	"function",
	"class",
	"private",
	"public",
	"readonly",
	"var",
	"let",
	"const",
	"new",
	"catch",
	"try",
	"while",
	"for",
	"of",
	"in",
	"throw",
	"using",
	"as",
	"return",
}

func (tokenizer *Tokenizer) readChar() {
	if tokenizer.rpos >= len(tokenizer.input) {
		tokenizer.ch = EofRune
	} else {
		tokenizer.ch = rune(tokenizer.input[tokenizer.rpos])
	}

	tokenizer.pos = tokenizer.rpos
	tokenizer.rpos++
}

func (tokenizer *Tokenizer) NextToken() Token {
	tokenizer.skipWhitespaces()

	symbols := []rune{
		'+', '=', '>', '<', '-', '%', '&', '|', '^', '*', ':',
	}

	token := Token{
		Illegal,
		string(tokenizer.ch),
	}

	switch tokenizer.ch {
	case '{':
		token.kind = LCurly
	case '}':
		token.kind = RCurly
	case '(':
		token.kind = LParen
	case ')':
		token.kind = RParen
	case ',':
		token.kind = Comma
	case ';':
		token.kind = Semi
	case '\x00':
		token.kind = Eof
	}

	if slices.Contains(symbols, tokenizer.ch) {
		token.kind = Symbol
	} else if tokenizer.ch == '/' {
		next := tokenizer.peekNext()

		if next == '/' {
			token.kind = Comment
			token.literal = tokenizer.readSingleLineComment()
		} else if next == '*' {
			token.kind = Comment
			token.literal = tokenizer.readMultiLineComment()
		} else {
			token.kind = Symbol
		}
	} else if tokenizer.ch == '"' {
		token.kind = String
		token.literal = tokenizer.readString()

		return token
	} else if unicode.IsDigit(tokenizer.ch) {
		token.kind = Num
		token.literal = tokenizer.readNum()

		return token
	} else if unicode.IsLetter(tokenizer.ch) {
		token.literal = tokenizer.readIdent()

		isKwd := slices.Contains(Keywords, token.literal)
		if isKwd {
			token.kind = Keyword
		} else {
			token.kind = Ident
		}

		return token
	}

	tokenizer.readChar()

	return token
}

func (tokenizer *Tokenizer) skipWhitespaces() {
	for unicode.IsSpace(tokenizer.ch) {
		tokenizer.readChar()
	}
}

func (tokenizer *Tokenizer) peekNext() rune {
	if tokenizer.rpos < len(tokenizer.input) {
		return rune(tokenizer.input[tokenizer.rpos])
	}

	return EofRune
}

func (tokenizer *Tokenizer) readSingleLineComment() string {
	from := tokenizer.pos

	for tokenizer.ch != '\n' {
		tokenizer.readChar()
	}

	return tokenizer.input[from:tokenizer.pos]
}

func (tokenizer *Tokenizer) readMultiLineComment() string {
	from := tokenizer.pos

	for tokenizer.ch != '*' && tokenizer.peekNext() != '/' {
		tokenizer.readChar()
	}

	return tokenizer.input[from:tokenizer.pos]
}

func (tokenizer *Tokenizer) readNum() string {
	from := tokenizer.pos

	for unicode.IsDigit(tokenizer.ch) || tokenizer.ch == '_' || tokenizer.ch == '.' {
		tokenizer.readChar()
	}

	return tokenizer.input[from:tokenizer.pos]
}

func (tokenizer *Tokenizer) readString() string {
	from := tokenizer.pos
	tokenizer.readChar()

	for tokenizer.ch != '"' {
		tokenizer.readChar()
	}

	tokenizer.readChar()
	return tokenizer.input[from:tokenizer.pos]
}

func (tokenizer *Tokenizer) readIdent() string {
	from := tokenizer.pos

	for unicode.IsLetter(tokenizer.ch) || unicode.IsDigit(tokenizer.ch) || tokenizer.ch == '_' || tokenizer.ch == '$' || tokenizer.ch == '[' || tokenizer.ch == '.' || tokenizer.ch == ']' {
		tokenizer.readChar()
	}

	return tokenizer.input[from:tokenizer.pos]
}

func NewTokenizer(input string) Tokenizer {
	tokenizer := Tokenizer{
		pos:   0,
		rpos:  0,
		input: input,
		ch:    rune(input[0]),
	}

	tokenizer.readChar()

	return tokenizer
}
