def solve_group_anagrams(strs: list[str]) -> list[list[str]]:
    """Group strings that are anagrams using sorted-key hashing.

    Time: O(n * k log k) where k is max string length, Space: O(n * k)
    """
    groups: dict[tuple[int, ...], list[str]] = {}
    for s in strs:
        key = [0] * 26
        for ch in s:
            key[ord(ch) - ord('a')] += 1
        key_tuple = tuple(key)
        if key_tuple not in groups:
            groups[key_tuple] = []
        groups[key_tuple].append(s)
    return list(groups.values())
