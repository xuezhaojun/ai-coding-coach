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
