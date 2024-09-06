package stack_test

import "testing"

func TestEvalRPN(t *testing.T) {
	testCases := []struct {
		desc     string
		token    []string
		expected int
	}{
		{
			desc:     "Testing reverse polish notation",
			token:    []string{"2", "1", "+", "3", "*"},
			expected: 9,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
