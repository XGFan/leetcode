package main

//
//type Node struct {
//	prev  *Node
//	key   int
//	value int
//	next  *Node
//}
//
//type LRUCache struct {
//	capacity int
//	nodeMap  map[int]*Node
//	head     *Node
//	tail     *Node
//}
//
//func Constructor(capacity int) LRUCache {
//	return LRUCache{
//		capacity: capacity,
//		nodeMap:  make(map[int]*Node),
//		head:     nil,
//		tail:     nil,
//	}
//}
//
//func (this *LRUCache) Get(key int) int {
//	node, exist := this.nodeMap[key]
//	if exist {
//		this.putFirst(node)
//		return node.value
//	} else {
//		return -1
//	}
//}
//
//func (this *LRUCache) putFirst(node *Node) {
//	if this.head == node {
//		return
//	}
//	node.prev.next = node.next
//	if this.tail == node {
//		this.tail = node.prev
//	} else {
//		node.next.prev = node.prev
//	}
//	node.prev = nil
//	node.next = this.head
//	this.head.prev = node
//	this.head = node
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	node, exist := this.nodeMap[key]
//	if !exist {
//		newHead := &Node{
//			prev:  nil,
//			key:   key,
//			value: value,
//			next:  this.head,
//		}
//		if newHead.next != nil {
//			newHead.next.prev = newHead
//		} else {
//			this.tail = newHead
//		}
//		this.head = newHead
//		if len(this.nodeMap) == this.capacity {
//			delete(this.nodeMap, this.tail.key)
//			newTail := this.tail.prev
//			this.tail.prev = nil
//			this.tail = newTail
//			this.tail.next = nil
//		}
//		this.nodeMap[key] = newHead
//	} else {
//		node.value = value
//		this.putFirst(node)
//	}
//}
//
///**
// * Your LRUCache object will be instantiated and called as such:
// * obj := Constructor(capacity);
// * param_1 := obj.Get(key);
// * obj.Put(key,value);
// */
//
//func main() {
//	lruCache := Constructor(2)
//	lruCache.Put(2, 1)
//	lruCache.Put(2, 2)
//	lruCache.Get(2)
//	lruCache.Put(3, 3)
//	lruCache.Get(2)
//	lruCache.Put(4, 4)
//	lruCache.Get(1)
//	lruCache.Get(3)
//	lruCache.Get(4)
//}
