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
		newRoute("GET", "/ping", ping),
		newRoute("POST", "/users",
			middlewares.HandleRequestID(
				middlewares.RequestLogger(userController.SignUp, logger))),

		newRoute("GET", "/users",
			middlewares.HandleRequestID(
				middlewares.RequestLogger(
					middlewares.EnsureAuth(
						userController.GetUser, authProvider, logger), logger))),

		newRoute("POST", "/token",
			middlewares.HandleRequestID(
				middlewares.RequestLogger(authController.Authenticate, logger))),

		newRoute("POST", "/secrets",
			middlewares.HandleRequestID(
				middlewares.RequestLogger(
					middlewares.EnsureAuth(
						secretController.CreateSecret, authProvider, logger), logger))),

		newRoute("GET", "/secrets",
			middlewares.HandleRequestID(
				middlewares.RequestLogger(
					middlewares.EnsureAuth(
						secretController.GetSecrets, authProvider, logger), logger))),

		newRoute("GET", "/secrets/([^/]+)",
			middlewares.HandleRequestID(
				middlewares.RequestLogger(
					middlewares.EnsureAuth(
						secretController.GetSecret, authProvider, logger), logger))),

		//newRoute("PUT", "/secret/([0-9]+)", updateSecret),
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
