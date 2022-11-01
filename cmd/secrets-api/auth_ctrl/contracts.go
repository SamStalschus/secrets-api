package auth_ctrl

import "net/http"

type IController interface {
	Authenticate(w http.ResponseWriter, r *http.Request)
}
