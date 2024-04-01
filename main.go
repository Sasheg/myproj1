package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	    r.URL.Path = r.URL.Path[1:]
		if r.URL.Path == "" {
		    r.URL.Path = "World"
		}
		fmt.Fprintf(w, "Hello, %s\n", r.URL.Path)
	})
	http.ListenAndServe(":3000", nil)
}
