/**
 * Solve the N-Queens problem using backtracking.
 * Time: O(n!), Space: O(n) for column and diagonal tracking
 */
export function solveSolveNQueens(n: number): string[][] {
    const result: string[][] = [];
    const cols = new Array<boolean>(n).fill(false);
    const diag1 = new Map<number, boolean>(); // row - col
    const diag2 = new Map<number, boolean>(); // row + col
    const board: string[][] = Array.from({ length: n }, () =>
        new Array<string>(n).fill('.')
    );

    const backtrack = (row: number): void => {
        if (row === n) {
            result.push(board.map((r) => r.join('')));
            return;
        }
        for (let col = 0; col < n; col++) {
            if (cols[col] || diag1.get(row - col) || diag2.get(row + col)) {
                continue;
            }
            board[row][col] = 'Q';
            cols[col] = true;
            diag1.set(row - col, true);
            diag2.set(row + col, true);
            backtrack(row + 1);
            board[row][col] = '.';
            cols[col] = false;
            diag1.set(row - col, false);
            diag2.set(row + col, false);
        }
    };

    backtrack(0);
    return result;
}
