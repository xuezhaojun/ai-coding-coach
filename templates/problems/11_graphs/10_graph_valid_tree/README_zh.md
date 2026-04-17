# 以图判树

- **难度**: 中等
- **分类**: 图
- **考点**: 并查集, 图, 树
- **链接**: [NeetCode](https://neetcode.io/problems/valid-tree) | [力扣 261](https://leetcode.cn/problems/graph-valid-tree/) | [LintCode 178](https://www.lintcode.com/problem/178/)

## 题目描述

给定编号从 `0` 到 `n - 1` 的 `n` 个节点和一个无向边列表（每条边以节点对表示），请判断这些边是否能够形成一个有效的树。

有效的树是一个连通的、无环的无向图。等价地，一棵有 `n` 个节点的有效树恰好有 `n - 1` 条边且是连通的。

## 示例

**示例 1:**

```
输入: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
输出: true
```

**示例 2:**

```
输入: n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
输出: false
解释: 存在环：1 -> 2 -> 3 -> 1。
```

**示例 3:**

```
输入: n = 4, edges = [[0,1],[2,3]]
输出: false
解释: 图不连通（有两个独立的连通分量）。
```

## 约束条件

- `1 <= n <= 2000`
- `0 <= edges.length <= 5000`
- `edges[i].length == 2`
- `0 <= ai, bi < n`
- `ai != bi`
- 没有自环或重复的边。

## 函数签名

```go
func validTree(n int, edges [][]int) bool
```
