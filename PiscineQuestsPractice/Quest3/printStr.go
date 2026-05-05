package main

import "fmt"

func PrintStr(s string) {
	for _, character := range s {
		fmt.Print(string(character))
	}

}
func main() {
	PrintStr("Hello World!")
}
