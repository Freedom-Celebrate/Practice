package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func hex(s string) string {
	var w []string = strings.Fields(s)

	for i := 0; i < len(w); i++ {
		if w[i] == "(hex)" || w[i] == "(HEX)" && i > 0 {
			previousWord := w[i-1]
			hexstr, _ := strconv.ParseInt(previousWord, 16, 64)

			w[i-1] = fmt.Sprint(hexstr)
			// previousWord = fmt.Sprint(hexstr)
			// w[i-1] = previousWord
			// fmt.Println(previousWord)
			// fmt.Println(w[i-1])
			w = append(w[:i], w[i+1:]...)
			// i--

		}

	}
	return strings.Join(w, "_")
	// return w
}

func bin(s string) string {
	var w []string = strings.Fields(s)

	for i := 0; i < len(w); i++ {
		if w[i] == "(bin)" && i > 0 {
			previousWord := w[i-1]
			binstr, _ := strconv.ParseInt(previousWord, 2, 64)

			w[i-1] = fmt.Sprint(binstr)

			w = append(w[:i], w[i+1:]...)
			i--

		}

	}
	return strings.Join(w, " ")
}

func up(s string) string {
	var w []string = strings.Fields(s)

	for i := 0; i < len(w); i++ {
		if w[i] == "(up)" && i > 0 {

			w[i-1] = strings.ToUpper(w[i-1])

			w = append(w[:i], w[i+1:]...)
			i--

		}

	}
	return strings.Join(w, " ")
}

func mulUp(s string) string {
	w := strings.Fields(s)

	for i := 0; i < len(w)-1; i++ {
		if strings.HasPrefix(w[i], "(up,") && strings.HasSuffix(w[i+1], ")") {

			number := w[i+1][:len(w[i+1])-1]

			count, _ := strconv.Atoi(number)
			w = append(w[:i], w[i+2:]...)
			i--

			start := i - count + 1
			if start < 0 {
				start = 0
			}
			for j := start; j <= i; j++ {
				w[j] = strings.ToUpper(w[j])
			}
		}

	}
	return strings.Join(w, " ")

}
func mulLow(s string) string {
	w := strings.Fields(s)

	for i := 0; i < len(w)-1; i++ {
		if strings.HasPrefix(w[i], "(low,") && strings.HasSuffix(w[i+1], ")") {

			number := w[i+1][:len(w[i+1])-1]

			count, _ := strconv.Atoi(number)
			w = append(w[:i], w[i+2:]...)
			i--

			start := i - count + 1
			if start < 0 {
				start = 0
			}
			for j := start; j <= i; j++ {
				w[j] = strings.ToLower(w[j])
			}
		}

	}
	return strings.Join(w, " ")

}
func mulCap(s string) string {
	w := strings.Fields(s)

	for i := 0; i < len(w)-1; i++ {
		if strings.HasPrefix(w[i], "(cap,") && strings.HasSuffix(w[i+1], ")") {

			number := w[i+1][:len(w[i+1])-1]

			count, _ := strconv.Atoi(number)
			w = append(w[:i], w[i+2:]...)
			i--

			start := i - count + 1
			if start < 0 {
				start = 0
			}
			for j := start; j <= i; j++ {
				w[j] = strings.ToUpper(string(w[j][0])) + strings.ToLower(w[j][1:])
			}
		}

	}
	return strings.Join(w, " ")

}
func GoisFun(a string) []string {
	s := strings.Fields(a)
	var result []string
	for i := 0; i < len(s); i++ {
		if s[i] == "(up)" && i > 0 {
			result[i-1] = strings.ToUpper(result[i-1])
			continue
		}
		result = append(result, s[i])
	}
	return result
}

func low(s string) string {
	var w []string = strings.Fields(s)

	for i := 0; i < len(w); i++ {
		if w[i] == "(low)" && i > 0 {

			w[i-1] = strings.ToLower(w[i-1])

			w = append(w[:i], w[i+1:]...)
			i--

		}

	}
	return strings.Join(w, " ")
}

func cap(s string) string {
	var w []string = strings.Fields(s)

	for i := 0; i < len(w); i++ {
		if w[i] == "(cap)" && i > 0 {

			w[i-1] = strings.ToUpper(string(w[i-1][0])) + strings.ToLower(string(w[i-1][1:]))

			w = append(w[:i], w[i+1:]...)
			i--

		}

	}
	return strings.Join(w, " ")
}

func fixarticles(s string) string {
	words := strings.Fields(s)
	vowels := "aeiouhAEIOUH"

	for i := 0; i < len(words); i++ {
		if i < len(words) {
			word := words[i]
			if word == "a" && strings.ContainsAny(vowels, string(words[i+1][0])) {
				words[i] = "an"
			} else if word == "A" && strings.ContainsAny(vowels, string(words[i+1][0])) {
				words[i] = "An"
			} else if word == "an" && !strings.ContainsAny(vowels, string(words[i+1][0])) {
				words[i] = "a"
			}
		}

	}

	return strings.Join(words, " ")

}

func fixpunctuation(s string) string {
	SpaceBefore := regexp.MustCompile(`\s+([.?!,;:]+)`)
	s = SpaceBefore.ReplaceAllString(s, "$1 ")
	return s

}

func main() {

	fmt.Println(hex("1E (HEX) 1F (hex) files were added"))
	fmt.Println(bin("10 (bin) files were added"))
	fmt.Println(GoisFun("dgf (up) files were added"))
	fmt.Println(low("sdf SDK (low) files were added"))
	fmt.Println(cap("sdf SDK (cap) files were added"))
	fmt.Println(mulUp("sdf dsg (up, 2) were added"))
	fmt.Println(mulLow("ASDF SDF SFGD (low, 3) were added"))
	fmt.Println(mulCap("ASDF SDF SFGD (cap, 3) were added"))
	fmt.Println(fixarticles("A apple an book"))
	fmt.Println(fixpunctuation("A apple ...an ,book !!"))

}
