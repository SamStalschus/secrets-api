package auth_ctrl

import (
	"encoding/json"
	"fmt"
	apiErr "github.com/SamStalschus/secrets-api/infra/errors"
	"github.com/SamStalschus/secrets-api/infra/log"
	"github.com/SamStalschus/secrets-api/internal"
	"github.com/SamStalschus/secrets-api/internal/auth"
	"io/ioutil"
	"net/http"
)

type Controller struct {
	authService auth.IService
	logger      log.Provider
	apiErr      apiErr.Provider
}

func NewController(
	authService auth.IService,
	logger log.Provider,
	apiErr apiErr.Provider,
) *Controller {
	return &Controller{
		authService: authService,
		logger:      logger,
		apiErr:      apiErr,
	}
}

func (c Controller) Authenticate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		c.logger.Error(r.Context(), "Bad Request - Fail to read body", log.Body{})
		return
	}

	var authUser internal.AuthUser

	err = json.Unmarshal(body, &authUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		c.logger.Error(r.Context(), "Bad Request - Fail to unmarshal body", log.Body{})
		return
	}

	errResponse := c.validateBody(&authUser)
	if errResponse != nil {
		response, _ := json.Marshal(errResponse)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errResponse.ErrorStatus)
		w.Write(response)
		c.logger.Error(r.Context(),
			fmt.Sprintf("Error %v - Status %v", errResponse.Error, errResponse.ErrorStatus), log.Body{})
		return
	}

	token, errResponse := c.authService.GenToken(r.Context(), &authUser, c.readUserIP(r))
	if errResponse != nil {
		response, _ := json.Marshal(errResponse)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errResponse.ErrorStatus)
		w.Write(response)
		c.logger.Error(r.Context(),
			fmt.Sprintf("Error %v - Status %v", errResponse.Error, errResponse.ErrorStatus), log.Body{})
		return
	}

	response, _ := json.Marshal(token)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

	return
}

func (c Controller) validateBody(user *internal.AuthUser) (apiErr *apiErr.Message) {
	if user.Email == "" || user.Password == "" {
		apiErr = c.apiErr.BadRequest("Missing params", fmt.Errorf("missing params"))
	}
	return apiErr
}

func (c Controller) readUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
