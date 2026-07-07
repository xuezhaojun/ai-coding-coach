import { describe, it, expect } from 'vitest';
import { combinationSum } from './combinationSum';

function sortCombinations(result: number[][]): number[][] {
    for (const c of result) {
        c.sort((a, b) => a - b);
    }
    result.sort((a, b) => {
        if (a.length !== b.length) return a.length - b.length;
        for (let k = 0; k < a.length; k++) {
            if (a[k] !== b[k]) return a[k] - b[k];
        }
        return 0;
    });
    return result;
}

describe('CombinationSum', () => {
    it('example 1', () => {
        const got = combinationSum([2, 3, 6, 7], 7);
        const expected = [[2, 2, 3], [7]];
        expect(sortCombinations(got)).toEqual(sortCombinations(expected));
    });

    it('example 2', () => {
        const got = combinationSum([2, 3, 5], 8);
        const expected = [[2, 2, 2, 2], [2, 3, 3], [3, 5]];
        expect(sortCombinations(got)).toEqual(sortCombinations(expected));
    });

    it('no combination', () => {
        const got = combinationSum([2], 1);
        expect(sortCombinations(got)).toEqual(sortCombinations([]));
    });

    it('single candidate equals target', () => {
        const got = combinationSum([1], 1);
        expect(sortCombinations(got)).toEqual(sortCombinations([[1]]));
    });

    it('single candidate repeated', () => {
        const got = combinationSum([1], 3);
        expect(sortCombinations(got)).toEqual(sortCombinations([[1, 1, 1]]));
    });

    it('multiple solutions', () => {
        const got = combinationSum([2, 3, 7], 9);
        const expected = [[2, 7], [2, 2, 2, 3], [3, 3, 3]];
        expect(sortCombinations(got)).toEqual(sortCombinations(expected));
    });
});
