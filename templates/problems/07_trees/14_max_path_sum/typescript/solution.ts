import { TreeNode } from "./helpers";

export function solveMaxPathSum(root: TreeNode | null): number {
    let result = Number.MIN_SAFE_INTEGER;

    function dfs(node: TreeNode | null): number {
        if (node === null) {
            return 0;
        }
        let left = dfs(node.left);
        if (left < 0) {
            left = 0;
        }
        let right = dfs(node.right);
        if (right < 0) {
            right = 0;
        }
        const pathSum = node.val + left + right;
        if (pathSum > result) {
            result = pathSum;
        }
        return node.val + (left > right ? left : right);
    }

    dfs(root);
    return result;
}
