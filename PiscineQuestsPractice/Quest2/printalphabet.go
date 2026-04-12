// Write a program that prints the Latin alphabet in lowercase on a single line.

// A line is a sequence of characters preceding the end of line character ('\n').

package main

import "fmt"

func main() {
	for char := 'a'; char <= 'z'; char++ {
		fmt.Printf("%c", char)
	}
}
