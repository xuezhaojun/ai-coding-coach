import { describe, it, expect } from 'vitest';
import { exist } from './wordSearch';

describe('WordSearch', () => {
    it('word exists', () => {
        const board = [
            ['A', 'B', 'C', 'E'],
            ['S', 'F', 'C', 'S'],
            ['A', 'D', 'E', 'E'],
        ];
        expect(exist(board, 'ABCCED')).toBe(true);
    });

    it('word exists path SEE', () => {
        const board = [
            ['A', 'B', 'C', 'E'],
            ['S', 'F', 'C', 'S'],
            ['A', 'D', 'E', 'E'],
        ];
        expect(exist(board, 'SEE')).toBe(true);
    });

    it('word does not exist', () => {
        const board = [
            ['A', 'B', 'C', 'E'],
            ['S', 'F', 'C', 'S'],
            ['A', 'D', 'E', 'E'],
        ];
        expect(exist(board, 'ABCB')).toBe(false);
    });

    it('single cell match', () => {
        expect(exist([['A']], 'A')).toBe(true);
    });

    it('single cell no match', () => {
        expect(exist([['A']], 'B')).toBe(false);
    });

    it('word longer than board cells', () => {
        const board = [
            ['A', 'B'],
            ['C', 'D'],
        ];
        expect(exist(board, 'ABCDA')).toBe(false);
    });

    it('snake path', () => {
        const board = [
            ['A', 'B', 'C'],
            ['D', 'E', 'F'],
            ['G', 'H', 'I'],
        ];
        expect(exist(board, 'ABCFEDGHI')).toBe(true);
    });
});
