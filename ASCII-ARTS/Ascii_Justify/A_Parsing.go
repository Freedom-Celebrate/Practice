package main

import (
	"fmt"
	"os"
	"strings"
)

func ArgumentParsing() (string, string, string) {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		return "", "", ""
	}
	var flag string
	var str string
	var banner string

	if strings.HasPrefix(os.Args[1], "--align=") {
		flag = strings.TrimPrefix(os.Args[1], "--align=")
		// validation
		if flag == "left" || flag == "right" || flag == "justify" || flag == "center" {
			str = os.Args[2]
			if len(os.Args) == 4 {
				banner = os.Args[3]
			} else {
				banner = "standard.txt"
			}
			//if flag is invalid
		} else {
			fmt.Println("invalid alignment type")
			return "", "", ""
		}
		// if no flag is present
	} else {
		flag = "left"
		str = os.Args[1]
		// if flag is not present
		if len(os.Args) == 3 {
			banner = os.Args[2]
		} else {
			banner = "standard.txt"
		}

	}
	return flag, str, banner

}
