package lexer

import (
	"testing"
)

func TestTokenizerIllegal(t *testing.T) {
	input := "=+(){},;"

	lexer := NewTokenizer(input)

	tokens := []TokenType{
		Symbol,
		Symbol,
		LParen,
		RParen,
		LCurly,
		RCurly,
		Comma,
		Semi,
	}

	for _, token := range tokens {
		next := lexer.NextToken()

		if next.kind != token {
			t.Errorf(
				"Token %v expected %v", next.kind, token)
		}

	}
}

func TestTokenizer(t *testing.T) {
	input := `let name = "hello";

  const ten = 10;
  function add(x, y) {
  // this is a comment
    return x + y
  }

  var res += add(10, 1)

  for(let i = 10; i < arr.length; i++) {
    console.log(arr[i])
  }
  `

	lexer := NewTokenizer(input)

	tokens := []Token{
		{Keyword, "let"},
		{Ident, "name"},
		{Symbol, "="},
		{String, "\"hello\""},
		{Semi, ";"},

		{Keyword, "const"},
		{Ident, "ten"},
		{Symbol, "="},
		{Num, "10"},
		{Semi, ";"},

		{Keyword, "function"},
		{Ident, "add"},
		{LParen, "("},
		{Ident, "a"},
		{Comma, ","},
		{Ident, "b"},
		{RParen, ")"},
		{LCurly, "{"},
		{Comment, " this is a comment"},
		{Keyword, "return"},
		{Ident, "x"},
		{Symbol, "+"},
		{Ident, "y"},
		{RCurly, "}"},

		{Keyword, "var"},
		{Ident, "res"},
		{Symbol, "+"},
		{Symbol, "="},
		{Ident, "add"},
		{LParen, "("},
		{Num, "10"},
		{Comma, ","},
		{Num, "1"},
		{RParen, ")"},
		//
		// for(let i = 10; i < arr.length; i++) {
		//   console.log(arr[i])
		// }

		{Keyword, "for"},
		{LParen, "("},
		{Keyword, "let"},
		{Ident, "i"},
		{Symbol, "="},
		{Num, "10"},
		{Semi, ";"},
		{Ident, "i"},
		{Symbol, ">"},
		{Ident, "arr.length"},
		{Semi, ";"},
		{Ident, "i"},
		{Symbol, "+"},
		{Symbol, "+"},
		{RParen, ")"},
		{LCurly, "{"},
		{Ident, "console.log"},
		{LParen, "("},
		{Ident, "arr[i]"},
		{RParen, ")"},
		{RCurly, "}"},
	}

	for _, token := range tokens {
		next := lexer.NextToken()

		if next.kind != token.kind && next.literal != token.literal {
			t.Errorf("Token %v:%v expected %v:%v", next.kind, next.literal, token.kind, token.literal)
		}
	}
}
