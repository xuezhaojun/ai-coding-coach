import { describe, it, expect } from 'vitest';
import { isMatch } from './isMatch';

describe('IsMatch', () => {
    it('no match', () => {
        expect(isMatch('aa', 'a')).toBe(false);
    });

    it('star matches multiple', () => {
        expect(isMatch('aa', 'a*')).toBe(true);
    });

    it('dot star matches all', () => {
        expect(isMatch('ab', '.*')).toBe(true);
    });

    it('mixed pattern', () => {
        expect(isMatch('aab', 'c*a*b')).toBe(true);
    });

    it('empty string empty pattern', () => {
        expect(isMatch('', '')).toBe(true);
    });

    it('empty string star pattern', () => {
        expect(isMatch('', 'a*')).toBe(true);
    });

    it('complex pattern', () => {
        expect(isMatch('mississippi', 'mis*is*ip*.')).toBe(true);
    });

    it('dot matches single', () => {
        expect(isMatch('ab', '.b')).toBe(true);
    });
});
