import { TreeNode } from "./helpers";

export function solveSerialize(root: TreeNode | null): string {
    const parts: string[] = [];
    function preorder(node: TreeNode | null): void {
        if (node === null) {
            parts.push("N");
            return;
        }
        parts.push(String(node.val));
        preorder(node.left);
        preorder(node.right);
    }
    preorder(root);
    return parts.join(",");
}

export function solveDeserialize(data: string): TreeNode | null {
    const tokens = data.split(",");
    let idx = 0;
    function build(): TreeNode | null {
        if (idx >= tokens.length || tokens[idx] === "N" || tokens[idx] === "") {
            idx++;
            return null;
        }
        const val = parseInt(tokens[idx], 10);
        idx++;
        const node = new TreeNode(val);
        node.left = build();
        node.right = build();
        return node;
    }
    return build();
}
