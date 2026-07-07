import { describe, it, expect } from 'vitest';
import { groupAnagrams } from './groupAnagrams';

function sortGroups(groups: string[][]): string[][] {
    for (const group of groups) {
        group.sort();
    }
    groups.sort((a, b) => {
        if (a.length !== b.length) return a.length - b.length;
        for (let k = 0; k < a.length; k++) {
            if (a[k] !== b[k]) return a[k] < b[k] ? -1 : 1;
        }
        return 0;
    });
    return groups;
}

describe('GroupAnagrams', () => {
    it('mixed anagram groups', () => {
        const result = sortGroups(groupAnagrams(['eat', 'tea', 'tan', 'ate', 'nat', 'bat']));
        const expected = sortGroups([['bat'], ['nat', 'tan'], ['ate', 'eat', 'tea']]);
        expect(result).toEqual(expected);
    });

    it('single empty string', () => {
        const result = sortGroups(groupAnagrams(['']));
        const expected = sortGroups([['']]);
        expect(result).toEqual(expected);
    });

    it('single non-empty string', () => {
        const result = sortGroups(groupAnagrams(['a']));
        const expected = sortGroups([['a']]);
        expect(result).toEqual(expected);
    });

    it('no anagrams', () => {
        const result = sortGroups(groupAnagrams(['abc', 'def', 'ghi']));
        const expected = sortGroups([['abc'], ['def'], ['ghi']]);
        expect(result).toEqual(expected);
    });

    it('all anagrams', () => {
        const result = sortGroups(groupAnagrams(['abc', 'bca', 'cab']));
        const expected = sortGroups([['abc', 'bca', 'cab']]);
        expect(result).toEqual(expected);
    });

    it('empty input', () => {
        const result = groupAnagrams([]);
        expect(result).toEqual([]);
    });

    it('multiple empty strings', () => {
        const result = sortGroups(groupAnagrams(['', '']));
        const expected = sortGroups([['', '']]);
        expect(result).toEqual(expected);
    });
});
