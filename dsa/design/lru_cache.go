package main

import (
	"fmt"
	"sync"
)

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
	mutex    sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) insert(node *Node) {
	node.next = l.head.next
	node.prev = l.head

	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) Get(key int) (int, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if node, found := l.cache[key]; found {
		l.remove(node)
		l.insert(node)
		return node.value, true
	}

	return -1, false
}

func (l *LRUCache) Put(key, value int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if node, found := l.cache[key]; found {
		l.remove(node)
		node.value = value
		l.insert(node)
		return
	}

	if len(l.cache) >= l.capacity {
		lru := l.tail.prev
		l.remove(lru)
		delete(l.cache, lru.key)
	}

	node := &Node{key: key, value: value}
	l.insert(node)
	l.cache[key] = node
}

func main() {
	cache := NewLRUCache(2)
	cache.Put(1, 100)
	cache.Put(2, 200)
	fmt.Println(cache.Get(1)) // 100, true
	cache.Put(3, 300)         // Evicts key 2
	fmt.Println(cache.Get(2)) // -1, false
	cache.Put(4, 400)         // Evicts key 1
	fmt.Println(cache.Get(1)) // -1, false
	fmt.Println(cache.Get(3)) // 300, true
	fmt.Println(cache.Get(4)) // 400, true

}
