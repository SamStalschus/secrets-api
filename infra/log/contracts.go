package log

import "context"

//go:generate mockgen -destination=./mocks.go -package=log -source=./contracts.go

type Provider interface {
	Debug(ctx context.Context, title string, valueMaps ...Body)
	Info(ctx context.Context, title string, valueMaps ...Body)
	Warn(ctx context.Context, title string, valueMaps ...Body)
	Error(ctx context.Context, title string, valueMaps ...Body)
	Fatal(ctx context.Context, title string, valueMaps ...Body)
}

type Body = map[string]interface{}
