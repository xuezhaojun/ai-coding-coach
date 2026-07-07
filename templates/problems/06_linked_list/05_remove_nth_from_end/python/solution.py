from helpers import ListNode, Optional


def solve_remove_nth_from_end(head: Optional[ListNode], n: int) -> Optional[ListNode]:
    """Remove the nth node from the end using two pointers.

    Time: O(n), Space: O(1).
    """
    dummy = ListNode(0, head)
    fast: Optional[ListNode] = dummy
    slow: Optional[ListNode] = dummy
    for _ in range(n + 1):
        fast = fast.next  # type: ignore[union-attr]
    while fast is not None:
        fast = fast.next
        slow = slow.next  # type: ignore[union-attr]
    slow.next = slow.next.next  # type: ignore[union-attr]
    return dummy.next
