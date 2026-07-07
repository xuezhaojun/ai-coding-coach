from median_two_sorted import find_median_sorted_arrays


def test_odd_total():
    assert find_median_sorted_arrays([1, 3], [2]) == 2.0


def test_even_total():
    assert find_median_sorted_arrays([1, 2], [3, 4]) == 2.5


def test_first_empty():
    assert find_median_sorted_arrays([], [1]) == 1.0


def test_second_empty():
    assert find_median_sorted_arrays([2], []) == 2.0


def test_no_overlap():
    assert find_median_sorted_arrays([1, 2], [3, 4, 5]) == 3.0


def test_interleaved():
    assert find_median_sorted_arrays([1, 3, 5], [2, 4, 6]) == 3.5


def test_single_elements():
    assert find_median_sorted_arrays([1], [2]) == 1.5


def test_duplicates():
    assert find_median_sorted_arrays([1, 1, 1], [1, 1, 1]) == 1.0
