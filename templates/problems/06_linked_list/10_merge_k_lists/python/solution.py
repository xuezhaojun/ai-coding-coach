import heapq

from helpers import ListNode, Optional


def solve_merge_k_lists(lists: list[Optional[ListNode]]) -> Optional[ListNode]:
    """Merge k sorted lists using a min-heap.

    Time: O(N log k), Space: O(k).
    """
    heap: list[tuple[int, int, ListNode]] = []
    counter = 0
    for i, l in enumerate(lists):
        if l is not None:
            heapq.heappush(heap, (l.val, counter, l))
            counter += 1
    dummy = ListNode()
    curr = dummy
    while heap:
        _, _, node = heapq.heappop(heap)
        curr.next = node
        curr = curr.next
        if node.next is not None:
            heapq.heappush(heap, (node.next.val, counter, node.next))
            counter += 1
    return dummy.next
