import { describe, it, expect } from 'vitest';
import { numDistinct } from './numDistinct';

describe('NumDistinct', () => {
    it('example rabbbit', () => {
        expect(numDistinct('rabbbit', 'rabbit')).toBe(3);
    });

    it('example babgbag', () => {
        expect(numDistinct('babgbag', 'bag')).toBe(5);
    });

    it('no match', () => {
        expect(numDistinct('abc', 'def')).toBe(0);
    });

    it('t longer than s', () => {
        expect(numDistinct('ab', 'abc')).toBe(0);
    });

    it('equal strings', () => {
        expect(numDistinct('abc', 'abc')).toBe(1);
    });

    it('empty t', () => {
        expect(numDistinct('abc', '')).toBe(1);
    });

    it('single char repeated', () => {
        expect(numDistinct('aaa', 'a')).toBe(3);
    });
});
