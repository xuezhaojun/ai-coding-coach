import { describe, it, expect } from 'vitest';
import { spiralOrder } from './spiralMatrix';

describe('SpiralMatrix', () => {
    it('3x3', () => {
        expect(spiralOrder([[1, 2, 3], [4, 5, 6], [7, 8, 9]])).toEqual([1, 2, 3, 6, 9, 8, 7, 4, 5]);
    });

    it('3x4', () => {
        expect(spiralOrder([[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12]])).toEqual([1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7]);
    });

    it('1x1', () => {
        expect(spiralOrder([[1]])).toEqual([1]);
    });

    it('1 row', () => {
        expect(spiralOrder([[1, 2, 3]])).toEqual([1, 2, 3]);
    });

    it('1 column', () => {
        expect(spiralOrder([[1], [2], [3]])).toEqual([1, 2, 3]);
    });

    it('2x2', () => {
        expect(spiralOrder([[1, 2], [3, 4]])).toEqual([1, 2, 4, 3]);
    });
});
