package stack

func IsValid(s string) bool {

	if len(s)%2 == 1 || len(s) == 0 {
		return false
	}
	hash_map := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	stack := []rune{}

	for _, c := range s {
		if _, ok := hash_map[c]; ok {
			stack = append(stack, c)
		} else if len(stack) == 0 || hash_map[stack[len(stack)-1]] != c {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
