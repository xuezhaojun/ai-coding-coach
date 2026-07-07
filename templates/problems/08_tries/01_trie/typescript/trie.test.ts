import { describe, it, expect } from "vitest";
import { Trie } from "./trie";

describe("Trie", () => {
    it("basic operations", () => {
        const trie = new Trie();
        trie.insert("apple");
        expect(trie.search("apple")).toBe(true);
        expect(trie.search("app")).toBe(false);
        expect(trie.search("apples")).toBe(false);
        expect(trie.startsWith("app")).toBe(true);
        expect(trie.startsWith("apple")).toBe(true);
        expect(trie.startsWith("b")).toBe(false);
    });

    it("multiple words", () => {
        const trie = new Trie();
        for (const w of ["apple", "app", "banana"]) {
            trie.insert(w);
        }
        expect(trie.search("apple")).toBe(true);
        expect(trie.search("app")).toBe(true);
        expect(trie.search("banana")).toBe(true);
        expect(trie.search("ban")).toBe(false);
        expect(trie.startsWith("app")).toBe(true);
        expect(trie.startsWith("ban")).toBe(true);
        expect(trie.startsWith("cat")).toBe(false);
    });

    it("empty string", () => {
        const trie = new Trie();
        trie.insert("");
        expect(trie.search("")).toBe(true);
        expect(trie.search("a")).toBe(false);
        expect(trie.startsWith("")).toBe(true);
    });

    it("single char words", () => {
        const trie = new Trie();
        for (const w of ["a", "b", "c"]) {
            trie.insert(w);
        }
        expect(trie.search("a")).toBe(true);
        expect(trie.search("b")).toBe(true);
        expect(trie.search("d")).toBe(false);
        expect(trie.startsWith("a")).toBe(true);
        expect(trie.startsWith("d")).toBe(false);
    });

    it("overlapping words", () => {
        const trie = new Trie();
        for (const w of ["the", "then", "them", "there"]) {
            trie.insert(w);
        }
        expect(trie.search("the")).toBe(true);
        expect(trie.search("then")).toBe(true);
        expect(trie.search("they")).toBe(false);
        expect(trie.search("there")).toBe(true);
        expect(trie.startsWith("the")).toBe(true);
        expect(trie.startsWith("th")).toBe(true);
        expect(trie.startsWith("thx")).toBe(false);
    });
});
