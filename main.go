package main

import (
	"github.com/abassGarane/leetcode/linked_list"
)

func main() {
	head := &linked_list.ListNode{
		Val:  1,
		Next: nil,
	}
	head.Push(&linked_list.ListNode{
		Val:  2,
		Next: nil,
	})
	head.Push(&linked_list.ListNode{
		Val:  3,
		Next: nil,
	})

	head.Print()
	linked_list.Reverse(head).Print()
}
