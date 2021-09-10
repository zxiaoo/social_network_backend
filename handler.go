package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func uploadHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one post request")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

    // parse from the body of request to get a json object
	decoder := json.NewDecoder(r.Body)

	var p Post
	if err := decoder.Decode(&p); err != nil {
		panic(err)
	}

	// response body
	fmt.Fprintf(w, "Post received: %s\n", p.Message)
}