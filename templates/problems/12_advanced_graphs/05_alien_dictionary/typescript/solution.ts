/**
 * Determine character ordering from sorted alien words using topological sort.
 * Time: O(C) where C is total characters across all words,
 * Space: O(1) since alphabet is bounded.
 */
export function solveAlienOrder(words: string[]): string {
    // Build graph of character ordering
    const graph = new Map<string, Set<string>>();
    const inDegree = new Map<string, number>();

    // Initialize all characters
    for (const w of words) {
        for (const ch of w) {
            if (!graph.has(ch)) {
                graph.set(ch, new Set());
                inDegree.set(ch, 0);
            }
        }
    }

    // Compare adjacent words to find ordering constraints
    for (let i = 0; i < words.length - 1; i++) {
        const w1 = words[i];
        const w2 = words[i + 1];
        // Check prefix violation
        if (w1.length > w2.length) {
            let prefix = true;
            for (let j = 0; j < w2.length; j++) {
                if (w1[j] !== w2[j]) {
                    prefix = false;
                    break;
                }
            }
            if (prefix) return "";
        }
        for (let j = 0; j < w1.length && j < w2.length; j++) {
            if (w1[j] !== w2[j]) {
                if (!graph.get(w1[j])!.has(w2[j])) {
                    graph.get(w1[j])!.add(w2[j]);
                    inDegree.set(w2[j], inDegree.get(w2[j])! + 1);
                }
                break;
            }
        }
    }

    // BFS topological sort
    const queue: string[] = [];
    for (const ch of graph.keys()) {
        if (inDegree.get(ch) === 0) queue.push(ch);
    }

    const result: string[] = [];
    let head = 0;
    while (head < queue.length) {
        const ch = queue[head++];
        result.push(ch);
        for (const nxt of graph.get(ch)!) {
            inDegree.set(nxt, inDegree.get(nxt)! - 1);
            if (inDegree.get(nxt) === 0) queue.push(nxt);
        }
    }

    if (result.length !== graph.size) return "";
    return result.join("");
}
