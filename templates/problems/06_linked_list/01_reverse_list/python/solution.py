from helpers import ListNode, Optional


def solve_reverse_list(head: Optional[ListNode]) -> Optional[ListNode]:
    """Reverse a singly linked list iteratively.

    Time: O(n), Space: O(1).
    """
    prev: Optional[ListNode] = None
    curr = head
    while curr is not None:
        nxt = curr.next
        curr.next = prev
        prev = curr
        curr = nxt
    return prev
