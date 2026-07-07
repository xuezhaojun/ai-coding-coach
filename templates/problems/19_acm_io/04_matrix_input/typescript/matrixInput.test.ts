import { describe, it, expect } from 'vitest';
import { solve } from './matrixInput';

describe('MatrixInput', () => {
    it('2x3 transpose', () => {
        expect(solve("2 3\n1 2 3\n4 5 6\n")).toBe("1 4\n2 5\n3 6\n");
    });

    it('3x3 identity', () => {
        expect(solve("3 3\n1 0 0\n0 1 0\n0 0 1\n")).toBe("1 0 0\n0 1 0\n0 0 1\n");
    });

    it('1x1', () => {
        expect(solve("1 1\n42\n")).toBe("42\n");
    });

    it('1x3 row to 3x1', () => {
        expect(solve("1 3\n10 20 30\n")).toBe("10\n20\n30\n");
    });
});
