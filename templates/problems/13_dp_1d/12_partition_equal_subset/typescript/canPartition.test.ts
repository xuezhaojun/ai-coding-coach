import { describe, it, expect } from 'vitest';
import { canPartition } from './canPartition';

describe('CanPartition', () => {
    it('example can partition', () => {
        expect(canPartition([1, 5, 11, 5])).toBe(true);
    });

    it('example cannot partition', () => {
        expect(canPartition([1, 2, 3, 5])).toBe(false);
    });

    it('two equal elements', () => {
        expect(canPartition([1, 1])).toBe(true);
    });

    it('single element', () => {
        expect(canPartition([1])).toBe(false);
    });

    it('odd total sum', () => {
        expect(canPartition([1, 2, 4])).toBe(false);
    });

    it('larger example', () => {
        expect(canPartition([1, 2, 3, 4, 5, 6, 7])).toBe(true);
    });

    it('all zeros', () => {
        expect(canPartition([0, 0, 0, 0])).toBe(true);
    });
});
