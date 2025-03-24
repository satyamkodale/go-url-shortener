package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to URL Shortner")
}

func main() {
	fmt.Println("<<<<<<<< Starting Server At 1911 >>>>>>>>")
	http.HandleFunc("/", homeHandler)
	error := http.ListenAndServe(":1911", nil)
	if error != nil {
		fmt.Println("Error for starting the server")
	}
	fmt.Println("Server Started")
}
