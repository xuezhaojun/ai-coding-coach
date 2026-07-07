from container_most_water import max_area


def test_basic_case():
    assert max_area([1, 8, 6, 2, 5, 4, 8, 3, 7]) == 49


def test_two_elements():
    assert max_area([1, 1]) == 1


def test_decreasing_heights():
    assert max_area([4, 3, 2, 1, 4]) == 16


def test_equal_heights():
    assert max_area([5, 5, 5, 5]) == 15


def test_one_tall_wall():
    assert max_area([1, 2, 1]) == 2


def test_large_difference():
    assert max_area([1, 1000, 1000, 1]) == 1000
