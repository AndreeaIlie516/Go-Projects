package main

import (
	"math/rand"
	"encoding/json"
	"net/http"

	"github.com/AndreeaIlie516/Go-Projects/Go-REST-API-Clean-Architecture/entity"
	"github.com/AndreeaIlie516/Go-Projects/Go-REST-API-Clean-Architecture/repository"
)

var (
	repo repository.PostRepository = NewPostRepository()
)

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()

	result, err := json.Marshal(posts)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the posts"}`))
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func addPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var post entity.Post 
	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}
	post.ID, err = rand.Int()
	repo.Save(&post)

	resp.WriteHeader(http.StatusOK)
	json.NewEncode(response).Encode(post)
}