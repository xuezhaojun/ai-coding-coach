from jump_game_ii import jump


def test_example_1():
    assert jump([2, 3, 1, 1, 4]) == 2


def test_example_2():
    assert jump([2, 3, 0, 1, 4]) == 2


def test_single_element():
    assert jump([0]) == 0


def test_two_elements():
    assert jump([1, 2]) == 1


def test_already_at_end():
    assert jump([1]) == 0


def test_linear_jumps():
    assert jump([1, 1, 1, 1]) == 3


def test_large_first_jump():
    assert jump([5, 1, 1, 1, 1, 1]) == 1


def test_greedy_per_level():
    assert jump([1, 2, 1, 1, 1]) == 3


def test_level_boundary():
    assert jump([7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3]) == 2
