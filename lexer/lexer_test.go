package lexer

import (
	"testing"

	"github.com/udzura/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType token.TokenType
		expectedLit  string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for _, tt := range tests {
		var tok token.Token
		tok = l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tokentype exp:%q act:%q", tt.expectedType, tok.Type)
		}

		if tok.Type != tt.expectedType {
			t.Fatalf("tokentype exp:%q act:%q", tt.expectedType, tok.Type)
		}
	}
}

func testSpecifiedToken(t *testing.T, input string, size int) {
	res := make([]token.Token, 0)

	l := New(input)

	for {
		var tok token.Token
		tok = l.NextToken()

		res = append(res, tok)
		if tok.Type == token.ILLIGAL {
			t.Fatalf("expect not to contain illegal token: %v", res)
		}

		if tok.Type == token.EOF {
			break
		}
	}

	t.Logf("Parsed: %v\n", res)

	if len(res) != size {
		t.Fatalf("expect token size %d, got %d", 99, len(res))
	}
}

func TestRealToken(t *testing.T) {
	input := `
let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
`
	testSpecifiedToken(t, input, 37)
}

func TestExtendedToken(t *testing.T) {
	input := `
!-/*5;
5 < 10 > 5;
return 123;
`
	testSpecifiedToken(t, input, 16)
}
