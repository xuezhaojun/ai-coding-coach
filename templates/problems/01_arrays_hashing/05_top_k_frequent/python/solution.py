def solve_top_k_frequent(nums: list[int], k: int) -> list[int]:
    """Find k most frequent elements using bucket sort.

    Time: O(n), Space: O(n)
    """
    freq: dict[int, int] = {}
    for n in nums:
        freq[n] = freq.get(n, 0) + 1

    # Bucket sort: index = frequency, value = list of numbers with that frequency
    buckets: list[list[int]] = [[] for _ in range(len(nums) + 1)]
    for num, count in freq.items():
        buckets[count].append(num)

    result: list[int] = []
    for i in range(len(buckets) - 1, -1, -1):
        if len(result) >= k:
            break
        result.extend(buckets[i])
    return result[:k]
