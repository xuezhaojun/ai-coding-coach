import pytest

from helpers import tree_to_slice
from build_tree import build_tree_from_pre_in


@pytest.mark.parametrize("name, preorder, inorder, want", [
    ("example 1", [3, 9, 20, 15, 7], [9, 3, 15, 20, 7], [3, 9, 20, -101, -101, 15, 7]),
    ("single node", [1], [1], [1]),
    ("left only", [1, 2], [2, 1], [1, 2]),
    ("right only", [1, 2], [1, 2], [1, -101, 2]),
    ("empty", [], [], []),
    ("three nodes", [1, 2, 3], [2, 1, 3], [1, 2, 3]),
])
def test_build_tree_from_pre_in(name, preorder, inorder, want):
    got = tree_to_slice(build_tree_from_pre_in(preorder, inorder))
    assert got == want
