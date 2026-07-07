def solve_can_complete_circuit(gas: list[int], cost: list[int]) -> int:
    """Find the starting gas station for a circular route. O(n) time, O(1) space."""
    total_surplus = 0
    current_surplus = 0
    start = 0
    for i in range(len(gas)):
        total_surplus += gas[i] - cost[i]
        current_surplus += gas[i] - cost[i]
        if current_surplus < 0:
            start = i + 1
            current_surplus = 0
    if total_surplus < 0:
        return -1
    return start
