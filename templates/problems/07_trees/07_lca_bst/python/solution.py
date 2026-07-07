from helpers import TreeNode, Optional


def solve_lowest_common_ancestor(root: Optional[TreeNode], p: Optional[TreeNode], q: Optional[TreeNode]) -> Optional[TreeNode]:
    """Find the lowest common ancestor of two nodes in a BST.

    Time: O(h), Space: O(1).
    """
    cur = root
    while cur is not None:
        if p.val < cur.val and q.val < cur.val:
            cur = cur.left
        elif p.val > cur.val and q.val > cur.val:
            cur = cur.right
        else:
            return cur
    return None
