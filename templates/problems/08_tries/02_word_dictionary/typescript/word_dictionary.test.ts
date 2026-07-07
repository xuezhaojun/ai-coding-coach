import { describe, it, expect } from "vitest";
import { WordDictionary } from "./word_dictionary";

describe("WordDictionary", () => {
    it("basic with wildcards", () => {
        const wd = new WordDictionary();
        for (const w of ["bad", "dad", "mad"]) {
            wd.addWord(w);
        }
        expect(wd.search("pad")).toBe(false);
        expect(wd.search("bad")).toBe(true);
        expect(wd.search(".ad")).toBe(true);
        expect(wd.search("b..")).toBe(true);
        expect(wd.search("...")).toBe(true);
        expect(wd.search("....")).toBe(false);
    });

    it("exact match only", () => {
        const wd = new WordDictionary();
        wd.addWord("hello");
        expect(wd.search("hello")).toBe(true);
        expect(wd.search("hell")).toBe(false);
        expect(wd.search("helloo")).toBe(false);
    });

    it("all wildcards", () => {
        const wd = new WordDictionary();
        for (const w of ["ab", "cd"]) {
            wd.addWord(w);
        }
        expect(wd.search("..")).toBe(true);
        expect(wd.search(".")).toBe(false);
        expect(wd.search("...")).toBe(false);
    });

    it("single char", () => {
        const wd = new WordDictionary();
        wd.addWord("a");
        expect(wd.search(".")).toBe(true);
        expect(wd.search("a")).toBe(true);
        expect(wd.search("..")).toBe(false);
    });

    it("mixed wildcards", () => {
        const wd = new WordDictionary();
        for (const w of ["apple", "ample"]) {
            wd.addWord(w);
        }
        expect(wd.search("a.ple")).toBe(true);
        expect(wd.search("a..le")).toBe(true);
        expect(wd.search("a...e")).toBe(true);
        expect(wd.search("a....")).toBe(true);
        expect(wd.search("b....")).toBe(false);
    });

    it("no words added", () => {
        const wd = new WordDictionary();
        expect(wd.search("a")).toBe(false);
        expect(wd.search(".")).toBe(false);
    });
});
