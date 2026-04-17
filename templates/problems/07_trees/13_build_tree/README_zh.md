# 从前序与中序遍历序列构造二叉树

- **难度**: 中等
- **分类**: 树
- **考点**: 二叉树, 递归, 哈希表, 分治
- **链接**: [NeetCode](https://neetcode.io/problems/binary-tree-from-preorder-and-inorder-traversal) | [力扣 105](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

## 题目描述

给定两个整数数组 `preorder` 和 `inorder`，其中 `preorder` 是二叉树的前序遍历，`inorder` 是同一棵树的中序遍历，请构造并返回该二叉树。保证数组中的值唯一且两个数组长度相同。

## 示例

**示例 1:**

```
输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]
解释: 根节点为 3（前序遍历的第一个元素）。在中序遍历中，9 在 3 的左边（左子树），[15,20,7] 在 3 的右边（右子树）。
```

**示例 2:**

```
输入: preorder = [1], inorder = [1]
输出: [1]
解释: 只有一个节点的树。
```

**示例 3:**

```
输入: preorder = [1,2], inorder = [2,1]
输出: [1,2]
解释: 节点 2 在中序遍历中出现在 1 之前，因此 2 是 1 的左子节点。
```

## 约束条件

- `1 <= preorder.length <= 3000`
- `inorder.length == preorder.length`
- `-3000 <= preorder[i], inorder[i] <= 3000`
- `preorder` 和 `inorder` 均由唯一值组成。
- `inorder` 中的每个值也出现在 `preorder` 中。
- `preorder` 保证是前序遍历，`inorder` 保证是同一棵树的中序遍历。

## 函数签名

```go
func buildTreeFromPreIn(preorder []int, inorder []int) *TreeNode
```
