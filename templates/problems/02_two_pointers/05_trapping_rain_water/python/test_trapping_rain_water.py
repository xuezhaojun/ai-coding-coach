from trapping_rain_water import trap


def test_basic_case():
    assert trap([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]) == 6


def test_v_shape():
    assert trap([4, 2, 0, 3, 2, 5]) == 9


def test_no_water():
    assert trap([1, 2, 3, 4, 5]) == 0


def test_empty_input():
    assert trap([]) == 0


def test_single_bar():
    assert trap([5]) == 0


def test_two_bars():
    assert trap([3, 1]) == 0


def test_flat_surface():
    assert trap([3, 3, 3, 3]) == 0
