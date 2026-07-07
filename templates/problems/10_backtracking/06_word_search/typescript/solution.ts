/**
 * Search for the word in the board using backtracking/DFS.
 * Time: O(m * n * 4^L) where L = word length, Space: O(L) recursion depth
 */
export function solveExist(board: string[][], word: string): boolean {
    const rows = board.length;
    const cols = board[0].length;

    const dfs = (r: number, c: number, idx: number): boolean => {
        if (idx === word.length) return true;
        if (r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] !== word[idx]) {
            return false;
        }
        const tmp = board[r][c];
        board[r][c] = '#';
        const dirs = [[0, 1], [0, -1], [1, 0], [-1, 0]];
        for (const [dr, dc] of dirs) {
            if (dfs(r + dr, c + dc, idx + 1)) {
                board[r][c] = tmp;
                return true;
            }
        }
        board[r][c] = tmp;
        return false;
    };

    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if (dfs(r, c, 0)) return true;
        }
    }
    return false;
}
