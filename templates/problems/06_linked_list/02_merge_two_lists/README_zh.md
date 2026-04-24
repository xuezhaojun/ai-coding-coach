# 合并两个有序链表

- **难度**：简单
- **分类**：链表
- **考点**：链表、递归、迭代
- **链接**：[NeetCode](https://neetcode.io/problems/merge-two-sorted-linked-lists) | [力扣 21](https://leetcode.cn/problems/merge-two-sorted-lists/)

## 题目描述

将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

返回合并后的链表的头节点。

![merge-two-sorted-lists](assets/merge_ex1.jpg)

## 示例

**示例 1：**

```
输入：list1 = [1,3,5], list2 = [2,4,6]
输出：[1,2,3,4,5,6]
解释：两个有序链表交替合并为一个有序链表。
```

**示例 2：**

```
输入：list1 = [], list2 = [1,2,3]
输出：[1,2,3]
解释：当一个链表为空时，结果为另一个链表。
```

**示例 3：**

```
输入：list1 = [], list2 = []
输出：[]
解释：合并两个空链表得到空链表。
```

## 约束条件

- 两个链表的节点数目范围是 `[0, 50]`
- `-100 <= Node.Val <= 100`
- `list1` 和 `list2` 均按非递减顺序排列

## 函数签名

```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode
```
