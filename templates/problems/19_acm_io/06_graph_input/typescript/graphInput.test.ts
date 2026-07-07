import { describe, it, expect } from 'vitest';
import { solve } from './graphInput';

describe('GraphInput', () => {
    it('basic directed graph', () => {
        expect(solve("4 5\n1 2\n1 3\n2 3\n2 4\n3 4\n")).toBe("2\n2\n1\n0\n");
    });

    it('sparse graph', () => {
        expect(solve("3 2\n1 2\n3 2\n")).toBe("1\n0\n1\n");
    });

    it('single edge', () => {
        expect(solve("2 1\n1 2\n")).toBe("1\n0\n");
    });

    it('self loop', () => {
        expect(solve("2 2\n1 1\n1 2\n")).toBe("2\n0\n");
    });
});
