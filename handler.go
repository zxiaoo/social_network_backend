package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func uploadHandler (w http.ResponseWriter, r *http.Request) {

	// parse from the body of request to get a json object
	fmt.Println("Received one post request")
	decoder := json.NewDecoder(r.Body)

	var p Post
	if err := decoder.Decode(&p); err != nil {
		panic(err)
	}

	// response body
	fmt.Fprintf(w, "Post received: %s\n", p.Message)
}