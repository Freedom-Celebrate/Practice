// Write a function that prints in ascending order and on a single line: all possible combinations of two different two-digit numbers.

// These combinations are separated by a comma and a space.

// Write a function that prints, in ascending order and on a single line: all unique combinations of three different digits so that, the first digit is lower than the second, and the second is lower than the third.

// These combinations are separated by a comma and a space.
package main

import (
	"fmt"
)

func PrintComb2() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 8; j++ {
			for k := 0; k <= 9; k++ {
				for l := 1; l <= 9; l++ {

					comb := fmt.Sprintf("%d%d %d%d, ", i, j, k, l)
					if len(comb) == 7 {
						fmt.Print(comb)

					}

				}

			}

		}
	}

}
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
	PrintComb2()
	PrintComb()

}

//required output
//00 01, 00 02, 00 03, ..., 00 98, 00 99, 01 02, 01 03, ..., 97 98, 97 99, 98 99$
