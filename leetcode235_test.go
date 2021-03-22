package leetcode

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	root := generateTestTree()
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				root: root,
				p:    &TreeNode{
					Val: 1,
				},
				q:    &TreeNode{
					Val: 5,
				},
			},
			want: root,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func generateTestTree() *TreeNode {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	return root
}
