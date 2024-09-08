package stack_test

import (
	"slices"
	"testing"

	"github.com/abassGarane/leetcode/stack"
)

func TestDailyTemperatures(t *testing.T) {
	testCases := []struct {
		desc         string
		temperatures []int
		expected     []int
	}{
		{
			desc:         "First test",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected:     []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := stack.DailyTemperatures(tC.temperatures)
			if !slices.Equal(res, tC.expected) {
				t.Errorf("wanted %v, got %v", tC.expected, res)
			}
		})
	}
}
