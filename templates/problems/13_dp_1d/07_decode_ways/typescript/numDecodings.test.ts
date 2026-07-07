import { describe, it, expect } from 'vitest';
import { numDecodings } from './numDecodings';

describe('NumDecodings', () => {
    it('example 12', () => {
        expect(numDecodings('12')).toBe(2);
    });

    it('example 226', () => {
        expect(numDecodings('226')).toBe(3);
    });

    it('leading zero', () => {
        expect(numDecodings('06')).toBe(0);
    });

    it('single digit', () => {
        expect(numDecodings('1')).toBe(1);
    });

    it('example 11106', () => {
        expect(numDecodings('11106')).toBe(2);
    });

    it('all ones', () => {
        expect(numDecodings('1111')).toBe(5);
    });

    it('example 10', () => {
        expect(numDecodings('10')).toBe(1);
    });

    it('example 27', () => {
        expect(numDecodings('27')).toBe(1);
    });
});
