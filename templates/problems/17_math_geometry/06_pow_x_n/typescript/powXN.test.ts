import { describe, it, expect } from 'vitest';
import { myPow } from './powXN';

describe('PowXN', () => {
    it('positive exponent', () => {
        expect(myPow(2.0, 10)).toBeCloseTo(1024.0, 5);
    });

    it('negative exponent', () => {
        expect(myPow(2.0, -2)).toBeCloseTo(0.25, 5);
    });

    it('zero exponent', () => {
        expect(myPow(5.0, 0)).toBeCloseTo(1.0, 5);
    });

    it('exponent 1', () => {
        expect(myPow(3.0, 1)).toBeCloseTo(3.0, 5);
    });

    it('fractional base', () => {
        expect(myPow(2.1, 3)).toBeCloseTo(9.261, 5);
    });

    it('base 1', () => {
        expect(myPow(1.0, 100)).toBeCloseTo(1.0, 5);
    });

    it('base 0', () => {
        expect(myPow(0.0, 5)).toBeCloseTo(0.0, 5);
    });
});
