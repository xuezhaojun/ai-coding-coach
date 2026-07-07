class SolveWDNode {
    children: Map<string, SolveWDNode> = new Map();
    isEnd: boolean = false;
}

export class SolveWordDictionary {
    root: SolveWDNode;

    constructor() {
        this.root = new SolveWDNode();
    }

    addWord(word: string): void {
        let node = this.root;
        for (const c of word) {
            if (!node.children.has(c)) {
                node.children.set(c, new SolveWDNode());
            }
            node = node.children.get(c)!;
        }
        node.isEnd = true;
    }

    search(word: string): boolean {
        return searchNode(this.root, word, 0);
    }
}

function searchNode(node: SolveWDNode, word: string, i: number): boolean {
    if (i === word.length) {
        return node.isEnd;
    }
    const c = word[i];
    if (c === ".") {
        for (const child of node.children.values()) {
            if (searchNode(child, word, i + 1)) {
                return true;
            }
        }
        return false;
    }
    const child = node.children.get(c);
    if (child === undefined) {
        return false;
    }
    return searchNode(child, word, i + 1);
}
