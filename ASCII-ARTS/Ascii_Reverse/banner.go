package main

import (
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	var data, err = os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	line := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(line, "\n")
	lines = lines[1:]
	var result = make(map[rune][]string)

	for char := ' '; char <= '~'; char++ {
		start := int(char-32) * 9
		result[char] = lines[start : start+8]
	}
	return result, nil

}
