# 字符串相乘

- **难度**: 中等
- **分类**: 数学与几何
- **考点**: 数学, 字符串, 模拟
- **链接**: [NeetCode](https://neetcode.io/problems/multiply-strings) | [力扣 43](https://leetcode.cn/problems/multiply-strings/)

## 题目描述

给定两个以字符串形式表示的非负整数 `num1` 和 `num2`，返回 `num1` 和 `num2` 的乘积，其结果也用字符串表示。不能使用任何内置的大整数库或直接将输入转换为整数，必须逐位进行乘法运算。

## 示例

**示例 1:**

```
输入: num1 = "2", num2 = "3"
输出: "6"
解释: 2 * 3 = 6。
```

**示例 2:**

```
输入: num1 = "123", num2 = "456"
输出: "56088"
解释: 123 * 456 = 56088。
```

**示例 3:**

```
输入: num1 = "0", num2 = "52"
输出: "0"
解释: 任何数乘以 0 都为 0。
```

## 约束条件

- `1 <= num1.length, num2.length <= 200`
- `num1` 和 `num2` 只包含数字
- `num1` 和 `num2` 不包含任何前导零，除了数字 `0` 本身

## 函数签名

```go
func multiply(num1 string, num2 string) string
```
