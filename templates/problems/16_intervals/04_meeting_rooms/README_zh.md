# 会议室

- **难度**: 简单
- **分类**: 区间
- **考点**: 数组, 排序, 区间
- **链接**: [NeetCode](https://neetcode.io/problems/meeting-schedule) | [力扣 252](https://leetcode.cn/problems/meeting-rooms/) | [LintCode 920](https://www.lintcode.com/problem/920/)

## 题目描述

给定一个会议时间区间数组 `intervals`，其中 `intervals[i] = [start_i, end_i]`，判断一个人是否能参加所有会议。只有当没有两个会议在时间上重叠时，才能全部参加。一个会议恰好在另一个会议结束时开始（例如 `[1, 5]` 和 `[5, 10]`）不算冲突。

## 示例

**示例 1:**

```
输入: intervals = [[0,30],[5,10],[15,20]]
输出: false
解释: 会议 [5,10] 与 [0,30] 冲突，因为 5 < 30。
```

**示例 2:**

```
输入: intervals = [[7,10],[2,4]]
输出: true
解释: 两个会议不重叠。
```

**示例 3:**

```
输入: intervals = [[1,5],[5,10]]
输出: true
解释: 会议在边界点 5 处接触但不重叠。
```

## 约束条件

- `0 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= start_i < end_i <= 10^6`

## 函数签名

```go
func canAttendMeetings(intervals [][]int) bool
```
