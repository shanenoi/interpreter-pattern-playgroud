package main

import (
	"log"
	"strings"
)

func main() {
	exp := strings.Split("( 1 + 2 ) / 3 * 4", " ")
	var tokens []token
	for _, s := range exp {
		tokens = append(tokens, s)
	}

	expression := parseExpression(tokens)
	/*
		expression = Multiply(
			Divide(
				Add(Number(1), Number(2)),
				Number(3),
			),
			Number(4),
		)
	*/

	log.Println(expression.Evaluate())
}
