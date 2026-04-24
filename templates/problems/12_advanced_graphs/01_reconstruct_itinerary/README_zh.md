# 重新安排行程

- **难度**: 困难
- **分类**: 高级图
- **考点**: 欧拉路径, 深度优先搜索, Hierholzer 算法, 贪心
- **链接**: [NeetCode](https://neetcode.io/problems/reconstruct-flight-path) | [力扣 332](https://leetcode.cn/problems/reconstruct-itinerary/)

## 题目描述

给你一份航线列表 `tickets`，其中 `tickets[i] = [fromi, toi]` 表示一张从 `fromi` 飞往 `toi` 的机票。请你对该行程进行重新规划排序。

所有这些机票都属于一个从 `"JFK"`（肯尼迪国际机场）出发的人，所以该行程必须从 `"JFK"` 开始。如果存在多种有效的行程，请你按字典排序返回最小的行程组合。

你可以假设所有机票至少存在一种合理的行程。所有的机票必须都用一次且只能用一次。

## 示例

**示例 1:**


![示例 1](assets/itinerary1-graph.jpg)
```
输入: tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
输出: ["JFK","MUC","LHR","SFO","SJC"]
```

**示例 2:**


![示例 2](assets/itinerary2-graph.jpg)
```
输入: tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
输出: ["JFK","ATL","JFK","SFO","ATL","SFO"]
解释: 另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"]，但字典序更大。
```

**示例 3:**

```
输入: tickets = [["JFK","ATL"],["ATL","JFK"]]
输出: ["JFK","ATL","JFK"]
```

## 约束条件

- `1 <= tickets.length <= 300`
- `tickets[i].length == 2`
- `fromi.length == 3`
- `toi.length == 3`
- `fromi` 和 `toi` 由大写英文字母组成。
- `fromi != toi`

## 函数签名

```go
func findItinerary(tickets [][]string) []string
```
