package middlewares

import (
	"fmt"
	"net/http"

	"github.com/sstalschus/secrets-api/infra/cache"
	"github.com/sstalschus/secrets-api/infra/log"
)

func RateLimiter(h http.HandlerFunc, log log.Provider, cache cache.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cached int

		if cached = cache.GetInt(r.Context(), r.RemoteAddr); cached > 20 {
			log.Error(r.Context(), fmt.Sprintf("To many requests of %s", r.RemoteAddr))

			w.WriteHeader(http.StatusTooManyRequests)

			return
		}

		increment(&cached)

		cache.SetInt(r.Context(), r.RemoteAddr, cached, 5)

		h.ServeHTTP(w, r)
	}
}

func increment(c *int) {
	*c++
}
