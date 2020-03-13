package main

import "fmt"
import "bytes"
import "strconv"

func main() {
	//a := "11"
	//b := "1"
	//c := addBinary(a, b)

	s := "11"
	s2 := "1"
	res := addBinary(s, s2)
	fmt.Println(res)
}

func addBinary(a string, b string) string {
	var buffer bytes.Buffer
	i := len(a) - 1
	j := len(b) - 1
	carry := 0

	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			parta := getInt(i, a)
			sum += parta
		}
		if j >= 0 {
			partb := getInt(j, b)
			sum += partb
		}
		result := sum % 2
		carry = sum / 2

		buffer.WriteString(strconv.Itoa(result))
		i--
		j--
	}
	if carry != 0 {
		buffer.WriteString("1")
	}
	return reverseString(buffer.String())
}

func reverseString(a string) string {
	r := []rune(a)
	length := len(r)
	for i := 0; i < length/2; i++ {
		r[i], r[length-i-1] = r[length-i-1], r[i]
	}
	return string(r)
}

func getInt(index int, s string) int {
	part, _ := strconv.Atoi(string(s[index]))
	return part
}
