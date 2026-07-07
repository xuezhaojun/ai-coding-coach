import { TreeNode } from "./helpers";

function height(node: TreeNode | null): number {
    if (node === null) {
        return 0;
    }
    const left = height(node.left);
    if (left === -1) {
        return -1;
    }
    const right = height(node.right);
    if (right === -1) {
        return -1;
    }
    const diff = left - right;
    if (diff < -1 || diff > 1) {
        return -1;
    }
    return left > right ? left + 1 : right + 1;
}

export function solveIsBalanced(root: TreeNode | null): boolean {
    return height(root) !== -1;
}
