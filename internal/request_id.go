package internal

import (
	"context"

	"github.com/google/uuid"
)

// RequestIDKey needs to be a string, since its used on Fiber `.Locals()`
const RequestIDKey string = "request_id"

func GenerateRequestID() string {
	return uuid.NewString()
}

func GetRequestIDFromContext(ctx context.Context) string {
	body, _ := ctx.Value(logCtxKey{}).(ctxBody)
	if body == nil || body[RequestIDKey] == nil {
		return ""
	}
	return body[RequestIDKey].(string)
}
