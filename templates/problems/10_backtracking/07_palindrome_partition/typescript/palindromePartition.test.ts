import { describe, it, expect } from 'vitest';
import { partition } from './palindromePartition';

function sortStringSlices(result: string[][]): string[][] {
    for (const s of result) {
        s.sort();
    }
    result.sort((a, b) => {
        if (a.length !== b.length) return a.length - b.length;
        for (let k = 0; k < a.length; k++) {
            if (a[k] !== b[k]) return a[k] < b[k] ? -1 : 1;
        }
        return 0;
    });
    return result;
}

describe('PalindromePartition', () => {
    it('example aab', () => {
        const got = partition('aab');
        const expected = [['a', 'a', 'b'], ['aa', 'b']];
        expect(sortStringSlices(got)).toEqual(sortStringSlices(expected));
    });

    it('single char', () => {
        const got = partition('a');
        const expected = [['a']];
        expect(sortStringSlices(got)).toEqual(sortStringSlices(expected));
    });

    it('all same chars', () => {
        const got = partition('aaa');
        const expected = [['a', 'a', 'a'], ['a', 'aa'], ['aa', 'a'], ['aaa']];
        expect(sortStringSlices(got)).toEqual(sortStringSlices(expected));
    });

    it('no palindrome partitions beyond singles', () => {
        const got = partition('abc');
        const expected = [['a', 'b', 'c']];
        expect(sortStringSlices(got)).toEqual(sortStringSlices(expected));
    });

    it('full palindrome', () => {
        const got = partition('aba');
        const expected = [['a', 'b', 'a'], ['aba']];
        expect(sortStringSlices(got)).toEqual(sortStringSlices(expected));
    });

    it('two chars same', () => {
        const got = partition('bb');
        const expected = [['b', 'b'], ['bb']];
        expect(sortStringSlices(got)).toEqual(sortStringSlices(expected));
    });
});
