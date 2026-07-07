import { describe, it, expect } from 'vitest';
import { longestIncreasingPath } from './longestIncreasingPath';

describe('LongestIncreasingPath', () => {
    it('example 1', () => {
        expect(longestIncreasingPath([[9, 9, 4], [6, 6, 8], [2, 1, 1]])).toBe(4);
    });

    it('example 2', () => {
        expect(longestIncreasingPath([[3, 4, 5], [3, 2, 6], [2, 2, 1]])).toBe(4);
    });

    it('single cell', () => {
        expect(longestIncreasingPath([[1]])).toBe(1);
    });

    it('single row', () => {
        expect(longestIncreasingPath([[1, 2, 3, 4]])).toBe(4);
    });

    it('single column', () => {
        expect(longestIncreasingPath([[1], [2], [3]])).toBe(3);
    });

    it('all same values', () => {
        expect(longestIncreasingPath([[5, 5], [5, 5]])).toBe(1);
    });
});
