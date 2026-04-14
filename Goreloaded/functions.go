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

			//w[i-1] = fmt.Sprint(hexstr)
			previousWord = fmt.Sprint(hexstr)
			w[i-1] = previousWord
			//fmt.Print(previousWord)
			//fmt.Print(w[i-1])
			w = append(w[:i], w[i+1:]...)
			i--

		}

	}
	return strings.Join(w, " ")
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
			//i--

		}

	}
	return strings.Join(w, " ")
}

func mulUp(s string) string {
	w := strings.Fields(s)

	for i := 0; i < len(w); i++ {
		//fmt.Println(i, "for")
		if w[i] == "(up," && strings.HasSuffix(w[i+1], ")") {

			number := string(w[i+1][0])

			count, _ := strconv.Atoi(number)

			//fmt.Println(i, "if")

			start := i - count
			if start < 0 {
				start = 0
			}
			for j := start; j <= i; j++ {
				w[j] = strings.ToUpper(w[j])
			}
			w = append(w[:i], w[i+2:]...)
			i--

			//fmt.Print(number) // "3)"
		}

	}
	return strings.Join(w, " ")

}
func mulLow(s string) string {
	w := strings.Fields(s)

	for i := 0; i < len(w); i++ {
		if w[i] == "(low," && strings.HasSuffix(w[i+1], ")") {

			number := string(w[i+1][0])
			// number := w[i+1][:len(w[i+1])-1]

			count, _ := strconv.Atoi(number)

			start := i - count
			if start < 0 {
				start = 0
			}
			for j := start; j <= i; j++ {
				w[j] = strings.ToLower(w[j])
			}
			w = append(w[:i], w[i+2:]...)
			i--

		}

	}
	return strings.Join(w, " ")

}
func mulCap(s string) string {
	w := strings.Fields(s)

	for i := 0; i < len(w); i++ {
		if w[i] == "(cap," && strings.HasSuffix(w[i+1], ")") {

			number := string(w[i+1][0])

			count, _ := strconv.Atoi(number)

			start := i - count //
			if start < 0 {
				start = 0
			}
			for j := start; j <= i; j++ {
				w[j] = strings.ToUpper(string(w[j][0])) + strings.ToLower(w[j][1:])
			}
			w = append(w[:i], w[i+2:]...)
			i--

		}

	}
	return strings.Join(w, " ")

}

// func GoisFun(a string) []string {
// 	s := strings.Fields(a)
// 	var result []string
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == "(up)" && i > 0 {
// 			result[i-1] = strings.ToUpper(result[i-1])
// 			continue
// 		}
// 		result = append(result, s[i])
// 	}
// 	return result
// }

// func GoisFun(a string) []string {
// 	s := strings.Fields(a)
// 	var result []string
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == "(up)" {
// 			if len(result) > 0 {
// 				result[i-1] = strings.ToUpper(result[i-1])
// 			}

// 		} else {
// 			result = append(result, s[i])
// 		}

//		}
//		return result
//	}

func fixQuote(s string) string {

	pattern := regexp.MustCompile(`(^|\s)'\s+`)
	s = pattern.ReplaceAllString(s, "$1'")

	pattern1 := regexp.MustCompile(`\s+'(\s|$)`)
	s = pattern1.ReplaceAllString(s, "'$1")

	pattern2 := regexp.MustCompile(`([;:])\s+'`)
	s = pattern2.ReplaceAllString(s, "$1 '")

	return s
}

func fixPunctuation(s string) string {
	SpaceBefore := regexp.MustCompile(`\s*([.?!,;:]+)\s*`)
	s = SpaceBefore.ReplaceAllString(s, "$1 ")
	return s

}

func fixArticles(s string) string {
	words := strings.Fields(s)
	vowels := "aeiouhAEIOUH"
	last := len(words) - 1

	for i := 0; i < len(words); i++ {
		if i < len(words) && i != last {

			if words[i] == "a" && strings.ContainsAny(vowels, string(words[i+1][0])) {
				words[i] = "an"
			} else if words[i] == "A" && strings.ContainsAny(vowels, string(words[i+1][0])) {
				words[i] = "An"
			} else if words[i] == "an" && !strings.ContainsAny(vowels, string(words[i+1][0])) {
				words[i] = "a"
			} else if words[i] == "An" && !strings.ContainsAny(vowels, string(words[i+1][0])) {
				words[i] = "A"
			}
		}

	}

	return strings.Join(words, " ")

}
func processor(s string) string {
	s = hex(s)
	s = bin(s)
	s = mulUp(s)
	s = mulLow(s)
	s = mulCap(s)
	s = up(s)
	s = low(s)
	s = cap(s)
	s = fixPunctuation(s)
	s = fixQuote(s)

	s = fixArticles(s) // Always last
	return s
}
