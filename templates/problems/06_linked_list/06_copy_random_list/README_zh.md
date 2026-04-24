# 复制带随机指针的链表

- **难度**：中等
- **分类**：链表
- **考点**：链表、哈希表、深拷贝
- **链接**：[NeetCode](https://neetcode.io/problems/copy-linked-list-with-random-pointer) | [力扣 138](https://leetcode.cn/problems/copy-list-with-random-pointer/)

## 题目描述

给你一个长度为 `n` 的链表，每个节点包含一个额外增加的随机指针 `random`，该指针可以指向链表中的任何节点或空节点。

构造这个链表的深拷贝。深拷贝应该正好由 `n` 个全新节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 `next` 指针和 `random` 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点。

返回复制链表的头节点。

## 示例

**示例 1：**


![示例 1](assets/e1.png)
```
输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
解释：每对 [val, random_index] 描述一个节点。深拷贝保持相同的结构。
```

**示例 2：**


![示例 2](assets/e2.png)
```
输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]
解释：节点 0 的 random 指向节点 1，节点 1 的 random 也指向节点 1。
```

**示例 3：**


![示例 3](assets/e3.png)
```
输入：head = [[3,null],[3,0],[3,null]]
输出：[[3,null],[3,0],[3,null]]
解释：节点可以有相同的值。随机指针可以为 null。
```

## 约束条件

- `0 <= n <= 1000`
- `-10^4 <= Node.Val <= 10^4`
- `Node.Random` 为 `null` 或指向链表中的某个节点

## 函数签名

```go
func copyRandomList(head *RandomNode) *RandomNode
```
