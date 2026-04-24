# 柱状图中最大的矩形

- **难度**：困难
- **分类**：栈
- **考点**：栈、单调栈、数组
- **链接**：[NeetCode](https://neetcode.io/problems/largest-rectangle-in-histogram) | [力扣 84](https://leetcode.cn/problems/largest-rectangle-in-histogram/)

## 题目描述

给定 `n` 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 `1`。求在该柱状图中，能够勾勒出来的矩形的最大面积。

矩形必须由连续的柱子组成，其高度受限于所跨范围内最矮的柱子。

## 示例

**示例 1：**


![示例 1](assets/histogram.jpg)
```
输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形面积为 5 * 2 = 10（由下标 2 和 3 处的柱子组成，高度分别为 5 和 6，取最小高度 5，宽度 2）。
```

**示例 2：**


![示例 2](assets/histogram-1.jpg)
```
输入：heights = [5]
输出：5
解释：一根高度为 5、宽度为 1 的柱子，面积为 5。
```

**示例 3：**

```
输入：heights = [3,3,3,3]
输出：12
解释：所有柱子高度都是 3，矩形横跨全部 4 根柱子：3 * 4 = 12。
```

## 约束条件

- `1 <= heights.length <= 10^5`
- `0 <= heights[i] <= 10^4`

## 函数签名

```go
func largestRectangleArea(heights []int) int
```
