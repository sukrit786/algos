func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := ListNode{}
    carry, sum := 0, 0
    p, p1, p2 := &dummy, l1, l2
    
    for p1!=nil || p2!=nil || carry!=0 {
        sum = carry;
        if p1!=nil {
            sum+=p1.Val
            p1 = p1.Next
        }
        if(p2!=nil) {
            sum+=p2.Val
            p2=p2.Next
        }
        carry = sum/10;
        p.Next = &ListNode{Val:sum%10,Next:nil}
        p = p.Next
    }
    return dummy.Next
}