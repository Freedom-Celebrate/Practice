package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("file.txt")
	//banner, _ := LoadBanner("standard.txt")
	line := strings.Split(string(data), "\n")
	fmt.Println(line)
	//slice := []string{}

	//result := ""
	fmt.Println(len(string(data)))
	// i := 0
	// for i >= 0 && i < len(line) {
	// 	slice = append(slice, line[i:i+8]...)
	// 	i += 9
	// }
	// fmt.Println(result)
}
