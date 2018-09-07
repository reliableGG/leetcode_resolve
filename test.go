package main

import "fmt"

func test() (result []int) {
	return result
}

func main() {
	var mainRes []int
	mainRes = test()
	fmt.Printf("%v\n", mainRes)
}
