package main

// creating an Abstract Syntax Tree

type Node struct {
	val  string
	next *Node
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

type NullNode struct {
	val string
}

type ArrayNode struct {
	values []Node
}

type ObjectNode struct {
	properties map[string]Node
}
