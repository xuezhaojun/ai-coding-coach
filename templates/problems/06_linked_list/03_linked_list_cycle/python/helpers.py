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


def build_cyclic_list(vals: list[int], cycle_pos: int) -> Optional[ListNode]:
    """Build a linked list from vals and connect the tail to the node at index cycle_pos.

    A cycle_pos of -1 means no cycle.
    """
    head = build_list(vals)
    if cycle_pos < 0 or head is None:
        return head
    tail = head
    while tail.next is not None:
        tail = tail.next
    target = head
    for _ in range(cycle_pos):
        target = target.next  # type: ignore[assignment]
    tail.next = target
    return head
