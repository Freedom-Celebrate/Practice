package main

import (
	"fmt"
	"strconv"
	"strings"
)

func upN(input string) {
	cmd := ""
	inBracket := false
	start := 0
	cleanedStr := ""
	for i, char := range input {
		if !inBracket && char == '(' {
			start = i + 1
			inBracket = true
			continue
		}
		if inBracket && char == ')' {
			cmd = input[start:i]
			inBracket = false
			continue
		}
		if !inBracket {
			cleanedStr += string(char)
		}
	}
	cleanedStr = strings.ReplaceAll(cleanedStr, "  ", " ")
	words := strings.Fields(cleanedStr)
	parts := strings.Fields((cmd))
	numStr := strings.TrimSpace(parts[len(parts)-1])
	num, _ := strconv.Atoi(numStr)
	for i := len(words) - num; i < len(words); i++ {
		words[i] = strings.ToUpper(string(words[i]))
	}
	fmt.Println(strings.Join(words, " "))
}

func main() {
	upN("i am a boy and i love eating rice and stew (up, 2) ")
}
