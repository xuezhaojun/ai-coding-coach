/**
 * Find shortest transformation sequence using BFS.
 * Time: O(n * m * 26) where n=wordList length, m=word length,
 * Space: O(n * m).
 */
export function solveLadderLength(beginWord: string, endWord: string, wordList: string[]): number {
    const wordSet = new Set(wordList);
    if (!wordSet.has(endWord)) return 0;
    if (beginWord === endWord) return 0;

    const queue: string[] = [beginWord];
    const visited = new Set<string>([beginWord]);
    let length = 1;
    let head = 0;

    while (head < queue.length) {
        const size = queue.length - head;
        for (let i = 0; i < size; i++) {
            const word = queue[head++];
            const chars = word.split("");
            for (let j = 0; j < chars.length; j++) {
                const original = chars[j];
                for (let c = 97; c <= 122; c++) {
                    const ch = String.fromCharCode(c);
                    if (ch === original) continue;
                    chars[j] = ch;
                    const next = chars.join("");
                    if (next === endWord) return length + 1;
                    if (wordSet.has(next) && !visited.has(next)) {
                        visited.add(next);
                        queue.push(next);
                    }
                }
                chars[j] = original;
            }
        }
        length++;
    }
    return 0;
}
