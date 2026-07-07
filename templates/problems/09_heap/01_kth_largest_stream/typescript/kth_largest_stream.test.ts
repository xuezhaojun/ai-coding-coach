import { describe, it, expect } from "vitest";
import { KthLargest } from "./kth_largest_stream";

describe("KthLargest", () => {
  it("example from problem", () => {
    const kl = new KthLargest(3, [4, 5, 8, 2]);
    const adds = [3, 5, 10, 9, 4];
    const expects = [4, 5, 5, 8, 8];
    for (let i = 0; i < adds.length; i++) {
      expect(kl.add(adds[i])).toBe(expects[i]);
    }
  });

  it("k equals 1", () => {
    const kl = new KthLargest(1, []);
    const adds = [-1, 1, -2, -4, 3];
    const expects = [-1, 1, 1, 1, 3];
    for (let i = 0; i < adds.length; i++) {
      expect(kl.add(adds[i])).toBe(expects[i]);
    }
  });

  it("single initial element", () => {
    const kl = new KthLargest(1, [5]);
    expect(kl.add(3)).toBe(5);
    expect(kl.add(7)).toBe(7);
  });

  it("all same values", () => {
    const kl = new KthLargest(2, [0]);
    expect(kl.add(0)).toBe(0);
    expect(kl.add(0)).toBe(0);
    expect(kl.add(0)).toBe(0);
  });

  it("negative numbers", () => {
    const kl = new KthLargest(2, [-5, -3]);
    expect(kl.add(-1)).toBe(-3);
    expect(kl.add(-7)).toBe(-3);
    expect(kl.add(0)).toBe(-1);
  });

  it("large k with enough initial elements", () => {
    const kl = new KthLargest(3, [1, 2, 3, 4, 5]);
    expect(kl.add(6)).toBe(4);
  });
});
