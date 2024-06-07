package main

import (
	"net/http"
	"strconv"

	"github.com/Katalcha/rss-scraper/internal/database"
	"github.com/Katalcha/rss-scraper/internal/helpers"
)

func (a *apiConfig) getPostsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifiedLimit
	}

	posts, err := a.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "could not get posts for user")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
