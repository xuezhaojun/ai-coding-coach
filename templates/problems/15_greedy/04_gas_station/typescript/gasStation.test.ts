import { describe, it, expect } from 'vitest';
import { canCompleteCircuit } from './gasStation';

describe('GasStation', () => {
    it('example 1', () => {
        expect(canCompleteCircuit([1, 2, 3, 4, 5], [3, 4, 5, 1, 2])).toBe(3);
    });

    it('no solution', () => {
        expect(canCompleteCircuit([2, 3, 4], [3, 4, 3])).toBe(-1);
    });

    it('single station enough', () => {
        expect(canCompleteCircuit([5], [3])).toBe(0);
    });

    it('single station not enough', () => {
        expect(canCompleteCircuit([2], [4])).toBe(-1);
    });

    it('start at index 0', () => {
        expect(canCompleteCircuit([3, 1, 1], [1, 2, 2])).toBe(0);
    });

    it('all equal', () => {
        expect(canCompleteCircuit([1, 1, 1], [1, 1, 1])).toBe(0);
    });

    it('start at last', () => {
        expect(canCompleteCircuit([1, 1, 5], [2, 3, 1])).toBe(2);
    });
});
