import { describe, it, expect } from 'vitest';
import { subsets } from './subsets';

function sortSubsets(result: number[][]): number[][] {
    for (const s of result) {
        s.sort((a, b) => a - b);
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

describe('Subsets', () => {
    it('three elements', () => {
        const got = subsets([1, 2, 3]);
        const expected = [[], [1], [2], [3], [1, 2], [1, 3], [2, 3], [1, 2, 3]];
        expect(sortSubsets(got)).toEqual(sortSubsets(expected));
    });

    it('single element', () => {
        const got = subsets([0]);
        const expected = [[], [0]];
        expect(sortSubsets(got)).toEqual(sortSubsets(expected));
    });

    it('two elements', () => {
        const got = subsets([1, 2]);
        const expected = [[], [1], [2], [1, 2]];
        expect(sortSubsets(got)).toEqual(sortSubsets(expected));
    });

    it('empty input', () => {
        const got = subsets([]);
        const expected = [[]];
        expect(sortSubsets(got)).toEqual(sortSubsets(expected));
    });

    it('negative numbers', () => {
        const got = subsets([-1, 0]);
        const expected = [[], [-1], [0], [-1, 0]];
        expect(sortSubsets(got)).toEqual(sortSubsets(expected));
    });

    it('five elements - slice sharing regression', () => {
        const got = subsets([9, 0, 3, 5, 7]);
        const expected = [
            [], [9], [0], [3], [5], [7],
            [9, 0], [9, 3], [9, 5], [9, 7], [0, 3], [0, 5], [0, 7], [3, 5], [3, 7], [5, 7],
            [9, 0, 3], [9, 0, 5], [9, 0, 7], [9, 3, 5], [9, 3, 7], [9, 5, 7], [0, 3, 5], [0, 3, 7], [0, 5, 7], [3, 5, 7],
            [9, 0, 3, 5], [9, 0, 3, 7], [9, 0, 5, 7], [9, 3, 5, 7], [0, 3, 5, 7],
            [9, 0, 3, 5, 7],
        ];
        expect(sortSubsets(got)).toEqual(sortSubsets(expected));
    });

    it('four elements', () => {
        const got = subsets([1, 2, 3, 4]);
        const expected = [
            [], [1], [2], [3], [4],
            [1, 2], [1, 3], [1, 4], [2, 3], [2, 4], [3, 4],
            [1, 2, 3], [1, 2, 4], [1, 3, 4], [2, 3, 4],
            [1, 2, 3, 4],
        ];
        expect(sortSubsets(got)).toEqual(sortSubsets(expected));
    });
});
