import { describe, it, expect } from 'vitest';
import { solveNQueens } from './nQueens';

function sortBoards(boards: string[][]): string[][] {
    boards.sort((a, b) => {
        for (let k = 0; k < a.length; k++) {
            if (a[k] !== b[k]) return a[k] < b[k] ? -1 : 1;
        }
        return 0;
    });
    return boards;
}

describe('NQueens', () => {
    it('n=1', () => {
        const got = solveNQueens(1);
        const expected = [['Q']];
        expect(got.length).toBe(1);
        expect(sortBoards(got)).toEqual(sortBoards(expected));
    });

    it('n=2 no solution', () => {
        const got = solveNQueens(2);
        expect(got.length).toBe(0);
        expect(sortBoards(got)).toEqual(sortBoards([]));
    });

    it('n=3 no solution', () => {
        const got = solveNQueens(3);
        expect(got.length).toBe(0);
        expect(sortBoards(got)).toEqual(sortBoards([]));
    });

    it('n=4', () => {
        const got = solveNQueens(4);
        const expected = [
            ['.Q..', '...Q', 'Q...', '..Q.'],
            ['..Q.', 'Q...', '...Q', '.Q..'],
        ];
        expect(got.length).toBe(2);
        expect(sortBoards(got)).toEqual(sortBoards(expected));
    });

    it('n=5', () => {
        const got = solveNQueens(5);
        expect(got.length).toBe(10);
    });

    it('n=6', () => {
        const got = solveNQueens(6);
        expect(got.length).toBe(4);
    });

    it('n=8', () => {
        const got = solveNQueens(8);
        expect(got.length).toBe(92);
    });
});
