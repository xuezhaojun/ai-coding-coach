import { describe, it, expect } from "vitest";
import { alienOrder } from "./alien_dictionary";

function verify(words: string[], got: string, want: string): void {
    if (want === "") {
        expect(got).toBe("");
        return;
    }
    expect(got.length).toBe(want.length);
    const charPos = new Map<string, number>();
    for (let i = 0; i < got.length; i++) charPos.set(got[i], i);
    for (let i = 0; i < words.length - 1; i++) {
        const w1 = words[i];
        const w2 = words[i + 1];
        for (let j = 0; j < w1.length && j < w2.length; j++) {
            if (w1[j] !== w2[j]) {
                expect(charPos.get(w1[j])! <= charPos.get(w2[j])!).toBe(true);
                break;
            }
        }
    }
}

describe("alienOrder", () => {
    it.each([
        { name: "standard order", words: ["wrt", "wrf", "er", "ett", "rftt"], want: "wertf" },
        { name: "simple two words", words: ["z", "x"], want: "zx" },
        { name: "invalid order", words: ["z", "x", "z"], want: "" },
        { name: "single word", words: ["abc"], want: "abc" },
        { name: "prefix violation", words: ["abc", "ab"], want: "" },
        { name: "single characters", words: ["z", "z"], want: "z" },
    ])("case: $name", ({ words, want }) => {
        const got = alienOrder(words);
        verify(words, got, want);
    });
});
