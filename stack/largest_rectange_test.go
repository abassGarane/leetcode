package stack_test

import (
	"github.com/abassGarane/leetcode/stack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	test_cases := []struct {
		heights  []int
		max_area int
	}{
		{
			heights:  []int{2, 4},
			max_area: 4,
		},
		{
			heights:  []int{2, 1, 5, 6, 2, 3},
			max_area: 10,
		},
	}

	for _, tC := range test_cases {
		res := stack.LargestRectangleArea(tC.heights)
		assert.Equal(t, tC.max_area, res)
	}
}
