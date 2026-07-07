import { describe, it, expect } from 'vitest';
import { isInterleave } from './isInterleave';

describe('IsInterleave', () => {
    it('example true', () => {
        expect(isInterleave('aabcc', 'dbbca', 'aadbbcbcac')).toBe(true);
    });

    it('example false', () => {
        expect(isInterleave('aabcc', 'dbbca', 'aadbbbaccc')).toBe(false);
    });

    it('both empty', () => {
        expect(isInterleave('', '', '')).toBe(true);
    });

    it('s1 empty', () => {
        expect(isInterleave('', 'abc', 'abc')).toBe(true);
    });

    it('s2 empty', () => {
        expect(isInterleave('abc', '', 'abc')).toBe(true);
    });

    it('length mismatch', () => {
        expect(isInterleave('a', 'b', 'abc')).toBe(false);
    });

    it('single chars true', () => {
        expect(isInterleave('a', 'b', 'ab')).toBe(true);
    });
});
