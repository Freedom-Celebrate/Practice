package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . sample.txt result.txt")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	input, _ := os.ReadFile(inputFile)

	result := string(input)
	result = processor(result)
	result1 := []byte(result)

	os.WriteFile(outputFile, result1, 0644)

}
