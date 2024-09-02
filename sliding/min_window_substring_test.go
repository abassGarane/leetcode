package sliding_test

import (
	"testing"

	"github.com/abassGarane/leetcode/sliding"
)

type testCase3 struct {
	s        string
	t        string
	expected string
}

func TestMinWindow(t *testing.T) {
	var test_cases = []*testCase3{
		{
			s:        "ADOBECODEBANC",
			t:        "ABC",
			expected: "BANC",
		},
		{
			s:        "a",
			t:        "a",
			expected: "a",
		},
	}
	for _, v := range test_cases {
		res := sliding.MinWindow(v.s, v.t)
		if res != v.expected {
			t.Errorf("expected %s got %s", v.expected, res)
		}
	}
}
