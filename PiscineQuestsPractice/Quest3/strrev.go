package main

import "fmt"

func StrRev(s string) string {
	result := ""
	for char := len(s) - 1; char >= 0; char-- {
		result += string(s[char])
	}
	return result
}

func main() {
	s := "Hello World!"
	StrRev(s)
	fmt.Println(s)
}
