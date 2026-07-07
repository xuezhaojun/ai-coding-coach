import { describe, it, expect } from 'vitest';
import { isHappy } from './happyNumber';

describe('HappyNumber', () => {
    it('happy 19', () => {
        expect(isHappy(19)).toBe(true);
    });

    it('happy 1', () => {
        expect(isHappy(1)).toBe(true);
    });

    it('not happy 2', () => {
        expect(isHappy(2)).toBe(false);
    });

    it('happy 7', () => {
        expect(isHappy(7)).toBe(true);
    });

    it('not happy 4', () => {
        expect(isHappy(4)).toBe(false);
    });

    it('happy 100', () => {
        expect(isHappy(100)).toBe(true);
    });

    it('not happy 20', () => {
        expect(isHappy(20)).toBe(false);
    });
});
