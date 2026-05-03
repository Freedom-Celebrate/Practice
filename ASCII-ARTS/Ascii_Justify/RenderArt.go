package main

func RenderArt(text string, banner map[rune][]string) []string {
	var slice []string
	for row := 0; row < 8; row++ {
		word := ""
		for _, char := range text {
			word += banner[char][row]
		}
		slice = append(slice, word)
	}
	return slice
}
