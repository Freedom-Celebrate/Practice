package main

import (
	"fmt"
	"strings"
)

func PrintAligned(row []string, alignment string, termWidth int) {
	var n int

	for str := 0; str < 8; str++ {
		currentRow := row[str]
		if alignment == "right" {
			n = termWidth - len(currentRow)

		} else if alignment == "center" {
			n = (termWidth - len(currentRow)) / 2
		} else {
			n = 0
		}
		padding := strings.Repeat(" ", n)

		fmt.Println(padding + currentRow)
	}
}
