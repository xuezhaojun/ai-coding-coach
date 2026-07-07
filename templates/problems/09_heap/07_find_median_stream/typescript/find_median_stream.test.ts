import { describe, it, expect } from "vitest";
import { MedianFinder } from "./find_median_stream";

describe("MedianFinder", () => {
  it("example from problem", () => {
    const mf = new MedianFinder();
    mf.addNum(1);
    expect(mf.findMedian()).toBe(1.0);
    mf.addNum(2);
    expect(mf.findMedian()).toBe(1.5);
    mf.addNum(3);
    expect(mf.findMedian()).toBe(2.0);
  });

  it("single element", () => {
    const mf = new MedianFinder();
    mf.addNum(5);
    expect(mf.findMedian()).toBe(5.0);
  });

  it("two elements", () => {
    const mf = new MedianFinder();
    mf.addNum(1);
    expect(mf.findMedian()).toBe(1.0);
    mf.addNum(2);
    expect(mf.findMedian()).toBe(1.5);
  });

  it("negative numbers", () => {
    const mf = new MedianFinder();
    mf.addNum(-1);
    expect(mf.findMedian()).toBe(-1.0);
    mf.addNum(-2);
    expect(mf.findMedian()).toBe(-1.5);
    mf.addNum(-3);
    expect(mf.findMedian()).toBe(-2.0);
  });

  it("duplicates", () => {
    const mf = new MedianFinder();
    for (let i = 0; i < 4; i++) mf.addNum(1);
    expect(mf.findMedian()).toBe(1.0);
  });

  it("descending order", () => {
    const mf = new MedianFinder();
    const adds = [5, 4, 3, 2, 1];
    const medians = [5.0, 4.5, 4.0, 3.5, 3.0];
    for (let i = 0; i < adds.length; i++) {
      mf.addNum(adds[i]);
      expect(mf.findMedian()).toBe(medians[i]);
    }
  });

  it("large gap", () => {
    const mf = new MedianFinder();
    mf.addNum(0);
    expect(mf.findMedian()).toBe(0.0);
    mf.addNum(100);
    expect(mf.findMedian()).toBe(50.0);
  });
});
