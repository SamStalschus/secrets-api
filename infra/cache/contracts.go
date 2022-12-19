package cache

import "context"

//go:generate mockgen -destination=./mocks.go -package=cache -source=./contracts.go

type Provider interface {
	GetInt(ctx context.Context, key string) int
	SetInt(ctx context.Context, key string, value, ttl int)
	GetMap(ctx context.Context, key string) map[string]string
	SetMap(ctx context.Context, key string, value map[string]string, ttl int)
}
