package linked_list_test

import (
	"testing"

	"github.com/abassGarane/leetcode/linked_list"
	"github.com/stretchr/testify/assert"
)

func TestReOrderList(t *testing.T) {
	testCases := []struct {
		desc      string
		given     []int
		reordered []int
	}{
		{
			desc:      "equal list reordering(even elements)",
			given:     []int{1, 2, 3, 4},
			reordered: []int{1, 4, 2, 3},
		},
		{
			desc:      "given an odd elements linked list",
			given:     []int{1, 2, 3, 4, 5},
			reordered: []int{1, 5, 2, 4, 3},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ll := linked_list.ToLinkedList(tC.given)
			linked_list.ReOrderList(ll)
			assert.Equal(t, tC.reordered, ll.Traverse())
		})
	}
}
