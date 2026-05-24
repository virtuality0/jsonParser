package main

import "strconv"

// creating an Abstract Syntax Tree

type Node interface {
	node()
}

type StringNode struct {
	val string
}

type NumberNode struct {
	val float64
}

type BoolNode struct {
	val bool
}

type NullNode struct{}

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

func (*StringNode) node() {}
func (*ObjectNode) node() {}
func (*NumberNode) node() {}
func (*NullNode) node()   {}
func (*BoolNode) node()   {}
func (*ArrayNode) node()  {}

func (p *Parser) current() Token {
	if p.pos < len(p.tokens) {
		return p.tokens[p.pos]
	}
	return Token{tokenType: EOF, value: ""}
}

func (p *Parser) advance() {
	p.pos++
}

func (p *Parser) peek() Token {
	next := p.pos + 1

	if next < len(p.tokens) {
		return p.tokens[next]
	}

	return Token{
		tokenType: EOF,
		value:     "",
	}
}

func (p *Parser) parseString() *StringNode {
	node := &StringNode{p.current().value}
	p.advance()
	return node
}

func (p *Parser) parseNumber() *NumberNode {
	num, err := strconv.ParseFloat(p.current().value, 64)
	if err != nil {
		panic("Error converting string to float")
	}

	p.advance()
	return &NumberNode{num}
}

func (p *Parser) parseBool() *BoolNode {
	value := p.current().value == "true"
	p.advance()
	return &BoolNode{value}
}

func (p *Parser) parseNull() *NullNode {
	p.advance()
	return &NullNode{}
}

func (p *Parser) parseArray() *ArrayNode {
	// consume [
	p.advance()
	arr := &ArrayNode{}

	// empty array check
	if p.current().tokenType == CLOSE_BRACKET {
		p.advance()
		return arr
	}

	for {
		value := p.parseValue()
		arr.values = append(arr.values, value)

		if p.current().tokenType == COMMA {
			p.advance()
			continue
		}

		if p.current().tokenType == CLOSE_BRACKET {
			break
		}

		panic("Expected comma or closing brace")
	}

	// consume ]
	p.advance()

	return arr
}

func (p *Parser) parseObject() *ObjectNode {
	// consume {
	p.advance()

	obj := &ObjectNode{
		properties: make(map[string]Node),
	}

	// empty object check
	if p.current().tokenType == CLOSE_BRACE {
		p.advance()
		return obj
	}

	for {
		if p.current().tokenType != STRING {
			panic("Expected string key")
		}

		key := p.current().value
		p.advance()

		if p.current().tokenType != COLON {
			panic("Expected :")
		}

		// consume :
		p.advance()

		value := p.parseValue()
		obj.properties[key] = value

		if p.current().tokenType == COMMA {
			p.advance()
			continue
		}

		if p.current().tokenType == CLOSE_BRACE {
			break
		}

		panic("Expected comma or closing brace")
	}

	// comsume }
	p.advance()
	return obj
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

	case OPEN_BRACE:
		return p.parseObject()

	default:
		panic("Invalid token found : ")
	}
}

func (p *Parser) Parse() Node {
	val := p.parseValue()
	if p.current().tokenType == EOF {
		panic("unexpected tokens after json value")
	}

	return val
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
		pos:    0,
	}
}
