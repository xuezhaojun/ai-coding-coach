import sys


def solve(input_data: str) -> str:
    """Read a directed graph with n nodes and m edges, output the out-degree
    of each node (1 to n), one per line.

    ACM-style: input format is:
        n m
        u1 v1
        u2 v2
        ...
        um vm

    Output is n lines: the out-degree of node i for i = 1..n.
    """
    # Parse n, m, then m edges. Count out-degree per node (1-indexed).
    #
    # 适用题型：图的度数/连通性、BFS/DFS遍历、拓扑排序、Dijkstra最短路
    # 典型输入：第一行 n m，接着 m 行每行 u v 表示一条边
    return ""


# For actual ACM usage:
# if __name__ == "__main__":
#     data = sys.stdin.read()
#     print(solve(data), end="")
