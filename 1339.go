package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxProduct(root *TreeNode) int {
	const mod = 1000_000_007
	ans := int64(0)
	vis := make(map[*TreeNode]int64)
	
	var dfs func(node *TreeNode) int64
	dfs = func(node *TreeNode) int64 {
		if node == nil {
			return 0
		}
		// 计算当前子树的总和
		currentSum := int64(node.Val) + dfs(node.Left) + dfs(node.Right)
		vis[node] = currentSum
		return currentSum
	}
	
	totalSum := dfs(root)

	var dfs2 func(node *TreeNode) 
	dfs2 = func(node *TreeNode) {
		if node == nil {
			return
		}

		product := (totalSum - vis[node]) * vis[node]
		ans = max64(ans, product)
		
		dfs2(node.Left)
		dfs2(node.Right)
	}
	
	dfs2(root)
	return int(ans % mod) 
}

func max64(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}
