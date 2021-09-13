package main

import (
	"reflect"

	"github.com/olivere/elastic/v7"
)

const (
	POST_INDEX = "post"
)

type Post struct {
	Id      string `json:"id"`
	User    string `json:"user"`
	Message string `json:"message"`
	Url     string `json:"url"`
	Type    string `json:"type"`
}

func searchPostsByUser(user string) ([]Post, error) {
	// send query to ES to get all posts from the sepcified user
	query := elastic.NewTermQuery("user", user)
	searchResult, err := readFromES(query, POST_INDEX)

	if err != nil {
		return nil, err
	}
	return getPostFromSearchResult(searchResult), nil

}

func getPostFromSearchResult(searchResult *elastic.SearchResult) []Post {
	var ptype Post
	var posts []Post

	// Loop all the returned results from ES and add to a result array []Post.
	// First, do a sanity check of the type of the data returned from ES, since there is no constraint about
	// the data type when storing the data into ES
	for _, item := range searchResult.Each(reflect.TypeOf(ptype)) {
		p := item.(Post)
		posts = append(posts, p)
	}
	return posts
}
