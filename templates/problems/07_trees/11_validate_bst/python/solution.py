from helpers import TreeNode, Optional


def solve_is_valid_bst(root: Optional[TreeNode]) -> bool:
    """Return true if the binary tree is a valid BST.

    Time: O(n), Space: O(h).
    """
    return _validate_bst(root, float('-inf'), float('inf'))


def _validate_bst(node: Optional[TreeNode], min_val: float, max_val: float) -> bool:
    if node is None:
        return True
    if node.val <= min_val or node.val >= max_val:
        return False
    return _validate_bst(node.left, min_val, node.val) and _validate_bst(node.right, node.val, max_val)
