import { describe, it, expect } from 'vitest';
import { findTargetSumWays } from './findTargetSumWays';

describe('FindTargetSumWays', () => {
    it('example 1', () => {
        expect(findTargetSumWays([1, 1, 1, 1, 1], 3)).toBe(5);
    });

    it('single element match', () => {
        expect(findTargetSumWays([1], 1)).toBe(1);
    });

    it('single element negative', () => {
        expect(findTargetSumWays([1], -1)).toBe(1);
    });

    it('impossible target', () => {
        expect(findTargetSumWays([1], 2)).toBe(0);
    });

    it('two elements', () => {
        expect(findTargetSumWays([1, 2], 1)).toBe(1);
    });

    it('all zeros', () => {
        expect(findTargetSumWays([0, 0, 0], 0)).toBe(8);
    });
});
