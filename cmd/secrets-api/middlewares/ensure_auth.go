package middlewares

import (
	"fmt"
	"github.com/sstalschus/secrets-api/infra/cache"
	"net/http"

	"github.com/sstalschus/secrets-api/infra/hash"
	"github.com/sstalschus/secrets-api/infra/log"
	"github.com/sstalschus/secrets-api/internal"
)

func EnsureAuth(h http.HandlerFunc, auth hash.Provider, logger log.Provider, cache cache.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userID, err := auth.ValidateJwt(token)

		if userID != "" && err == nil {
			cached := cache.GetMap(r.Context(), userID)

			if cached["USER_STATUS"] == internal.BlockedStatus || cached["USER_STATUS"] == internal.CancelledStatus {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("Your user has been blocked or deleted. Please contact our support support@email.com"))
				return
			}

			ctx := internal.CtxWithValues(r.Context(), log.Body{
				"user_id": userID,
			})

			h.ServeHTTP(w, r.WithContext(ctx))
		} else {
			logger.Info(r.Context(), fmt.Sprintf("Error in hash of user %s Error %v", userID, err), log.Body{})
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
}
