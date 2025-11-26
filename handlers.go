package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Kondou727/blogging-platform-api/internal/database"
)

func (cfg *apiConfig) createBlogHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Title    string          `json:"title"`
		Content  string          `json:"content"`
		Category string          `json:"category"`
		Tags     json.RawMessage `json:"tags"`
	}
	params := parameters{}

	err := decodeJSON(r, &params)
	if err != nil {
		log.Printf("failed to decode json: %s", err)
		respondWithJSON(w, http.StatusInternalServerError, errResp{Error: err.Error()})
		return
	}

	blog, err := cfg.BlogsDBQueries.CreateBlog(r.Context(), database.CreateBlogParams{
		Title:    params.Title,
		Content:  params.Content,
		Category: params.Category,
		Tags:     params.Tags,
	})
	if err != nil {
		log.Printf("failed to create blog: %s", err)
		respondWithJSON(w, http.StatusBadRequest, errResp{Error: err.Error()})
		return
	}

	respondWithJSON(w, http.StatusCreated, blogResp{
		ID:        int(blog.ID),
		Title:     blog.Title,
		Category:  blog.Category,
		Tags:      blog.Tags,
		CreatedAt: blog.Createdat,
		UpdatedAt: blog.Updatedat,
	})
}
func (cfg *apiConfig) updateBlogHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Title    string          `json:"title"`
		Content  string          `json:"content"`
		Category string          `json:"category"`
		Tags     json.RawMessage `json:"tags"`
	}
	params := parameters{}

	err := decodeJSON(r, &params)
	if err != nil {
		log.Printf("failed to decode json: %s", err)
		respondWithJSON(w, http.StatusInternalServerError, errResp{Error: err.Error()})
		return
	}

	blogID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("failed to get ID: %s", err)
		respondWithJSON(w, http.StatusBadRequest, errResp{Error: err.Error()})
		return
	}
	blog, err := cfg.BlogsDBQueries.UpdateBlog(r.Context(), database.UpdateBlogParams{
		Title:    params.Title,
		Content:  params.Content,
		Category: params.Category,
		Tags:     params.Tags,
		ID:       int64(blogID),
	})

	if err == sql.ErrNoRows {
		log.Printf("blog not found: %s\n", err)
		respondWithJSON(w, http.StatusNotFound, errResp{Error: "blog not found"})
		return
	} else if err != nil {
		log.Printf("failed to update blog: %s\n", err)
		respondWithJSON(w, http.StatusBadRequest, errResp{Error: err.Error()})
		return
	}

	respondWithJSON(w, http.StatusOK, blogResp{
		ID:        int(blog.ID),
		Title:     blog.Title,
		Category:  blog.Category,
		Tags:      blog.Tags,
		CreatedAt: blog.Createdat,
		UpdatedAt: blog.Updatedat,
	})
}
func (cfg *apiConfig) deleteBlogHandler(w http.ResponseWriter, r *http.Request) {
	blogID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("failed to get ID: %s", err)
		respondWithJSON(w, http.StatusBadRequest, errResp{Error: err.Error()})
		return
	}

	title, _ := cfg.BlogsDBQueries.DeleteBlog(r.Context(), int64(blogID))
	if title == "" {
		log.Println("blog not found")
		respondWithJSON(w, http.StatusNotFound, errResp{Error: "blog not found"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cfg *apiConfig) getBlogHandler(w http.ResponseWriter, r *http.Request) {
	blogID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("failed to get ID: %s", err)
		respondWithJSON(w, http.StatusBadRequest, errResp{Error: err.Error()})
		return
	}

	blog, err := cfg.BlogsDBQueries.GetBlog(r.Context(), int64(blogID))
	if err == sql.ErrNoRows {
		log.Printf("blog not found: %s\n", err)
		respondWithJSON(w, http.StatusNotFound, errResp{Error: "blog not found"})
		return
	}

	respondWithJSON(w, http.StatusOK, blogResp{
		ID:        int(blog.ID),
		Title:     blog.Title,
		Category:  blog.Category,
		Tags:      blog.Tags,
		CreatedAt: blog.Createdat,
		UpdatedAt: blog.Updatedat,
	})
}

func (cfg *apiConfig) getAllBlogsHandler(w http.ResponseWriter, r *http.Request) {
	blogs, err := cfg.BlogsDBQueries.GetAllBlogs(r.Context())
	if err != nil {
		log.Printf("failed to get blogs: %s", err)
		respondWithJSON(w, http.StatusInternalServerError, errResp{Error: err.Error()})
		return
	}

	var jsonBlogs []blogResp
	for _, blog := range blogs {
		jsonBlogs = append(jsonBlogs, blogResp{
			ID:        int(blog.ID),
			Title:     blog.Title,
			Category:  blog.Category,
			Tags:      blog.Tags,
			CreatedAt: blog.Createdat,
			UpdatedAt: blog.Updatedat,
		})
	}
	respondWithJSON(w, http.StatusOK, jsonBlogs)
}
