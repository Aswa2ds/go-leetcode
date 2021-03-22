package leetcode

type listNode struct {
	val  int
	next *listNode
}

func getPath(root, p *TreeNode) *listNode {
	curr := &listNode{
		val:  root.Val,
		next: nil,
	}
	if root.Val == p.Val {
		return curr
	}
	if root.Val > p.Val {
		curr.next = getPath(root.Left, p)
		return curr
	}
	if root.Val < p.Val {
		curr.next = getPath(root.Right, p)
		return curr
	}
	return nil
}

func getNode(root *TreeNode, target int) *TreeNode {
	var result *TreeNode
	if root.Val == target {
		result = root
	}
	if root.Val < target {
		result = getNode(root.Right, target)
	}
	if root.Val > target {
		result = getNode(root.Left, target)
	}
	return result
}

func lowestCommonAncestor(root, a, b *TreeNode) *TreeNode {
	path1 := getPath(root, a)
	path2 := getPath(root, b)
	result := 0
	p, q := path1, path2
	for {
		if p.val == q.val {
			result = p.val
		} else {
			break
		}
		p, q = p.next, q.next
	}
	return getNode(root, result)
}
