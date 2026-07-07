import { describe, it, expect } from "vitest";
import { ladderLength } from "./word_ladder";

describe("ladderLength", () => {
    it.each([
        {
            name: "standard transformation",
            beginWord: "hit", endWord: "cog",
            wordList: ["hot", "dot", "dog", "lot", "log", "cog"],
            want: 5,
        },
        {
            name: "no valid transformation",
            beginWord: "hit", endWord: "cog",
            wordList: ["hot", "dot", "dog", "lot", "log"],
            want: 0,
        },
        { name: "one step", beginWord: "hot", endWord: "dot", wordList: ["dot"], want: 2 },
        { name: "same begin and end", beginWord: "hit", endWord: "hit", wordList: ["hit"], want: 0 },
        { name: "end word not in list", beginWord: "abc", endWord: "xyz", wordList: ["abc", "xbc"], want: 0 },
        { name: "longer path", beginWord: "a", endWord: "c", wordList: ["a", "b", "c"], want: 2 },
        { name: "two letter words", beginWord: "ab", endWord: "cd", wordList: ["ad", "cb", "cd"], want: 3 },
    ])("case: $name", ({ beginWord, endWord, wordList, want }) => {
        const got = ladderLength(beginWord, endWord, wordList);
        expect(got).toBe(want);
    });
});
