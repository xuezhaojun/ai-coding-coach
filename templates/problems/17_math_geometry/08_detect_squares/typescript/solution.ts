/**
 * Use a hash map of point counts. O(1) Add, O(n) Count per query.
 */
export class SolveDetectSquares {
    private pointCount: Map<string, number> = new Map();
    private points: number[][] = [];

    add(point: number[]): void {
        const key = `${point[0]},${point[1]}`;
        this.pointCount.set(key, (this.pointCount.get(key) ?? 0) + 1);
        this.points.push(point);
    }

    count(point: number[]): number {
        const px = point[0], py = point[1];
        let total = 0;
        for (const p of this.points) {
            const qx = p[0], qy = p[1];
            const dx = qx - px;
            const dy = qy - py;
            if (Math.abs(dx) !== Math.abs(dy) || dx === 0) {
                continue;
            }
            total += (this.pointCount.get(`${px},${qy}`) ?? 0) * (this.pointCount.get(`${qx},${py}`) ?? 0);
        }
        return total;
    }
}
