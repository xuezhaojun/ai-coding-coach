from __future__ import annotations
from typing import Optional


class TreeNode:
    def __init__(self, val: int = 0, left: Optional['TreeNode'] = None, right: Optional['TreeNode'] = None):
        self.val = val
        self.left = left
        self.right = right


def build_tree(vals: list[int]) -> Optional[TreeNode]:
    if len(vals) == 0 or vals[0] == -101:
        return None
    root = TreeNode(vals[0])
    queue: list[TreeNode] = [root]
    i = 1
    while queue and i < len(vals):
        node = queue.pop(0)
        if i < len(vals) and vals[i] != -101:
            node.left = TreeNode(vals[i])
            queue.append(node.left)
        i += 1
        if i < len(vals) and vals[i] != -101:
            node.right = TreeNode(vals[i])
            queue.append(node.right)
        i += 1
    return root


def tree_to_slice(root: Optional[TreeNode]) -> list[int]:
    if root is None:
        return []
    result: list[int] = []
    queue: list[Optional[TreeNode]] = [root]
    while queue:
        node = queue.pop(0)
        if node is None:
            result.append(-101)
        else:
            result.append(node.val)
            queue.append(node.left)
            queue.append(node.right)
    while result and result[-1] == -101:
        result.pop()
    return result
