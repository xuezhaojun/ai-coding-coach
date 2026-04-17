# 合并 K 个升序链表

- **难度**：困难
- **分类**：链表
- **考点**：链表、堆、分治
- **链接**：[NeetCode](https://neetcode.io/problems/merge-k-sorted-linked-lists) | [力扣 23](https://leetcode.cn/problems/merge-k-sorted-lists/)

## 题目描述

给你一个链表数组，每个链表都已经按升序排列。请你将所有链表合并到一个升序链表中，返回合并后的链表。

## 示例

**示例 1：**

```
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：三个有序链表通过比较头节点、每次取最小值的方式合并为一个有序链表。
```

**示例 2：**

```
输入：lists = []
输出：[]
解释：空的链表数组产生空结果。
```

**示例 3：**

```
输入：lists = [[]]
输出：[]
解释：单个空链表合并后仍为空链表。
```

## 约束条件

- `k == lists.length`
- `0 <= k <= 10^4`
- `0 <= lists[i].length <= 500`
- `-10^4 <= lists[i][j] <= 10^4`
- `lists[i]` 按升序排列
- `lists[i].length` 的总和不超过 `10^4`

## 函数签名

```go
func mergeKLists(lists []*ListNode) *ListNode
```
