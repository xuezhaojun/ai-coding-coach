from combination_sum import combination_sum


def sort_combinations(result: list[list[int]]) -> list[list[int]]:
    for c in result:
        c.sort()
    result.sort(key=lambda g: (len(g), g))
    return result


def test_example_1():
    got = combination_sum([2, 3, 6, 7], 7)
    expected = [[2, 2, 3], [7]]
    assert sort_combinations(got) == sort_combinations(expected)


def test_example_2():
    got = combination_sum([2, 3, 5], 8)
    expected = [[2, 2, 2, 2], [2, 3, 3], [3, 5]]
    assert sort_combinations(got) == sort_combinations(expected)


def test_no_combination():
    got = combination_sum([2], 1)
    assert sort_combinations(got) == sort_combinations([])


def test_single_candidate_equals_target():
    got = combination_sum([1], 1)
    assert sort_combinations(got) == sort_combinations([[1]])


def test_single_candidate_repeated():
    got = combination_sum([1], 3)
    assert sort_combinations(got) == sort_combinations([[1, 1, 1]])


def test_multiple_solutions():
    got = combination_sum([2, 3, 7], 9)
    expected = [[2, 7], [2, 2, 2, 3], [3, 3, 3]]
    assert sort_combinations(got) == sort_combinations(expected)
