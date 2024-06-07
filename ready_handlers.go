package main

import (
	"net/http"

	"github.com/Katalcha/rss-scraper/internal/helpers"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	type OkType struct {
		Status string `json:"status"`
	}
	helpers.RespondWithJSON(w, http.StatusOK, OkType{Status: "ok"})
}

func errHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
