package middlewares

import (
	"fmt"
	"github.com/sstalschus/secrets-api/infra/hash"
	"github.com/sstalschus/secrets-api/infra/log"
	"github.com/sstalschus/secrets-api/internal"
	"net/http"
)

func EnsureAuth(h http.HandlerFunc, auth hash.Provider, logger log.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userID, err := auth.ValidateJwt(token)

		if userID != "" && err == nil {
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
