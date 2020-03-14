package cacheS

import (
	"github.com/Alter/blog/pkg/util/cacheS/byteview"
	"github.com/Alter/blog/pkg/util/cacheS/lru3"
	"sync"
)


// 在 add 方法中，判断了 c.lru 是否为 nil，如果不等于 nil 再创建实例。
// 这种方法称之为延迟初始化(Lazy Initialization)，
// 一个对象的延迟初始化意味着该对象的创建将会延迟至第一次使用该对象时。
//主要用于提高性能，并减少程序内存要求。
type cache struct {
	mu         sync.Mutex
	lruCache   *lru3.Cache
	cacheBytes int64
}

func (c *cache) add(key string, value byteview.ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lruCache == nil {
		c.lruCache = lru3.New(c.cacheBytes, nil)
	}
	c.lruCache.Add(key, value)
}


func (c *cache) get(key string) (value byteview.ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lruCache == nil {
		return
	}

	if v, ok := c.lruCache.Get(key); ok {
		return v.(byteview.ByteView), true
	}
	return
}



type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}



