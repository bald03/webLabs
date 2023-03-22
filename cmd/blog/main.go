package main

import (
	"fmt"
	"net/http"
)

const port = ":3000"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", index)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(".static"))))

	fmt.Println("Start server")
	http.ListenAndServe(port, mux)
}
