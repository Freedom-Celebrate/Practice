package main

import "fmt"

func add(num int) int {

	if num == 1 {
		return num
	}
	return num + add(num-1)
	// for i := num; i > 0; i-- {

	// 	result += i

}

func reverse(s string) string {
	if len(s) == 1 {
		return s
	}
	// return reverse(s[1:]) + string(s[0])
	return string(s[len(s)-1]) + reverse(s[:len(s)-1])
}

func main() {
	fmt.Println(add(5))
	fmt.Println(reverse("favour"))
}
