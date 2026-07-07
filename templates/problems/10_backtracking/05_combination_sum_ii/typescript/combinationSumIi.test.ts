import { describe, it, expect } from 'vitest';
import { combinationSum2 } from './combinationSumIi';

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

describe('CombinationSum2', () => {
    it('example 1', () => {
        const got = combinationSum2([10, 1, 2, 7, 6, 1, 5], 8);
        for (const c of got) c.sort((a, b) => a - b);
        const expected = [[1, 1, 6], [1, 2, 5], [1, 7], [2, 6]];
        expect(sortCombinations(got)).toEqual(sortCombinations(expected));
    });

    it('example 2', () => {
        const got = combinationSum2([2, 5, 2, 1, 2], 5);
        for (const c of got) c.sort((a, b) => a - b);
        const expected = [[1, 2, 2], [5]];
        expect(sortCombinations(got)).toEqual(sortCombinations(expected));
    });

    it('no combination', () => {
        const got = combinationSum2([3, 5], 1);
        expect(sortCombinations(got)).toEqual(sortCombinations([]));
    });

    it('single element matches', () => {
        const got = combinationSum2([1], 1);
        expect(sortCombinations(got)).toEqual(sortCombinations([[1]]));
    });

    it('all duplicates', () => {
        const got = combinationSum2([2, 2, 2], 4);
        expect(sortCombinations(got)).toEqual(sortCombinations([[2, 2]]));
    });

    it('target zero', () => {
        const got = combinationSum2([1, 2, 3], 0);
        expect(sortCombinations(got)).toEqual(sortCombinations([[]]));
    });
});
