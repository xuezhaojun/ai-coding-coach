from helpers import ListNode, Optional


def solve_has_cycle(head: Optional[ListNode]) -> bool:
    """Detect a cycle using Floyd's tortoise and hare algorithm.

    Time: O(n), Space: O(1).
    """
    slow = head
    fast = head
    while fast is not None and fast.next is not None:
        slow = slow.next
        fast = fast.next.next
        if slow is fast:
            return True
    return False
