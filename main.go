package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	jsonBytes, err := os.ReadFile("data.json")

	if err != nil {
		log.Fatal("Error reading json file")
	}

	jsonString := string(jsonBytes)
	// tokens array
	tokens := []Token{}

	t := &Tokenizer{
		input: jsonString,
		pos:   0,
		ch:    jsonString[0],
	}

	for {
		tok := t.NextToken()

		if tok.tokenType == EOF {
			break
		}

		tokens = append(tokens, tok)
	}

	p := NewParser(tokens)
	ast := p.Parse()

	result := Eval(ast)
	obj, ok := result.(map[string]any)

	if !ok {
		log.Fatal("Invalid object")
	}

	for key, _ := range obj {
		fmt.Println("key : ", key, " value : ", obj[key])
	}
}
