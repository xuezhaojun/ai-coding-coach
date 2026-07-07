import { describe, it, expect } from "vitest";
import { kClosest } from "./k_closest_points";

function sortPoints(pts: number[][]): number[][] {
  return pts
    .slice()
    .sort((a, b) => (a[0] !== b[0] ? a[0] - b[0] : a[1] - b[1]));
}

describe("kClosest", () => {
  it("example 1", () => {
    expect(kClosest([[1, 3], [-2, 2]], 1)).toEqual([[-2, 2]]);
  });

  it("example 2", () => {
    const got = kClosest([[3, 3], [5, -1], [-2, 4]], 2);
    expect(sortPoints(got)).toEqual(sortPoints([[-2, 4], [3, 3]]));
  });

  it("single point", () => {
    expect(kClosest([[0, 1]], 1)).toEqual([[0, 1]]);
  });

  it("all points same distance", () => {
    const got = kClosest([[1, 0], [0, 1], [-1, 0], [0, -1]], 2);
    expect(got.length).toBe(2);
  });

  it("origin point", () => {
    expect(kClosest([[0, 0], [1, 1], [2, 2]], 1)).toEqual([[0, 0]]);
  });

  it("k equals length", () => {
    const got = kClosest([[1, 2], [3, 4]], 2);
    expect(sortPoints(got)).toEqual(sortPoints([[1, 2], [3, 4]]));
  });
});
