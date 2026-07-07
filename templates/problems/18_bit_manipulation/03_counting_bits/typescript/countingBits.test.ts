import { describe, it, expect } from 'vitest';
import { countBits } from './countingBits';

describe('CountingBits', () => {
    it('n=2', () => {
        expect(countBits(2)).toEqual([0, 1, 1]);
    });

    it('n=5', () => {
        expect(countBits(5)).toEqual([0, 1, 1, 2, 1, 2]);
    });

    it('n=0', () => {
        expect(countBits(0)).toEqual([0]);
    });

    it('n=1', () => {
        expect(countBits(1)).toEqual([0, 1]);
    });

    it('n=8', () => {
        expect(countBits(8)).toEqual([0, 1, 1, 2, 1, 2, 2, 3, 1]);
    });

    it('n=3', () => {
        expect(countBits(3)).toEqual([0, 1, 1, 2]);
    });
});
