package main

import (
	"sort"
	"sync"
)

type LRUCache struct {
	Cap int
	M   map[int]int
	L   map[int]int
	sync.RWMutex
}

func Constructor(capacity int) *LRUCache {
	lrucache := &LRUCache{Cap: capacity,
		M: make(map[int]int{}, capacity),
		L: make(map[int]int{}, cap),
	}
	return lrucache
}

func (this *LRUCache) Get(key int) int {
	this.RLock()
	defer this.Unlock()

	if value, exist := this.M[key]; exist {
		this.L[key]++
		return value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	this.Lock()
	defer this.Unlock()

	if !this.reachMax() {
		this.M[key] = value
		this.L[key] = 0
	} else {
		this.evicts(key)
	}
}

func (this *LRUCache) reachMax() bool {
	return len(this.M) == this.Cap
}

func (this *LRUCache) evicts(key int) {
	container := make(map[int]int, this.Cap)
	for k, v := range this.M {
		container[v] = k
	}

	sort.Sort(container)

	delete(this.M, container[:1])
	delete(this.L, container[:1])
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Get(1)
}
