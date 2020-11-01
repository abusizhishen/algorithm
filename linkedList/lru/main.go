package main

import (
	"fmt"
	"log"
)

type Node struct {
	Key,Value int
	Pre,Next *Node
}

type LRUCache struct {
	m map[int]*Node
	cap,length int
	head,tail *Node
}

func construct(n int) *LRUCache {
	head := &Node{
		Key:   0,
		Value: 0,
		Pre:   nil,
		Next:  nil,
	}

	tail := &Node{
		Key:   0,
		Value: 0,
		Pre:   nil,
		Next:  nil,
	}

	head.Next = tail
	tail.Pre = head

	return &LRUCache{
		m: map[int]*Node{},
		cap:    n,
		length: 0,
		head:   head,
		tail:   tail,
	}
}

func (lru *LRUCache) Put(k,v int) {
	fmt.Printf("add k:%d v:%d\n", k,v)

	node,ok := lru.m[k]
	if ok{
		node.Value = v
		lru.moveToHead(node)
		return
	}

	node = &Node{
		Key:   k,
		Value: v,
		Pre:   nil,
		Next:  nil,
	}

	if lru.length == lru.cap{
		fmt.Println("need remove")
		lru.removeTail()
		lru.length--
	}
	lru.addToHead(node)
	lru.length++
	lru.m[k] = node
}

func (lru *LRUCache)Get(k int)int{
	if node,ok := lru.m[k];ok{
		lru.moveToHead(node)
		return node.Value
	}
	return -1
}

func (lru *LRUCache)addToHead(node *Node){
	head := lru.head.Next
	head.Pre = node
	node.Next = head
	node.Pre = lru.head
	lru.head.Next = node

}

func (lru *LRUCache)moveToHead(node *Node)  {
	lru.remove(node)
	lru.addToHead(node)
}

func (lru *LRUCache)remove(node *Node){
	node.Next.Pre,node.Pre.Next = node.Pre,node.Next
}
func (lru *LRUCache)removeTail(){
	node := lru.tail.Pre
	log.Printf("remove node %v", node)
	lru.remove(node)
	delete(lru.m,node.Key)
}

func main()  {
	var cache = construct(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}

func printLinkedList(cache LRUCache)  {
	var node = cache.head
	for node != nil{
		log.Printf("k:%d v:%d", node.Key,node.Value)
		fmt.Println()
	}
}

