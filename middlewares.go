package main

import (
	"net/http"

	"github.com/Katalcha/rss-scraper/internal/auth"
	"github.com/Katalcha/rss-scraper/internal/database"
	"github.com/Katalcha/rss-scraper/internal/helpers"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (a *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			helpers.RespondWithError(w, http.StatusUnauthorized, "could not find api key")
			return
		}

		user, err := a.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			helpers.RespondWithError(w, http.StatusNotFound, "could not get user")
			return
		}

		handler(w, r, user)
	}
}
