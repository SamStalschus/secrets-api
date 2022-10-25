package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/SamStalschus/secrets-api/cmd/secrets-api/middlewares"
	"github.com/SamStalschus/secrets-api/infra/log"
	"github.com/SamStalschus/secrets-api/internal"
)

var routes []route

func initializeRoutes() {
	routes = []route{
		newRoute("POST", "/users",
			middlewares.HandleRequestID(middlewares.RequestLogger(userController.SignUp, logger))),

		newRoute("GET", "/users/([^/]+)",
			middlewares.HandleRequestID(middlewares.RequestLogger(userController.GetUserByEmail, logger))),

		newRoute("PUT", "/users", updateUser),
		newRoute("GET", "/users/secrets/([^/]+)", getUserSecrets),
		newRoute("POST", "/secrets", createSecret),
		newRoute("GET", "/secret/([^/]+)", getSecret),
		newRoute("PUT", "/secret/([0-9]+)", updateSecret),
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

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "updateUser\n")
}

func getUserSecrets(w http.ResponseWriter, r *http.Request) {
	slug := internal.GetFields(r, "CtxKey", 0)
	id, _ := strconv.Atoi(internal.GetFields(r, "CtxKey", 1))
	fmt.Fprintf(w, "getUserSecrets %s %d\n", slug, id)
}

func createSecret(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "createSecret\n")
}

func getSecret(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "getSecret\n")
}

func updateSecret(w http.ResponseWriter, r *http.Request) {
	slug := internal.GetFields(r, "CtxKey", 0)
	fmt.Fprintf(w, "updateSecret %s\n", slug)
}
