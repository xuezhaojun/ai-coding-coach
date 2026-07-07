/**
 * Capture surrounded regions by marking border-connected O's.
 * Time: O(m * n), Space: O(m * n) recursion stack.
 */
export function solveSolve(board: string[][]): void {
    if (board.length === 0) return;
    const rows = board.length;
    const cols = board[0].length;

    const dfs = (r: number, c: number): void => {
        if (r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] !== "O") {
            return;
        }
        board[r][c] = "T";
        dfs(r + 1, c);
        dfs(r - 1, c);
        dfs(r, c + 1);
        dfs(r, c - 1);
    };

    for (let r = 0; r < rows; r++) {
        dfs(r, 0);
        dfs(r, cols - 1);
    }
    for (let c = 0; c < cols; c++) {
        dfs(0, c);
        dfs(rows - 1, c);
    }

    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if (board[r][c] === "O") {
                board[r][c] = "X";
            } else if (board[r][c] === "T") {
                board[r][c] = "O";
            }
        }
    }
}
