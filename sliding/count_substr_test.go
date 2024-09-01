package sliding_test

import (
	"testing"

	"github.com/abassGarane/leetcode/sliding"
)

type testCase2 struct {
	sub_str  string
	k        int
	expected int
}

func TestCountKConstraintSubstrings(t *testing.T) {
	test_cases := []*testCase2{
		{
			sub_str:  "1010101",
			k:        2,
			expected: 25,
		},
		{
			sub_str:  "11111",
			k:        1,
			expected: 15,
		},
	}
	for _, v := range test_cases {
		res := sliding.CountKConstraintSubstrings(v.sub_str, v.k)
		if v.expected != res {
			t.Errorf("expected %v got %v", v.expected, res)
		}

	}
}
