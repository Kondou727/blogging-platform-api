package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) createBlogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Called createdBlogHandler")
}
