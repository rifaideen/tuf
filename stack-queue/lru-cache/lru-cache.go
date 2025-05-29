package main

import (
	"fmt"
	"strings"
)

type Node struct {
	next     *Node
	previous *Node
	key      int
	value    int
}

type LRUCache struct {
	head     *Node
	tail     *Node
	capacity int
	cache    map[int]*Node
}

func (t LRUCache) String() string {
	nodes := strings.Builder{}

	node := t.head

	for node != nil {
		if node.next != nil {
			nodes.WriteString(fmt.Sprintf("%d <-> ", node.value))
		} else {
			nodes.WriteString(fmt.Sprintf("%d ", node.value))
		}

		node = node.next
	}

	return nodes.String()
}

func Constructor(capacity int) LRUCache {
	// create dummy head and tail and link them
	head := &Node{
		key:   -1,
		value: -1,
	}

	tail := &Node{
		key:      -2,
		value:    -2,
		previous: head,
	}

	head.next = tail

	return LRUCache{
		head:     head,
		tail:     tail,
		cache:    make(map[int]*Node),
		capacity: capacity,
	}
}

func (t *LRUCache) Get(key int) int {
	if node, ok := t.cache[key]; ok {
		t.remove(node)
		t.insert(node)

		return node.value
	}

	// node doesn't exists
	return -1
}

func (t *LRUCache) Put(key int, value int) {
	if node, ok := t.cache[key]; ok {
		node.value = value
		t.remove(node)
		t.insert(node)

		return
	}

	if len(t.cache) == t.capacity {
		t.remove(t.tail.previous)
	}

	node := &Node{
		key:   key,
		value: value,
	}

	t.insert(node)
}

func (t *LRUCache) remove(node *Node) {
	prev := node.previous
	next := node.next

	prev.next = next
	next.previous = prev

	delete(t.cache, node.key)
}

func (t *LRUCache) insert(node *Node) {
	t.cache[node.key] = node

	node.next = t.head.next
	node.next.previous = node
	t.head.next = node
	node.previous = t.head
}
