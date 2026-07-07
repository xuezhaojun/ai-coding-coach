import { describe, it, expect } from 'vitest';
import { robII } from './robII';

describe('RobII', () => {
    it('example 1', () => {
        expect(robII([2, 3, 2])).toBe(3);
    });

    it('example 2', () => {
        expect(robII([1, 2, 3, 1])).toBe(4);
    });

    it('single house', () => {
        expect(robII([5])).toBe(5);
    });

    it('two houses', () => {
        expect(robII([1, 2])).toBe(2);
    });

    it('example 3', () => {
        expect(robII([1, 2, 3])).toBe(3);
    });

    it('four houses', () => {
        expect(robII([1, 3, 1, 3, 100])).toBe(103);
    });

    it('all same', () => {
        expect(robII([4, 4, 4, 4, 4])).toBe(8);
    });
});
