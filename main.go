package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//the ascii character are 2dimensional with vertical rows and horizontal columns
	var aschar [][]string

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <string> st.txt")
		return
	}

	inputText := os.Args[1]
	outputText := os.Args[2]
	words := strings.Split(inputText, "\\n")
	fmt.Println(words)
	data, err := os.ReadFile(outputText)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := string(data)
	for _, val := range result {
		fmt.Println(string(val))
	}
	content := strings.Split(result, "\n\n")

	//spliting to get the rows of each character
	for _, char := range content {
		rows := strings.Split(char, "\n")
		aschar = append(aschar, rows)
	}
	//Contains the first, to the last row of each character
	// PRINTS ASCHAR
	for _, word := range words {

		for row := 0; row <= 7; row++ {
			//holds the index of each character in the string)
			for _, ch := range word {
				columns := int(ch - 32)
				fmt.Print(aschar[columns][row])
			}
			fmt.Println()
		}

	}

}
