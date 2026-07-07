from last_stone_weight import last_stone_weight


def test_example_from_problem():
    assert last_stone_weight([2, 7, 4, 1, 8, 1]) == 1


def test_single_stone():
    assert last_stone_weight([1]) == 1


def test_two_equal_stones():
    assert last_stone_weight([3, 3]) == 0


def test_two_different_stones():
    assert last_stone_weight([1, 5]) == 4


def test_all_same_weight():
    assert last_stone_weight([2, 2, 2, 2]) == 0


def test_descending_weights():
    assert last_stone_weight([10, 5, 3, 1]) == 1


def test_three_stones():
    assert last_stone_weight([3, 7, 2]) == 2
