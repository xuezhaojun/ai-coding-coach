import { describe, it, expect } from 'vitest';
import { solve } from './stringLines';

describe('StringLines', () => {
    it('basic reversal', () => {
        expect(solve("hello world\nI love coding\n")).toBe("world hello\ncoding love I\n");
    });

    it('four words', () => {
        expect(solve("a b c d\n")).toBe("d c b a\n");
    });

    it('single word', () => {
        expect(solve("hello\n")).toBe("hello\n");
    });

    it('multiple spaces', () => {
        expect(solve("Go   is   great\n")).toBe("great is Go\n");
    });
});
