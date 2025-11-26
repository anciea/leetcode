```go
// 计算满足 check(x) == true 的最大整数 x
func binarySearchMax(nums []int) int {
    // 二分猜答案：判断 mid 是否满足题目要求
    check := func(mid int) bool {
        
    }

    left :=  // 循环不变量：check(left) 恒为 true
    right :=  // 循环不变量：check(right) 恒为 false
    for left+1 < right {
        mid := left + (right-left)/2
        if check(mid) {
            left = mid // 注意这里更新的是 left，和上面的模板反过来
        } else {
            right = mid
        }
    }
    // 循环结束后 left+1 = right
    // 此时 check(left) == true 而 check(left+1) == check(right) == false
    // 所以 left 就是最大的满足 check 的值
    return left // check 更新的是谁，最终就返回谁
}


```


这是一个非常经典的面试题，核心原因只有一个：**防止整型溢出（Integer Overflow）。**

虽然在数学上 `(left + right) / 2` 和 `left + (right - left) / 2` 是相等的，但在计算机的世界里，它们的安全性完全不同。

### 1. 为什么会溢出？

假设我们使用的是 32 位有符号整数（`int32`），它的最大值（`MAX_INT`）是 `2,147,483,647`。

设想以下场景：
*   **left** = `1,500,000,000` （很大的数，但在范围内）
*   **right** = `2,000,000,000` （很大的数，但在范围内）

如果你直接做加法：
```go
sum = left + right 
// 结果是 3,500,000,000
```
这个结果 **超过了 `MAX_INT`**！
在计算机中，溢出后通常会变成一个 **负数**（这是由二进制补码机制决定的）。

所以：
*   `(left + right) / 2` 可能会变成一个负数，导致程序去访问数组的负数下标，引发 `Index Out of Bounds` 异常或者逻辑错误（无限循环）。

### 2. 为什么减法写法是安全的？

```go
mid := left + (right - left) / 2
```

在这个公式中：
1.  **`right - left`**：因为 `right` 和 `left` 都在合法范围内，且 `right >= left`，所以它们的**差值**一定比 `right` 小，绝对不会溢出。
2.  **`(right - left) / 2`**：差的一半，当然更不会溢出。
3.  **`left + ...`**：因为 `mid` 必定在 `left` 和 `right` 之间，既然 `right` 没有溢出，那么小于 `right` 的 `mid` 自然也不会溢出。

**数学推导：**
$$ left + \frac{right - left}{2} = \frac{2 \cdot left + right - left}{2} = \frac{left + right}{2} $$
虽然结果一样，但计算过程避开了“先加出一个超大数”的风险。

### 3. 这个问题的历史背景

这是一个真实存在过的著名 Bug。
在 2006 年之前，**Java 标准库** 中的 `Arrays.binarySearch` 用的就是 `(low + high) / 2`。
直到有人在处理超大数组时发现了这个问题，Sun 公司才修复了这个 Bug。你可以搜索 "Google Research: Nearly All Binary Searches and Mergesorts are Broken" 了解这段历史。

### 4. 在 Go 语言中的情况

在现代 64 位系统上，Go 语言的 `int` 默认是 64 位的。
`int64` 的最大值约为 $9 \times 10^{18}$（900 亿亿）。

*   **现状**：除非你的数组长度达到了天文数字（内存都存不下），否则在 64 位 Go 程序里直接写 `(left + right) / 2` **通常也不会报错**。
*   **建议**：尽管如此，使用 `left + (right - left) / 2` 依然是**标准最佳实践**。
    *   它能保证你的代码移植到 32 位系统时不出错。
    *   它是算法竞赛和工程代码中的“行话”，表明你考虑到了边界情况。

### 总结
写 `left + (right - left) / 2` 是为了**代码健壮性**，防止当 `left` 和 `right` 都很大时，相加导致数值超出整数类型的上限。
