# 矩阵中的最长递增路径

- **难度**: 困难
- **分类**: 二维动态规划
- **考点**: 动态规划, 深度优先搜索, 记忆化搜索, 矩阵
- **链接**: [NeetCode](https://neetcode.io/problems/longest-increasing-path-in-matrix) | [力扣 329](https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/)

## 题目描述

给定一个 `m x n` 整数矩阵 `matrix`，找出其中最长递增路径的长度。对于每个单元格，你可以往上、下、左、右四个方向移动。不能沿对角线方向移动或移动到边界外。路径必须是严格递增的。

## 示例

**示例 1:**

```
输入: matrix = [[9,9,4],[6,6,8],[2,1,1]]
输出: 4
解释: 最长递增路径为 [1, 2, 6, 9]。
```

**示例 2:**

```
输入: matrix = [[3,4,5],[3,2,6],[2,2,1]]
输出: 4
解释: 最长递增路径为 [3, 4, 5, 6]。注意不允许沿对角线移动。
```

**示例 3:**

```
输入: matrix = [[1]]
输出: 1
```

## 约束条件

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 200`
- `0 <= matrix[i][j] <= 2^31 - 1`

## 函数签名

```go
func longestIncreasingPath(matrix [][]int) int
```
