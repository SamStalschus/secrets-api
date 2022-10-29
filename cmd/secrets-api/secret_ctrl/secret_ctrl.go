package secret_ctrl

import (
	"encoding/json"
	"fmt"
	apiErr "github.com/SamStalschus/secrets-api/infra/errors"
	"github.com/SamStalschus/secrets-api/infra/log"
	"github.com/SamStalschus/secrets-api/internal"
	"github.com/SamStalschus/secrets-api/internal/secret"
	"io/ioutil"
	"net/http"
)

type Controller struct {
	secretService secret.IService
	logger        log.Provider
	apiErr        apiErr.Provider
}

func NewController(
	secretService secret.IService,
	logger log.Provider,
	apiErr apiErr.Provider,
) *Controller {
	return &Controller{
		secretService: secretService,
		logger:        logger,
		apiErr:        apiErr,
	}
}

func (c Controller) CreateSecret(w http.ResponseWriter, r *http.Request) {
	userID := internal.GetField(r.Context(), "user_id")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		c.logger.Error(r.Context(), "Bad Request - Fail to read body", log.Body{})
		return
	}

	var secret internal.Secret

	err = json.Unmarshal(body, &secret)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		c.logger.Error(r.Context(), "Bad Request - Fail to unmarshal body", log.Body{})
		return
	}

	errResponse := c.validateBody(&secret)
	if errResponse != nil {
		response, _ := json.Marshal(errResponse)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errResponse.ErrorStatus)
		w.Write(response)
		c.logger.Error(r.Context(),
			fmt.Sprintf("Error %v - Status %v", errResponse.Error, errResponse.ErrorStatus), log.Body{})
		return
	}

	errResponse = c.secretService.CreateSecret(r.Context(), &secret, userID)
	if errResponse != nil {
		response, _ := json.Marshal(errResponse)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errResponse.ErrorStatus)
		w.Write(response)
		c.logger.Error(r.Context(),
			fmt.Sprintf("Error %v - Status %v", errResponse.Error, errResponse.ErrorStatus), log.Body{})
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (c Controller) validateBody(secret *internal.Secret) (apiErr *apiErr.Message) {
	if secret.Value == "" || secret.Name == "" {
		apiErr = c.apiErr.BadRequest("Missing params", fmt.Errorf(""))
	}
	return apiErr
}
