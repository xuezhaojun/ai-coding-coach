import { describe, it, expect } from "vitest";
import { findMedianSortedArrays } from "./median_two_sorted";

describe("findMedianSortedArrays", () => {
  it("odd total", () => {
    expect(findMedianSortedArrays([1, 3], [2])).toBe(2.0);
  });

  it("even total", () => {
    expect(findMedianSortedArrays([1, 2], [3, 4])).toBe(2.5);
  });

  it("first empty", () => {
    expect(findMedianSortedArrays([], [1])).toBe(1.0);
  });

  it("second empty", () => {
    expect(findMedianSortedArrays([2], [])).toBe(2.0);
  });

  it("no overlap", () => {
    expect(findMedianSortedArrays([1, 2], [3, 4, 5])).toBe(3.0);
  });

  it("interleaved", () => {
    expect(findMedianSortedArrays([1, 3, 5], [2, 4, 6])).toBe(3.5);
  });

  it("single elements", () => {
    expect(findMedianSortedArrays([1], [2])).toBe(1.5);
  });

  it("duplicates", () => {
    expect(findMedianSortedArrays([1, 1, 1], [1, 1, 1])).toBe(1.0);
  });
});
