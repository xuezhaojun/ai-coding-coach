class SolveTrieNode {
    children: Map<string, SolveTrieNode> = new Map();
    isEnd: boolean = false;
}

export class SolveTrie {
    root: SolveTrieNode;

    constructor() {
        this.root = new SolveTrieNode();
    }

    insert(word: string): void {
        let node = this.root;
        for (const c of word) {
            if (!node.children.has(c)) {
                node.children.set(c, new SolveTrieNode());
            }
            node = node.children.get(c)!;
        }
        node.isEnd = true;
    }

    search(word: string): boolean {
        const node = this.findNode(word);
        return node !== null && node.isEnd;
    }

    startsWith(prefix: string): boolean {
        return this.findNode(prefix) !== null;
    }

    private findNode(prefix: string): SolveTrieNode | null {
        let node = this.root;
        for (const c of prefix) {
            const next = node.children.get(c);
            if (next === undefined) {
                return null;
            }
            node = next;
        }
        return node;
    }
}
