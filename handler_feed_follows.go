package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/wlucjan/rssagg/internal/database"
)

func (api *apiConfig) handlerFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedID, err := uuid.Parse(chi.URLParam(r, "feed_id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing feed id: %v", err))
		return
	}

	follow, err := api.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feedID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error following feed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedFollowToAPIFeedFollow(follow))
}

func (api *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	follows, err := api.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting feed follows: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToAPIFeedFollows(follows))
}

func (api *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedID, err := uuid.Parse(chi.URLParam(r, "feed_follow_id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing feed follow id: %v", err))
		return
	}

	err = api.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error deleting feed follow: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
