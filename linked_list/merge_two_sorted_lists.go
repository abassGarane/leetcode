package linked_list

func Merge(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			current = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			current = l2
			l2 = l2.Next
		}
	}
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}
	return head.Next
}
