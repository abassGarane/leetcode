package linked_list_test

import (
	"testing"

	"github.com/abassGarane/leetcode/linked_list"
	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		desc   string
		l1     []int
		l2     []int
		result []int
	}{
		{
			desc:   "first test",
			l1:     []int{1, 2, 4, 5},
			l2:     []int{3, 5, 6, 7},
			result: []int{1, 2, 3, 4, 5, 5, 6, 7},
		},
		{
			desc:   "second test merge two empty lists",
			l1:     []int{},
			l2:     []int{},
			result: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			l1_ll := linked_list.ToLinkedList(tC.l1)
			l2_ll := linked_list.ToLinkedList(tC.l2)
			merged := linked_list.Merge(l1_ll, l2_ll)
			assert.Equal(t, tC.result, merged.Traverse())
		})
	}
}
