package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	bstIterator := BSTIterator{
		stack: []*TreeNode{},
	}
	for root != nil {
		bstIterator.stack = append(bstIterator.stack, root)
		root = root.Left
	}
	return bstIterator
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	peek := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	if peek.Right != nil {
		p := peek.Right
		for p != nil {
			this.stack = append(this.stack, p)
			p = p.Left
		}
	}

	return peek.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) != 0

}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
func main() {

	n1 := TreeNode{Val: 1}
	n2 := TreeNode{Val: 2}
	n3 := TreeNode{Val: 3}
	n4 := TreeNode{Val: 4}
	n5 := TreeNode{Val: 5}
	n6 := TreeNode{Val: 6}

	n1.Left = &n2
	n1.Right = &n5
	n2.Left = &n3
	n2.Right = &n4
	n3.Left = nil
	n3.Right = nil
	n4.Left = nil
	n4.Right = nil
	n5.Left = nil
	n5.Right = &n6
	n6.Left = nil
	n6.Right = nil

	obj := Constructor(&n1)
	//param_1 := obj.Next()
	//fmt.Println(param_1)
	//param_2 := obj.HasNext()
	//fmt.Println(param_2)
	for obj.HasNext() {
		n := obj.Next()
		fmt.Println(n)
	}

}
