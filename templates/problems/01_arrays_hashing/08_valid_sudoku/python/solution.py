def solve_is_valid_sudoku(board: list[list[str]]) -> bool:
    """Validate a 9x9 sudoku board using hash sets for rows, columns, and boxes.

    Time: O(1) - fixed 9x9 board, Space: O(1)
    """
    rows = [[False] * 9 for _ in range(9)]
    cols = [[False] * 9 for _ in range(9)]
    boxes = [[False] * 9 for _ in range(9)]

    for r in range(9):
        for c in range(9):
            if board[r][c] == ".":
                continue
            d = ord(board[r][c]) - ord("1")
            box = (r // 3) * 3 + c // 3

            if rows[r][d] or cols[c][d] or boxes[box][d]:
                return False
            rows[r][d] = True
            cols[c][d] = True
            boxes[box][d] = True
    return True
