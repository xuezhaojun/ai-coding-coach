# 接雨水

- **难度**：困难
- **分类**：双指针
- **考点**：数组、双指针、动态规划、栈
- **链接**：[NeetCode](https://neetcode.io/problems/trapping-rain-water) | [力扣 42](https://leetcode.cn/problems/trapping-rain-water/)

## 题目描述

给定 `n` 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

每个柱子上方能接住的水量取决于其左侧最大高度和右侧最大高度中的较小值，减去该柱子本身的高度。如果该值为负数，则该位置不接水。

![trapping-rain-water](assets/rainwatertrap.png)

## 示例

**示例 1：**

```
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面的高度图表示的柱子可以接住 6 个单位的雨水。水填充了柱子之间的间隙。
```

**示例 2：**

```
输入：height = [4,2,0,3,2,5]
输出：9
```

**示例 3：**

```
输入：height = [1,2,3,4,5]
输出：0
解释：柱子严格递增，无法接住任何雨水。
```

## 约束条件

- `n == height.length`
- `1 <= n <= 2 * 10^4`
- `0 <= height[i] <= 10^5`

## 函数签名

```go
func trap(height []int) int
```
