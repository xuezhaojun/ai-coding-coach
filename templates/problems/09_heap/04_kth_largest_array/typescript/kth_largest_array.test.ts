import { describe, it, expect } from "vitest";
import { findKthLargest } from "./kth_largest_array";

describe("findKthLargest", () => {
  it("example 1", () => {
    expect(findKthLargest([3, 2, 1, 5, 6, 4], 2)).toBe(5);
  });

  it("example 2", () => {
    expect(findKthLargest([3, 2, 3, 1, 2, 4, 5, 5, 6], 4)).toBe(4);
  });

  it("single element", () => {
    expect(findKthLargest([1], 1)).toBe(1);
  });

  it("k equals length", () => {
    expect(findKthLargest([5, 3, 1], 3)).toBe(1);
  });

  it("all same elements", () => {
    expect(findKthLargest([7, 7, 7, 7], 2)).toBe(7);
  });

  it("negative numbers", () => {
    expect(findKthLargest([-1, -2, -3, -4], 1)).toBe(-1);
  });

  it("mixed positive and negative", () => {
    expect(findKthLargest([-1, 2, 0], 2)).toBe(0);
  });
});
