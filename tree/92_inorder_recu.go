package main

import (
	"errors"
	"fmt"
	"sync"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	lock  sync.Mutex
	nodes []*TreeNode
}

func NewStack() *Stack {
	return &Stack{sync.Mutex{}, []*TreeNode{}}
}

func (s *Stack) Push(node *TreeNode) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.nodes = append(s.nodes, node)
}

func (s *Stack) Pop() (*TreeNode, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.nodes)
	if l == 0 {
		return nil, errors.New("Empty")
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

	stack := NewStack()
	current := root

	for {
		for current != nil {
			stack.Push(current)
			current = current.Left
		}

		node, err := stack.Pop()
		if err != nil {
			break
		}

		result = append(result, node.Val)

		if node.Right != nil {
			current = node.Right
		}
	}

	return result
}

func main() {
	n1 := TreeNode{Val: 1}
	n2 := TreeNode{Val: 2}
	n3 := TreeNode{Val: 3}
	n1.Left = nil
	n1.Right = &n2
	n2.Left = &n3
	n2.Right = nil
	n3.Left = nil
	n3.Right = nil
	output := inorderTraversal(&n1)
	fmt.Printf("output is %v\n", output)
	for _, x := range output {
		fmt.Println(x)
	}
}
