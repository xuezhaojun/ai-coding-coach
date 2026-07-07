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

export function buildCyclicList(vals: number[], cyclePos: number): ListNode | null {
    const head = buildList(vals);
    if (cyclePos < 0 || head === null) {
        return head;
    }
    let tail = head;
    while (tail.next !== null) {
        tail = tail.next;
    }
    let target: ListNode | null = head;
    for (let i = 0; i < cyclePos; i++) {
        target = target!.next;
    }
    tail.next = target;
    return head;
}
