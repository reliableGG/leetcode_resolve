package main

import "fmt"

type Node struct {
	Val    int
	lchild *Node
	rchild *Node
}

func MaxDepthOfBTree(node *Node) int {
	if node == nil {
		return 0
	}
	if node.lchild == nil && node.rchild == nil {
		return 1
	}
	return Max(MaxDepthOfBTree(node.lchild), MaxDepthOfBTree(node.rchild)) + 1
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Println("vim-go")
	root := new(Node)
	l1 := new(Node)
	l2 := new(Node)
	l3 := new(Node)
	root.lchild = l1
	root.rchild = l2
	l1.lchild = l3
	dep := MaxDepthOfBTree(root)
	fmt.Println("depth is ", dep)

}
