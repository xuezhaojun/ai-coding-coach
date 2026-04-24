# 冗余连接

- **难度**: 中等
- **分类**: 图
- **考点**: 并查集, 图, 环检测
- **链接**: [NeetCode](https://neetcode.io/problems/redundant-connection) | [力扣 684](https://leetcode.cn/problems/redundant-connection/)

## 题目描述

树可以看成是一个连通且无环的无向图。给定一个往一棵 `n` 个节点（节点标签从 `1` 到 `n`）的树中添加一条额外边后的图。添加的边连接了树中已有的两个节点，恰好形成了一个环。

返回一条可以删去的边，使得结果图是一棵有 `n` 个节点的树。如果有多个答案，则返回数组 `edges` 中最后出现的边。

## 示例

**示例 1:**


![示例 1](assets/reduntant1-1-graph.jpg)
```
输入: edges = [[1,2],[1,3],[2,3]]
输出: [2,3]
```

**示例 2:**


![示例 2](assets/reduntant1-2-graph.jpg)
```
输入: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
输出: [1,4]
```

**示例 3:**

```
输入: edges = [[1,2],[2,3],[3,4],[4,5],[5,1]]
输出: [5,1]
解释: 删除最后一条边 [5,1] 可以打破环。
```

## 约束条件

- `n == edges.length`
- `3 <= n <= 1000`
- `edges[i].length == 2`
- `1 <= ai < bi <= edges.length`
- `ai != bi`
- 没有重复的边。
- 给定的图是连通的。

## 函数签名

```go
func findRedundantConnection(edges [][]int) []int
```
