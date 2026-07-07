import { describe, it, expect } from "vitest";
import { leastInterval } from "./task_scheduler";

describe("leastInterval", () => {
  it("example 1", () => {
    expect(leastInterval(["A", "A", "A", "B", "B", "B"], 2)).toBe(8);
  });

  it("no cooldown", () => {
    expect(leastInterval(["A", "A", "A", "B", "B", "B"], 0)).toBe(6);
  });

  it("example 3", () => {
    expect(
      leastInterval(
        ["A", "A", "A", "A", "A", "A", "B", "C", "D", "E", "F", "G"],
        2,
      ),
    ).toBe(16);
  });

  it("single task", () => {
    expect(leastInterval(["A"], 5)).toBe(1);
  });

  it("all different tasks", () => {
    expect(leastInterval(["A", "B", "C", "D"], 3)).toBe(4);
  });

  it("two tasks with cooldown 1", () => {
    expect(leastInterval(["A", "A", "B", "B"], 1)).toBe(4);
  });
});
