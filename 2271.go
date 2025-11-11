package main

import "sort"

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
    sort.Slice(tiles, func(i, j int) bool {
        return tiles[i][0] < tiles[j][0]
    })

    n := len(tiles)
    maxCover := 0
    currentCover := 0 // 窗口 [l, r] 内的瓷砖总长度
    
    // r 是地毯 *可能* 覆盖的最后一个瓷砖
    // l 是地毯 *开始* 覆盖的第一个瓷砖
    
    l := 0 
    for r := 0; r < n; r++ {
        // 假设地毯的 *起始点* 是 tiles[l][0]
        carpetStart := tiles[l][0]
        carpetEnd := carpetStart + carpetLen - 1
        
        // 1. 扩大窗口：把 r 包含进来
        currentCover += (tiles[r][1] - tiles[r][0] + 1)
        
        // 2. 收缩窗口：
        // 只要 tiles[r] 的右边界已经超出了地毯的覆盖范围
        // 我们就需要移动地毯的起始点 (l++)
        for l <= r && tiles[r][1] > carpetEnd {
            // 在移动 l 之前，先计算 l 对应的地毯的覆盖量
            
            // 此时 r 已经超出了地毯 [carpetStart, carpetEnd]
            // 我们需要减去 r 多出来的部分
            overCount := 0
            if tiles[r][0] <= carpetEnd { // r 和地毯有交集
                overCount = tiles[r][1] - carpetEnd
            } else { // r 和地毯完全没有交集
                // 这种情况理论上在 currentCover 累加时就被 overCount 覆盖了
                // 但为了严谨，我们减去 r 的全部长度
                overCount = (tiles[r][1] - tiles[r][0] + 1)
            }
            
            actualCover := currentCover - overCount
            maxCover = max(maxCover, actualCover)
            
            // 窗口 l 移出
            currentCover -= (tiles[l][1] - tiles[l][0] + 1)
            l++ // 尝试下一个地毯起始点
            
            // 更新地毯边界
            if l <= r {
                 carpetStart = tiles[l][0]
                 carpetEnd = carpetStart + carpetLen - 1
            }
        }
        
        // 3. 检查当前窗口 [l, r] 是否完全被 [carpetStart, carpetEnd] 覆盖
        if l <=r {
             // 此时 r 还在地毯内 (tiles[r][1] <= carpetEnd)
             // 我们不需要减去任何东西，currentCover 就是覆盖量
             maxCover = max(maxCover, currentCover)
        }
    }
    
    return maxCover
}
