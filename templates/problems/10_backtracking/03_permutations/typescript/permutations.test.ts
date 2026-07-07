import { describe, it, expect } from 'vitest';
import { permute } from './permutations';

function sortPermutations(result: number[][]): number[][] {
    result.sort((a, b) => {
        for (let k = 0; k < a.length && k < b.length; k++) {
            if (a[k] !== b[k]) return a[k] - b[k];
        }
        return a.length - b.length;
    });
    return result;
}

describe('Permutations', () => {
    it('three elements', () => {
        const got = permute([1, 2, 3]);
        const expected = [
            [1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1],
        ];
        expect(sortPermutations(got)).toEqual(sortPermutations(expected));
    });

    it('single element', () => {
        const got = permute([1]);
        const expected = [[1]];
        expect(sortPermutations(got)).toEqual(sortPermutations(expected));
    });

    it('two elements', () => {
        const got = permute([0, 1]);
        const expected = [[0, 1], [1, 0]];
        expect(sortPermutations(got)).toEqual(sortPermutations(expected));
    });

    it('negative numbers', () => {
        const got = permute([-1, 0, 1]);
        const expected = [
            [-1, 0, 1], [-1, 1, 0], [0, -1, 1], [0, 1, -1], [1, -1, 0], [1, 0, -1],
        ];
        expect(sortPermutations(got)).toEqual(sortPermutations(expected));
    });

    it('four elements count', () => {
        const got = permute([1, 2, 3, 4]);
        expect(got.length).toBe(24);
    });
});
