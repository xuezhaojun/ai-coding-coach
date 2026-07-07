from helpers import TreeNode, Optional


def solve_level_order(root: Optional[TreeNode]) -> list[list[int]]:
    """Return the level order traversal of a binary tree's values.

    Time: O(n), Space: O(n).
    """
    if root is None:
        return []
    result: list[list[int]] = []
    queue: list[TreeNode] = [root]
    while queue:
        size = len(queue)
        level: list[int] = []
        for _ in range(size):
            node = queue.pop(0)
            level.append(node.val)
            if node.left is not None:
                queue.append(node.left)
            if node.right is not None:
                queue.append(node.right)
        result.append(level)
    return result
