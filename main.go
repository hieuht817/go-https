package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}