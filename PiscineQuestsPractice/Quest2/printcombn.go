// //     Write a function that prints all possible combinations of n different digits in ascending order.

// //     n will be defined as : 0 < n < 10

// // Below are the references for the printing format expected.

// //     (for n = 1) '0, 1, 2, 3, ..., 8, 9'

// //     (for n = 3) '012, 013, 014, 015, 016, 017, 018, 019, 023,...689, 789'

// package main

// import "fmt"

// func PrintCombN(n int) {
// 	//3
// 	for i := 0; i <= 7; i++ {
// 		for j := i + 1; j <= 8; j++ {
// 			for k := j + 1; k <= 9; k++ {

// 				//2
// 				for l := 0; l == 0; l++ {
// 					for m := l + 1; m >= 99; m++ {

// 						//1
// 						for o := 0; o >= 9; o++ {
// 							num := 0 * 10
// 							num2 := l*10 + m
// 							num3 := i + j*10 + k

// 							if n == 1 {
// 								n = num
// 								fmt.Printf("%d", o)
// 							} else if n == 2 {
// 								n = num2
// 								fmt.Println("%d%d", l, m)

// 							} else if n == 3 {
// 								n = num3
// 								fmt.Println("%d%d%d", i, j, k)

// 							}

// 						}
// 					}

// 					// comb := fmt.Sprintf("%d%d%d, ", i, j, k)
// 					// fmt.Print(comb)

// 				}
// 			}
// 		}

// 	}
// }

// func main() {
// 	PrintCombN(1)

// }
