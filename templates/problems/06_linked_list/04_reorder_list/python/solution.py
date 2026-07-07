from helpers import ListNode, Optional


def solve_reorder_list(head: Optional[ListNode]) -> None:
    """Reorder the list L0->Ln->L1->Ln-1->... in place.

    Time: O(n), Space: O(1).
    """
    if head is None or head.next is None:
        return
    # Find the middle
    slow = head
    fast = head
    while fast.next is not None and fast.next.next is not None:
        slow = slow.next  # type: ignore[assignment]
        fast = fast.next.next  # type: ignore[assignment]
    # Reverse the second half
    prev: Optional[ListNode] = None
    curr = slow.next
    slow.next = None
    while curr is not None:
        nxt = curr.next
        curr.next = prev
        prev = curr
        curr = nxt
    # Merge the two halves
    first = head
    second = prev
    while second is not None:
        tmp1 = first.next
        tmp2 = second.next
        first.next = second
        second.next = tmp1
        first = tmp1  # type: ignore[assignment]
        second = tmp2
