import { ListNode } from "./helpers";

export function solveReorderList(head: ListNode | null): void {
    if (head === null || head.next === null) {
        return;
    }
    // Find the middle
    let slow = head;
    let fast: ListNode | null = head;
    while (fast!.next !== null && fast!.next.next !== null) {
        slow = slow.next!;
        fast = fast!.next.next;
    }
    // Reverse the second half
    let prev: ListNode | null = null;
    let curr: ListNode | null = slow.next;
    slow.next = null;
    while (curr !== null) {
        const next = curr.next;
        curr.next = prev;
        prev = curr;
        curr = next;
    }
    // Merge the two halves
    let first: ListNode | null = head;
    let second: ListNode | null = prev;
    while (second !== null) {
        const tmp1: ListNode | null = first!.next;
        const tmp2: ListNode | null = second.next;
        first!.next = second;
        second.next = tmp1;
        first = tmp1;
        second = tmp2;
    }
}
