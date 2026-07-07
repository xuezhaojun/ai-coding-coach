class TrieNode {
    children: (TrieNode | null)[] = new Array(26).fill(null);
    word: string = "";
}

export function solveFindWords(board: string[][], words: string[]): string[] {
    const root = new TrieNode();
    for (const w of words) {
        let node = root;
        for (const ch of w) {
            const idx = ch.charCodeAt(0) - "a".charCodeAt(0);
            if (node.children[idx] === null) {
                node.children[idx] = new TrieNode();
            }
            node = node.children[idx]!;
        }
        node.word = w;
    }

    const rows = board.length;
    const cols = rows > 0 ? board[0].length : 0;
    const result: string[] = [];

    function dfs(r: number, c: number, node: TrieNode): void {
        if (r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] === "#") {
            return;
        }
        const ch = board[r][c];
        const next = node.children[ch.charCodeAt(0) - "a".charCodeAt(0)];
        if (next === null) {
            return;
        }
        if (next.word !== "") {
            result.push(next.word);
            next.word = "";
        }
        board[r][c] = "#";
        dfs(r + 1, c, next);
        dfs(r - 1, c, next);
        dfs(r, c + 1, next);
        dfs(r, c - 1, next);
        board[r][c] = ch;
    }

    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            dfs(r, c, root);
        }
    }
    return result;
}
