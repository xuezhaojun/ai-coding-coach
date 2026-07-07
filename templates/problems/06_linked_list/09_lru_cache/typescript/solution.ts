class DLNode {
    key: number;
    val: number;
    prev: DLNode | null;
    next: DLNode | null;
    constructor(key: number = 0, val: number = 0) {
        this.key = key;
        this.val = val;
        this.prev = null;
        this.next = null;
    }
}

export class SolveLRUCache {
    private capacity: number;
    private cache: Map<number, DLNode>;
    private head: DLNode;
    private tail: DLNode;

    constructor(capacity: number) {
        this.capacity = capacity;
        this.cache = new Map<number, DLNode>();
        this.head = new DLNode();
        this.tail = new DLNode();
        this.head.next = this.tail;
        this.tail.prev = this.head;
    }

    get(key: number): number {
        const node = this.cache.get(key);
        if (node === undefined) {
            return -1;
        }
        this.moveToFront(node);
        return node.val;
    }

    put(key: number, value: number): void {
        const node = this.cache.get(key);
        if (node !== undefined) {
            node.val = value;
            this.moveToFront(node);
            return;
        }
        const newNode = new DLNode(key, value);
        this.cache.set(key, newNode);
        this.addToFront(newNode);
        if (this.cache.size > this.capacity) {
            const removed = this.removeLast();
            if (removed !== null) {
                this.cache.delete(removed.key);
            }
        }
    }

    private addToFront(node: DLNode): void {
        node.prev = this.head;
        node.next = this.head.next;
        this.head.next!.prev = node;
        this.head.next = node;
    }

    private removeNode(node: DLNode): void {
        if (node.prev !== null) {
            node.prev.next = node.next;
        }
        if (node.next !== null) {
            node.next.prev = node.prev;
        }
    }

    private moveToFront(node: DLNode): void {
        this.removeNode(node);
        this.addToFront(node);
    }

    private removeLast(): DLNode | null {
        const node = this.tail.prev;
        if (node === null || node === this.head) {
            return null;
        }
        this.removeNode(node);
        return node;
    }
}
