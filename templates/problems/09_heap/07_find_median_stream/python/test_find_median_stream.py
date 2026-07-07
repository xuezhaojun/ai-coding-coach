from find_median_stream import MedianFinder


def test_example_from_problem():
    mf = MedianFinder()
    mf.add_num(1)
    assert mf.find_median() == 1.0
    mf.add_num(2)
    assert mf.find_median() == 1.5
    mf.add_num(3)
    assert mf.find_median() == 2.0


def test_single_element():
    mf = MedianFinder()
    mf.add_num(5)
    assert mf.find_median() == 5.0


def test_two_elements():
    mf = MedianFinder()
    mf.add_num(1)
    assert mf.find_median() == 1.0
    mf.add_num(2)
    assert mf.find_median() == 1.5


def test_negative_numbers():
    mf = MedianFinder()
    mf.add_num(-1)
    assert mf.find_median() == -1.0
    mf.add_num(-2)
    assert mf.find_median() == -1.5
    mf.add_num(-3)
    assert mf.find_median() == -2.0


def test_duplicates():
    mf = MedianFinder()
    for _ in range(4):
        mf.add_num(1)
    assert mf.find_median() == 1.0


def test_descending_order():
    mf = MedianFinder()
    adds = [5, 4, 3, 2, 1]
    medians = [5.0, 4.5, 4.0, 3.5, 3.0]
    for num, expected in zip(adds, medians):
        mf.add_num(num)
        assert mf.find_median() == expected


def test_large_gap():
    mf = MedianFinder()
    mf.add_num(0)
    assert mf.find_median() == 0.0
    mf.add_num(100)
    assert mf.find_median() == 50.0
