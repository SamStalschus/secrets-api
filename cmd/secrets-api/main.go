package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/cors"
	"github.com/sstalschus/secrets-api/cmd/secrets-api/auth_ctrl"
	"github.com/sstalschus/secrets-api/cmd/secrets-api/secret_ctrl"
	"github.com/sstalschus/secrets-api/infra/cache"
	"github.com/sstalschus/secrets-api/infra/cache/memorycache"
	"github.com/sstalschus/secrets-api/infra/mongodb/secret_repo"
	authService "github.com/sstalschus/secrets-api/internal/auth"
	"github.com/sstalschus/secrets-api/internal/secret"

	"github.com/sstalschus/secrets-api/infra/hash"

	"github.com/sstalschus/secrets-api/cmd/secrets-api/user_ctrl"
	"github.com/sstalschus/secrets-api/infra/env"
	apierr "github.com/sstalschus/secrets-api/infra/errors"
	"github.com/sstalschus/secrets-api/infra/log"
	"github.com/sstalschus/secrets-api/infra/log/jsonlogs"
	"github.com/sstalschus/secrets-api/infra/mongodb"
	"github.com/sstalschus/secrets-api/infra/mongodb/user_repo"
	"github.com/sstalschus/secrets-api/internal"
	"github.com/sstalschus/secrets-api/internal/user"
)

var (
	userController   *user_ctrl.Controller
	authController   *auth_ctrl.Controller
	secretController *secret_ctrl.Controller

	logger       log.Provider
	apiErrors    apierr.Provider
	authProvider hash.Provider
	cacheClient  cache.Provider
)

func main() {
	port := env.GetString("PORT", "8080")
	logLevel := env.GetString("LOG_LEVEL", "INFO")
	databaseURI := env.GetString("DATABASE_URI", "")
	secretKey := env.GetString("AUTH_SECRET_KEY", "")
	cacheTTL := env.GetInt("CACHE_TTL", 24)
	cachePurge := env.GetInt("CACHE_PURGE", 5)

	logger = jsonlogs.New(logLevel, internal.GetCtxValues)
	cacheClient = memorycache.New(time.Duration(cacheTTL)*time.Hour, time.Duration(cachePurge)*time.Minute)
	apiErrors = apierr.New()
	authProvider = hash.NewClient(secretKey)

	db, ctx := mongodb.GetConnection(logger, databaseURI)
	defer db.Disconnect(ctx)

	mongoRepository := mongodb.NewRepository(db)

	userRepository := user_repo.NewRepository(&mongoRepository)
	userService := user.NewService(logger, &userRepository, apiErrors, authProvider, cacheClient)
	userController = user_ctrl.NewController(userService, logger, apiErrors)

	authService := authService.NewService(apiErrors, authProvider, cacheClient, logger, &userService)
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
	cors := cors.AllowAll()

	handler := http.HandlerFunc(Server)
	return http.ListenAndServe(":"+port, cors.Handler(handler))
}
