package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hex(s string) string {
	var w []string = strings.Fields(s)

	for i := 0; i < len(w); i++ {
		if w[i] == "(hex)" && i > 0 {
			pword := w[i-1]
			hexstr, _ := strconv.ParseInt(pword, 16, 64)

			w[i-1] = fmt.Sprint(hexstr)

			w = append(w[:i], w[i+1:]...)
			i--

		}

	}
	return strings.Join(w, " ")
}
func bin(s string) string {
	var w []string = strings.Fields(s)

	for i := 0; i < len(w); i++ {
		if w[i] == "(bin)" && i > 0 {
			pword := w[i-1]
			hexstr, _ := strconv.ParseInt(pword, 2, 64)

			w[i-1] = fmt.Sprint(hexstr)

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

	for i := 0; i < len(w); i++ {
		if strings.HasPrefix(w[i], "(up, ") && strings.HasSuffix(w[i], ")") {
			w[i] = strings.TrimPrefix(w[i], "")
			w[i] = strings.TrimSuffix(w[i], "")
			number, _ := strconv.Atoi(w[i])

			w[i-number] = strings.ToUpper(w[i-number])

			w = append(w[:i], w[i+1:]...)
			i--

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

			w[i-1] = strings.ToUpper(w[i-1])

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

func main() {
	fmt.Println(hex("1E (hex) files were added"))
	fmt.Println(bin("10 (bin) files were added"))
	fmt.Println(GoisFun("dgf (up, 2) files were added"))
	fmt.Println(low("sdf SDK (low) files were added"))
	fmt.Println(cap("sdf SDK (cap) files were added"))
	fmt.Println(mulUp("sdf dsg (up,2) files were added"))

}
