package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	slice := []string{}
	blocks := [][]string{}
	// extracting flag from it prefix
	if strings.HasPrefix(os.Args[1], "--reverse=") {
		flag := strings.TrimPrefix(os.Args[1], "--reverse=")
		fmt.Println(flag)

		// reading the file(flag)
		data, _ := os.ReadFile(flag)
		line := strings.Split(string(data), "\n")
		fmt.Println(line)

		// trimming $ from every line in the file
		for i := 0; i < len(line); i++ {
			if strings.HasSuffix(line[i], "$") {
				line[i] = strings.TrimSuffix(line[i], "$")
				slice = append(slice, line[i])
			} else {
				slice = append(slice, line[i])
			}
		}
		fmt.Println(slice)

		for i := 0; i < len(slice); i += 9 {
			block := slice[i : i+8]
			blocks = append(blocks, block)

		}

		fmt.Println(len(blocks))

		banner, _ := LoadBanner(os.Args[2])
		reverseMap := make(map[string]rune)

		for key, value := range banner {
			value1 := strings.Join(value, "\n")

			reverseMap[value1] = key
		}

	} else {
		fmt.Println("Usage: go run . [OPTION]")
	}
}
