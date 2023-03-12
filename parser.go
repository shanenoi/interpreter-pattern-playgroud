package main

import (
	"log"
	"strconv"
)

// parseExpression parses the tokens and returns an expression
func parseExpression(tokens []token) Expression {
	tokens = groupBrackets(tokens)
	tokens = groupMulAndDiv(tokens)
	tokens = groupAddAndSub(tokens)

	if len(tokens) != 1 {
		return parseExpression(tokens)
	}
	return getExpression(tokens[0])
}

type token interface{}

// getExpression returns an expression from a token
func getExpression(token token) Expression {
	switch rv := token.(type) {
	case Expression:
		return rv
	case string:
		value, err := strconv.ParseFloat(token.(string), 64)
		if err != nil {
			log.Panicln(err)
		}
		return &Number{value}
	}

	return nil
}

// replace replaces a slice of tokens from head to tail with new tokens
func replace(tokens []token, head int, tail int, newTokens ...token) []token {
	updatedTokens := append(tokens[:head], newTokens...)
	updatedTokens = append(updatedTokens, tokens[tail+1:]...)
	return updatedTokens
}

// groupMulAndDiv groups the tokens that are multiplication or division operations
func groupMulAndDiv(tokens []token) []token {
	idx := 0
	for idx < len(tokens) {
		token := tokens[idx]
		// check if token is a string
		if _, ok := token.(string); !ok {
			idx++
			continue
		}

		if token == "*" {
			tokens = replace(
				tokens, idx-1, idx+1,
				&Multiply{
					getExpression(tokens[idx-1]),
					getExpression(tokens[idx+1]),
				},
			)
			continue
		}

		if token == "/" {
			tokens = replace(
				tokens, idx-1, idx+1,
				&Divide{
					getExpression(tokens[idx-1]),
					getExpression(tokens[idx+1]),
				},
			)
			continue
		}

		idx++
	}

	return tokens
}

// groupAddAndSub groups the tokens that are addition or subtraction operations
func groupAddAndSub(tokens []token) []token {
	idx := 0
	for idx < len(tokens) {
		token := tokens[idx]
		// check if token is a string
		if _, ok := token.(string); !ok {
			idx++
			continue
		}

		if token == "+" {
			tokens = replace(
				tokens, idx-1, idx+1,
				&Add{
					getExpression(tokens[idx-1]),
					getExpression(tokens[idx+1]),
				},
			)
			continue
		}

		if token == "-" {
			tokens = replace(
				tokens, idx-1, idx+1,
				&Subtract{
					getExpression(tokens[idx-1]),
					getExpression(tokens[idx+1]),
				},
			)
			continue
		}

		idx++
	}

	return tokens
}

// groupBrackets groups the tokens that are inside brackets
func groupBrackets(tokens []token) []token {
	head := 0
	tail := 0

	var open []int
	for idx, token := range tokens {
		// check if token is a string
		if _, ok := token.(string); !ok {
			continue
		}

		if token == "(" {
			open = append(open, idx)
		}

		if token == ")" {
			head = open[len(open)-1] // get the last element of open
			tail = idx
			break
		}
	}

	if head == 0 && tail == 0 {
		return tokens
	}

	_tokens := tokens[head+1 : tail]
	_tokens = groupMulAndDiv(_tokens)
	_tokens = groupAddAndSub(_tokens)
	tokens = replace(tokens, head, tail, _tokens...)

	return groupBrackets(tokens)
}
