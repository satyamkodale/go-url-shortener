package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("<<<<<<<< Starting Server At 1911 >>>>>>>>")
	error := http.ListenAndServe(":1911", nil)
	if error != nil {
		fmt.Println("Error for starting the server")
	}
	fmt.Println("Server Started")
}
