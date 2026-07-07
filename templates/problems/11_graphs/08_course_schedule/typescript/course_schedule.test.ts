import { describe, it, expect } from "vitest";
import { canFinish } from "./course_schedule";

describe("canFinish", () => {
    it.each([
        { name: "simple chain", numCourses: 2, prerequisites: [[1, 0]], want: true },
        { name: "cycle detected", numCourses: 2, prerequisites: [[1, 0], [0, 1]], want: false },
        { name: "no prerequisites", numCourses: 3, prerequisites: [], want: true },
        { name: "diamond dependency", numCourses: 4, prerequisites: [[1, 0], [2, 0], [3, 1], [3, 2]], want: true },
        { name: "three node cycle", numCourses: 3, prerequisites: [[0, 1], [1, 2], [2, 0]], want: false },
        { name: "single course", numCourses: 1, prerequisites: [], want: true },
        { name: "disconnected courses", numCourses: 5, prerequisites: [[1, 0], [3, 2]], want: true },
        { name: "node with two prerequisites and cycle", numCourses: 3, prerequisites: [[1, 0], [1, 2], [0, 1]], want: false },
    ])("case: $name", ({ numCourses, prerequisites, want }) => {
        const got = canFinish(numCourses, prerequisites);
        expect(got).toBe(want);
    });
});
