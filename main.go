package main

import (
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

	for _, t := range tokens {
		println(t.tokenType, " ", t.value)
	}
}
