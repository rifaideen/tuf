package main

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	// -1 -1
	cache := Constructor(2)
	// -1 1 -1
	cache.Put(1, 1)

	// // -1 2 1 -1
	cache.Put(2, 2)

	fmt.Println(cache.String())

	cache.Put(1, 10)
	fmt.Println(cache.String())
	// // -1 1 2 -1
	// assert.Equal(t, 1, cache.Get(1))
	// // -1 3 1 -1
	// cache.Put(3, 3)
	// fmt.Println(cache.String())

	// assert.Equal(t, -1, cache.Get(2))

	// cache.Put(4, 4)
	// assert.Equal(t, -1, cache.Get(1))
	// assert.Equal(t, 3, cache.Get(3))
	// assert.Equal(t, 4, cache.Get(4))
}
