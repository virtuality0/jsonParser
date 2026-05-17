package main

import "strconv"

// creating an Abstract Syntax Tree

type Node interface{}

type StringNode struct {
	val string
}

type NumberNode struct {
	val float64
}

type BoolNode struct {
	val bool
}

type NullNode struct {
	val string
}

type ArrayNode struct {
	values []Node
}

type ObjectNode struct {
	properties map[string]Node
}

type Parser struct {
	tokens []Token
	pos    int
}

func (p *Parser) current() Token {
	if p.pos <= len(p.tokens) {
		return p.tokens[p.pos]
	}
	return Token{tokenType: EOF, value: ""}
}

func (p *Parser) advance() {
	p.pos++
}

func (p *Parser) parseString() *StringNode {
	return &StringNode{p.current().value}
}

func (p *Parser) parseNumber() *NumberNode {
	num, err := strconv.ParseFloat(p.current().value, 64)
	if err != nil {
		panic("Error converting string to float")
	}
	return &NumberNode{num}
}

func (p *Parser) parseBool() *BoolNode {
	value := If(p.current().value == "true", true, false)
	return &BoolNode{value}
}

func (p *Parser) parseNull() *NullNode {
	return &NullNode{}
}

func (p *Parser) parseArray() *ArrayNode {
	curr := p.current()
	aNode := &ArrayNode{}

	for curr.tokenType != CLOSE_BRACKET {
		t := p.current()
		if t.tokenType == EOF {
			panic("Error ! Couldn't find array end")
		}
		aNode.values = append(aNode.values, p.parseValue())
		p.advance()
	}

	return aNode
}

func (p *Parser) parseObject() {
	curr := p.current()

	for curr.tokenType != CLOSE_BRACE {

		p.advance()
	}
}

func (p *Parser) parseValue() Node {
	tok := p.current()

	switch tok.tokenType {
	case TRUE, FALSE:
		return p.parseBool()

	case NULL:
		return p.parseNull()

	case NUMBER:
		return p.parseNumber()

	case STRING:
		return p.parseString()

	case OPEN_BRACKET:
		return p.parseArray()

	default:
		panic("Invalid token found : ")
	}
}

func (p *Parser) Parse() Node {
	return p.parseValue()
}
