import { ListNode } from "./helpers";

export function solveRemoveNthFromEnd(head: ListNode | null, n: number): ListNode | null {
    const dummy = new ListNode(0, head);
    let fast: ListNode | null = dummy;
    let slow: ListNode | null = dummy;
    for (let i = 0; i <= n; i++) {
        fast = fast!.next;
    }
    while (fast !== null) {
        fast = fast.next;
        slow = slow!.next;
    }
    slow!.next = slow!.next!.next;
    return dummy.next;
}
