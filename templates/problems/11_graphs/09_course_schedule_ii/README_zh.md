# 课程表 II

- **难度**: 中等
- **分类**: 图
- **考点**: 拓扑排序, 广度优先搜索, Kahn 算法, 有向图
- **链接**: [NeetCode](https://neetcode.io/problems/course-schedule-ii) | [力扣 210](https://leetcode.cn/problems/course-schedule-ii/)

## 题目描述

现在你总共有 `numCourses` 门课需要选，记为 `0` 到 `numCourses - 1`。给你一个数组 `prerequisites`，其中 `prerequisites[i] = [ai, bi]` 表示在选修课程 `ai` 前必须先选修 `bi`。

返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回任意一种就可以了。如果不可能完成所有课程（存在环），返回一个空数组。

## 示例

**示例 1:**

```
输入: numCourses = 2, prerequisites = [[1,0]]
输出: [0,1]
解释: 先修课程 0，再修课程 1。
```

**示例 2:**

```
输入: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
输出: [0,1,2,3] 或 [0,2,1,3]
解释: 存在多个有效的顺序，课程 0 必须排在最前面。
```

**示例 3:**

```
输入: numCourses = 2, prerequisites = [[1,0],[0,1]]
输出: []
解释: 存在环，无法得到有效的排序。
```

## 约束条件

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= numCourses * (numCourses - 1)`
- `prerequisites[i].length == 2`
- `0 <= ai, bi < numCourses`
- `ai != bi`
- 所有 `[ai, bi]` 对互不相同。

## 函数签名

```go
func findOrder(numCourses int, prerequisites [][]int) []int
```
