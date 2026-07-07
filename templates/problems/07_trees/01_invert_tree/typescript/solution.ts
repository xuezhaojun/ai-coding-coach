import { TreeNode } from "./helpers";

export function solveInvertTree(root: TreeNode | null): TreeNode | null {
    if (root === null) {
        return null;
    }
    const temp = root.left;
    root.left = root.right;
    root.right = temp;
    solveInvertTree(root.left);
    solveInvertTree(root.right);
    return root;
}
