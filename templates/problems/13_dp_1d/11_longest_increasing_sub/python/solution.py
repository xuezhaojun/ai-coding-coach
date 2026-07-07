import bisect


def solve_length_of_lis(nums: list[int]) -> int:
    """Return the length of the longest increasing subsequence.

    tails[i] is the smallest tail element for an increasing subsequence of length i+1.
    Time: O(n log n), Space: O(n)
    """
    tails: list[int] = []
    for num in nums:
        pos = bisect.bisect_left(tails, num)
        if pos == len(tails):
            tails.append(num)
        else:
            tails[pos] = num
    return len(tails)
