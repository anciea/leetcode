package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	var dfs func(curr *TreeNode, tmp int)
	dfs = func(curr *TreeNode, tmp int) {
		if curr == nil {
			return
		}
		if len(ans) < tmp+1 {
			ans = append(ans, make([]int, 0))
		}
		ans[tmp] = append(ans[tmp], curr.Val)

		if curr.Left != nil {
			dfs(curr.Left, tmp+1)
		}
		if curr.Right != nil {
			dfs(curr.Right, tmp+1)
		}
	}

	dfs(root, 0)
	return ans
}
