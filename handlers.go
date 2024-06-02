package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Katalcha/rss-scraper/internal/auth"
	"github.com/Katalcha/rss-scraper/internal/database"
	"github.com/Katalcha/rss-scraper/internal/helpers"
	"github.com/google/uuid"
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

func (a *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type requestParameters struct {
		Name string
	}

	decoder := json.NewDecoder(r.Body)
	params := requestParameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "could not decode parameters")
		return
	}
	if params.Name == "" {
		helpers.RespondWithError(w, http.StatusBadRequest, "malformed request body")
		return
	}

	user, err := a.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "could not create user")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (a *apiConfig) getUserByAPIKeyHandler(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		helpers.RespondWithError(w, http.StatusUnauthorized, "could not find api key")
		return
	}

	user, err := a.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		helpers.RespondWithError(w, http.StatusNotFound, "could not find user")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
