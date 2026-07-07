def solve_word_break(s: str, word_dict: list[str]) -> bool:
    """Check if s can be segmented into dictionary words.

    Time: O(n^2), Space: O(n)
    """
    word_set = set(word_dict)
    n = len(s)
    dp = [False] * (n + 1)
    dp[0] = True
    for i in range(1, n + 1):
        for j in range(i):
            if dp[j] and s[j:i] in word_set:
                dp[i] = True
                break
    return dp[n]
