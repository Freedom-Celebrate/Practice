package main

import (
	"fmt"
)

func main() {
	// if len(os.Args) < 3 || len(os.Args) >= 3 {
	// 	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	// 	return
	// }
	// flag := os.Args[1]
	// str := os.Args[2]
	// banner := "standard.txt"

	// result, err := LoadBanner(banner)

	fmt.Println(getTerminalWidth())

}
