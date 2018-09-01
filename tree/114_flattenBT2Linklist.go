package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func faltten(root *TreeNode) {
	pn := root
	for pn != nil {
		if pn.Left == nil {
			pn = pn.Right
		} else {
			p1 := pn.Left
			for p1.Right != nil {
				p1 = p1.Right
			}
			p1.Right = pn.Right
			pn.Right = pn.Left
			pn.Left = nil
			pn = pn.Right
		}
	}
}
