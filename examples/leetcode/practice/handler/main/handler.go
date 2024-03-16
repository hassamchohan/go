package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	err := http.ListenAndServe(":8080", handler())
	if err != nil {
		log.Fatal(err)
	}
}

func handler() http.Handler {
	r := http.NewServeMux() //router
	r.HandleFunc("/double", doubleHandler)
	return r
}

func doubleHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("param")
	if text == "" {
		http.Error(w, "missing parameter", http.StatusBadRequest)
		return
	}

	v, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "not a number: "+text, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, v*2)

}
