import { describe, it, expect } from 'vitest';
import { letterCombinations } from './letterCombinations';

describe('LetterCombinations', () => {
    it('example 23', () => {
        const got = letterCombinations('23').sort();
        const expected = ['ad', 'ae', 'af', 'bd', 'be', 'bf', 'cd', 'ce', 'cf'].sort();
        expect(got).toEqual(expected);
    });

    it('empty string', () => {
        const got = letterCombinations('').sort();
        const expected: string[] = [];
        expect(got).toEqual(expected);
    });

    it('single digit 2', () => {
        const got = letterCombinations('2').sort();
        const expected = ['a', 'b', 'c'].sort();
        expect(got).toEqual(expected);
    });

    it('digit 7 with four letters', () => {
        const got = letterCombinations('7').sort();
        const expected = ['p', 'q', 'r', 's'].sort();
        expect(got).toEqual(expected);
    });

    it('digit 9 with four letters', () => {
        const got = letterCombinations('9').sort();
        const expected = ['w', 'x', 'y', 'z'].sort();
        expect(got).toEqual(expected);
    });

    it('three digits', () => {
        const got = letterCombinations('234').sort();
        const expected = [
            'adg', 'adh', 'adi', 'aeg', 'aeh', 'aei', 'afg', 'afh', 'afi',
            'bdg', 'bdh', 'bdi', 'beg', 'beh', 'bei', 'bfg', 'bfh', 'bfi',
            'cdg', 'cdh', 'cdi', 'ceg', 'ceh', 'cei', 'cfg', 'cfh', 'cfi',
        ].sort();
        expect(got).toEqual(expected);
    });
});
