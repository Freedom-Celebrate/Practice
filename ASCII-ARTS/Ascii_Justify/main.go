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
	if flag == "justify" {
		Justify(str, bannerMap, flag, term_width)

	} else {
		PrintAligned(slice, flag, term_width, str)

	}

}
