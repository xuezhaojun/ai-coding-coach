export function solveFindDuplicate(nums: number[]): number {
    // Phase 1: Find intersection point
    let slow = nums[0];
    let fast = nums[nums[0]];
    while (slow !== fast) {
        slow = nums[slow];
        fast = nums[nums[fast]];
    }
    // Phase 2: Find entrance to the cycle
    slow = 0;
    while (slow !== fast) {
        slow = nums[slow];
        fast = nums[fast];
    }
    return slow;
}
