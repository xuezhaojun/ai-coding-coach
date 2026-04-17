# 无向图中连通分量的数目

- **难度**: 中等
- **分类**: 图
- **考点**: 并查集, 图, 连通分量
- **链接**: [NeetCode](https://neetcode.io/problems/count-connected-components) | [力扣 323](https://leetcode.cn/problems/number-of-connected-components-in-an-undirected-graph/) | [LintCode 431](https://www.lintcode.com/problem/431/)

## 题目描述

你有一个包含 `n` 个节点的图。给你一个整数 `n` 和一个数组 `edges`，其中 `edges[i] = [ai, bi]` 表示图中节点 `ai` 和 `bi` 之间存在一条无向边。

返回图中连通分量的数目。连通分量是一个最大的节点集合，其中每对节点之间都存在路径。

## 示例

**示例 1:**

```
输入: n = 5, edges = [[0,1],[1,2],[3,4]]
输出: 2
解释: 连通分量为 {0,1,2} 和 {3,4}。
```

**示例 2:**

```
输入: n = 5, edges = [[0,1],[1,2],[2,3],[3,4]]
输出: 1
解释: 所有节点都连通在一个分量中。
```

**示例 3:**

```
输入: n = 4, edges = []
输出: 4
解释: 没有边，每个节点都是独立的分量。
```

## 约束条件

- `1 <= n <= 2000`
- `1 <= edges.length <= 5000`
- `edges[i].length == 2`
- `0 <= ai, bi < n`
- `ai != bi`
- 没有重复的边。

## 函数签名

```go
func countComponents(n int, edges [][]int) int
```
