from typing import Optional


class _DLNode:
    __slots__ = ("key", "val", "prev", "next")

    def __init__(self, key: int = 0, val: int = 0):
        self.key = key
        self.val = val
        self.prev: Optional['_DLNode'] = None
        self.next: Optional['_DLNode'] = None


class SolveLRUCache:
    """LRU cache with a doubly linked list and hash map. O(1) Get and Put."""

    def __init__(self, capacity: int):
        self.capacity = capacity
        self.cache: dict[int, _DLNode] = {}
        self.head = _DLNode()  # dummy head
        self.tail = _DLNode()  # dummy tail
        self.head.next = self.tail
        self.tail.prev = self.head

    def get(self, key: int) -> int:
        node = self.cache.get(key)
        if node is None:
            return -1
        self._move_to_front(node)
        return node.val

    def put(self, key: int, value: int) -> None:
        node = self.cache.get(key)
        if node is not None:
            node.val = value
            self._move_to_front(node)
            return
        node = _DLNode(key, value)
        self.cache[key] = node
        self._add_to_front(node)
        if len(self.cache) > self.capacity:
            removed = self._remove_last()
            if removed is not None:
                del self.cache[removed.key]

    def _add_to_front(self, node: _DLNode) -> None:
        node.prev = self.head
        node.next = self.head.next
        self.head.next.prev = node  # type: ignore[union-attr]
        self.head.next = node

    def _remove_node(self, node: _DLNode) -> None:
        if node.prev is not None:
            node.prev.next = node.next
        if node.next is not None:
            node.next.prev = node.prev

    def _move_to_front(self, node: _DLNode) -> None:
        self._remove_node(node)
        self._add_to_front(node)

    def _remove_last(self) -> Optional[_DLNode]:
        node = self.tail.prev
        if node is None or node is self.head:
            return None
        self._remove_node(node)
        return node
