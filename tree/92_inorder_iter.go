package main

import (
	"errors"
	"fmt"
	"sync"
)

type Node struct {
	Val   int
	left  *Node
	right *Node
}

func InorderTraversal(root *Node) []int {
	result := []int{}
	if root == nil {
		fmt.Printf("%v\n", result)
		return nil
	}

	result = append(result, InorderTraversal(root.left)...)
	result = append(result, root.Val)
	result = append(result, InorderTraversal(root.right)...)

	return result
}

type TreeNode struct {
	Val   int
	left  *TreeNode
	right *TreeNode
}

type Stack struct {
	sync.Mutex
	nodes []*TreeNode
}

func Newstack() *Stack {
	return &Stack{sync.Mutex{}, []*TreeNode{}}
}

func (s *Stack) Push(node *TreeNode) {
	s.Lock()
	defer s.Unlock()

	s.nodes = append(s.nodes, node)
}

func (s *Stack) Pop() (*TreeNode, error) {
	s.Lock()
	defer s.Unlock()

	l := len(s.nodes)
	if l == 0 {
		return nil, errors.New("Empty Stack")
	}

	node := s.nodes[l-1]
	s.nodes = s.nodes[:l-1]
	return node, nil
}

func inorderTraversal(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}
	current := root
	stack := Newstack()

	for {
		for current != nil {
			stack.Push(current)
			current = current.left
		}

		node, err := stack.Pop()
		if err != nil {
			break
		}

		result = append(result, node.Val)

		if node.right != nil {
			current = node.right
		}
	}

	return result
}

func main() {
	n1 := TreeNode{Val: 1}
	n2 := TreeNode{Val: 2}
	n3 := TreeNode{Val: 3}
	n1.left = nil
	n1.right = &n2
	n2.left = &n3
	n2.right = nil
	n3.left = nil
	n3.right = nil
	output := inorderTraversal(&n1)
	fmt.Printf("output is %v\n", output)
	for _, x := range output {
		fmt.Println(x)
	}
}
