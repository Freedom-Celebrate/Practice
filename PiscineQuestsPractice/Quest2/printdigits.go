// Write a program that prints the decimal digits in ascending order (from 0 to 9) on a single line.

// A line is a sequence of characters preceding the end of line character ('\n').

package main

import "fmt"

func main() {
	for char := '0'; char <= '9'; char++ {
		fmt.Printf("%c", char)
	}
	fmt.Println("\n")
}
