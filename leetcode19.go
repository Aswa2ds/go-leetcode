package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{
		Val:  -1,
		Next: head,
	}
	p := newHead
	for i := 0; i < n; i++ {
		p = p.Next
	}
	q := newHead
	for p.Next != nil {
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next
	return newHead.Next

}
