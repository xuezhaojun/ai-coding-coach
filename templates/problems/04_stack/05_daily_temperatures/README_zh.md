# 每日温度

- **难度**：中等
- **分类**：栈
- **考点**：栈、单调栈、数组
- **链接**：[NeetCode](https://neetcode.io/problems/daily-temperatures) | [力扣 739](https://leetcode.cn/problems/daily-temperatures/)

## 题目描述

给定一个整数数组 `temperatures`，表示每天的温度，返回一个数组 `answer`，其中 `answer[i]` 是指对于第 `i` 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 `0` 来代替。

## 示例

**示例 1：**

```
输入：temperatures = [73,74,75,71,69,72,76,73]
输出：[1,1,4,2,1,1,0,0]
解释：第 0 天，第 1 天温度更高（74 > 73），所以 answer[0] = 1。第 2 天，下一个更高温度在第 6 天（76 > 75），所以 answer[2] = 4。
```

**示例 2：**

```
输入：temperatures = [30,20,10]
输出：[0,0,0]
解释：温度严格递减，没有任何一天有更高的未来温度。
```

**示例 3：**

```
输入：temperatures = [10,20,30]
输出：[1,1,0]
解释：每一天（除了最后一天）的下一天就是更高温度。
```

## 约束条件

- `1 <= temperatures.length <= 10^5`
- `30 <= temperatures[i] <= 100`

## 函数签名

```go
func dailyTemperatures(temperatures []int) []int
```
