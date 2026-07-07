from helpers import ListNode, Optional


def solve_merge_two_lists(list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
    """Merge two sorted linked lists.

    Time: O(n+m), Space: O(1).
    """
    dummy = ListNode()
    curr = dummy
    while list1 is not None and list2 is not None:
        if list1.val <= list2.val:
            curr.next = list1
            list1 = list1.next
        else:
            curr.next = list2
            list2 = list2.next
        curr = curr.next
    if list1 is not None:
        curr.next = list1
    else:
        curr.next = list2
    return dummy.next
