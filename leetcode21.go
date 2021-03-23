package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}

}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2
	if p == nil {
		return q
	}
	if q == nil {
		return p
	}
	emptyHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	tail := emptyHead
	for p != nil && q != nil {
		if p.Val < q.Val {
			tail.Next = p
			p = p.Next
		} else {
			tail.Next = q
			q = q.Next
		}
		tail = tail.Next
	}
	if p == nil {
		p = q
	}
	tail.Next = p
	return emptyHead.Next
}
