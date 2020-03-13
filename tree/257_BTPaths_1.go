package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePathsHelper(root *TreeNode, path []int, result *[]string) {
	if root == nil {
		return
	}

	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if len(path) > 0 {
			s := []string{}
			for i := 0; i < len(path); i++ {
				s = append(s, strconv.Itoa(path[i]))
			}
			*result = append(*result, strings.Join(s, "->"))
		}
		return
	}

	binaryTreePathsHelper(root.Left, path, result)
	binaryTreePathsHelper(root.Right, path, result)
}

func binaryTreePaths(root *TreeNode) []string {
	var result []string
	var path []int
	binaryTreePathsHelper(root, path, &result)
	return result
}

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

	node := binaryTreePaths(&n1)
	fmt.Println(node)
}
