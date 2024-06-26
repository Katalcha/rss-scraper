package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Katalcha/rss-scraper/internal/database"
	"github.com/Katalcha/rss-scraper/internal/helpers"
	"github.com/google/uuid"
)

func (a *apiConfig) createFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type requestParameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := requestParameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "could not decode parameters")
		return
	}

	feed, err := a.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "could not create feed")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
}

func (a *apiConfig) getFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := a.DB.GetFeeds(r.Context())
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "could not get feeds")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}
