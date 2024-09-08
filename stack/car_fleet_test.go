package stack_test

import (
	"testing"

	"github.com/abassGarane/leetcode/stack"
	"github.com/stretchr/testify/assert"
)

func TestCarFleet(t *testing.T) {
	testCases := []struct {
		desc     string
		target   int
		position []int
		speed    []int
		expected int
	}{
		{
			desc:     "when target is 12",
			target:   12,
			position: []int{10, 8, 0, 5, 3},
			speed:    []int{2, 4, 1, 1, 3},
			expected: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.expected, stack.CarFleet(tC.target, tC.position, tC.speed))
		})
	}
}
