from permutations import permute


def sort_permutations(result: list[list[int]]) -> list[list[int]]:
    result.sort(key=lambda g: (g, len(g)))
    return result


def test_three_elements():
    got = permute([1, 2, 3])
    expected = [
        [1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1],
    ]
    assert sort_permutations(got) == sort_permutations(expected)


def test_single_element():
    got = permute([1])
    expected = [[1]]
    assert sort_permutations(got) == sort_permutations(expected)


def test_two_elements():
    got = permute([0, 1])
    expected = [[0, 1], [1, 0]]
    assert sort_permutations(got) == sort_permutations(expected)


def test_negative_numbers():
    got = permute([-1, 0, 1])
    expected = [
        [-1, 0, 1], [-1, 1, 0], [0, -1, 1], [0, 1, -1], [1, -1, 0], [1, 0, -1],
    ]
    assert sort_permutations(got) == sort_permutations(expected)


def test_four_elements_count():
    got = permute([1, 2, 3, 4])
    assert len(got) == 24
