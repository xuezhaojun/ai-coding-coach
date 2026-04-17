# 二分查找

- **难度**：简单
- **分类**：二分查找
- **考点**：二分查找、数组
- **链接**：[NeetCode](https://neetcode.io/problems/binary-search) | [力扣 704](https://leetcode.cn/problems/binary-search/)

## 题目描述

给定一个 `n` 个元素有序的（升序）整型数组 `nums` 和一个目标值 `target`，写一个函数搜索 `nums` 中的 `target`，如果目标值存在返回下标，否则返回 `-1`。

你必须设计并实现时间复杂度为 O(log n) 的算法。

## 示例

**示例 1：**

```
输入：nums = [-1,0,3,5,9,12], target = 9
输出：4
解释：9 出现在 nums 中并且下标为 4。
```

**示例 2：**

```
输入：nums = [-1,0,3,5,9,12], target = 2
输出：-1
解释：2 不存在 nums 中因此返回 -1。
```

**示例 3：**

```
输入：nums = [5], target = 5
输出：0
解释：5 出现在 nums 中并且下标为 0。
```

## 约束条件

- `1 <= nums.length <= 10^4`
- `-10^4 < nums[i], target < 10^4`
- `nums` 中的所有元素互不相同
- `nums` 按升序排列

## 函数签名

```go
func search(nums []int, target int) int
```
