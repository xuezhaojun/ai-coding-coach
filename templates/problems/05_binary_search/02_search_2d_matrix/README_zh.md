# 搜索二维矩阵

- **难度**：中等
- **分类**：二分查找
- **考点**：二分查找、矩阵
- **链接**：[NeetCode](https://neetcode.io/problems/search-2d-matrix) | [力扣 74](https://leetcode.cn/problems/search-a-2d-matrix/)

## 题目描述

给你一个满足下述两条属性的 `m x n` 整数矩阵：
- 每行中的整数从左到右按非递减顺序排列。
- 每行的第一个整数大于前一行的最后一个整数。

给你一个整数 `target`，如果 `target` 在矩阵中，返回 `true`；否则，返回 `false`。你必须设计时间复杂度为 O(log(m * n)) 的解法。

## 示例

**示例 1：**

```
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true
解释：3 存在于第一行中。
```

**示例 2：**

```
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出：false
解释：13 不存在于矩阵的任何一行中。
```

**示例 3：**

```
输入：matrix = [[1]], target = 1
输出：true
解释：矩阵只有一个元素，与目标值相等。
```

## 约束条件

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 100`
- `-10^4 <= matrix[i][j], target <= 10^4`

## 函数签名

```go
func searchMatrix(matrix [][]int, target int) bool
```
