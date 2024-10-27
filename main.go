package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	expression = strings.Replace(expression, " ", "", -1)
	if expression == "" {
		return 0, errors.New("пустая строка")
	}

	expression_postfix, err := toPostfix(expression)
	if err != nil {
		return 0, err
	}

	return evalPostfix(expression_postfix)
}

func toPostfix(expression string) ([]string, error) {
	var expression_postfix []string
	var stack []rune

	priority := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'(': 0,
	}

	var number string

	for _, char := range expression {
		if char >= '0' && char <= '9' || char == '.' {
			number += string(char)
		} else {
			if len(number) > 0 {
				expression_postfix = append(expression_postfix, number)
				number = ""
			}

			if char == '(' {
				stack = append(stack, char)
			} else if char == ')' {
				for len(stack) > 0 && stack[len(stack)-1] != '(' {
					expression_postfix = append(expression_postfix, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				if len(stack) == 0 {
					return nil, errors.New("неправильная последовательность скобок")
				}
				stack = stack[:len(stack)-1]
			} else if char == '+' || char == '-' || char == '*' || char == '/' {
				for len(stack) > 0 && priority[char] <= priority[stack[len(stack)-1]] {
					expression_postfix = append(expression_postfix, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, char)
			} else {
				return nil, errors.New("неизвестный символ")
			}
		}
	}

	if len(number) > 0 {
		expression_postfix = append(expression_postfix, number)
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == '(' {
			return nil, errors.New("неправильная последовательность скобок")
		}
		expression_postfix = append(expression_postfix, string(stack[len(stack)-1]))
		stack = stack[:len(stack)-1]
	}

	return expression_postfix, nil
}

func evalPostfix(expression_postfix []string) (float64, error) {
	var stack []float64

	for _, char := range expression_postfix {
		if char == "+" || char == "-" || char == "*" || char == "/" {
			if len(stack) < 2 {
				return 0, errors.New("неправильное выражение")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch char {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, errors.New("деление на ноль")
				}
				stack = append(stack, a/b)
			}
		} else {
			num, err := strconv.ParseFloat(char, 64)
			if err != nil {
				return 0, errors.New("неизвестный символ")
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("неправильное выражение")
	}

	return stack[0], nil
}

func main() {
	fmt.Println(Calc("(3 + 6) * 2 / 0.5"))
}
