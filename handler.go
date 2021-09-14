package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// handler to upload a post
func uploadHandler(w http.ResponseWriter, r *http.Request) {
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

// handler to search posts based on "user" or "keywords"
func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one request for search")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		return
	}

	// get "user" and "keywords" from the url
	user := r.URL.Query().Get("user")
	keywords := r.URL.Query().Get("keywords")

	var posts []Post
	var err error
	if user != "" {
		posts, err = searchPostsByUser(user)
	} else {
		posts, err = searchPostsByKeywords(keywords)
	}

	if err != nil {
		http.Error(w, "Failed to read post from Elasticsearch", http.StatusInternalServerError)
		fmt.Printf("Failed to read posts from Elasticsearch %v.\n", err)
		return
	}

	// convert post type to json
	js, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, "Failed to parse posts into JSON format", http.StatusInternalServerError)
		fmt.Printf("Failed to parse posts into JSON format %v.\n", err)
		return
	}
	// write json into the body of response
	w.Write(js)
}
