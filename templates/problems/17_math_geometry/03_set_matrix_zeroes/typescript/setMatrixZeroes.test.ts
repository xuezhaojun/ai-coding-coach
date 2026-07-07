import { describe, it, expect } from 'vitest';
import { setZeroes } from './setMatrixZeroes';

describe('SetMatrixZeroes', () => {
    it('example 1', () => {
        const matrix = [[1, 1, 1], [1, 0, 1], [1, 1, 1]];
        setZeroes(matrix);
        expect(matrix).toEqual([[1, 0, 1], [0, 0, 0], [1, 0, 1]]);
    });

    it('example 2', () => {
        const matrix = [[0, 1, 2, 0], [3, 4, 5, 2], [1, 3, 1, 5]];
        setZeroes(matrix);
        expect(matrix).toEqual([[0, 0, 0, 0], [0, 4, 5, 0], [0, 3, 1, 0]]);
    });

    it('no zeros', () => {
        const matrix = [[1, 2], [3, 4]];
        setZeroes(matrix);
        expect(matrix).toEqual([[1, 2], [3, 4]]);
    });

    it('all zeros', () => {
        const matrix = [[0, 0], [0, 0]];
        setZeroes(matrix);
        expect(matrix).toEqual([[0, 0], [0, 0]]);
    });

    it('single element zero', () => {
        const matrix = [[0]];
        setZeroes(matrix);
        expect(matrix).toEqual([[0]]);
    });

    it('single element nonzero', () => {
        const matrix = [[5]];
        setZeroes(matrix);
        expect(matrix).toEqual([[5]]);
    });

    it('corner zero', () => {
        const matrix = [[0, 1], [1, 1]];
        setZeroes(matrix);
        expect(matrix).toEqual([[0, 0], [0, 1]]);
    });
});
