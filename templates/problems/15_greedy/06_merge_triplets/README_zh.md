# 合并三元组以形成目标三元组

- **难度**: 中等
- **分类**: 贪心
- **考点**: 数组, 贪心
- **链接**: [NeetCode](https://neetcode.io/problems/merge-triplets-to-form-target) | [力扣 1899](https://leetcode.cn/problems/merge-triplets-to-form-target-triplet/)

## 题目描述

三元组是一个由三个整数组成的数组。给你一个二维整数数组 `triplets`，其中 `triplets[i] = [ai, bi, ci]` 表示第 `i` 个三元组。同时给你一个整数数组 `target = [x, y, z]`，表示你想要得到的三元组。为了得到 `target`，你可以对 `triplets` 执行下面的操作任意次：选择两个下标 `i` 和 `j`（`i != j`），将 `triplets[j]` 更新为 `[max(ai, aj), max(bi, bj), max(ci, cj)]`。如果通过以上操作可以使 `target` 成为 `triplets` 的一个元素，返回 `true`，否则返回 `false`。

## 示例

**示例 1:**

```
输入: triplets = [[2,5,3],[1,8,4],[1,7,5]], target = [2,7,5]
输出: true
解释: 合并三元组 [2,5,3] 和 [1,7,5] 得到 [2,7,5]。
```

**示例 2:**

```
输入: triplets = [[3,4,5],[4,5,6]], target = [3,2,5]
输出: false
解释: 无法产生第二个位置值为 2 的组合，因为两个三元组的第二个位置值都 >= 4。
```

**示例 3:**

```
输入: triplets = [[2,5,3],[10,1,1],[1,7,5]], target = [2,7,5]
输出: true
解释: 三元组 [10,1,1] 因为 10 > target[0] 被过滤掉。合并 [2,5,3] 和 [1,7,5] 得到 [2,7,5]。
```

## 约束条件

- `1 <= triplets.length <= 10^5`
- `triplets[i].length == target.length == 3`
- `1 <= ai, bi, ci, x, y, z <= 1000`

## 函数签名

```go
func mergeTriplets(triplets [][]int, target []int) bool
```
