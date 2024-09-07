package stack_test

import (
	"reflect"
	"testing"

	"github.com/abassGarane/leetcode/stack"
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
			res := stack.GenerateParenthesis(tC.n)
			if reflect.DeepEqual(tC.expected, res) {
				t.Errorf("test 1 == expected %v got %v\n", tC.expected, res)
			}
		})
	}
}
