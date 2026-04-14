// Write a function that prints an int passed in parameter. All possible values of type int have to go through. You cannot convert to int64.

package main

import (
	"fmt"
)

func PrintCombN(n int) {
	fmt.Print(n)

}

func main() {
	PrintCombN(-128)
	PrintCombN(0)
	PrintCombN(128)

}
