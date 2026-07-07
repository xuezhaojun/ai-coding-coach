/**
 * Add two integers using bit manipulation. O(1) time (bounded iterations), O(1) space.
 * JavaScript bitwise operators work on 32-bit signed integers, so the carry
 * naturally shifts out after 32 bits.
 */
export function solveGetSum(a: number, b: number): number {
    let aa = a, bb = b;
    while (bb !== 0) {
        const carry = aa & bb;
        aa = aa ^ bb;
        bb = carry << 1;
    }
    return aa;
}
