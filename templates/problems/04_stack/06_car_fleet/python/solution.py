def solve_car_fleet(target: int, position: list[int], speed: list[int]) -> int:
    """Count car fleets by sorting by position and using a stack of arrival times.

    Time: O(n log n), Space: O(n)
    """
    n = len(position)
    if n == 0:
        return 0
    cars = sorted(zip(position, speed), key=lambda c: c[0], reverse=True)
    fleets = 0
    last_time = 0.0
    for pos, sp in cars:
        time = (target - pos) / sp
        if time > last_time:
            fleets += 1
            last_time = time
    return fleets
