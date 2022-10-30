package main

import (
	"fmt"
	"github.com/SamStalschus/secrets-api/cmd/secrets-api/auth_ctrl"
	"github.com/SamStalschus/secrets-api/cmd/secrets-api/secret_ctrl"
	"github.com/SamStalschus/secrets-api/infra/mongodb/secret_repo"
	authService "github.com/SamStalschus/secrets-api/internal/auth"
	"github.com/SamStalschus/secrets-api/internal/secret"
	"net/http"

	"github.com/SamStalschus/secrets-api/infra/auth"

	"github.com/SamStalschus/secrets-api/cmd/secrets-api/user_ctrl"
	"github.com/SamStalschus/secrets-api/infra/env"
	apierr "github.com/SamStalschus/secrets-api/infra/errors"
	"github.com/SamStalschus/secrets-api/infra/log"
	"github.com/SamStalschus/secrets-api/infra/log/jsonlogs"
	"github.com/SamStalschus/secrets-api/infra/mongodb"
	"github.com/SamStalschus/secrets-api/infra/mongodb/user_repo"
	"github.com/SamStalschus/secrets-api/internal"
	"github.com/SamStalschus/secrets-api/internal/user"
)

var (
	userController   *user_ctrl.Controller
	authController   *auth_ctrl.Controller
	secretController *secret_ctrl.Controller
	logger           log.Provider
	apiErrors        apierr.Provider
	authProvider     auth.Provider
)

func main() {
	port := env.GetString("PORT", "8080")
	logLevel := env.GetString("LOG_LEVEL", "INFO")
	databaseURI := env.GetString("DATABASE_URI", "")
	secretKey := env.GetString("AUTH_SECRET_KEY", "")

	logger = jsonlogs.New(logLevel, internal.GetCtxValues)
	apiErrors = apierr.New()
	authProvider = auth.NewClient(secretKey)

	db, ctx := mongodb.GetConnection(logger, databaseURI)
	defer db.Disconnect(ctx)

	mongoRepository := mongodb.NewRepository(db)

	userRepository := user_repo.NewRepository(&mongoRepository)
	userService := user.NewService(logger, &userRepository, apiErrors, authProvider)
	userController = user_ctrl.NewController(userService, logger, apiErrors)

	authService := authService.NewService(apiErrors, authProvider, logger, &userRepository)
	authController = auth_ctrl.NewController(authService, logger, apiErrors)

	secretRepository := secret_repo.NewRepository(&mongoRepository)
	secretService := secret.NewService(logger, secretRepository, apiErrors, authProvider)
	secretController = secret_ctrl.NewController(secretService, logger, apiErrors)

	logger.Info(ctx, fmt.Sprintf("Listening on port %s", port), log.Body{})
	if err := run(port); err != nil {
		logger.Fatal(ctx, fmt.Sprintf("Error to start server on port: %s - Erro: %s ", port, err), log.Body{})
	}
}

func run(port string) error {
	handler := http.HandlerFunc(Server)
	return http.ListenAndServe(":"+port, handler)
}
