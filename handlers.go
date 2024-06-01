package main

import (
	"net/http"

	"github.com/Katalcha/rss-scraper/internal/helpers"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	type OkType struct {
		Status string `json:"status"`
	}

	payload := OkType{Status: "ok"}
	helpers.RespondWithJSON(w, http.StatusOK, payload)
}

func errHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Internal Server Error"
	helpers.RespondWithError(w, http.StatusInternalServerError, msg)
}
