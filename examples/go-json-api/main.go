package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/users", makeHTTPHandler(handleGetUserByID))
	http.ListenAndServe(":3000", nil)
}
