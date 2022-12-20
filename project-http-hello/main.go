package main

import (
	"fmt"
	"net/http"
)

var calls []string
var stats map[string]int

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	calls = append(calls, name)
	stats[name] += 1

	fmt.Printf("calls: %#v\n", calls)
	fmt.Printf("stats: %#v\n\n", stats)

	fmt.Fprint(w, "Hello, ", name)
}

func main() {
	// Your solution goes here. Good luck!

	stats = map[string]int{}

	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8080", nil)
}
