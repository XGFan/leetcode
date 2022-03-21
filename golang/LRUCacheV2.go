package main

import "container/list"

type Item struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	elemMap  map[int]*list.Element
	list     *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		elemMap:  make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, exist := this.elemMap[key]; exist {
		this.list.MoveToFront(elem)
		return elem.Value.(Item).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exist := this.elemMap[key]; exist {
		this.list.MoveToFront(node)
		node.Value = Item{key, value}
	} else {
		if len(this.elemMap) == this.capacity {
			delete(this.elemMap, this.list.Back().Value.(Item).key)
			this.list.Remove(this.list.Back())
		}
		this.list.PushFront(Item{
			key:   key,
			value: value,
		})
		this.elemMap[key] = this.list.Front()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lruCache := Constructor(2)
	lruCache.Put(2, 1)
	lruCache.Put(2, 2)
	lruCache.Get(2)
	lruCache.Put(3, 3)
	lruCache.Get(2)
	lruCache.Put(4, 4)
	lruCache.Get(1)
	lruCache.Get(3)
	lruCache.Get(4)
}
