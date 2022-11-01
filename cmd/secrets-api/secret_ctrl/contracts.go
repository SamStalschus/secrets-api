package secret_ctrl

import "net/http"

type IController interface {
	CreateSecret(w http.ResponseWriter, r *http.Request)
	GetSecrets(w http.ResponseWriter, r *http.Request)
	GetSecret(w http.ResponseWriter, r *http.Request)
}
