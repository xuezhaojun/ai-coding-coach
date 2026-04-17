# 存在重复元素

- **难度**：简单
- **分类**：数组与哈希
- **考点**：数组、哈希集合
- **链接**：[NeetCode](https://neetcode.io/problems/duplicate-integer) | [力扣 217](https://leetcode.cn/problems/contains-duplicate/)

## 题目描述

给定一个整数数组 `nums`，如果任一值在数组中出现 **至少两次**，返回 `true`；如果数组中每个元素互不相同，返回 `false`。

## 示例

**示例 1：**

```
输入：nums = [1,2,3,1]
输出：true
解释：元素 1 在下标 0 和 3 处重复出现。
```

**示例 2：**

```
输入：nums = [1,2,3,4]
输出：false
解释：所有元素各不相同。
```

**示例 3：**

```
输入：nums = [1,1,1,3,3,4,3,2,4,2]
输出：true
```

## 约束条件

- `1 <= nums.length <= 10^5`
- `-10^9 <= nums[i] <= 10^9`

## 函数签名

```go
func containsDuplicate(nums []int) bool
```
