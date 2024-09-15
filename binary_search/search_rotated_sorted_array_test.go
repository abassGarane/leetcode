package binarysearch_test

import (
	"testing"

	binarysearch "github.com/abassGarane/leetcode/binary_search"
	"github.com/stretchr/testify/assert"
)

func TestSearchRotated(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		target int
		index  int
	}{
		{
			desc:   "first test",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			index:  4,
		},
		{
			desc:   "second test",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			index:  -1,
		},
		{
			desc:   "third test",
			nums:   []int{1},
			target: 0,
			index:  -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.index, binarysearch.SearchRotated(tC.nums, tC.target))
		})
	}
}
