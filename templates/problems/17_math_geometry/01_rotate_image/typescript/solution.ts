/**
 * Rotate the matrix 90 degrees clockwise in-place. O(n^2) time, O(1) space.
 */
export function solveRotate(matrix: number[][]): void {
    const n = matrix.length;
    // Transpose
    for (let i = 0; i < n; i++) {
        for (let j = i + 1; j < n; j++) {
            [matrix[i][j], matrix[j][i]] = [matrix[j][i], matrix[i][j]];
        }
    }
    // Reverse each row
    for (let i = 0; i < n; i++) {
        let l = 0, r = n - 1;
        while (l < r) {
            [matrix[i][l], matrix[i][r]] = [matrix[i][r], matrix[i][l]];
            l++;
            r--;
        }
    }
}
