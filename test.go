package main

import "github.com/lewilewilewi/lru-cache/lewilru"

func main() {
    lru, _ := lewilru.New(5)

    lru.Set("a", 1)
    lru.Set("b", 2)
    lru.Set("c", 3)
    lru.Set("d", 4)
    lru.Set("e", 5)
    lru.Set("f", 6)
    lru.Set("g", 7)
    lru.Set("b", 2)
    lru.Set("g", 7)
    lru.Get("foo")
    lru.Set("foo", 100)
    lru.Get("foo")
    lru.Set("foo", 199)
    lru.Get("foo")
    lru.Set("e", 1)
}