package factory

import (
	"github.com/pkg/errors"
	"sync"
)

/*
属于：创建型模式
优点：保持了简单工厂模式的优点，而且克服了它的缺点。
缺点：在添加新产品时，在一定程度上增加了系统的复杂度。
适合：客户端不需要知道具体产品类的类名，只需要知道所对应的工厂即可。
*/

// 定义一个抽象的cache工厂
type CacheFactory interface {
	Create() Cache
}

// 实现具体的工厂：redis工厂
type RedisCacheFactory struct {
}

func (rf RedisCacheFactory) Create() Cache {
	return &RedisCache{
		data: make(map[string]interface{}),
		mux:  sync.RWMutex{},
	}
}

// 实现具体的工厂：mem工厂
type MemCacheFactory struct {
}

func (mf MemCacheFactory) Create() Cache {
	return &Memcache{
		data: make(map[string]interface{}),
		mux:  sync.RWMutex{},
	}
}

type Cache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, error)
}

type RedisCache struct {
	data map[string]interface{}
	mux  sync.RWMutex
}

func (r *RedisCache) Set(key string, value interface{}) {
	r.mux.Lock()
	defer r.mux.Unlock()
	r.data[key] = value
	return
}

func (r *RedisCache) Get(key string) (interface{}, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()
	if v, ok := r.data[key]; ok {
		return v, nil
	}
	return nil, errors.New("key is not found")
}

// ------------------------------------------------------------------------
// 实现具体的Cache: Memcache
// ------------------------------------------------------------------------
type Memcache struct {
	data map[string]interface{}
	mux  sync.RWMutex
}

func (m *Memcache) Set(key string, value interface{}) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.data[key] = value
	return
}

func (m *Memcache) Get(key string) (interface{}, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	if v, ok := m.data[key]; ok {
		return v, nil
	}
	return nil, errors.New("key is not found")
}
