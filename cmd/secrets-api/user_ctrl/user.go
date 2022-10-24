package user_ctrl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"secrets-api/domain"
	"secrets-api/domain/user"
	apiErr "secrets-api/infra/errors"
	"secrets-api/infra/log"
	"strconv"
)

type Controller struct {
	usersService user.Service
	logger       log.Provider
	apiErr       apiErr.Provider
}

func NewController(
	userService user.Service,
	logger log.Provider,
	apiErr apiErr.Provider,
) Controller {
	return Controller{
		usersService: userService,
		logger:       logger,
		apiErr:       apiErr,
	}
}

func (c Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		c.logger.Error(ctx, "Bad Request - Fail to read body", log.Body{})
		return
	}

	var user domain.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		c.logger.Error(ctx, "Bad Request - Fail to unmarshal body", log.Body{})
		return
	}

	errResponse := c.validateBody(&user)
	if errResponse != nil {
		response, _ := json.Marshal(errResponse)
		w.WriteHeader(errResponse.ErrorStatus)
		w.Write(response)
		c.logger.Error(ctx, fmt.Sprintf("Bad Request - Error %v", errResponse.Error), log.Body{})
		return
	}

	errResponse = c.usersService.CreateUser(r.Context(), &user)
	if errResponse != nil {
		response, _ := json.Marshal(errResponse)
		w.WriteHeader(errResponse.ErrorStatus)
		w.Write(response)
		c.logger.Error(ctx, fmt.Sprintf("Bad Request - Error %v", errResponse.Error), log.Body{})
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (c Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(domain.GetFields(r, "CtxKey", 0))
	userText := c.usersService.GetUser(r.Context(), id)

	w.Write([]byte(userText))
}

func (c Controller) validateBody(user *domain.User) (apiErr *apiErr.Message) {
	if user.Email == "" || user.Name == "" || user.Password == "" {
		apiErr = c.apiErr.BadRequest("Missing params", fmt.Errorf(""))
	}
	return apiErr
}
