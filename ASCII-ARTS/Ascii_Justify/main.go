package main

import "fmt"

func main() {
	flag, str, banner := ArgumentParsing()
	bannerMap, err := LoadBanner(banner)
	if err != nil {
		fmt.Println(err)
		return
	}
	slice := RenderArt(str, bannerMap)
	term_width := getTerminalWidth()
	PrintAligned(slice, flag, term_width)
}
