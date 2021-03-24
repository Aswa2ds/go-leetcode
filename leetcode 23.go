package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	v := length/2
	left := mergeKLists(lists[:v])
	right := mergeKLists(lists[v:])
	return mergeTwoListsSubFunc(left, right)
}

func mergeTwoListsSubFunc(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsSubFunc(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsSubFunc(l1, l2.Next)
		return l2
	}

}