package leetcode

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// if root.Left == nil && root.Right == nil {
		// return []int{root.Val}
	// }
	var ret []int
	ret = append(ret, inorderTraversal(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inorderTraversal(root.Right)...)
	return ret
}