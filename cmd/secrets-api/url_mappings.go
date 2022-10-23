package main

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var routes = []route{
	newRoute("POST", "/users", createUser),
	newRoute("PUT", "/users", updateUser),
	newRoute("GET", "/users/([0-9]+)", getUser),
	newRoute("GET", "/users/secrets/([^/]+)", getUserSecrets),
	newRoute("POST", "/secrets", createSecret),
	newRoute("GET", "/secret/([^/]+)", getSecret),
	newRoute("PUT", "/secret/([0-9]+)", updateSecret),
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
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

type ctxKey struct{}

func getField(r *http.Request, index int) string {
	fields := r.Context().Value(ctxKey{}).([]string)
	return fields[index]
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "createUser\n")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "updateUser\n")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	slug := getField(r, 0)
	id, _ := strconv.Atoi(getField(r, 0))
	fmt.Fprintf(w, "getUser %s %d\n", slug, id)
}

func getUserSecrets(w http.ResponseWriter, r *http.Request) {
	slug := getField(r, 0)
	id, _ := strconv.Atoi(getField(r, 1))
	fmt.Fprintf(w, "getUserSecrets %s %d\n", slug, id)
}

func createSecret(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "createSecret\n")
}

func getSecret(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "getSecret\n")
}

func updateSecret(w http.ResponseWriter, r *http.Request) {
	slug := getField(r, 0)
	fmt.Fprintf(w, "updateSecret %s\n", slug)
}
