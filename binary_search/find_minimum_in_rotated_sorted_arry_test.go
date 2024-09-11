package binarysearch_test

import (
	"testing"

	binarysearch "github.com/abassGarane/leetcode/binary_search"
	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	testCases := []struct {
		desc    string
		nums    []int
		minimum int
	}{
		{
			desc:    "first test",
			nums:    []int{3, 4, 5, 1, 2},
			minimum: 1,
		},
		{
			desc:    "second test",
			nums:    []int{4, 5, 6, 7, 0, 1, 2},
			minimum: 0,
		},
		{
			desc:    "third test",
			nums:    []int{11, 13, 15, 17},
			minimum: 11,
		},
		{
			desc:    "fourth test",
			nums:    []int{1},
			minimum: 1,
		},
		{
			desc:    "fifth test",
			nums:    []int{2, 1},
			minimum: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.minimum, binarysearch.FindMin(tC.nums))
		})
	}
}
