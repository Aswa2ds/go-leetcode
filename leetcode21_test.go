package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "empty test",
			args: args{
				l1: fromArray([]int{}),
				l2: fromArray([]int{}),
			},
			want: nil,
		},
		{
			name: "one empty test",
			args: args{
				l1: fromArray([]int{1}),
				l2: fromArray([]int{}),
			},
			want: fromArray([]int{1}),
		},
		{
			name: "one longer test",
			args: args{
				l1: fromArray([]int{1,2,3}),
				l2: fromArray([]int{1,4}),
			},
			want: fromArray([]int{1,1,2,3,4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
