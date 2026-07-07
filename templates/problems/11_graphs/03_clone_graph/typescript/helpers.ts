export class Node {
    val: number;
    neighbors: Node[];

    constructor(val: number = 0, neighbors: Node[] = []) {
        this.val = val;
        this.neighbors = neighbors;
    }
}
