/**
 * Detect if all courses can be finished (no cycle) using DFS.
 * Time: O(V + E), Space: O(V + E).
 */
export function solveCanFinish(numCourses: number, prerequisites: number[][]): boolean {
    const graph: number[][] = Array.from({ length: numCourses }, () => []);
    for (const p of prerequisites) {
        graph[p[1]].push(p[0]);
    }

    // 0=unvisited, 1=visiting, 2=visited
    const state = new Array(numCourses).fill(0);

    const hasCycle = (node: number): boolean => {
        if (state[node] === 1) return true;
        if (state[node] === 2) return false;
        state[node] = 1;
        for (const nxt of graph[node]) {
            if (hasCycle(nxt)) return true;
        }
        state[node] = 2;
        return false;
    };

    for (let i = 0; i < numCourses; i++) {
        if (hasCycle(i)) return false;
    }
    return true;
}
