package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}
func maxLevelSum(root *TreeNode) int {
	ans, maxn := 1, root.Val
	vis := make(map[int]int)
	var dfs func(node *TreeNode, cnt int)
	dfs = func (node *TreeNode, cnt int) {
		if node == nil {
			return
		}
		vis[cnt] += node.Val
		dfs(node.Left, cnt + 1)
		dfs(node.Right, cnt + 1)
	}
	dfs(root, 1)

	for k,v := range vis {
		if maxn <= v {
			if maxn < v {
				ans = k
			} else {
				ans = min(ans, k)
			}
			maxn = v
		}
	}
	return ans
}
