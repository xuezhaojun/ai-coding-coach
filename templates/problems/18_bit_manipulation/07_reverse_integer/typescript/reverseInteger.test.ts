import { describe, it, expect } from 'vitest';
import { reverse } from './reverseInteger';

describe('ReverseInteger', () => {
    it('positive', () => {
        expect(reverse(123)).toBe(321);
    });

    it('negative', () => {
        expect(reverse(-123)).toBe(-321);
    });

    it('trailing zero', () => {
        expect(reverse(120)).toBe(21);
    });

    it('zero', () => {
        expect(reverse(0)).toBe(0);
    });

    it('single digit', () => {
        expect(reverse(5)).toBe(5);
    });

    it('overflow positive', () => {
        expect(reverse(1534236469)).toBe(0);
    });

    it('overflow negative', () => {
        expect(reverse(-2147483648)).toBe(0);
    });
});
