import { describe, it, expect } from "vitest";
import { search } from "./binary_search";

describe("search", () => {
  it("found middle", () => {
    expect(search([-1, 0, 3, 5, 9, 12], 9)).toBe(4);
  });

  it("found first", () => {
    expect(search([-1, 0, 3, 5, 9, 12], -1)).toBe(0);
  });

  it("found last", () => {
    expect(search([-1, 0, 3, 5, 9, 12], 12)).toBe(5);
  });

  it("not found", () => {
    expect(search([-1, 0, 3, 5, 9, 12], 2)).toBe(-1);
  });

  it("single element found", () => {
    expect(search([5], 5)).toBe(0);
  });

  it("single element not found", () => {
    expect(search([5], 3)).toBe(-1);
  });

  it("two elements", () => {
    expect(search([1, 3], 3)).toBe(1);
  });

  it("empty array", () => {
    expect(search([], 1)).toBe(-1);
  });
});
