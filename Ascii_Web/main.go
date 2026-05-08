package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, m *http.Request) {
		fmt.Fprintln(response, "Welcome to ASCII Art Web")
	})

	port := ":8080"
	fmt.Println("server is running on port:8080")

	log.Fatal(http.ListenAndServe(port, nil))

}
