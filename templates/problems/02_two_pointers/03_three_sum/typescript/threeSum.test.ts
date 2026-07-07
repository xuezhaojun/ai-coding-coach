import { describe, it, expect } from 'vitest';
import { threeSum } from './threeSum';

function sortResult(result: number[][]): number[][] {
    for (const triplet of result) {
        triplet.sort((a, b) => a - b);
    }
    result.sort((a, b) => {
        for (let k = 0; k < 3; k++) {
            if (a[k] !== b[k]) return a[k] - b[k];
        }
        return 0;
    });
    return result;
}

describe('ThreeSum', () => {
    it('basic case', () => {
        const result = sortResult(threeSum([-1, 0, 1, 2, -1, -4]));
        const expected = sortResult([[-1, -1, 2], [-1, 0, 1]]);
        expect(result).toEqual(expected);
    });

    it('no triplets', () => {
        expect(threeSum([0, 1, 1])).toEqual([]);
    });

    it('all zeros', () => {
        const result = sortResult(threeSum([0, 0, 0]));
        const expected = sortResult([[0, 0, 0]]);
        expect(result).toEqual(expected);
    });

    it('empty input', () => {
        expect(threeSum([])).toEqual([]);
    });

    it('two elements only', () => {
        expect(threeSum([-1, 1])).toEqual([]);
    });

    it('multiple triplets', () => {
        const result = sortResult(threeSum([-2, 0, 1, 1, 2]));
        const expected = sortResult([[-2, 0, 2], [-2, 1, 1]]);
        expect(result).toEqual(expected);
    });

    it('all positive', () => {
        expect(threeSum([1, 2, 3, 4, 5])).toEqual([]);
    });

    it('duplicate left right values', () => {
        const result = sortResult(threeSum([-2, 0, 0, 2, 2]));
        const expected = sortResult([[-2, 0, 2]]);
        expect(result).toEqual(expected);
    });

    it('all same value after anchor', () => {
        const result = sortResult(threeSum([-4, 2, 2, 2, 2]));
        const expected = sortResult([[-4, 2, 2]]);
        expect(result).toEqual(expected);
    });
});
