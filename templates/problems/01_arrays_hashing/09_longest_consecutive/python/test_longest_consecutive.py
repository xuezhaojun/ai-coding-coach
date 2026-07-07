from longest_consecutive import longest_consecutive


def test_basic_case():
    assert longest_consecutive([100, 4, 200, 1, 3, 2]) == 4


def test_longer_sequence():
    assert longest_consecutive([0, 3, 7, 2, 5, 8, 4, 6, 0, 1]) == 9


def test_empty_array():
    assert longest_consecutive([]) == 0


def test_single_element():
    assert longest_consecutive([1]) == 1


def test_no_consecutive():
    assert longest_consecutive([10, 20, 30]) == 1


def test_duplicates_in_sequence():
    assert longest_consecutive([1, 2, 2, 3, 3, 4]) == 4


def test_negative_numbers():
    assert longest_consecutive([-3, -2, -1, 0, 1]) == 5


def test_all_same_elements():
    assert longest_consecutive([5, 5, 5, 5]) == 1
