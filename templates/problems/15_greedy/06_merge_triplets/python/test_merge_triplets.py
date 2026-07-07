from merge_triplets import merge_triplets


def test_example_true():
    assert merge_triplets([[2, 5, 3], [1, 8, 4], [1, 7, 5]], [2, 7, 5]) == True


def test_example_false():
    assert merge_triplets([[3, 4, 5], [4, 5, 6]], [3, 2, 5]) == False


def test_exact_match_single():
    assert merge_triplets([[2, 5, 3]], [2, 5, 3]) == True


def test_no_valid_triplet():
    assert merge_triplets([[1, 1, 1]], [2, 2, 2]) == False


def test_filter_out_exceeding():
    assert merge_triplets([[2, 5, 3], [10, 1, 1], [1, 7, 5]], [2, 7, 5]) == True


def test_all_exceed_one_dim():
    assert merge_triplets([[3, 1, 1], [3, 2, 2]], [2, 2, 2]) == False


def test_multiple_combos():
    assert merge_triplets([[1, 2, 3], [2, 1, 3], [2, 2, 1]], [2, 2, 3]) == True
