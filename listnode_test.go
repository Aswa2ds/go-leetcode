package leetcode

import (
	"reflect"
	"testing"
)

func Test_fromArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "empty test",
			args: args{nums: []int{}},
			want: nil,
		},
		{
			name: "one node test",
			args: args{nums: []int{1}},
			want: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
		{
			name: "multi nodes test",
			args: args{nums: []int{1, 2, 3}},
			want: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fromArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_printList(t *testing.T) {
	fromArray([]int{1,2,3,4,5}).printList()
}