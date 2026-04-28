package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode

	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		mid := node == p || node == q

		// 三种情况成立 -> 当前是 LCA
		if (mid && left) || (mid && right) || (left && right) {
			ans = node
		}

		return mid || left || right
	}

	dfs(root)
	return ans
}
