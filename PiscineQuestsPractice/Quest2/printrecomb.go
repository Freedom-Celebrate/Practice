package main

import (
	"fmt"
)

func PrintReComb() {

	for i := 7; i >= 0; i-- {
		for j := 8; j >= 1; j-- {
			for k := 9; k >= 2; k-- {

				comb := fmt.Sprintf("%d%d%d, ", i, j, k)
				fmt.Print(comb)

			}
		}
	}

}

func main() {
	PrintReComb()

}
