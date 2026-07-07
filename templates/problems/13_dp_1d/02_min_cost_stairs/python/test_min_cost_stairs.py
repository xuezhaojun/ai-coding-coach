from min_cost_stairs import min_cost_climbing_stairs


def test_example_1():
    assert min_cost_climbing_stairs([10, 15, 20]) == 15


def test_example_2():
    assert min_cost_climbing_stairs([1, 100, 1, 1, 1, 100, 1, 1, 100, 1]) == 6


def test_two_steps_equal_cost():
    assert min_cost_climbing_stairs([5, 5]) == 5


def test_two_steps_different_cost():
    assert min_cost_climbing_stairs([1, 100]) == 1


def test_increasing_cost():
    assert min_cost_climbing_stairs([1, 2, 3, 4, 5]) == 6


def test_all_zeros():
    assert min_cost_climbing_stairs([0, 0, 0, 0]) == 0
