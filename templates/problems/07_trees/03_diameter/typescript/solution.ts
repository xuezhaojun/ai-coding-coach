import { TreeNode } from "./helpers";

export function solveDiameterOfBinaryTree(root: TreeNode | null): number {
    let result = 0;

    function dfs(node: TreeNode | null): number {
        if (node === null) {
            return 0;
        }
        const left = dfs(node.left);
        const right = dfs(node.right);
        if (left + right > result) {
            result = left + right;
        }
        return left > right ? left + 1 : right + 1;
    }

    dfs(root);
    return result;
}
