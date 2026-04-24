# 重排链表

- **难度**：中等
- **分类**：链表
- **考点**：链表、双指针、反转
- **链接**：[NeetCode](https://neetcode.io/problems/reorder-linked-list) | [力扣 143](https://leetcode.cn/problems/reorder-list/)

## 题目描述

给定一个单链表的头节点 `head`，链表可以表示为：

`L0 -> L1 -> ... -> Ln-1 -> Ln`

请将其重新排列为：

`L0 -> Ln -> L1 -> Ln-1 -> L2 -> Ln-2 -> ...`

不能只是单纯地改变节点内部的值，而是需要实际进行节点交换。必须原地完成重排。

## 示例

**示例 1：**


![示例 1](assets/reorder1linked-list.jpg)
```
输入：head = [1,2,3,4]
输出：[1,4,2,3]
解释：从两端交替取节点进行重排。
```

**示例 2：**


![示例 2](assets/reorder2-linked-list.jpg)
```
输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]
解释：重排顺序：第一个、最后一个、第二个、倒数第二个、中间。
```

**示例 3：**

```
输入：head = [1,2,3]
输出：[1,3,2]
解释：L0->L2->L1。
```

## 约束条件

- 链表的长度范围为 `[1, 5 * 10^4]`
- `1 <= Node.Val <= 1000`

## 函数签名

```go
func reorderList(head *ListNode)
```
