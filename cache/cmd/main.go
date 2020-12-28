package main

import (
	"fmt"

	"github.com/HassankSalim/datastructure/cache"
)

func main() {
	capacity := 10
	c := cache.NewLRUCache(capacity)
	c.Set("Hello", "World")
	val := c.Get("Hello")
	fmt.Println(val)
}
