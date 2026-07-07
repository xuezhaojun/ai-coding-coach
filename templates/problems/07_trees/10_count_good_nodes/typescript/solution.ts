import { TreeNode } from "./helpers";

export function solveGoodNodes(root: TreeNode | null): number {
    function dfs(node: TreeNode | null, maxSoFar: number): number {
        if (node === null) {
            return 0;
        }
        let count = 0;
        if (node.val >= maxSoFar) {
            count = 1;
            maxSoFar = node.val;
        }
        return count + dfs(node.left, maxSoFar) + dfs(node.right, maxSoFar);
    }
    return dfs(root, root!.val);
}
