//Write a function that takes a pointer to an int as argument and gives to this int the value of 1.

package main

import "fmt"

func PointOne(n *int) {
	*n = 7

}
func main() {

	m := 5
	PointOne(&m)
	fmt.Print(m)

}
