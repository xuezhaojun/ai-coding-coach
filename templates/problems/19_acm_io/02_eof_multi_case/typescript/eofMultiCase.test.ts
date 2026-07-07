import { describe, it, expect } from 'vitest';
import { solve } from './eofMultiCase';

describe('EofMultiCase', () => {
    it('basic multi-case', () => {
        expect(solve("1 2\n3 4\n5 6\n")).toBe("3\n7\n11\n");
    });

    it('negative numbers', () => {
        expect(solve("-1 2\n0 0\n")).toBe("1\n0\n");
    });

    it('single case', () => {
        expect(solve("100 200\n")).toBe("300\n");
    });

    it('large values', () => {
        expect(solve("1000000000 1000000000\n-1000000000 1\n")).toBe("2000000000\n-999999999\n");
    });
});
