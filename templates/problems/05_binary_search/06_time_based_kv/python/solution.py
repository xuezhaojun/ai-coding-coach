import bisect


class SolveTimeMap:
    """Time-based key-value store.

    Set: O(1), Get: O(log n). Space: O(n).
    """

    def __init__(self) -> None:
        self._store: dict[str, list[tuple[int, str]]] = {}

    def set(self, key: str, value: str, timestamp: int) -> None:
        self._store.setdefault(key, []).append((timestamp, value))

    def get(self, key: str, timestamp: int) -> str:
        entries = self._store.get(key, [])
        if not entries:
            return ""
        # Binary search for the largest timestamp <= given timestamp.
        # bisect_right gives the index after the last timestamp <= target.
        idx = bisect.bisect_right(entries, (timestamp, chr(0x10FFFF)))
        if idx == 0:
            return ""
        return entries[idx - 1][1]
