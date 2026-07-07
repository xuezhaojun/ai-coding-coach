import { describe, it, expect } from 'vitest';
import { solve } from './fixedFormat';

describe('FixedFormat', () => {
    it('basic sum', () => {
        expect(solve("5\n1 2 3 4 5\n")).toBe("15\n");
    });

    it('negative numbers', () => {
        expect(solve("3\n-1 0 1\n")).toBe("0\n");
    });

    it('single element', () => {
        expect(solve("1\n42\n")).toBe("42\n");
    });

    it('large values', () => {
        expect(solve("2\n1000000000 -1000000000\n")).toBe("0\n");
    });
});
