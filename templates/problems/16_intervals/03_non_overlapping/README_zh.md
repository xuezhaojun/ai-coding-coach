# 无重叠区间

- **难度**: 中等
- **分类**: 区间
- **考点**: 数组, 贪心, 排序, 区间
- **链接**: [NeetCode](https://neetcode.io/problems/non-overlapping-intervals) | [力扣 435](https://leetcode.cn/problems/non-overlapping-intervals/)

## 题目描述

给定一个区间数组 `intervals`，其中 `intervals[i] = [start_i, end_i]`，返回需要移除的最少区间数量，使得剩余区间互不重叠。

两个区间 `[a, b]` 和 `[c, d]` 如果共享任何内部点（即 `a < d` 且 `c < b`）则认为重叠。仅在边界点接触的区间（例如 `[1, 2]` 和 `[2, 3]`）不算重叠。

## 示例

**示例 1:**

```
输入: intervals = [[1,2],[2,3],[3,4],[1,3]]
输出: 1
解释: 移除 [1,3] 后，剩余 [[1,2],[2,3],[3,4]] 互不重叠。
```

**示例 2:**

```
输入: intervals = [[1,2],[2,3]]
输出: 0
解释: 区间不重叠（仅在点 2 处接触），无需移除。
```

**示例 3:**

```
输入: intervals = [[1,2],[1,2],[1,2]]
输出: 2
解释: 三个相同的区间中必须移除两个。
```

## 约束条件

- `1 <= intervals.length <= 10^5`
- `intervals[i].length == 2`
- `-5 * 10^4 <= start_i < end_i <= 5 * 10^4`

## 函数签名

```go
func eraseOverlapIntervals(intervals [][]int) int
```
