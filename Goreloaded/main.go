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

	input, err := os.ReadFile(inputFile)
	if err != nil{
		fmt.Println(err)
		return
	}

	result := string(input)
	result = processor(result)
	result1 := []byte(result)

	err = os.WriteFile(outputFile, result1, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
