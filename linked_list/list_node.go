package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) Traverse() []int {
	elements := []int{}
	for ln != nil {
		elements = append(elements, ln.Val)
		ln = ln.Next
	}
	return elements
}

func ToLinkedList(list []int) *ListNode {
	var head *ListNode
	for _, v := range list {
		node := &ListNode{
			Val:  v,
			Next: nil,
		}
		if head == nil {
			head = node
			continue
		}
		head.Push(node)
	}
	return head
}

func (ln *ListNode) Print() {
	items := ln.Traverse()
	for _, v := range items {
		println(v)
	}
}
func (ln *ListNode) Push(item *ListNode) {
	for ln != nil {
		if ln.Next == nil {
			ln.Next = item
			break
		}
		ln = ln.Next
	}
}
