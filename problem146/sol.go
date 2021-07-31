package sol

import "container/list"

/*
	LeetCode Problem 146: LRUCache
	Level: Hard
	Description: Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
	Implement the LRUCache class:
	1. LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
	2. int get(int key) Return the value of the key if the key exists, otherwise return -1.
	3. void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key
	value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least
	recently used key.
	The functions get and put must each run in O(1) average time complexity.
	Link: https://leetcode.com/problems/lru-cache/
*/

// LRUCache store key-value pairs.
type LRUCache struct {
	capacity int
	meta     *list.List
	data     map[int]*list.Element
}

// Entry store key-value pair
type Entry struct {
	Key   int
	Value int
}

// Constructor return the LRUCache initialized with the capacity
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		meta:     list.New(),
		data:     make(map[int]*list.Element, capacity),
	}
}

// Get return the value of the key if the key exists otherwise return -1
func (l *LRUCache) Get(key int) int {
	ele, ok := l.data[key]
	if !ok {
		return -1
	}
	l.meta.MoveToFront(ele)
	entry := ele.Value.(*Entry)
	return entry.Value
}

// Put update the value of the key if the key exists. Otherwise, add the key value pair
// to the cache. If the number of keys exceed the capacity from this operation,
// evict the least recetly userd key.
func (l *LRUCache) Put(key int, value int) {
	ele, ok := l.data[key]
	if ok {
		l.meta.MoveToFront(ele)
		ele.Value.(*Entry).Value = value
		return
	}
	if l.meta.Len() >= l.capacity {
		lru := l.meta.Back()
		l.meta.Remove(lru)
		delete(l.data, lru.Value.(*Entry).Key)
	}
	entry := &Entry{Key: key, Value: value}
	newele := l.meta.PushFront(entry)
	l.data[key] = newele
}
