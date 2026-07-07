from longest_increasing_sub import length_of_lis


def test_example_1():
    assert length_of_lis([10, 9, 2, 5, 3, 7, 101, 18]) == 4


def test_all_increasing():
    assert length_of_lis([1, 2, 3, 4, 5]) == 5


def test_all_decreasing():
    assert length_of_lis([5, 4, 3, 2, 1]) == 1


def test_single_element():
    assert length_of_lis([7]) == 1


def test_example_2():
    assert length_of_lis([0, 1, 0, 3, 2, 3]) == 4


def test_duplicates():
    assert length_of_lis([7, 7, 7, 7, 7]) == 1


def test_valley_then_rise():
    assert length_of_lis([1, 3, 6, 7, 9, 4, 10, 5, 6]) == 6
