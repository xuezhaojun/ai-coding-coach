# 组合总和 II

- **难度**: 中等
- **分类**: 回溯
- **考点**: 回溯, 排序, 去重
- **链接**: [NeetCode](https://neetcode.io/problems/combination-target-sum-ii) | [力扣 40](https://leetcode.cn/problems/combination-sum-ii/)

## 题目描述

给定一个候选人编号的集合 `candidates` 和一个目标数 `target`，找出 `candidates` 中所有可以使数字和为 `target` 的组合。`candidates` 中的每个数字在每个组合中只能使用一次。

注意：解集不能包含重复的组合。

## 示例

**示例 1:**

```
输入: candidates = [10, 1, 2, 7, 6, 1, 5], target = 8
输出: [[1,1,6], [1,2,5], [1,7], [2,6]]
```

**示例 2:**

```
输入: candidates = [2, 5, 2, 1, 2], target = 5
输出: [[1,2,2], [5]]
```

**示例 3:**

```
输入: candidates = [3, 5], target = 1
输出: []
解释: 没有组合能够求和为 1。
```

## 约束条件

- `1 <= candidates.length <= 100`
- `1 <= candidates[i] <= 50`
- `1 <= target <= 30`

## 函数签名

```go
func combinationSum2(candidates []int, target int) [][]int
```
