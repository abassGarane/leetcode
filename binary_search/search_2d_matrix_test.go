package binarysearch_test

import (
	"testing"

	binarysearch "github.com/abassGarane/leetcode/binary_search"
	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	testCases := []struct {
		desc    string
		matrix  [][]int
		target  int
		present bool
	}{
		{
			desc:    "first test",
			matrix:  [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:  13,
			present: false,
		},
		{
			desc:    "second test",
			matrix:  [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:  3,
			present: true,
		},
		{
			desc:    "third test",
			matrix:  [][]int{{1}, {3}},
			target:  4,
			present: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.present, binarysearch.SearchMatrixAlt(tC.matrix, tC.target))
		})
	}
}
