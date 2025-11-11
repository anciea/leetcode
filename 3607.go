package main

import (
	redblacktree "github.com/emirpasic/gods/v2/trees/redblacktree"
)


func processQueries(c int, connections [][]int, queries [][]int) []int {
	ans := make([]int, 0)
	parent := make([]int, c+1)
	// regionMap 存储每个 *根节点* 对应的在线电站集合（红黑树）
	regionMap := make(map[int]*redblacktree.Tree[int, struct{}])

	for i := range parent {
		parent[i] = i
		// 初始化时，每个电站都是一个独立的根
		regionMap[i] = redblacktree.NewWith[int, struct{}](func(x, y int) int {
			return x - y // 使用标准整数比较
		})
		regionMap[i].Put(i, struct{}{}) // 初始时所有电站都在线
	}

	var foundFunc func(x int) int
	var unionFunc func(x int, y int)

	// DSU 的 Find 操作（带路径压缩）
	foundFunc = func(x int) int {
		if x != parent[x] {
			parent[x] = foundFunc(parent[x])
		}
		return parent[x]
	}

	// DSU 的 Union 操作（带“按大小合并”优化）
	unionFunc = func(x int, y int) {
		xRoot := foundFunc(x)
		yRoot := foundFunc(y)

		if xRoot != yRoot {
			xRootTree := regionMap[xRoot]
			yRootTree := regionMap[yRoot]

			// 优化：总是将节点数（在线电站数）较少的树合并到较多的树
			// 这样可以减少总的插入次数
			if xRootTree.Size() > yRootTree.Size() {
				// 交换 xRoot 和 yRoot，确保 xRoot 是较小的一方
				xRoot, yRoot = yRoot, xRoot
				xRootTree, yRootTree = yRootTree, xRootTree
			}

			// 将 xRoot (较小的) 合并到 yRoot (较大的)
			parent[xRoot] = yRoot

			for _, key := range xRootTree.Keys() {
				yRootTree.Put(key, struct{}{})
			}
			// 释放掉旧的树（现在已经合并）
			regionMap[xRoot] = nil
		}
	}

	// 1. 构建初始的电网（并查集）
	for _, connection := range connections {
		unionFunc(connection[0], connection[1])
	}

	// vis[i] == 1 表示电站 i 已离线
	vis := make([]int, c+1)

	// 2. 处理查询
	for _, query := range queries {
		if query[0] == 1 { // 类型 1：检查
			x := query[1]
			if vis[x] == 0 {
				// 电站 x 在线，它自己解决
				ans = append(ans, x)
			} else {
				// 电站 x 离线，需要找到电网中编号最小的 *在线* 电站
				root := foundFunc(x)
				tree := regionMap[root]

				// 如果 tree.Empty() 为 true，说明该电网没有在线电站了
				// 此时调用 tree.Left() 会返回 nil，tree.Left().Key 会导致 nil pointer panic
				if tree == nil || tree.Empty() {
					ans = append(ans, -1) // 按题目要求返回 -1
				} else {
					// .Left() 返回树中最小的节点
					ans = append(ans, tree.Left().Key)
				}
			}
		} else { // 类型 2：离线
			x := query[1]
			if vis[x] == 0 { // 只有在它还在线时才需要处理
				vis[x] = 1 // 标记为离线
				root := foundFunc(x)
				tree := regionMap[root]
				
				// 从该电网的“在线电站树”中移除 x
				if tree != nil {
					tree.Remove(x)
				}
			}
		}
	}

	return ans
}
