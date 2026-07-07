import pytest

from helpers import build_tree
from same_tree import is_same_tree


@pytest.mark.parametrize("name, p, q, want", [
    ("identical", [1, 2, 3], [1, 2, 3], True),
    ("different structure", [1, 2], [1, -101, 2], False),
    ("different values", [1, 2, 1], [1, 1, 2], False),
    ("both empty", [], [], True),
    ("one empty", [1], [], False),
    ("single same", [1], [1], True),
])
def test_is_same_tree(name, p, q, want):
    p_root = build_tree(p)
    q_root = build_tree(q)
    assert is_same_tree(p_root, q_root) == want
