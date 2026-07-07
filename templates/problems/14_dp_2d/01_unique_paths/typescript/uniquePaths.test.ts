import { describe, it, expect } from 'vitest';
import { uniquePaths } from './uniquePaths';

describe('UniquePaths', () => {
    it('3x7 grid', () => {
        expect(uniquePaths(3, 7)).toBe(28);
    });

    it('3x2 grid', () => {
        expect(uniquePaths(3, 2)).toBe(3);
    });

    it('1x1 grid', () => {
        expect(uniquePaths(1, 1)).toBe(1);
    });

    it('1xN single row', () => {
        expect(uniquePaths(1, 5)).toBe(1);
    });

    it('Nx1 single column', () => {
        expect(uniquePaths(5, 1)).toBe(1);
    });

    it('2x2 grid', () => {
        expect(uniquePaths(2, 2)).toBe(2);
    });

    it('3x3 grid', () => {
        expect(uniquePaths(3, 3)).toBe(6);
    });
});
