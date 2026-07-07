import { describe, it, expect } from 'vitest';
import { encode, decode } from './encodeDecodeStrings';

describe('EncodeDecodeStrings', () => {
    it('basic strings', () => {
        const strs = ['hello', 'world'];
        expect(decode(encode(strs))).toEqual(strs);
    });

    it('empty list', () => {
        const strs: string[] = [];
        expect(decode(encode(strs))).toEqual(strs);
    });

    it('single empty string', () => {
        const strs = [''];
        expect(decode(encode(strs))).toEqual(strs);
    });

    it('multiple empty strings', () => {
        const strs = ['', '', ''];
        expect(decode(encode(strs))).toEqual(strs);
    });

    it('strings with special characters', () => {
        const strs = ['he:llo', 'wor#ld', 'foo;bar'];
        expect(decode(encode(strs))).toEqual(strs);
    });

    it('strings with delimiters and colons', () => {
        const strs = ['4:abcd', '3:xyz'];
        expect(decode(encode(strs))).toEqual(strs);
    });

    it('single character strings', () => {
        const strs = ['a', 'b', 'c'];
        expect(decode(encode(strs))).toEqual(strs);
    });

    it('mixed empty and non-empty', () => {
        const strs = ['', 'a', '', 'b', ''];
        expect(decode(encode(strs))).toEqual(strs);
    });
});
