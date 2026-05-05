package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	slice := []string{}
	blocks := [][]string{}


	if strings.HasPrefix(os.Args[1], "--reverse=") {
		flag := strings.TrimPrefix(os.Args[1], "--reverse=")

		data, _ := os.ReadFile(flag)
		line := strings.Split(string(data), "\n")

		for i := 0; i < len(line); i++ {

			if strings.HasSuffix(line[i], "$") {
				line1 := strings.TrimSuffix(line[i], "$")
				slice = append(slice, line1)

			}
			fmt.Println(slice)

		}

		for i := 0; i < len(slice); i += 9 {
			block := slice[i : i+8]
			blocks = append(blocks, block)

		}
		fmt.Println(len(blocks))

		//banner, _ := LoadBanner(os.Args[2])

	} else {
		fmt.Println("Usage: go run . [OPTION]")
	}
}
