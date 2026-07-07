from jump_game import can_jump


def test_reachable():
    assert can_jump([2, 3, 1, 1, 4]) == True


def test_not_reachable():
    assert can_jump([3, 2, 1, 0, 4]) == False


def test_single_element():
    assert can_jump([0]) == True


def test_two_elements_reachable():
    assert can_jump([1, 0]) == True


def test_all_zeros_except_first():
    assert can_jump([5, 0, 0, 0, 0]) == True


def test_large_first_jump():
    assert can_jump([10, 0, 0, 0, 0, 0, 0, 0, 1]) == True


def test_stuck_at_zero():
    assert can_jump([1, 0, 1]) == False
