package main

func evalArray(n *ArrayNode) []any {
	arr := []any{}

	for _, val := range n.Values {
		arr = append(arr, Eval(val))
	}

	return arr
}

func evalObject(obj *ObjectNode) map[string]any {
	result := map[string]any{}

	for key, value := range obj.Properties {
		result[key] = Eval(value)
	}

	return result
}

func Eval(node Node) any {
	switch n := node.(type) {
	case *StringNode:
		return n.Value

	case *NumberNode:
		return n.Value

	case *NullNode:
		return nil

	case *BoolNode:
		return n.Value

	case *ArrayNode:
		return evalArray(n)

	case *ObjectNode:
		return evalObject(n)
	}

	panic("Invalid node")
}
