// LRUNode represents a single entry in the LRU Cache, holding both a key and a value.
// It includes pointers to the previous and next nodes, to allow efficient reordering of the cache.
type LRUNode struct {
	Prev, Next *LRUNode // references to the nodes preceding and succeeding this one in the cache.
	Key, Val   int      // The key used to look up this node, and this node's value.
}

// The LRUCache structure holds all the nodes in our cache, the location of most (Head) and
// least (Tail) recently used items, has a map to nodes for quick access, and a capacity limit.
type LRUCache struct {
	head, tail *LRUNode         // dummy nodes to simplify operations
	m          map[int]*LRUNode // map to allow O(1) retrieval of nodes by key
	c          int              // The capacity limit of the cache
}

// Constructor creates a new LRUCache, with a provided cache size, also initiates head and tail nodes.
func Constructor(capacity int) LRUCache {
	head := new(LRUNode) // Create new dummy head node
	tail := new(LRUNode) // Create new dummy tail node
	head.Next = tail     // Set head's next to tail
	tail.Prev = head     // Set tail's previous to head
	return LRUCache{
		head: head,
		tail: tail,
		m:    make(map[int]*LRUNode), // Empty map of Nodes
		c:    capacity,
	}
}

// Get retrieves the value of key from cache and moves this node to the head because it's recently used.
// If the key doesn't exist, -1 will be returned.
func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.moveToFront(node) // Move accessed node to the front
		return node.Val        // Return node's value
	}
	return -1
}

// Put adds a new key-value node to the cache. If the key exists, update the value and move it to the head.
// When the cache reaches its capacity, it should invalidate/remove the least recently used item before inserting.
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.Val = value       // Update node's value
		this.moveToFront(node) // Move accessed/modified node to the front
		return
	}
	if len(this.m) == this.c {
		tailKey := this.removeTail() // Remove LRU node when cache size is at limit
		delete(this.m, tailKey)
	}
	node := &LRUNode{ // Create new node
		Key: key,
		Val: value,
	}
	this.insertToFront(node) // Insert new node to front
	this.m[node.Key] = node  // Store reference to node in map
}

// removeTail removes the least recently used node from the double linked list and from the map, returning its key.
func (this *LRUCache) removeTail() int {
	tail := this.tail.Prev
	tail.Prev.Next = this.tail
	this.tail.Prev = tail.Prev
	return tail.Key
}

// moveToFront moves the specified node to the head of the double linked list (most recently used).
func (this *LRUCache) moveToFront(node *LRUNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	nxt := this.head.Next
	node.Next = nxt
	this.head.Next = node
	node.Prev = this.head
	nxt.Prev = node
}

// insertToFront inserts the specified node to the head of the double linked list (most recently used).
func (this *LRUCache) insertToFront(node *LRUNode) {
	nxt := this.head.Next
	node.Next = nxt
	this.head.Next = node
	node.Prev = this.head
	nxt.Prev = node
}

// keyValue represents a single entry in the LRU Cache, holding both a key and a value.
type keyValue struct {
	key, val int
}

// The LRUCache structure holds all the keyValues in our cache using a doubly linked list and a map.
// It also has a capacity limit.
type LRUCache struct {
	l *list.List            // list holds keyValues in order of recency
	m map[int]*list.Element // map to allow O(1) retrieval of list elements by key
	c int                   // The capacity limit of the cache
}

// Constructor creates a new LRUCache, with a provided cache size
func Constructor(capacity int) LRUCache {
	return LRUCache{
		l: list.New(),                  // Create new empty List
		m: make(map[int]*list.Element), // Create new empty Map of list elements
		c: capacity,
	}
}

// Get retrieves the value of a specified key from cache and moves this keyValue
// to the front of the list because it's recently used. If the key doesn't exist, -1 will be returned.
func (this *LRUCache) Get(key int) int {
	if e, ok := this.m[key]; ok {
		this.l.MoveToFront(e)          // Move accessed list element to front
		return e.Value.(*keyValue).val // Return the value of the keyValue
	}
	return -1
}

// Put adds a new key-value to the cache. If the key exists, update the value and move it to the front.
// If the cache is full, it removes the least recently used (LRU) item before inserting the new keyValue.
func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.m[key]; ok {
		this.l.MoveToFront(e)           // Move existing list element to front
		e.Value.(*keyValue).val = value // Update value of existing keyValue
		return
	}
	if this.l.Len() == this.c {
		delete(this.m, this.l.Back().Value.(*keyValue).key) // Delete the least recently used keyValue from the map
		this.l.Remove(this.l.Back())                        // Remove the least recently used keyValue from the list
	}
	e := this.l.PushFront(&keyValue{key: key, val: value}) // Insert new keyValue at front of list
	this.m[key] = e                                        // Store reference to list element in map
}
