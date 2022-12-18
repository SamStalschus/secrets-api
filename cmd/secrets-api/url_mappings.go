package main

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/sstalschus/secrets-api/cmd/secrets-api/middlewares"
	"github.com/sstalschus/secrets-api/infra/log"
	"github.com/sstalschus/secrets-api/internal"
)

var routes []route

func initializeRoutes() {
	routes = []route{
		newRoute("GET", "/ping",
			middlewares.RateLimiter(ping, logger, cacheClient)),

		newRoute("POST", "/users",
			middlewares.RateLimiter(
				middlewares.HandleRequestID(
					middlewares.RequestLogger(userController.SignUp, logger)),
				logger, cacheClient)),

		newRoute("GET", "/users",
			middlewares.RateLimiter(
				middlewares.HandleRequestID(
					middlewares.RequestLogger(
						middlewares.EnsureAuth(
							userController.GetUser, authProvider, logger), logger)),
				logger, cacheClient)),

		newRoute("POST", "/token",
			middlewares.RateLimiter(
				middlewares.HandleRequestID(
					middlewares.RequestLogger(authController.Authenticate, logger)),
				logger, cacheClient)),

		newRoute("POST", "/secrets",
			middlewares.RateLimiter(
				middlewares.HandleRequestID(
					middlewares.RequestLogger(
						middlewares.EnsureAuth(
							secretController.CreateSecret, authProvider, logger), logger)),
				logger, cacheClient)),

		newRoute("GET", "/secrets",
			middlewares.RateLimiter(
				middlewares.HandleRequestID(
					middlewares.RequestLogger(
						middlewares.EnsureAuth(
							secretController.GetSecrets, authProvider, logger), logger)),
				logger, cacheClient)),

		newRoute("GET", "/secrets/([^/]+)",
			middlewares.RateLimiter(
				middlewares.HandleRequestID(
					middlewares.RequestLogger(
						middlewares.EnsureAuth(
							secretController.GetSecret, authProvider, logger), logger)),
				logger, cacheClient)),
	}
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func Server(w http.ResponseWriter, r *http.Request) {
	initializeRoutes()
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := internal.CtxWithValues(r.Context(), log.Body{
				"CtxKey": matches[1:],
			})
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong"))
}
