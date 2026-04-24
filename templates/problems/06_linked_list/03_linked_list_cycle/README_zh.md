# 环形链表

- **难度**：简单
- **分类**：链表
- **考点**：链表、双指针、Floyd 判圈算法
- **链接**：[NeetCode](https://neetcode.io/problems/linked-list-cycle-detection) | [力扣 141](https://leetcode.cn/problems/linked-list-cycle/)

## 题目描述

给你一个链表的头节点 `head`，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 `next` 指针再次到达，则链表中存在环。为了表示给定链表中的环，评测系统内部使用整数 `pos` 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：`pos` 不作为参数进行传递。

如果链表中存在环，则返回 `true`。否则，返回 `false`。

## 示例

**示例 1：**


![示例 1](assets/circularlinkedlist.png)
```
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第 1 个节点（索引从 0 开始）。
```

**示例 2：**


![示例 2](assets/circularlinkedlist_test2.png)
```
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第 0 个节点。
```

**示例 3：**


![示例 3](assets/circularlinkedlist_test3.png)
```
输入：head = [1], pos = -1
输出：false
解释：链表中没有环。
```

## 约束条件

- 链表中节点的数目范围是 `[0, 10^4]`
- `-10^5 <= Node.Val <= 10^5`
- `pos` 为 `-1` 或者链表中的一个有效索引

## 函数签名

```go
func hasCycle(head *ListNode) bool
```
