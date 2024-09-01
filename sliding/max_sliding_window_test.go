package sliding_test

import (
	"slices"
	"testing"

	"github.com/abassGarane/leetcode/sliding"
)

type testStruct struct {
	have []int
	k    int
	need []int
}

func TestMaxSlidingWindow(t *testing.T) {
	ts := []*testStruct{
		{
			have: []int{},
			k:    2,
			need: []int{},
		},
		{
			have: []int{1, 2, 3},
			k:    3,
			need: []int{3},
		},
		{
			have: []int{1, 2, 3, 1},
			k:    3,
			need: []int{3, 3},
		},
	}
	for _, v := range ts {
		res := sliding.MaxSlidingWindow(v.have, v.k)
		if !slices.Equal(v.need, res) {
			t.Errorf("expected %v got %v", v.need, res)
		}

	}
}
