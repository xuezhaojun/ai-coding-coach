package lru_cache

// SolveLRUCache implements an LRU cache with a doubly linked list and hash map. O(1) Get and Put.
type SolveLRUCache struct {
	capacity int
	cache    map[int]*dlNode
	head     *dlNode // dummy head
	tail     *dlNode // dummy tail
}

type dlNode struct {
	key, val   int
	prev, next *dlNode
}

func SolveConstructor(capacity int) SolveLRUCache {
	head := &dlNode{}
	tail := &dlNode{}
	head.next = tail
	tail.prev = head
	return SolveLRUCache{
		capacity: capacity,
		cache:    make(map[int]*dlNode),
		head:     head,
		tail:     tail,
	}
}

func (c *SolveLRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.moveToFront(node)
		return node.val
	}
	return -1
}

func (c *SolveLRUCache) Put(key int, value int) {
	if node, ok := c.cache[key]; ok {
		node.val = value
		c.moveToFront(node)
		return
	}
	node := &dlNode{key: key, val: value}
	c.cache[key] = node
	c.addToFront(node)
	if len(c.cache) > c.capacity {
		removed := c.removeLast()
		delete(c.cache, removed.key)
	}
}

func (c *SolveLRUCache) addToFront(node *dlNode) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *SolveLRUCache) removeNode(node *dlNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *SolveLRUCache) moveToFront(node *dlNode) {
	c.removeNode(node)
	c.addToFront(node)
}

func (c *SolveLRUCache) removeLast() *dlNode {
	node := c.tail.prev
	c.removeNode(node)
	return node
}
