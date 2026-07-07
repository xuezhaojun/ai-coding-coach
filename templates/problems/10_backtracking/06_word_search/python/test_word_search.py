from word_search import exist


def test_word_exists():
    board = [
        ["A", "B", "C", "E"],
        ["S", "F", "C", "S"],
        ["A", "D", "E", "E"],
    ]
    assert exist(board, "ABCCED") == True


def test_word_exists_path_see():
    board = [
        ["A", "B", "C", "E"],
        ["S", "F", "C", "S"],
        ["A", "D", "E", "E"],
    ]
    assert exist(board, "SEE") == True


def test_word_does_not_exist():
    board = [
        ["A", "B", "C", "E"],
        ["S", "F", "C", "S"],
        ["A", "D", "E", "E"],
    ]
    assert exist(board, "ABCB") == False


def test_single_cell_match():
    assert exist([["A"]], "A") == True


def test_single_cell_no_match():
    assert exist([["A"]], "B") == False


def test_word_longer_than_board_cells():
    board = [
        ["A", "B"],
        ["C", "D"],
    ]
    assert exist(board, "ABCDA") == False


def test_snake_path():
    board = [
        ["A", "B", "C"],
        ["D", "E", "F"],
        ["G", "H", "I"],
    ]
    assert exist(board, "ABCFEDGHI") == True
