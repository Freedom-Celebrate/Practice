// Write a function that prints, in ascending order and on a single line: all unique combinations of three different digits so that, the first digit is lower than the second, and the second is lower than the third.

// These combinations are separated by a comma and a space.
package main

import (
	"fmt"
)

func PrintComb() {

	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {

				comb := fmt.Sprintf("%d%d%d, ", i, j, k)
				fmt.Print(comb)

			}
		}
	}

}
func main() {
	PrintComb()

}
