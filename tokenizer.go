package main

import (
	"fmt"
	"unicode"
)

// different possible tokens

type TokenType int

const (
	OPEN_BRACE TokenType = iota
	CLOSE_BRACE
	COLON
	QUOTES
	COMMA
	OPEN_BRACKET
	CLOSE_BRACKET
	STRING
	NUMBER
	NULL
	TRUE
	FALSE
	EOF
)

type Token struct {
	tokenType TokenType
	value     string
}

type Tokenizer struct {
	input string
	pos   int
	ch    byte
}

// moves your cursor to the next token
func (t *Tokenizer) NextChar() {
	fmt.Println("Position : ", t.pos, " char : ", string(t.ch))
	t.pos++
	if t.pos >= len(t.input) {
		t.ch = 0
	} else {
		t.ch = t.input[t.pos]
	}
}

func (t *Tokenizer) ReadLiteral() string {
	start := t.pos
	for IsLetter(t.ch) {
		t.NextChar()
	}

	return t.input[start:t.pos]
}

// skips the whitespaces
func (t *Tokenizer) SkipWhitespace() {
	for t.ch == '\n' || t.ch == '\t' || t.ch == ' ' || t.ch == '\r' {
		t.NextChar()
	}
}

// TO-DO : escape char and unicode support
func (t *Tokenizer) ReadString() string {
	start := t.pos + 1 // skipping starting "
	for {
		t.NextChar()
		if t.ch == 0 || t.ch == '"' {
			break
		}
	}

	return t.input[start:t.pos]
}

// TO-DO : complex, negative and exponent nums support
func (t *Tokenizer) ReadNumber() string {
	start := t.pos

	for unicode.IsDigit(rune(t.ch)) {
		t.NextChar()
	}

	return t.input[start:t.pos]
}

func (t *Tokenizer) NextToken() Token {
	t.SkipWhitespace()
	tok := Token{}

	switch t.ch {
	case '{':
		tok = Token{tokenType: OPEN_BRACE, value: "{"}

	case '}':
		tok = Token{tokenType: CLOSE_BRACE, value: "}"}

	case '[':
		tok = Token{tokenType: OPEN_BRACKET, value: "["}

	case ']':
		tok = Token{tokenType: CLOSE_BRACKET, value: "]"}

	case ':':
		tok = Token{tokenType: COLON, value: ":"}

	case ',':
		tok = Token{tokenType: COMMA, value: ","}

	case '"':
		str := t.ReadString()
		tok = Token{tokenType: CLOSE_BRACE, value: str}

	default:
		tok = Token{tokenType: EOF, value: ""}
	}

	if unicode.IsDigit(rune(t.ch)) {
		return Token{tokenType: NUMBER, value: t.ReadNumber()}
	}

	if IsLetter(t.ch) {
		str := t.ReadLiteral()

		println(str)
		switch str {
		case "true":
			return Token{tokenType: TRUE, value: "true"}

		case "false":
			return Token{tokenType: FALSE, value: "false"}

		case "null":
			return Token{tokenType: NULL, value: "null"}
		}
	}

	t.NextChar()
	return tok
}
