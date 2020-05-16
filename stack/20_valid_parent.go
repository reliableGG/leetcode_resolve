package main

import "fmt"

func main() {
	ca := "{{{}}}"
	fmt.Println(isValid(ca))
}

func isValid(s string) bool {
	stack := Stack{}
	for _, b := range s {
		c := string(b)
		if c == "{" || c == "[" || c == "(" {
			stack.Push(c)
		} else {
			peek := stack.Pop()
			switch {
			case c == "}" && peek != "{":
				return false
			case c == "]" && peek != "[":
				return false
			case c == ")" && peek != "(":
				return false
			}
		}
	}
	return len(stack.e) == 0
}

type Stack struct {
	e []string
}

func (s *Stack) Pop() string {
	res := ""
	if len(s.e)-1 >= 0 {
		res = s.e[len(s.e)-1]
		s.e = s.e[:len(s.e)-1]
	}
	return res
}

func (s *Stack) Push(str string) {
	s.e = append(s.e, str)
}
