from hand_of_straights import is_n_straight_hand


def test_example_true():
    assert is_n_straight_hand([1, 2, 3, 6, 2, 3, 4, 7, 8], 3) == True


def test_example_false():
    assert is_n_straight_hand([1, 2, 3, 4, 5], 4) == False


def test_single_group():
    assert is_n_straight_hand([1, 2, 3], 3) == True


def test_group_size_1():
    assert is_n_straight_hand([5, 3, 1], 1) == True


def test_not_divisible():
    assert is_n_straight_hand([1, 2, 3, 4], 3) == False


def test_duplicates_needed():
    assert is_n_straight_hand([1, 1, 2, 2, 3, 3], 3) == True


def test_gap_in_sequence():
    assert is_n_straight_hand([1, 2, 4, 5, 6, 7], 3) == False
