func removeNthFromEnd(head *ListNode, k int) *ListNode {
	p, x := head, head
	n := 0
	for p != nil {
		n++
		p = p.Next
	}
	k = n - k
	index := 0
	if x.Next == nil || k == 0 {
		return head.Next
	}
	for x != nil {
		if index == k-1 {
			next := x.Next
			x.Next = next.Next
		}
		index++
		x = x.Next
	}
	return head
}