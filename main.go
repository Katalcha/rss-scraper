package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv(ENV_PORT)

	// debug flag parsing e.g: --debug
	dbg := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	if dbg != nil && *dbg {
		fmt.Println("Nothing happens here...")
	}

	mux := http.NewServeMux()

	// readiness
	mux.HandleFunc(GET+HEALTHZ, healthzHandler)

	server := &http.Server{
		Addr:    HOST + PORT,
		Handler: mux,
	}

	log.Printf("Ready for takeoff...\nServing from . on port %s\n", PORT)
	log.Fatal(server.ListenAndServe())
}
