package binarysearch_test

import (
	"testing"

	binarysearch "github.com/abassGarane/leetcode/binary_search"
	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	testCases := []struct {
		desc   string
		num1   []int
		num2   []int
		medium float64
	}{
		{
			desc:   "first test with odd number of elements",
			num1:   []int{1, 3},
			num2:   []int{2},
			medium: 2.00,
		},
		{
			desc:   "second test even elements",
			num1:   []int{1, 2, 3, 4, 5, 6},
			num2:   []int{3, 4, 5, 5},
			medium: 4.0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.medium, binarysearch.FindMedianSortedArrays(tC.num1, tC.num2))
		})
	}
}
