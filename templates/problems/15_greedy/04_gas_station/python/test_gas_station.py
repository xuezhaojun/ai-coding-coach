from gas_station import can_complete_circuit


def test_example_1():
    assert can_complete_circuit([1, 2, 3, 4, 5], [3, 4, 5, 1, 2]) == 3


def test_no_solution():
    assert can_complete_circuit([2, 3, 4], [3, 4, 3]) == -1


def test_single_station_enough():
    assert can_complete_circuit([5], [3]) == 0


def test_single_station_not_enough():
    assert can_complete_circuit([2], [4]) == -1


def test_start_at_index_0():
    assert can_complete_circuit([3, 1, 1], [1, 2, 2]) == 0


def test_all_equal():
    assert can_complete_circuit([1, 1, 1], [1, 1, 1]) == 0


def test_start_at_last():
    assert can_complete_circuit([1, 1, 5], [2, 3, 1]) == 2
