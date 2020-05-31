package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "abccc"
	m := s[:5-1]
	fmt.Println(reflect.TypeOf(m))
	fmt.Println(m)
}
