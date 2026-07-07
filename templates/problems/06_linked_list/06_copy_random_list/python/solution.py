from helpers import RandomNode, Optional


def solve_copy_random_list(head: Optional[RandomNode]) -> Optional[RandomNode]:
    """Create a deep copy using a hash map.

    Time: O(n), Space: O(n).
    """
    if head is None:
        return None
    old_to_new: dict[int, RandomNode] = {}
    curr = head
    while curr is not None:
        old_to_new[id(curr)] = RandomNode(val=curr.val)
        curr = curr.next
    curr = head
    while curr is not None:
        node = old_to_new[id(curr)]
        if curr.next is not None:
            node.next = old_to_new[id(curr.next)]
        if curr.random is not None:
            node.random = old_to_new[id(curr.random)]
        curr = curr.next
    return old_to_new[id(head)]
