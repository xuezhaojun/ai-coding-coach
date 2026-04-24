# 删除链表的倒数第 N 个结点

- **难度**：中等
- **分类**：链表
- **考点**：链表、双指针
- **链接**：[NeetCode](https://neetcode.io/problems/remove-node-from-end-of-linked-list) | [力扣 19](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/)

## 题目描述

给你一个链表，删除链表的倒数第 `n` 个结点，并且返回链表的头结点。

题目保证 `n` 总是有效的，即 `n` 总是小于或等于链表的长度。

![remove-nth-node-from-end-of-list](assets/remove_ex1.jpg)

## 示例

**示例 1：**

```
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
解释：倒数第 2 个节点是值为 4 的节点，将其删除。
```

**示例 2：**

```
输入：head = [1], n = 1
输出：[]
解释：唯一的节点被删除，返回空链表。
```

**示例 3：**

```
输入：head = [1,2], n = 1
输出：[1]
解释：最后一个节点（值为 2）被删除。
```

## 约束条件

- 链表中结点的数目为 `sz`
- `1 <= sz <= 30`
- `0 <= Node.Val <= 100`
- `1 <= n <= sz`

## 函数签名

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode
```
