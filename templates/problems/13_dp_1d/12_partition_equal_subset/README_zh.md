# 分割等和子集

- **难度**: 中等
- **分类**: 一维动态规划
- **考点**: 动态规划, 0/1 背包
- **链接**: [NeetCode](https://neetcode.io/problems/partition-equal-subset-sum) | [力扣 416](https://leetcode.cn/problems/partition-equal-subset-sum/)

## 题目描述

给你一个只包含正整数的非空数组 `nums`。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

## 示例

**示例 1:**

```
输入: nums = [1,5,11,5]
输出: true
解释: 数组可以分割为 [1,5,5] 和 [11]，两者和均为 11。
```

**示例 2:**

```
输入: nums = [1,2,3,5]
输出: false
解释: 数组不能被分割成两个元素和相等的子集。
```

**示例 3:**

```
输入: nums = [1,2,3,4,5,6,7]
输出: true
解释: 数组可以分割为 [1,6,7] 和 [2,3,4,5]，两者和均为 14。
```

## 约束条件

- `1 <= nums.length <= 200`
- `1 <= nums[i] <= 100`

## 函数签名

```go
func canPartition(nums []int) bool
```
