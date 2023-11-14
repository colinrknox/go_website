package main

import (
	"net/http"
	"fmt"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Listent and server error: ", err)
	}
}
