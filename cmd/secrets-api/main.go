package main

import (
	"fmt"
	"net/http"
	"secrets-api/infra/bcrypt"

	"secrets-api/cmd/secrets-api/user_ctrl"
	"secrets-api/domain"
	"secrets-api/domain/user"
	"secrets-api/infra/env"
	apierr "secrets-api/infra/errors"
	"secrets-api/infra/log"
	"secrets-api/infra/log/jsonlogs"
	"secrets-api/infra/mongodb"
	"secrets-api/infra/mongodb/user_repo"
)

var (
	userController *user_ctrl.Controller
	logger         log.Provider
	apiErrors      apierr.Provider
)

func main() {
	port := env.GetString("PORT", "8080")
	logLevel := env.GetString("LOG_LEVEL", "INFO")
	databaseURI := env.GetString("DATABASE_URI", "")

	logger = jsonlogs.New(logLevel, domain.GetCtxValues)
	apiErrors := apierr.New()
	bcryptClient := bcrypt.NewClient()

	db, ctx := mongodb.GetConnection(logger, databaseURI)
	defer db.Disconnect(ctx)

	mongoRepository := mongodb.NewRepository(db)

	userRepository := user_repo.NewRepository(&mongoRepository)
	userService := user.NewService(logger, &userRepository, apiErrors, bcryptClient)
	userController = user_ctrl.NewController(userService, logger, apiErrors)

	logger.Info(ctx, fmt.Sprintf("Listening on port %s", port), log.Body{})
	if err := run(port); err != nil {
		logger.Fatal(ctx, fmt.Sprintf("Error to start server on port: %s - Erro: %s ", port, err), log.Body{})
	}
}

func run(port string) error {
	handler := http.HandlerFunc(Server)
	return http.ListenAndServe(":"+port, handler)
}
