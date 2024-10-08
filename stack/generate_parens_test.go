package stack_test

import (
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	testCases := []struct {
		desc     string
		n        int
		expected []string
	}{
		{
			desc:     "Generate parenthesis when input is three",
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			desc:     "Generate parens when n is 1",
			n:        1,
			expected: []string{"()"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
		})
	}
}
