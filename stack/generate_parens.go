package stack

import "strings"

func GenerateParenthesis(n int) []string {
	parens := []string{}
	str_stack := []string{}
	backtrack(0, 0, n, str_stack, &parens)
	return parens
}

func backtrack(left, right int, n int, str_stack []string, parens *[]string) {
	if left == n && right == n {
		str := strings.Join(str_stack, "")
		*parens = append(*parens, str)
		return
	}

	if left < n || left == 0 {
		str_stack = append(str_stack, "(")
		backtrack(left+1, right, n, str_stack, parens)
		str_stack = str_stack[:len(str_stack)-1]
	}
	if right < left {
		str_stack = append(str_stack, ")")
		backtrack(left, right+1, n, str_stack, parens)
	}
}
func GenerateParenthesisAlt(n int) []string {
	result := []string{}

	var generate func(left, right int, current string)
	generate = func(left, right int, current string) {
		if left < n {
			generate(left+1, right, current+"(")
		}

		if right < left {
			generate(left, right+1, current+")")
		}

		if right == n {
			result = append(result, current)
		}
	}

	generate(0, 0, "")
	return result

}
