package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Katalcha/rss-scraper/internal/database"
	"github.com/Katalcha/rss-scraper/internal/helpers"
	"github.com/google/uuid"
)

func (a *apiConfig) getFeedFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := a.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "could not get feed follow")
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (a *apiConfig) createFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type responseParameters struct {
		FeedID uuid.UUID
	}
	decoder := json.NewDecoder(r.Body)
	params := responseParameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "could not decode parameters")
		return
	}

	feedFollow, err := a.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "could not create feed follow")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))
}

func (a *apiConfig) deleteFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := r.PathValue("feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "invalid feed follow ID")
		return
	}

	err = a.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFollowID,
	})
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "could not delete feed follow")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, struct{}{})
}
