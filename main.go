package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)

	fmt.Println("hi")
}
