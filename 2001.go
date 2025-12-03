package main

func interchangeableRectangles(rectangles [][]int) int64 {
	vis := map[float64]int{}
	ans := int64(0)
	
	for _, rect := range rectangles {
		width, height := float64(rect[0]), float64(rect[1])
		ratio := width / height
		
		ans += int64(vis[ratio])
		vis[ratio]++
	}
	
	return ans
}

/*
注：这里为什么可以使用float64作为map的key？

// 关于浮点数精度：
// double 的有效精度约为 2^-52 ≈ 1e-16。
// 若两数的差小于其数量级 × 2^-52，则它们可能被舍入为同一个值。
// 对于 (a+1)/a 与 a/(a-1)，它们的差约为 1/(a^2)。
// 当 a > 2^26 ≈ 6.7e7 时，1/(a^2) < 2^-52，double 无法区分这两个值，结果会错误。
// 本题 a ≤ 1e5，远小于危险阈值，可以安全使用双精度浮点数。

*/
