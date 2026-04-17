# 单词接龙

- **难度**: 困难
- **分类**: 图
- **考点**: 广度优先搜索, 字符串, 哈希集合
- **链接**: [NeetCode](https://neetcode.io/problems/word-ladder) | [力扣 127](https://leetcode.cn/problems/word-ladder/)

## 题目描述

字典 `wordList` 中从单词 `beginWord` 到 `endWord` 的转换序列是一个按下述规格形成的序列 `beginWord -> s1 -> s2 -> ... -> sk`：

- 每一对相邻的单词只差一个字母。
- 对于 `1 <= i <= k` 时，每个 `si` 都在 `wordList` 中。注意，`beginWord` 不需要在 `wordList` 中。
- `sk == endWord`

给你两个单词 `beginWord` 和 `endWord` 和一个字典 `wordList`，返回从 `beginWord` 到 `endWord` 的最短转换序列中的单词数目，如果不存在这样的转换序列，返回 `0`。

## 示例

**示例 1:**

```
输入: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
输出: 5
解释: "hit" -> "hot" -> "dot" -> "dog" -> "cog"（序列中有 5 个单词）。
```

**示例 2:**

```
输入: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
输出: 0
解释: endWord "cog" 不在 wordList 中，所以无法进行转换。
```

**示例 3:**

```
输入: beginWord = "hot", endWord = "dot", wordList = ["dot"]
输出: 2
解释: "hot" -> "dot"（序列中有 2 个单词）。
```

## 约束条件

- `1 <= beginWord.length <= 10`
- `endWord.length == beginWord.length`
- `1 <= wordList.length <= 5000`
- `wordList[i].length == beginWord.length`
- `beginWord`、`endWord` 和 `wordList[i]` 由小写英文字母组成。
- `beginWord != endWord`
- `wordList` 中的所有单词互不相同。

## 函数签名

```go
func ladderLength(beginWord string, endWord string, wordList []string) int
```
