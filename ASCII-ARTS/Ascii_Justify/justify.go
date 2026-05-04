package main

import (
	"fmt"
	"strings"
)

func Justify(text string, banner map[rune][]string, alignment string, termWidth int) {
	parts := strings.Split(text, " ")

	for r := 0; r < 8; r++ {
		totalWidth := 0
		result := ""

		for i := 0; i < len(parts); i++ {
			parts1 := RenderArt(parts[i], banner)
			currentRow := len(parts1[r])

			totalWidth += currentRow

		}
		leftover := termWidth - totalWidth
		gaps := len(parts) - 1
		n := leftover / gaps
		padding := strings.Repeat(" ", n)

		for i := 0; i < len(parts); i++ {
			parts1 := RenderArt(parts[i], banner)

			if i == len(parts)-1 {
				result += parts1[r]
			} else {
				result += parts1[r] + padding
			}
		}
		fmt.Println(result)

	}
}
