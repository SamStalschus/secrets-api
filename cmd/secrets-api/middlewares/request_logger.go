package middlewares

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sstalschus/secrets-api/infra/log"
)

func RequestLogger(h http.HandlerFunc, logger log.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		h.ServeHTTP(w, r)

		body, _ := ioutil.ReadAll(r.Body)

		logger.Info(r.Context(), "request-completed", log.Body{
			"route":        r.Method + " " + r.URL.Path,
			"request_body": body,
			"duration_ms":  time.Since(startTime).Milliseconds(),
		})
	}
}
