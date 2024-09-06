package stack

import "strconv"

// You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

// Evaluate the expression. Return an integer that represents the value of the expression.

// Note that:

// The valid operators are '+', '-', '*', and '/'.
// Each operand may be an integer or another expression.
// The division between two integers always truncates toward zero.
// There will not be any division by zero.
// The input represents a valid arithmetic expression in a reverse polish notation.
// The answer and all the intermediate calculations can be represented in a 32-bit integer.

func EvalRPN(tokens []string) int {
	stack := Stack{}
	for _, c := range tokens {
		if c == "*" || c == "/" || c == "-" || c == "+" {
			first, _ := stack.Pop()
			second, _ := stack.Pop()
			result := 0
			if c == "*" {
				result = first * second
			} else if c == "/" {
				result = second / first
			} else if c == "-" {
				result = second - first
			} else {
				result = second + first
			}
			stack.Push(result)
		} else {
			num, _ := strconv.Atoi(c)
			stack.Push(num)
		}
	}
	res, _ := stack.Peek()
	return res
}
