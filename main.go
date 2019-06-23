package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, "hello"); err != nil {
			panic(err)
		}
	})
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
