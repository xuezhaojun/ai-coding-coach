import { describe, it, expect } from "vitest";
import { findWords } from "./word_search_ii";

describe("findWords", () => {
    it.each([
        {
            name: "example 1",
            board: [["o", "a", "a", "n"], ["e", "t", "a", "e"], ["i", "h", "k", "r"], ["i", "f", "l", "v"]],
            words: ["oath", "pea", "eat", "rain"],
            want: ["eat", "oath"],
        },
        {
            name: "example 2",
            board: [["a", "b"], ["c", "d"]],
            words: ["abcb"],
            want: [],
        },
        {
            name: "single cell match",
            board: [["a"]],
            words: ["a"],
            want: ["a"],
        },
        {
            name: "single cell no match",
            board: [["a"]],
            words: ["b"],
            want: [],
        },
        {
            name: "multiple found",
            board: [["a", "b"], ["c", "d"]],
            words: ["ab", "ac", "abdc", "abcd", "dcba"],
            want: ["ab", "abdc", "ac"],
        },
        {
            name: "no words",
            board: [["a", "b"], ["c", "d"]],
            words: [],
            want: [],
        },
    ])("case: $name", ({ board, words, want }) => {
        // copy board to avoid mutation across parameterized cases
        const boardCopy = board.map((row) => [...row]);
        const got = findWords(boardCopy, words);
        expect([...got].sort()).toEqual([...want].sort());
    });
});
