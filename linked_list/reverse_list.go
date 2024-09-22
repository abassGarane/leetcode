package linked_list

func Reverse(head *ListNode) *ListNode {
	var previous *ListNode
	for head != nil {
		next_node := head.Next
		head.Next = previous
		previous = head
		head = next_node
	}
	return previous
}
