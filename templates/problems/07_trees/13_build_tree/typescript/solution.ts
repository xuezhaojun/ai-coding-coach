import { TreeNode } from "./helpers";

export function solveBuildTreeFromPreIn(preorder: number[], inorder: number[]): TreeNode | null {
    const inMap = new Map<number, number>();
    for (let i = 0; i < inorder.length; i++) {
        inMap.set(inorder[i], i);
    }
    let idx = 0;

    function build(lo: number, hi: number): TreeNode | null {
        if (lo > hi) {
            return null;
        }
        const rootVal = preorder[idx];
        idx++;
        const node = new TreeNode(rootVal);
        const mid = inMap.get(rootVal)!;
        node.left = build(lo, mid - 1);
        node.right = build(mid + 1, hi);
        return node;
    }

    return build(0, inorder.length - 1);
}
