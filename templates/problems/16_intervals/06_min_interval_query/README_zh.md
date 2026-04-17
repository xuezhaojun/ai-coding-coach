# 包含每个查询的最小区间

- **难度**: 困难
- **分类**: 区间
- **考点**: 数组, 排序, 堆, 区间
- **链接**: [NeetCode](https://neetcode.io/problems/minimum-interval-including-query) | [力扣 1851](https://leetcode.cn/problems/minimum-interval-to-include-each-query/)

## 题目描述

给你一个二维整数数组 `intervals`，其中 `intervals[i] = [left_i, right_i]` 表示第 `i` 个区间从 `left_i` 开始到 `right_i` 结束（包含两端）。区间的大小定义为它包含的整数个数，即 `right_i - left_i + 1`。

同时给你一个整数数组 `queries`。第 `j` 个查询的答案是包含 `queries[j]` 的最小区间 `i` 的大小，即满足 `left_i <= queries[j] <= right_i` 的最小区间大小。如果不存在这样的区间，答案为 `-1`。

返回一个包含所有查询答案的数组。

## 示例

**示例 1:**

```
输入: intervals = [[1,4],[2,4],[3,6],[4,4]], queries = [2,3,4,5]
输出: [3,3,1,4]
解释:
- 查询 2：区间 [1,4]（大小 4）和 [2,4]（大小 3）包含 2，最小为 3。
- 查询 3：区间 [1,4]（大小 4）、[2,4]（大小 3）和 [3,6]（大小 4）包含 3，最小为 3。
- 查询 4：所有四个区间都包含 4，[4,4]（大小 1）最小。
- 查询 5：仅 [3,6]（大小 4）包含 5。
```

**示例 2:**

```
输入: intervals = [[2,3],[2,5],[1,8],[20,25]], queries = [2,19,5,22]
输出: [2,-1,4,6]
解释:
- 查询 2：[2,3]（大小 2）是包含 2 的最小区间。
- 查询 19：没有区间包含 19，返回 -1。
- 查询 5：[2,5]（大小 4）是包含 5 的最小区间。
- 查询 22：[20,25]（大小 6）包含 22。
```

## 约束条件

- `1 <= intervals.length <= 10^5`
- `1 <= queries.length <= 10^5`
- `intervals[i].length == 2`
- `1 <= left_i <= right_i <= 10^7`
- `1 <= queries[j] <= 10^7`

## 函数签名

```go
func minInterval(intervals [][]int, queries []int) []int
```
