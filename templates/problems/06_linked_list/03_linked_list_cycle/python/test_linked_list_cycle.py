from helpers import build_cyclic_list
from linked_list_cycle import has_cycle


def test_nil_list():
    assert has_cycle(build_cyclic_list([], -1)) is False


def test_single_no_cycle():
    assert has_cycle(build_cyclic_list([1], -1)) is False


def test_single_self_cycle():
    assert has_cycle(build_cyclic_list([1], 0)) is True


def test_cycle_at_head():
    assert has_cycle(build_cyclic_list([3, 2, 0, -4], 0)) is True


def test_cycle_at_middle():
    assert has_cycle(build_cyclic_list([3, 2, 0, -4], 1)) is True


def test_no_cycle():
    assert has_cycle(build_cyclic_list([1, 2, 3, 4, 5], -1)) is False


def test_two_nodes_cycle():
    assert has_cycle(build_cyclic_list([1, 2], 0)) is True
