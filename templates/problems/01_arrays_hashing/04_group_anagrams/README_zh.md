# 字母异位词分组

- **难度**：中等
- **分类**：数组与哈希
- **考点**：字符串、哈希表、排序
- **链接**：[NeetCode](https://neetcode.io/problems/anagram-groups) | [力扣 49](https://leetcode.cn/problems/group-anagrams/)

## 题目描述

给你一个字符串数组 `strs`，请你将字母异位词组合在一起。可以按任意顺序返回结果列表。

字母异位词是由重新排列源单词的所有字母得到的一个新单词，每个字母恰好使用一次。

## 示例

**示例 1：**

```
输入：strs = ["eat","tea","tan","ate","nat","bat"]
输出：[["bat"],["nat","tan"],["ate","eat","tea"]]
解释："eat"、"tea" 和 "ate" 互为字母异位词，"tan" 和 "nat" 互为字母异位词，"bat" 没有其他字母异位词。
```

**示例 2：**

```
输入：strs = [""]
输出：[[""]]
```

**示例 3：**

```
输入：strs = ["a"]
输出：[["a"]]
```

## 约束条件

- `1 <= strs.length <= 10^4`
- `0 <= strs[i].length <= 100`
- `strs[i]` 仅由小写英文字母组成。

## 函数签名

```go
func groupAnagrams(strs []string) [][]string
```
