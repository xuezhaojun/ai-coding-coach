# 一手顺子

- **难度**: 中等
- **分类**: 贪心
- **考点**: 数组, 哈希表, 贪心, 排序
- **链接**: [NeetCode](https://neetcode.io/problems/hand-of-straights) | [力扣 846](https://leetcode.cn/problems/hand-of-straights/)

## 题目描述

Alice 手中有一些牌，用整数数组 `hand` 表示，其中 `hand[i]` 是第 `i` 张牌上写的数值。她想把牌重新排列成组，每组大小为 `groupSize`，且每组由 `groupSize` 张连续的牌组成。给定她的手牌和组大小，如果她能重新排列成功，返回 `true`，否则返回 `false`。

## 示例

**示例 1:**

```
输入: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
输出: true
解释: Alice 的手牌可以重新排列为 [1,2,3]、[2,3,4]、[6,7,8]。
```

**示例 2:**

```
输入: hand = [1,2,3,4,5], groupSize = 4
输出: false
解释: Alice 的手牌无法被重新排列成大小为 4 的连续牌组。
```

**示例 3:**

```
输入: hand = [1,2,3], groupSize = 3
输出: true
```

## 约束条件

- `1 <= hand.length <= 10^4`
- `0 <= hand[i] <= 10^9`
- `1 <= groupSize <= hand.length`

## 函数签名

```go
func isNStraightHand(hand []int, groupSize int) bool
```
