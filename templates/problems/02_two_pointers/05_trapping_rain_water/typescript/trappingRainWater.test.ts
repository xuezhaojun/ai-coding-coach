import { describe, it, expect } from 'vitest';
import { trap } from './trappingRainWater';

describe('TrappingRainWater', () => {
    it('basic case', () => {
        expect(trap([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1])).toBe(6);
    });

    it('v shape', () => {
        expect(trap([4, 2, 0, 3, 2, 5])).toBe(9);
    });

    it('no water', () => {
        expect(trap([1, 2, 3, 4, 5])).toBe(0);
    });

    it('empty input', () => {
        expect(trap([])).toBe(0);
    });

    it('single bar', () => {
        expect(trap([5])).toBe(0);
    });

    it('two bars', () => {
        expect(trap([3, 1])).toBe(0);
    });

    it('flat surface', () => {
        expect(trap([3, 3, 3, 3])).toBe(0);
    });
});
