package main

import (
	"log"
	"strings"
)

func main() {
	exp := strings.Split("( ( 1 + 2 ) * ( 3 - 4 ) ) / 2 + 20 - ( 0.5 * 9 )", " ")
	var tokens []token
	for _, s := range exp {
		tokens = append(tokens, s)
	}

	expression := parseExpression(tokens)
	log.Println(expression.Evaluate())
}
