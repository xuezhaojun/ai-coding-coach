from contains_duplicate import contains_duplicate


def test_has_duplicates():
    assert contains_duplicate([1, 2, 3, 1]) == True


def test_no_duplicates():
    assert contains_duplicate([1, 2, 3, 4]) == False


def test_empty_slice():
    assert contains_duplicate([]) == False


def test_single_element():
    assert contains_duplicate([1]) == False


def test_all_same_elements():
    assert contains_duplicate([5, 5, 5, 5]) == True


def test_duplicates_at_end():
    assert contains_duplicate([1, 2, 3, 4, 5, 5]) == True


def test_negative_numbers_with_duplicates():
    assert contains_duplicate([-1, -2, -3, -1]) == True
