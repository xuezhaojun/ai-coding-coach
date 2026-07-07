from helpers import ListNode, Optional


def solve_add_two_numbers(l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
    """Add two numbers represented as reversed linked lists.

    Time: O(max(m,n)), Space: O(max(m,n)).
    """
    dummy = ListNode()
    curr = dummy
    carry = 0
    while l1 is not None or l2 is not None or carry > 0:
        total = carry
        if l1 is not None:
            total += l1.val
            l1 = l1.next
        if l2 is not None:
            total += l2.val
            l2 = l2.next
        carry = total // 10
        curr.next = ListNode(total % 10)
        curr = curr.next
    return dummy.next
