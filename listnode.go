package leetcode

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func fromArray(nums []int) *ListNode {
    emptyHead := &ListNode{
        Val:  -1,
        Next: nil,
    }
    tail := emptyHead
    for _, num := range nums {
        node := &ListNode{
            Val:  num,
            Next: nil,
        }
        tail.Next = node
        tail = node
    }
    return emptyHead.Next
}

func (list *ListNode) printList() {
    for p := list; p != nil; p = p.Next {
        fmt.Print(p.Val, " ")
    }
}
