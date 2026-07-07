import { describe, it, expect } from "vitest";
import { findMin } from "./find_min_rotated";

describe("findMin", () => {
  it("rotated", () => {
    expect(findMin([3, 4, 5, 1, 2])).toBe(1);
  });

  it("rotated once", () => {
    expect(findMin([4, 5, 6, 7, 0, 1, 2])).toBe(0);
  });

  it("not rotated", () => {
    expect(findMin([11, 13, 15, 17])).toBe(11);
  });

  it("single element", () => {
    expect(findMin([1])).toBe(1);
  });

  it("two elements rotated", () => {
    expect(findMin([2, 1])).toBe(1);
  });

  it("two elements sorted", () => {
    expect(findMin([1, 2])).toBe(1);
  });

  it("min at end", () => {
    expect(findMin([2, 3, 4, 5, 1])).toBe(1);
  });

  it("min at start", () => {
    expect(findMin([1, 2, 3, 4, 5])).toBe(1);
  });
});
