from helpers import TreeNode, Optional


def solve_is_subtree(root: Optional[TreeNode], sub_root: Optional[TreeNode]) -> bool:
    """Return true if sub_root is a subtree of root.

    Time: O(m*n), Space: O(h).
    """
    if root is None:
        return sub_root is None
    if _same_tree(root, sub_root):
        return True
    return solve_is_subtree(root.left, sub_root) or solve_is_subtree(root.right, sub_root)


def _same_tree(p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
    if p is None and q is None:
        return True
    if p is None or q is None or p.val != q.val:
        return False
    return _same_tree(p.left, q.left) and _same_tree(p.right, q.right)
