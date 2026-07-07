import { describe, it, expect } from 'vitest';
import { solve } from './sentinelTerminate';

describe('SentinelTerminate', () => {
    it('basic with sentinel', () => {
        expect(solve("1 2\n3 4\n0 0\n")).toBe("3\n7\n");
    });

    it('immediate sentinel', () => {
        expect(solve("0 0\n")).toBe("");
    });

    it('negative before sentinel', () => {
        expect(solve("-5 10\n0 0\n")).toBe("5\n");
    });

    it('zeros not at sentinel position', () => {
        expect(solve("0 5\n5 0\n0 0\n")).toBe("5\n5\n");
    });
});
