import { ListNode } from "./helpers";

export function solveReverseKGroup(head: ListNode | null, k: number): ListNode | null {
    const dummy = new ListNode(0, head);
    let groupPrev: ListNode = dummy;

    while (true) {
        // Check if there are k nodes remaining
        let kth: ListNode | null = groupPrev;
        for (let i = 0; i < k; i++) {
            kth = kth!.next;
            if (kth === null) {
                return dummy.next;
            }
        }
        const groupNext: ListNode | null = kth!.next;

        // Reverse the group
        let prev: ListNode | null = kth!.next;
        let curr: ListNode | null = groupPrev.next;
        while (curr !== groupNext) {
            const next = curr!.next;
            curr!.next = prev;
            prev = curr;
            curr = next;
        }

        // Connect with previous part
        const tmp = groupPrev.next;
        groupPrev.next = kth!;
        groupPrev = tmp!;
    }
}
