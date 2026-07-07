import { describe, it, expect } from 'vitest';
import { DetectSquares } from './detectSquares';

describe('DetectSquares', () => {
    it('basic square', () => {
        const ds = new DetectSquares();
        for (const p of [[3, 10], [11, 10], [11, 2]]) {
            ds.add(p);
        }
        expect(ds.count([3, 2])).toBe(1);
    });

    it('no square possible', () => {
        const ds = new DetectSquares();
        for (const p of [[1, 1], [2, 2]]) {
            ds.add(p);
        }
        expect(ds.count([3, 3])).toBe(0);
    });

    it('duplicate points multiply count', () => {
        const ds = new DetectSquares();
        for (const p of [[3, 10], [3, 10], [11, 10], [11, 2]]) {
            ds.add(p);
        }
        expect(ds.count([3, 2])).toBe(2);
    });

    it('multiple squares from one query', () => {
        const ds = new DetectSquares();
        for (const p of [[0, 0], [1, 0], [1, 1], [0, 1], [2, 0], [2, 1]]) {
            ds.add(p);
        }
        expect(ds.count([0, 0])).toBe(2);
    });

    it('no points added', () => {
        const ds = new DetectSquares();
        expect(ds.count([0, 0])).toBe(0);
    });

    it('collinear points', () => {
        const ds = new DetectSquares();
        for (const p of [[1, 1], [2, 1], [3, 1]]) {
            ds.add(p);
        }
        expect(ds.count([0, 1])).toBe(0);
    });
});
