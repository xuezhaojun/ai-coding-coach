import { describe, it, expect } from 'vitest';
import { reverseBits } from './reverseBits';

describe('ReverseBits', () => {
    it('example 1', () => {
        expect(reverseBits(0b00000010100101000001111010011100)).toBe(964176192);
    });

    it('example 2', () => {
        expect(reverseBits(0b11111111111111111111111111111101)).toBe(3221225471);
    });

    it('zero', () => {
        expect(reverseBits(0)).toBe(0);
    });

    it('all ones', () => {
        expect(reverseBits(0xFFFFFFFF)).toBe(0xFFFFFFFF);
    });

    it('one', () => {
        expect(reverseBits(1)).toBe(0x80000000);
    });

    it('power of two', () => {
        expect(reverseBits(0x80000000)).toBe(1);
    });
});
