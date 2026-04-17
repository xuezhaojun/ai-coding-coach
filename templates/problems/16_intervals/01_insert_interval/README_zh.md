# 插入区间

- **难度**: 中等
- **分类**: 区间
- **考点**: 数组, 区间
- **链接**: [NeetCode](https://neetcode.io/problems/insert-new-interval) | [力扣 57](https://leetcode.cn/problems/insert-interval/)

## 题目描述

给你一个无重叠的、按照区间起始端点排序的区间列表 `intervals`，其中 `intervals[i] = [start_i, end_i]` 表示第 `i` 个区间的开始和结束。同样给定一个区间 `newInterval = [start, end]` 表示另一个区间的开始和结束。

在 `intervals` 中插入 `newInterval`，使得 `intervals` 仍然按起始端点升序排列，且不存在重叠区间（如有必要则合并重叠区间）。返回插入后的 `intervals`。

注意你不需要原地修改 `intervals`，可以创建一个新数组并返回。

## 示例

**示例 1:**

```
输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
输出: [[1,5],[6,9]]
解释: 新区间 [2,5] 与 [1,3] 重叠，合并为 [1,5]。
```

**示例 2:**

```
输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出: [[1,2],[3,10],[12,16]]
解释: 新区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠，合并为 [3,10]。
```

**示例 3:**

```
输入: intervals = [], newInterval = [5,7]
输出: [[5,7]]
解释: 没有已存在的区间，直接插入新区间。
```

## 约束条件

- `0 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= start_i <= end_i <= 10^5`
- `intervals` 按 `start_i` 升序排列
- `newInterval.length == 2`
- `0 <= start <= end <= 10^5`

## 函数签名

```go
func insert(intervals [][]int, newInterval []int) [][]int
```
