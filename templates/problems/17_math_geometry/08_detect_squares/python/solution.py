class SolveDetectSquares:
    """Use a hash map of point counts. O(1) Add, O(n) Count per query."""

    def __init__(self) -> None:
        self.point_count: dict[tuple[int, int], int] = {}
        self.points: list[list[int]] = []

    def add(self, point: list[int]) -> None:
        key = (point[0], point[1])
        self.point_count[key] = self.point_count.get(key, 0) + 1
        self.points.append(point)

    def count(self, point: list[int]) -> int:
        px, py = point[0], point[1]
        total = 0
        for p in self.points:
            qx, qy = p[0], p[1]
            dx = qx - px
            dy = qy - py
            if abs(dx) != abs(dy) or dx == 0:
                continue
            total += self.point_count.get((px, qy), 0) * self.point_count.get((qx, py), 0)
        return total
