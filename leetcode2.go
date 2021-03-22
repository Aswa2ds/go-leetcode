package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c := 0
	ret := &ListNode{
		Val:  -1,
		Next: nil,
	}
	tail := ret
	sum := 0
	p, q := l1, l2
	for ;p != nil && q != nil; p, q = p.Next, q.Next {
		sum = (p.Val + q.Val + c) % 10
		c = (p.Val + q.Val + c) / 10
		tail.Next = &ListNode{
			Val:  sum,
			Next: nil,
		}
		tail = tail.Next
	}
	if q != nil {
		p = q
	}
	for ; p != nil; p = p.Next {
		sum = (p.Val + c) % 10
		c = (p.Val + c) / 10
		tail.Next = &ListNode{
			Val:  sum,
			Next: nil,
		}
		tail = tail.Next
	}
	if c != 0 {
		tail.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
		tail = tail.Next
	}
	return ret.Next
}
