package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv(ENV_PORT)

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    HOST + PORT,
		Handler: mux,
	}

	log.Printf("Ready for takeoff...\nServing from . on port %s\n", PORT)
	log.Fatal(server.ListenAndServe())
}
