package linked_list_test

import (
	"testing"

	"github.com/abassGarane/leetcode/linked_list"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		desc      string
		head_list []int
		reversed  []int
	}{
		{
			desc:      "first test",
			head_list: []int{1, 2, 3, 4},
			reversed:  []int{4, 3, 2, 1},
		},
		{
			desc:      "empty array",
			head_list: []int{},
			reversed:  []int{},
		},
	}
	for _, tC := range testCases {
		var head *linked_list.ListNode
		for _, v := range tC.head_list {
			node := &linked_list.ListNode{
				Val:  v,
				Next: nil,
			}
			if head == nil {
				head = node
				continue
			}
			head.Push(node)
		}
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.reversed, linked_list.Reverse(head).Traverse())
		})
	}
}
