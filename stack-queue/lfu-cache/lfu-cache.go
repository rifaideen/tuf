package main

import "math"

type LFUCache struct {
	capacity     int
	frequency    map[int]*LRUCache
	cache        map[int]*Node
	minFrequency int
}

func Constructor(capacity int) LFUCache {
	frequency := map[int]*LRUCache{
		1: NewLRUCache(),
	}

	return LFUCache{
		capacity:     capacity,
		frequency:    frequency,
		cache:        make(map[int]*Node),
		minFrequency: 1,
	}
}

func (t *LFUCache) Get(key int) int {
	if node, ok := t.cache[key]; ok {
		// remove from existing frequency
		t.frequency[node.frequency].remove(node)

		// increment the frequency and put it into updated frequency
		node.frequency++

		if cache, ok := t.frequency[node.frequency]; ok {
			cache.insert(node)
		} else {
			cache := NewLRUCache()
			cache.insert(node)

			t.frequency[node.frequency] = cache
		}

		// clear the old map if it's empty after evicting from older frequency
		if len(t.frequency[node.frequency-1].cache) == 0 {
			delete(t.frequency, node.frequency-1)

			t.updateMinFrequency()
		}

		return node.value
	}

	return -1
}

func (t *LFUCache) Put(key int, value int) {
	if node, ok := t.cache[key]; ok {
		// // remove from existing frequency
		// t.frequency[node.frequency].remove(node)

		// // update the value
		// node.value = value

		// // increment the frequency and put it into updated frequency
		// node.frequency++

		// if cache, ok := t.frequency[node.frequency]; ok {
		// 	cache.insert(node)
		// } else {
		// 	cache := NewLRUCache()
		// 	cache.insert(node)

		// 	t.frequency[node.frequency] = cache
		// }

		// // clear the old map if it's empty after evicting from older frequency
		// if len(t.frequency[node.frequency-1].cache) == 0 {
		// 	delete(t.frequency, node.frequency-1)

		// 	t.updateMinFrequency()
		// }
		node.value = value
		t.Get(node.key)

		return
	}

	if len(t.cache) == t.capacity {
		lf := t.frequency[t.minFrequency]
		victim := lf.tail.previous
		keyToDelete := victim.key
		lf.remove(victim)

		delete(t.cache, keyToDelete)
	}

	node := &Node{
		key:       key,
		value:     value,
		frequency: 1,
	}

	if _, ok := t.frequency[1]; !ok {
		t.frequency[1] = NewLRUCache()
	}

	t.minFrequency = 1
	t.frequency[1].insert(node)
	t.cache[key] = node
}

func (t *LFUCache) updateMinFrequency() {
	newMinFreq := math.MaxInt32

	for freq := range t.frequency {
		if freq < newMinFreq {
			newMinFreq = freq
		}
	}

	if newMinFreq == math.MaxInt32 {
		t.minFrequency = 0 // means no entries left
	} else {
		t.minFrequency = newMinFreq
	}
}

type Node struct {
	next      *Node
	previous  *Node
	key       int
	value     int
	frequency int
}

type LRUCache struct {
	cache map[int]*Node
	head  *Node
	tail  *Node
}

func NewLRUCache() *LRUCache {
	head := &Node{
		key:   -1,
		value: -1,
	}

	tail := &Node{
		previous: head,
		key:      -1,
		value:    -1,
	}

	head.next = tail

	return &LRUCache{
		cache: make(map[int]*Node),
		head:  head,
		tail:  tail,
	}
}

func (t *LRUCache) Get(key int) int {
	// cache present, remove from current place and insert it to front
	if node, ok := t.cache[key]; ok {
		t.remove(node)
		t.insert(node)

		return node.value
	}

	// cache does not found
	return -1
}

func (t *LRUCache) Put(key int, value int) {
	var node *Node
	var ok bool

	// key present in cache, replace the value and put it to front
	if node, ok = t.cache[key]; ok {
		node.value = value

		t.remove(node)
		t.insert(node)

		return
	}

	node = &Node{
		key:       key,
		value:     value,
		frequency: 1,
	}

	t.insert(node)
}

func (t *LRUCache) insert(node *Node) {
	// node's previous should be head
	node.previous = t.head
	// node's next should head's next
	node.next = t.head.next
	// update head's next node's previous to node
	t.head.next.previous = node
	// update the head's next to point node
	t.head.next = node

	// keep it cache
	t.cache[node.key] = node
}

func (t *LRUCache) remove(node *Node) {
	// previous node's next should points to node's next
	node.previous.next = node.next
	// next node's previous should points to node's previous
	node.next.previous = node.previous

	delete(t.cache, node.key)
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
