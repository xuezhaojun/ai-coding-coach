import { describe, it, expect } from "vitest";
import { searchRotated } from "./search_rotated";

describe("searchRotated", () => {
  it("found in left half", () => {
    expect(searchRotated([4, 5, 6, 7, 0, 1, 2], 0)).toBe(4);
  });

  it("found in right half", () => {
    expect(searchRotated([4, 5, 6, 7, 0, 1, 2], 5)).toBe(1);
  });

  it("not found", () => {
    expect(searchRotated([4, 5, 6, 7, 0, 1, 2], 3)).toBe(-1);
  });

  it("single element found", () => {
    expect(searchRotated([1], 1)).toBe(0);
  });

  it("single element not found", () => {
    expect(searchRotated([1], 0)).toBe(-1);
  });

  it("not rotated found", () => {
    expect(searchRotated([1, 2, 3, 4, 5], 3)).toBe(2);
  });

  it("two elements", () => {
    expect(searchRotated([3, 1], 1)).toBe(1);
  });

  it("target at pivot", () => {
    expect(searchRotated([6, 7, 1, 2, 3, 4, 5], 1)).toBe(2);
  });
});
