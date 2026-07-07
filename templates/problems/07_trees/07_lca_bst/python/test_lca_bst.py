import pytest

from helpers import build_tree, TreeNode
from lca_bst import lowest_common_ancestor


def find_node(root, val):
    if root is None or root.val == val:
        return root
    left = find_node(root.left, val)
    if left is not None:
        return left
    return find_node(root.right, val)


@pytest.mark.parametrize("name, vals, p_val, q_val, want_v", [
    ("example 1", [6, 2, 8, 0, 4, 7, 9, -101, -101, 3, 5], 2, 8, 6),
    ("example 2", [6, 2, 8, 0, 4, 7, 9, -101, -101, 3, 5], 2, 4, 2),
    ("root is lca", [2, 1, 3], 1, 3, 2),
    ("same node", [2, 1, 3], 1, 1, 1),
    ("right subtree", [6, 2, 8, 0, 4, 7, 9], 7, 9, 8),
    ("deep nodes", [6, 2, 8, 0, 4, 7, 9, -101, -101, 3, 5], 3, 5, 4),
])
def test_lowest_common_ancestor(name, vals, p_val, q_val, want_v):
    root = build_tree(vals)
    p = find_node(root, p_val)
    q = find_node(root, q_val)
    got = lowest_common_ancestor(root, p, q)
    got_v = -1 if got is None else got.val
    assert got_v == want_v
