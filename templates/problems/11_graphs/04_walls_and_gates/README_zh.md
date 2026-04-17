# 墙与门

- **难度**: 中等
- **分类**: 图
- **考点**: 广度优先搜索, 多源BFS, 矩阵
- **链接**: [NeetCode](https://neetcode.io/problems/islands-and-treasure) | [力扣 286](https://leetcode.cn/problems/walls-and-gates/) | [LintCode 663](https://www.lintcode.com/problem/663/)

## 题目描述

给你一个 `m x n` 的网格 `rooms`，使用以下三种值初始化：

- `-1` -- 墙壁或障碍物。
- `0` -- 门。
- `INF`（2147483647）-- 空房间。

请你将每个空房间填入该房间到最近的门的距离。如果无法到达门，则该房间保持 `INF`。

距离定义为沿上、下、左、右四个基本方向移动的步数。

## 示例

**示例 1:**

```
输入: rooms = [
  [INF, -1,  0, INF],
  [INF, INF, INF, -1],
  [INF, -1, INF, -1],
  [ 0,  -1, INF, INF]
]
输出: [
  [3, -1, 0,  1],
  [2,  2, 1, -1],
  [1, -1, 2, -1],
  [0, -1, 3,  4]
]
```

**示例 2:**

```
输入: rooms = [[0]]
输出: [[0]]
```

**示例 3:**

```
输入: rooms = [[INF, INF],[INF, INF]]
输出: [[INF, INF],[INF, INF]]
解释: 没有门，所有房间保持 INF。
```

## 约束条件

- `m == rooms.length`
- `n == rooms[i].length`
- `1 <= m, n <= 250`
- `rooms[i][j]` 是 `-1`、`0` 或 `2147483647`。

## 函数签名

```go
func wallsAndGates(rooms [][]int)
```
