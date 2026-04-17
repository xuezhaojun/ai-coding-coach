# 字符串编码与解码

- **难度**：中等
- **分类**：数组与哈希
- **考点**：字符串、设计
- **链接**：[NeetCode](https://neetcode.io/problems/string-encode-and-decode) | [力扣 271](https://leetcode.cn/problems/encode-and-decode-strings/)

## 题目描述

请你设计一个算法，将一个字符串列表编码为一个单独的字符串。编码后的字符串可以被解码回原始的字符串列表。

请实现 `encode` 和 `decode` 两个函数。编码后的字符串需要能够处理任何可能的输入，包括包含特殊字符、分隔符或为空的字符串。编码必须是无损的——解码编码后的字符串必须产生完全相同的原始字符串列表。

## 示例

**示例 1：**

```
输入：["hello","world"]
输出：["hello","world"]
解释：encode(["hello","world"]) 可能生成 "5#hello5#world"，而 decode("5#hello5#world") 返回 ["hello","world"]。
```

**示例 2：**

```
输入：[""]
输出：[""]
解释：列表包含一个空字符串，必须在编码和解码过程中保留。
```

**示例 3：**

```
输入：["he:llo","wor#ld","foo;bar"]
输出：["he:llo","wor#ld","foo;bar"]
解释：字符串中的特殊字符必须被正确处理。编码不能被看起来像分隔符的字符所混淆。
```

## 约束条件

- `0 <= strs.length <= 200`
- `0 <= strs[i].length <= 200`
- `strs[i]` 可以包含 256 个有效 ASCII 字符中的任意字符。

## 函数签名

```go
func encode(strs []string) string
func decode(s string) []string
```
