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
				for l := j + 1; l <= 9; l++ {
					num := i*10 + j
					num2 := k*10 + l
					/*
						OR THIS:
						num := fmt.Sprintf("%d%d", i,j)
						num2 := fmt.Sprintf("%d%d", k,l)
					*/
					if num > num2 {
						continue
					} else if i == 9 && j == 8 && k == 9 && l == 9 {
						fmt.Printf("%d%d %d%d\n", i, j, k, l)
						break
					} else {
						fmt.Printf("%d%d %d%d, ", i, j, k, l)
					}
				}

			}

		}
	}

}

func main() {
	PrintComb2()

}

//required output
//00 01, 00 02, 00 03, ..., 00 98, 00 99, 01 02, 01 03, ..., 97 98, 97 99, 98 99$
