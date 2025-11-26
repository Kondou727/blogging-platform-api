package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Kondou727/blogging-platform-api/internal/database"
	_ "modernc.org/sqlite"
)

type apiConfig struct {
	BlogsDB        *sql.DB
	BlogsDBQueries *database.Queries
}

func main() {
	mux := http.NewServeMux()
	db, err := LoadDB()
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	apiCfg := apiConfig{BlogsDB: db, BlogsDBQueries: dbQueries}

	mux.HandleFunc("POST /posts", apiCfg.createBlogHandler)
	mux.HandleFunc("PUT /posts/{id}", apiCfg.updateBlogHandler)
	mux.HandleFunc("DELETE /posts/{id}", apiCfg.deleteBlogHandler)

	server := http.Server{Handler: mux, Addr: ":8080"}
	log.Println("starting server...")
	log.Fatal(server.ListenAndServe())
}
