def solve_solve_n_queens(n: int) -> list[list[str]]:
    """Solve the N-Queens problem using backtracking.

    Time: O(n!), Space: O(n) for column and diagonal tracking
    """
    result: list[list[str]] = []
    cols = [False] * n
    diag1: dict[int, bool] = {}  # row - col
    diag2: dict[int, bool] = {}  # row + col
    board = [["." for _ in range(n)] for _ in range(n)]

    def backtrack(row: int) -> None:
        if row == n:
            result.append(["".join(r) for r in board])
            return
        for col in range(n):
            if cols[col] or diag1.get(row - col, False) or diag2.get(row + col, False):
                continue
            board[row][col] = "Q"
            cols[col] = True
            diag1[row - col] = True
            diag2[row + col] = True
            backtrack(row + 1)
            board[row][col] = "."
            cols[col] = False
            diag1[row - col] = False
            diag2[row + col] = False

    backtrack(0)
    return result
