from three_sum import three_sum


def sort_result(result: list[list[int]]) -> list[list[int]]:
    for triplet in result:
        triplet.sort()
    result.sort()
    return result


def test_basic_case():
    result = three_sum([-1, 0, 1, 2, -1, -4])
    assert sort_result(result) == sort_result([[-1, -1, 2], [-1, 0, 1]])


def test_no_triplets():
    result = three_sum([0, 1, 1])
    assert result == []


def test_all_zeros():
    result = three_sum([0, 0, 0])
    assert sort_result(result) == sort_result([[0, 0, 0]])


def test_empty_input():
    result = three_sum([])
    assert result == []


def test_two_elements_only():
    result = three_sum([-1, 1])
    assert result == []


def test_multiple_triplets():
    result = three_sum([-2, 0, 1, 1, 2])
    assert sort_result(result) == sort_result([[-2, 0, 2], [-2, 1, 1]])


def test_all_positive():
    result = three_sum([1, 2, 3, 4, 5])
    assert result == []


def test_duplicate_left_right_values():
    result = three_sum([-2, 0, 0, 2, 2])
    assert sort_result(result) == sort_result([[-2, 0, 2]])


def test_all_same_value_after_anchor():
    result = three_sum([-4, 2, 2, 2, 2])
    assert sort_result(result) == sort_result([[-4, 2, 2]])
