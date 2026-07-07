import { describe, it, expect } from 'vitest';
import { longestConsecutive } from './longestConsecutive';

describe('LongestConsecutive', () => {
    it('basic case', () => {
        expect(longestConsecutive([100, 4, 200, 1, 3, 2])).toBe(4);
    });

    it('longer sequence', () => {
        expect(longestConsecutive([0, 3, 7, 2, 5, 8, 4, 6, 0, 1])).toBe(9);
    });

    it('empty array', () => {
        expect(longestConsecutive([])).toBe(0);
    });

    it('single element', () => {
        expect(longestConsecutive([1])).toBe(1);
    });

    it('no consecutive', () => {
        expect(longestConsecutive([10, 20, 30])).toBe(1);
    });

    it('duplicates in sequence', () => {
        expect(longestConsecutive([1, 2, 2, 3, 3, 4])).toBe(4);
    });

    it('negative numbers', () => {
        expect(longestConsecutive([-3, -2, -1, 0, 1])).toBe(5);
    });

    it('all same elements', () => {
        expect(longestConsecutive([5, 5, 5, 5])).toBe(1);
    });
});
