package main

import "fmt"

func StrLen(s string) int {
	// result := []rune(s)
	// return len(result)
	count := 0

	for length := 0; length < len(s); length++ {
		if length < len(s) {
			count++
		}
	}
	return count

}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
