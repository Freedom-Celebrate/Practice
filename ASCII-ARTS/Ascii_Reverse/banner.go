package main

import (
	"os"
	"strings"
)

func LoadBanner(filename string) (map[string]rune, error) {
	var data, err = os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	line := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(line, "\n")
	lines = lines[1:]
	var result = make(map[string]rune)

	for char := ' '; char <= '~'; char++ {
		slice := ""
		start := int(char-32) * 9
		slice = strings.Join(lines[start:start+8], "")
		result[slice] = char
	}
	return result, nil

}
