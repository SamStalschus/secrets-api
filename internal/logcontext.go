package internal

import (
	"context"

	"github.com/sstalschus/secrets-api/infra/maps"
)

type logCtxKey struct{}

type ctxBody = map[string]interface{}

func CtxWithValues(ctx context.Context, values ctxBody) context.Context {
	m, _ := ctx.Value(logCtxKey{}).(ctxBody)
	return context.WithValue(ctx, logCtxKey{}, mergeMaps(m, values))
}

// GetCtxValues extracts the ctxBody currently stored on the input ctx.
func GetCtxValues(ctx context.Context) ctxBody {
	m, _ := ctx.Value(logCtxKey{}).(ctxBody)
	if m == nil {
		m = ctxBody{}
	}
	m["request_id"] = GetRequestIDFromContext(ctx)
	return m
}

func mergeMaps(bodies ...ctxBody) ctxBody {
	body := ctxBody{}
	maps.Merge(&body, bodies...)
	return body
}

func GetFields(ctx context.Context, key string, index int) string {
	keys := GetCtxValues(ctx)[key].([]string)
	return keys[index]
}

func GetField(ctx context.Context, key string) string {
	return GetCtxValues(ctx)[key].(string)
}
