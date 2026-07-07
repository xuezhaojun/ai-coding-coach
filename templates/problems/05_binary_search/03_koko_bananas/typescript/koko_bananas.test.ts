import { describe, it, expect } from "vitest";
import { minEatingSpeed } from "./koko_bananas";

describe("minEatingSpeed", () => {
  it("example 1", () => {
    expect(minEatingSpeed([3, 6, 7, 11], 8)).toBe(4);
  });

  it("example 2", () => {
    expect(minEatingSpeed([30, 11, 23, 4, 20], 5)).toBe(30);
  });

  it("example 3", () => {
    expect(minEatingSpeed([30, 11, 23, 4, 20], 6)).toBe(23);
  });

  it("single pile exact", () => {
    expect(minEatingSpeed([10], 1)).toBe(10);
  });

  it("single pile slow", () => {
    expect(minEatingSpeed([10], 10)).toBe(1);
  });

  it("all ones", () => {
    expect(minEatingSpeed([1, 1, 1], 3)).toBe(1);
  });

  it("large h", () => {
    expect(minEatingSpeed([3, 6, 7, 11], 100)).toBe(1);
  });
});
