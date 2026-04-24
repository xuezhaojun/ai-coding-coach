# 盛最多水的容器

- **难度**：中等
- **分类**：双指针
- **考点**：数组、双指针、贪心
- **链接**：[NeetCode](https://neetcode.io/problems/max-water-container) | [力扣 11](https://leetcode.cn/problems/container-with-most-water/)

## 题目描述

给定一个长度为 `n` 的整数数组 `height`。有 `n` 条垂线，第 `i` 条线的两个端点是 `(i, 0)` 和 `(i, height[i])`。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。返回容器可以储存的最大水量。

注意：你不能倾斜容器。

![container-with-most-water](assets/question_11.jpg)

## 示例

**示例 1：**

```
输入：height = [1,8,6,2,5,4,8,3,7]
输出：49
解释：下标 1（高度 8）和下标 8（高度 7）处的线段构成的容器宽度为 7。面积为 min(8,7) * 7 = 49。
```

**示例 2：**

```
输入：height = [1,1]
输出：1
```

**示例 3：**

```
输入：height = [4,3,2,1,4]
输出：16
解释：下标 0（高度 4）和下标 4（高度 4）处的线段构成的容器宽度为 4。面积为 min(4,4) * 4 = 16。
```

## 约束条件

- `n == height.length`
- `2 <= n <= 10^5`
- `0 <= height[i] <= 10^4`

## 函数签名

```go
func maxArea(height []int) int
```
