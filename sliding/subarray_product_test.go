package sliding_test

import (
	"testing"

	"github.com/abassGarane/leetcode/sliding"
)

type testCase struct {
	input []int
	k     int
	want  int
}

func TestNumSubarrayProductLessThanK(t *testing.T) {
	test_cases := []*testCase{
		{
			input: []int{10, 5, 2, 6},
			k:     100,
			want:  8,
		},
		{
			input: []int{},
			k:     10,
			want:  0,
		},
		{
			input: []int{1, 2, 3},
			k:     0,
			want:  0,
		},
	}
	for _, tc := range test_cases {
		got := sliding.NumSubarrayProductLessThanK(tc.input, tc.k)
		if got != tc.want {
			t.Errorf("expected %d got %d", tc.want, got)
		}

	}
}
