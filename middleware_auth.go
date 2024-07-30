package main

import (
	"fmt"
	"net/http"

	"github.com/wlucjan/rssagg/internal/auth"
	"github.com/wlucjan/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (api *apiConfig) middlewareAuth(next authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Error getting API key: %v", err))
			return
		}

		user, err := api.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting user: %v", err))
			return
		}

		next(w, r, user)
	}
}
