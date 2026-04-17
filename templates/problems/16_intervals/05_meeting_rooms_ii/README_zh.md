# 会议室 II

- **难度**: 中等
- **分类**: 区间
- **考点**: 数组, 排序, 堆, 区间, 扫描线
- **链接**: [NeetCode](https://neetcode.io/problems/meeting-schedule-ii) | [力扣 253](https://leetcode.cn/problems/meeting-rooms-ii/) | [LintCode 919](https://www.lintcode.com/problem/919/)

## 题目描述

给定一个会议时间区间数组 `intervals`，其中 `intervals[i] = [start_i, end_i]`，求所有会议不冲突地进行所需的最少会议室数量。如果一个会议在另一个会议结束时或之后开始，它们可以共用同一个会议室。

## 示例

**示例 1:**

```
输入: intervals = [[0,30],[5,10],[15,20]]
输出: 2
解释: 会议 [0,30] 占用一个房间。会议 [5,10] 和 [15,20] 可以共用第二个房间，因为 [15,20] 在 [5,10] 结束后开始。
```

**示例 2:**

```
输入: intervals = [[7,10],[2,4]]
输出: 1
解释: 会议不重叠，只需一个房间。
```

**示例 3:**

```
输入: intervals = [[1,5],[2,6],[3,7]]
输出: 3
解释: 三个会议互相重叠，需要三个房间。
```

## 约束条件

- `0 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= start_i < end_i <= 10^6`

## 函数签名

```go
func minMeetingRooms(intervals [][]int) int
```
