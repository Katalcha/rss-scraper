package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Katalcha/rss-scraper/internal/database"
	"github.com/Katalcha/rss-scraper/internal/helpers"
	"github.com/google/uuid"
)

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

func (a *apiConfig) getUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	helpers.RespondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
