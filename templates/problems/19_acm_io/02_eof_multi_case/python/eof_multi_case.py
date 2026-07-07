import sys


def solve(input_data: str) -> str:
    """Read pairs (a, b) until EOF. For each pair, return a+b on its own line.

    ACM-style: input is multiple lines, each with two integers.
    Output is one line per pair: a + b.
    """
    # Parse input_data line by line, summing pairs until all input is consumed.
    # Use split() to tokenize, then process tokens two at a time.
    #
    # 适用题型：多组数据求和/差/积、多轮模拟、连续查询处理
    # 典型输入：若干行 a b，读到文件末尾为止
    return ""


# For actual ACM usage:
# if __name__ == "__main__":
#     data = sys.stdin.read()
#     print(solve(data), end="")
