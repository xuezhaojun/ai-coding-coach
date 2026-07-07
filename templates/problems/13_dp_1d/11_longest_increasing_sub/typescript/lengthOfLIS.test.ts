import { describe, it, expect } from 'vitest';
import { lengthOfLIS } from './lengthOfLIS';

describe('LengthOfLIS', () => {
    it('example 1', () => {
        expect(lengthOfLIS([10, 9, 2, 5, 3, 7, 101, 18])).toBe(4);
    });

    it('all increasing', () => {
        expect(lengthOfLIS([1, 2, 3, 4, 5])).toBe(5);
    });

    it('all decreasing', () => {
        expect(lengthOfLIS([5, 4, 3, 2, 1])).toBe(1);
    });

    it('single element', () => {
        expect(lengthOfLIS([7])).toBe(1);
    });

    it('example 2', () => {
        expect(lengthOfLIS([0, 1, 0, 3, 2, 3])).toBe(4);
    });

    it('duplicates', () => {
        expect(lengthOfLIS([7, 7, 7, 7, 7])).toBe(1);
    });

    it('valley then rise', () => {
        expect(lengthOfLIS([1, 3, 6, 7, 9, 4, 10, 5, 6])).toBe(6);
    });
});
