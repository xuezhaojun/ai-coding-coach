from kth_largest_stream import KthLargest


def test_example_from_problem():
    kl = KthLargest(3, [4, 5, 8, 2])
    adds = [3, 5, 10, 9, 4]
    expects = [4, 5, 5, 8, 8]
    for val, expected in zip(adds, expects):
        assert kl.add(val) == expected


def test_k_equals_1():
    kl = KthLargest(1, [])
    adds = [-1, 1, -2, -4, 3]
    expects = [-1, 1, 1, 1, 3]
    for val, expected in zip(adds, expects):
        assert kl.add(val) == expected


def test_single_initial_element():
    kl = KthLargest(1, [5])
    adds = [3, 7]
    expects = [5, 7]
    for val, expected in zip(adds, expects):
        assert kl.add(val) == expected


def test_all_same_values():
    kl = KthLargest(2, [0])
    adds = [0, 0, 0]
    expects = [0, 0, 0]
    for val, expected in zip(adds, expects):
        assert kl.add(val) == expected


def test_negative_numbers():
    kl = KthLargest(2, [-5, -3])
    adds = [-1, -7, 0]
    expects = [-3, -3, -1]
    for val, expected in zip(adds, expects):
        assert kl.add(val) == expected


def test_large_k_with_enough_initial_elements():
    kl = KthLargest(3, [1, 2, 3, 4, 5])
    assert kl.add(6) == 4
