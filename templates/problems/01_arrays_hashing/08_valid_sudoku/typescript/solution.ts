/**
 * Validate a 9x9 sudoku board using hash sets for rows, columns, and boxes.
 * Time: O(1) - fixed 9x9 board, Space: O(1)
 */
export function solveIsValidSudoku(board: string[][]): boolean {
    const rows: boolean[][] = Array.from({ length: 9 }, () => new Array(9).fill(false));
    const cols: boolean[][] = Array.from({ length: 9 }, () => new Array(9).fill(false));
    const boxes: boolean[][] = Array.from({ length: 9 }, () => new Array(9).fill(false));

    for (let r = 0; r < 9; r++) {
        for (let c = 0; c < 9; c++) {
            if (board[r][c] === '.') {
                continue;
            }
            const d = board[r][c].charCodeAt(0) - '1'.charCodeAt(0);
            const box = Math.floor(r / 3) * 3 + Math.floor(c / 3);

            if (rows[r][d] || cols[c][d] || boxes[box][d]) {
                return false;
            }
            rows[r][d] = true;
            cols[c][d] = true;
            boxes[box][d] = true;
        }
    }
    return true;
}
