from n_queens import solve_n_queens


def sort_boards(boards: list[list[str]]) -> list[list[str]]:
    boards.sort()
    return boards


def test_n_1():
    got = solve_n_queens(1)
    expected = [["Q"]]
    assert len(got) == 1
    assert sort_boards(got) == sort_boards(expected)


def test_n_2_no_solution():
    got = solve_n_queens(2)
    assert len(got) == 0
    assert sort_boards(got) == sort_boards([])


def test_n_3_no_solution():
    got = solve_n_queens(3)
    assert len(got) == 0
    assert sort_boards(got) == sort_boards([])


def test_n_4():
    got = solve_n_queens(4)
    expected = [
        [".Q..", "...Q", "Q...", "..Q."],
        ["..Q.", "Q...", "...Q", ".Q.."],
    ]
    assert len(got) == 2
    assert sort_boards(got) == sort_boards(expected)


def test_n_5():
    got = solve_n_queens(5)
    assert len(got) == 10


def test_n_6():
    got = solve_n_queens(6)
    assert len(got) == 4


def test_n_8():
    got = solve_n_queens(8)
    assert len(got) == 92
