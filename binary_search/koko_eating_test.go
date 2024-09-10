package binarysearch_test

import (
	"testing"

	binarysearch "github.com/abassGarane/leetcode/binary_search"
	"github.com/stretchr/testify/assert"
)

func TestMinEatingSpeed(t *testing.T) {
	testCases := []struct {
		desc  string
		piles []int
		h     int
		k     int
	}{
		{
			desc:  "first test",
			piles: []int{3, 6, 7, 11},
			h:     8,
			k:     4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.k, binarysearch.MinEatingSpeed(tC.piles, tC.h))
		})
	}
}
