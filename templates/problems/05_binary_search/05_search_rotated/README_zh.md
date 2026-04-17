# 搜索旋转排序数组

- **难度**：中等
- **分类**：二分查找
- **考点**：二分查找、数组
- **链接**：[NeetCode](https://neetcode.io/problems/search-in-rotated-sorted-array) | [力扣 33](https://leetcode.cn/problems/search-in-rotated-sorted-array/)

## 题目描述

整数数组 `nums` 按升序排列，数组中的值互不相同。在传递给函数之前，`nums` 在预先未知的某个下标 `k`（`1 <= k < nums.length`）上进行了旋转，使数组变为 `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]`（下标从 0 开始计数）。例如，`[0,1,2,4,5,6,7]` 在下标 3 处经旋转后可能变为 `[4,5,6,7,0,1,2]`。

给你旋转后的数组 `nums` 和一个整数 `target`，如果 `nums` 中存在这个目标值，则返回它的下标，否则返回 `-1`。你必须设计一个时间复杂度为 O(log n) 的算法。

## 示例

**示例 1：**

```
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
解释：0 在下标 4 处被找到。
```

**示例 2：**

```
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
解释：3 不在数组中，返回 -1。
```

**示例 3：**

```
输入：nums = [1], target = 1
输出：0
解释：数组只有一个元素，与目标值相等。
```

## 约束条件

- `1 <= nums.length <= 5000`
- `-10^4 <= nums[i] <= 10^4`
- `nums` 中的所有值互不相同
- `nums` 是一个可能经过旋转的升序数组
- `-10^4 <= target <= 10^4`

## 函数签名

```go
func searchRotated(nums []int, target int) int
```
