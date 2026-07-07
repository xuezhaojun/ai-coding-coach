export class TreeNode {
    val: number;
    left: TreeNode | null;
    right: TreeNode | null;
    constructor(val: number = 0, left: TreeNode | null = null, right: TreeNode | null = null) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}

export function buildTree(vals: number[]): TreeNode | null {
    if (vals.length === 0 || vals[0] === -101) {
        return null;
    }
    const root = new TreeNode(vals[0]);
    const queue: (TreeNode | null)[] = [root];
    let i = 1;
    while (queue.length > 0 && i < vals.length) {
        const node = queue.shift()!;
        if (i < vals.length && vals[i] !== -101) {
            node.left = new TreeNode(vals[i]);
            queue.push(node.left);
        }
        i++;
        if (i < vals.length && vals[i] !== -101) {
            node.right = new TreeNode(vals[i]);
            queue.push(node.right);
        }
        i++;
    }
    return root;
}

export function treeToSlice(root: TreeNode | null): number[] {
    if (root === null) {
        return [];
    }
    const result: number[] = [];
    const queue: (TreeNode | null)[] = [root];
    while (queue.length > 0) {
        const node = queue.shift()!;
        if (node === null) {
            result.push(-101);
        } else {
            result.push(node.val);
            queue.push(node.left);
            queue.push(node.right);
        }
    }
    while (result.length > 0 && result[result.length - 1] === -101) {
        result.pop();
    }
    return result;
}
