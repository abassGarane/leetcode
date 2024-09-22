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
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.reversed, linked_list.Reverse(linked_list.ToLinkedList(tC.head_list)).Traverse())
		})
	}
}
