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
	sync.Mutex
	Nodes []*TreeNode
}

func NewStack() Stack {
	return Stack{sync.Mutex{}, []*TreeNode{}}
}

func (s *Stack) Push(node *TreeNode) {
	s.Lock()
	defer s.Unlock()

	s.Nodes = append(s.Nodes, node)
}

func (s *Stack) Pop() (*TreeNode, error) {
	s.Lock()
	defer s.Unlock()

	l := len(s.Nodes)
	if l == 0 {
		return nil, errors.New("Empty Stack")
	}
	node := s.Nodes[l-1]
	s.Nodes = s.Nodes[:l-1]
	return node, nil
}

func preOrderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	stack := NewStack()

	for {
		for root != nil {
			result = append(result, root.Val)
			stack.Push(root)
			root = root.Left
		}
		current, err := stack.Pop()
		if err != nil {
			break
		}
		if current.Right != nil {
			root = current.Right
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
	res := preOrderTraversal(&n1)
	for _, ele := range res {
		fmt.Println(ele)
	}
}
