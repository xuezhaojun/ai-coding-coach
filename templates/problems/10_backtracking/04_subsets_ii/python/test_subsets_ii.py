from subsets_ii import subsets_with_dup


def sort_subsets(result: list[list[int]]) -> list[list[int]]:
    for s in result:
        s.sort()
    result.sort(key=lambda g: (len(g), g))
    return result


def test_example_with_duplicates():
    got = subsets_with_dup([1, 2, 2])
    expected = [[], [1], [2], [1, 2], [2, 2], [1, 2, 2]]
    assert sort_subsets(got) == sort_subsets(expected)


def test_single_element():
    got = subsets_with_dup([0])
    expected = [[], [0]]
    assert sort_subsets(got) == sort_subsets(expected)


def test_all_duplicates():
    got = subsets_with_dup([1, 1, 1])
    expected = [[], [1], [1, 1], [1, 1, 1]]
    assert sort_subsets(got) == sort_subsets(expected)


def test_no_duplicates():
    got = subsets_with_dup([1, 2, 3])
    expected = [[], [1], [2], [3], [1, 2], [1, 3], [2, 3], [1, 2, 3]]
    assert sort_subsets(got) == sort_subsets(expected)


def test_two_pairs_of_duplicates():
    got = subsets_with_dup([1, 1, 2, 2])
    expected = [
        [], [1], [2], [1, 1], [1, 2], [2, 2],
        [1, 1, 2], [1, 2, 2], [1, 1, 2, 2],
    ]
    assert sort_subsets(got) == sort_subsets(expected)


def test_unsorted_input_with_duplicates():
    got = subsets_with_dup([2, 1, 2])
    expected = [[], [1], [2], [1, 2], [2, 2], [1, 2, 2]]
    assert sort_subsets(got) == sort_subsets(expected)


def test_empty_input():
    got = subsets_with_dup([])
    expected = [[]]
    assert sort_subsets(got) == sort_subsets(expected)
