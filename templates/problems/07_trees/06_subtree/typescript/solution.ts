import { TreeNode } from "./helpers";

function sameTree(p: TreeNode | null, q: TreeNode | null): boolean {
    if (p === null && q === null) {
        return true;
    }
    if (p === null || q === null || p.val !== q.val) {
        return false;
    }
    return sameTree(p.left, q.left) && sameTree(p.right, q.right);
}

export function solveIsSubtree(root: TreeNode | null, subRoot: TreeNode | null): boolean {
    if (root === null) {
        return subRoot === null;
    }
    if (sameTree(root, subRoot)) {
        return true;
    }
    return solveIsSubtree(root.left, subRoot) || solveIsSubtree(root.right, subRoot);
}
