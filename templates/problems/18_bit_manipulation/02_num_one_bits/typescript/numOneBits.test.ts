import { describe, it, expect } from 'vitest';
import { hammingWeight } from './numOneBits';

describe('NumOneBits', () => {
    it('three ones', () => {
        expect(hammingWeight(0b00000000000000000000000000001011)).toBe(3);
    });

    it('one bit', () => {
        expect(hammingWeight(0b00000000000000000000000010000000)).toBe(1);
    });

    it('all ones 32bit', () => {
        expect(hammingWeight(0b11111111111111111111111111111101)).toBe(31);
    });

    it('zero', () => {
        expect(hammingWeight(0)).toBe(0);
    });

    it('power of two', () => {
        expect(hammingWeight(16)).toBe(1);
    });

    it('all ones', () => {
        expect(hammingWeight(0xFFFFFFFF)).toBe(32);
    });
});
