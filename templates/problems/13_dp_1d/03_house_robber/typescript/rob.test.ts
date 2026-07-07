import { describe, it, expect } from 'vitest';
import { rob } from './rob';

describe('Rob', () => {
    it('example 1', () => {
        expect(rob([1, 2, 3, 1])).toBe(4);
    });

    it('example 2', () => {
        expect(rob([2, 7, 9, 3, 1])).toBe(12);
    });

    it('single house', () => {
        expect(rob([5])).toBe(5);
    });

    it('two houses pick larger', () => {
        expect(rob([1, 2])).toBe(2);
    });

    it('all same values', () => {
        expect(rob([3, 3, 3, 3])).toBe(6);
    });

    it('empty', () => {
        expect(rob([])).toBe(0);
    });

    it('large values alternating', () => {
        expect(rob([100, 1, 100, 1, 100])).toBe(300);
    });
});
