from __future__ import annotations
from typing import Optional


class ListNode:
    def __init__(self, val: int = 0, next: Optional['ListNode'] = None):
        self.val = val
        self.next = next


def build_list(vals: list[int]) -> Optional[ListNode]:
    dummy = ListNode()
    cur = dummy
    for v in vals:
        cur.next = ListNode(v)
        cur = cur.next
    return dummy.next


def list_to_slice(head: Optional[ListNode]) -> list[int]:
    result: list[int] = []
    cur = head
    while cur:
        result.append(cur.val)
        cur = cur.next
    return result


class RandomNode:
    def __init__(self, val: int = 0, next: Optional['RandomNode'] = None, random: Optional['RandomNode'] = None):
        self.val = val
        self.next = next
        self.random = random


def build_random_list(vals: list[int], random_indices: list[int]) -> Optional[RandomNode]:
    """Build a RandomNode list from vals and random_indices.

    random_indices[i] is the index that node i's Random points to, or -1 for None.
    """
    if not vals:
        return None
    nodes = [RandomNode(val=v) for v in vals]
    for i in range(len(nodes) - 1):
        nodes[i].next = nodes[i + 1]
    for i, ri in enumerate(random_indices):
        if ri >= 0:
            nodes[i].random = nodes[ri]
    return nodes[0]


def random_list_to_slice(head: Optional[RandomNode]) -> tuple[list[int], list[int]]:
    """Return (vals, random_indices) for the list. random index is -1 for None."""
    index_map: dict[int, int] = {}
    idx = 0
    cur = head
    while cur:
        index_map[id(cur)] = idx
        idx += 1
        cur = cur.next
    vals: list[int] = []
    randoms: list[int] = []
    cur = head
    while cur:
        vals.append(cur.val)
        if cur.random is not None:
            randoms.append(index_map[id(cur.random)])
        else:
            randoms.append(-1)
        cur = cur.next
    return vals, randoms
