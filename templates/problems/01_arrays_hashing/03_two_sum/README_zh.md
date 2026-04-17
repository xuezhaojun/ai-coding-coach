# 两数之和

- **难度**：简单
- **分类**：数组与哈希
- **考点**：数组、哈希表
- **链接**：[NeetCode](https://neetcode.io/problems/two-integer-sum) | [力扣 1](https://leetcode.cn/problems/two-sum/)

## 题目描述

给定一个整数数组 `nums` 和一个整数目标值 `target`，请你在该数组中找出和为目标值 `target` 的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用同一个元素两次。你可以按任意顺序返回答案。

## 示例

**示例 1：**

```
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9，返回 [0, 1]。
```

**示例 2：**

```
输入：nums = [3,2,4], target = 6
输出：[1,2]
解释：因为 nums[1] + nums[2] == 6，返回 [1, 2]。
```

**示例 3：**

```
输入：nums = [3,3], target = 6
输出：[0,1]
```

## 约束条件

- `2 <= nums.length <= 10^4`
- `-10^9 <= nums[i] <= 10^9`
- `-10^9 <= target <= 10^9`
- 只会存在一个有效答案。

## 函数签名

```go
func twoSum(nums []int, target int) []int
```
