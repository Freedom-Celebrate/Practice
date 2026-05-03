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
	lines := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines1 := strings.Split(lines, "\n")
	lines1 = lines1[1:]
	var result = make(map[rune][]string)

	for char := ' '; char <= '~'; char++ {
		for row := 0; row < 8; row++ {
			result[char] = append(result[char], lines1[int(char-32)*9+row])

		}
	}
	return result, nil

}
