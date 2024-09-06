package stack_test

import (
	"testing"

	"github.com/abassGarane/leetcode/stack"
)

type isValidTestCases struct {
	s     string
	valid bool
}

func TestIsValid(t *testing.T) {
	var test_cases = []*isValidTestCases{
		{
			s:     "()",
			valid: true,
		},
		{
			s:     "()[]{}",
			valid: true,
		},
		{
			s:     "([)",
			valid: false,
		},
		{
			s:     "){",
			valid: false,
		},
	}
	for _, tc := range test_cases {
		res := stack.IsValid(tc.s)
		if res != tc.valid {
			t.Errorf("expected %v got %v", tc.valid, res)
		}
	}
}
