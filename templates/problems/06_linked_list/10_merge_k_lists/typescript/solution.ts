import { ListNode } from "./helpers";

class MinHeap {
    private data: ListNode[] = [];

    get length(): number {
        return this.data.length;
    }

    push(node: ListNode): void {
        this.data.push(node);
        this.bubbleUp(this.data.length - 1);
    }

    pop(): ListNode | undefined {
        if (this.data.length === 0) {
            return undefined;
        }
        const top = this.data[0];
        const last = this.data.pop()!;
        if (this.data.length > 0) {
            this.data[0] = last;
            this.bubbleDown(0);
        }
        return top;
    }

    private bubbleUp(i: number): void {
        while (i > 0) {
            const parent = (i - 1) >> 1;
            if (this.data[i].val < this.data[parent].val) {
                [this.data[i], this.data[parent]] = [this.data[parent], this.data[i]];
                i = parent;
            } else {
                break;
            }
        }
    }

    private bubbleDown(i: number): void {
        const n = this.data.length;
        while (true) {
            let smallest = i;
            const left = 2 * i + 1;
            const right = 2 * i + 2;
            if (left < n && this.data[left].val < this.data[smallest].val) {
                smallest = left;
            }
            if (right < n && this.data[right].val < this.data[smallest].val) {
                smallest = right;
            }
            if (smallest === i) {
                break;
            }
            [this.data[i], this.data[smallest]] = [this.data[smallest], this.data[i]];
            i = smallest;
        }
    }
}

export function solveMergeKLists(lists: (ListNode | null)[]): ListNode | null {
    const heap = new MinHeap();
    for (const l of lists) {
        if (l !== null) {
            heap.push(l);
        }
    }
    const dummy = new ListNode();
    let curr = dummy;
    while (heap.length > 0) {
        const node = heap.pop()!;
        curr.next = node;
        curr = curr.next;
        if (node.next !== null) {
            heap.push(node.next);
        }
    }
    return dummy.next;
}
