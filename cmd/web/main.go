package main

import (
	"log"
	"net/http"
)

func main(){
	// create custom servemux
	mux := http.NewServeMux()

	// register routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// start the server  with custom  servemux
	log.Print("Starting server on http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}