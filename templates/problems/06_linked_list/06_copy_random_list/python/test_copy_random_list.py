from helpers import build_random_list, random_list_to_slice
from copy_random_list import copy_random_list


def _verify(vals, random_indices):
    head = build_random_list(vals, random_indices)
    copied = copy_random_list(head)
    if head is None:
        assert copied is None
        return
    got_vals, got_randoms = random_list_to_slice(copied)
    assert got_vals == list(vals)
    assert got_randoms == list(random_indices)
    # verify deep copy (no shared nodes)
    orig_cur = head
    copy_cur = copied
    while orig_cur is not None:
        assert orig_cur is not copy_cur
        orig_cur = orig_cur.next
        copy_cur = copy_cur.next  # type: ignore[union-attr]


def test_nil_list():
    _verify([], [])


def test_single_no_random():
    _verify([1], [-1])


def test_single_self_random():
    _verify([1], [0])


def test_two_nodes():
    _verify([1, 2], [1, 0])


def test_three_nodes_mixed():
    _verify([7, 13, 11], [-1, 0, 2])


def test_all_random_nil():
    _verify([1, 2, 3, 4], [-1, -1, -1, -1])


def test_chain_random():
    _verify([1, 2, 3], [2, 0, 1])
