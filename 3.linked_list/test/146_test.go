package test

import (
	"container/list"
	"fmt"
	"testing"
)

func Test_146(t *testing.T) {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	value := obj.Get(1)
	fmt.Printf("k:v = %d:%d\n", 1, value)

	obj.Put(3, 3)
	value = obj.Get(2)
	fmt.Printf("k:v = %d:%d\n", 2, value)

	obj.Put(4, 4)
	value = obj.Get(1)
	fmt.Printf("k:v = %d:%d\n", 1, value)
	value = obj.Get(3)
	fmt.Printf("k:v = %d:%d\n", 3, value)
	value = obj.Get(4)
	fmt.Printf("k:v = %d:%d\n", 4, value)
}

type entry struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	ll       *list.List
	cache    map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		ll:       list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if this.cache == nil {
		return -1
	}
	if ele, hit := this.cache[key]; hit {
		this.ll.MoveToFront(ele)
		return ele.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.cache == nil {
		this.cache = make(map[int]*list.Element)
		this.ll = list.New()
	}
	if ee, ok := this.cache[key]; ok {
		this.ll.MoveToFront(ee)
		ee.Value.(*entry).value = value
		return
	}
	ele := this.ll.PushFront(&entry{key, value})
	this.cache[key] = ele
	if this.capacity != 0 && this.ll.Len() > this.capacity {
		oldest := this.ll.Back()
		if oldest != nil {
			this.ll.Remove(oldest)
			kv := oldest.Value.(*entry)
			delete(this.cache, kv.key)
		}
	}
}
