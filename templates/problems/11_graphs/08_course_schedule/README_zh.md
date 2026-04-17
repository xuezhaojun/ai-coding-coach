# 课程表

- **难度**: 中等
- **分类**: 图
- **考点**: 深度优先搜索, 拓扑排序, 环检测, 有向图
- **链接**: [NeetCode](https://neetcode.io/problems/course-schedule) | [力扣 207](https://leetcode.cn/problems/course-schedule/)

## 题目描述

你这个学期必须选修 `numCourses` 门课程，记为 `0` 到 `numCourses - 1`。在选修某些课程之前需要一些先修课程。先修课程按数组 `prerequisites` 给出，其中 `prerequisites[i] = [ai, bi]` 表示如果要学习课程 `ai` 则必须先学习课程 `bi`。

例如，先修课程对 `[0, 1]` 表示：想要学习课程 `0`，你需要先完成课程 `1`。

请你判断是否可能完成所有课程的学习。如果可以，返回 `true`；否则返回 `false`。

## 示例

**示例 1:**

```
输入: numCourses = 2, prerequisites = [[1,0]]
输出: true
解释: 先修课程 0，再修课程 1。
```

**示例 2:**

```
输入: numCourses = 2, prerequisites = [[1,0],[0,1]]
输出: false
解释: 课程 0 和课程 1 相互依赖，形成了循环。
```

**示例 3:**

```
输入: numCourses = 3, prerequisites = []
输出: true
解释: 没有先修课程要求，所有课程可以独立学习。
```

## 约束条件

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= 5000`
- `prerequisites[i].length == 2`
- `0 <= ai, bi < numCourses`
- 所有先修课程对 `prerequisites[i]` 互不相同。

## 函数签名

```go
func canFinish(numCourses int, prerequisites [][]int) bool
```
