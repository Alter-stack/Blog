
package lru3

import "container/list"

// Cache is a LRU cache. It is not safe for concurrent access.
type Cache struct {
	maxBytes int64
	nbytes   int64
	ll       *list.List
	cache    map[string]*list.Element
	// optional and executed when an entry is purged.
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool){
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return
}
// RemoveOldest removes the oldest item
func (c *Cache) RemoveOldest() {
	back := c.ll.Back()
	if back != nil {
		kv := back.Value.(*entry)
		delete(c.cache, kv.key)
		c.ll.Remove(back)
		c.nbytes -= int64(kv.value.Len()) + int64(len(kv.key))
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add item / modify item
func (c *Cache) Add(key string, value Value)  {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
		return
	}
	newEntry := &entry{key:key, value:value}
	newEle := c.ll.PushFront(newEntry)
	c.cache[key] = newEle
	c.nbytes += int64(len(key)) + int64(value.Len())
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
