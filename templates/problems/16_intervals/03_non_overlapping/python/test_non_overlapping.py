from non_overlapping import erase_overlap_intervals


def test_one_removal():
    assert erase_overlap_intervals([[1, 2], [2, 3], [3, 4], [1, 3]]) == 1


def test_no_removal():
    assert erase_overlap_intervals([[1, 2], [2, 3]]) == 0


def test_all_overlap():
    assert erase_overlap_intervals([[1, 2], [1, 2], [1, 2]]) == 2


def test_single_interval():
    assert erase_overlap_intervals([[1, 5]]) == 0


def test_nested_intervals():
    assert erase_overlap_intervals([[1, 10], [2, 3], [4, 5]]) == 1


def test_chain_overlap():
    assert erase_overlap_intervals([[1, 3], [2, 4], [3, 5]]) == 1


def test_negative_start():
    assert erase_overlap_intervals([[-50000, 1]]) == 0
