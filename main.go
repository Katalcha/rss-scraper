package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Katalcha/rss-scraper/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DBURL")
	if dbURL == "" {
		log.Fatal("DBURL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	// debug flag parsing e.g: --debug
	dbg := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	if dbg != nil && *dbg {
		fmt.Println("Nothing happens here...")
	}

	apiConfig := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()

	// /v1/ readiness, error
	mux.HandleFunc("GET /v1/healthz", healthzHandler)
	mux.HandleFunc("GET /v1/err", errHandler)

	// /v1/users
	mux.HandleFunc("POST /v1/users", apiConfig.createUserHandler)
	mux.HandleFunc("GET /v1/users", apiConfig.middlewareAuth(apiConfig.getUserHandler))

	// /v1/feeds
	mux.HandleFunc("POST /v1/feeds", apiConfig.middlewareAuth(apiConfig.createFeedHandler))
	mux.HandleFunc("GET /v1/feeds", apiConfig.getFeedsHandler)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Ready for takeoff...\nServing on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
