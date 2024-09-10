package binarysearch_test

import (
	"testing"

	binarysearch "github.com/abassGarane/leetcode/binary_search"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		target int
		output int
	}{
		{
			desc:   "first search",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			output: 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.output, binarysearch.Search(tC.nums, tC.target))
		})
	}
}
