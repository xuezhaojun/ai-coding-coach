import { describe, it, expect } from 'vitest';
import { topKFrequent } from './topKFrequent';

describe('TopKFrequent', () => {
    it('basic case', () => {
        const result = topKFrequent([1, 1, 1, 2, 2, 3], 2).sort((a, b) => a - b);
        expect(result).toEqual([1, 2]);
    });

    it('single element', () => {
        const result = topKFrequent([1], 1).sort((a, b) => a - b);
        expect(result).toEqual([1]);
    });

    it('all same frequency k equals length', () => {
        const result = topKFrequent([1, 2, 3], 3).sort((a, b) => a - b);
        expect(result).toEqual([1, 2, 3]);
    });

    it('negative numbers', () => {
        const result = topKFrequent([-1, -1, -2, -2, -2, -3], 1).sort((a, b) => a - b);
        expect(result).toEqual([-2]);
    });

    it('k equals 1 with clear winner', () => {
        const result = topKFrequent([4, 4, 4, 1, 2, 3], 1).sort((a, b) => a - b);
        expect(result).toEqual([4]);
    });

    it('two elements', () => {
        const result = topKFrequent([1, 2], 2).sort((a, b) => a - b);
        expect(result).toEqual([1, 2]);
    });
});
