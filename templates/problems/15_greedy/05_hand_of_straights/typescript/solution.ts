/**
 * Use sorting and a frequency map. O(n log n) time, O(n) space.
 */
export function solveIsNStraightHand(hand: number[], groupSize: number): boolean {
    if (hand.length % groupSize !== 0) return false;
    const freq = new Map<number, number>();
    for (const card of hand) {
        freq.set(card, (freq.get(card) ?? 0) + 1);
    }
    hand = [...hand].sort((a, b) => a - b);
    for (const card of hand) {
        if ((freq.get(card) ?? 0) === 0) continue;
        for (let i = 0; i < groupSize; i++) {
            if ((freq.get(card + i) ?? 0) === 0) return false;
            freq.set(card + i, (freq.get(card + i) ?? 0) - 1);
        }
    }
    return true;
}
