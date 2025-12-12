from typing import List
from collections import defaultdict


class Solution:
    def countCoveredBuildings(self, n: int, buildings: List[List[int]]) -> int:
        # g1[x] 存储所有 x 坐标相同的建筑物的 y 坐标
        # g2[y] 存储所有 y 坐标相同的建筑物的 x 坐标
        g1 = defaultdict(list)
        g2 = defaultdict(list)
        for x, y in buildings:
            g1[x].append(y)
            g2[y].append(x)

        # 排序以便快速判断边界
        for x in g1:
            g1[x].sort()
        for y in g2:
            g2[y].sort()

        ans = 0
        for x, y in buildings:
            l1 = g1[x]  # 同一列的所有 y 坐标
            l2 = g2[y]  # 同一行的所有 x 坐标
            # 检查四个方向是否都有建筑物
            # l2[0] < x < l2[-1] 表示左右都有建筑物
            # l1[0] < y < l1[-1] 表示上下都有建筑物
            if l2[0] < x < l2[-1] and l1[0] < y < l1[-1]:
                ans += 1
        return ans
