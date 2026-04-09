package main

import (
	"fmt"
	"strconv"
	"strings"
)

// func hex(s string) string {

// 	word := strings.Split(s, " ")

// 	for i, _ := range word {
// 		if i < len(word) {
// 			if word[i] == "(hex)" || word[i] == "(HEX)" {
// 				decimal, _ := strconv.ParseInt(word[i-1], 16, 64)

// 				word[i-1] = strconv.FormatInt(decimal, 10)

// 				word = append(word[:i], word[i+1:]...)
// 			}
// 		}

// 	}
// 	return strings.Join(word, " ")

}

func main() {
	fmt.Println(hex("1E (hex) 1F (HEX) first fire"))

}
