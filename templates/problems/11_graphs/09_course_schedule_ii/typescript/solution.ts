/**
 * Return a topological ordering of courses using Kahn's algorithm (BFS).
 * Time: O(V + E), Space: O(V + E).
 */
export function solveFindOrder(numCourses: number, prerequisites: number[][]): number[] {
    const graph: number[][] = Array.from({ length: numCourses }, () => []);
    const inDegree = new Array(numCourses).fill(0);
    for (const p of prerequisites) {
        graph[p[1]].push(p[0]);
        inDegree[p[0]]++;
    }

    const queue: number[] = [];
    for (let i = 0; i < numCourses; i++) {
        if (inDegree[i] === 0) queue.push(i);
    }
    const order: number[] = [];
    let head = 0;
    while (head < queue.length) {
        const node = queue[head++];
        order.push(node);
        for (const nxt of graph[node]) {
            inDegree[nxt]--;
            if (inDegree[nxt] === 0) queue.push(nxt);
        }
    }

    if (order.length !== numCourses) return [];
    return order;
}
