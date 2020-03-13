package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Node struct {
	Key   int
	Val   int
	left  *Node
	right *Node
}

type LRUCache struct {
	head     *Node
	M        map[int]*Node
	length   int
	capacity int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		M:        make(map[int]*Node, capacity),
		capacity: capacity,
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, exist := this.M[key]; !exist {
		return -1
	} else {
		leftNode := node.left
		rightNode := node.right
		leftNode.right = rightNode
		rightNode.left = leftNode

		this.head.left = node
		node.right = this.head
		node.left = nil
		this.head = node
		return node.Val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, exist := this.M[key]; !exist {
		n := &Node{
			Key:   key,
			Value: value,
		}
		if this.length == 0 {
			this.head = n
		}

		if this.length+1 <= this.capacity {
			this.head.left = node
			n.right = this.head
			n.left = nil
			this.head = n
		}
		this.length++
		this.M[key] = n
	}

}
