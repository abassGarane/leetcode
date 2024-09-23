package linked_list

// Given the list 1-> 2 -> 3 -> 4, reorder it to 1-> 4 -> 2 -> 3
func ReOrderList(head *ListNode) {
	// special case
	if head == nil || head.Next == nil {
		return
	}
	// break the linked list into two
	breakList(head)
	// reverse one part/second part
	reverseList(head)
	// merge the two arrays ie head to head and continue that way
}
func breakList(list *ListNode) (l1, l2 *ListNode) {}
func reverseList(list *ListNode) *ListNode        {}
func merge(l1, l2 *ListNode) *ListNode            {}
