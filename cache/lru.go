package cache

import "sync"

//can also intialise with timestamp
type CacheItem struct {
	key   string
	value interface{}
	next  *CacheItem
	prev  *CacheItem
}
//and with duration. So that after a certain time we done get outdated data
type LRU struct {
	capacity int
	cache    map[string]*CacheItem
	head     *CacheItem
	tail     *CacheItem
	mutex    sync.Mutex
}

func NewLRUCache(capacity int) *LRU {
	head := &CacheItem{}
	tail := &CacheItem{}
	head.next = tail
	tail.prev = head

	return &LRU{
		capacity: capacity,
		cache:    make(map[string]*CacheItem),
		head:     head,
		tail:     tail,
	}

}

func (c *LRU) addFront(node *CacheItem) {
	node.next = c.head.next
	node.prev = c.head
	c.head.prev = node
	c.head.next = node
}

func (c *LRU) remove(node *CacheItem) {
	node.prev.next = node.next
	node.next.prev = node.prev

}

func (c *LRU) Get(key string) (interface{}, bool) {

	if item, ok := c.cache[key]; ok {
		c.remove(item)
		c.addFront(item)
		return item.value, true
	}

	return nil, false

}

func (c *LRU) Set(key string, value interface{}) {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, ok := c.cache[key]; ok {
		item.value = value
		c.remove(item)
		c.addFront(item)
		return
	}

	if len(c.cache) >= c.capacity {
		nodeRemoved := c.tail.prev
		c.remove(nodeRemoved)
		delete(c.cache, nodeRemoved.key)
	}

	newItem := &CacheItem{
		key:   key,
		value: value,
	}

	c.cache[key] = newItem
	c.addFront(newItem)

}