# 验证回文串

- **难度**：简单
- **分类**：双指针
- **考点**：字符串、双指针
- **链接**：[NeetCode](https://neetcode.io/problems/is-palindrome) | [力扣 125](https://leetcode.cn/problems/valid-palindrome/)

## 题目描述

如果在将所有大写字母转换为小写字母、并移除所有非字母数字字符之后，短语正着读和反着读都一样，则可以认为该短语是一个回文串。字母和数字都属于字母数字字符。

给你一个字符串 `s`，如果它是回文串，返回 `true`；否则，返回 `false`。

## 示例

**示例 1：**

```
输入：s = "A man, a plan, a canal: Panama"
输出：true
解释：移除非字母数字字符并转为小写后，"amanaplanacanalpanama" 是回文串。
```

**示例 2：**

```
输入：s = "race a car"
输出：false
解释：处理后，"raceacar" 不是回文串。
```

**示例 3：**

```
输入：s = " "
输出：true
解释：移除非字母数字字符后，s 为空字符串 ""。空字符串正着读和反着读都一样，所以是回文串。
```

## 约束条件

- `1 <= s.length <= 2 * 10^5`
- `s` 仅由可打印的 ASCII 字符组成。

## 函数签名

```go
func isPalindrome(s string) bool
```
