import { describe, it, expect } from "vitest";
import { findOrder } from "./course_schedule_ii";

describe("findOrder", () => {
    it.each([
        { name: "two courses with dependency", numCourses: 2, prerequisites: [[1, 0]], wantLen: 2, hasOrder: true },
        { name: "four courses", numCourses: 4, prerequisites: [[1, 0], [2, 0], [3, 1], [3, 2]], wantLen: 4, hasOrder: true },
        { name: "cycle returns empty", numCourses: 2, prerequisites: [[1, 0], [0, 1]], wantLen: 0, hasOrder: false },
        { name: "single course", numCourses: 1, prerequisites: [], wantLen: 1, hasOrder: true },
        { name: "no prerequisites", numCourses: 3, prerequisites: [], wantLen: 3, hasOrder: true },
        { name: "three node cycle", numCourses: 3, prerequisites: [[0, 1], [1, 2], [2, 0]], wantLen: 0, hasOrder: false },
    ])("case: $name", ({ numCourses, prerequisites, wantLen, hasOrder }) => {
        const got = findOrder(numCourses, prerequisites);
        if (!hasOrder) {
            expect(got).toEqual([]);
            return;
        }
        expect(got.length).toBe(wantLen);
        const pos = new Map<number, number>();
        got.forEach((c, i) => pos.set(c, i));
        for (const p of prerequisites) {
            if (pos.get(p[1])! > pos.get(p[0])!) {
                throw new Error(`findOrder(): prerequisite ${p} violated in order ${got}`);
            }
        }
    });
});
