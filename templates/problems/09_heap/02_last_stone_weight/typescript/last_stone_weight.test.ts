import { describe, it, expect } from "vitest";
import { lastStoneWeight } from "./last_stone_weight";

describe("lastStoneWeight", () => {
  it("example from problem", () => {
    expect(lastStoneWeight([2, 7, 4, 1, 8, 1])).toBe(1);
  });

  it("single stone", () => {
    expect(lastStoneWeight([1])).toBe(1);
  });

  it("two equal stones", () => {
    expect(lastStoneWeight([3, 3])).toBe(0);
  });

  it("two different stones", () => {
    expect(lastStoneWeight([1, 5])).toBe(4);
  });

  it("all same weight", () => {
    expect(lastStoneWeight([2, 2, 2, 2])).toBe(0);
  });

  it("descending weights", () => {
    expect(lastStoneWeight([10, 5, 3, 1])).toBe(1);
  });

  it("three stones", () => {
    expect(lastStoneWeight([3, 7, 2])).toBe(2);
  });
});
