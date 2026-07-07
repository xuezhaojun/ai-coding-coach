import sys


def solve(input_data: str) -> str:
    """Read pairs (a, b) until both are 0. For each pair (except 0 0),
    return a+b on its own line.

    ACM-style: input ends with the sentinel line `0 0`.
    The sentinel pair is NOT processed/output.
    """
    # Parse input_data, processing pairs until encountering (0, 0).
    #
    # 适用题型：多组查询直到特定条件停止、模拟直到终止信号
    # 典型输入：若干行 a b，遇到 0 0 结束
    return ""


# For actual ACM usage:
# if __name__ == "__main__":
#     data = sys.stdin.read()
#     print(solve(data), end="")
