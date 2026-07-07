import { TreeNode } from "./helpers";

function validateBST(node: TreeNode | null, min: number, max: number): boolean {
    if (node === null) {
        return true;
    }
    if (node.val <= min || node.val >= max) {
        return false;
    }
    return validateBST(node.left, min, node.val) && validateBST(node.right, node.val, max);
}

export function solveIsValidBST(root: TreeNode | null): boolean {
    return validateBST(root, Number.MIN_SAFE_INTEGER, Number.MAX_SAFE_INTEGER);
}
