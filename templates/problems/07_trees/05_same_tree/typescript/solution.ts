import { TreeNode } from "./helpers";

export function solveIsSameTree(p: TreeNode | null, q: TreeNode | null): boolean {
    if (p === null && q === null) {
        return true;
    }
    if (p === null || q === null || p.val !== q.val) {
        return false;
    }
    return solveIsSameTree(p.left, q.left) && solveIsSameTree(p.right, q.right);
}
