from subsets import subsets


def sort_subsets(result: list[list[int]]) -> list[list[int]]:
    for s in result:
        s.sort()
    result.sort(key=lambda g: (len(g), g))
    return result


def test_three_elements():
    got = subsets([1, 2, 3])
    expected = [[], [1], [2], [3], [1, 2], [1, 3], [2, 3], [1, 2, 3]]
    assert sort_subsets(got) == sort_subsets(expected)


def test_single_element():
    got = subsets([0])
    expected = [[], [0]]
    assert sort_subsets(got) == sort_subsets(expected)


def test_two_elements():
    got = subsets([1, 2])
    expected = [[], [1], [2], [1, 2]]
    assert sort_subsets(got) == sort_subsets(expected)


def test_empty_input():
    got = subsets([])
    expected = [[]]
    assert sort_subsets(got) == sort_subsets(expected)


def test_negative_numbers():
    got = subsets([-1, 0])
    expected = [[], [-1], [0], [-1, 0]]
    assert sort_subsets(got) == sort_subsets(expected)


def test_five_elements_slice_sharing_regression():
    got = subsets([9, 0, 3, 5, 7])
    expected = [
        [], [9], [0], [3], [5], [7],
        [9, 0], [9, 3], [9, 5], [9, 7], [0, 3], [0, 5], [0, 7], [3, 5], [3, 7], [5, 7],
        [9, 0, 3], [9, 0, 5], [9, 0, 7], [9, 3, 5], [9, 3, 7], [9, 5, 7], [0, 3, 5], [0, 3, 7], [0, 5, 7], [3, 5, 7],
        [9, 0, 3, 5], [9, 0, 3, 7], [9, 0, 5, 7], [9, 3, 5, 7], [0, 3, 5, 7],
        [9, 0, 3, 5, 7],
    ]
    assert sort_subsets(got) == sort_subsets(expected)


def test_four_elements():
    got = subsets([1, 2, 3, 4])
    expected = [
        [], [1], [2], [3], [4],
        [1, 2], [1, 3], [1, 4], [2, 3], [2, 4], [3, 4],
        [1, 2, 3], [1, 2, 4], [1, 3, 4], [2, 3, 4],
        [1, 2, 3, 4],
    ]
    assert sort_subsets(got) == sort_subsets(expected)
