package main

import (
	"fmt"
	"net/http"
)

func main() {


	http.HandleFunc("/add", handleAdd)
	http.HandleFunc("/subtract", handleSubtract)
	http.HandleFunc("/multiply", handleMutliply)
	http.HandleFunc("/divide", handleDivide)

	http.ListenAndServe(":8080", nil)

	fmt.Println("Hello world")
}