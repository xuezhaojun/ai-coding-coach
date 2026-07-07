import { describe, it, expect } from 'vitest';
import { checkValidString } from './validParenString';

describe('ValidParenString', () => {
    it('simple valid', () => {
        expect(checkValidString('()')).toBe(true);
    });

    it('star as empty', () => {
        expect(checkValidString('(*)')).toBe(true);
    });

    it('star as paren', () => {
        expect(checkValidString('(*))')).toBe(true);
    });

    it('empty string', () => {
        expect(checkValidString('')).toBe(true);
    });

    it('only stars', () => {
        expect(checkValidString('***')).toBe(true);
    });

    it('unmatched open', () => {
        expect(checkValidString('((')).toBe(false);
    });

    it('unmatched close', () => {
        expect(checkValidString('))')).toBe(false);
    });

    it('star as open', () => {
        expect(checkValidString('*(')).toBe(false);
    });
});
