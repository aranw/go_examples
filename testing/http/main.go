package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Add(x, y int) int {
	return x + y
}

func addHandler(w http.ResponseWriter, r *http.Request) {

	first, second := r.FormValue("x"), r.FormValue("y")
	x, err := strconv.Atoi(first)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	y, err := strconv.Atoi(second)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	// The main core logic of this handler should be abstracted out of the handler
	// This allows for that code to be tested independently
	z := Add(x, y)

	fmt.Fprintf(w, strconv.Itoa(z))
}

// Having the http ServeMux in a separate method makes it easier to test the routing of the application.
func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/add", addHandler)
	return r
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", handler()))
}
