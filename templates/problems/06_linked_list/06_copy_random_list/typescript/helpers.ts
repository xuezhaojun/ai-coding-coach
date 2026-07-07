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

export class RandomNode {
    val: number;
    next: RandomNode | null;
    random: RandomNode | null;
    constructor(val: number = 0, next: RandomNode | null = null, random: RandomNode | null = null) {
        this.val = val;
        this.next = next;
        this.random = random;
    }
}

export function buildRandomList(vals: number[], randomIndices: number[]): RandomNode | null {
    if (vals.length === 0) {
        return null;
    }
    const nodes: RandomNode[] = vals.map((v) => new RandomNode(v));
    for (let i = 0; i < nodes.length - 1; i++) {
        nodes[i].next = nodes[i + 1];
    }
    for (let i = 0; i < randomIndices.length; i++) {
        const ri = randomIndices[i];
        if (ri >= 0) {
            nodes[i].random = nodes[ri];
        }
    }
    return nodes[0];
}

export function randomListToSlice(head: RandomNode | null): { vals: number[]; randoms: number[] } {
    const indexMap = new Map<RandomNode, number>();
    let idx = 0;
    let cur = head;
    while (cur) {
        indexMap.set(cur, idx);
        idx++;
        cur = cur.next;
    }
    const vals: number[] = [];
    const randoms: number[] = [];
    cur = head;
    while (cur) {
        vals.push(cur.val);
        if (cur.random !== null) {
            randoms.push(indexMap.get(cur.random)!);
        } else {
            randoms.push(-1);
        }
        cur = cur.next;
    }
    return { vals, randoms };
}
