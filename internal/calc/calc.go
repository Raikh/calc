package calc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Calc(expression string) (float64, error) {
	// Remove all whitespace from the expression
	expression = strings.ReplaceAll(expression, " ", "")

	fmt.Println(expression)
	// Check if the expression is empty
	if len(expression) == 0 {
		return 0, errors.New("empty expression")
	}

	tokens, err := tokenize(expression)
	if err != nil {
		return 0, err
	}

	postfix, err := infixToPostfix(tokens)
	if err != nil {
		return 0, err
	}

	return evaluatePostfix(postfix)
}

func tokenize(expr string) ([]string, error) {
	var tokens []string
	var current string

	for i := 0; i < len(expr); i++ {
		ch := expr[i]
		if unicode.IsDigit(rune(ch)) || ch == '.' {
			current += string(ch)
		} else {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			tokens = append(tokens, string(ch))
		}
	}
	if current != "" {
		tokens = append(tokens, current)
	}
	return tokens, nil
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

func infixToPostfix(tokens []string) ([]string, error) {
	var output []string
	var stack []string

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			for len(stack) > 0 && precedence(stack[len(stack)-1]) >= precedence(token) {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		case "(":
			stack = append(stack, token)
		case ")":
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, errors.New("mismatched parentheses")
			}
			stack = stack[:len(stack)-1] // Pop the "("
		default: // Number
			output = append(output, token)
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return nil, errors.New("mismatched parentheses")
		}
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return output, nil
}

func evaluatePostfix(tokens []string) (float64, error) {
	var stack []float64

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0, errors.New("invalid expression")
			}
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				result = a / b
			}
			stack = append(stack, result)
		default:
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %s", token)
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}

	return stack[0], nil
}
