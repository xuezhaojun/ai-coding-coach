from helpers import ListNode, Optional


def solve_reverse_k_group(head: Optional[ListNode], k: int) -> Optional[ListNode]:
    """Reverse nodes in k-group. Remaining nodes stay as-is.

    Time: O(n), Space: O(1).
    """
    dummy = ListNode(0, head)
    group_prev = dummy

    while True:
        # Check if there are k nodes remaining
        kth: Optional[ListNode] = group_prev
        for _ in range(k):
            kth = kth.next  # type: ignore[union-attr]
            if kth is None:
                return dummy.next
        group_next = kth.next  # type: ignore[union-attr]

        # Reverse the group
        prev: Optional[ListNode] = kth.next  # type: ignore[assignment]
        curr: Optional[ListNode] = group_prev.next
        while curr is not group_next:
            nxt = curr.next  # type: ignore[union-attr]
            curr.next = prev  # type: ignore[union-attr]
            prev = curr
            curr = nxt

        # Connect with previous part
        tmp = group_prev.next
        group_prev.next = kth  # type: ignore[assignment]
        group_prev = tmp  # type: ignore[assignment]
