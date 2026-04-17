# 检测正方形

- **难度**: 中等
- **分类**: 数学与几何
- **考点**: 数组, 哈希表, 设计, 计数
- **链接**: [NeetCode](https://neetcode.io/problems/count-squares) | [力扣 2013](https://leetcode.cn/problems/detect-squares/)

## 题目描述

给你一个在 X-Y 平面上的点流。设计一个数据结构，支持以下操作：

- **添加**一个新的点到数据结构中。允许重复点，且应分别计数。
- **计算**以给定查询点为四个角之一，能形成多少个轴对齐正方形。轴对齐正方形的所有边都平行于 X 轴和 Y 轴，且面积为正（边长 > 0）。

实现 `DetectSquares` 结构体，包括 `Constructor()`、`Add(point []int)` 和 `Count(point []int) int`。

## 示例

**示例 1:**

```
输入:
  添加点: [3,10], [11,10], [11,2]
  查询: [3,2]
输出: 1
解释: 点 [3,10]、[11,10]、[11,2] 和查询点 [3,2] 形成一个边长为 8 的轴对齐正方形。
```

**示例 2:**

```
输入:
  添加点: [3,10], [3,10], [11,10], [11,2]
  查询: [3,2]
输出: 2
解释: 重复的点 [3,10] 使得可以计数两个不同的正方形。
```

**示例 3:**

```
输入:
  添加点: [1,1], [2,2]
  查询: [3,3]
输出: 0
解释: 无法用查询点形成轴对齐正方形。
```

## 约束条件

- `point.length == 2`
- `0 <= x, y <= 1000`
- `Add` 和 `Count` 的总调用次数最多为 `3000` 次

## 函数签名

```go
type DetectSquares struct {}

func Constructor() DetectSquares

func (ds *DetectSquares) Add(point []int)

func (ds *DetectSquares) Count(point []int) int
```
