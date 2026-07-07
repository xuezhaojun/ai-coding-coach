from combination_sum_ii import combination_sum_2


def sort_combinations(result: list[list[int]]) -> list[list[int]]:
    for c in result:
        c.sort()
    result.sort(key=lambda g: (len(g), g))
    return result


def test_example_1():
    got = combination_sum_2([10, 1, 2, 7, 6, 1, 5], 8)
    expected = [[1, 1, 6], [1, 2, 5], [1, 7], [2, 6]]
    for c in got:
        c.sort()
    assert sort_combinations(got) == sort_combinations(expected)


def test_example_2():
    got = combination_sum_2([2, 5, 2, 1, 2], 5)
    expected = [[1, 2, 2], [5]]
    for c in got:
        c.sort()
    assert sort_combinations(got) == sort_combinations(expected)


def test_no_combination():
    got = combination_sum_2([3, 5], 1)
    assert sort_combinations(got) == sort_combinations([])


def test_single_element_matches():
    got = combination_sum_2([1], 1)
    assert sort_combinations(got) == sort_combinations([[1]])


def test_all_duplicates():
    got = combination_sum_2([2, 2, 2], 4)
    assert sort_combinations(got) == sort_combinations([[2, 2]])


def test_target_zero():
    got = combination_sum_2([1, 2, 3], 0)
    assert sort_combinations(got) == sort_combinations([[]])
