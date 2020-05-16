package main

import "fmt"

func main() {
	//a := ")(())())"
	a := "())()"
	res := longestValidParentheses(a)
	fmt.Println(res)
}

func longestValidParentheses(s string) int {
	stack := Stack{}
	maxLen := 0
	last := -1
	for idx, c := range s {
		if string(c) == "(" {
			stack.Add(idx)
		} else {
			if stack.Empty() {
				last = idx
			} else {
				stack.Pop()
				if !stack.Empty() {
					maxLen = max(maxLen, idx-stack.Peek())
				} else {
					maxLen = max(maxLen, idx-last)
				}
			}

		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Stack struct {
	elems []int
}

func (s *Stack) Add(n int) {
	s.elems = append(s.elems, n)
}

func (s *Stack) Pop() (n int) {
	n = s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return
}

func (s *Stack) Peek() (n int) {
	if len(s.elems) != 0 {
		n = s.elems[len(s.elems)-1]
	}
	return
}

func (s *Stack) Empty() bool {
	return len(s.elems) == 0
}
