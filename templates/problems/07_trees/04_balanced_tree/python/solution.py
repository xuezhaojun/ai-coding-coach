from helpers import TreeNode, Optional


def solve_is_balanced(root: Optional[TreeNode]) -> bool:
    """Return true if the binary tree is height-balanced.

    Time: O(n), Space: O(h).
    """
    return _height(root) != -1


def _height(node: Optional[TreeNode]) -> int:
    if node is None:
        return 0
    left = _height(node.left)
    if left == -1:
        return -1
    right = _height(node.right)
    if right == -1:
        return -1
    diff = left - right
    if diff < -1 or diff > 1:
        return -1
    return left + 1 if left > right else right + 1
