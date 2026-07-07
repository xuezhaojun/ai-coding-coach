import { TreeNode } from "./helpers";

export function solveMaxDepth(root: TreeNode | null): number {
    if (root === null) {
        return 0;
    }
    const left = solveMaxDepth(root.left);
    const right = solveMaxDepth(root.right);
    return left > right ? left + 1 : right + 1;
}
