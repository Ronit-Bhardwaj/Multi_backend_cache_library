package cache

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type LRU struct {
	capacity int
	items    map[string]*list.Element
	evict    *list.List
	mutex    sync.Mutex
}

type data struct {
	key     string
	Value   interface{}
	expired time.Time
}

func Newlru(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		evict:    list.New(),
	}
}

func (lru *LRU) Set(key string, val interface{}, ttl time.Duration) error {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	if ele, ok := lru.items[key]; ok {
		ele.Value.(*data).Value = val
		lru.evict.MoveToFront(ele)
		ele.Value.(*data).expired = time.Now().Add(ttl)
		return nil
	} else {
		if lru.evict.Len() >= lru.capacity {
			last_ele := lru.evict.Back()
			if last_ele != nil {
				lru.evict.Remove(last_ele)
				delete(lru.items, last_ele.Value.(*data).key)
			}
		}
		data := &data{key, val, time.Now().Add(ttl)}
		ele := lru.evict.PushFront(data)
		lru.items[key] = ele
		return nil
	}
}

func (lru *LRU) Get(key string) (interface{}, error) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	if ele, ok := lru.items[key]; ok {
		if time.Now().Before(ele.Value.(*data).expired) {
			lru.evict.MoveToFront(ele)
			return ele.Value.(*data).Value, nil
		}
	}
	return nil, fmt.Errorf("cache miss")
}


func (lru *LRU) GetAllKeys() []string {
    lru.mutex.Lock()
    defer lru.mutex.Unlock()

    keys := make([]string, 0, len(lru.items))
    for key := range lru.items {
        keys = append(keys, key)
    }
    return keys
}


func (lru *LRU) Delete(key string) error {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	if ele, ok := lru.items[key]; ok {
		lru.evict.Remove(ele)
		delete(lru.items, key)
		return nil
	}
	return fmt.Errorf("cache miss")
}

func (lru *LRU) Clear() {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	lru.items = make(map[string]*list.Element)
	lru.evict.Init()
}