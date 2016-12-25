package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if (len(name) > 2) {
		fmt.Fprintf(w, "Hello, %s ! \n", name)
	} else {
		fmt.Fprintf(w, "Broken request \n")
		w.WriteHeader(400)
	}
}

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8000", nil)
}