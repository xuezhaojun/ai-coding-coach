import { describe, it, expect } from "vitest";
import { TimeMap } from "./time_based_kv";

describe("TimeMap", () => {
  it("basic set and get", () => {
    const tm = new TimeMap();
    tm.set("foo", "bar", 1);
    expect(tm.get("foo", 1)).toBe("bar");
    expect(tm.get("foo", 3)).toBe("bar");
    tm.set("foo", "bar2", 4);
    expect(tm.get("foo", 4)).toBe("bar2");
    expect(tm.get("foo", 5)).toBe("bar2");
  });

  it("get before any set", () => {
    const tm = new TimeMap();
    expect(tm.get("foo", 1)).toBe("");
  });

  it("get with timestamp before first set", () => {
    const tm = new TimeMap();
    tm.set("key", "val", 5);
    expect(tm.get("key", 3)).toBe("");
  });

  it("multiple keys", () => {
    const tm = new TimeMap();
    tm.set("a", "v1", 1);
    tm.set("b", "v2", 1);
    expect(tm.get("a", 1)).toBe("v1");
    expect(tm.get("b", 1)).toBe("v2");
  });

  it("get exact timestamp", () => {
    const tm = new TimeMap();
    tm.set("k", "a", 1);
    tm.set("k", "b", 2);
    tm.set("k", "c", 3);
    expect(tm.get("k", 1)).toBe("a");
    expect(tm.get("k", 2)).toBe("b");
    expect(tm.get("k", 3)).toBe("c");
  });
});
