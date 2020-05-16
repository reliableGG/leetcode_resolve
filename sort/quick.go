package main

import (
	"fmt"
)

const MAX = 10

var sortArray = []int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}

func main() {

	//fmt.Println("before sort：")
	//show()

	//quickSort(sortArray, 0, MAX-1)
	quickSort(sortArray)

	//fmt.Println("after sort:")
	//show()

}

func quickSort(arr []int) []int {
	sort(arr, 0, len(arr)-1)

	return arr
}

func sort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		sort(arr, low, pi-1)
		sort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			fmt.Println(arr[i], "<->", arr[j])
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	fmt.Println(arr[i+1], "<->", arr[high])
	arr[i+1], arr[high] = arr[high], arr[i+1]

	fmt.Println(sortArray)
	fmt.Println("partion ", i+1, "=>", sortArray[i+1])
	return i + 1
}

// foreach
func show() {
	for _, value := range sortArray {
		fmt.Printf("%d\t", value)
	}
}
