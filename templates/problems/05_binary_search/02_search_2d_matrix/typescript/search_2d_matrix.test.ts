import { describe, it, expect } from "vitest";
import { searchMatrix } from "./search_2d_matrix";

describe("searchMatrix", () => {
  it("found in middle row", () => {
    expect(
      searchMatrix(
        [
          [1, 3, 5, 7],
          [10, 11, 16, 20],
          [23, 30, 34, 60],
        ],
        3,
      ),
    ).toBe(true);
  });

  it("not found", () => {
    expect(
      searchMatrix(
        [
          [1, 3, 5, 7],
          [10, 11, 16, 20],
          [23, 30, 34, 60],
        ],
        13,
      ),
    ).toBe(false);
  });

  it("found first element", () => {
    expect(
      searchMatrix(
        [
          [1, 3, 5, 7],
          [10, 11, 16, 20],
        ],
        1,
      ),
    ).toBe(true);
  });

  it("found last element", () => {
    expect(
      searchMatrix(
        [
          [1, 3, 5, 7],
          [10, 11, 16, 20],
        ],
        20,
      ),
    ).toBe(true);
  });

  it("single element found", () => {
    expect(searchMatrix([[1]], 1)).toBe(true);
  });

  it("single element not found", () => {
    expect(searchMatrix([[1]], 2)).toBe(false);
  });

  it("single row", () => {
    expect(searchMatrix([[1, 3, 5]], 3)).toBe(true);
  });

  it("single column", () => {
    expect(searchMatrix([[1], [3], [5]], 5)).toBe(true);
  });
});
