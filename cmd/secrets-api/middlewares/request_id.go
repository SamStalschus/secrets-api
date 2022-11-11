package middlewares

import (
	"net/http"

	"github.com/sstalschus/secrets-api/infra/log"
	"github.com/sstalschus/secrets-api/internal"
)

func HandleRequestID(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := internal.GetField(r.Context(), internal.RequestIDKey)
		if requestID == "" {
			requestID = internal.GenerateRequestID()
		}

		ctx := internal.CtxWithValues(r.Context(), log.Body{
			internal.RequestIDKey: requestID,
		})

		h.ServeHTTP(w, r.WithContext(ctx))
	}
}
