package middlewares

import (
	"net/http"

	"github.com/SamStalschus/secrets-api/domain"
	"github.com/SamStalschus/secrets-api/infra/log"
)

func HandleRequestID(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := domain.GetField(r.Context(), domain.RequestIDKey)
		if requestID == "" {
			requestID = domain.GenerateRequestID()
		}

		ctx := domain.CtxWithValues(r.Context(), log.Body{
			domain.RequestIDKey: requestID,
		})

		h.ServeHTTP(w, r.WithContext(ctx))
	}
}
