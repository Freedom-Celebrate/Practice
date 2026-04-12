//Write a function that prints 'T' (true) on a single line if the int passed as parameter is negative, otherwise it prints 'F' (false).

package main

import "fmt"

func IsNegative(nb int) {
	if nb >= 0 {
		fmt.Println(string('F'))
	} else if nb < 0 {
		fmt.Println(string('T'))

	}

}
func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)

}
