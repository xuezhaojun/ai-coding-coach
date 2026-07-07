from helpers import TreeNode, Optional


def solve_right_side_view(root: Optional[TreeNode]) -> list[int]:
    """Return the values of the nodes visible from the right side.

    Time: O(n), Space: O(n).
    """
    if root is None:
        return []
    result: list[int] = []
    queue: list[TreeNode] = [root]
    while queue:
        size = len(queue)
        for i in range(size):
            node = queue.pop(0)
            if i == size - 1:
                result.append(node.val)
            if node.left is not None:
                queue.append(node.left)
            if node.right is not None:
                queue.append(node.right)
    return result
