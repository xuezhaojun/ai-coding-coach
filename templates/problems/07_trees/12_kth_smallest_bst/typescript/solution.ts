import { TreeNode } from "./helpers";

export function solveKthSmallest(root: TreeNode | null, k: number): number {
    let count = 0;
    let result = 0;

    function inorder(node: TreeNode | null): void {
        if (node === null || count >= k) {
            return;
        }
        inorder(node.left);
        count++;
        if (count === k) {
            result = node.val;
            return;
        }
        inorder(node.right);
    }

    inorder(root);
    return result;
}
