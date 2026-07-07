export class ListNode {
    val: number;
    next: ListNode | null;
    constructor(val: number = 0, next: ListNode | null = null) {
        this.val = val;
        this.next = next;
    }
}

export function buildList(vals: number[]): ListNode | null {
    const dummy = new ListNode();
    let cur = dummy;
    for (const v of vals) {
        cur.next = new ListNode(v);
        cur = cur.next;
    }
    return dummy.next;
}

export function listToSlice(head: ListNode | null): number[] {
    const result: number[] = [];
    let cur = head;
    while (cur) {
        result.push(cur.val);
        cur = cur.next;
    }
    return result;
}
