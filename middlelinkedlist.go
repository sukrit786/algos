func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		fmt.Println(slow.Val, fast.Val)
		if fast.Next == nil {
			fast = fast.Next
		} else {
			slow = slow.Next
			fast = fast.Next.Next
		}
	}
	fmt.Print(slow)
	return slow

}