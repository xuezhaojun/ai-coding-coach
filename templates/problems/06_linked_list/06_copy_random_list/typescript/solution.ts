import { RandomNode } from "./helpers";

export function solveCopyRandomList(head: RandomNode | null): RandomNode | null {
    if (head === null) {
        return null;
    }
    const oldToNew = new Map<RandomNode, RandomNode>();
    let curr: RandomNode | null = head;
    while (curr !== null) {
        oldToNew.set(curr, new RandomNode(curr.val));
        curr = curr.next;
    }
    curr = head;
    while (curr !== null) {
        const node = oldToNew.get(curr)!;
        if (curr.next !== null) {
            node.next = oldToNew.get(curr.next)!;
        }
        if (curr.random !== null) {
            node.random = oldToNew.get(curr.random)!;
        }
        curr = curr.next;
    }
    return oldToNew.get(head)!;
}
